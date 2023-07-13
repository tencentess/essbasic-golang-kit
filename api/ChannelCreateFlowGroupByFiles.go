package api

import (
	"fmt"

	"essbasic-golang-kit_/utils"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	essbasic "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/essbasic/v20210526"
)

// ChannelCreateFlowGroupByFiles 用于通过多文件创建合同组签署流程。
// 详细参考 https://cloud.tencent.com/document/api/1420/80390
func ChannelCreateFlowGroupByFiles(agent *essbasic.Agent, flowGroupName *string,
	flowFileInfos []*essbasic.FlowFileInfo) *essbasic.ChannelCreateFlowGroupByFilesResponse {
	// 实例化一个认证对象，入参需要传入腾讯云账户secretId，secretKey,此处还需注意密钥对的保密
	client, err := utils.InitClient()
	if err != nil {
		panic(err)
	}

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := essbasic.NewChannelCreateFlowGroupByFilesRequest()

	// 第三方平台应用相关信息
	// 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
	request.Agent = agent
	// 合同组名称，长度不超过200个字符
	request.FlowGroupName = flowGroupName
	// 每个子合同的发起所需的信息，数量限制2-100
	// 详细参考 https://cloud.tencent.com/document/product/1420/61534
	request.FlowFileInfos = flowFileInfos

	// 返回的resp是一个ChannelCreateFlowGroupByFilesResponse的实例，与请求对象对应
	response, err := client.ChannelCreateFlowGroupByFiles(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return nil
	}
	if err != nil {
		panic(err)
	}
	return response
}
