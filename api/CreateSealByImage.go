package api

import (
	"essbasic-golang-kit_/utils"
	"fmt"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	essbasic "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/essbasic/v20210526"
)

// CreateSealByImage
// 渠道通过图片为子客代创建印章，图片最大5m；此接口不可直接使用，请联系运营/客服咨询相关流程
func CreateSealByImage(agent *essbasic.Agent, sealName, sealImage *string) *essbasic.CreateSealByImageResponse {
	// 实例化一个认证对象，入参需要传入腾讯云账户secretId，secretKey,此处还需注意密钥对的保密
	client, err := utils.InitClient()
	if err != nil {
		panic(err)
	}

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := essbasic.NewCreateSealByImageRequest()

	// 渠道应用相关信息
	request.Agent = agent
	// 印章名称，最大长度不超过30字符
	request.SealName = sealName
	// 印章图片base64
	request.SealImage = sealImage

	// 返回的resp是一个CreateSealByImageResponse的实例，与请求对象对应
	response, err := client.CreateSealByImage(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return nil
	}
	if err != nil {
		panic(err)
	}
	return response
}
