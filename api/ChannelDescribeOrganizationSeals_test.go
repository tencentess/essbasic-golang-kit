package api

import (
	"essbasic-golang-kit_/utils"
	"fmt"
	"testing"
)

func TestChannelDescribeOrganizationSeals(t *testing.T) {
	agent := utils.SetAgent()

	limit := int64(10)
	offset := int64(0)
	infoType := int64(1)
	sealId := ""

	response := ChannelDescribeOrganizationSeals(agent, &limit, &offset, &infoType, &sealId)
	// 输出json格式的字符串回包
	fmt.Printf("%s", response.ToJsonString())
}
