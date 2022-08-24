package api

import (
	"essbasic-golang-kit_/utils"
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	essbasic "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/essbasic/v20210526"
)

// ChannelCreateFlowByFiles
//  用来通过上传后的pdf资源编号来创建待签署的合同流程。
//  适用场景1：适用非制式的合同文件签署。一般开发者自己有完整的签署文件，可以通过该接口传入完整的PDF文件及流程信息生成待签署的合同流程。
//  适用场景2：可通过改接口传入制式合同文件，同时在指定位置添加签署控件。可以起到接口创建临时模板的效果。如果是标准的制式文件，建议使用模板功能生成模板ID进行合同流程的生成。
//  注意事项：该接口需要依赖“多文件上传”接口生成pdf资源编号（FileIds）进行使用。
func ChannelCreateFlowByFiles(agent *essbasic.Agent, flowName string, fileIds []*string, flowApproverInfos []*essbasic.FlowApproverInfo) *essbasic.ChannelCreateFlowByFilesResponse {
	// 实例化一个认证对象，入参需要传入腾讯云账户secretId，secretKey,此处还需注意密钥对的保密
	client, err := utils.InitClient()
	if err != nil {
		panic(err)
	}

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := essbasic.NewChannelCreateFlowByFilesRequest()

	// 渠道应用相关信息
	request.Agent = agent
	// 签署流程签约方列表，最多不超过5个参与方
	request.FlowApprovers = flowApproverInfos
	// 签署流程名称，长度不超过200个字符
	request.FlowName = &flowName
	// 签署文件资源Id列表，目前仅支持单个文件
	request.FileIds = fileIds

	// 返回的resp是一个ChannelCreateFlowByFilesResponse的实例，与请求对象对应
	response, err := client.ChannelCreateFlowByFiles(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return nil
	}
	if err != nil {
		panic(err)
	}
	return response
}
