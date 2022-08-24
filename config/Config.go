package config

const (
	// SecretId SecretKey 调用API的密钥对，通过腾讯云控制台获取
	SecretId  = "****************"
	SecretKey = "****************"

	AppId = "****************"

	// ProxyAppId 腾讯电子签颁发给渠道侧合作企业的应用ID
	ProxyAppId = "****************"

	// ProxyOrganizationOpenId 渠道/平台合作企业的企业ID
	ProxyOrganizationOpenId = "****************"

	// ProxyOperatorOpenId 渠道/平台合作企业经办人（操作员）ID
	ProxyOperatorOpenId = "****************"

	// ServerSignSealId 企业方静默签用的印章Id，电子签控制台获取
	ServerSignSealId = "****************"

	// EndPoint API域名，现网使用 ess.tencentcloudapi.com
	EndPoint = "essbasic.test.ess.tencent.cn"

	// FileServiceEndPoint 文件服务域名，现网使用 file.ess.tencent.cn
	FileServiceEndPoint = "file.test.ess.tencent.cn"

	// TemplateId 模板ID
	TemplateId = "****************"

	// COUNT 批量发起时数量设置
	COUNT = 1
)
