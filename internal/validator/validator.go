package validator

import (
	"slices"
	"strings"
	"unicode/utf8"
)

// Определяем новую структуру валидатора, которая содержит карту сообщений об ошибках валидации
// для полей нашей формы.
type Validator struct {
	FieldErrors map[string]string
}

// Valid() возвращает true если карта FieldErrors не содержит никаких элементов
func (v *Validator) Valid() bool {
	return len(v.FieldErrors) == 0
}

// AddFieldError() добавляет сообщение об ошибке в карту FieldErrors (до тех пор, пока
// для данного ключа еще не существует записи ).
func (v *Validator) AddFieldError(key, message string) {

	if v.FieldErrors == nil {
		v.FieldErrors = make(map[string]string)
	}

	if _, exists := v.FieldErrors[key]; !exists {
		v.FieldErrors[key] = message
	}
}

// CheckField() добавляет сообщение об ошибке в карту FieldErrors только если
// проеверка валидации не 'ok'.
func (v *Validator) CheckField(ok bool, key, message string) {
	if !ok {
		v.AddFieldError(key, message)
	}
}

// NotBlank() возвращает true если value не пустая строка
func NotBlank(value string) bool {
	return strings.TrimSpace(value) != ""
}

// MaxChars() возвращет true если value содержит не больше n символов.
func MaxChars(value string, n int) bool {
	return utf8.RuneCountInString(value) <= n
}

// PermittedValue() возвращает true если value в списке особенных разрешенных
// значений.
func PermittedValue[T comparable](value T, permittedValues ...T) bool {
	return slices.Contains(permittedValues, value)
}
