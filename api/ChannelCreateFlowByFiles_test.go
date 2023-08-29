package api

import (
	"fmt"
	"testing"

	"essbasic-golang-kit_/utils"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	essbasic "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/essbasic/v20210526"
)

// 通过上传后的pdf资源编号来创建待签署的合同流程测试
func TestChannelCreateFlowByFiles(t *testing.T) {

	fileId := "**********************"

	flowName := "我的第一份文件合同"

	agent := utils.SetAgent()

	// 签署方参与信息
	// 个人签署方参数
	var flowApproverInfos = []*essbasic.FlowApproverInfo{
		{
			ApproverType: common.StringPtr("PERSON"),

			Name: common.StringPtr("*****"),

			Mobile: common.StringPtr("**********"),
			//（使用PDF文件直接发起合同时，签署人指定的签署控件
			SignComponents: []*essbasic.Component{
				{
					ComponentPosX: common.Float64Ptr(146.15625),

					ComponentPosY: common.Float64Ptr(472.78125),

					ComponentWidth: common.Float64Ptr(112),

					ComponentHeight: common.Float64Ptr(40),

					FileIndex: common.Int64Ptr(0),

					ComponentPage: common.Int64Ptr(1),

					ComponentType: common.StringPtr("SIGN_SIGNATURE"),
	
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
