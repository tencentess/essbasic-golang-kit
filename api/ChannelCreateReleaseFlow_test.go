package api

import (
	"essbasic-golang-kit_/utils"
	"fmt"
	"testing"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	essbasic "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/essbasic/v20210526"
)

func TestChannelCreateReleaseFlow(t *testing.T) {
	agent := utils.SetAgent()

	needRelievedFlowId := "**********************"

	reliveInfo := &essbasic.RelieveInfo{
		Reason:                    common.StringPtr("**********************"),
		RemainInForceItem:         common.StringPtr("**********************"),
		OriginalExpenseSettlement: common.StringPtr("**********************"),
		OriginalOtherSettlement:   common.StringPtr("**********************"),
		OtherDeals:                common.StringPtr("**********************"),
	}

	releasedApprovers := []*essbasic.ReleasedApprover{
		{
			OrganizationName:   common.StringPtr("**********************"),
			ApproverNumber:     common.Uint64Ptr(0),
			ApproverType:       common.StringPtr("**********************"),
			Name:               common.StringPtr("**********************"),
			IdCardType:         common.StringPtr("**********************"),
			IdCardNumber:       common.StringPtr("**********************"),
			Mobile:             common.StringPtr("**********************"),
			OrganizationOpenId: common.StringPtr("**********************"),
			OpenId:             common.StringPtr("**********************"),
		},
	}

	callbackUrl := "**********************"

	response := ChannelCreateReleaseFlow(agent, &needRelievedFlowId, reliveInfo, releasedApprovers, &callbackUrl)
	// 输出json格式的字符串回包
	fmt.Printf("%s", response.ToJsonString())
}
