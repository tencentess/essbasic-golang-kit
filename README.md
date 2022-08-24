# 腾讯电子签企业版API接入工具包

## 项目说明
项目通过go get引入了腾讯云sdk，补充了调用电子签企业版API所需要的内容，并提供了调用的样例。使用前请先在项目中导入腾讯云sdk。

## 通过go get安装（推荐）

推荐使用腾讯云镜像加速下载：

1. Linux 或 MacOS:

    ```bash
    export GOPROXY=https://mirrors.tencent.com/go/
    ```

2. Windows:

    ```cmd
    set GOPROXY=https://mirrors.tencent.com/go/
    ```

### 按需安装（推荐）

注意：此安装方式仅支持使用 **Go Modules** 模式进行依赖管理，即环境变量 `GO111MODULE=auto`或者`GO111MODULE=on`, 并且在您的项目中执行了 `go mod init xxx`.

如果您使用 GOPATH, 请参考下节： 全部安装

v1.0.170后可以按照产品下载，您只需下载基础包和对应的产品包(如ess)即可，不需要下载全部的产品，从而加快您构建镜像或者编译的速度：

1. 安装公共基础包：

    ```bash
    go get -v -u github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common
    ```

2. 安装对应的产品包(如ess):

    ```bash
    go get -v -u github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/essbasic
    ```

### 全部安装

此模式支持 GOPATH 和 Go Modules

此方式会一次性下载腾讯云所有产品的包：

```bash
go get -v -u github.com/tencentcloud/tencentcloud-sdk-go
```

注意：为了支持 go mod，SDK 版本号从 v3.x 降到了 v1.x。并于2021.05.10移除了所有`v3.0.*`和`3.0.*`的tag，如需追溯以前的tag，请参考项目根目录下的 `commit2tag` 文件。

## 通过源码安装

前往代码托管地址 [Github](https://github.com/tencentcloud/tencentcloud-sdk-go) 或者 [Gitee](https://gitee.com/tencentcloud/tencentcloud-sdk-go) 下载最新代码，解压后安装到 $GOPATH/src/github.com/tencentcloud 目录下。

## 目录文件说明
### api
api目录是对电子签渠道版所有API的简单封装，以及调用的Example。
如果需要API更加高级的功能，需要结合业务修改api的封装。

### byfile
byfile目录是电子签渠道版的核心场景之一 - 通过文件发起的快速上手样例。
业务方可以结合自己的业务调整，用于正式对接。

### bytemplate
byfile目录是电子签渠道版的核心场景之一 - 通过模版发起的快速上手样例。
业务方可以结合自己的业务调整，用于正式对接。

### testdata
testdata是一个空白的pdf用于快速发起合同，测试。

### utils
tools目录提供了调用电子签渠道版API时涉及到的各种算法样例。
如果不使用sdk调用电子签服务，可参考其中的签名计算方式。

### config
里面定义调用电子签渠道版API需要的一些核心参数。

## 电子签企业版官网入口
[腾讯电子签企业版](https://cloud.tencent.com/document/product/1323)