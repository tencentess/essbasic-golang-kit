package api

import (
	"essbasic-golang-kit_/utils"
	"fmt"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	essbasic "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/essbasic/v20210526"
)

// ChannelBatchCancelFlows
// 指定需要批量撤销的签署流程Id，批量撤销合同
// 客户指定需要撤销的签署流程Id，最多100个，超过100不处理；接口失败后返回错误信息
// 注意:
// 能撤回合同的只能是合同的发起人或者发起企业的超管、法人
func ChannelBatchCancelFlows(agent *essbasic.Agent, flowIds []*string,
	cancelMessage *string, cancelMessageFormat *int64) *essbasic.ChannelBatchCancelFlowsResponse {
	// 实例化一个认证对象，入参需要传入腾讯云账户secretId，secretKey,此处还需注意密钥对的保密
	client, err := utils.InitClient()
	if err != nil {
		panic(err)
	}

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := essbasic.NewChannelBatchCancelFlowsRequest()

	// 渠道应用相关信息
	request.Agent = agent
	// 签署流程Id数组，最多100个，超过100不处理
	request.FlowIds = flowIds
	// 撤销理由
	request.CancelMessage = cancelMessage
	// 撤销理由自定义格式；选项：
	// 0 默认格式
	// 1 只保留身份信息：展示为【发起方】
	// 2 保留身份信息+企业名称：展示为【发起方xxx公司】
	// 3 保留身份信息+企业名称+经办人名称：展示为【发起方xxxx公司-经办人姓名】
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
