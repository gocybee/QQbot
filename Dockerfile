FROM golang:alpine as builder

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPORXY=https://goproxy.cn

# 移动到工作目录：/build
WORKDIR /build

# 将代码复制到容器中
COPY . .

# 将我们的代码编译成二进制可执行文件app
RUN go build -o app ./cmd/main.go

FROM scratch

# 移动到用于存放生成的二进制文件的 /bot 目录
WORKDIR /bot

# 将二进制文件从 /build 目录复制到这里
COPY --from=builder /build/app .
# COPY --from=builder /build/config/ ./config/
# 自行映射数据卷

# 声明服务端口
EXPOSE 8080

# 启动容器时运行的命令
CMD ["/bot/app"]