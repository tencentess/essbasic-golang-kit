package api

import (
	"essbasic-golang-kit_/utils"
	"fmt"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	essbasic "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/essbasic/v20210526"
)

// ChannelCreateBoundFlows
// 此接口（ChannelCreateBoundFlows）用于渠道子客领取合同，经办人需要有相应的角色，领取后的合同不能重复领取。
// 详细参考 https://cloud.tencent.com/document/api/1420/83118
func ChannelCreateBoundFlows(agent *essbasic.Agent, flowIds []*string) *essbasic.ChannelCreateBoundFlowsResponse {
	// 实例化一个认证对象，入参需要传入腾讯云账户secretId，secretKey,此处还需注意密钥对的保密
	client, err := utils.InitClient()
	if err != nil {
		panic(err)
	}

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := essbasic.NewChannelCreateBoundFlowsRequest()

	// 应用信息
	// 此接口Agent.AppId、Agent.ProxyOrganizationOpenId 和 Agent. ProxyOperator.OpenId 必填
	request.Agent = agent
	// 领取的合同id列表
	request.FlowIds = flowIds

	// 返回的resp是一个ChannelCreateBoundFlowsResponse的实例，与请求对象对应
	response, err := client.ChannelCreateBoundFlows(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return nil
	}
	if err != nil {
		panic(err)
	}
	return response
}
