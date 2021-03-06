// tool/jwt.go
package tool

import (
	"time"

	"github.com/iris-contrib/middleware/jwt"
)

// 获取JWT字符串和20分钟期约
func GetJWTString(name string, id int64) (string, error) {
	token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		// 根据需求，可以存一些必要的数据
		"userName": name,
		"userId":   id,
		// 签发人
		"iss": "iris",
		// 签发时间
		"iat": time.Now().Unix(),
		// 设定过期时间，设置一天过期
		"exp": time.Now().Add(365 * 24 * 60 * time.Minute * time.Duration(1)).Unix(),
	})

	// 使用设置的秘钥，签名生成jwt字符串
	tokenString, err := token.SignedString([]byte("My Secret"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
