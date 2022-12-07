package api

import (
	"essbasic-golang-kit_/utils"
	"fmt"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	essbasic "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/essbasic/v20210526"
)

// ChannelCreateBatchCancelFlowUrl
// 指定需要批量撤销的签署流程Id，获取批量撤销链接
// 客户指定需要撤销的签署流程Id，最多100个，超过100不处理；
// 接口调用成功返回批量撤销合同的链接，通过链接跳转到电子签小程序完成批量撤销;
// 注意:
// 能撤回合同的只能是合同的发起人或者发起企业的超管、法人
func ChannelCreateBatchCancelFlowUrl(agent *essbasic.Agent,
	flowIds []*string) *essbasic.ChannelCreateBatchCancelFlowUrlResponse {
	// 实例化一个认证对象，入参需要传入腾讯云账户secretId，secretKey,此处还需注意密钥对的保密
	client, err := utils.InitClient()
	if err != nil {
		panic(err)
	}

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := essbasic.NewChannelCreateBatchCancelFlowUrlRequest()

	// 渠道应用相关信息
	request.Agent = agent
	// 签署流程Id数组
	request.FlowIds = flowIds

	// 返回的resp是一个ChannelCreateBatchCancelFlowUrlResponse的实例，与请求对象对应
	response, err := client.ChannelCreateBatchCancelFlowUrl(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return nil
	}
	if err != nil {
		panic(err)
	}
	return response
}
