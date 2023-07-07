using System;
using TencentCloud.Common;
using TencentCloud.Essbasic.V20210526;
using TencentCloud.Essbasic.V20210526.Models;

// ChannelDescribeEmployees
// 查询企业员工列表
// 详细参考 https://cloud.tencent.com/document/api/1420/81119
namespace api
{
    class ChannelDescribeEmployeesService
    {
        public static ChannelDescribeEmployeesResponse ChannelDescribeEmployees(Agent agent, 
            Filter[] filters, long limit, long offset)
        {
            try
            {
                // 构造客户端调用实例
                EssbasicClient client = Common.GetClientInstance(Configs.secretId, Configs.secretKey, Configs.endPoint);

                // 实例化一个请求对象,每个接口都会对应一个request对象
                ChannelDescribeEmployeesRequest req = new ChannelDescribeEmployeesRequest();

                // 第三方平台应用相关信息
                // 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
                req.Agent = agent;

                // 查询过滤实名用户，Key为Status，Values为["IsVerified"]
                // 根据第三方系统openId过滤查询员工时,Key为StaffOpenId,Values为["OpenId","OpenId",...]
                // 查询离职员工时，Key为Status，Values为["QuiteJob"]
                req.Filters = filters;
                // 返回最大数量，最大为20
                req.Limit = limit;
                // 偏移量，默认为0，最大为20000
                req.Offset = offset;
                
                // 返回的resp是一个ChannelDescribeEmployeesResponse的实例，与请求对象对应
                ChannelDescribeEmployeesResponse resp = client.ChannelDescribeEmployeesSync(req);
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
            Filter filter = new Filter();
            filter.Key = "";
            filter.Values = new[] { "" };
            Filter[] filters = { filter };
            long limit = 10;
            long offset = 0;           
            ChannelDescribeEmployeesResponse resp = ChannelDescribeEmployees(Common.getAgent(), 
                filters, limit,  offset);
            Console.WriteLine(AbstractModel.ToJsonString(resp));
        }
    }
}