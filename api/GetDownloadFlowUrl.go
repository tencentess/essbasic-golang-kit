package api

import (
	"essbasic-golang-kit_/utils"
	"fmt"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	essbasic "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/essbasic/v20210526"
)

// GetDownloadFlowUrl
// 此接口（GetDownloadFlowUrl）用于创建电子签批量下载地址，让合作企业进入控制台直接下载，支持客户合同（流程）按照自定义文件夹形式 分类下载。
// 当前接口限制最多合同（流程）50个.
func GetDownloadFlowUrl(agent *essbasic.Agent,
	downLoadFlows []*essbasic.DownloadFlowInfo) *essbasic.GetDownloadFlowUrlResponse {
	// 实例化一个认证对象，入参需要传入腾讯云账户secretId，secretKey,此处还需注意密钥对的保密
	client, err := utils.InitClient()
	if err != nil {
		panic(err)
	}

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := essbasic.NewGetDownloadFlowUrlRequest()

	// 渠道应用相关信息
	request.Agent = agent
	// 文件夹数组，签署流程总数不能超过50个，一个文件夹下，不能超过20个签署流程
	request.DownLoadFlows = downLoadFlows

	// 返回的resp是一个GetDownloadFlowUrlResponse的实例，与请求对象对应
	response, err := client.GetDownloadFlowUrl(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return nil
	}
	if err != nil {
		panic(err)
	}
	return response
}
