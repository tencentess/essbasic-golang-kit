package api

import (
	"essbasic-golang-kit_/utils"
	"fmt"
	"testing"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
)

// 创建跳转小程序查看或签署的链接测试
func TestCreateSignUrls(t *testing.T) {
	agent := utils.SetAgent()
	// 资源所对应的签署流程Id
	flowIds := []*string{common.StringPtr("******************")}

	response := CreateSignUrls(flowIds, agent)
	// 输出json格式的字符串回包
	fmt.Printf("%s", response.ToJsonString())

}
