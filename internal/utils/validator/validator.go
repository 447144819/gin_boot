package validator

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
	"sync"
)

var (
	trans    ut.Translator
	validate *validator.Validate
	once     sync.Once
)

// ValidationError 验证错误结构
type ValidationError struct {
	Field   string `json:"field"`
	Tag     string `json:"tag"`
	Message string `json:"message"`
}

// ValidationErrors 验证错误列表
type ValidationErrors []ValidationError

// Error 实现 error 接口
func (v ValidationErrors) Error() string {
	var msgs []string
	for _, err := range v {
		msgs = append(msgs, err.Message)
	}
	return strings.Join(msgs, "; ")
}

// Init 初始化验证器（支持中文）
func Init() {
	once.Do(func() {
		// 获取gin默认的validator
		if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
			validate = v

			// 注册中文翻译器
			chinese := zh.New()
			uni := ut.New(chinese, chinese)
			trans, _ = uni.GetTranslator("zh")

			// 注册中文翻译
			zh_translations.RegisterDefaultTranslations(validate, trans)

			// 注册自定义字段名翻译
			registerFieldTranslations()

			// 注册自定义验证器
			registerCustomValidators()
		}
	})
}

// registerFieldTranslations 注册字段名中文翻译
func registerFieldTranslations() {
	// 使用json tag作为字段名
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	// 自定义字段名翻译
	fieldTranslations := map[string]string{
		"username": "用户名",
		"password": "密码",
		"email":    "邮箱",
		"phone":    "手机号",
		"age":      "年龄",
		"name":     "姓名",
	}

	for field, translation := range fieldTranslations {
		validate.RegisterTranslation(field, trans, func(ut ut.Translator) error {
			return ut.Add(field, translation, true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T(field)
			return t
		})
	}
}

// Validate 验证结构体
func Validate(obj interface{}) ValidationErrors {
	Init() // 确保已初始化

	err := validate.Struct(obj)
	if err != nil {
		var errors ValidationErrors
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, ValidationError{
				Field:   err.Field(),
				Tag:     err.Tag(),
				Message: err.Translate(trans),
			})
			break
		}
		return errors
	}
	return nil
}

// ValidateVar 验证单个变量
func ValidateVar(field interface{}, tag string) error {
	Init()
	return validate.Var(field, tag)
}

// GinBind Gin绑定并验证（支持中文错误）
func GinBind(c *gin.Context, obj interface{}) ValidationErrors {
	Init()

	// 先绑定
	if err := c.ShouldBind(obj); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			var errors ValidationErrors
			for _, err := range validationErrors {
				errors = append(errors, ValidationError{
					Field:   err.Field(),
					Tag:     err.Tag(),
					Message: err.Translate(trans),
				})
				break
			}
			return errors
		}
		// 非验证错误（如JSON解析错误）
		return ValidationErrors{{
			Field:   "",
			Tag:     "bind",
			Message: "请求参数格式错误",
		}}
	}
	return nil
}
