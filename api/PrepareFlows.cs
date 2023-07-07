using System;
using TencentCloud.Common;
using TencentCloud.Essbasic.V20210526;
using TencentCloud.Essbasic.V20210526.Models;

// PrepareFlows
// 该接口 (PrepareFlows) 用于创建待发起文件
// 用户通过该接口进入签署流程发起的确认页面，进行发起信息二次确认， 如果确认则进行正常发起。
// 目前该接口只支持B2C，不建议使用，将会废弃。
// 详细参考 https://cloud.tencent.com/document/api/1420/61519
namespace api
{
    class PrepareFlowsService
    {
        public static PrepareFlowsResponse PrepareFlows(Agent agent, String jumpUrl, FlowInfo[] flowInfos)
        {
            try
            {
                // 构造客户端调用实例
                EssbasicClient client = Common.GetClientInstance(Configs.secretId, Configs.secretKey, Configs.endPoint);
                
                // 实例化一个请求对象,每个接口都会对应一个request对象
                PrepareFlowsRequest req = new PrepareFlowsRequest();

                // 第三方平台应用相关信息
                // 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
                req.Agent = agent;
                // 多个合同（签署流程）信息，最大支持20个签署流程。
                req.JumpUrl = jumpUrl;
                // 操作完成后的跳转地址，最大长度200
                req.FlowInfos = flowInfos;
                
                // 返回的resp是一个PrepareFlowsResponse的实例，与请求对象对应
                PrepareFlowsResponse resp = client.PrepareFlowsSync(req);
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
            String jumpUrl = "************************";

            FlowInfo flowInfo = new FlowInfo();
            flowInfo.FlowName = "***";
            flowInfo.TemplateId = "************************";

            FlowApproverInfo flowApprover = new FlowApproverInfo();
            flowApprover.Name = "***";
            flowApprover.Mobile = "***********";
            flowInfo.FlowApprovers = new[] { flowApprover };
            
            FlowInfo[] flowInfos = { flowInfo };
            PrepareFlowsResponse resp = PrepareFlows(Common.getAgent(), 
                jumpUrl, flowInfos);
            Console.WriteLine(AbstractModel.ToJsonString(resp));
        }
    }
}