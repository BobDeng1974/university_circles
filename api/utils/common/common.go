package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gomodule/redigo/redis"
	"github.com/satori/go.uuid"
	"go.uber.org/zap"
	"math/rand"
	"strconv"
	"time"

	"university_circles/api/utils/logger"
	myRedis "university_circles/service/user_service/databases/redis"
)

const (
	// redis key
	LOGINTOKENCACHEPREFIX     = "university_circles:login:token:"
	LOGINPHONECODECACHEPREFIX = "university_circles:login:phone:code:"

	// im
	IMAUTHTOKENCACHEPREFIX = "access_token_"
	USERIMAUTHTOKENCACHEKEY = "user_im_access_token"

	// 阿里云oss
	bucket       = "university-circles-static-resources"
	endPoint     = "oss-cn-shenzhen-internal.aliyuncs.com"
	accessKeyId  = "LTAIuZkI49CABYVV"
	accessSecret = "PN5eRkdrtaGoL0oshDoQWmL6LnE7cb"

	KC_RAND_KIND_NUM   = 0 // 纯数字
	KC_RAND_KIND_LOWER = 1 // 小写字母
	KC_RAND_KIND_UPPER = 2 // 大写字母
	KC_RAND_KIND_ALL   = 3 // 数字、大小写字母

	UPLOADFILEPATHPREFIX = "ee11cbb19052e40b07aac0ca060c23ee/96ab4e163f4ee03aaa4d1051aa51d204/%d/%s"
)

func SetLoginToken(uid, nickName, deviceType string, universityId int64) (token string, err error) {
	rd := myRedis.DefaultRedisPool.Get()
	defer rd.Close()

	cacheKey := LOGINTOKENCACHEPREFIX + deviceType + "_" + uid

	expiration := time.Hour * 24 * 30
	expireTime := time.Now().Unix() + 259200
	fmt.Println(expireTime)
	token, err = GenerateToken(uid, nickName, universityId, expiration)
	// 设置缓存
	_, err = rd.Do("Set", cacheKey, token, "EX", expireTime)
	if err != nil {
		logger.Logger.Warn("set user token into redis failed", zap.Any("uid", uid), zap.Any("key", cacheKey), zap.Error(err))
		return "", err
	}
	return
}

func GetLoginToken(uid string) (token string, err error) {
	rd := myRedis.DefaultRedisPool.Get()
	defer rd.Close()

	cacheKey := LOGINTOKENCACHEPREFIX + uid

	// 设置缓存
	token, err = redis.String(rd.Do("Get", cacheKey))
	if err != nil {
		logger.Logger.Warn("get user token from redis failed", zap.Any("uid", uid), zap.Any("key", cacheKey), zap.Error(err))
		return "", err
	}
	return
}

func RefreshLoginToken(uid, nickName string, universityId int64) (token string, err error) {
	rd := myRedis.DefaultRedisPool.Get()
	defer rd.Close()

	cacheKey := LOGINTOKENCACHEPREFIX + uid

	expiration := time.Hour * 24 * 3
	expireTime := time.Now().Unix() + 259200
	token, err = GenerateToken(uid, nickName, universityId, expiration)
	// 设置缓存
	_, err = rd.Do("Set", cacheKey, token, "EX", expireTime)
	if err != nil {
		logger.Logger.Warn("reset user token into redis failed", zap.Any("uid", uid), zap.Any("key", cacheKey), zap.Error(err))
		return "", err
	}
	return
}

func GenIMToken(appid, uid int64) string {
	u2 := uuid.NewV4()

	return Md5V(u2.String() + strconv.FormatInt(appid, 10) + strconv.FormatInt(uid, 10))
}

func SaveIMToken(token string, user map[string]interface{}) (err error) {
	rd := myRedis.DefaultRedisPool.Get()
	defer rd.Close()

	cacheKey := IMAUTHTOKENCACHEPREFIX + token
	_, err = rd.Do("hMSet", redis.Args{}.Add(cacheKey).AddFlat(user)...)
	if err != nil {
		logger.Logger.Warn("im auth user token hmset failed", zap.Any("user", user), zap.Any("cacheKey", cacheKey), zap.Error(err))
		return
	}

	return
}

func SetUserIMToken(token string, uid int64) (err error) {
	rd := myRedis.DefaultRedisPool.Get()
	defer rd.Close()

	cacheKey := USERIMAUTHTOKENCACHEKEY
	_, err = rd.Do("hSet", cacheKey, uid, token)
	if err != nil {
		logger.Logger.Warn("im auth user token hset failed", zap.Any("cacheKey", cacheKey), zap.Error(err))
		return
	}
	return
}

func GetUserIMToken(uid int64) (token string, err error) {
	rd := myRedis.DefaultRedisPool.Get()
	defer rd.Close()

	cacheKey := USERIMAUTHTOKENCACHEKEY
	token, err = redis.String(rd.Do("hGet", cacheKey, uid))
	if err != nil {
		if err.Error() != "redigo: nil returned" {
			logger.Logger.Warn("im auth user token hget failed", zap.Any("cacheKey", cacheKey), zap.Error(err))
			return
		}
	}

	return token, nil
}

func DelIMToken(token string) (err error) {
	rd := myRedis.DefaultRedisPool.Get()
	defer rd.Close()

	cacheKey := IMAUTHTOKENCACHEPREFIX + token
	_, err = rd.Do("Del", cacheKey)
	if err != nil {
		logger.Logger.Warn("im auth user token del failed", zap.Any("cacheKey", cacheKey), zap.Error(err))
		return
	}
	return
}

func Md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func UploadFileToOSS(localFilename, file string) error {
	// 创建OSSClient实例
	client, err := oss.New(endPoint, accessKeyId, accessSecret)
	if err != nil {
		logger.Logger.Warn("upload file to oss failed", zap.String("filename", file), zap.Error(err))
		return err
	}

	bucket, err := client.Bucket(bucket)
	if err != nil {
		logger.Logger.Warn("upload file to oss failed", zap.String("filename", file), zap.Error(err))
		return err
	}

	err = bucket.PutObjectFromFile(file, localFilename)
	if err != nil {
		logger.Logger.Warn("upload file to oss failed", zap.String("filename", file), zap.Error(err))
		return err
	}

	return nil
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

