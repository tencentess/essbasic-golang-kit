package api

import (
	"essbasic-golang-kit_/utils"
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"testing"
)

// 根据签署流程信息批量获取资源下载链接测试
func TestDescribeResourceUrlsByFlows(t *testing.T) {
	agent := utils.SetAgent()
	// 资源所对应的签署流程Id
	flowIds := []*string{common.StringPtr("******************")}

	response := DescribeResourceUrlsByFlows(flowIds, agent)
	// 输出json格式的字符串回包
	fmt.Printf("%s", response.ToJsonString())
}
