package api

import (
	"fmt"
	"testing"

	"essbasic-golang-kit_/config"
	"essbasic-golang-kit_/utils"
)

// 查询该企业在电子签 第三方平台中配置的有效模板列表测试
func TestDescribeTemplates(t *testing.T) {
	agent := utils.SetAgent()
	// 模板Id
	templateId := config.TemplateId

	response := DescribeTemplates(agent, &templateId)
	// 输出json格式的字符串回包
	fmt.Printf("%s", response.ToJsonString())
}
