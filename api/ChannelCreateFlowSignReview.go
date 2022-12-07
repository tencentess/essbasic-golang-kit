package api

import (
	"essbasic-golang-kit_/utils"
	"fmt"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	essbasic "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/essbasic/v20210526"
)

// ChannelCreateFlowSignReview
// 提交企业签署流程审批结果
//
// 在通过接口(CreateFlowsByTemplates或者ChannelCreateFlowByFiles)创建签署流程时，
// 若指定了参数 NeedSignReview 为true,则可以调用此接口提交企业内部签署审批结果。
// 若签署流程状态正常，且本企业存在签署方未签署，同一签署流程可以多次提交签署审批结果，签署时的最后一个“审批结果”有效。
func ChannelCreateFlowSignReview(agent *essbasic.Agent, flowId, reviewType,
	reviewMessage, recipientId *string) *essbasic.ChannelCreateFlowSignReviewResponse {
	// 实例化一个认证对象，入参需要传入腾讯云账户secretId，secretKey,此处还需注意密钥对的保密
	client, err := utils.InitClient()
	if err != nil {
		panic(err)
	}

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := essbasic.NewChannelCreateFlowSignReviewRequest()

	// 渠道应用相关信息
	request.Agent = agent
	// 签署流程编号
	request.FlowId = flowId
	// 企业内部审核结果
	// PASS: 通过
	// REJECT: 拒绝
	// SIGN_REJECT:拒签(流程结束)	request.ReviewType = flowFileInfos
	request.ReviewType = reviewType
	// 审核原因
	// 当ReviewType 是REJECT 时此字段必填,字符串长度不超过200
	request.ReviewMessage = reviewMessage
	// 签署节点审核时需要指定
	request.RecipientId = recipientId

	// 返回的resp是一个ChannelCreateFlowSignReviewResponse的实例，与请求对象对应
	response, err := client.ChannelCreateFlowSignReview(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return nil
	}
	if err != nil {
		panic(err)
	}
	return response
}
