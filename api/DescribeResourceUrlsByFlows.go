package api

import (
	"essbasic-golang-kit_/utils"
	"fmt"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	essbasic "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/essbasic/v20210526"
)

// DescribeResourceUrlsByFlows 根据签署流程信息批量获取资源下载链接，需合作企业先进行授权
func DescribeResourceUrlsByFlows(flowIds []*string, agent *essbasic.Agent) *essbasic.DescribeResourceUrlsByFlowsResponse {
	// 实例化一个认证对象，入参需要传入腾讯云账户secretId，secretKey,此处还需注意密钥对的保密
	client, err := utils.InitClient()
	if err != nil {
		panic(err)
	}

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := essbasic.NewDescribeResourceUrlsByFlowsRequest()

	//资源所对应的签署流程Id
	request.FlowIds = flowIds
	// 渠道应用相关信息
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
