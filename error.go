package errors

import (
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
)

type (
	Error struct {
		Code       string `json:"code"`
		HTTPStatus int    `json:"http_status"`
		Message    string `json:"message"`
		Stack      string `json:"-"`
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
	// ErrCodeESy00001 Error no especificado
	ErrCodeESy00001 = "ESY-00001"

	// ErrCodeEdb00001 Problema de red con el servidor de base de datos.
	ErrCodeEdb00001 = "EDB-00001"
	// ErrCodeEdb00002 EL servidor de base de datos tarda demasiado en responder.
	ErrCodeEdb00002 = "EDB-00002"
	// ErrCodeEdb00003 El valor de la clave duplicada infringe el valor único.
	ErrCodeEdb00003 = "EDB-00003"

	// ErrCodeEap00001 El correo electronico ya está en uso.
	ErrCodeEap00001 = "EAP-00001"
	// ErrCodeEap00002 La contraseña debe tener minimo 8 caracteres.
	ErrCodeEap00002 = "EAP-00002"

	HttpStatus401 = 401
	HttpStatus400 = 400
	HttpStatus500 = 500
)

func (err *Error) Error() string {
	return err.Message
}

func NewErrorFromErrorMongo(err error) Error {
	customError := Error{}
	msj := err.Error()

	if mongo.IsDuplicateKeyError(err) {
		field := strings.TrimSpace(msj[strings.Index(msj, "{")+1 : strings.LastIndex(msj, ":")])
		customError.dbUnique(field, msj)
		return customError
	} else if mongo.IsNetworkError(err) {
		customError.dbNetworkError(msj)
		return customError
	} else if mongo.IsTimeout(err) {
		customError.dbTimeOut(msj)
		return customError
	}

	customError.unspecified(msj)
	return customError
}

func (e *Error) dbUnique(field, stack string) {
	e.Code = ErrCodeEdb00003
	e.HTTPStatus = HttpStatus500
	e.Message = strings.TrimSpace("Campo único duplicado: " + field + ".")
	e.Stack = stack
}

func (e *Error) dbNetworkError(stack string) {
	e.Code = ErrCodeEdb00001
	e.HTTPStatus = HttpStatus500
	e.Message = "No se realizo la conexion al servidor de base de datos."
	e.Stack = stack
}

func (e *Error) dbTimeOut(stack string) {
	e.Code = ErrCodeEdb00002
	e.HTTPStatus = HttpStatus500
	e.Message = "Tardo demasiado en responder el servidor de base de datos."
	e.Stack = stack
}

func (e *Error) unspecified(stack string) {
	e.Code = ErrCodeESy00001
	e.HTTPStatus = HttpStatus500
	e.Message = "Error interno en el sistema."
	e.Stack = stack
}
