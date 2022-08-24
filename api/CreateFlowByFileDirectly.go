package api

import (
	"essbasic-golang-kit_/utils"
	v20210526 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/essbasic/v20210526"
)

// CreateFlowByFileDirectly 通过文件base64直接发起签署流程，返回flowId和签署链接
func CreateFlowByFileDirectly(fileBase64, flowName string, flowApproverInfos []*v20210526.FlowApproverInfo) map[string]*string {
	flowIdAndUrl := make(map[string]*string)

	// 设置渠道相关参数
	agent := utils.SetAgent()

	// 设置uploadFile参数
	var uploadFiles []*v20210526.UploadFile
	uploadFiles = append(uploadFiles, &v20210526.UploadFile{
		FileBody: &fileBase64,
	})

	// 上传文件获取fileId
	uploadRes := UploadFiles(agent, uploadFiles)
	fileId := uploadRes.Response.FileIds[0]

	var fileIds = []*string{fileId}
	// 创建签署流程
	createFlowRes := ChannelCreateFlowByFiles(agent, flowName, fileIds, flowApproverInfos)

	// 获取flowId
	flowId := createFlowRes.Response.FlowId

	// 获取签署链接
	var flowIds = []*string{flowId}
	createSignUrlsRes := CreateSignUrls(flowIds, agent)
	url := createSignUrlsRes.Response.SignUrlInfos[0].SignUrl

	flowIdAndUrl["flowId"] = flowId
	flowIdAndUrl["url"] = url

	return flowIdAndUrl
}
