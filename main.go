package main

import (
	"github.com/IrvanWijayaSardam/Bengkel/config"
	"github.com/IrvanWijayaSardam/Bengkel/controller"
	"github.com/IrvanWijayaSardam/Bengkel/middleware"
	"github.com/IrvanWijayaSardam/Bengkel/repository"
	"github.com/IrvanWijayaSardam/Bengkel/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                         = config.SetupDatabaseConnection()
	userRepository repository.UserRepository        = repository.NewUserRepository(db)
	trxRepository  repository.TransactionRepository = repository.NewTransactionRepository(db)
	jwtService     service.JWTService               = service.NewJWTService()
	trxService     service.TransactionService       = service.NewTransactionService(trxRepository)
	authService    service.AuthService              = service.NewAuthService(userRepository)
	authController controller.AuthController        = controller.NewAuthController(authService, jwtService)
	trxController  controller.TransactionContoller  = controller.NewTransactionController(trxService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	trxRoutes := r.Group("api/trx", middleware.AuthorizeJWT(jwtService))
	{
		trxRoutes.GET("/", trxController.All)
		trxRoutes.POST("/", trxController.Insert)
	}

	r.Run(":8001")
}
