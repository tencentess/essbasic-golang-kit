package api

import (
	"essbasic-golang-kit_/utils"
	"fmt"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	essbasic "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/essbasic/v20210526"
)

// CreateSealByImage
// 渠道通过图片为子客代创建印章，图片最大5MB
// 详细参考 https://cloud.tencent.com/document/api/1420/73067
func CreateSealByImage(agent *essbasic.Agent, sealName, sealImage *string) *essbasic.CreateSealByImageResponse {
	// 实例化一个认证对象，入参需要传入腾讯云账户secretId，secretKey,此处还需注意密钥对的保密
	client, err := utils.InitClient()
	if err != nil {
		panic(err)
	}

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := essbasic.NewCreateSealByImageRequest()

	// 渠道应用相关信息
	// 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 和 Agent.ProxyAppId 均必填。
	request.Agent = agent
	// 印章名称，最大长度不超过50字符
	request.SealName = sealName
	// 印章图片base64，大小不超过10M（原始图片不超过7.6M）
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
