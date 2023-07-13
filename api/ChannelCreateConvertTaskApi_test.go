package api

import (
	"essbasic-golang-kit_/utils"
	"fmt"
	"testing"
)

func TestChannelCreateConvertTaskApi(t *testing.T) {
	agent := utils.SetAgent()

	resourceType := "html"
	resourceName := "resourceName"
	resourceId := "**********************"

	response := ChannelCreateConvertTaskApi(agent, &resourceType, &resourceName, &resourceId)
	// 输出json格式的字符串回包
	fmt.Printf("%s", response.ToJsonString())
}
