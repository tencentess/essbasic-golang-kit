using System;
using System.Threading.Tasks;
using TencentCloud.Common;
using TencentCloud.Common.Profile;
using TencentCloud.Essbasic.V20210526;
using TencentCloud.Essbasic.V20210526.Models;

// ChannelCreateFlowByFiles
// 用于第三方应用集成通过文件创建签署流程。
// 注意事项：该接口需要依赖“多文件上传”接口生成pdf资源编号（FileIds）进行使用。
// 此接口静默签能力不可直接使用，需要运营申请
// 详细参考 https://cloud.tencent.com/document/api/1420/73068
namespace api
{
    class ChannelCreateFlowByFilesService
    {
        public static ChannelCreateFlowByFilesResponse ChannelCreateFlowByFiles(Agent agent, FlowApproverInfo[] flowApproverInfos, String flowName, String fileId)
        {
            try
            {
                // 构造客户端调用实例
                EssbasicClient client = Common.GetClientInstance(Configs.secretId, Configs.secretKey, Configs.endPoint);
                
                // 实例化一个请求对象,每个接口都会对应一个request对象
                ChannelCreateFlowByFilesRequest req = new ChannelCreateFlowByFilesRequest();

                // 第三方平台应用相关信息。
                // 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
                req.Agent = agent;

                 // 签署pdf文件的资源编号列表，通过UploadFiles接口获取
                req.FileIds = new string[] { fileId };

                // 签署流程参与者信息
                req.FlowApprovers = flowApproverInfos;

                // 签署流程名称,最大长度200个字符
                req.FlowName = flowName;

                // 其他更多参数和控制，参考文档 https://cloud.tencent.com/document/api/1420/73068
	            // 也可以结合test case传参
                
                // 返回的resp是一个ChannelCreateFlowByFilesResponse的实例，与请求对象对应
                ChannelCreateFlowByFilesResponse resp = client.ChannelCreateFlowByFilesSync(req);
                // 输出json格式的字符串回包
                // Console.WriteLine(AbstractModel.ToJsonString(resp));
                return resp;
            }
            catch (Exception e)
            {
                Console.WriteLine(e.ToString());
            }
            // Console.Read();
            return null;
        }
    }
}