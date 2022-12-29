package api

import (
	"essbasic-golang-kit_/utils"
	"fmt"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	essbasic "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/essbasic/v20210526"
)

// SyncProxyOrganizationOperators
// 用于同步渠道子客企业经办人列表，主要是同步经办人的离职状态。
// 子客Web控制台的组织架构管理，是依赖于渠道平台的，无法针对员工做新增/更新/离职等操作。
// 若经办人信息有误，或者需要修改，也可以先将之前的经办人做离职操作，然后重新使用控制台链接CreateConsoleLoginUrl让经办人重新实名。
// 详细参考 https://cloud.tencent.com/document/api/1420/61517
func SyncProxyOrganizationOperators(agent *essbasic.Agent, operatorType *string,
	proxyOrganizationOperators []*essbasic.ProxyOrganizationOperator) *essbasic.SyncProxyOrganizationOperatorsResponse {
	// 实例化一个认证对象，入参需要传入腾讯云账户secretId，secretKey,此处还需注意密钥对的保密
	client, err := utils.InitClient()
	if err != nil {
		panic(err)
	}

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := essbasic.NewSyncProxyOrganizationOperatorsRequest()

	// 渠道应用相关信息。 
	// 此接口Agent.AppId 和 Agent.ProxyOrganizationOpenId必填。
	request.Agent = agent
	// 操作类型，新增:"CREATE"，修改:"UPDATE"，离职:"RESIGN"
	request.OperatorType = operatorType
	// 经办人信息列表，最大长度200
	request.ProxyOrganizationOperators = proxyOrganizationOperators

	// 返回的resp是一个SyncProxyOrganizationOperatorsResponse的实例，与请求对象对应
	response, err := client.SyncProxyOrganizationOperators(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return nil
	}
	if err != nil {
		panic(err)
	}
	return response
}
