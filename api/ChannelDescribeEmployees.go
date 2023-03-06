package api

import (
	"fmt"

	"essbasic-golang-kit_/utils"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	essbasic "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/essbasic/v20210526"
)

// ChannelDescribeEmployees 查询企业员工列表
// 详细参考 https://cloud.tencent.com/document/api/1420/81119
func ChannelDescribeEmployees(agent *essbasic.Agent, filters []*essbasic.Filter,
	offset *int64) *essbasic.ChannelDescribeEmployeesResponse {
	// 实例化一个认证对象，入参需要传入腾讯云账户secretId，secretKey,此处还需注意密钥对的保密
	client, err := utils.InitClient()
	if err != nil {
		panic(err)
	}

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := essbasic.NewChannelDescribeEmployeesRequest()

	// 第三方平台应用相关信息
	// 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	request.Agent = agent
	// 查询过滤实名用户，Key为Status，Values为["IsVerified"]
	// 根据第三方系统openId过滤查询员工时,Key为StaffOpenId,Values为["OpenId","OpenId",...]
	// 查询离职员工时，Key为Status，Values为["QuiteJob"]
	request.Filters = filters
	// 偏移量，默认为0，最大为20000
	request.Offset = offset

	// 返回的resp是一个ChannelDescribeEmployeesResponse的实例，与请求对象对应
	response, err := client.ChannelDescribeEmployees(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return nil
	}
	if err != nil {
		panic(err)
	}
	return response
}
