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
	// flowApproverInfos = append(flowApproverInfos, BuildOrganizationApprover(&organizationName, &organizationOpenId, &openId))

	// 传入企业静默签署
	// flowApproverInfos = append(flowApproverInfos, BuildServerSignApprover())

	return flowApproverInfos
}

// BuildPersonApprover 打包个人签署方参与者信息
func BuildPersonApprover(name, mobile *string) *v20210526.FlowApproverInfo {
	// 签署参与者信息
	// 个人签署方
	flowApproverInfo := &v20210526.FlowApproverInfo{}

	// 签署人类型
	// PERSON-个人/自然人；
	// ORGANIZATION-企业（企业签署方或模版发起时的企业静默签）；
	// ENTERPRISESERVER-企业静默签（文件发起时的企业静默签字）。
	approverType := "PERSON"
	flowApproverInfo.ApproverType = &approverType

	// 签署人姓名，最大长度50个字符
	flowApproverInfo.Name = name
	// 签署人手机号，脱敏显示。大陆手机号为11位，暂不支持海外手机号
	flowApproverInfo.Mobile = mobile

	// 控件，包括填充控件、签署控件，具体查看
	// https://cloud.tencent.com/document/api/1420/61525#Component
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

	// 签署人类型
	// PERSON-个人/自然人；
	// ORGANIZATION-企业（企业签署方或模版发起时的企业静默签）；
	// ENTERPRISESERVER-企业静默签（文件发起时的企业静默签字）。
	approverType := "ORGANIZATION"
	flowApproverInfo.ApproverType = &approverType

	// 企业签署方工商营业执照上的企业名称，签署方为非发起方企业场景下必传，最大长度64个字符；
	flowApproverInfo.OrganizationName = organizationName

	// 如果签署方是子客企业，此处需要传子客企业的OrganizationOpenId
	// 企业签署方在同一渠道下的其他合作企业OpenId，签署方为非发起方企业场景下必传，最大长度64个字符；
	flowApproverInfo.OrganizationOpenId = organizationOpenId
	// 如果签署方是子客企业，此处需要传子客企业经办人的OpenId
	// 当签署方为同一渠道下的员工时，该字段若不指定，则发起【待领取】的流程
	flowApproverInfo.OpenId = openId

	// 控件，包括填充控件、签署控件，具体查看
	// https://cloud.tencent.com/document/api/1420/61525#Component
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

	// 签署人类型
	// PERSON-个人/自然人；
	// ORGANIZATION-企业（企业签署方或模版发起时的企业静默签）；
	// ENTERPRISESERVER-企业静默签（文件发起时的企业静默签字）。
	approverType := "ENTERPRISESERVER"
	flowApproverInfo.ApproverType = &approverType

	// 控件，包括填充控件、签署控件，具体查看
	// https://cloud.tencent.com/document/api/1420/61525#Component
	var components []*v20210526.Component

	// 这里简单定义一个个人手写签名的签署控件
	component := BuildComponent(146.15625, 472.78125, 112,
		40, 0, 1, "SIGN_SIGNATURE", "")
	components = append(components, component)
	flowApproverInfo.SignComponents = components

	return flowApproverInfo
}

// BuildComponent 构建（签署）控件信息
// 详细参考 https://cloud.tencent.com/document/api/1420/61525#Component

// 在通过文件发起合同时，对应的component有三种定位方式
// 绝对定位方式
// 表单域(FIELD)定位方式
// 关键字(KEYWORD)定位方式
// 可以参考官网说明
// https://cloud.tencent.com/document/product/1323/78346#component-.E4.B8.89.E7.A7.8D.E5.AE.9A.E4.BD.8D.E6.96.B9.E5.BC.8F.E8.AF.B4.E6.98.8E
func BuildComponent(componentPosX, componentPosY, componentWidth, componentHeight float64,
	fileIndex, componentPage int64, componentType, componentValue string) *v20210526.Component {
	var component = v20210526.Component{
		// 位置信息 包括：
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

		// 控件类型与对应值，这里以官网说明为准
		// https://cloud.tencent.com/document/api/1420/61525#Component
		ComponentType: &componentType,
		ComponentValue: &componentValue,
	}
	
	return &component
}
