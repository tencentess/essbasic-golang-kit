using System;
using TencentCloud.Common;
using TencentCloud.Essbasic.V20210526;
using TencentCloud.Essbasic.V20210526.Models;

// ChannelCreateBatchCancelFlowUrl
// 指定需要批量撤销的签署流程Id，获取批量撤销链接
// 客户指定需要撤销的签署流程Id，最多100个，超过100不处理；
// 接口调用成功返回批量撤销合同的链接，通过链接跳转到电子签小程序完成批量撤销;
// 可以撤回：未全部签署完成；不可以撤回（终态）：已全部签署完成、已拒签、已过期、已撤回。
// 注意:
// 能撤回合同的只能是合同的发起人或者发起企业的超管、法人
// 详细参考 https://cloud.tencent.com/document/api/1420/78264
namespace api
{
    class ChannelCreateBatchCancelFlowUrlService
    {
        public static ChannelCreateBatchCancelFlowUrlResponse ChannelCreateBatchCancelFlowUrl(Agent agent, String[] flowIds)
        {
            try
            {
                // 构造客户端调用实例
                EssbasicClient client = Common.GetClientInstance(Configs.secretId, Configs.secretKey, Configs.endPoint);

                // 实例化一个请求对象,每个接口都会对应一个request对象
                ChannelCreateBatchCancelFlowUrlRequest req = new ChannelCreateBatchCancelFlowUrlRequest();

                // 第三方平台应用相关信息
                // 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
                req.Agent = agent;
                // 签署流程Id数组
                req.FlowIds = flowIds;
                
                // 返回的resp是一个ChannelCreateBatchCancelFlowUrlResponse的实例，与请求对象对应
                ChannelCreateBatchCancelFlowUrlResponse resp = client.ChannelCreateBatchCancelFlowUrlSync(req);
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
            ChannelCreateBatchCancelFlowUrlResponse resp = ChannelCreateBatchCancelFlowUrl(Common.getAgent(),
                flowIds);
            Console.WriteLine(AbstractModel.ToJsonString(resp));
        }
    }
}