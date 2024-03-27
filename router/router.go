package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	authhandler "main.go/handlers/auth"
	bookinghandler "main.go/handlers/booking"
	filmhandler "main.go/handlers/film"
	locationhandler "main.go/handlers/location"
	seathandler "main.go/handlers/seat"
	userhandler "main.go/handlers/users"
	"main.go/middlewares"
	bookingrepository "main.go/repositories/booking"
	filmrepository "main.go/repositories/film"
	"main.go/repositories/is_booked"
	locationrepository "main.go/repositories/location"
	seatrepository "main.go/repositories/seat"
	userrepository "main.go/repositories/users"
	authservice "main.go/services/auth"
	bookingservice "main.go/services/booking"
	filmservice "main.go/services/film"
	locationservice "main.go/services/location"
	seatservice "main.go/services/seat"
	userservice "main.go/services/user"
)

func SetupRouter(r *gin.Engine, db *gorm.DB) {
	var (
		//repo
		authRepository     userrepository.UserRepository         = userrepository.NewUserRepository(db)
		userRepository     userrepository.UserRepository         = userrepository.NewUserRepository(db)
		filmRepository     filmrepository.FilmRepository         = filmrepository.NewFilmRepository(db)
		bookingRepository  bookingrepository.BookingRepository   = bookingrepository.NewBookingRepository(db)
		seatRepository     seatrepository.SeatRepository         = seatrepository.NewSeatRepository(db)
		locationRepository locationrepository.LocationRepository = locationrepository.NewLocationRepository(db)
		isBookedRepository is_booked.IsBookedRepository          = is_booked.NewIsBookedRepository(db)

		//service
		authService     authservice.AuthService         = authservice.NewAuthService(authRepository)
		userService     userservice.UserService         = userservice.NewUserService(userRepository)
		filmService     filmservice.FilmService         = filmservice.NewFilmService(filmRepository)
		bookingService  bookingservice.BookingService   = bookingservice.NewBookingService(bookingRepository, isBookedRepository)
		seatService     seatservice.SeatService         = seatservice.NewSeatService(seatRepository)
		locationService locationservice.LocationService = locationservice.NewLocationService(locationRepository)

		//handler
		authHandler     authhandler.AuthHandler         = authhandler.NewAuthHandler(authService)
		userHandler     userhandler.UserHandler         = userhandler.NewUserHandler(userService)
		filmHandler     filmhandler.FilmHandler         = filmhandler.NewFilmHandler(filmService)
		bookingHandler  bookinghandler.BookingHandler   = bookinghandler.NewBookingHandler(bookingService)
		seatHandler     seathandler.SeatHandler         = seathandler.NewSeatHandler(seatService)
		locationHandler locationhandler.LocationHandler = locationhandler.NewLocationHandler(locationService)

		//middleware
		authMiddleware = middlewares.AuthMiddleware()
		//adminMiddleware = middlewares.AdminMiddleware()
		//superuserMiddleware = middlewares.SuperUserMiddleware()
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
			userRoutes.PUT("/:id", authMiddleware, userHandler.UpdateUser)
			userRoutes.DELETE("/:id", authMiddleware, userHandler.DeleteUserByID)
			userRoutes.GET("/:id", authMiddleware, userHandler.GetUserByID)

			//superuser role
			userRoutes.PUT("/update-role/:id", userHandler.UpdateUserRoleBySuperuser)
			//userRoutes.GET("", superuserMiddleware, userHandler.ListUser)
		}

		filmRoutes := apiRoutes.Group("films")
		{
			filmRoutes.GET("/:id", filmHandler.GetFilmByID)
			filmRoutes.GET("", filmHandler.ListFilm)

			//admin role
			filmRoutes.POST("", filmHandler.CreateFilm)
			filmRoutes.PUT("/:id", filmHandler.UpdateFilm)
			filmRoutes.DELETE("/:id", filmHandler.DeleteFilmByID)
		}

		bookingRoutes := apiRoutes.Group("bookings")
		{
			bookingRoutes.POST("", bookingHandler.CreateBooking)
			bookingRoutes.GET("/:id", bookingHandler.GetBookingByID)
		}

		seatRoutes := apiRoutes.Group("seats")
		{
			seatRoutes.GET("", seatHandler.ListUnbookedSeatByFilmIDAndLocationID)
		}

		locationRoutes := apiRoutes.Group("locations")
		{
			locationRoutes.GET("", locationHandler.ListLocation)
		}
	}

}
