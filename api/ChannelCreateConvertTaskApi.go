package api

import (
	"essbasic-golang-kit_/utils"
	"fmt"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	essbasic "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/essbasic/v20210526"
)

// ChannelCreateConvertTaskApi
// 渠道创建文件转换任务
func ChannelCreateConvertTaskApi(agent *essbasic.Agent,
	resourceType, resourceName, resourceId *string) *essbasic.ChannelCreateConvertTaskApiResponse {
	// 实例化一个认证对象，入参需要传入腾讯云账户secretId，secretKey,此处还需注意密钥对的保密
	client, err := utils.InitClient()
	if err != nil {
		panic(err)
	}

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := essbasic.NewChannelCreateConvertTaskApiRequest()

	// 渠道应用相关信息
	request.Agent = agent
	// 资源类型 取值范围doc,docx,html,excel之一
	request.ResourceType = resourceType
	// 资源名称，长度限制为256字符
	request.ResourceName = resourceName
	// 资源Id，通过UploadFiles获取
	request.ResourceId = resourceId

	// 返回的resp是一个ChannelCreateConvertTaskApiResponse的实例，与请求对象对应
	response, err := client.ChannelCreateConvertTaskApi(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return nil
	}
	if err != nil {
		panic(err)
	}
	return response
}
