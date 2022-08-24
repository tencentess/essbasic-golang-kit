package api

import (
	"essbasic-golang-kit_/utils"
	"fmt"
	essbasic "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/essbasic/v20210526"
	"testing"
)

// 用于生成pdf资源编号测试
func TestUploadFiles(t *testing.T) {
	agent := utils.SetAgent()
	filePath := "testdata/blank.pdf"
	fileBase64 := utils.ConvertImageFileToBase64(filePath)
	uploadFiles := []*essbasic.UploadFile{{
		FileBody: &fileBase64,
	}}

	response := UploadFiles(agent, uploadFiles)
	// 输出json格式的字符串回包
	fmt.Printf("%s", response.ToJsonString())
}
