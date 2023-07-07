using System;
using System.Threading.Tasks;
using TencentCloud.Common;
using TencentCloud.Common.Profile;
using TencentCloud.Essbasic.V20210526;
using TencentCloud.Essbasic.V20210526.Models;

// ChannelCreateMultiFlowSignQRCode
//  用于创建一码多扫流程签署二维码。
//  适用场景：无需填写签署人信息，可通过模板id生成签署二维码，签署人可通过扫描二维码补充签署信息进行实名签署。常用于提前不知道签署人的身份信息场景，例如：劳务工招工、大批量员工入职等场景。
//  适用的模板仅限于B2C（1、无序签署，2、顺序签署时B静默签署，3、顺序签署时B非首位签署）、单C的模板，且模板中发起方没有填写控件。
// 详细参考 https://cloud.tencent.com/document/api/1420/75452
namespace api
{
    class ChannelCreateMultiFlowSignQRCodeService
    {
        static ChannelCreateMultiFlowSignQRCodeResponse ChannelCreateMultiFlowSignQRCode(Agent agent, String templateId, String flowName)
        {
            try
            {
                // 构造客户端调用实例
                EssbasicClient client = Common.GetClientInstance(Configs.secretId, Configs.secretKey, Configs.endPoint);
                // 实例化一个请求对象,每个接口都会对应一个request对象
                ChannelCreateMultiFlowSignQRCodeRequest req = new ChannelCreateMultiFlowSignQRCodeRequest();

                // 第三方平台应用相关信息。
                // 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
                req.Agent = agent;
                // 签署流程名称，最大长度200个字符。
                req.FlowName = flowName;
                // 模板ID
                req.TemplateId = templateId;
                
                // 返回的resp是一个ChannelCreateMultiFlowSignQRCodeResponse的实例，与请求对象对应
                ChannelCreateMultiFlowSignQRCodeResponse resp = client.ChannelCreateMultiFlowSignQRCodeSync(req);
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
            string templateId = "********";
            string flowName = "************";
            ChannelCreateMultiFlowSignQRCodeResponse resp = ChannelCreateMultiFlowSignQRCode(Common.getAgent(),templateId,flowName);
            Console.WriteLine(AbstractModel.ToJsonString(resp));
        }
    }
}
            