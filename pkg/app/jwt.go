package app

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/xielizyh/goprj-blog_service/global"
	"github.com/xielizyh/goprj-blog_service/pkg/util"
)

type Claims struct {
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	jwt.StandardClaims
}

// GetJWTSecret 获取JWT Secret
func GetJWTSecret() []byte {
	// SigningMethodHMAC的Sign方法判断了key是否为字节数组类型 参考https://www.jianshu.com/p/664e6d83ea69
	return []byte(global.JWTSetting.Secret)
}

// GenerateToken 生成JWT Token
func GenerateToken(appKey, appSecret string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(global.JWTSetting.Expire)
	claims := Claims{
		AppKey:    util.EncodeMD5(appKey),
		AppSecret: util.EncodeMD5(appSecret),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    global.JWTSetting.Issuer,
		},
	}
	// 根据Claims创建Token实例
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 生成签名字符串
	token, err := tokenClaims.SignedString(GetJWTSecret())

	return token, err
}

// ParseToken 解析和校验Token
func ParseToken(token string) (*Claims, error) {
	// 解析鉴权的声明
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
