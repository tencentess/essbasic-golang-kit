using System;
using System.Threading.Tasks;
using TencentCloud.Common;
using TencentCloud.Common.Profile;
using TencentCloud.Essbasic.V20210526;
using TencentCloud.Essbasic.V20210526.Models;

// DescribeTemplates 查询该子客企业在电子签拥有的有效模板，不包括平台企业模板
// 详细参考 https://cloud.tencent.com/document/api/1420/61521
namespace api
{
    class DescribeTemplatesService
    {
        public static DescribeTemplatesResponse DescribeTemplates(Agent agent, String TemplateId)
        {
            try
            {
                // 构造客户端调用实例
                EssbasicClient client = Common.GetClientInstance(Configs.secretId, Configs.secretKey, Configs.endPoint);
                
                // 实例化一个请求对象,每个接口都会对应一个request对象
                DescribeTemplatesRequest req = new DescribeTemplatesRequest();

                // 第三方平台应用相关信息
                // 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
                req.Agent = agent;
                // 模板唯一标识，查询单个模板时使用
                req.TemplateId = TemplateId;

                // 其他查询参数参考官网文档
	            // https://cloud.tencent.com/document/api/1420/61521
                
                // 返回的resp是一个DescribeTemplatesResponse的实例，与请求对象对应
                DescribeTemplatesResponse resp = client.DescribeTemplatesSync(req);
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
            String templateId = "**************";
            DescribeTemplatesResponse resp = DescribeTemplates(Common.getAgent(), templateId);
            Console.WriteLine(AbstractModel.ToJsonString(resp));
        }
    }
}
            