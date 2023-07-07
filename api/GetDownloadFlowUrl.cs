using System;
using System.Threading.Tasks;
using TencentCloud.Common;
using TencentCloud.Common.Profile;
using TencentCloud.Essbasic.V20210526;
using TencentCloud.Essbasic.V20210526.Models;

// GetDownloadFlowUrl
// 用于创建电子签批量下载地址，让合作企业进入控制台直接下载，支持客户合同（流程）按照自定义文件夹形式 分类下载。
// 当前接口限制最多合同（流程）50个.
// 详细参考 https://cloud.tencent.com/document/api/1420/66368
namespace api
{
    class GetDownloadFlowUrlService
    {
        static GetDownloadFlowUrlResponse GetDownloadFlowUrl(Agent agent, DownloadFlowInfo[] downLoadFlows)
        {
            try
            {
                // 构造客户端调用实例
                EssbasicClient client = Common.GetClientInstance(Configs.secretId, Configs.secretKey, Configs.endPoint);
                // 实例化一个请求对象,每个接口都会对应一个request对象
                GetDownloadFlowUrlRequest req = new GetDownloadFlowUrlRequest();

                // 第三方平台应用相关信息
                // 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
                req.Agent = agent;
                // 文件夹数组，签署流程总数不能超过50个，一个文件夹下，不能超过20个签署流程
                req.DownLoadFlows = downLoadFlows;
                
                // 返回的resp是一个GetDownloadFlowUrlResponse的实例，与请求对象对应
                GetDownloadFlowUrlResponse resp = client.GetDownloadFlowUrlSync(req);
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
            DownloadFlowInfo downloadFlowInfo = new DownloadFlowInfo();
            downloadFlowInfo.FileName = "***";
            string flowId = "********";
            downloadFlowInfo.FlowIdList = new String[]{flowId};
            GetDownloadFlowUrlResponse resp = GetDownloadFlowUrl(Common.getAgent(), new DownloadFlowInfo[]{downloadFlowInfo});
            Console.WriteLine(AbstractModel.ToJsonString(resp));
        }
    }
}
            