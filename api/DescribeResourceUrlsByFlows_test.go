package api

import (
	"essbasic-golang-kit_/utils"
	"fmt"
	"testing"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
)

// 根据签署流程信息批量获取资源下载链接测试
func TestDescribeResourceUrlsByFlows(t *testing.T) {
	agent := utils.SetAgent()

	flowIds := []*string{common.StringPtr("******************")}

	response := DescribeResourceUrlsByFlows(flowIds, agent)
	// 输出json格式的字符串回包
	fmt.Printf("%s", response.ToJsonString())
}
