package utils

import (
	"encoding/base64"
	"essbasic-golang-kit_/config"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	essbasic "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/essbasic/v20210526"
	"io/ioutil"
)

// InitClient 初始化Client
func InitClient() (*essbasic.Client, error) {
	// 实例化一个认证对象，入参需要传入腾讯云账户secretId，secretKey,此处还需注意密钥对的保密
	// 密钥可前往https://console.cloud.tencent.com/cam/capi网站进行获取
	credential := common.NewCredential(
		config.SecretId,
		config.SecretKey,
	)
	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = config.EndPoint
	// 实例化要请求产品的client对象,clientProfile是可选的
	client, err := essbasic.NewClient(credential, "", cpf)
	if err != nil {
		return nil, err
	}
	return client, err
}

// SetAgent 设置Agent
func SetAgent() *essbasic.Agent {
	appId := config.AppId
	proxyAppId := config.ProxyAppId
	proxyOrganizationOpenId := config.ProxyOrganizationOpenId
	proxyOperatorOpenId := config.ProxyOperatorOpenId
	userInfo := &essbasic.UserInfo{}
	userInfo.OpenId = &proxyOperatorOpenId
	var agent = &essbasic.Agent{
		AppId:                   &appId,
		ProxyAppId:              &proxyAppId,
		ProxyOrganizationOpenId: &proxyOrganizationOpenId,
		ProxyOperator:           userInfo,
	}
	return agent
}

// FillFlowInfo 设置FlowInfo
func FillFlowInfo(TemplateId, FlowName string, flowApproverInfos []*essbasic.FlowApproverInfo) *essbasic.FlowInfo {
	FlowType := "合同"
	var flowInfo = &essbasic.FlowInfo{
		TemplateId:    &TemplateId,
		FlowName:      &FlowName,
		FlowApprovers: flowApproverInfos,
		FlowType:      &FlowType,
	}
	return flowInfo
}

// ConvertImageFileToBase64 将图片文件转化为字符串，并对其进行Base64编码处理
func ConvertImageFileToBase64(filePath string) string {
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return ""
	}
	bs64 := base64.StdEncoding.EncodeToString(fileBytes)
	return bs64
}

func BuildFormField(componentName, componentValue string) *essbasic.FormField {
	return &essbasic.FormField{
		ComponentName:  common.StringPtr(componentName),
		ComponentValue: common.StringPtr(componentValue),
	}
}
