package common

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	w "github.com/edwingeng/wuid/redis/wuid"
	goRedis "github.com/go-redis/redis"
	redigo "github.com/gomodule/redigo/redis"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"regexp"
	"strings"
	"time"
	myRedis "university_circles/service/user_service/databases/redis"
	"university_circles/service/user_service/utils/errcode"
	"university_circles/service/user_service/utils/logger"
	"unsafe"
)

const (
	// redis key
	RANDINCRIDCHEPREFIX       = "university_circles:user:rand:id:"
	LOGINTOKENCACHEPREFIX     = "university_circles:login:token:"
	LOGINPHONECODECACHEPREFIX = "university_circles:login:phone:code:"

	ACCESSKEYID     = "LTAIuZkI49CABYVV"
	ACCESSKEYSECRET = "PN5eRkdrtaGoL0oshDoQWmL6LnE7cb"

	// es index
	HOMEPUBLISHMSGUSER    = "home_publish_msg_user"
	HOMEPUBLISHMSG        = "home_publish_msg"
	HOMEPUBLISHMSGCOMMENT = "home_publish_msg_comment"

	// user index
	// 用户关注关系列表
	USERFOLLOWList = "user_follow_list:"
	// 用户被关注关系列表
	USERFOLLOWINGList = "user_following_list:"
	// 用户session
	STUDENTLOGINSESSIONPREKEY = "university_circles:student:login:session:"

	// im
	IMAUTHTOKENCACHEPREFIX = "access_token_"

	HOMEFILE             = 1
	USERAVATARFILE       = 2
	USERREGISTERFILE     = 3
	UPLOADFILEPATHPREFIX = "ee11cbb19052e40b07aac0ca060c23ee/96ab4e163f4ee03aaa4d1051aa51d204/%d/%s"
	OSSFILEURLPREFIX     = "https://university-circles-static-resources.oss-cn-shenzhen.aliyuncs.com/"

	KC_RAND_KIND_NUM   = 0 // 纯数字
	KC_RAND_KIND_LOWER = 1 // 小写字母
	KC_RAND_KIND_UPPER = 2 // 大写字母
	KC_RAND_KIND_ALL   = 3 // 数字、大小写字母
)

//json数据解析
type Message struct {
	Message   string
	RequestId string
	BizId     string
	Code      string
}

func SendVerifyCode(phone, code string) (int64, error) {
	client, err := sdk.NewClientWithAccessKey("cn-hangzhou", ACCESSKEYID, ACCESSKEYSECRET)
	if err != nil {
		logger.Logger.Warn("sms New Client With AccessKey failed", zap.String("phone", phone), zap.String("code", code))
		return -1, err
	}

	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Scheme = "http" // https | http
	request.Domain = "dysmsapi.aliyuncs.com"
	request.Version = "2017-05-25"
	request.ApiName = "SendSms"
	request.QueryParams["RegionId"] = "cn-hangzhou"
	request.QueryParams["PhoneNumbers"] = phone
	request.QueryParams["SignName"] = "火仁"
	request.QueryParams["TemplateCode"] = "SMS_171746086"
	request.QueryParams["TemplateParam"] = "{\"code\":\"" + code + "\"}"

	response, err := client.ProcessCommonRequest(request)
	if err != nil {
		logger.Logger.Warn("send Verify Code failed", zap.String("response", response.GetHttpContentString()), zap.String("phone", phone), zap.String("code", code))
		return -1, err
	}
	fmt.Println("发短信body", response.GetHttpContentString())
	var message Message //阿里云返回的json信息对应的类
	//记得判断错误信息
	err = json.Unmarshal(response.GetHttpContentBytes(), &message)
	if err != nil {
		logger.Logger.Warn("send Verify Code json Unmarshal failed", zap.Any("message", "message"), zap.String("phone", phone), zap.String("code", code))
		return -1, err
	}

	fmt.Println(err, "isv.BUSINESS_LIMIT_CONTROL" == message.Code, message.Message != "OK")

	if "isv.BUSINESS_LIMIT_CONTROL" == message.Code {
		logger.Logger.Warn("send Verify Code failed", zap.Any("message", "message"), zap.String("phone", phone), zap.String("code", code))
		return errcode.ErrBusinessDayLimitControlFailed.Code, err
	}

	if message.Message != "OK" {
		logger.Logger.Warn("send Verify Code failed", zap.Any("message", "message"), zap.String("phone", phone), zap.String("code", code))
		return errcode.ErrSendVerifyCodeFailed.Code, err
	}

	return 0, nil
}

func GenValidateCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}

func GenRandomRecommendMsgId(n int) string {
	var letterBytes = "0123456789abcdefghijklmnopqrstuvwxyz"
	const (
		letterIdxBits = 6                    // 6 bits to represent a letter index
		letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
		letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	)

	b := make([]byte, n)
	var src = rand.NewSource(time.Now().UnixNano())
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}

func GetRandomID() uint64 {
	newClient := func() (goRedis.Cmdable, bool, error) {
		return goRedis.NewClient(&goRedis.Options{
			Addr:     "127.0.0.1:6379",
			Password: "123456",
		}), true, nil
	}

	// Setup
	g := w.NewWUID("default", nil, w.WithSection(15))
	_ = g.LoadH28FromRedis(newClient, "university_circles_msg_orderId")
	// Generate
	return g.Next()
}

// 生成密码
func GeneratePassword(password string, cost int) (string, error) {
	if cost < 1 {
		cost = 1
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// 校验密码
func ValidatePassword(hashPwd, pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(pwd))
	if err != nil {
		return false
	}
	return true
}

// 校验邮箱
func VerifyEmailFormat(email string) bool {
	//pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`

	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

// 校验手机
func VerifyMobileFormat(mobileNum string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"

	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}

func Md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// 随机字符串
func KRand(size int, kind int) string {
	iKind, kinds, result := kind, [][]int{{10, 48}, {26, 97}, {26, 65}}, make([]byte, size)
	isAll := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if isAll { // random ikind
			iKind = rand.Intn(3)
		}
		scope, base := kinds[iKind][0], kinds[iKind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return string(result)
}

// incr随机字符串
func RandIncrId(randType string) (randId int64, err error) {
	rd := myRedis.DefaultRedisPool.Get()
	defer rd.Close()

	cacheKey := RANDINCRIDCHEPREFIX + randType

	// 设置缓存
	randId, err = redigo.Int64(rd.Do("Incr", cacheKey))
	if err != nil {
		logger.Logger.Warn("generate rand id by redis incr failed", zap.Any("randType", randType), zap.Any("key", cacheKey), zap.Error(err))
		return randId, err
	}

	return
}



