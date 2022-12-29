package config

const (
	// SecretId SecretKey 调用云API需要使用的密钥对，是渠道平台企业的密钥对（子客企业无需准备此参数）。
	// 联调/生产环境都是可以由渠道企业电子签控制台应用集成模块自主获取。
	// https://cloud.tencent.com/document/product/1323/83102#.E6.B8.A0.E9.81.93.E7.89.88
	SecretId  = "****************"
	SecretKey = "****************"

	// 应用的唯一标识，对应通用参数AppId。不同的业务系统可以采用不同的AppId，
	// 不同AppId下的数据是隔离的。联调/生产环境都是可以由控制台开发者中心-应用集成自主生成。
	// https://cloud.tencent.com/document/product/1323/83102#.E6.B8.A0.E9.81.93.E7.89.88
	AppId = "****************"

	// 在子客企业开通电子签后，会生成唯一的子客应用Id（ProxyAppId）用于代理调用时的鉴权
	// 在回调与跳转-电子签开通通知里会返回此数据ProxyAppId
	ProxyAppId = "****************"

	// 道子客企业OpenId（OrganizationOpenId）：简称OrgOpenId，对应通用参数ProxyOrganizationOpenId。
	// 是渠道平台自定义，对于渠道子客企业的唯一标识。
	// 一个渠道子客企业主体与子客企业ProxyOrganizationOpenId是一一对应的，不可更改，不可重复使用。
	// （比如，可以使用企业名称的hash值，或者社会统一信用代码的hash值，或者随机hash值，需要渠道平台保存）
	// 最大64位字符串
	ProxyOrganizationOpenId = "****************"

	// 渠道子客企业员工/经办人OpenId（OperatorOpenId）：简称员工/经办人OpenId。对应通用ProxyOperator中的参数OpenId。
	// 是渠道平台自定义，对渠道子客企业员的唯一标识。
	// 一个OpenId在一个子客企业内唯一对应一个真实员工，不可在其他子客企业内重复使用。
	// （比如，可以使用经办人企业名+员工身份证的hash值，需要渠道平台保存）
	// 最大64位字符串
	ProxyOperatorOpenId = "****************"

	// ServerSignSealId 企业方静默签用的印章Id，电子签控制台印章模块获取
	ServerSignSealId = "****************"

	// EndPoint API域名，现网使用 ess.tencentcloudapi.com
	EndPoint = "essbasic.test.ess.tencent.cn"

	// FileServiceEndPoint 文件服务域名，现网使用 file.ess.tencent.cn
	// UploadFiles 接口使用此域名进行调用
	FileServiceEndPoint = "file.test.ess.tencent.cn"

	// TemplateId 模板ID，电子签控制台模版模块获取，仅在通过模板发起时使用
	TemplateId = "****************"

	// COUNT 批量发起时数量设置
	COUNT = 1
)
