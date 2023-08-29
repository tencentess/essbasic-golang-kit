package api

import (
	"essbasic-golang-kit_/utils"
)

func DescribeFileUrls(flowId *string) *string {
	// 设置agent参数
	agent := utils.SetAgent()

	var flowIds = []*string{flowId}
	urlResp := DescribeResourceUrlsByFlows(flowIds, agent)
	url := urlResp.Response.FlowResourceUrlInfos[0].ResourceUrlInfos[0].Url

	return url
}
