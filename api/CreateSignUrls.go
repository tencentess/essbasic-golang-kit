package api

import (
	"essbasic-golang-kit_/utils"
	"fmt"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	essbasic "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/essbasic/v20210526"
)

// CreateSignUrls 创建跳转小程序查看或签署的链接；自动签署的签署方不创建签署链接；
func CreateSignUrls(flowIds []*string, agent *essbasic.Agent) *essbasic.CreateSignUrlsResponse {
	// 实例化一个认证对象，入参需要传入腾讯云账户secretId，secretKey,此处还需注意密钥对的保密
	client, err := utils.InitClient()
	if err != nil {
		panic(err)
	}

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := essbasic.NewCreateSignUrlsRequest()

	// 渠道应用相关信息
	request.Agent = agent
	// 资源所对应的签署流程Id
	request.FlowIds = flowIds

	// 返回的resp是一个CreateSignUrlsResponse的实例，与请求对象对应
	response, err := client.CreateSignUrls(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return nil
	}
	if err != nil {
		panic(err)
	}
	return response
}
