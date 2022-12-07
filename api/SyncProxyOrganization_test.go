package api

import (
	"essbasic-golang-kit_/utils"
	"fmt"
	"testing"
)

func TestSyncProxyOrganization(t *testing.T) {
	agent := utils.SetAgent()

	proxyOrganizationName := "*************"

	businessLicensePath := "../testdata/test_businessLicense.png"
	businessLicense := utils.ConvertImageFileToBase64(businessLicensePath)

	uniformSocialCreditCode := "*************"
	proxyLegalName := "*************"

	response := SyncProxyOrganization(agent, &proxyOrganizationName, &businessLicense,
		&uniformSocialCreditCode, &proxyLegalName)
	// 输出json格式的字符串回包
	fmt.Printf("%s", response.ToJsonString())
}
