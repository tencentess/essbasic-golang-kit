using System;
using System.Threading.Tasks;
using System.Collections.Generic;
using TencentCloud.Common;
using TencentCloud.Common.Profile;
using TencentCloud.Essbasic.V20210526;
using TencentCloud.Essbasic.V20210526.Models;

// CreateFlowByTemplateDirectly 通过合同名和模板Id直接发起签署流程
// 本接口是对于发起合同几个接口的封装，详细参数需要根据自身业务进行调整
// CreateFlowsByTemplates--CreateSignUrls
namespace api
{
    class CreateFlowByTemplateDirectly
    {
        public static Dictionary<String, String[]> ChannelCreateFlowByTemplateDirectly(String flowName,
                                                                     String templateId,
                                                                     FlowApproverInfo[] flowApproverInfos)
        {
            String[] FlowIds = { };
            // 设置agent参数
            Agent agent = Common.getAgent();
            
            Dictionary<String, String[]> resp = new Dictionary<String, String[]>();

            // 创建签署流程
            // 签署数量
            int count = Configs.count;
            FlowInfo[] FlowInfos = new FlowInfo[count];
            for (int i = 0; i < count; i++)
            {
                FlowInfo flowInfo = new FlowInfo();
                flowInfo.TemplateId = templateId;
                flowInfo.FlowName = flowName;
                flowInfo.FlowApprovers = flowApproverInfos;
                flowInfo.FlowType = "合同";

                // 填写控件填充
                // FormField formField = new FormField();
                // formField.ComponentName = "name";
                // formField.ComponentValue = "zhangsan";
                // flowInfo.FormFields = new FormField[] {formField};

                FlowInfos[i] = flowInfo;
                
            }

            // 发起签署
            CreateFlowsByTemplatesResponse flowResponse = CreateFlowsByTemplatesService.CreateFlowsByTemplates(agent, FlowInfos);
            if (!flowResponse.Equals(null))
            {
                FlowIds = flowResponse.FlowIds;
                resp.Add("FlowIds", FlowIds);
            }


            // 获取签署链接
            CreateSignUrlsResponse createSignUrlsRes = CreateSignUrlsService.CreateSignUrls(FlowIds, agent);
            if (!createSignUrlsRes.Equals(null))
            {
                if (createSignUrlsRes.SignUrlInfos.Length != 0)
                {
                    String[] Urls = new String[count];
                    for (int i = 0; i < count; i++)
                    {
                        Urls[i] = createSignUrlsRes.SignUrlInfos[i].SignUrl;
                    }
                    resp.Add("Urls", Urls);
                }
            }

            return resp;
        }
    }
}
