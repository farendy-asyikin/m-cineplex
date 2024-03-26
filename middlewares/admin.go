package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"main.go/constants"
	"main.go/utils"
	"net/http"
	"slices"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := ctx.MustGet("user").(*jwt.MapClaims)
		roleNames := (*user)["role"].([]any)

		isSuperUser := slices.ContainsFunc(roleNames, func(roleName any) bool {
			roleNameStr := roleName.(string)
			return roleNameStr == constants.ROLE["ADMIN"]
		})

		if !isSuperUser {
			ctx.Abort()
			utils.ApiResponse(ctx, http.StatusUnauthorized, "you can login as admin to access this endpoint", nil, nil)
			return
		}

		ctx.Next()
	}
}
