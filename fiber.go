package errors

import (
	"github.com/gofiber/fiber/v2"
	"log"
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

	if er, o := err.(*Error); o {
		log.Println(er.Message, er.Stack, er.Code)
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
