package course

import (
	"duolingo_api/internal/domain"
	"duolingo_api/internal/usecase/course"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type CourseHandler struct {
	cu     course.ICourseUsecase
	logger *zap.Logger
}

func NewCourseHandler(cu course.ICourseUsecase, logger *zap.Logger) *CourseHandler {
	return &CourseHandler{
		cu:     cu,
		logger: logger.Named("CourseHandler"),
	}
}

// Create godoc
// @Summary Создать курс
// @Tags Course
// @Security BearerAuth
// @Produce json
// @Param input body CreateCourseDTO true "Название и описание проекта"
// @Success 201 {object} CourseResponse
// @Router /course [post]
func (h *CourseHandler) Create(c *gin.Context) {
	var req CreateCourseDTO
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	course, err := h.cu.Create(c.Request.Context(), &domain.Course{
		Name:        req.Name,
		Description: req.Description,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := ToCourseResponse(course)
	c.JSON(http.StatusCreated, response)
}
