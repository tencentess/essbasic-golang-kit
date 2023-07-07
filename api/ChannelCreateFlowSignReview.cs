using System;
using TencentCloud.Common;
using TencentCloud.Essbasic.V20210526;
using TencentCloud.Essbasic.V20210526.Models;

// ChannelCreateFlowSignReview
// 在通过接口(CreateFlowsByTemplates 或者ChannelCreateFlowByFiles)创建签署流程时
// 若指定了参数 NeedSignReview 为true,则可以调用此接口提交企业内部签署审批结果。
// 若签署流程状态正常，且本企业存在签署方未签署，同一签署流程可以多次提交签署审批结果，签署时的最后一个“审批结果”有效。
// 详细参考 https://cloud.tencent.com/document/api/1420/78953
namespace api
{
    class ChannelCreateFlowSignReviewService
    {
        public static ChannelCreateFlowSignReviewResponse ChannelCreateFlowSignReview(Agent agent, String flowId,
            String reviewType, String reviewMessage, String recipientId)
        {
            try
            {
                // 构造客户端调用实例
                EssbasicClient client = Common.GetClientInstance(Configs.secretId, Configs.secretKey, Configs.endPoint);

                // 实例化一个请求对象,每个接口都会对应一个request对象
                ChannelCreateFlowSignReviewRequest req = new ChannelCreateFlowSignReviewRequest();

                // 第三方平台应用相关信息
                // 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
                req.Agent = agent;

                // 签署流程编号
                req.FlowId = flowId;
                // 企业内部审核结果
                // PASS: 通过
                // REJECT: 拒绝
                // SIGN_REJECT:拒签(流程结束)
                req.ReviewType = reviewType;
                // 审核原因
	            // 当ReviewType 是REJECT 时此字段必填,字符串长度不超过200
                req.ReviewMessage = reviewMessage;
                // 签署节点审核时需要指定
                req.RecipientId = recipientId;
                
                // 返回的resp是一个ChannelCreateFlowSignReviewResponse的实例，与请求对象对应
                ChannelCreateFlowSignReviewResponse resp = client.ChannelCreateFlowSignReviewSync(req);
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
            String reviewType = "************************";
            String reviewMessage = "************************";           
            String recipientId = "************************";           
            ChannelCreateFlowSignReviewResponse resp = ChannelCreateFlowSignReview(Common.getAgent(), 
                flowId, reviewType,  reviewMessage,  recipientId);
            Console.WriteLine(AbstractModel.ToJsonString(resp));
        }
    }
}