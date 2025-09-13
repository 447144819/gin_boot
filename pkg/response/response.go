package response

import (
	"fmt"
	"gin_boot/internal/utils/logs"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	ParamsError   = "参数错误"
	DoSuccess     = "操作成功"
	DoError       = "操作失败"
	AddSuccess    = "添加成功"
	AddError      = "添加失败"
	EditSuccess   = "修改成功"
	EditError     = "修改失败"
	DeleteSuccess = "删除成功"
	DeleteError   = "删除失败"
	NoExists      = "数据不存在"
)

// 统一返回结构
type Response struct {
	Code    int         `json:"code"`    // 业务状态码
	Message string      `json:"message"` // 提示信息
	Data    interface{} `json:"data"`    // 返回数据
}

// 分页数据
type PageResult struct {
	List  interface{} `json:"list"`
	Total int64       `json:"count"`
	Page  int         `json:"page"`
	Limit int         `json:"limit"`
}

// 分页返回
func PageSuccess(c *gin.Context, list interface{}, total int64, page int, limit int) {
	pageResult := PageResult{
		List:  list,
		Total: total,
		Page:  page,
		Limit: limit,
	}
	SuccessData(c, pageResult)
}

// 成功返回
func Success(ctx *gin.Context, msg ...string) {
	message := DoSuccess
	if len(msg) == 1 {
		message = msg[0]
	}
	Custom(ctx, http.StatusOK, message, nil)
}

// 成功返回
func SuccessData(ctx *gin.Context, data interface{}, msg ...string) {
	message := DoSuccess
	if len(msg) == 1 {
		message = msg[0]
	}
	Custom(ctx, http.StatusOK, message, data)
}

// 失败返回
func Error(ctx *gin.Context, err error, msg ...string) {
	message := DoError
	if len(msg) > 0 {
		message = msg[0]
	}
	// 记录错误日志
	logs.Info(fmt.Sprintf("%v:%v", message, err.Error()))
	Custom(ctx, http.StatusCreated, message, nil)
}

// 失败返回
func ErrorWithCode(ctx *gin.Context, code int, err error, msg ...string) {
	message := DoError
	if len(msg) > 0 {
		message = msg[0]
	}
	// 记录错误日志
	logs.Error(fmt.Sprintf("%v:%v", message, err.Error()))
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
