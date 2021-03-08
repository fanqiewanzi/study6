package util

import (
	"github.com/dgrijalva/jwt-go"
	"study6/manage/pkg/setting"
	"time"
)

//获取JwtSecret
var jwtSecret = []byte(setting.JwtSecret)

//自定义载荷结构，包含了一个标准的JWT载荷
type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

//创建JWt
func CreatJwt(username, password string) (string, error) {
	//通过现在时间得到token过期时间
	nowTime := time.Now()
	expTime := nowTime.Add(30 * time.Minute)

	//初始化自定义的载荷
	claim := Claims{username, password, jwt.StandardClaims{ExpiresAt: expTime.Unix(), Issuer: "bbb"}}

	//初始化tokenClaim数据结构
	tokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	//生成签名字符串，再用于获取完整、已签名的token
	token, err := tokenClaim.SignedString(jwtSecret)
	return token, err
}

//解析token字符串
func ParseToken(token string) (*Claims, error) {
	//用密匙解析出token声明
	tokenClaim, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	//token声明解析完成后再次解析出自定义的声明并检查token声明是否有效，有效则返回自定义声明
	if tokenClaim != nil {
		if claims, ok := tokenClaim.Claims.(*Claims); ok && tokenClaim.Valid {
			return claims, nil
		}
	}

	//解析失败
	return nil, err
}
