package api

import (
	"essbasic-golang-kit_/utils"
	"fmt"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	essbasic "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/essbasic/v20210526"
)

// OperateChannelTemplate
// 此接口（OperateChannelTemplate）用于针对渠道模板库中的模板对子客企业可见性的查询和设置，不会直接分配渠道模板给子客企业。
// 1、OperateType=select时：
// 查询渠道模板库
// 2、OperateType=update或者delete时：
// 对子客企业进行模板库中模板可见性的修改、删除操作。
func OperateChannelTemplate(agent *essbasic.Agent, operateType, templateId,
	proxyOrganizationOpenIds, authTag *string) *essbasic.OperateChannelTemplateResponse {
	// 实例化一个认证对象，入参需要传入腾讯云账户secretId，secretKey,此处还需注意密钥对的保密
	client, err := utils.InitClient()
	if err != nil {
		panic(err)
	}

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := essbasic.NewOperateChannelTemplateRequest()

	// 渠道应用相关信息
	request.Agent = agent
	// 操作类型，查询:"SELECT"，删除:"DELETE"，更新:"UPDATE"
	request.OperateType = operateType
	// 渠道方模板库模板唯一标识
	request.TemplateId = templateId
	// 合作企业方第三方机构唯一标识数据，支持多个， 用","进行分隔
	request.ProxyOrganizationOpenIds = proxyOrganizationOpenIds
	// 模板可见性, 全部可见-"all", 部分可见-"part"
	request.AuthTag = authTag

	// 返回的resp是一个OperateChannelTemplateResponse的实例，与请求对象对应
	response, err := client.OperateChannelTemplate(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return nil
	}
	if err != nil {
		panic(err)
	}
	return response
}
