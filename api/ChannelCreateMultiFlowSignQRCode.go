package api

import (
	"fmt"

	"essbasic-golang-kit_/utils"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	essbasic "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/essbasic/v20210526"
)

// ChannelCreateMultiFlowSignQRCode 用于创建一码多扫流程签署二维码。
//
//	适用场景：无需填写签署人信息，可通过模板id生成签署二维码，签署人可通过扫描二维码补充签署信息进行实名签署。常用于提前不知道签署人的身份信息场景，例如：劳务工招工、大批量员工入职等场景。
//	适用的模板仅限于B2C（1、无序签署，2、顺序签署时B静默签署，3、顺序签署时B非首位签署）、单C的模板，且模板中发起方没有填写控件。
//
// 详细参考 https://cloud.tencent.com/document/api/1420/75452
func ChannelCreateMultiFlowSignQRCode(agent *essbasic.Agent,
	templateId, flowName string) *essbasic.ChannelCreateMultiFlowSignQRCodeResponse {
	// 实例化一个认证对象，入参需要传入腾讯云账户secretId，secretKey,此处还需注意密钥对的保密
	client, err := utils.InitClient()
	if err != nil {
		panic(err)
	}

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := essbasic.NewChannelCreateMultiFlowSignQRCodeRequest()

	// 第三方平台应用相关信息。
	// 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
	request.Agent = agent
	// 模板ID
	request.TemplateId = &templateId
	// 签署流程名称，最大长度200个字符。
	request.FlowName = &flowName

	// 返回的resp是一个ChannelCreateMultiFlowSignQRCodeResponse的实例，与请求对象对应
	response, err := client.ChannelCreateMultiFlowSignQRCode(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return nil
	}
	if err != nil {
		panic(err)
	}
	return response
}
