using System;
using TencentCloud.Common;
using TencentCloud.Essbasic.V20210526;
using TencentCloud.Essbasic.V20210526.Models;

// ChannelVerifyPdf
// 合同文件验签
// 详细参考 https://cloud.tencent.com/document/api/1420/80799
namespace api
{
    class ChannelVerifyPdfService
    {
        public static ChannelVerifyPdfResponse ChannelVerifyPdf(Agent agent, 
            String flowId)
        {
            try
            {
                // 构造客户端调用实例
                EssbasicClient client = Common.GetClientInstance(Configs.secretId, Configs.secretKey, Configs.endPoint);

                // 实例化一个请求对象,每个接口都会对应一个request对象
                ChannelVerifyPdfRequest req = new ChannelVerifyPdfRequest();

                // 第三方平台应用相关信息
                // 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
                req.Agent = agent;
                // 合同Id，流程Id
                req.FlowId = flowId;
                
                // 返回的resp是一个ChannelVerifyPdfResponse的实例，与请求对象对应
                ChannelVerifyPdfResponse resp = client.ChannelVerifyPdfSync(req);
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
            ChannelVerifyPdfResponse resp = ChannelVerifyPdf(Common.getAgent(), flowId);
            Console.WriteLine(AbstractModel.ToJsonString(resp));
        }
    }
}