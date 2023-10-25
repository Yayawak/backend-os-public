package errors

import "errors"

type ApiError interface {
	Code() string
	StatusCode() int
	ThMessage() string
	EnMessage() string
	Error() string
	ErrorResponse() (int, ApiErrorResponse)
}
type HttpError struct {
	error string
}
type ApiErrorResponse struct {
	Code      string `json:"code"`
	Error     string `json:"error"` // for dev to read
	ThMessage string `json:"thMessage"`
	EnMessage string `json:"enMessage"`
}

func New(ErrorText string) *HttpError {
	return &HttpError{error: ErrorText}
}
func InitError(e error) error {
	if apiError, ok := e.(ApiError); ok {
		return apiError
	} else {
		return errors.New(e.Error())
	}
}
func (e *HttpError) Error() string {
	return e.error
}
func (e *HttpError) Code() string {
	if code, ok := CustomCode[e]; ok {
		return code
	} else {
		return InternalServerError.Code()
	}
}
func (e *HttpError) StatusCode() int {
	if status, ok := HttpStatusCode[e]; ok {
		return status
	} else {
		return InternalServerError.StatusCode()
	}
}
func (e *HttpError) ThMessage() string {
	if th, ok := ThMessages[e]; ok {
		return th
	} else {
		return InternalServerError.ThMessage()
	}
}
func (e *HttpError) EnMessage() string {
	if en, ok := EnMessages[e]; ok {
		return en
	} else {
		return InternalServerError.EnMessage()
	}
}
func (e *HttpError) ErrorResponse() (int, ApiErrorResponse) {
	if code, ok := HttpStatusCode[e]; ok {
		return code, ApiErrorResponse{
			Code:      e.Code(),
			Error:     e.Error(),
			ThMessage: e.ThMessage(),
			EnMessage: e.EnMessage(),
		}
	} else {
		return InternalServerError.StatusCode(), ApiErrorResponse{
			Code:      InternalServerError.Code(),
			Error:     InternalServerError.Error(),
			ThMessage: InternalServerError.ThMessage(),
			EnMessage: InternalServerError.EnMessage(),
		}
	}
}
func GetErrorResponse(e error) (int, ApiErrorResponse) {
	if statusCode, ok := HttpStatusCode[e]; ok {
		if customCode, ok := CustomCode[e]; ok {
			if thMessage, ok := ThMessages[e]; ok {
				if enMessage, ok := EnMessages[e]; ok {
					return statusCode, ApiErrorResponse{
						Code:      customCode,
						Error:     e.Error(),
						ThMessage: thMessage,
						EnMessage: enMessage,
					}
				}
			}
		}
	}
	return InternalServerError.StatusCode(), ApiErrorResponse{
		Code:      InternalServerError.Code(),
		Error:     InternalServerError.Error(),
		ThMessage: InternalServerError.ThMessage(),
		EnMessage: InternalServerError.EnMessage(),
	}
}

var (
	InternalServerError ApiError = New("internal_server_error")
	DataNotFound        ApiError = New("data_not_found")
	InvalidBodyRequest  ApiError = New("invalid_body_request")
	MissingBodyRequest  ApiError = New("missing_body_request")
	PageNotFound        ApiError = New("page_not_found")
)

var CustomCode = map[error]string{
	InternalServerError: "10000",
	DataNotFound:        "10001",
	InvalidBodyRequest:  "20000",
	MissingBodyRequest:  "20001",
	PageNotFound:        "20002",
}
var HttpStatusCode = map[error]int{
	InternalServerError: 500,
	DataNotFound:        404,
	InvalidBodyRequest:  400,
	MissingBodyRequest:  400,
	PageNotFound:        404,
}
var ThMessages = map[error]string{
	InternalServerError: "เกิดข้อผิดผลาดขึ้นที่เซิฟเวอร์  กรุณาลองใหม่อีกครั้งในภายหลัง",
	DataNotFound:        "ไม่พบข้อมูลที่ต้องการ  กรุณาตรวจสอบอีกครั้งก่อนทำการส่งใหม่",
	InvalidBodyRequest:  "ข้อมูลไม่ตรงตามเงื่อนไข  กรุณาตรวจสอบความถูกต้องก่อนลองใหม่อีกครั้ง",
	MissingBodyRequest:  "ข้อมูลไม่ครบจามเงื่อนไข  กรุณาตรวจสอบความถูกต้องก่อนลองใหม่อีกครั้ง",
	PageNotFound:        "ไม่พบหน้าเว็บที่ต้องการ  กรุณาเช็ค url ของท่านและลองใหม่อีกครั้ง",
}
var EnMessages = map[error]string{
	InternalServerError: "Server error. Please try again later",
	DataNotFound:        "Data is not exist. Please check your request and sent it again later.",
	InvalidBodyRequest:  "Invalid body request. Please check it and try again.",
	MissingBodyRequest:  "Missing body content. Please check it and try again.",
	PageNotFound:        "Page not found. Please check your url and try again.",
}
