package api

import (
	"fmt"

	"essbasic-golang-kit_/utils"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	essbasic "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/essbasic/v20210526"
)

// CreateChannelFlowEvidenceReport 创建出证报告，返回报告 ID
// 详细参考 https://cloud.tencent.com/document/api/1420/79688
func CreateChannelFlowEvidenceReport(agent *essbasic.Agent,
	flowId *string) *essbasic.CreateChannelFlowEvidenceReportResponse {
	// 实例化一个认证对象，入参需要传入腾讯云账户secretId，secretKey,此处还需注意密钥对的保密
	client, err := utils.InitClient()
	if err != nil {
		panic(err)
	}

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := essbasic.NewCreateChannelFlowEvidenceReportRequest()

	// 第三方平台应用相关信息
	// 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填
	request.Agent = agent
	// 签署流程编号，合同id
	request.FlowId = flowId

	// 返回的resp是一个CreateChannelFlowEvidenceReportResponse的实例，与请求对象对应
	response, err := client.CreateChannelFlowEvidenceReport(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return nil
	}
	if err != nil {
		panic(err)
	}
	return response
}
