package response

import "github.com/gin-gonic/gin"

const (
	SuccessCode      = 200 // 成功
	ErrorCode        = 202 // 失败
	BadRequestCode   = 400 // 请求错误
	UnauthorizedCode = 401 // 未授权
	ForbiddenCode    = 403 // 禁止访问
	NotFoundCode     = 404 // 资源不存在
	ServerErrorCode  = 500 // 服务器错误
)

// 统一返回结构
type Response struct {
	Code    int         `json:"code"`    // 业务状态码
	Message string      `json:"message"` // 提示信息
	Data    interface{} `json:"data"`    // 返回数据
}

// 分页数据
type PageResult struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}

// 分页返回
func PageSuccess(c *gin.Context, list interface{}, total int64, page int, pageSize int) {
	pageResult := PageResult{
		List:     list,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}
	SuccessData(c, pageResult)
}

// 成功返回
func Success(ctx *gin.Context, msg ...string) {
	message := "操作成功"
	if len(msg) == 1 {
		message = msg[0]
	}
	Custom(ctx, SuccessCode, message, nil)
}

// 成功返回
func SuccessData(ctx *gin.Context, data interface{}, msg ...string) {
	message := "操作成功"
	if len(msg) == 1 {
		message = msg[0]
	}
	Custom(ctx, SuccessCode, message, data)
}

// 失败返回
func Error(ctx *gin.Context, msg ...string) {
	message := "操作失败"
	if len(msg) == 1 {
		message = msg[0]
	}
	Custom(ctx, ErrorCode, message, nil)
}

// 失败返回
func ErrorWithCode(ctx *gin.Context, code int, msg ...string) {
	message := "操作失败"
	if len(msg) == 1 {
		message = msg[0]
	}
	Custom(ctx, code, message, nil)
}

// 自定义返回
func Custom(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(200, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
