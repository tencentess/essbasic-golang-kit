package api

import (
	"fmt"

	"essbasic-golang-kit_/utils"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	essbasic "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/essbasic/v20210526"
)

// ChannelCreateFlowByFiles 用于第三方应用集成通过文件创建签署流程。
// 注意事项：该接口需要依赖“多文件上传”接口生成pdf资源编号（FileIds）进行使用。
// 此接口静默签能力不可直接使用，需要运营申请
// 详细参考 https://cloud.tencent.com/document/api/1420/73068
func ChannelCreateFlowByFiles(
	agent *essbasic.Agent, flowName string, fileIds []*string,
	flowApproverInfos []*essbasic.FlowApproverInfo) *essbasic.ChannelCreateFlowByFilesResponse {
	// 实例化一个认证对象，入参需要传入腾讯云账户secretId，secretKey,此处还需注意密钥对的保密
	client, err := utils.InitClient()
	if err != nil {
		panic(err)
	}

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := essbasic.NewChannelCreateFlowByFilesRequest()

	// 第三方平台应用相关信息。
	// 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
	request.Agent = agent
	// 签署流程签约方列表，最多不超过5个参与方
	request.FlowApprovers = flowApproverInfos
	// 签署流程名称，长度不超过200个字符
	request.FlowName = &flowName
	// 签署文件资源Id列表，目前仅支持单个文件
	request.FileIds = fileIds

	// 其他更多参数和控制，参考文档 https://cloud.tencent.com/document/api/1420/73068
	// 也可以结合test case传参

	// 返回的resp是一个ChannelCreateFlowByFilesResponse的实例，与请求对象对应
	response, err := client.ChannelCreateFlowByFiles(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return nil
	}
	if err != nil {
		panic(err)
	}
	return response
}
