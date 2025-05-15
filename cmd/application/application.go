package application

import (
	_ "diprec_api/docs"
	"diprec_api/internal/config"
	"diprec_api/internal/service"
	course_handler "diprec_api/internal/transport/http/course"
	"diprec_api/internal/transport/http/middleware"
	user_handler "diprec_api/internal/transport/http/user"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Application struct {
	config *config.Config
	logger *zap.Logger
	db     *gorm.DB
}

func NewApplication(config *config.Config, logger *zap.Logger, db *gorm.DB) *Application {
	return &Application{
		config: config,
		logger: logger,
		db:     db,
	}
}

// Start @title Baumlingo Backend API
// @version 1.0
// @description API для дипломного проекта
// @termsOfService http://swagger.io/terms/
// @host localhost:8080
// @BasePath /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func (a *Application) Start(user_handler *user_handler.UserHandler, course_handler *course_handler.CourseHandler, auth_service *service.AuthService) {
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
		{
			course := protected.Group("/course")
			{
				course.GET("")
				course.POST("", course_handler.Create)
				course.GET("/:id")
				course.DELETE("/:id")
				course.PUT("/:id")
			}

			test := protected.Group("/test")
			{
				test.GET("")
				test.POST("")
				test.GET("/:id")
				test.DELETE("/:id")
				test.PUT("/:id")
			}

			question := protected.Group("/question")
			{
				question.GET("")
				question.POST("")
				question.GET("/:id")
				question.DELETE("/:id")
				question.PUT("/:id")
			}
		}
	}

	serverAddr := fmt.Sprintf("%s:%d", a.config.Server.Host, a.config.Server.Port)
	a.logger.Info("Starting server", zap.String("address", serverAddr))
	if err := router.Run(serverAddr); err != nil {
		a.logger.Fatal("Failed to start server", zap.Error(err))
		log.Fatal(err)
	}
}
