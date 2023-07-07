using System;
using System.Threading.Tasks;
using System.Collections.Generic;
using TencentCloud.Common;
using TencentCloud.Common.Profile;
using TencentCloud.Essbasic.V20210526;
using TencentCloud.Essbasic.V20210526.Models;
using api;

namespace bytemplate
{
    class ByTemplate
    {
        // 构造签署人 - 以个人为例, 实际请根据自己的场景构造签署方、控件
        public static FlowApproverInfo[] BuildApprovers(List<Recipient> recipients)
        {

            // 个人签署方参数
            String personName = "********";
            String personMobile = "*********";

            // 企业签署方参数
            String organizationName = "*********";
            String organizationOpenId = "***************";
            String openId = "*********";

            List<FlowApproverInfo> flowApproverInfoList = new List<FlowApproverInfo>();
            foreach (Recipient recipient in recipients)
            {
                switch (recipient.RecipientType)
                {
                    case "INDIVIDUAL":
                        flowApproverInfoList.Add(BuildPersonApprover(personName, personMobile, recipient.RecipientId));
                        break;
                    case "ENTERPRISE":
                        flowApproverInfoList.Add(BuildOrganizationApprover(organizationName, organizationOpenId, openId, recipient.RecipientId));
                        break;

                }
            }

            // 转换为对象数组
            FlowApproverInfo[] flowApproverInfo = flowApproverInfoList.ToArray();

            // 内容控件填充结构，详细说明参考
            // https://cloud.tencent.com/document/api/1420/61525#FormField

            return flowApproverInfo;
        }

        // 打包个人签署方参与者信息
        public static FlowApproverInfo BuildPersonApprover(String name, String mobile, String recipientId)
        {

            // 签署参与者信息
            // 个人签署方
            FlowApproverInfo flowApproverInfo = new FlowApproverInfo();

            // 签署人类型
            // PERSON-个人/自然人；
            // ORGANIZATION-企业（企业签署方或模板发起时的企业静默签）；
            // ENTERPRISESERVER-企业静默签（文件发起时的企业静默签字）。
            flowApproverInfo.ApproverType = "PERSON";

            // 签署人姓名，最大长度50个字符
            flowApproverInfo.Name = name;
            // 签署人手机号，脱敏显示。大陆手机号为11位，暂不支持海外手机号
            flowApproverInfo.Mobile = mobile;

            // 模板中对应签署方的参与方id
            flowApproverInfo.RecipientId = recipientId;

            return flowApproverInfo;
        }

        // 打包企业签署方参与者信息
        public static FlowApproverInfo BuildOrganizationApprover(String organizationName, String organizationOpenId, String openId, String recipientId)
        {

            // 签署参与者信息
            // 企业签署方
            FlowApproverInfo flowApproverInfo = new FlowApproverInfo();

            // 签署人类型
            // PERSON-个人/自然人；
            // ORGANIZATION-企业（企业签署方或模板发起时的企业静默签）；
            // ENTERPRISESERVER-企业静默签（文件发起时的企业静默签字）。
            flowApproverInfo.ApproverType = "ORGANIZATION";

            // 企业签署方工商营业执照上的企业名称，签署方为非发起方企业场景下必传，最大长度64个字符；
            flowApproverInfo.OrganizationName = organizationName;

            // 如果签署方是子客企业，此处需要传子客企业的OrganizationOpenId
            // 企业签署方在同一第三方应用集成下的其他合作企业OpenId，签署方为非发起方企业场景下必传，最大长度64个字符；
            flowApproverInfo.OrganizationOpenId = organizationOpenId;
            // 如果签署方是子客企业，此处需要传子客企业经办人的OpenId
            // 当签署方为同一平台下的员工时，该字段若不指定，则发起【待领取】的流程
            flowApproverInfo.OpenId = openId;

            // 模板中对应签署方的参与方id
            flowApproverInfo.RecipientId = recipientId;

            return flowApproverInfo;
        }

        // 打包企业静默签署方参与者信息
        public static FlowApproverInfo BuildServerSignApprover()
        {
            // 签署参与者信息
            // 企业静默签
            FlowApproverInfo flowApproverInfo = new FlowApproverInfo();

            // 签署人类型
            // PERSON-个人/自然人；
            // ORGANIZATION-企业（企业签署方或模板发起时的企业静默签）；
            // ENTERPRISESERVER-企业静默签（文件发起时的企业静默签字）。
            flowApproverInfo.ApproverType = "ENTERPRISESERVER";

            // 注：此时发起方会替换为接口调用的企业+经办人，所以不需要传签署方信息

            return flowApproverInfo;
        }

        //获取所有签署人信息
        public static Recipient[] GetRecipients(String templateId)
        {
            Agent agent = Common.getAgent();
            DescribeTemplatesResponse templatesResponse = DescribeTemplatesService.DescribeTemplates(agent,
                    templateId);
            return templatesResponse.Templates[0].Recipients;
        }


    }
}
