package authorization

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	t "api-gateway/internal/https/token"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {}

func (a *AuthMiddleware) MiddleWare() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		fmt.Println("Authorization header:", token)
		url := ctx.Request.URL.Path
		fmt.Println("Request URL:", url)

		if strings.Contains(url, "swagger") || url == "/api/v1/auth/login" {
			ctx.Next()
			return
		}

		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "Authorization header is missing",
			})
			return
		}

		if !strings.HasPrefix(token, "Bearer ") {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "Authorization token is missing Bearer prefix",
			})
			return
		}

		token = strings.TrimPrefix(token, "Bearer ")

		claims, err := t.ExtractClaim(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": err.Error(),
			})
			return
		}
		log.Println(claims)

		user_id, ok := claims["user_id"].(string)
		if !ok || user_id == "" {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "ID not found in token",
			})
			return
		}

		user_identifier, ok := claims["user_identifier"].(string)
		if !ok || user_id == "" {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "Identifier not found in token",
			})
			return
		}

		ctx.Set("user_id", user_id)
		ctx.Set("user_identifier", user_identifier)
		
		ctx.Next()
	}
}

