package api

import (
	"essbasic-golang-kit_/utils"
	"fmt"
	"testing"

	essbasic "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/essbasic/v20210526"
)

func TestChannelDescribeEmployees(t *testing.T) {
	agent := utils.SetAgent()

	// 查询实名用户
	//filters := []*essbasic.Filter{{
	//	Key:    common.StringPtr("Status"),
	//	Values: []*string{common.StringPtr("IsVerified")},
	//}}

	// 根据第三方系统openId过滤查询员工
	//filters := []*essbasic.Filter{{
	//	Key:    common.StringPtr("StaffOpenId"),
	//	Values: []*string{common.StringPtr("starkjin")},
	//}}

	// 查询离职员工
	//filters := []*essbasic.Filter{{
	//	Key:    common.StringPtr("Status"),
	//	Values: []*string{common.StringPtr("QuiteJob")},
	//}}

	filters := []*essbasic.Filter{{}}
	offset := int64(0)

	response := ChannelDescribeEmployees(agent, filters, &offset)
	// 输出json格式的字符串回包
	fmt.Printf("%s", response.ToJsonString())
}
