package main

import (
	"essbasic-golang-kit_/api"
	"essbasic-golang-kit_/config"
	"essbasic-golang-kit_/utils"

	v20210526 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/essbasic/v20210526"
)

// BuildApprovers 构造签署人 - 以BtoC为例, 实际请根据自己的场景构造签署方、控件
func BuildApprovers(recipients []*v20210526.Recipient) []*v20210526.FlowApproverInfo {

	// 个人签署方参数
	personName := "***"
	personMobile := "*********"

	// 企业签署方参数
	organizationName := "*********"
	organizationOpenId := config.ProxyOrganizationOpenId
	openId := config.ProxyOperatorOpenId
	var flowApproverInfos []*v20210526.FlowApproverInfo

	// 传入个人签署方
	flowApproverInfos = append(flowApproverInfos, BuildPersonApprover(&personName,
		&personMobile, recipients[0].RecipientId))

	// 传入企业签署方
	flowApproverInfos = append(flowApproverInfos, BuildOrganizationApprover(&organizationName,
		&organizationOpenId, &openId, recipients[0].RecipientId))

	// 传入企业静默签署，此处需要在config.php中设置一个持有的印章值serverSignSealId
	// flowApproverInfos = append(flowApproverInfos, BuildServerSignApprover())

	// 内容控件填充结构，详细说明参考
	// https://cloud.tencent.com/document/api/1420/61525#FormField

	return flowApproverInfos
}

// BuildPersonApprover 打包个人签署方参与者信息
func BuildPersonApprover(name, mobile, recipient *string) *v20210526.FlowApproverInfo {
	// 签署参与者信息
	// 个人签署方
	flowApproverInfo := &v20210526.FlowApproverInfo{}

	approverType := "PERSON"
	flowApproverInfo.ApproverType = &approverType

	flowApproverInfo.Name = name

	flowApproverInfo.Mobile = mobile

	// 模板中对应签署方的参与方id
	flowApproverInfo.RecipientId = recipient

	return flowApproverInfo
}

// BuildOrganizationApprover 打包企业签署方参与者信息
func BuildOrganizationApprover(
	organizationName, organizationOpenId, openId, recipient *string) *v20210526.FlowApproverInfo {
	// 签署参与者信息
	// 企业签署方
	flowApproverInfo := &v20210526.FlowApproverInfo{}

	approverType := "ORGANIZATION"
	flowApproverInfo.ApproverType = &approverType

	flowApproverInfo.OrganizationName = organizationName

	flowApproverInfo.OrganizationOpenId = organizationOpenId

	flowApproverInfo.OpenId = openId

	// 模板中对应签署方的参与方id
	flowApproverInfo.RecipientId = recipient

	return flowApproverInfo
}

// BuildServerSignApprover 打包企业静默签署方参与者信息
func BuildServerSignApprover() *v20210526.FlowApproverInfo {
	// 签署参与者信息
	// 企业静默签
	flowApproverInfo := &v20210526.FlowApproverInfo{}

	approverType := "ENTERPRISESERVER"

	// 注：此时发起方会替换为接口调用的企业+经办人，所以不需要传签署方信息
	flowApproverInfo.ApproverType = &approverType

	return flowApproverInfo
}

// BuildComponent 构建（签署）控件信息
func BuildComponent(componentPosX, componentPosY, componentWidth, componentHeight float64,
	fileIndex, componentPage int64, componentType, componentValue string) *v20210526.Component {
	var component = v20210526.Component{

		ComponentPosX: &componentPosX,

		ComponentPosY: &componentPosY,

		ComponentWidth: &componentWidth,

		ComponentHeight: &componentHeight,

		FileIndex: &fileIndex,

		ComponentPage: &componentPage,

		ComponentType:  &componentType,
		ComponentValue: &componentValue,
	}

	return &component
}

// GetRecipients 从模板中获取参与人信息，用于模板发起合同
func GetRecipients(templateId string) []*v20210526.Recipient {
	agent := utils.SetAgent()
	templatesResp := api.DescribeTemplates(agent, &templateId)
	return templatesResp.Response.Templates[0].Recipients
}
