package api

import (
	"essbasic-golang-kit_/utils"
	"fmt"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	essbasic "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/essbasic/v20210526"
)

// SyncProxyOrganization
// 用于同步渠道子客企业信息，主要是子客企业的营业执照，便于子客企业开通过程中不用手动上传。
// 若有需要调用此接口，需要在创建控制链接CreateConsoleLoginUrl之后即刻进行调用。
// 详细参考 https://cloud.tencent.com/document/api/1420/61518
func SyncProxyOrganization(agent *essbasic.Agent, proxyOrganizationName,
	businessLicense, uniformSocialCreditCode, proxyLegalName *string) *essbasic.SyncProxyOrganizationResponse {
	// 实例化一个认证对象，入参需要传入腾讯云账户secretId，secretKey,此处还需注意密钥对的保密
	client, err := utils.InitClient()
	if err != nil {
		panic(err)
	}

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := essbasic.NewSyncProxyOrganizationRequest()

	// 应用信息
	// 此接口Agent.AppId、Agent.ProxyOrganizationOpenId必填
	request.Agent = agent
	// 渠道侧合作企业名称，最大长度64个字符
	request.ProxyOrganizationName = proxyOrganizationName
	// 营业执照正面照(PNG或JPG) base64格式, 大小不超过5M
	request.BusinessLicense = businessLicense
	// 渠道侧合作企业统一社会信用代码，最大长度200个字符
	request.UniformSocialCreditCode = uniformSocialCreditCode
	// 渠道侧合作企业法人/负责人姓名
	request.ProxyLegalName = proxyLegalName

	// 返回的resp是一个SyncProxyOrganizationResponse的实例，与请求对象对应
	response, err := client.SyncProxyOrganization(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return nil
	}
	if err != nil {
		panic(err)
	}
	return response
}
