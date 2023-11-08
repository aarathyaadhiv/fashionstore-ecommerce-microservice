package middleware


import (
	
	"fmt"
	"net/http"
	"strings"

	
	"github.com/aarathyaadhiv/fashionstore-api-gateway/pkg/utils/response"
	"github.com/gin-gonic/gin"

	"github.com/golang-jwt/jwt"
)

type CustomUserClaim struct {
	ID    uint
	Email string
	Role  string
	jwt.StandardClaims
}
func UserAuthorizationMiddleware(c *gin.Context) {
	s := c.Request.Header.Get("Authorization")

	var tokenString string
	if strings.HasPrefix(s, "Bearer") {
		tokenString = strings.TrimPrefix(s, "Bearer ")
	} else {
		tokenString = s
	}

	token, err := validateUserToken(tokenString)
	if err != nil || !token.Valid {
		errRes := response.Responses(http.StatusUnauthorized, "not authorised", nil, err.Error())
		c.JSON(http.StatusUnauthorized, errRes)
		c.Abort()
		return
	}
	claims,ok:=token.Claims.(*CustomUserClaim)
	
	if !ok{
		errRes := response.Responses(http.StatusUnauthorized, "not authorised", nil, fmt.Errorf("claim not retrieved").Error())
		c.JSON(http.StatusUnauthorized, errRes)
		c.Abort()
		return
	}
	
	id := claims.ID
	c.Set("userId", id)
	c.Next()
}

func validateUserToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString,&CustomUserClaim{} ,func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte("secret"), nil
	})

	return token, err
}