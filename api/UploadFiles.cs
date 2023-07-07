using System;
using System.Threading.Tasks;
using TencentCloud.Common;
using TencentCloud.Common.Profile;
using TencentCloud.Essbasic.V20210526;
using TencentCloud.Essbasic.V20210526.Models;

// UploadFiles 用于生成pdf资源编号（FileIds）来配合“用PDF创建流程”接口使用，使用场景可详见“用PDF创建流程”接口说明。
// 调用时需要设置Domain, 正式环境为 file.ess.tencent.cn
// 测试环境为 https://file.test.ess.tencent.cn
// 详细参考 https://cloud.tencent.com/document/api/1420/71479
namespace api
{
    class UploadFilesService
    {
        public static UploadFilesResponse UploadFiles(Agent agent,string fileBase64, string filename)
        {
            try
            {
                // 构造客户端调用实例
                EssbasicClient client = Common.GetClientInstance(Configs.secretId, Configs.secretKey, Configs.fileServiceEndPoint);
                // 实例化一个请求对象,每个接口都会对应一个request对象
                UploadFilesRequest req = new UploadFilesRequest();

                // 应用相关信息，AppId 和 ProxyOrganizationOpenId 必填
                req.Agent = agent;
                // 文件对应业务类型
                // 1. TEMPLATE - 模板； 文件类型：.pdf/.doc/.docx/.html
                // 2. DOCUMENT - 签署过程及签署后的合同文档/图片控件 文件类型：.pdf/.doc/.docx/.jpg/.png/.xls.xlsx/.html
                req.BusinessType="DOCUMENT";

                // 上传文件内容数组，最多支持20个文件
                UploadFile file = new UploadFile();
                // Base64编码后的文件内容
                file.FileBody = fileBase64;
                // 文件名，最大长度不超过200字符
                file.FileName = filename;

                req.FileInfos = new UploadFile[] { file };

                // 返回的resp是一个UploadFilesResponse的实例，与请求对象对应
                UploadFilesResponse resp = client.UploadFilesSync(req);
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
            string fileBase64 = "********";
            string filename = "********";
            UploadFilesResponse resp = UploadFiles(Common.getAgent(), fileBase64, filename);
            Console.WriteLine(AbstractModel.ToJsonString(resp));
        }
    }
}
            