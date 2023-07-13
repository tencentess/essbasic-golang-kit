package api

import (
	"essbasic-golang-kit_/utils"
	"fmt"
	"testing"
)

func TestChannelCreateBoundFlows(t *testing.T) {
	agent := utils.SetAgent()

	flowId := "**********************"

	flowIds := []*string{&flowId}

	response := ChannelCreateBoundFlows(agent, flowIds)
	// 输出json格式的字符串回包
	fmt.Printf("%s", response.ToJsonString())
}
