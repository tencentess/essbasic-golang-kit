using System;
using TencentCloud.Common;
using TencentCloud.Essbasic.V20210526;
using TencentCloud.Essbasic.V20210526.Models;

// ChannelGetTaskResultApi
// 第三方应用集成查询转换任务状态
// 详细参考 https://cloud.tencent.com/document/api/1420/78773
namespace api
{
    class ChannelGetTaskResultApiService
    {
        public static ChannelGetTaskResultApiResponse ChannelGetTaskResultApi(Agent agent, 
            String taskId)
        {
            try
            {
                // 构造客户端调用实例
                EssbasicClient client = Common.GetClientInstance(Configs.secretId, Configs.secretKey, Configs.endPoint);

                // 实例化一个请求对象,每个接口都会对应一个request对象
                ChannelGetTaskResultApiRequest req = new ChannelGetTaskResultApiRequest();

                // 第三方平台应用相关信息
                // 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
                req.Agent = agent;
                // 任务Id，通过ChannelCreateConvertTaskApi接口获得
                req.TaskId = taskId;
                
                // 返回的resp是一个ChannelGetTaskResultApiResponse的实例，与请求对象对应
                ChannelGetTaskResultApiResponse resp = client.ChannelGetTaskResultApiSync(req);
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
            String taskId = "************************";       
            ChannelGetTaskResultApiResponse resp = ChannelGetTaskResultApi(Common.getAgent(), taskId);
            Console.WriteLine(AbstractModel.ToJsonString(resp));
        }
    }
}