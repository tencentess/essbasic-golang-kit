using System;
using System.IO;
using TencentCloud.Common;
using TencentCloud.Essbasic.V20210526;
using TencentCloud.Essbasic.V20210526.Models;

// SyncProxyOrganizationOperators
// 用于同步子客企业经办人列表，主要是同步经办人的离职状态。
// 子客Web控制台的组织架构管理，是依赖于平台企业的，无法针对员工做新增/更新/离职等操作。
// 若经办人信息有误，或者需要修改，也可以先将之前的经办人做离职操作，然后重新使用控制台链接CreateConsoleLoginUrl让经办人重新实名。
// 详细参考 https://cloud.tencent.com/document/api/1420/61517
namespace api
{
    class SyncProxyOrganizationOperatorsService
    {
        public static SyncProxyOrganizationOperatorsResponse SyncProxyOrganizationOperators(Agent agent, 
            String operatorType, ProxyOrganizationOperator[] proxyOrganizationOperators)
        {
            try
            {
                // 构造客户端调用实例
                EssbasicClient client = Common.GetClientInstance(Configs.secretId, Configs.secretKey, Configs.endPoint);
                
                // 实例化一个请求对象,每个接口都会对应一个request对象
                SyncProxyOrganizationOperatorsRequest req = new SyncProxyOrganizationOperatorsRequest();

                // 第三方平台应用相关信息
                // 此接口Agent.AppId 和 Agent.ProxyOrganizationOpenId必填。
                req.Agent = agent;
                // 操作类型，新增:"CREATE"，修改:"UPDATE"，离职:"RESIGN"
                req.OperatorType = operatorType;
                // 经办人信息列表，最大长度200
                req.ProxyOrganizationOperators = proxyOrganizationOperators;
                
                // 返回的resp是一个SyncProxyOrganizationOperatorsResponse的实例，与请求对象对应
                SyncProxyOrganizationOperatorsResponse resp = client.SyncProxyOrganizationOperatorsSync(req);
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
            String operatorType = "************************";


            ProxyOrganizationOperator proxyOrganizationOperator = new ProxyOrganizationOperator();
            proxyOrganizationOperator.Id = "************************";
            proxyOrganizationOperator.Name = "************************";
            proxyOrganizationOperator.IdCardType = "************************";
            proxyOrganizationOperator.IdCardNumber = "************************";
            proxyOrganizationOperator.Mobile = "************************";
            
            ProxyOrganizationOperator[] proxyOrganizationOperators = { proxyOrganizationOperator };
            
            SyncProxyOrganizationOperatorsResponse resp = SyncProxyOrganizationOperators(Common.getAgent(), 
                operatorType, proxyOrganizationOperators);
            Console.WriteLine(AbstractModel.ToJsonString(resp));
        }
    }
}