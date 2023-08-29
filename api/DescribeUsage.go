package api

import (
	"fmt"

	"essbasic-golang-kit_/utils"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	essbasic "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/essbasic/v20210526"
)

func DescribeUsage(agent *essbasic.Agent, startDate, endDate *string,
	needAggregate *bool, limit, offset *uint64) *essbasic.DescribeUsageResponse {
	// 实例化一个认证对象，入参需要传入腾讯云账户secretId，secretKey,此处还需注意密钥对的保密
	client, err := utils.InitClient()
	if err != nil {
		panic(err)
	}

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := essbasic.NewDescribeUsageRequest()


	request.Agent = agent

	request.StartDate = startDate

	request.EndDate = endDate

	request.NeedAggregate = needAggregate

	request.Limit = limit

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
