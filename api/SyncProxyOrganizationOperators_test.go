package api

import (
	"essbasic-golang-kit_/utils"
	"fmt"
	"testing"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	essbasic "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/essbasic/v20210526"
)

func TestSyncProxyOrganizationOperators(t *testing.T) {
	agent := utils.SetAgent()

	operatorType := "********"

	proxyOrganizationOperators := []*essbasic.ProxyOrganizationOperator{
		{
			Id:           common.StringPtr("*******"),
			Name:         common.StringPtr("*******"),
			IdCardType:   common.StringPtr("*******"),
			IdCardNumber: common.StringPtr("*******"),
			Mobile:       common.StringPtr("*******"),
		},
	}

	response := SyncProxyOrganizationOperators(agent, &operatorType, proxyOrganizationOperators)
	// 输出json格式的字符串回包
	fmt.Printf("%s", response.ToJsonString())
}
