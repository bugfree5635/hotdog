FROM golang:1.23-alpine AS build

# 设置国内代理，解决连接被拒绝
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o hotdog .

FROM alpine
RUN apk add --no-cache ca-certificates
WORKDIR /root/
COPY --from=build /app/hotdog .
EXPOSE 8080
CMD ["./hotdog"]