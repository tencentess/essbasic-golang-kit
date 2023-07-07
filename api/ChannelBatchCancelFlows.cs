using System;
using TencentCloud.Common;
using TencentCloud.Essbasic.V20210526;
using TencentCloud.Essbasic.V20210526.Models;

// ChannelBatchCancelFlows
// 指定需要批量撤销的签署流程Id，批量撤销合同
// 客户指定需要撤销的签署流程Id，最多100个，超过100不处理；接口失败后返回错误信息
// 注意:
// 能撤回合同的只能是合同的发起人或者发起企业的超管、法人
// 详细参考 https://cloud.tencent.com/document/api/1420/80391
namespace api
{
    class ChannelBatchCancelFlowsService
    {
        public static ChannelBatchCancelFlowsResponse ChannelBatchCancelFlows(Agent agent, String[] flowIds, 
            String cancelMessage, long cancelMessageFormat)
        {
            try
            {
                // 构造客户端调用实例
                EssbasicClient client = Common.GetClientInstance(Configs.secretId, Configs.secretKey, Configs.endPoint);

                // 实例化一个请求对象,每个接口都会对应一个request对象
                ChannelBatchCancelFlowsRequest req = new ChannelBatchCancelFlowsRequest();

                // 第三方平台应用相关信息
                // 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
                req.Agent = agent;

                // 签署流程Id数组，最多100个，超过100不处理
                req.FlowIds = flowIds;
                // 撤回原因，最大不超过200字符
                req.CancelMessage = cancelMessage;
                // 撤销理由自定义格式；选项：
                // 0 默认格式
                // 1 只保留身份信息：展示为【发起方】
                // 2 保留身份信息+企业名称：展示为【发起方xxx公司】
                // 3 保留身份信息+企业名称+经办人名称：展示为【发起方xxxx公司-经办人姓名】
                req.CancelMessageFormat = cancelMessageFormat;
                
                // 返回的resp是一个ChannelBatchCancelFlowsResponse的实例，与请求对象对应
                ChannelBatchCancelFlowsResponse resp = client.ChannelBatchCancelFlowsSync(req);
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
            String[] flowIds = { "************************" };
            String cancelMessage = "";
            long cancelMessageFormat = 0;
            ChannelBatchCancelFlowsResponse resp = ChannelBatchCancelFlows(Common.getAgent(),
                flowIds, cancelMessage, cancelMessageFormat);
            Console.WriteLine(AbstractModel.ToJsonString(resp));
        }
    }
}