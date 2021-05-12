# 构建
FROM tetafro/golang-gcc:1.16-alpine as builder
WORKDIR /img2webp
ENV GOPROXY=https://goproxy.cn
COPY ./admin/ ./
RUN go mod download && \
sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
apk add --no-cache tzdata
RUN go build -ldflags "-s -w" -o admin

# 打包
FROM alpine as runner
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=builder /img2webp/admin /img2webp/
WORKDIR /img2webp
CMD ["./admin", "-host", "0.0.0.0", "-port", "8861"]
