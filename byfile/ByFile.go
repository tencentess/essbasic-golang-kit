package main

import v20210526 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/essbasic/v20210526"

// BuildApprovers 构造签署人 - 以个人为例, 实际请根据自己的场景构造签署方、控件
func BuildApprovers() []*v20210526.FlowApproverInfo {

	// 个人签署方参数
	personName := "***"
	personMobile := "*********"

	// 企业签署方参数
	//organizationName := "*********"
	//organizationOpenId := "***************"
	//openId := "*********"
	var flowApproverInfos []*v20210526.FlowApproverInfo

	// 传入个人签署方
	flowApproverInfos = append(flowApproverInfos, BuildPersonApprover(&personName, &personMobile))

	// 传入企业签署方
	// flowApproverInfos = append(flowApproverInfos,
	//BuildOrganizationApprover(&organizationName, &organizationOpenId, &openId))

	// 传入企业静默签署
	// flowApproverInfos = append(flowApproverInfos, BuildServerSignApprover())

	return flowApproverInfos
}

// BuildPersonApprover 打包个人签署方参与者信息
func BuildPersonApprover(name, mobile *string) *v20210526.FlowApproverInfo {
	// 签署参与者信息
	// 个人签署方
	flowApproverInfo := &v20210526.FlowApproverInfo{}

	approverType := "PERSON"
	flowApproverInfo.ApproverType = &approverType

	flowApproverInfo.Name = name

	flowApproverInfo.Mobile = mobile

	var components []*v20210526.Component

	// 这里简单定义一个个人手写签名的签署控件
	component := BuildComponent(146.15625, 472.78125, 112,
		40, 0, 1, "SIGN_SIGNATURE", "")
	components = append(components, component)
	flowApproverInfo.SignComponents = components

	return flowApproverInfo
}

// BuildOrganizationApprover 打包企业签署方参与者信息
func BuildOrganizationApprover(organizationName, organizationOpenId, openId *string) *v20210526.FlowApproverInfo {
	// 签署参与者信息
	// 企业签署方
	flowApproverInfo := &v20210526.FlowApproverInfo{}

	approverType := "ORGANIZATION"
	flowApproverInfo.ApproverType = &approverType

	flowApproverInfo.OrganizationName = organizationName

	flowApproverInfo.OrganizationOpenId = organizationOpenId

	flowApproverInfo.OpenId = openId

	var components []*v20210526.Component

	// 这里简单定义一个个人手写签名的签署控件
	component := BuildComponent(146.15625, 472.78125, 112,
		40, 0, 1, "SIGN_SIGNATURE", "")
	components = append(components, component)
	flowApproverInfo.SignComponents = components

	return flowApproverInfo
}

// BuildServerSignApprover 打包企业静默签署方参与者信息
func BuildServerSignApprover() *v20210526.FlowApproverInfo {
	// 签署参与者信息
	// 企业静默签
	flowApproverInfo := &v20210526.FlowApproverInfo{}

	approverType := "ENTERPRISESERVER"
	flowApproverInfo.ApproverType = &approverType

	var components []*v20210526.Component

	// 这里简单定义一个个人手写签名的签署控件
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
