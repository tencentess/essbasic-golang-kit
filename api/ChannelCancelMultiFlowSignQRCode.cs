using System;
using System.Threading.Tasks;
using TencentCloud.Common;
using TencentCloud.Common.Profile;
using TencentCloud.Essbasic.V20210526;
using TencentCloud.Essbasic.V20210526.Models;

// ChannelCancelMultiFlowSignQRCode 
// 用于取消一码多扫二维码。该接口对传入的二维码ID，若还在有效期内，可以提前失效
// 详细参考 https://cloud.tencent.com/document/api/1420/75453
namespace api
{
    class ChannelCancelMultiFlowSignQRCodeService
    {
        static ChannelCancelMultiFlowSignQRCodeResponse ChannelCancelMultiFlowSignQRCode(Agent agent,String qrCodeId)
        {
            try
            {
                // 构造客户端调用实例
                EssbasicClient client = Common.GetClientInstance(Configs.secretId, Configs.secretKey, Configs.endPoint);
                // 实例化一个请求对象,每个接口都会对应一个request对象
                ChannelCancelMultiFlowSignQRCodeRequest req = new ChannelCancelMultiFlowSignQRCodeRequest();

                // 第三方平台应用相关信息。
	            // 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
                // Agent

                // 二维码ID
                // QrCodeId
                
                // 返回的resp是一个ChannelCancelMultiFlowSignQRCodeResponse的实例，与请求对象对应
                ChannelCancelMultiFlowSignQRCodeResponse resp = client.ChannelCancelMultiFlowSignQRCodeSync(req);
                // 输出json格式的字符串回包
                Console.WriteLine(AbstractModel.ToJsonString(resp));
                return resp;
            }
            catch (Exception e)
            {
                Console.WriteLine(e.ToString());
            }
            Console.Read();
            return null;
        }


         public static void Main1(string[] args)
        {
            string qrCodeId = "********";
            ChannelCancelMultiFlowSignQRCodeResponse resp = ChannelCancelMultiFlowSignQRCode(Common.getAgent(),qrCodeId);
            Console.WriteLine(AbstractModel.ToJsonString(resp));
        }
    }
}
            