package test

import (
	"fmt"
	w "github.com/edwingeng/wuid/redis/wuid"
	"github.com/go-redis/redis"
	"math/rand"
	"regexp"
	"strings"
	"testing"
	"time"
	"unsafe"
)

func TestGetRandomString(t *testing.T) {

	newClient := func() (redis.Cmdable, bool, error) {
		return redis.NewClient(&redis.Options{
			Addr:     "127.0.0.1:6379",
			Password: "123456",
		}), true, nil
	}

	// Setup
	g := w.NewWUID("default", nil, w.WithSection(15))
	_ = g.LoadH28FromRedis(newClient, "university_circles_uid")
	// Generate
	for i := 0; i < 10; i++ {
		//time.Sleep(1*time.Second)
		fmt.Println(fmt.Sprintf("%d%s", g.Next(), time.Now().Format("20160102150405")))
	}
}

func TestRandomStr(t *testing.T) {
	var letterBytes = "0123456789abcdefghijklmnopqrstuvwxyz"
	const (
		letterIdxBits = 6                    // 6 bits to represent a letter index
		letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
		letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	)
	n := 32
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

	fmt.Println(*(*string)(unsafe.Pointer(&b)))
}

func TestUrlsInText(t *testing.T) {
	flysnowRegexp := regexp.MustCompile(`(https?|ftp|file)://[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]`)
	params := flysnowRegexp.FindAllString("测试发送啦啦啦啦啦啦啦啦阿拉http://www.baidu.com哈哈哈哈哈哈。不知道怎样http://www.hao123.com嗯嗯额", 30)
	fmt.Println(params)
	fmt.Println(len(params))
}

func BenchmarkAddStringWithJoin(b *testing.B) {
	hello := "hello"
	world := "world"
	var str string
	for i := 0; i < b.N; i++ {
		str = strings.Join([]string{hello, world}, ",")
	}

	fmt.Println(str)
}
