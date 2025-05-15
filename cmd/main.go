package main

import (
	"diprec_api/cmd/application"
	"diprec_api/internal/config"
	"diprec_api/internal/infrastructure/db/postgres"
	"diprec_api/internal/pkg/logger"
	"diprec_api/internal/service"
	"fmt"
	"log"

	"go.uber.org/zap"

	user_repo "diprec_api/internal/repository/user"
	user_handler "diprec_api/internal/transport/http/user"
	user_usecase "diprec_api/internal/usecase/user"

	course_repo "diprec_api/internal/repository/course"
	course_handler "diprec_api/internal/transport/http/course"
	course_usecase "diprec_api/internal/usecase/course"
)

func main() {
	cfg := config.MustLoad()

	fmt.Printf("Server starting on %s:%d\n", cfg.Server.Host, cfg.Server.Port)

	db, err := postgres.NewPostgresDB(postgres.Config{
		Host:            cfg.DB.Host,
		Port:            cfg.DB.Port,
		User:            cfg.DB.User,
		Password:        cfg.DB.Password,
		DBName:          cfg.DB.DBName,
		SSLMode:         cfg.DB.SSLMode,
		MaxIdleConns:    cfg.DB.MaxIdleConns,
		MaxOpenConns:    cfg.DB.MaxOpenConns,
		ConnMaxLifetime: cfg.DB.ConnMaxLifetime,
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := postgres.AutoMigrate(db); err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}

	custom_logger, err := logger.New(cfg.Logging)

	auth_service := service.NewAuthService(&service.JWTConfig{
		SecretKey:     cfg.Auth.JWTSecret,
		AccessExpiry:  cfg.Auth.AccessTokenExpire,
		RefreshExpiry: cfg.Auth.RefreshTokenExpire,
	})

	if err != nil {
		custom_logger.Error("grpc detector client failed", zap.Error(err))
	}

	ur := user_repo.NewUserRepository(db)
	uc := user_usecase.NewUserUseCase(ur, auth_service, custom_logger)
	uh := user_handler.NewUserHandler(uc, custom_logger)

	cr := course_repo.NewCourseRepository(db)
	cu := course_usecase.NewCourseUseCase(cr, custom_logger)
	ch := course_handler.NewCourseHandler(cu, custom_logger)

	app := application.NewApplication(cfg, custom_logger, db)

	app.Start(uh, ch, auth_service)
}
