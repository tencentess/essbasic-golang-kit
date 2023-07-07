using System;
using System.IO;
using TencentCloud.Common;
using TencentCloud.Essbasic.V20210526;
using TencentCloud.Essbasic.V20210526.Models;

// SyncProxyOrganization
// 用于同步子客企业信息，主要是子客企业的营业执照，便于子客企业开通过程中不用手动上传。
// 若有需要调用此接口，需要在创建控制链接CreateConsoleLoginUrl之后即刻进行调用。
// 详细参考 https://cloud.tencent.com/document/api/1420/61518
namespace api
{
    class SyncProxyOrganizationService
    {
        public static SyncProxyOrganizationResponse SyncProxyOrganization(Agent agent, 
            String proxyOrganizationName, String pusinessLicense, String uniformSocialCreditCode, String proxyLegalName)
        {
            try
            {
                // 构造客户端调用实例
                EssbasicClient client = Common.GetClientInstance(Configs.secretId, Configs.secretKey, Configs.endPoint);
                
                // 实例化一个请求对象,每个接口都会对应一个request对象
                SyncProxyOrganizationRequest req = new SyncProxyOrganizationRequest();

                // 第三方平台应用相关信息
                // 此接口Agent.AppId、Agent.ProxyOrganizationOpenId必填
                req.Agent = agent;
                // 子客企业名称，最大长度64个字符
                req.ProxyOrganizationName = proxyOrganizationName;
                // 营业执照正面照(PNG或JPG) base64格式, 大小不超过5M
                req.BusinessLicense = pusinessLicense;
                // 子客企业统一社会信用代码，最大长度200个字符
                req.UniformSocialCreditCode = uniformSocialCreditCode;
                // 子客企业法人/负责人姓名
                req.ProxyLegalName = proxyLegalName;
                
                // 返回的resp是一个SyncProxyOrganizationResponse的实例，与请求对象对应
                SyncProxyOrganizationResponse resp = client.SyncProxyOrganizationSync(req);
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
            String proxyOrganizationName = "************************";
            String pusinessLicense = "************************";
            Byte[] bytes = File.ReadAllBytes("testdata/test_businessLicense.png");
            String uniformSocialCreditCode = Convert.ToBase64String(bytes);
            String proxyLegalName = "************************";
            
            SyncProxyOrganizationResponse resp = SyncProxyOrganization(Common.getAgent(), 
                proxyOrganizationName, pusinessLicense, uniformSocialCreditCode, proxyLegalName);
            Console.WriteLine(AbstractModel.ToJsonString(resp));
        }
    }
}