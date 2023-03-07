:: Windows下编译Mac平台64位可执行程序
set CGO_ENABLED=0
set GOOS=darwin
set GOARCH=amd64
go build
