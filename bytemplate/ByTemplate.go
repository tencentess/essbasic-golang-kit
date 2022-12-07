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
	flowApproverInfos = append(flowApproverInfos, BuildPersonApprover(&personName, &personMobile, recipients[0].RecipientId))
	// 传入企业签署方
	flowApproverInfos = append(flowApproverInfos, BuildOrganizationApprover(&organizationName, &organizationOpenId, &openId, recipients[0].RecipientId))
	// 传入企业静默签署
	// flowApproverInfos = append(flowApproverInfos, BuildServerSignApprover())

	return flowApproverInfos
}

// BuildPersonApprover 打包个人签署方参与者信息
func BuildPersonApprover(name, mobile, recipient *string) *v20210526.FlowApproverInfo {
	// 签署参与者信息
	// 个人签署方
	flowApproverInfo := &v20210526.FlowApproverInfo{}
	// 签署人类型，PERSON-个人；
	// ORGANIZATION-企业；
	// ENTERPRISESERVER-企业静默签;
	// 注：ENTERPRISESERVER 类型仅用于使用文件创建流程（ChannelCreateFlowByFiles）接口；并且仅能指定发起方企业签署方为静默签署；
	approverType := "PERSON"
	flowApproverInfo.ApproverType = &approverType
	// 本环节需要操作人的名字
	flowApproverInfo.Name = name
	// 本环节需要操作人的手机号
	flowApproverInfo.Mobile = mobile
	flowApproverInfo.RecipientId = recipient
	var components []*v20210526.Component

	component := BuildComponent(146.15625, 472.78125, 112,
		40, 0, 1, "SIGN_SIGNATURE", "")

	components = append(components, component)

	flowApproverInfo.SignComponents = components
	return flowApproverInfo
}

// BuildOrganizationApprover 打包企业签署方参与者信息
func BuildOrganizationApprover(organizationName, organizationOpenId, openId, recipient *string) *v20210526.FlowApproverInfo {
	// 签署参与者信息
	// 个人签署方
	flowApproverInfo := &v20210526.FlowApproverInfo{}
	// 签署人类型，PERSON-个人；
	// ORGANIZATION-企业；
	// ENTERPRISESERVER-企业静默签;
	// 注：ENTERPRISESERVER 类型仅用于使用文件创建流程（ChannelCreateFlowByFiles）接口；并且仅能指定发起方企业签署方为静默签署；
	approverType := "ORGANIZATION"
	flowApproverInfo.ApproverType = &approverType
	// 本环节需要企业操作人的企业名称
	flowApproverInfo.OrganizationName = organizationName
	// 本环节需要企业的OpenId
	flowApproverInfo.OrganizationOpenId = organizationOpenId
	// 本环节需要操作人的OpenId
	flowApproverInfo.OpenId = openId
	flowApproverInfo.RecipientId = recipient
	var components []*v20210526.Component

	component := BuildComponent(146.15625, 472.78125, 112,
		40, 0, 1, "SIGN_SIGNATURE", "")

	components = append(components, component)

	flowApproverInfo.SignComponents = components
	return flowApproverInfo
}

// BuildServerSignApprover 打包企业静默签署方参与者信息
func BuildServerSignApprover() *v20210526.FlowApproverInfo {
	// 签署参与者信息
	// 个人签署方
	flowApproverInfo := &v20210526.FlowApproverInfo{}
	// 签署人类型，PERSON-个人；
	// ORGANIZATION-企业；
	// ENTERPRISESERVER-企业静默签;
	// 注：ENTERPRISESERVER 类型仅用于使用文件创建流程（ChannelCreateFlowByFiles）接口；并且仅能指定发起方企业签署方为静默签署；
	approverType := "ENTERPRISESERVER"
	flowApproverInfo.ApproverType = &approverType
	var components []*v20210526.Component

	component := BuildComponent(146.15625, 472.78125, 112,
		40, 0, 1, "SIGN_SIGNATURE", "")

	components = append(components, component)

	flowApproverInfo.SignComponents = components
	return flowApproverInfo
}

// BuildComponent 构建（签署）控件信息
func BuildComponent(componentPosX, componentPosY, componentWidth, componentHeight float64,
	fileIndex, componentPage int64, componentType, componentValue string) *v20210526.Component {
	var component = v20210526.Component{
		// 参数控件X位置，单位px
		ComponentPosX: &componentPosX,
		// 参数控件Y位置，单位px
		ComponentPosY: &componentPosY,
		// 参数控件宽度，默认100，单位px，表单域和关键字转换控件不用填
		ComponentWidth: &componentWidth,
		// 参数控件高度，默认100，单位px，表单域和关键字转换控件不用填
		ComponentHeight: &componentHeight,
		// 控件所属文件的序号 (文档中文件的排列序号，从0开始)
		FileIndex: &fileIndex,
		// 参数控件所在页码，从1开始
		ComponentPage: &componentPage,
		// 如果是Component控件类型，则可选的字段为：
		//TEXT - 普通文本控件；
		//DATE - 普通日期控件；跟TEXT相比会有校验逻辑
		//DYNAMIC_TABLE- 动态表格控件
		//如果是SignComponent控件类型，则可选的字段为
		//SIGN_SEAL - 签署印章控件；
		//SIGN_DATE - 签署日期控件；
		//SIGN_SIGNATURE - 用户签名控件；
		//SIGN_PERSONAL_SEAL - 个人签署印章控件；
		//表单域的控件不能作为印章和签名控件
		ComponentType: &componentType,
		// 印章 ID，传参 DEFAULT_COMPANY_SEAL 表示使用默认印章。
		// 控件填入内容，印章控件里面，如果是手写签名内容为PNG图片格式的base64编码。
		ComponentValue: &componentValue,
	}
	return &component
}

func GetRecipients(templateId string) []*v20210526.Recipient {
	agent := utils.SetAgent()
	templatesResp := api.DescribeTemplates(agent, &templateId)
	return templatesResp.Response.Templates[0].Recipients
}
