package test

import (
	"diprec_api/internal/domain"
	"diprec_api/internal/usecase/test"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type TestHandler struct {
	tu     test.ITestUsecase
	logger *zap.Logger
}

func NewTestHandler(tu test.ITestUsecase, logger *zap.Logger) *TestHandler {
	return &TestHandler{
		tu:     tu,
		logger: logger.Named("TestHandler"),
	}
}

// Create godoc
// @Summary Создать тест
// @Tags Test
// @Security BearerAuth
// @Produce json
// @Param input body CreateTestDTO true "Название, описание и дедлайн теста"
// @Param id path int true "ID курса"
// @Success 201 {object} domain.TestResponse
// @Error 400 {object} domain.Error
// @Error 400 {object} domain.Error
// @Error 401 {object} domain.Error
// @Error 500 {object} domain.Error
// @Router /test/{id} [post]
func (h *TestHandler) Create(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.logger.Warn("Validation error", zap.Error(err))
		c.JSON(http.StatusBadRequest, domain.Error{Message: domain.ErrInvalidRequestBody.Error()})
		return
	}

	var req CreateTestDTO
	if err := c.ShouldBind(&req); err != nil {
		h.logger.Warn("Validation error", zap.Error(err))
		c.JSON(http.StatusBadRequest, domain.Error{Message: domain.ErrInvalidRequestBody.Error()})
		return
	}

	test, err := h.tu.Create(c.Request.Context(), &domain.Test{
		Name:        req.Name,
		Description: req.Description,
		Deadline:    req.Deadline,
	}, uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.Error{Message: err.Error()})
		return
	}

	response := test.ToTestResponse()
	c.JSON(http.StatusCreated, response)
}

// GetByID Get godoc
// @Summary Получить тест по ID
// @Tags Test
// @Security BearerAuth
// @Produce json
// @Param id path int true "ID теста"
// @Success 200 {object} domain.TestResponseWithQuestions
// @Failure 400 {object} domain.Error
// @Failure 401 {object} domain.Error
// @Failure 500 {object} domain.Error
// @Router /test/{id} [get]
func (h *TestHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.logger.Warn("Validation error", zap.Error(err))
		c.JSON(http.StatusBadRequest, domain.Error{Message: domain.ErrInvalidRequestBody.Error()})
		return
	}

	test, err := h.tu.GetByID(c.Request.Context(), uint(id))
	if err != nil {
		h.logger.Warn("Internal error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, domain.Error{Message: err.Error()})
		return
	}

	response := test.ToTestResponseWithQuestions()
	c.JSON(http.StatusOK, response)
}

// Update godoc
// @Summary Обновить тест
// @Tags Test
// @Security BearerAuth
// @Produce json
// @Param id path int true "ID теста"
// @Param input body UpdateTestDTO true "Название, описание и дедлайн курса"
// @Success 200 {object} domain.TestResponse
// @Success 400 {object} domain.Error
// @Success 400 {object} domain.Error
// @Success 500 {object} domain.Error
// @Router /test/{id} [put]
func (h *TestHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.logger.Warn("Validation error", zap.Error(err))
		c.JSON(http.StatusBadRequest, domain.Error{Message: domain.ErrInvalidRequestBody.Error()})
		return
	}

	var req UpdateTestDTO
	if err := c.ShouldBind(&req); err != nil {
		h.logger.Warn("Validation error", zap.Error(err))
		c.JSON(http.StatusBadRequest, domain.Error{Message: domain.ErrInvalidRequestBody.Error()})
		return
	}

	test, err := h.tu.Update(c.Request.Context(), &domain.Test{
		ID:          uint(id),
		Name:        req.Name,
		Description: req.Description,
		Deadline:    req.Deadline,
	})
	if err != nil {
		h.logger.Warn("Internal error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, domain.Error{Message: err.Error()})
		return
	}

	response := test.ToTestResponse()
	c.JSON(http.StatusOK, response)
}

// Delete godoc
// @Summary Удалить тест
// @Tags Test
// @Security BearerAuth
// @Produce json
// @Param id path int true "ID теста"
// @Success 200
// @Failure 400 {object} domain.Error
// @Failure 401 {object} domain.Error
// @Failure 500 {object} domain.Error
// @Router /test/{id} [delete]
func (h *TestHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.logger.Warn("Validation error", zap.Error(err))
		c.JSON(http.StatusBadRequest, domain.Error{Message: domain.ErrInvalidRequestBody.Error()})
		return
	}

	if err := h.tu.Delete(c.Request.Context(), uint(id)); err != nil {
		h.logger.Warn("Internal error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, domain.Error{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}

// AttachQuestion godoc
// @Summary Прикрепить вопрос к тесту
// @Tags Test
// @Security BearerAuth
// @Produce json
// @Param id path int true "ID теста"
// @Param input body AttachQuestionDTO true "ID вопроса"
// @Success 200
// @Failure 400 {object} domain.Error
// @Failure 401 {object} domain.Error
// @Failure 500 {object} domain.Error
// @Router /test/{id}/question [post]
func (h *TestHandler) AttachQuestion(c *gin.Context) {
	idStr := c.Param("id")
	testID, err := strconv.Atoi(idStr)
	if err != nil {
		h.logger.Warn("Validation error", zap.Error(err))
		c.JSON(http.StatusBadRequest, domain.Error{Message: domain.ErrInvalidRequestBody.Error()})
		return
	}

	var req AttachQuestionDTO
	if err := c.ShouldBind(&req); err != nil {
		h.logger.Warn("Validation error", zap.Error(err))
		c.JSON(http.StatusBadRequest, domain.Error{Message: domain.ErrInvalidRequestBody.Error()})
		return
	}

	err = h.tu.AttachQuestion(c.Request.Context(), uint(testID), uint(req.QuestionID))
	if err != nil {
		h.logger.Warn("Internal error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, domain.Error{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}
