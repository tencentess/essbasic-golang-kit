using System;
using System.Security.Cryptography;
using System.Text;

// 使用callbackKey解密
namespace TencentCloudExamples
{
    class EssCallback
    {

        static void Main1(string[] args)
        {
            try
            {
                // 传入CallbackUrlKey
                String key = "*************";
                // 传入密文
                String content = "*************";

                String plaintext = AESDecrypt(content, Encoding.ASCII.GetBytes(key));

                Console.WriteLine(plaintext);
            }
            catch (Exception e)
            {
                Console.WriteLine(e.ToString());
            }

        }

        public static string AESDecrypt(string encryptStr, byte[] key)
        {
            byte[] toEncryptArray = Convert.FromBase64String(encryptStr);
            RijndaelManaged rDel = new RijndaelManaged();
            rDel.Key = key;
            byte[] iv = new byte[16];
            Array.Copy(key, iv, iv.Length);
            rDel.IV = iv;
            rDel.Mode = CipherMode.CBC;
            rDel.Padding = PaddingMode.PKCS7;
            ICryptoTransform cTransform = rDel.CreateDecryptor();
            byte[] resultArray = cTransform.TransformFinalBlock(toEncryptArray, 0, toEncryptArray.Length);
            return UTF8Encoding.UTF8.GetString(resultArray);
        }

    }

}