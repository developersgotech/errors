package errors

import "strings"

type (
	Error struct {
		Code       string
		HTTPStatus int
		Message    string
	}
)

/*
200 - Bien	Todo funcionó como se esperaba.
400 - Petición Incorrecta	La solicitud era inaceptable, a menudo debido a que faltaba un parámetro obligatorio.
401 - No autorizado	No se ha proporcionado ninguna clave de API válida.
402 - Solicitud fallida	Los parámetros eran válidos pero la solicitud falló.
403 - Prohibido	La clave API no tiene permisos para realizar la solicitud.
404 - No encontrado	El recurso solicitado no existe.
409 - Conflicto	La solicitud entra en conflicto con otra solicitud (quizás debido al uso de la misma clave idempotente).
429 - Demasiadas solicitudes	Demasiadas solicitudes llegan a la API demasiado rápido. Recomendamos un retroceso exponencial de sus solicitudes.
500, 502, 503, 504 - Errores del servidor	Algo salió mal por parte de Stripe. (Estos son raros.)
*/

const (
	ErrCodeEdb00001 = "EDB-00001"

	ErrCodeEap00001 = "EAP-00001"

	HttpStatus400 = 400
	HttpStatus500 = 500
)

func (err *Error) Error() string {
	return err.Message
}

func New(code, message string, status int) Error {
	return Error{
		Code:       code,
		HTTPStatus: status,
		Message:    message,
	}
}

func (err *Error) DBUnique(field string) {
	err.Code = ErrCodeEdb00001
	err.HTTPStatus = HttpStatus500
	err.Message = strings.TrimSpace("Campo único duplicado: " + field + ".")
}

func (err *Error) APPFieldFormat(field, format string) {
	err.Code = ErrCodeEap00001
	err.HTTPStatus = HttpStatus400
	err.Message = strings.TrimSpace("El campo " + field + " requiere el siguiente formato " + format + ".")
}
