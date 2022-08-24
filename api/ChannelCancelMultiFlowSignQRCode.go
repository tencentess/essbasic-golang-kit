package api

import (
	"essbasic-golang-kit_/utils"
	"fmt"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	essbasic "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/essbasic/v20210526"
)

// ChannelCancelMultiFlowSignQRCode 用于取消一码多扫二维码。该接口对传入的二维码ID，若还在有效期内，可以提前失效
func ChannelCancelMultiFlowSignQRCode(agent *essbasic.Agent, qrCodeId *string) *essbasic.ChannelCancelMultiFlowSignQRCodeResponse {
	// 实例化一个认证对象，入参需要传入腾讯云账户secretId，secretKey,此处还需注意密钥对的保密
	client, err := utils.InitClient()
	if err != nil {
		panic(err)
	}

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := essbasic.NewChannelCancelMultiFlowSignQRCodeRequest()

	// 渠道应用相关信息
	request.Agent = agent
	// 二维码ID
	request.QrCodeId = qrCodeId

	// 返回的resp是一个ChannelCancelMultiFlowSignQRCodeResponse的实例，与请求对象对应
	response, err := client.ChannelCancelMultiFlowSignQRCode(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return nil
	}
	if err != nil {
		panic(err)
	}
	return response
}
