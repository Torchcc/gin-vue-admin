FROM golang:1.13.10-alpine AS builder

ENV \
    GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn,direct \
    SERVICE_NAME=mk-admin

# 移动到工作目录
WORKDIR /src/$SERVICE_NAME/

# 复制项目中的 go.mod 和 go.sum文件并下载依赖信息
COPY go.mod .
COPY go.sum .
RUN go mod download

# 将代码复制到容器中
COPY . .

# 将我们的代码编译成二进制可执行文
RUN go build -o app .

###################
# 接下来创建一个小镜像
###################
FROM scratch

ENV \
    ENV_DEPLOY=1 \
    BRANCH=prod\
    PORT=8888 \
    GIN_MODE=release \
    TZ=Asia/Shanghai \
    SERVICE_NAME=mk-admin




# 使用 tar 命令打包系统的时区 /usr/share/zoneinfo 定义压缩文件 zoneinfo.tar.gz
# 结合 docker ADD 命令可以让时区文件自动解压到 /usr/share/zoneinfo
# 预先将 zoneinfo 中的 /usr/share/zoneinfo/PRC 文件拷贝到根目录
# 将 PRC 复制到 /etc/localtime 可以改变本地时区，从而不需要改 golang 代码
ADD PRC /etc/localtime
ADD zoneinfo.tar.gz /

# 需要用https 的时候打开这行
ADD https://curl.haxx.se/ca/cacert.pem /etc/ssl/certs/

COPY ./deployment  /src/mk-admin/deployment
COPY ./resource /src/mk-admin/resource

# 从builder镜像中把/src/mk-api/app 拷贝到当前目录
COPY --from=builder /src/mk-admin/app /src/mk-admin/
WORKDIR /src/mk-admin/
VOLUME /src/mk-admin /home/ubuntu/admin

EXPOSE 8888

CMD ["/src/mk-admin/app"]
