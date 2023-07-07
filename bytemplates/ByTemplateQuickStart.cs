using System;
using System.Threading.Tasks;
using System.Collections.Generic;
using TencentCloud.Common;
using TencentCloud.Common.Profile;
using TencentCloud.Essbasic.V20210526;
using TencentCloud.Essbasic.V20210526.Models;
using api;

/*
本示例用于第三方应用集成接口对接，通过模板快速发起第一份合同
建议配合文档进行操作，先修改config里的基本参数以及对应环境域名，然后跑一次
第三方应用集成主要针对平台企业-代理子客发起合同，简要步骤主要是
    1. 通过CreateConsoleLoginUrl引导子客企业完成电子签的实名认证 - 子客企业在电子签配置印章等
    2. 通过简单封装的CreateFlowByTemplateDirectly接口快速发起一份合同，并得到签署链接
    3. 在小程序签署合同，通过API下载合同
基于具体业务上的参数调用，可以参考官网的接口说明 
https://cloud.tencent.com/document/product/1420/61534
每个API的封装在api目录下可以自己配合相关参数进行调用
*/
namespace bytemplate
{
    class ByTemplateQuickStart
    {
        static void Main(string[] args)
        {
            // Step 1 登录子客控制台
            // 子客企业真实名称
            String proxyOrganizationName = "我的企业";

            // 创建控制台链接
            CreateConsoleLoginUrlResponse loginUrlResponse =
                    CreateConsoleLoginUrlService.CreateConsoleLoginUrl(Common.getAgent(), proxyOrganizationName);

            //Step 2 发合同
            // 定义合同名
            String flowName = "我的第2个合同";
            // 模板Id，根据自己传入的模板需求修改第三个参数为对应的TemplateId，在Config中配置
            String templateId = Configs.templateId;

            // 获取模板里面的参与方RecipientId
            Recipient[] recipients = ByTemplate.GetRecipients(templateId);
            if (recipients == null)
            {
                throw new Exception("签署人不能为空");
            }

           // 此处为快速发起的签署方；如果是正式接入，构造签署方，请参考函数内说明，构造需要的场景参数
            FlowApproverInfo[] flowApproverInfos = ByTemplate.BuildApprovers(new List<Recipient>(recipients));

            // 发起合同
            // 样例为BtoC
            Dictionary<String, String[]> resp = CreateFlowByTemplateDirectly.ChannelCreateFlowByTemplateDirectly(flowName,
                     templateId, flowApproverInfos);

            //  返回相关信息
            Console.WriteLine("您的控制台入口为：");
            Console.WriteLine(loginUrlResponse.ConsoleUrl);
            Console.WriteLine("\r\n\r\n");
            
            int count = Configs.count;
            for (int i = 0; i < count; i++)
            {
                // 返回合同Id
                Console.WriteLine("您创建的合同id为：");
                Console.WriteLine(resp["FlowIds"][i]);
                Console.WriteLine("\r\n\r\n");

                // 返回签署的链接
                Console.WriteLine("签署链接为：");
                Console.WriteLine(resp["Urls"][i]);
                Console.WriteLine("\r\n\r\n");

                // Step 3 下载合同
                // 返回合同下载链接
                DescribeResourceUrlsByFlowsResponse urlResp = DescribeResourceUrlsByFlowsService.DescribeResourceUrlsByFlows(Common.getAgent(),new String[]{resp["FlowIds"][i]} );
                Console.WriteLine("请访问以下地址下载您的合同：");
                Console.WriteLine(urlResp.FlowResourceUrlInfos[0].ResourceUrlInfos[0].Url);
                Console.WriteLine("\r\n\r\n");
            }
        }
    }
}
