using System;
using System.Threading.Tasks;
using System.Collections.Generic;
using TencentCloud.Common;
using TencentCloud.Common.Profile;
using TencentCloud.Essbasic.V20210526;
using TencentCloud.Essbasic.V20210526.Models;

// CreateFlowByFileDirectly 通过文件base64直接发起签署流程，返回flowId和签署链接
// 本接口是对于发起合同几个接口的封装，详细参数需要根据自身业务进行调整
// UploadFiles--ChannelCreateFlowByFiles--CreateSignUrls
namespace api
{
    class CreateFlowByFileDirectly
    {
        public static Dictionary<String, String> ChannelCreateFlowByFilesDirectly(String fileBase64,
                                                               FlowApproverInfo[] flowApproverInfos, String flowName)
        {
            Dictionary<String, String> flowIdAndUrl = new Dictionary<String, String>();

            // 设置agent参数
            Agent agent = Common.getAgent();
            String fileId = "";
            String flowId = "";
            String url = "";


            // 设置uploadFile参数,这里可以修改传入数量
            UploadFile[] uploadFiles = new UploadFile[] { new UploadFile() };
            uploadFiles[0].FileBody = fileBase64;

            // 上传文件获取fileId
            UploadFilesResponse uploadRes = UploadFilesService.UploadFiles(agent, fileBase64, flowName);

            // fileId
            if (!uploadRes.Equals(null))
            {
                fileId = uploadRes.FileIds[0];
            }

            // 创建签署流程
            ChannelCreateFlowByFilesResponse createFlowRes = ChannelCreateFlowByFilesService.
                    ChannelCreateFlowByFiles(agent, flowApproverInfos, flowName, fileId);

            // 获取flowId
            if (!createFlowRes.Equals(null))
            {
                flowId = createFlowRes.FlowId;
            }

            // 获取签署链接
            String[] flowIds = new String[1];
            flowIds[0] = flowId;
            CreateSignUrlsResponse createSignUrlsRes = CreateSignUrlsService.CreateSignUrls(flowIds, agent);
            if (!createSignUrlsRes.Equals(null))
            {
                url = createSignUrlsRes.SignUrlInfos[0].SignUrl;
            }


            flowIdAndUrl.Add("flowId", flowId);
            flowIdAndUrl.Add("url", url);

            return flowIdAndUrl;
        }
    }
}