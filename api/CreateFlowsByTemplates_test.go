package api

import (
	"essbasic-golang-kit_/utils"
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	essbasic "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/essbasic/v20210526"
	"testing"
)

// 用于使用多个模板批量创建签署流程测试
func TestCreateFlowsByTemplates(t *testing.T) {
	agent := utils.SetAgent()
	// 通过DescribeTemplates接口获取
	recipientId := "**************"

	flowInfos := []*essbasic.FlowInfo{
		{
			// 模板Id
			TemplateId: common.StringPtr("***************"),
			// 签署流程名称，长度不超过200个字符
			FlowName: common.StringPtr("我的第一份合同"),
			// 签署类型
			FlowType: common.StringPtr("合同"),
			// 签署人信息
			FlowApprovers: []*essbasic.FlowApproverInfo{
				{
					ApproverType: common.StringPtr("PERSON"),
					Name:         common.StringPtr("***"),
					Mobile:       common.StringPtr("****************"),
					RecipientId:  &recipientId,
				},
			},
		},
	}

	response := CreateFlowsByTemplates(agent, flowInfos)
	// 输出json格式的字符串回包
	fmt.Printf("%s", response.ToJsonString())
}
