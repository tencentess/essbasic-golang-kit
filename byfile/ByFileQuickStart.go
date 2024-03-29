package main

import (
	"fmt"

	"essbasic-golang-kit_/api"
	"essbasic-golang-kit_/utils"
)

/*
本示例用于第三方应用集成接口对接，通过文件快速发起第一份合同
建议配合文档进行操作，先修改config里的基本参数以及对应环境域名，然后跑一次
第三方应用集成主要针对平台企业-代理子客发起合同，简要步骤主要是
 1. 通过CreateConsoleLoginUrl引导子客企业完成电子签的实名认证 - 子客企业在电子签配置印章等
 2. 通过简单封装的CreateFlowByFileDirectly接口上传文件并快速发起一份合同，并得到签署链接
 3. 在小程序签署合同，通过API下载合同

基于具体业务上的参数调用，可以参考官网的接口说明
每个API的封装在api目录下可以自己配合相关参数进行调用
*/
func main() {
	// Step 1 登录子客控制台
	// 子客企业真实名称
	proxyOrganizationName := "我的企业"

	// 创建控制台链接
	loginUrlResp := api.CreateConsoleLoginUrl(utils.SetAgent(), proxyOrganizationName)

	// Step 2 发合同
	// 定义文件所在的路径
	filePath := "testdata/blank.pdf"
	// 定义合同名
	flowName := "我的第一个合同"
	// 将文件处理为Base64编码后的文件内容
	fileBase64 := utils.ConvertImageFileToBase64(filePath)

	// 此处为快速发起定义签署方；如果是正式接入，构造签署方，请参考函数内说明，构造需要的场景参数
	flowApproverInfos := BuildApprovers()

	// 发起合同
	resp := api.CreateFlowByFileDirectly(fileBase64, flowName, flowApproverInfos)

	fmt.Println("您的控制台入口为：")
	fmt.Println(*loginUrlResp.Response.ConsoleUrl)
	fmt.Print("\r\n\r\n")

	// 返回合同Id
	fmt.Println("您创建的合同id为：")
	fmt.Println(*resp["flowId"])
	fmt.Print("\r\n\r\n")

	// 返回合同Id
	fmt.Println("签署链接为：")
	fmt.Println(*resp["url"])
	fmt.Print("\r\n\r\n")

	// Step 3 下载合同
	// 返回合同下载链接
	url := api.DescribeFileUrls(resp["flowId"])
	fmt.Println("请访问以下地址下载您的合同：")
	fmt.Println(*url)
	fmt.Print("\r\n\r\n")
}
