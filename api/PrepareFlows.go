package api

import (
	"essbasic-golang-kit_/utils"
	"fmt"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	essbasic "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/essbasic/v20210526"
)

// PrepareFlows
// 该接口 (PrepareFlows) 用于创建待发起文件
// 用户通过该接口进入签署流程发起的确认页面，进行发起信息二次确认， 如果确认则进行正常发起。
// 目前该接口只支持B2C，不建议使用，将会废弃。
func PrepareFlows(agent *essbasic.Agent, flowInfos []*essbasic.FlowInfo,
	jumpUrl *string) *essbasic.PrepareFlowsResponse {
	// 实例化一个认证对象，入参需要传入腾讯云账户secretId，secretKey,此处还需注意密钥对的保密
	client, err := utils.InitClient()
	if err != nil {
		panic(err)
	}

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := essbasic.NewPrepareFlowsRequest()

	// 渠道应用相关信息
	request.Agent = agent
	// 多个合同（签署流程）信息，最大支持20个签署流程。
	request.FlowInfos = flowInfos
	// 操作完成后的跳转地址，最大长度200
	request.JumpUrl = jumpUrl

	// 返回的resp是一个PrepareFlowsResponse的实例，与请求对象对应
	response, err := client.PrepareFlows(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return nil
	}
	if err != nil {
		panic(err)
	}
	return response
}
