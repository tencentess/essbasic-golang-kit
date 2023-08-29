package api

import (
	"fmt"

	"essbasic-golang-kit_/utils"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	essbasic "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/essbasic/v20210526"
)

func ChannelBatchCancelFlows(agent *essbasic.Agent, flowIds []*string,
	cancelMessage *string, cancelMessageFormat *int64) *essbasic.ChannelBatchCancelFlowsResponse {
	// 实例化一个认证对象，入参需要传入腾讯云账户secretId，secretKey,此处还需注意密钥对的保密
	client, err := utils.InitClient()
	if err != nil {
		panic(err)
	}

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := essbasic.NewChannelBatchCancelFlowsRequest()

	request.Agent = agent

	request.FlowIds = flowIds

	request.CancelMessage = cancelMessage

	request.CancelMessageFormat = cancelMessageFormat

	// 返回的resp是一个ChannelBatchCancelFlowsResponse的实例，与请求对象对应
	response, err := client.ChannelBatchCancelFlows(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return nil
	}
	if err != nil {
		panic(err)
	}
	return response
}
