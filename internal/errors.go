package internal

import "errors"

var (
	ErrInvalidData         = errors.New("Некорректные данные")
	ErrUnserUnathorized    = errors.New("Пользователь не авторизован")
	ErrNoAcces             = errors.New("Пользователь не имеет доступа")
	ErrBannerNotFound      = errors.New("Баннер для не найден")
	ErrInternalServerError = errors.New("Внутренняя ошибка сервера")
)
