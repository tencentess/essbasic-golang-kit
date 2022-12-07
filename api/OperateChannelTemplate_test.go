package api

import (
	"essbasic-golang-kit_/utils"
	"fmt"
	"testing"
)

func TestOperateChannelTemplate(t *testing.T) {
	agent := utils.SetAgent()

	operateType := "SELECT"
	templateId := "**********************"
	proxyOrganizationOpenIds := "**********************"
	authTag := "all"

	response := OperateChannelTemplate(agent, &operateType, &templateId, &proxyOrganizationOpenIds, &authTag)
	// 输出json格式的字符串回包
	fmt.Printf("%s", response.ToJsonString())
}
