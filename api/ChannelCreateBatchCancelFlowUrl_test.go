package api

import (
	"essbasic-golang-kit_/utils"
	"fmt"
	"testing"
)

func TestChannelCreateBatchCancelFlowUrl(t *testing.T) {
	agent := utils.SetAgent()

	flowId := "**********************"

	flowIds := []*string{&flowId}

	response := ChannelCreateBatchCancelFlowUrl(agent, flowIds)
	// 输出json格式的字符串回包
	fmt.Printf("%s", response.ToJsonString())
}
