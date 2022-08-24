package api

import (
	"essbasic-golang-kit_/utils"
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"testing"
)

// 用于查询合同(签署流程)的详细信息测试
func TestDescribeFlowDetailInfo(t *testing.T) {
	agent := utils.SetAgent()
	// 资源所对应的签署流程Id
	flowIds := []*string{common.StringPtr("******************")}

	response := DescribeFlowDetailInfo(agent, flowIds)
	// 输出json格式的字符串回包
	fmt.Printf("%s", response.ToJsonString())
}
