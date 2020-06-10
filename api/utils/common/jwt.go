package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte("W4SIGw7fw40Q7OkzSWaM2Obl#zfpxRGFDHTR&$%$%^#$GGFG5545gdgdf-90345")

// JWT 参考：http://www.ruanyifeng.com/blog/2018/07/json_web_token-tutorial.html
// 自定义 Payload（负载）
type Claims struct {
	Uid          string `json:"uid"`
	ScreenName   string `json:"screenName"`
	UniversityId int64  `json:"university_id"`
	LoginTime    int64  `json:loginTime`
	jwt.StandardClaims
}

// 创建token
func GenerateToken(uid, nickName string, UniversityId int64, expiration time.Duration) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(expiration)
	claims := Claims{
		uid,
		nickName,
		UniversityId,
		time.Now().Unix(),
		jwt.StandardClaims{
			Issuer:    "com.huoren",      // 签发人
			ExpiresAt: expireTime.Unix(), // 过期时间
			Subject:   uid,               // 使用人
			Audience:  "all",             // 接收对象
			NotBefore: nowTime.Unix(),    // 生效时间
			IssuedAt:  nowTime.Unix(),    //签发时间
			//Id:        nil,               // 唯一身份标识，主要用来作为一次性token,从而回避重放攻击。
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// 解析 token
func ParseToken(token string) (*Claims, error) {
	// 解析token
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, keyFunc)
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

// 返回key
func keyFunc(token *jwt.Token) (interface{}, error) {
	return jwtSecret, nil
}
