package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Ctx        *gin.Context `json:"-"`
	StatusCode int          `json:"-"`
	OK         bool         `json:"ok"`
	Message    string       `json:"message,omitempty"`
	Data       any          `json:"data,omitempty"`
	Errors     any          `json:"errors,omitempty"`
}

func New(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

func (r *Response) Ok(data any) {
	r.Success(http.StatusOK, data)
}

func (r *Response) Forbidden() {
	r.Ctx.JSON(http.StatusForbidden, nil)
}

func (r *Response) BadRequest(msg string) {
	r.Ctx.JSON(http.StatusBadRequest, map[string]string{
		"message": msg,
	})
}

func (r *Response) NotFound() {
	r.Error(http.StatusNotFound, "")
}

func (r *Response) ServerError() {
	r.Error(http.StatusInternalServerError, "")
}

func (r *Response) Conflict() {
	r.Error(http.StatusConflict, "")
}

func (r *Response) json() {
	r.Ctx.JSON(r.StatusCode, r)
}

func (r *Response) setStatusCode(code int) {
	r.StatusCode = code
}

func (r *Response) setOk(v bool) {
	r.OK = v
}

func (r *Response) setData(i any) {
	r.Data = i
}

func (r *Response) SetMessage(message string) *Response {
	r.Message = message

	return r
}

func (r *Response) Error(statusCode int, v any) {
	r.setOk(false)
	r.setStatusCode(statusCode)

	r.Errors = v

	r.json()
}

func (r *Response) Success(statusCode int, v any) {
	r.setOk(true)
	r.setData(v)
	r.setStatusCode(statusCode)

	r.json()
}

func (r *Response) UnprocessableEntity(v any) {
	r.Error(http.StatusUnprocessableEntity, v)
}

func (r *Response) Created(v any) {
	r.Success(http.StatusCreated, v)
}
