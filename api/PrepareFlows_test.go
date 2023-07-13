package api

import (
	"essbasic-golang-kit_/utils"
	"fmt"
	"testing"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	essbasic "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/essbasic/v20210526"
)

func TestPrepareFlows(t *testing.T) {
	agent := utils.SetAgent()

	flowInfos := []*essbasic.FlowInfo{
		{
			FlowName:   common.StringPtr("合同名称"),
			TemplateId: common.StringPtr("*******************"),
			FlowApprovers: []*essbasic.FlowApproverInfo{
				{
					Name:                   common.StringPtr("*****"),
					IdCardType:             nil,
					IdCardNumber:           nil,
					Mobile:                 common.StringPtr("***********"),
					OrganizationName:       nil,
					NotChannelOrganization: nil,
					OpenId:                 nil,
					OrganizationOpenId:     nil,
					ApproverType:           nil,
					RecipientId:            nil,
					Deadline:               nil,
					CallbackUrl:            nil,
					SignComponents:         nil,
					ComponentLimitType:     nil,
					PreReadTime:            nil,
					JumpUrl:                nil,
					ApproverOption:         nil,
					ApproverNeedSignReview: nil,
				},
			},
			FormFields: []*essbasic.FormField{
				{
					ComponentValue: common.StringPtr("*****"),
					ComponentName:  common.StringPtr("*****"),
				},
				{
					ComponentValue: common.StringPtr("*****"),
					ComponentName:  common.StringPtr("*****"),
				},
			},
			CallbackUrl:     nil,
			FlowType:        nil,
			FlowDescription: nil,
			CustomerData:    nil,
			CustomShowMap:   nil,
			CcInfos:         nil,
			NeedSignReview:  nil,
		},
		{
			FlowName:   common.StringPtr("合同名称1"),
			TemplateId: common.StringPtr("*******************"),
			FlowApprovers: []*essbasic.FlowApproverInfo{
				{
					Name:                   common.StringPtr("*****"),
					IdCardType:             nil,
					IdCardNumber:           nil,
					Mobile:                 common.StringPtr("***********"),
					OrganizationName:       nil,
					NotChannelOrganization: nil,
					OpenId:                 nil,
					OrganizationOpenId:     nil,
					ApproverType:           nil,
					RecipientId:            nil,
					Deadline:               nil,
					CallbackUrl:            nil,
					SignComponents:         nil,
					ComponentLimitType:     nil,
					PreReadTime:            nil,
					JumpUrl:                nil,
					ApproverOption:         nil,
					ApproverNeedSignReview: nil,
				},
			},
			FormFields: []*essbasic.FormField{
				{
					ComponentValue: common.StringPtr("*****"),
					ComponentName:  common.StringPtr("*****"),
				},
				{
					ComponentValue: common.StringPtr("*****"),
					ComponentName:  common.StringPtr("*****"),
				},
			},
			CallbackUrl:     nil,
			FlowType:        nil,
			FlowDescription: nil,
			CustomerData:    nil,
			CustomShowMap:   nil,
			CcInfos:         nil,
			NeedSignReview:  nil,
		},
	}
	jumpUrl := "***************"

	response := PrepareFlows(agent, flowInfos, &jumpUrl)
	// 输出json格式的字符串回包
	fmt.Printf("%s", response.ToJsonString())
}
