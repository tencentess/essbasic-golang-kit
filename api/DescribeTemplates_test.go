package api

import (
	"essbasic-golang-kit_/config"
	"essbasic-golang-kit_/utils"
	"fmt"
	"testing"
)

// 查询该企业在电子签渠道版中配置的有效模板列表测试
func TestDescribeTemplates(t *testing.T) {
	agent := utils.SetAgent()
	// 模板Id
	templateId := config.TemplateId

	response := DescribeTemplates(agent, &templateId)
	// 输出json格式的字符串回包
	fmt.Printf("%s", response.ToJsonString())
}
