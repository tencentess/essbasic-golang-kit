package api

import (
	"fmt"
	"testing"
)

// 查询文件下载URL测试
func TestDescribeFileUrls(t *testing.T) {
	// 资源所对应的签署流程Id
	flowId := "*************"

	response := DescribeFileUrls(&flowId)
	fmt.Printf("%s", *response)
}
