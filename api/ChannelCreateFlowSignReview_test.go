package api

import (
	"essbasic-golang-kit_/utils"
	"fmt"
	"testing"
)

func TestChannelCreateFlowSignReview(t *testing.T) {
	agent := utils.SetAgent()

	flowId := "**********************"
	reviewType := "PASS"
	reviewMessage := "**********************"
	recipientId := ""

	response := ChannelCreateFlowSignReview(agent, &flowId, &reviewType, &reviewMessage, &recipientId)
	// 输出json格式的字符串回包
	fmt.Printf("%s", response.ToJsonString())
}
