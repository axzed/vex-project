package common

import "regexp"

// VerifyMobile 验证手机号
func VerifyMobile(mobile string) bool {
	if mobile == "" {
		return false
	}
	// 正则表达式
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	// 编译正则表达式
	reg := regexp.MustCompile(regular)
	// 匹配手机号
	return reg.MatchString(mobile)
}
