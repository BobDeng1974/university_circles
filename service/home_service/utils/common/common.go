package common

import (
	"math/rand"
	"time"
	"unsafe"
)

const (
	// es index
	HOMEPUBLISHMSGUSER    = "home_publish_msg_user"
	HOMEPUBLISHMSG        = "home_publish_msg"
	HOMEPUBLISHMSGCOMMENT = "home_publish_msg_comment"

	// user index
	// 用户关注关系列表
	USERFOLLOWList = "user_follow_list:"
	// 用户被关注关系列表
	USERFOLLOWINGList = "user_following_list:"

	UPLOADFILEPATHPREFIX = "ee11cbb19052e40b07aac0ca060c23ee/96ab4e163f4ee03aaa4d1051aa51d204/%d/%s"
	OSSFILEURLPREFIX     = "https://university-circles-static-resources.oss-cn-shenzhen.aliyuncs.com/"
)

//json数据解析
type Message struct {
	Message   string
	RequestId string
	BizId     string
	Code      string
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
