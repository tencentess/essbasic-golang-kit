package api

import (
	"essbasic-golang-kit_/utils"
)

// DescribeFileUrls
//  查询文件下载URL 对于DescribeResourceUrlsByFlows的封装
//  适用场景：通过传参合同流程编号，下载对应的合同PDF文件流到本地。
func DescribeFileUrls(flowId *string) *string {
	// 设置agent参数
	agent := utils.SetAgent()

	var flowIds = []*string{flowId}
	urlResp := DescribeResourceUrlsByFlows(flowIds, agent)
	url := urlResp.Response.FlowResourceUrlInfos[0].ResourceUrlInfos[0].Url

	return url
}
