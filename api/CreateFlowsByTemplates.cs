using System;
using System.Threading.Tasks;
using TencentCloud.Common;
using TencentCloud.Common.Profile;
using TencentCloud.Essbasic.V20210526;
using TencentCloud.Essbasic.V20210526.Models;

// CreateFlowsByTemplates
// 用于使用多个模板批量创建签署流程。当前可批量发起合同（签署流程）数量最大为20个。
// 如若在模板中配置了动态表格, 上传的附件必须为A4大小
// 合同发起人必须在电子签已经进行实名。
// 详细参考 https://cloud.tencent.com/document/api/1420/61523
namespace api
{
    class CreateFlowsByTemplatesService
    {
        public static CreateFlowsByTemplatesResponse CreateFlowsByTemplates(Agent agent, FlowInfo[] flowInfos)
        {
            try
            {
                // 构造客户端调用实例
                EssbasicClient client = Common.GetClientInstance(Configs.secretId, Configs.secretKey, Configs.endPoint);
                // 实例化一个请求对象,每个接口都会对应一个request对象
                CreateFlowsByTemplatesRequest req = new CreateFlowsByTemplatesRequest();

                // 第三方平台应用相关信息
                // 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
                req.Agent = agent;
                // 多个合同（签署流程）信息
                // 详细参考 https://cloud.tencent.com/document/api/1420/61525#FlowInfo
	            // 签署人 https://cloud.tencent.com/document/api/1420/61525#FlowApproverInfo
                req.FlowInfos = flowInfos;
                
                // 返回的resp是一个CreateFlowsByTemplatesResponse的实例，与请求对象对应
                CreateFlowsByTemplatesResponse resp = client.CreateFlowsByTemplatesSync(req);
                // 输出json格式的字符串回包
                Console.WriteLine(AbstractModel.ToJsonString(resp));
                return resp;
            }
            catch (Exception e)
            {
                Console.WriteLine(e.ToString());
            }
            return null;
        }
    }
}