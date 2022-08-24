package api

import (
	"essbasic-golang-kit_/utils"
	"fmt"
	"testing"
)

// 用于创建一码多扫流程签署二维码测试
func TestChannelCreateMultiFlowSignQRCode(t *testing.T) {
	agent := utils.SetAgent()
	templateId := "************"
	flowName := "我的第一份合同"

	response := ChannelCreateMultiFlowSignQRCode(agent, templateId, flowName)
	// 输出json格式的字符串回包
	fmt.Printf("%s", response.ToJsonString())
}
