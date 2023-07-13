package api

import (
	"essbasic-golang-kit_/utils"
	"fmt"
	"testing"
)

func TestChannelGetTaskResultApi(t *testing.T) {
	agent := utils.SetAgent()

	taskId := "**********************"

	response := ChannelGetTaskResultApi(agent, &taskId)
	// 输出json格式的字符串回包
	fmt.Printf("%s", response.ToJsonString())
}
