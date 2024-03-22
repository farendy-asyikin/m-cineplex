package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	authhandler "main.go/handlers/auth"
	filmhandler "main.go/handlers/film"
	userhandler "main.go/handlers/users"
	"main.go/middlewares"
	bookingrepository "main.go/repositories/booking"
	filmrepository "main.go/repositories/film"
	userrepository "main.go/repositories/users"
	authservice "main.go/services/auth"
	filmservice "main.go/services/film"
	userservice "main.go/services/user"
)

func SetupRouter(r *gin.Engine, db *gorm.DB) {
	var (
		//repo
		authRepository    userrepository.UserRepository       = userrepository.NewUserRepository(db)
		userRepository    userrepository.UserRepository       = userrepository.NewUserRepository(db)
		filmRepository    filmrepository.FilmRepository       = filmrepository.NewFilmRepository(db)
		bookingRepository bookingrepository.BookingRepository = bookingrepository.NewBookingRepository(db)

		//service
		authService    authservice.AuthService       = authservice.NewAuthService(authRepository)
		userService    userservice.UserService       = userservice.NewUserService(userRepository)
		filmService    filmservice.FilmService       = filmservice.NewFilmService(filmRepository)
		bookingService bookingservice.BookingService = bookingservice.NewBookingServixe(bookingRepository)

		//handler
		authHandler    authhandler.AuthHandler       = authhandler.NewAuthHandler(authService)
		userHandler    userhandler.UserHandler       = userhandler.NewUserHandler(userService)
		filmHandler    filmhandler.FilmHandler       = filmhandler.NewFilmHandler(filmService)
		bookingHandler bookinghandler.BookingHandler = bookinghandler.NewFilmHandler(bookingService)

		//middleware
		authMiddleware  = middlewares.AuthMiddleware()
		adminMiddleware = middlewares.AdminMiddleware()
	)

	apiRoutes := r.Group("api")
	{
		authRoutes := apiRoutes.Group("auth")
		{
			authRoutes.POST("/login", authHandler.Login)
		}

		userRoutes := apiRoutes.Group("users")
		{
			userRoutes.POST("", userHandler.CreateUser)
			userRoutes.PUT(":id/update-user", authMiddleware, userHandler.UpdateUser)
			userRoutes.PUT(":id/delete", authMiddleware, userHandler.DeleteUserByID)
			userRoutes.GET("/:id", authMiddleware, userHandler.GetUserByID)
		}

		filmRoutes := apiRoutes.Group("films")
		{
			filmRoutes.POST("", adminMiddleware, filmHandler.CreateFilm)
			filmRoutes.GET("", authMiddleware, filmHandler.ListFilm)
			filmRoutes.GET("/:id", authMiddleware, filmHandler.DetailFilm)
			filmRoutes.PUT("/:id", adminMiddleware, filmHandler.UpdateFilm)
			filmRoutes.DELETE("/:id", adminMiddleware, filmHandler.DeleteFilm)
		}

		bookingRoutes := apiRoutes.Group("booking")
		{
			bookingRoutes.POST("", authMiddleware, bookingHandler.CreateBooking)
			bookingRoutes.GET("", adminMiddleware, bookingHandler.ListBooking)
			bookingRoutes.GET("", authMiddleware, bookingHandler.DetailBooking)
			bookingRoutes.GET("", authMiddleware, bookingHandler.ListByUserIdBooking)
			bookingRoutes.PUT("", adminMiddleware, bookingHandler.UpdateBooking)
			bookingRoutes.DELETE("", adminMiddleware, bookingHandler.DeleteBooking)
		}
	}

}
