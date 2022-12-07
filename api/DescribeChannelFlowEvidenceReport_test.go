package api

import (
	"essbasic-golang-kit_/utils"
	"fmt"
	"testing"
)

func TestDescribeChannelFlowEvidenceReport(t *testing.T) {
	agent := utils.SetAgent()

	reportId := "**********************"

	response := DescribeChannelFlowEvidenceReport(agent, &reportId)
	// 输出json格式的字符串回包
	fmt.Printf("%s", response.ToJsonString())
}
