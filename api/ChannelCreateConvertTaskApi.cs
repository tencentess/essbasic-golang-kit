using System;
using TencentCloud.Common;
using TencentCloud.Essbasic.V20210526;
using TencentCloud.Essbasic.V20210526.Models;

// ChannelCreateConvertTaskApi
// 平台企业创建文件转换任务
// 详细参考 https://cloud.tencent.com/document/api/1420/78774
namespace api
{
    class ChannelCreateConvertTaskApiService
    {
        public static ChannelCreateConvertTaskApiResponse ChannelCreateConvertTaskApi(Agent agent, String resourceName,
            String resourceId, String resourceType)
        {
            try
            {
                // 构造客户端调用实例
                EssbasicClient client = Common.GetClientInstance(Configs.secretId, Configs.secretKey, Configs.endPoint);

                // 实例化一个请求对象,每个接口都会对应一个request对象
                ChannelCreateConvertTaskApiRequest req = new ChannelCreateConvertTaskApiRequest();

                // 第三方平台应用相关信息
                // 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
                req.Agent = agent;

                // 资源名称，长度限制为256字符
                req.ResourceName = resourceName;
                // 资源Id，通过UploadFiles获取
                req.ResourceId = resourceId;
                // 资源类型 取值范围doc,docx,html,xls,xlsx之一
                req.ResourceType = resourceType;
                
                // 返回的resp是一个ChannelCreateConvertTaskApiResponse的实例，与请求对象对应
                ChannelCreateConvertTaskApiResponse resp = client.ChannelCreateConvertTaskApiSync(req);
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
            String resourceName = "hello.docx";
            String resourceId = "************************";
            String resourceType = "docx";           
            ChannelCreateConvertTaskApiResponse resp = ChannelCreateConvertTaskApi(Common.getAgent(), 
                resourceName, resourceId, resourceType);
            Console.WriteLine(AbstractModel.ToJsonString(resp));
        }
    }
}