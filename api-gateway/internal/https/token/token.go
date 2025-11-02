package token 

import (
	"errors"
	"log"
	"time"
	
	config "api-gateway/internal/config"

	"github.com/form3tech-oss/jwt-go"
	pb "github.com/xadichamakhkamova/YouthUnionContracts/genproto/userpb"
)

type Tokens struct {
	Token     string
	ExpiresIn string
}

var cfg config.Config
var tokenKey = cfg.TokenKey

func GenereteJWTToken(user *pb.GetUserByIdentifierResponse) *Tokens {

	refreshToken := jwt.New(jwt.SigningMethodHS256)

	rftclaims := refreshToken.Claims.(jwt.MapClaims)
	rftclaims["user_id"] = user.Id
	rftclaims["user_identifier"] = user.Identifier
	rftclaims["iat"] = time.Now().Unix()
	rftclaims["exp"] = time.Now().Add(24 * time.Hour).Unix()
	refresh, err := refreshToken.SignedString([]byte(tokenKey))
	if err != nil {
		log.Fatal("error while genereting refresh token : ", err)
	}

	expirationUnix := rftclaims["exp"].(int64)
	expirationTime := time.Unix(expirationUnix, 0).Format("2006-01-02 15:04:05")

	return &Tokens{
		Token:     refresh,
		ExpiresIn: expirationTime,
	}
}

func ExtractClaim(tokenStr string) (jwt.MapClaims, error) {
	var (
		token *jwt.Token
		err   error
	)

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenKey), nil
	}
	token, err = jwt.Parse(tokenStr, keyFunc)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
