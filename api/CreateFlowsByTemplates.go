package api

import (
	"essbasic-golang-kit_/utils"
	"fmt"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	essbasic "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/essbasic/v20210526"
)

// CreateFlowsByTemplates
// 用于使用多个模板批量创建签署流程。当前可批量发起合同（签署流程）数量最大为20个。
// 如若在模板中配置了动态表格, 上传的附件必须为A4大小
// 合同发起人必须在电子签已经进行实名。
// 详细参考 https://cloud.tencent.com/document/api/1420/61523
func CreateFlowsByTemplates(agent *essbasic.Agent, flowInfos []*essbasic.FlowInfo) *essbasic.CreateFlowsByTemplatesResponse {
	// 实例化一个认证对象，入参需要传入腾讯云账户secretId，secretKey,此处还需注意密钥对的保密
	client, err := utils.InitClient()
	if err != nil {
		panic(err)
	}

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := essbasic.NewCreateFlowsByTemplatesRequest()

	// 渠道应用相关信息。 
	// 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	request.Agent = agent
	// 多个合同（签署流程）信息，最多支持20个
	// 详细参考 https://cloud.tencent.com/document/api/1420/61525#FlowInfo
	// 签署人 https://cloud.tencent.com/document/api/1420/61525#FlowApproverInfo
	request.FlowInfos = flowInfos

	// NeedPreview 可以设置预览模式，在发起合同前的预览

	// 返回的resp是一个CreateFlowsByTemplatesResponse的实例，与请求对象对应
	response, err := client.CreateFlowsByTemplates(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return nil
	}
	if err != nil {
		panic(err)
	}
	return response
}
