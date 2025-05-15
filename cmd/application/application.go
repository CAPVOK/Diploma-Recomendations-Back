package application

import (
	_ "diprec_api/docs"
	"diprec_api/internal/config"
	"diprec_api/internal/service"
	course_handler "diprec_api/internal/transport/http/course"
	"diprec_api/internal/transport/http/middleware"
	question_handler "diprec_api/internal/transport/http/question"
	test_handler "diprec_api/internal/transport/http/test"
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
func (a *Application) Start(
	user_handler *user_handler.UserHandler,
	course_handler *course_handler.CourseHandler,
	test_handler *test_handler.TestHandler,
	question_handler *question_handler.QuestionHandler,
	auth_service *service.AuthService,
) {
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
				course.GET("", course_handler.Get)
				course.POST("", course_handler.Create)
				course.GET("/:id", course_handler.GetByID)
				course.DELETE("/:id", course_handler.Delete)
				course.PUT("/:id", course_handler.Update)
			}

			test := protected.Group("/test")
			{
				test.GET("/:id", test_handler.GetByID)
				test.POST("/:id", test_handler.Create)
				test.DELETE("/:id", test_handler.Delete)
				test.PUT("/:id", test_handler.Update)
			}

			question := protected.Group("/question")
			{
				question.POST("/:id", question_handler.Create)
				question.GET("/:id", question_handler.GetByID)
				question.DELETE("/:id", question_handler.Delete)
				question.PUT("/:id", question_handler.Update)
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
