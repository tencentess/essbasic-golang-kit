using System;
using System.Text;
using System.Security.Cryptography;

// 验证回调消息中的签名
namespace callback
{
    class CallbackVerify
    {
        static void Main1(string[] args)
        {
            // 回调消息体
            String payload = "**********";
            // secretToken 创建应用号时配置的
            String secretToken = "**********";

            // 1. 取出header [Content-Signature]
            String signFromHeader = "***********";

            // 2. 验证签名
            String hash = "sha256=" + VerifySignature(payload, secretToken);

            //3. 如果验证通过，继续处理。如果不通过，忽略该请求
            Console.WriteLine(hash == signFromHeader);
        }

        public static string VerifySignature(string data, string secret)
        {
            var encoding = new System.Text.UTF8Encoding();
            byte[] keyByte = encoding.GetBytes(secret);
            byte[] messageBytes = encoding.GetBytes(data);
            using (var hmacsha256 = new HMACSHA256(keyByte))
            {
                byte[] hashmessage = hmacsha256.ComputeHash(messageBytes);
                StringBuilder builder = new StringBuilder();
                for (int i = 0; i < hashmessage.Length; i++)
                {
                    builder.Append(hashmessage[i].ToString("x2"));
                }
                return builder.ToString();
            }
        }

    }
}
