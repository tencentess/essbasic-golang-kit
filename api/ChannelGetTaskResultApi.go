package api

import (
	"essbasic-golang-kit_/utils"
	"fmt"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	essbasic "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/essbasic/v20210526"
)

// ChannelGetTaskResultApi
// 渠道版查询转换任务状态
func ChannelGetTaskResultApi(agent *essbasic.Agent, taskId *string) *essbasic.ChannelGetTaskResultApiResponse {
	// 实例化一个认证对象，入参需要传入腾讯云账户secretId，secretKey,此处还需注意密钥对的保密
	client, err := utils.InitClient()
	if err != nil {
		panic(err)
	}

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := essbasic.NewChannelGetTaskResultApiRequest()

	// 渠道应用相关信息
	request.Agent = agent
	// 任务Id，通过ChannelCreateConvertTaskApi接口获得
	request.TaskId = taskId

	// 返回的resp是一个ChannelGetTaskResultApiResponse的实例，与请求对象对应
	response, err := client.ChannelGetTaskResultApi(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return nil
	}
	if err != nil {
		panic(err)
	}
	return response
}
