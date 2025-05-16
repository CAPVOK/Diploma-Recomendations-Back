package domain

import "errors"

type Error struct {
	Message string
}

var (
	ErrUserNotFound        = errors.New("Пользователь не найден")
	ErrInvalidCredentials  = errors.New("Неверные имя пользователя или пароль")
	ErrUserExists          = errors.New("Пользователь с таким именем уже существует")
	ErrInvalidRefreshToken = errors.New("Рефреш токен невалиден")
	ErrInvalidTokenType    = errors.New("Неверный тип токена")
	ErrUnauthorized        = errors.New("Неавторизован")
	ErrInvalidRequestBody  = errors.New("Данные неверны")
	ErrCourseNotFound      = errors.New("Курс не найден")
	ErrTestNotFound        = errors.New("Тест не найден")
	ErrQuestionNotFound    = errors.New("Вопрос не найден")
	ErrInvalidRole         = errors.New("Данный функционал доступен только преподавателю!")
)
