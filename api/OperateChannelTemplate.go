package api

import (
	"fmt"

	"essbasic-golang-kit_/utils"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	essbasic "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/essbasic/v20210526"
)

func OperateChannelTemplate(agent *essbasic.Agent, operateType, templateId,
	proxyOrganizationOpenIds, authTag *string) *essbasic.OperateChannelTemplateResponse {
	// 实例化一个认证对象，入参需要传入腾讯云账户secretId，secretKey,此处还需注意密钥对的保密
	client, err := utils.InitClient()
	if err != nil {
		panic(err)
	}

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := essbasic.NewOperateChannelTemplateRequest()

	request.Agent = agent

	request.OperateType = operateType

	request.TemplateId = templateId

	request.ProxyOrganizationOpenIds = proxyOrganizationOpenIds

	request.AuthTag = authTag

	// 返回的resp是一个OperateChannelTemplateResponse的实例，与请求对象对应
	response, err := client.OperateChannelTemplate(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return nil
	}
	if err != nil {
		panic(err)
	}
	return response
}
