package api

import (
	"essbasic-golang-kit_/utils"
	"fmt"
	"testing"
)

func TestDescribeUsage(t *testing.T) {
	agent := utils.SetAgent()

	startDate := "****-**-**"
	endDate := "****-**-**"
	needAggregate := false
	limit := uint64(1000)
	offset := uint64(0)

	response := DescribeUsage(agent, &startDate, &endDate, &needAggregate, &limit, &offset)
	// 输出json格式的字符串回包
	fmt.Printf("%s", response.ToJsonString())
}
