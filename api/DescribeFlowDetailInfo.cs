using System;
using System.Threading.Tasks;
using TencentCloud.Common;
using TencentCloud.Common.Profile;
using TencentCloud.Essbasic.V20210526;
using TencentCloud.Essbasic.V20210526.Models;

// DescribeFlowDetailInfo 此接口用于查询合同(签署流程)的详细信息。
// 详细参考 https://cloud.tencent.com/document/api/1420/66683
namespace api
{
    class DescribeFlowDetailInfoService
    {
        static DescribeFlowDetailInfoResponse DescribeFlowDetailInfo(Agent agent, String[] flowIds)
        {
            try
            {
                // 构造客户端调用实例
                EssbasicClient client = Common.GetClientInstance(Configs.secretId, Configs.secretKey, Configs.endPoint);
                // 实例化一个请求对象,每个接口都会对应一个request对象
                DescribeFlowDetailInfoRequest req = new DescribeFlowDetailInfoRequest();

                // 第三方平台应用相关信息
                // 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
                req.Agent = agent;
                // 合同(流程)编号数组，最多支持100个
                req.FlowIds = flowIds;
                
                // 返回的resp是一个DescribeFlowDetailInfoResponse的实例，与请求对象对应
                DescribeFlowDetailInfoResponse resp = client.DescribeFlowDetailInfoSync(req);
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
            DescribeFlowDetailInfoResponse resp = DescribeFlowDetailInfo(Common.getAgent(), new String[]{flowId});
            Console.WriteLine(AbstractModel.ToJsonString(resp));
        }
    }
}
            