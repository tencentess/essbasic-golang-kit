package api

import (
	"essbasic-golang-kit_/utils"
	"fmt"
	"testing"
)

func TestChannelCancelFlow(t *testing.T) {
	agent := utils.SetAgent()

	flowId := "**********************"
	cancelMessage := "撤销理由"
	cancelMessageFormat := int64(3)

	response := ChannelCancelFlow(agent, &flowId, &cancelMessage, &cancelMessageFormat)
	// 输出json格式的字符串回包
	fmt.Printf("%s", response.ToJsonString())
}
