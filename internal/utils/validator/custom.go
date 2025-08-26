package validator

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"regexp"
)

// registerCustomValidators 注册自定义验证器
func registerCustomValidators() {
	// 中国手机号验证
	validate.RegisterValidation("phone", validateChinesePhone)
	validate.RegisterTranslation("phone", trans, func(ut ut.Translator) error {
		return ut.Add("phone", "{0}必须是有效的中国手机号", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("phone", fe.Field())
		return t
	})

	// 身份证验证
	validate.RegisterValidation("id_card", validateIDCard)
	validate.RegisterTranslation("id_card", trans, func(ut ut.Translator) error {
		return ut.Add("id_card", "{0}必须是有效的身份证号", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("id_card", fe.Field())
		return t
	})

	// 强密码验证
	validate.RegisterValidation("strong_password", validateStrongPassword)
	validate.RegisterTranslation("strong_password", trans, func(ut ut.Translator) error {
		return ut.Add("strong_password", "{0}必须包含大小写字母、数字和特殊字符，长度8-20位", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("strong_password", fe.Field())
		return t
	})
}

// validateChinesePhone 中国手机号验证
func validateChinesePhone(fl validator.FieldLevel) bool {
	phone := fl.Field().String()
	matched, _ := regexp.MatchString(`^1[3-9]\d{9}$`, phone)
	return matched
}

// validateIDCard 身份证验证（简化版）
func validateIDCard(fl validator.FieldLevel) bool {
	idCard := fl.Field().String()
	// 18位身份证正则
	matched, _ := regexp.MatchString(`^[1-9]\d{5}(18|19|20)\d{2}((0[1-9])|(1[0-2]))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$`, idCard)
	return matched
}

// validateStrongPassword 强密码验证
func validateStrongPassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	// 至少8位，包含大小写字母、数字、特殊字符
	patterns := []string{
		`[a-z]`,      // 小写字母
		`[A-Z]`,      // 大写字母
		`\d`,         // 数字
		`[!@#$%^&*]`, // 特殊字符
	}

	if len(password) < 8 || len(password) > 20 {
		return false
	}

	for _, pattern := range patterns {
		matched, _ := regexp.MatchString(pattern, password)
		if !matched {
			return false
		}
	}

	return true
}
