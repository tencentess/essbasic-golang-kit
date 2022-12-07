package api

import (
	"essbasic-golang-kit_/utils"
	"fmt"
	"testing"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	essbasic "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/essbasic/v20210526"
)

// 通过上传后的pdf资源编号来创建待签署的合同流程测试
func TestChannelCreateFlowByFiles(t *testing.T) {
	// 从UploadFiles接口获取到的fileId
	fileId := "**********************"
	// 签署流程名称,最大长度200个字符
	flowName := "我的第一份文件合同"
	// 渠道应用相关信息
	agent := utils.SetAgent()

	// 签署方参与信息
	// 个人签署方参数
	var flowApproverInfos = []*essbasic.FlowApproverInfo{
		{
			// 签署人类型，PERSON-个人；
			// ORGANIZATION-企业；
			// ENTERPRISESERVER-企业静默签;
			// 注：ENTERPRISESERVER 类型仅用于使用文件创建流程（ChannelCreateFlowByFiles）接口；并且仅能指定发起方企业签署方为静默签署；
			ApproverType: common.StringPtr("PERSON"),
			// 操作人的名字
			Name: common.StringPtr("*****"),
			// 操作人的手机号
			Mobile: common.StringPtr("**********"),
			//（签署）控件信息
			SignComponents: []*essbasic.Component{
				{
					// 参数控件X位置，单位px
					ComponentPosX: common.Float64Ptr(146.15625),
					// 参数控件Y位置，单位px
					ComponentPosY: common.Float64Ptr(472.78125),
					// 参数控件宽度，默认100，单位px，表单域和关键字转换控件不用填
					ComponentWidth: common.Float64Ptr(112),
					// 参数控件高度，默认100，单位px，表单域和关键字转换控件不用填
					ComponentHeight: common.Float64Ptr(40),
					// 控件所属文件的序号 (文档中文件的排列序号，从0开始)
					FileIndex: common.Int64Ptr(0),
					// 参数控件所在页码，从1开始
					ComponentPage: common.Int64Ptr(1),
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
					ComponentType: common.StringPtr("SIGN_SIGNATURE"),
					// 印章 ID，传参 DEFAULT_COMPANY_SEAL 表示使用默认印章。
					// 控件填入内容，印章控件里面，如果是手写签名内容为PNG图片格式的base64编码。
					ComponentValue: common.StringPtr(""),
				},
			},
		},
	}
	var fileIds = []*string{&fileId}
	response := ChannelCreateFlowByFiles(agent, flowName, fileIds, flowApproverInfos)
	// 输出json格式的字符串回包
	fmt.Printf("%s", response.ToJsonString())
}
