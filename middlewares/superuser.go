package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"main.go/constants"
	"main.go/utils"
	"net/http"
)

func SuperUserMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, exists := ctx.Get("user")
		if !exists {
			ctx.Abort()
			utils.ApiResponse(ctx, http.StatusUnauthorized, "user not found in context", nil, nil)
			return
		}

		claims, ok := user.(*jwt.MapClaims)
		if !ok {
			ctx.Abort()
			utils.ApiResponse(ctx, http.StatusUnauthorized, "invalid user claims", nil, nil)
			return
		}

		role, roleExists := (*claims)["role"].(string)
		if !roleExists {
			ctx.Abort()
			utils.ApiResponse(ctx, http.StatusUnauthorized, "user role not found in claims", nil, nil)
			return
		}

		if role != constants.ROLE["SUPER_USER"] {
			ctx.Abort()
			utils.ApiResponse(ctx, http.StatusUnauthorized, "only superuser can access this endpoint", nil, nil)
			return
		}

		ctx.Next()
	}
}
