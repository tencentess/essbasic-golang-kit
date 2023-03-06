package api

import (
	"essbasic-golang-kit_/utils"
	"fmt"
	"testing"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	essbasic "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/essbasic/v20210526"
)

func TestChannelCreateFlowGroupByFiles(t *testing.T) {

	agent := utils.SetAgent()

	flowGroupName := "合同组名称"

	flowFileInfos := []*essbasic.FlowFileInfo{
		packFlowFileOne(),
		packFlowFileTwo(),
	}

	response := ChannelCreateFlowGroupByFiles(agent, &flowGroupName, flowFileInfos)
	// 输出json格式的字符串回包
	fmt.Printf("%s", response.ToJsonString())
}

func packFlowFileOne() *essbasic.FlowFileInfo {
	return &essbasic.FlowFileInfo{
		FileIds: []*string{
			common.StringPtr("**********************"),
		},
		FlowName: common.StringPtr("合同名称1"),
		FlowApprovers: []*essbasic.FlowApproverInfo{
			{
				Name:                   nil,
				IdCardType:             nil,
				IdCardNumber:           nil,
				Mobile:                 nil,
				OrganizationName:       common.StringPtr("****************"),
				NotChannelOrganization: nil,
				OpenId:                 common.StringPtr("****************"),
				OrganizationOpenId:     common.StringPtr("****************"),
				ApproverType:           common.StringPtr("****************"),
				RecipientId:            nil,
				Deadline:               nil,
				CallbackUrl:            nil,
				SignComponents: []*essbasic.Component{
					{
						ComponentId:           nil,
						ComponentType:         common.StringPtr("SIGN_SEAL"),
						ComponentName:         nil,
						ComponentRequired:     nil,
						ComponentRecipientId:  nil,
						FileIndex:             nil,
						GenerateMode:          nil,
						ComponentWidth:        common.Float64Ptr(100),
						ComponentHeight:       common.Float64Ptr(100),
						ComponentPage:         common.Int64Ptr(1),
						ComponentPosX:         common.Float64Ptr(100),
						ComponentPosY:         common.Float64Ptr(100),
						ComponentExtra:        nil,
						ComponentValue:        nil,
						ComponentDateFontSize: nil,
						DocumentId:            nil,
						ComponentDescription:  nil,
						OffsetX:               nil,
						OffsetY:               nil,
						KeywordPage:           nil,
						RelativeLocation:      nil,
						KeywordIndexes:        nil,
					},
				},
				ComponentLimitType:     nil,
				PreReadTime:            nil,
				JumpUrl:                nil,
				ApproverOption:         nil,
				ApproverNeedSignReview: nil,
			},
			{
				Name:                   common.StringPtr("***"),
				IdCardType:             nil,
				IdCardNumber:           nil,
				Mobile:                 common.StringPtr("***********"),
				OrganizationName:       nil,
				NotChannelOrganization: nil,
				OpenId:                 nil,
				OrganizationOpenId:     nil,
				ApproverType:           common.StringPtr("PERSON"),
				RecipientId:            nil,
				Deadline:               nil,
				CallbackUrl:            nil,
				SignComponents: []*essbasic.Component{
					{
						ComponentId:           nil,
						ComponentType:         common.StringPtr("SIGN_SIGNATURE"),
						ComponentName:         nil,
						ComponentRequired:     nil,
						ComponentRecipientId:  nil,
						FileIndex:             nil,
						GenerateMode:          nil,
						ComponentWidth:        common.Float64Ptr(100),
						ComponentHeight:       common.Float64Ptr(100),
						ComponentPage:         common.Int64Ptr(1),
						ComponentPosX:         common.Float64Ptr(100),
						ComponentPosY:         common.Float64Ptr(200),
						ComponentExtra:        nil,
						ComponentValue:        nil,
						ComponentDateFontSize: nil,
						DocumentId:            nil,
						ComponentDescription:  nil,
						OffsetX:               nil,
						OffsetY:               nil,
						KeywordPage:           nil,
						RelativeLocation:      nil,
						KeywordIndexes:        nil,
					},
				},
				ComponentLimitType:     nil,
				PreReadTime:            nil,
				JumpUrl:                nil,
				ApproverOption:         nil,
				ApproverNeedSignReview: nil,
			},
		},
		Deadline:        common.Int64Ptr(0),
		FlowDescription: common.StringPtr("合同流程的描述"),
		FlowType:        nil,
		CallbackUrl:     nil,
		CustomerData:    nil,
		Unordered:       nil,
		CustomShowMap:   nil,
		NeedSignReview:  nil,
	}
}

func packFlowFileTwo() *essbasic.FlowFileInfo {
	return &essbasic.FlowFileInfo{
		FileIds: []*string{
			common.StringPtr("**********************"),
		},
		FlowName: common.StringPtr("合同名称2"),
		FlowApprovers: []*essbasic.FlowApproverInfo{
			{
				Name:                   nil,
				IdCardType:             nil,
				IdCardNumber:           nil,
				Mobile:                 nil,
				OrganizationName:       common.StringPtr("****************"),
				NotChannelOrganization: nil,
				OpenId:                 common.StringPtr("****************"),
				OrganizationOpenId:     common.StringPtr("****************"),
				ApproverType:           common.StringPtr("ORGANIZATION"),
				RecipientId:            nil,
				Deadline:               nil,
				CallbackUrl:            nil,
				SignComponents: []*essbasic.Component{
					{
						ComponentId:           nil,
						ComponentType:         common.StringPtr("SIGN_SEAL"),
						ComponentName:         nil,
						ComponentRequired:     nil,
						ComponentRecipientId:  nil,
						FileIndex:             nil,
						GenerateMode:          nil,
						ComponentWidth:        common.Float64Ptr(100),
						ComponentHeight:       common.Float64Ptr(100),
						ComponentPage:         common.Int64Ptr(1),
						ComponentPosX:         common.Float64Ptr(100),
						ComponentPosY:         common.Float64Ptr(100),
						ComponentExtra:        nil,
						ComponentValue:        nil,
						ComponentDateFontSize: nil,
						DocumentId:            nil,
						ComponentDescription:  nil,
						OffsetX:               nil,
						OffsetY:               nil,
						KeywordPage:           nil,
						RelativeLocation:      nil,
						KeywordIndexes:        nil,
					},
				},
				ComponentLimitType:     nil,
				PreReadTime:            nil,
				JumpUrl:                nil,
				ApproverOption:         nil,
				ApproverNeedSignReview: nil,
			},
			{
				Name:                   common.StringPtr("***"),
				IdCardType:             nil,
				IdCardNumber:           nil,
				Mobile:                 common.StringPtr("***********"),
				OrganizationName:       nil,
				NotChannelOrganization: nil,
				OpenId:                 nil,
				OrganizationOpenId:     nil,
				ApproverType:           common.StringPtr("PERSON"),
				RecipientId:            nil,
				Deadline:               nil,
				CallbackUrl:            nil,
				SignComponents: []*essbasic.Component{
					{
						ComponentId:           nil,
						ComponentType:         common.StringPtr("SIGN_SIGNATURE"),
						ComponentName:         nil,
						ComponentRequired:     nil,
						ComponentRecipientId:  nil,
						FileIndex:             nil,
						GenerateMode:          nil,
						ComponentWidth:        common.Float64Ptr(100),
						ComponentHeight:       common.Float64Ptr(100),
						ComponentPage:         common.Int64Ptr(1),
						ComponentPosX:         common.Float64Ptr(100),
						ComponentPosY:         common.Float64Ptr(200),
						ComponentExtra:        nil,
						ComponentValue:        nil,
						ComponentDateFontSize: nil,
						DocumentId:            nil,
						ComponentDescription:  nil,
						OffsetX:               nil,
						OffsetY:               nil,
						KeywordPage:           nil,
						RelativeLocation:      nil,
						KeywordIndexes:        nil,
					},
				},
				ComponentLimitType:     nil,
				PreReadTime:            nil,
				JumpUrl:                nil,
				ApproverOption:         nil,
				ApproverNeedSignReview: nil,
			},
		},
		Deadline:        common.Int64Ptr(0),
		FlowDescription: common.StringPtr("合同流程的描述"),
		FlowType:        nil,
		CallbackUrl:     nil,
		CustomerData:    nil,
		Unordered:       nil,
		CustomShowMap:   nil,
		NeedSignReview:  nil,
	}
}
