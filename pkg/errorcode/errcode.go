package errorcode

import(
	"net/http"
	"fmt"
)

type Error struct {
	CodeInfo int `json:"code"`
	MsgInfo string `json:"msg"`
	DetailsInfo []string `json:"details"`
}

var codes = map[int]string{}

func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("錯誤 %d 已經存在，請更換一個", code))
	}
	codes[code] = msg
	return &Error{CodeInfo: code, MsgInfo: msg}
}

func (e *Error) Error() string {
	return fmt.Sprintf("錯誤：%d, 錯誤資訊：%s", e.Code(), e.Msg())
}

func (e *Error) Code() int {
	return e.CodeInfo
}

func (e *Error) Msg() string {
	return e.MsgInfo
}

func (e *Error) Msgf(args []interface{}) string {
	return fmt.Sprintf(e.MsgInfo, args...)
}

func (e *Error) Details() []string {
	return e.DetailsInfo
}

func (e *Error) WithDetails(details ...string) *Error {
	newError := *e
	newError.DetailsInfo = []string{}
	for i := range details {
		// https://stackoverflow.com/questions/38692998/strange-golang-append-behavior-overwriting-values-in-slice
		newError.DetailsInfo = append(newError.DetailsInfo, details[i])
	}

	return &newError
}

func (e *Error) StatusCode() int {
	switch e.Code() {
	case Success.Code():
		return http.StatusOK
	case ServerError.Code():
		return http.StatusInternalServerError
	case InvalidParams.Code():
		return http.StatusBadRequest
	case UnauthorizedAuthNotExist.Code():
		fallthrough
	case UnauthorizedTokenError.Code():
		fallthrough
	case UnauthorizedTokenGenerate.Code():
		fallthrough
	case UnauthorizedTokenTimeout.Code():
		return http.StatusUnauthorized
	case TooManyRequests.Code():
		return http.StatusTooManyRequests
	}

	return http.StatusInternalServerError
}