package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"slices"
)

func SuperUserMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := ctx.MustGet("user").(*jwt.MapClaims)
		roleNames := (*user)["role_names"].([]any)

		isSuperUser := slices.ContainsFunc(roleNames, func(roleName any) bool {
			roleNameStr := roleName.(string)
			return roleNameStr == constants.ROLE_NAMES["SUPER_USER"]
		})

		if !isSuperUser {
			ctx.Abort()
			utils.ApiResponse(ctx, http.StatusUnauthorized, "only superuser can access this endpoint", nil, nil)
			return
		}

		ctx.Next()
	}
}
