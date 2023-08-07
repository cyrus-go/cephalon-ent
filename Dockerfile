FROM --platform=$BUILDPLATFORM golang:1.20-alpine AS builder

LABEL maintainer="StarkSim<stark@cephalon.ai>"

# 在容器根目录创建 src 目录
WORKDIR /src

#RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk add --no-cache g++

COPY ./go.mod .

COPY ./go.sum .

ENV GOPROXY="https://goproxy.cn"

RUN go mod download

COPY . .

ARG TARGETOS
ARG TARGETARCH

RUN CGO_ENABLE=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -trimpath -ldflags "-s -w" -o main ./main.go

FROM alpine:latest

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk add --no-cache tzdata

WORKDIR /app

COPY --from=builder /src/entrypoint.sh /app/
COPY --from=builder /src/main /app/
COPY --from=builder /src/internal/db/migrations /app/internal/db/migrations/
COPY --from=builder /src/config.yaml /app/cephalon-ent/config.yaml

EXPOSE 8080

ENTRYPOINT ["./entrypoint.sh"]
