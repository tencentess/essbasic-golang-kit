package api

import (
	"essbasic-golang-kit_/config"
	"essbasic-golang-kit_/utils"

	v20210526 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/essbasic/v20210526"
)

// CreateFlowByTemplateDirectly 通过合同名和模板Id直接发起签署流程
func CreateFlowByTemplateDirectly(flowName, templateId string,
	flowApproverInfos []*v20210526.FlowApproverInfo) map[string][]*string {

	agent := utils.SetAgent()
	resp := make(map[string][]*string)

	// 创建签署流程
	// 签署数量
	count := config.COUNT
	var flowInfos []*v20210526.FlowInfo
	for i := 0; i < count; i++ {
		flowInfos = append(flowInfos, utils.FillFlowInfo(templateId, flowName, flowApproverInfos))
		// 构建内容控件填充结构(根据自己需求使用)
		// flowInfos[i].FormFields = []*v20210526.FormField{utils.BuildFormField("姓名", "张三")}
	}

	// 发起签署
	flowResponse := CreateFlowsByTemplates(agent, flowInfos)
	flowIds := flowResponse.Response.FlowIds
	resp["flowIds"] = flowIds

	// 获取签署链接
	createSignUrlsResp := CreateSignUrls(flowIds, agent)
	if len(createSignUrlsResp.Response.SignUrlInfos) != 0 {
		var urls []*string
		for i := 0; i < count; i++ {
			urls = append(urls, createSignUrlsResp.Response.SignUrlInfos[i].SignUrl)
		}
		resp["urls"] = urls
	}

	return resp
}
