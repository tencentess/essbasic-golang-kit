using System;
using TencentCloud.Common;
using TencentCloud.Essbasic.V20210526;
using TencentCloud.Essbasic.V20210526.Models;

// CreateChannelFlowEvidenceReport
// 创建出证报告，返回报告 ID
// 详细参考 https://cloud.tencent.com/document/api/1420/79688
namespace api
{
    class CreateChannelFlowEvidenceReportService
    {
        public static CreateChannelFlowEvidenceReportResponse CreateChannelFlowEvidenceReport(Agent agent, 
            String flowId)
        {
            try
            {
                // 构造客户端调用实例
                EssbasicClient client = Common.GetClientInstance(Configs.secretId, Configs.secretKey, Configs.endPoint);

                // 实例化一个请求对象,每个接口都会对应一个request对象
                CreateChannelFlowEvidenceReportRequest req = new CreateChannelFlowEvidenceReportRequest();

                // 第三方平台应用相关信息
                // 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填
                req.Agent = agent;
                // 签署流程编号，合同id
                req.FlowId = flowId;
                
                // 返回的resp是一个CreateChannelFlowEvidenceReportResponse的实例，与请求对象对应
                CreateChannelFlowEvidenceReportResponse resp = client.CreateChannelFlowEvidenceReportSync(req);
                // 输出json格式的字符串回包
                // Console.WriteLine(AbstractModel.ToJsonString(resp));
                return resp;
            }
            catch (Exception e)
            {
                Console.WriteLine(e.ToString());
            }
            return null;
        }

        public static void Main1(string[] args)
        {
            String flowId = "************************";       
            CreateChannelFlowEvidenceReportResponse resp = CreateChannelFlowEvidenceReport(Common.getAgent(), flowId);
            Console.WriteLine(AbstractModel.ToJsonString(resp));
        }
    }
}