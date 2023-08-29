package api

import (
	"essbasic-golang-kit_/config"
	"fmt"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	essbasic "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/essbasic/v20210526"
)

func UploadFiles(agent *essbasic.Agent, uploadFiles []*essbasic.UploadFile) *essbasic.UploadFilesResponse {
	// 实例化一个认证对象，入参需要传入腾讯云账户secretId，secretKey,此处还需注意密钥对的保密
	// 密钥可前往https://console.cloud.tencent.com/cam/capi网站进行获取
	credential := common.NewCredential(
		config.SecretId,
		config.SecretKey,
	)
	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = config.FileServiceEndPoint
	// 实例化要请求产品的client对象,clientProfile是可选的
	client, err := essbasic.NewClient(credential, "", cpf)
	if err != nil {
		return nil
	}
	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := essbasic.NewUploadFilesRequest()


	request.Agent = agent

	request.FileInfos = uploadFiles

	BusinessType := "DOCUMENT"
	request.BusinessType = &BusinessType
	// 返回的resp是一个UploadFilesResponse的实例，与请求对象对应
	response, err := client.UploadFiles(request)

	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return nil
	}
	if err != nil {
		panic(err)
	}
	return response
}
