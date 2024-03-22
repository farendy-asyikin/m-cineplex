package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"main.go/constants"
	"main.go/utils"
	"net/http"
	"slices"
)

func UserMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := ctx.MustGet("user").(*jwt.MapClaims)
		roleNames := (*user)["role_names"].([]any)

		isSuperUser := slices.ContainsFunc(roleNames, func(roleName any) bool {
			roleNameStr := roleName.(string)
			return roleNameStr == constants.ROLE_NAMES["USER"]
		})

		if !isSuperUser {
			ctx.Abort()
			utils.ApiResponse(ctx, http.StatusUnauthorized, "you can login as user to access this endpoint", nil, nil)
			return
		}

		ctx.Next()
	}
}
