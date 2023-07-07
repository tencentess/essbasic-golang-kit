using System;
using TencentCloud.Common;
using TencentCloud.Essbasic.V20210526;
using TencentCloud.Essbasic.V20210526.Models;

// DescribeUsage
// 此接口（DescribeUsage）用于获取第三方应用集成所有合作企业流量消耗情况。
// 注: 此接口每日限频2次，若要扩大限制次数,请提前与客服经理或邮件至e-contract@tencent.com进行联系。
// 详细参考 https://cloud.tencent.com/document/api/1420/61520
namespace api
{
    class DescribeUsageService
    {
        public static DescribeUsageResponse DescribeUsage(Agent agent, 
            String startDate, String endDate, Boolean needAggregate, ulong limit, ulong offset)
        {
            try
            {
                // 构造客户端调用实例
                EssbasicClient client = Common.GetClientInstance(Configs.secretId, Configs.secretKey, Configs.endPoint);
                
                // 实例化一个请求对象,每个接口都会对应一个request对象
                DescribeUsageRequest req = new DescribeUsageRequest();

                // 应用信息，此接口Agent.AppId必填
                req.Agent = agent;
                // 开始时间，例如：2021-03-21
                req.StartDate = startDate;
                // 结束时间，例如：2021-06-21；
	            // 开始时间到结束时间的区间长度小于等于90天。
                req.EndDate = endDate;
                // 是否汇总数据，默认不汇总。
                // 不汇总：返回在统计区间内第三方应用集成下所有企业的每日明细，即每个企业N条数据，N为统计天数；
                // 汇总：返回在统计区间内第三方应用集成下所有企业的汇总后数据，即每个企业一条数据；
                req.NeedAggregate = needAggregate;
                // 单次返回的最多条目数量。默认为1000，且不能超过1000。
                req.Limit = limit;
                // 偏移量，默认是0。
                req.Offset = offset;
                
                // 返回的resp是一个DescribeUsageResponse的实例，与请求对象对应
                DescribeUsageResponse resp = client.DescribeUsageSync(req);
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
            String startDate = "2022-10-01";
            String endDate = "2022-10-20";
            Boolean needAggregate = false;
            ulong limit = 10;
            ulong offset = 0;
            DescribeUsageResponse resp = DescribeUsage(Common.getAgent(), 
                startDate, endDate, needAggregate, limit, offset);
            Console.WriteLine(AbstractModel.ToJsonString(resp));
        }
    }
}