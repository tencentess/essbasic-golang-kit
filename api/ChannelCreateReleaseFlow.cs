using System;
using TencentCloud.Common;
using TencentCloud.Essbasic.V20210526;
using TencentCloud.Essbasic.V20210526.Models;

// ChannelCreateReleaseFlow
// 第三方应用集成发起解除协议，主要应用场景为：基于一份已经签署的合同，进行解除操作。
// 合同发起人必须在电子签已经进行实名。
// 详细参考 https://cloud.tencent.com/document/api/1420/83461
namespace api
{
    class ChannelCreateReleaseFlowService
    {
        public static ChannelCreateReleaseFlowResponse ChannelCreateReleaseFlow(Agent agent, String needRelievedFlowId, 
            RelieveInfo reliveInfo, ReleasedApprover[] releasedApprovers, String callbackUrl)
        {
            try
            {
                // 构造客户端调用实例
                EssbasicClient client = Common.GetClientInstance(Configs.secretId, Configs.secretKey, Configs.endPoint);

                // 实例化一个请求对象,每个接口都会对应一个request对象
                ChannelCreateReleaseFlowRequest req = new ChannelCreateReleaseFlowRequest();

                // 第三方平台应用相关信息
                // 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
                req.Agent = agent;

                // 待解除的流程编号（即原流程的编号）
                req.NeedRelievedFlowId = needRelievedFlowId;
                // 解除协议内容，其中Reason必填
                // 详细参考 https://cloud.tencent.com/document/api/1420/61525#RelieveInfo
                req.ReliveInfo = reliveInfo;
                // 非必须，解除协议的本企业签署人列表，默认使用原流程的签署人列表；
                // 当解除协议的签署人与原流程的签署人不能相同时（例如原流程签署人离职了），需要指定本企业的其他签署人来替换原流程中的原签署人，
                // 注意需要指明ApproverNumber来代表需要替换哪一个签署人，解除协议的签署人数量不能多于原流程的签署人数量
                req.ReleasedApprovers = releasedApprovers;
                // 签署完回调url，最大长度1000个字符
                req.CallbackUrl = callbackUrl;
                
                // 返回的resp是一个ChannelCreateReleaseFlowResponse的实例，与请求对象对应
                ChannelCreateReleaseFlowResponse resp = client.ChannelCreateReleaseFlowSync(req);
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
            String needRelievedFlowId = "************************";
            
            RelieveInfo reliveInfo = new RelieveInfo();
            reliveInfo.Reason = "************************";
            reliveInfo.RemainInForceItem = "************************";
            reliveInfo.OriginalExpenseSettlement = "************************";
            reliveInfo.OriginalOtherSettlement = "************************";
            reliveInfo.OtherDeals = "************************";

            ReleasedApprover releasedApprover = new ReleasedApprover();
            releasedApprover.OrganizationName = "************************";
            releasedApprover.ApproverNumber = 0;
            releasedApprover.ApproverType = "************************";
            releasedApprover.Name = "************************";
            releasedApprover.IdCardType = "************************";
            releasedApprover.IdCardNumber = "************************";
            releasedApprover.Mobile = "************************";
            releasedApprover.OrganizationOpenId = "************************";
            releasedApprover.OpenId = "************************";
            
            ReleasedApprover[] releasedApprovers = { releasedApprover };
            String callbackUrl = "************************";
            ChannelCreateReleaseFlowResponse resp = ChannelCreateReleaseFlow(Common.getAgent(),
                needRelievedFlowId, reliveInfo, releasedApprovers, callbackUrl);
            Console.WriteLine(AbstractModel.ToJsonString(resp));
        }
    }
}