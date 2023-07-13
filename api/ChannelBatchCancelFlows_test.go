package api

import (
	"essbasic-golang-kit_/utils"
	"fmt"
	"testing"
)

func TestChannelBatchCancelFlows(t *testing.T) {
	agent := utils.SetAgent()

	flowId := "**********************"

	flowIds := []*string{&flowId}
	cancelMessage := "撤销理由"
	cancelMessageFormat := int64(3)

	response := ChannelBatchCancelFlows(agent, flowIds, &cancelMessage, &cancelMessageFormat)
	// 输出json格式的字符串回包
	fmt.Printf("%s", response.ToJsonString())
}
