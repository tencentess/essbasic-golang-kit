package api

import (
	"essbasic-golang-kit_/utils"
	"fmt"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	essbasic "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/essbasic/v20210526"
)

// DescribeTemplates 查询该子客企业在电子签拥有的有效模板，不包括渠道模板
// 详细参考 https://cloud.tencent.com/document/api/1420/61521
func DescribeTemplates(agent *essbasic.Agent, templateId *string) *essbasic.DescribeTemplatesResponse {
	// 实例化一个认证对象，入参需要传入腾讯云账户secretId，secretKey,此处还需注意密钥对的保密
	client, err := utils.InitClient()
	if err != nil {
		panic(err)
	}

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := essbasic.NewDescribeTemplatesRequest()

	// 渠道应用相关信息。 
	// 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	request.Agent = agent
	// 模板唯一标识，查询单个模板时使用
	request.TemplateId = templateId

	// 其他查询参数参考官网文档
	// https://cloud.tencent.com/document/api/1420/61521

	// 返回的resp是一个DescribeTemplatesResponse的实例，与请求对象对应
	response, err := client.DescribeTemplates(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return nil
	}
	if err != nil {
		panic(err)
	}
	return response

}
