using System;
using TencentCloud.Common;
using TencentCloud.Essbasic.V20210526;
using TencentCloud.Essbasic.V20210526.Models;

// ChannelDescribeOrganizationSeals
// 查询子客企业电子印章，需要操作者具有管理印章权限
// 客户指定需要获取的印章数量和偏移量，数量最多100，超过100按100处理；
// 入参InfoType控制印章是否携带授权人信息，为1则携带，为0则返回的授权人信息为空数组。
// 接口调用成功返回印章的信息列表还有企业印章的总数。
// 详细参考 https://cloud.tencent.com/document/api/1420/82455
namespace api
{
    class ChannelDescribeOrganizationSealsService
    {
        public static ChannelDescribeOrganizationSealsResponse ChannelDescribeOrganizationSeals(Agent agent, 
            long infoType, String sealId, long limit, long offset)
        {
            try
            {
                // 构造客户端调用实例
                EssbasicClient client = Common.GetClientInstance(Configs.secretId, Configs.secretKey, Configs.endPoint);

                // 实例化一个请求对象,每个接口都会对应一个request对象
                ChannelDescribeOrganizationSealsRequest req = new ChannelDescribeOrganizationSealsRequest();

                // 第三方平台应用相关信息
                // 此接口Agent.ProxyOrganizationOpenId、Agent. ProxyOperator.OpenId、Agent.AppId 必填。
                req.Agent = agent;
                // 查询信息类型，为1时返回授权用户，为其他值时不返回
                req.InfoType = infoType;
                // 印章id（没有输入返回所有）
                req.SealId = sealId;
                // 返回最大数量，最大为100
                req.Limit = limit;
                // 偏移量，默认为0，最大为20000
                req.Offset = offset;
                
                // 返回的resp是一个ChannelDescribeOrganizationSealsResponse的实例，与请求对象对应
                ChannelDescribeOrganizationSealsResponse resp = client.ChannelDescribeOrganizationSealsSync(req);
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
            long infoType = 0;
            String sealId = "************************";
            long limit = 10;
            long offset = 0;           
            ChannelDescribeOrganizationSealsResponse resp = ChannelDescribeOrganizationSeals(Common.getAgent(), 
                infoType, sealId, limit,  offset);
            Console.WriteLine(AbstractModel.ToJsonString(resp));
        }
    }
}