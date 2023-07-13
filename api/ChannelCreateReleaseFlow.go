package api

import (
	"fmt"

	"essbasic-golang-kit_/utils"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	essbasic "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/essbasic/v20210526"
)

// ChannelCreateReleaseFlow 第三方应用集成发起解除协议，主要应用场景为：基于一份已经签署的合同，进行解除操作。
// 合同发起人必须在电子签已经进行实名。
// 详细参考 https://cloud.tencent.com/document/api/1420/83461
func ChannelCreateReleaseFlow(agent *essbasic.Agent, needRelievedFlowId *string,
	reliveInfo *essbasic.RelieveInfo, releasedApprovers []*essbasic.ReleasedApprover,
	callbackUrl *string) *essbasic.ChannelCreateReleaseFlowResponse {
	// 实例化一个认证对象，入参需要传入腾讯云账户secretId，secretKey,此处还需注意密钥对的保密
	client, err := utils.InitClient()
	if err != nil {
		panic(err)
	}

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := essbasic.NewChannelCreateReleaseFlowRequest()

	// 第三方平台应用相关信息
	// 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
	request.Agent = agent
	// 待解除的流程编号（即原流程的编号）
	request.NeedRelievedFlowId = needRelievedFlowId
	// 解除协议内容，其中Reason必填
	// 详细参考 https://cloud.tencent.com/document/api/1420/61525#RelieveInfo
	request.ReliveInfo = reliveInfo
	// 非必须，解除协议的本企业签署人列表，默认使用原流程的签署人列表；
	// 当解除协议的签署人与原流程的签署人不能相同时（例如原流程签署人离职了），需要指定本企业的其他签署人来替换原流程中的原签署人，
	// 注意需要指明ApproverNumber来代表需要替换哪一个签署人，解除协议的签署人数量不能多于原流程的签署人数量
	request.ReleasedApprovers = releasedApprovers
	// 签署完回调url，最大长度1000个字符
	request.CallbackUrl = callbackUrl

	// 返回的resp是一个ChannelCreateReleaseFlowResponse的实例，与请求对象对应
	response, err := client.ChannelCreateReleaseFlow(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return nil
	}
	if err != nil {
		panic(err)
	}
	return response
}
