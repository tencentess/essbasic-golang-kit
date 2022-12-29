package api

import (
	"essbasic-golang-kit_/utils"
	"fmt"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	essbasic "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/essbasic/v20210526"
)

// DescribeResourceUrlsByFlows
// 根据签署流程信息批量获取资源下载链接，可以下载签署中、签署完的合同，需合作企业先进行授权。
// 此接口直接返回下载的资源的url，与接口GetDownloadFlowUrl跳转到控制台的下载方式不同。
// 详细参考 https://cloud.tencent.com/document/api/1420/63220
func DescribeResourceUrlsByFlows(flowIds []*string, agent *essbasic.Agent) *essbasic.DescribeResourceUrlsByFlowsResponse {
	// 实例化一个认证对象，入参需要传入腾讯云账户secretId，secretKey,此处还需注意密钥对的保密
	client, err := utils.InitClient()
	if err != nil {
		panic(err)
	}

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := essbasic.NewDescribeResourceUrlsByFlowsRequest()

	// 渠道应用相关信息。
	// 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	request.FlowIds = flowIds
	// 查询资源所对应的签署流程Id，最多支持50个
	request.Agent = agent

	// 返回的resp是一个DescribeResourceUrlsByFlowsResponse的实例，与请求对象对应
	response, err := client.DescribeResourceUrlsByFlows(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return nil
	}
	if err != nil {
		panic(err)
	}
	return response
}
