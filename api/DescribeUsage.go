package api

import (
	"fmt"

	"essbasic-golang-kit_/utils"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	essbasic "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/essbasic/v20210526"
)

// DescribeUsage 此接口（DescribeUsage）用于获取第三方应用集成所有合作企业流量消耗情况。
// 注: 此接口每日限频2次，若要扩大限制次数,请提前与客服经理或邮件至e-contract@tencent.com进行联系。
// 详细参考 https://cloud.tencent.com/document/api/1420/61520
func DescribeUsage(agent *essbasic.Agent, startDate, endDate *string,
	needAggregate *bool, limit, offset *uint64) *essbasic.DescribeUsageResponse {
	// 实例化一个认证对象，入参需要传入腾讯云账户secretId，secretKey,此处还需注意密钥对的保密
	client, err := utils.InitClient()
	if err != nil {
		panic(err)
	}

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := essbasic.NewDescribeUsageRequest()

	// 应用信息，此接口Agent.AppId必填
	request.Agent = agent
	// 开始时间，例如：2021-03-21
	request.StartDate = startDate
	// 结束时间，例如：2021-06-21；
	// 开始时间到结束时间的区间长度小于等于90天。
	request.EndDate = endDate
	// 是否汇总数据，默认不汇总。
	// 不汇总：返回在统计区间内第三方应用集成下所有企业的每日明细，即每个企业N条数据，N为统计天数；
	// 汇总：返回在统计区间内第三方应用集成下所有企业的汇总后数据，即每个企业一条数据；
	request.NeedAggregate = needAggregate
	// 单次返回的最多条目数量。默认为1000，且不能超过1000。
	request.Limit = limit
	// 偏移量，默认是0。
	request.Offset = offset

	// 返回的resp是一个DescribeUsageResponse的实例，与请求对象对应
	response, err := client.DescribeUsage(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return nil
	}
	if err != nil {
		panic(err)
	}
	return response
}
