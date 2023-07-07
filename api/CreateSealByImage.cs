using System;
using System.IO;
using TencentCloud.Common;
using TencentCloud.Essbasic.V20210526;
using TencentCloud.Essbasic.V20210526.Models;

// CreateSealByImage
// 平台企业通过图片为子客代创建印章，图片最大5MB
// 详细参考 https://cloud.tencent.com/document/api/1420/73067
namespace api
{
    class CreateSealByImageService
    {
        public static CreateSealByImageResponse CreateSealByImage(Agent agent, 
            String sealName, String sealImage)
        {
            try
            {
                // 构造客户端调用实例
                EssbasicClient client = Common.GetClientInstance(Configs.secretId, Configs.secretKey, Configs.endPoint);

                // 实例化一个请求对象,每个接口都会对应一个request对象
                CreateSealByImageRequest req = new CreateSealByImageRequest();

                // 第三方平台应用相关信息
                // 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
                req.Agent = agent;
                // 印章名称，最大长度不超过50字符
                req.SealName = sealName;
                // 印章图片base64，大小不超过10M（原始图片不超过7.6M）
                req.SealImage = sealImage;
                
                // 返回的resp是一个CreateSealByImageResponse的实例，与请求对象对应
                CreateSealByImageResponse resp = client.CreateSealByImageSync(req);
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
            String sealName = "************************";   
            Byte[] bytes = File.ReadAllBytes("testdata/test_seal.png");
            String sealImage = Convert.ToBase64String(bytes);
            CreateSealByImageResponse resp = CreateSealByImage(Common.getAgent(), sealName, sealImage);
            Console.WriteLine(AbstractModel.ToJsonString(resp));
        }
    }
}