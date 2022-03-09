package errors

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)

const (
	ErrCodeApp    = "errCodeApp"
	ErrMessaje    = "errMessaje"
	ErrHttpStatus = "errHttpStatus"
)

func HandlerFiberError(ctx *fiber.Ctx, err error) error {
	status, _ := strconv.Atoi(ctx.Get(ErrHttpStatus, "500"))
	code := ctx.Get(ErrCodeApp, ErrCodeESy00001)
	msj := ctx.Get(ErrMessaje, "")

	if e, ok := err.(*fiber.Error); ok {
		status = e.Code
	}

	if err, isErr := err.(*Error); isErr {
		status = err.HTTPStatus
		code = err.Code
		msj = err.Message
	}

	if msj == "" {
		msj = err.Error()
	}

	return ctx.Status(status).JSON(Error{
		Code:       code,
		HTTPStatus: status,
		Message:    msj,
	})
}

func SetErrorContext(ctx *fiber.Ctx, err *Error) {
	ctx.Set(ErrCodeApp, err.Code)
	ctx.Set(ErrHttpStatus, string(rune(err.HTTPStatus)))
	ctx.Set(ErrMessaje, err.Message)
}
