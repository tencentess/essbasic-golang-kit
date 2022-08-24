package api

import (
	"essbasic-golang-kit_/utils"
	"fmt"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	essbasic "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/essbasic/v20210526"
)

// DescribeTemplates 查询该企业在电子签渠道版中配置的有效模板列表
func DescribeTemplates(agent *essbasic.Agent, templateId *string) *essbasic.DescribeTemplatesResponse {
	// 实例化一个认证对象，入参需要传入腾讯云账户secretId，secretKey,此处还需注意密钥对的保密
	client, err := utils.InitClient()
	if err != nil {
		panic(err)
	}

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := essbasic.NewDescribeTemplatesRequest()

	// 渠道应用相关信息
	request.Agent = agent
	// 模板唯一标识
	request.TemplateId = templateId

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
