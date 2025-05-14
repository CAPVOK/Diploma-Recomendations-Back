package application

import (
	_ "duolingo_api/docs"
	"duolingo_api/internal/config"
	"duolingo_api/internal/service"
	"duolingo_api/internal/transport/http/middleware"
	user_handler "duolingo_api/internal/transport/http/user"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Application struct {
	config *config.Config
	logger *zap.Logger
	db     *gorm.DB
	minio  *minio.Client
}

func NewApplication(config *config.Config, logger *zap.Logger, db *gorm.DB) *Application {
	return &Application{
		config: config,
		logger: logger,
		db:     db,
	}
}

// Start @title Thesis Backend API
// @version 1.0
// @description API для дипломного проекта
// @termsOfService http://swagger.io/terms/
// @host localhost:8080
// @BasePath /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func (a *Application) Start(user_handler *user_handler.UserHandler, auth_service *service.AuthService) {
	router := gin.Default()

	router.Use(middleware.CORSMiddleware())
	router.Use(gin.Recovery())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := router.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/register", user_handler.Register)
			auth.POST("/login", user_handler.Login)
			auth.POST("/refresh", user_handler.Refresh)
		}

		protected := v1.Group("")
		protected.Use(middleware.IsAuthenticated(auth_service, a.logger.Named("Auth Middleware")))
	}

	serverAddr := fmt.Sprintf("%s:%d", a.config.Server.Host, a.config.Server.Port)
	a.logger.Info("Starting server", zap.String("address", serverAddr))
	if err := router.Run(serverAddr); err != nil {
		a.logger.Fatal("Failed to start server", zap.Error(err))
		log.Fatal(err)
	}
}
