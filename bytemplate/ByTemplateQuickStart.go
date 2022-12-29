package main

import (
	"essbasic-golang-kit_/api"
	"essbasic-golang-kit_/config"
	"essbasic-golang-kit_/utils"
	"fmt"
)

/*
本示例用于渠道版接口对接，通过模板快速发起第一份合同
建议配合文档进行操作，先修改config里的基本参数以及对应环境域名，然后跑一次
渠道版主要针对渠道企业-代理子客发起合同，简要步骤主要是
	1. 通过CreateConsoleLoginUrl引导子客企业完成电子签的实名认证 - 子客企业在电子签配置印章等
	2. 通过简单封装的CreateFlowByTemplateDirectly接口快速发起一份合同，并得到签署链接
	3. 在小程序签署合同，通过API下载合同
基于具体业务上的参数调用，可以参考官网的接口说明 
https://cloud.tencent.com/document/product/1420/61534
每个API的封装在api目录下可以自己配合相关参数进行调用
*/
func main() {
	// Step 1 登录子客控制台
	// 渠道子客企业真实名称
	proxyOrganizationName := "我的企业"
	
	// 创建控制台链接
	loginUrlResp := api.CreateConsoleLoginUrl(utils.SetAgent(), proxyOrganizationName)
	
	// Step 2 发合同
	// 定义合同名
	flowName := "我的第一个合同"
	// 模板Id,根据自己传入的模板需求修改的参数,在config/Config中配置
	templateId := config.TemplateId

	// 获取模板里面的参与方RecipientId
	recipients := GetRecipients(templateId)

	// 此处为快速发起的签署方；如果是正式接入，构造签署方，请参考函数内说明，构造需要的场景参数
	flowApproverInfos := BuildApprovers(recipients)

	// 发起合同
	resp := api.CreateFlowByTemplateDirectly(flowName, templateId, flowApproverInfos)

	count := config.COUNT
	fmt.Println("您的控制台入口为：")
	fmt.Println(*loginUrlResp.Response.ConsoleUrl)
	fmt.Print("\r\n\r\n")

	for i := 0; i < count; i++ {
		// 返回合同Id
		fmt.Println("您创建的合同id为：")
		fmt.Println(*resp["flowIds"][i])
		fmt.Print("\r\n\r\n")
		
		// 返回签署的链接
		fmt.Println("签署链接为：")
		fmt.Println(*resp["urls"][i])
		fmt.Print("\r\n\r\n")

		// Step 3 下载合同
		// 返回合同下载链接
		url := api.DescribeFileUrls(resp["flowIds"][i])
		fmt.Println("请访问以下地址下载您的合同：")
		fmt.Println(*url)
		fmt.Print("\r\n\r\n")
	}

}
