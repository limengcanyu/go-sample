# go-sample

开发环境设置

1. 安装go1.12.5.windows-amd64.msi
2. GoLand中设置Go的 GOROOT 和 GOPATH ，GOPATH中设置项目GOPATH，指定GO的工作空间，即workspace
3. 新增main.go文件，即为应用的入口函数

GOROOT: go sdk 安装目录

/Users/rock/sdk/go1.14.4

GOPATH: 

ProjectPath:

/Users/rock/GolandProjects，其下有 src 目录，存放代码

go env -w GOPROXY=https://goproxy.cn,direct


## go 命令

go mod init example/hello

go mod init limengcanyu.com/example

go mod tidy



go env

go run .

go build

go install 

go get

## Go项目结构

代码保存在 $GOPATH/src 目录下。在工程运行 go build、go install 或者 go get 等命令，下载的第三方包源代码放在 $GOPATH/src 目录下，生成的二进制可执行文件放在 $GOPATH/bin 目录下，生成的中间缓存文件放在 $GOPATH/pkg 目录下。


## VSCode开发配置

1. 安装 Go 扩展

2. 更新 Go 工具

$ go env -w GO111MODULE=on
$ go env -w GOPROXY=https://goproxy.cn,direct

在 Visual Studio Code 中，打开命令面板的 HelpShow>All Commands。 或使用键盘快捷方式 (Ctrl+Shift+P)

Go: Install/Update tools搜索然后从托盘运行命令

出现提示时，选择所有可用的 Go 工具，然后单击“确定”。

等待 Go 工具完成更新。

5. 编写示例 Go 程序

选择main.go，按F5执行

