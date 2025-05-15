package question

import (
	"diprec_api/internal/domain"
	"diprec_api/internal/pkg/utils"
	"diprec_api/internal/usecase/question"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type QuestionHandler struct {
	qu     question.IQuestionUsecase
	logger *zap.Logger
}

func NewQuestionHandler(qu question.IQuestionUsecase, logger *zap.Logger) *QuestionHandler {
	return &QuestionHandler{
		qu:     qu,
		logger: logger.Named("QuestionHandler"),
	}
}

// Create godoc
// @Summary Создать вопрос
// @Tags Question
// @Security BearerAuth
// @Produce json
// @Param input body CreateQuestionDTO true "ДТО создания вопроса"
// @Param id path int true "ID теста"
// @Success 201 {object} domain.QuestionResponse
// @Error 400 {object} domain.Error
// @Error 400 {object} domain.Error
// @Error 401 {object} domain.Error
// @Error 500 {object} domain.Error
// @Router /question/{id} [post]
func (h *QuestionHandler) Create(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.logger.Warn("Validation error", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": domain.ErrInvalidRequestBody.Error()})
		return
	}

	var req CreateQuestionDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Warn("Validation error", zap.Error(err))
		c.JSON(http.StatusBadRequest, domain.Error{Message: domain.ErrInvalidRequestBody.Error()})
		return
	}

	question, err := h.qu.Create(c.Request.Context(), &domain.Question{
		Title:    req.Title,
		Type:     req.Type,
		Variants: utils.ParseMapToJSON(req.Variants),
		Answer:   utils.ParseMapToJSON(req.Answer),
	}, uint(id))
	if err != nil {
		h.logger.Warn("Create error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := question.ToQuestionResponse()
	c.JSON(http.StatusCreated, response)
}

// GetByID Get godoc
// @Summary Получить вопрос по ID
// @Tags Question
// @Security BearerAuth
// @Produce json
// @Param id path int true "ID вопроса"
// @Success 200 {object} domain.QuestionResponse
// @Failure 400 {object} domain.Error
// @Failure 401 {object} domain.Error
// @Failure 500 {object} domain.Error
// @Router /question/{id} [get]
func (h *QuestionHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.logger.Warn("Validation error", zap.Error(err))
		c.JSON(http.StatusBadRequest, domain.Error{Message: domain.ErrInvalidRequestBody.Error()})
		return
	}

	question, err := h.qu.GetByID(c.Request.Context(), uint(id))
	if err != nil {
		h.logger.Warn("GetByID error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := question.ToQuestionResponse()
	c.JSON(http.StatusOK, response)
}

// Update godoc
// @Summary Обновить вопрос
// @Tags Question
// @Security BearerAuth
// @Produce json
// @Param id path int true "ID вопроса"
// @Param input body UpdateQuestionDTO true "ДТО обновления вопроса"
// @Success 200 {object} domain.QuestionResponse
// @Success 400 {object} domain.Error
// @Success 400 {object} domain.Error
// @Success 500 {object} domain.Error
// @Router /question/{id} [put]
func (h *QuestionHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.logger.Warn("Validation error", zap.Error(err))
		c.JSON(http.StatusBadRequest, domain.Error{Message: domain.ErrInvalidRequestBody.Error()})
		return
	}

	var req UpdateQuestionDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Warn("Validation error", zap.Error(err))
		c.JSON(http.StatusBadRequest, domain.Error{Message: domain.ErrInvalidRequestBody.Error()})
		return
	}

	question, err := h.qu.Update(c.Request.Context(), &domain.Question{
		ID:       uint(id),
		Title:    req.Title,
		Type:     req.Type,
		Variants: utils.ParseMapToJSON(req.Variants),
		Answer:   utils.ParseMapToJSON(req.Answer),
	})
	if err != nil {
		h.logger.Warn("Update error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := question.ToQuestionResponse()
	c.JSON(http.StatusOK, response)
}

// Delete godoc
// @Summary Удалить тест
// @Tags Question
// @Security BearerAuth
// @Produce json
// @Param id path int true "ID вопроса"
// @Success 200
// @Failure 400 {object} domain.Error
// @Failure 401 {object} domain.Error
// @Failure 500 {object} domain.Error
// @Router /question/{id} [delete]
func (h *QuestionHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.logger.Warn("Validation error", zap.Error(err))
		c.JSON(http.StatusBadRequest, domain.Error{Message: domain.ErrInvalidRequestBody.Error()})
		return
	}

	err = h.qu.Delete(c.Request.Context(), uint(id))
	if err != nil {
		h.logger.Warn("Delete error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}
