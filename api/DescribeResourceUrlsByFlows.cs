using System;
using System.Threading.Tasks;
using TencentCloud.Common;
using TencentCloud.Common.Profile;
using TencentCloud.Essbasic.V20210526;
using TencentCloud.Essbasic.V20210526.Models;

// DescribeResourceUrlsByFlows
// 根据签署流程信息批量获取资源下载链接，可以下载签署中、签署完的合同，需合作企业先进行授权。
// 此接口直接返回下载的资源的url，与接口GetDownloadFlowUrl跳转到控制台的下载方式不同。
// 详细参考 https://cloud.tencent.com/document/api/1420/63220
namespace api
{
    class DescribeResourceUrlsByFlowsService
    {
        public static DescribeResourceUrlsByFlowsResponse DescribeResourceUrlsByFlows(Agent agent, String[] flowIds)
        {
            try
            {
                // 构造客户端调用实例
                EssbasicClient client = Common.GetClientInstance(Configs.secretId, Configs.secretKey, Configs.endPoint);
                // 实例化一个请求对象,每个接口都会对应一个request对象
                DescribeResourceUrlsByFlowsRequest req = new DescribeResourceUrlsByFlowsRequest();

                // 第三方平台应用相关信息
                // 设置agent参数
                // 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
                req.Agent = agent;
                // 查询资源所对应的签署流程Id，最多支持50个
                req.FlowIds = flowIds;

                
                
                // 返回的resp是一个DescribeResourceUrlsByFlowsResponse的实例，与请求对象对应
                DescribeResourceUrlsByFlowsResponse resp = client.DescribeResourceUrlsByFlowsSync(req);
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
            string flowId = "********";
            DescribeResourceUrlsByFlowsResponse resp = DescribeResourceUrlsByFlows(Common.getAgent(), new String[]{flowId});
            Console.WriteLine(AbstractModel.ToJsonString(resp));
        }
    }
}
            