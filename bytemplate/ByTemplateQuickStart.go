package main

import (
	"essbasic-golang-kit_/api"
	"essbasic-golang-kit_/config"
	"fmt"
)

func main() {
	// Step 1
	// 定义合同名
	flowName := "我的第一个合同"
	// 定义企业名
	proxyOrganizationName := "我的企业"

	// 模板Id,根据自己传入的模板需求修改第参数,在config/Config中配置
	templateId := config.TemplateId

	// step 2
	// 获取模板里面的RecipientId
	recipients := GetRecipients(templateId)

	// step3
	// 此处为快速发起；如果是正式接入，构造签署人，请参考函数内说明，构造需要的场景参数
	flowApproverInfos := BuildApprovers(recipients)

	// Step 4
	// 发起合同
	// 样例为BtoC
	resp := api.CreateFlowByTemplateDirectly(flowName, proxyOrganizationName, templateId, flowApproverInfos)

	count := config.COUNT
	for i := 0; i < count; i++ {
		// 返回合同Id
		fmt.Println("您创建的合同id为：")
		fmt.Println(*resp["flowIds"][i])
		fmt.Print("\r\n\r\n")
		// 返回合同Id
		fmt.Println("签署链接为：")
		fmt.Println(*resp["urls"][i])
		fmt.Print("\r\n\r\n")
		// Step 3
		// 下载合同
		// 返回合同下载链接
		url := api.DescribeFileUrls(resp["flowIds"][i])
		fmt.Println("请访问以下地址下载您的合同：")
		fmt.Println(*url)
		fmt.Print("\r\n\r\n")
	}

}
