package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/middlewares"
	"net/http"
)

func SetupRouter(r *gin.Engine, db *gorm.DB) {
	var (
		//repo
		authRepository
		userRepository userrepository.UserRepository = userrepository.NewUserRepository(db)

		//service
		userService userservice.UserService = userservice.NewUserService(userRepository, roleRepository, departmentRepository, levelRepository, workUnitRepository)

		//handler
		userHandler userhandler.UserHandler = userhandler.NewUserHandler(userService)

		//middleware
		superUserMiddleware = middlewares.SuperUserMiddleware()
		adminMiddleware     = middlewares.AdminMiddleware()
		userMiddleware      = middlewares.UserMiddleware()
	)

	authRoutes := apiRoutes.Group("auth")
	{
		authRoutes.POST("/login", authHandler.Login)
		authRoutes.POST("/forgot-password", authHandler.ForgotPassword)
		authRoutes.POST("/reset-password", authHandler.ResetPassword)
		authRoutes.GET("/token", authMiddleware, authHandler.GenerateNoExpireToken)
	}

	userRoutes := apiRoutes.Group("users", authMiddleware)
	{
		userRoutes.POST("", superUserMiddleware, userHandler.CreateUser)
		userRoutes.GET("", superUserMiddleware, userHandler.ListUserPagination)
		userRoutes.PUT(":id/update-user", superUserMiddleware, userHandler.UpdateUser)
		userRoutes.PUT(":id/soft-delete", superUserMiddleware, userHandler.SoftDeleteUserByID)
		userRoutes.POST("/bulk-soft-delete", superUserMiddleware, userHandler.BulkSoftDeleteUser)
		userRoutes.GET("/:id", superUserMiddleware, userHandler.GetUserByID)
		userRoutes.GET("/all", userHandler.ListUser)
		userRoutes.GET("/pic-work-unit/:work_unit_id", userHandler.GetUserPICByWorkUnitID)
	}

}
