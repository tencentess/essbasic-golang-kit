package api

import (
	"essbasic-golang-kit_/utils"
	"fmt"
	"testing"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	essbasic "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/essbasic/v20210526"
)

func TestGetDownloadFlowUrl(t *testing.T) {
	agent := utils.SetAgent()

	downLoadFlows := []*essbasic.DownloadFlowInfo{
		{
			FileName: common.StringPtr("文件夹名称"),
			FlowIdList: []*string{
				common.StringPtr("**********************"),
			},
		},
	}

	response := GetDownloadFlowUrl(agent, downLoadFlows)
	// 输出json格式的字符串回包
	fmt.Printf("%s", response.ToJsonString())
}
