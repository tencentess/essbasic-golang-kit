package api

import (
	"essbasic-golang-kit_/utils"
	"fmt"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	essbasic "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/essbasic/v20210526"
)

// ChannelDescribeOrganizationSeals
// 查询渠道子客企业电子印章，需要操作者具有管理印章权限
// 客户指定需要获取的印章数量和偏移量，数量最多100，超过100按100处理；
// 入参InfoType控制印章是否携带授权人信息，为1则携带，为0则返回的授权人信息为空数组。
// 接口调用成功返回印章的信息列表还有企业印章的总数。
func ChannelDescribeOrganizationSeals(agent *essbasic.Agent, limit, offset, infoType *int64,
	sealId *string) *essbasic.ChannelDescribeOrganizationSealsResponse {
	// 实例化一个认证对象，入参需要传入腾讯云账户secretId，secretKey,此处还需注意密钥对的保密
	client, err := utils.InitClient()
	if err != nil {
		panic(err)
	}

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := essbasic.NewChannelDescribeOrganizationSealsRequest()

	// 渠道应用相关信息
	request.Agent = agent
	// 返回最大数量，最大为100
	request.Limit = limit
	// 偏移量，默认为0，最大为20000
	request.Offset = offset
	// 查询信息类型，为0时不返回授权用户，为1时返回
	request.InfoType = infoType
	// 印章id（没有输入返回所有）
	request.SealId = sealId

	// 返回的resp是一个ChannelDescribeOrganizationSealsResponse的实例，与请求对象对应
	response, err := client.ChannelDescribeOrganizationSeals(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return nil
	}
	if err != nil {
		panic(err)
	}
	return response
}
