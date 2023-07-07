using System;
using TencentCloud.Common;
using TencentCloud.Essbasic.V20210526;
using TencentCloud.Essbasic.V20210526.Models;

// ChannelCreateBoundFlows
// 此接口（ChannelCreateBoundFlows）用于子客领取合同，经办人需要有相应的角色，领取后的合同不能重复领取。
// 详细参考 https://cloud.tencent.com/document/api/1420/83118
namespace api
{
    class ChannelCreateBoundFlowsService
    {
        public static ChannelCreateBoundFlowsResponse ChannelCreateBoundFlows(Agent agent, String[] flowIds)
        {
            try
            {
                // 构造客户端调用实例
                EssbasicClient client = Common.GetClientInstance(Configs.secretId, Configs.secretKey, Configs.endPoint);

                // 实例化一个请求对象,每个接口都会对应一个request对象
                ChannelCreateBoundFlowsRequest req = new ChannelCreateBoundFlowsRequest();

                // 第三方平台应用相关信息
                // 此接口Agent.AppId、Agent.ProxyOrganizationOpenId 和 Agent. ProxyOperator.OpenId 必填
                req.Agent = agent;
                // 领取的合同id列表
                req.FlowIds = flowIds;
                
                // 返回的resp是一个ChannelCreateBoundFlowsResponse的实例，与请求对象对应
                ChannelCreateBoundFlowsResponse resp = client.ChannelCreateBoundFlowsSync(req);
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
            String[] flowIds = new[] { "************************" };
            ChannelCreateBoundFlowsResponse resp = ChannelCreateBoundFlows(Common.getAgent(), flowIds);
            Console.WriteLine(AbstractModel.ToJsonString(resp));
        }
    }
}