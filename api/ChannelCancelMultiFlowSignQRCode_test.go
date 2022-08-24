package api

import (
	"essbasic-golang-kit_/utils"
	"fmt"
	"testing"
)

// 用于取消一码多扫二维码测试
func TestChannelCancelMultiFlowSignQRCode(t *testing.T) {
	agent := utils.SetAgent()
	qrCodeId := "************"

	response := ChannelCancelMultiFlowSignQRCode(agent, &qrCodeId)
	// 输出json格式的字符串回包
	fmt.Printf("%s", response.ToJsonString())
}
