package api

import (
	"essbasic-golang-kit_/config"
	"fmt"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	essbasic "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/essbasic/v20210526"
)

// UploadFiles 用于生成pdf资源编号（FileIds）来配合“用PDF创建流程”接口使用，使用场景可详见“用PDF创建流程”接口说明。
// 调用时需要设置Domain, 正式环境为 file.ess.tencent.cn
// 测试环境为 https://file.test.ess.tencent.cn
// 详细参考 https://cloud.tencent.com/document/api/1420/71479
func UploadFiles(agent *essbasic.Agent, uploadFiles []*essbasic.UploadFile) *essbasic.UploadFilesResponse {
	// 实例化一个认证对象，入参需要传入腾讯云账户secretId，secretKey,此处还需注意密钥对的保密
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

	// 应用相关信息，appid 和proxyappid 必填
	request.Agent = agent
	// 上传文件内容数组，最多支持20个文件
	request.FileInfos = uploadFiles
	// 文件对应业务类型
	// 1. TEMPLATE - 模板； 文件类型：.pdf/.doc/.docx/.html
	// 2. DOCUMENT - 签署过程及签署后的合同文档/图片控件 文件类型：.pdf/.doc/.docx/.jpg/.png/.xls.xlsx/.html
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
