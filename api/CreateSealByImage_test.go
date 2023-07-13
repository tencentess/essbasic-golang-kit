package api

import (
	"essbasic-golang-kit_/utils"
	"fmt"
	"testing"
)

func TestCreateSealByImage(t *testing.T) {
	agent := utils.SetAgent()

	sealName := "印章名称"

	sealPath := "../testdata/test_seal.png"
	sealImage := utils.ConvertImageFileToBase64(sealPath)

	response := CreateSealByImage(agent, &sealName, &sealImage)
	// 输出json格式的字符串回包
	fmt.Printf("%s", response.ToJsonString())
}
