package api

import (
	"essbasic-golang-kit_/utils"
	"fmt"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	essbasic "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/essbasic/v20210526"
)

// DescribeFlowDetailInfo 此接口用于查询合同(签署流程)的详细信息。
// 详细参考 https://cloud.tencent.com/document/api/1420/66683
func DescribeFlowDetailInfo(agent *essbasic.Agent, flowIds []*string) *essbasic.DescribeFlowDetailInfoResponse {
	// 实例化一个认证对象，入参需要传入腾讯云账户secretId，secretKey,此处还需注意密钥对的保密
	client, err := utils.InitClient()
	if err != nil {
		panic(err)
	}

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := essbasic.NewDescribeFlowDetailInfoRequest()

	// 渠道应用相关信息。 
	// 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	request.Agent = agent
	// 合同(流程)编号数组，最多支持100个
	request.FlowIds = flowIds

	// 返回的resp是一个DescribeFlowDetailInfoResponse的实例，与请求对象对应
	response, err := client.DescribeFlowDetailInfo(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return nil
	}
	if err != nil {
		panic(err)
	}
	return response
}
