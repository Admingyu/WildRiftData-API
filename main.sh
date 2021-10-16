export PATH=$PATH:/usr/local/go/bin/go
/usr/local/go/bin/go env -w GOPROXY=https://goproxy.cn
/usr/local/go/bin/go env
/usr/local/go/bin/go mod tidy
/usr/local/go/bin/go run main.go

