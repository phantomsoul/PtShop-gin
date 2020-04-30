package drivers

import (
	"encoding/json"
	"errors"
	"fmt"
	jwtLib "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strings"
	"time"
	"pt-gin/config"
	"pt-gin/dto"
	"pt-gin/modules/log"
)

type jwtAuthManager struct {
	secret string
	exp    time.Duration
	alg    string
}

func NewJwtAuthDriver() *jwtAuthManager {
	return &jwtAuthManager{
		secret: config.GetJwtConfig().SECRET,
		exp:    config.GetJwtConfig().EXP,
		alg:    config.GetJwtConfig().ALG,
	}
}

// 创建Token
func (jwtAuth *jwtAuthManager) CreateToken(userInfo map[string]interface{}) interface{} {

	token := jwtLib.New(jwtLib.GetSigningMethod(jwtAuth.alg))
	// Set some claims
	userStr, err := json.Marshal(userInfo)
	if err != nil {
		return nil
	}
	log.Sugar(string(userStr))
	token.Claims = jwtLib.MapClaims{
		"user": string(userStr),
		"exp":  time.Now().Add(jwtAuth.exp).Unix(),
	}
	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(jwtAuth.secret))
	if err != nil {
		return nil
	}

	return tokenString
}

// 验证Token
func (jwtAuth *jwtAuthManager) Check(c *gin.Context) bool {
	token := c.Request.Header.Get("zd-jwt")
	token = strings.Replace(token, "Bearer ", "", -1)
	if token == "" {
		return false
	}
	//var keyFun = func(token *jwtLib.Token) (interface{}, error) {
	//	b := []byte(jwtAuth.secret)
	//	return b, nil
	//}
	//authJwtToken, err := request.ParseFromRequest(c.Request, request.OAuth2Extractor, keyFun)
	authJwtToken, err := jwtLib.Parse(token, func(token *jwtLib.Token) (interface{}, error) {
		b := []byte(jwtAuth.secret)
		return b, nil
	})

	if err != nil {
		log.Error("jwt check error", zap.Error(err))
		return false
	}

	c.Set("jwt_auth_token", authJwtToken)

	// 解析payload
	if claims, ok := authJwtToken.Claims.(jwtLib.MapClaims); ok && authJwtToken.Valid {
		c.Set("auth_user", claims["user"].(string))
		return true
	} else {
		log.Error("decode jwt user claims fail")
		return false
	}
}

// 快捷获取Token中用户信息
func (jwtAuth *jwtAuthManager) User(c *gin.Context) dto.JwtUserInfo {
	jwtUserInfo := fmt.Sprintf("%s", c.MustGet("auth_user"))
	userInfo := dto.JwtUserInfo{}
	if jwtUserInfo != "" {
		if err := json.Unmarshal([]byte(jwtUserInfo), &userInfo); err != nil {
			panic(err)
		}
		return userInfo
	}

	var jwtToken *jwtLib.Token
	if jwtAuthToken, exist := c.Get("jwt_auth_token"); !exist {
		tokenStr := strings.Replace(c.Request.Header.Get("zd-jwt"), "Bearer ", "", -1)
		if tokenStr == "" {
			return dto.JwtUserInfo{}
		}
		var err error
		jwtToken, err = jwtLib.Parse(tokenStr, func(token *jwtLib.Token) (interface{}, error) {
			b := []byte(jwtAuth.secret)
			return b, nil
		})
		if err != nil {
			panic(err)
		}
	} else {
		jwtToken = jwtAuthToken.(*jwtLib.Token)
	}

	if claims, ok := jwtToken.Claims.(jwtLib.MapClaims); ok && jwtToken.Valid {
		if err := json.Unmarshal([]byte(claims["user"].(string)), &userInfo); err != nil {
			panic(err)
		}
		c.Set("auth_user", claims["user"].(string))
		return userInfo
	} else {
		panic(errors.New("decode jwt user claims fail"))
	}
}
