using System;
using TencentCloud.Common;
using TencentCloud.Essbasic.V20210526;
using TencentCloud.Essbasic.V20210526.Models;

// OperateChannelTemplate
// 用于针对平台企业模板库中的模板对子客企业可见性的查询和设置，不会直接分配平台企业模板给子客企业。
// 1、OperateType=select时：
// 查询平台企业模板库
// 2、OperateType=update或者delete时：
// 对子客企业进行模板库中模板可见性的修改、删除操作。
// 详细参考 https://cloud.tencent.com/document/api/1420/66367
namespace api
{
    class OperateChannelTemplateService
    {
        public static OperateChannelTemplateResponse OperateChannelTemplate(Agent agent, 
            String operateType, String templateId, String proxyOrganizationOpenIds, String authTag)
        {
            try
            {
                // 构造客户端调用实例
                EssbasicClient client = Common.GetClientInstance(Configs.secretId, Configs.secretKey, Configs.endPoint);
                
                // 实例化一个请求对象,每个接口都会对应一个request对象
                OperateChannelTemplateRequest req = new OperateChannelTemplateRequest();

                // 第三方平台应用相关信息
                // 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
                req.Agent = agent;
                // 操作类型，查询:"SELECT"，删除:"DELETE"，更新:"UPDATE"
                req.OperateType = operateType;
                // 平台企业模板库模板唯一标识
                req.TemplateId = templateId;
                // 合作企业方第三方机构唯一标识数据，支持多个， 用","进行分隔
                req.ProxyOrganizationOpenIds = proxyOrganizationOpenIds;
                // 模板可见性, 全部可见-"all", 部分可见-"part"
                req.AuthTag = authTag;
                
                // 返回的resp是一个OperateChannelTemplateResponse的实例，与请求对象对应
                OperateChannelTemplateResponse resp = client.OperateChannelTemplateSync(req);
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
            String operateType = "SELECT";
            String templateId = "************************";
            String proxyOrganizationOpenIds = "";
            String authTag = "all";
            OperateChannelTemplateResponse resp = OperateChannelTemplate(Common.getAgent(), 
                operateType, templateId, proxyOrganizationOpenIds, authTag);
            Console.WriteLine(AbstractModel.ToJsonString(resp));
        }
    }
}