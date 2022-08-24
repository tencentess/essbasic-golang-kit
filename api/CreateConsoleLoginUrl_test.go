package api

import (
	"essbasic-golang-kit_/utils"
	"fmt"
	"testing"
)

// 用于创建电子签控制台登录链接测试
func TestCreateConsoleLoginUrl(t *testing.T) {
	agent := utils.SetAgent()
	proxyOrganizationName := "************"

	response := CreateConsoleLoginUrl(agent, proxyOrganizationName)
	// 输出json格式的字符串回包
	fmt.Printf("%s", response.ToJsonString())
}
