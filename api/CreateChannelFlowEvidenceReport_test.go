package api

import (
	"essbasic-golang-kit_/utils"
	"fmt"
	"testing"
)

func TestCreateChannelFlowEvidenceReport(t *testing.T) {
	agent := utils.SetAgent()

	flowId := "**********************"

	response := CreateChannelFlowEvidenceReport(agent, &flowId)
	// 输出json格式的字符串回包
	fmt.Printf("%s", response.ToJsonString())
}
