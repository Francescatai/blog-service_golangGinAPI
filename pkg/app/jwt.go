package app

import(
	"time"

	"github.com/dgrijalva/jwt-go"

	"go_gin_blog/global"
	"go_gin_blog/pkg/util"
)

type Claims struct {
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	jwt.StandardClaims //payload
}

func GetJWTSecret() []byte {
	return []byte(global.JWTSetting.Secret)
}

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

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) //根據Claims結構創建Token實例
	token, err := tokenClaims.SignedString(GetJWTSecret()) //根據傳入secret生成簽名字符串並返回標準Token
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) { //ParseWithClaims:用於解析鑑權的聲明，最終返回*token
		return GetJWTSecret(), nil
	})
    if err != nil {
	    return nil, err
	}
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid { //Valid:驗證基於時間的聲明ex.過期時間/簽發者/生效時間
			return claims, nil
		}
	}

	return nil, err
}