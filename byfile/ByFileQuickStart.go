package main

import (
	"essbasic-golang-kit_/api"
	"essbasic-golang-kit_/utils"
	"fmt"
)

func main() {
	// Step 1
	// 定义文件所在的路径
	filePath := "testdata/blank.pdf"
	// 定义合同名
	flowName := "我的第一个合同"

	// 此处为快速发起；如果是正式接入，构造签署人，请参考函数内说明，构造需要的场景参数
	flowApproverInfos := BuildApprovers()

	// Step 2
	// 将文件处理为Base64编码后的文件内容
	fileBase64 := utils.ConvertImageFileToBase64(filePath)

	// 发起合同
	resp := api.CreateFlowByFileDirectly(fileBase64, flowName, flowApproverInfos)

	// 返回合同Id
	fmt.Println("您创建的合同id为：")
	fmt.Println(*resp["flowId"])
	fmt.Print("\r\n\r\n")
	// 返回合同Id
	fmt.Println("签署链接为：")
	fmt.Println(*resp["url"])
	fmt.Print("\r\n\r\n")
	// Step 3
	// 下载合同
	// 返回合同下载链接
	url := api.DescribeFileUrls(resp["flowId"])
	fmt.Println("请访问以下地址下载您的合同：")
	fmt.Println(*url)
	fmt.Print("\r\n\r\n")
}
