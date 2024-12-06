FROM golang:1.20 AS builder


COPY . /src

WORKDIR /src

RUN GOPROXY=https://goproxy.cn make build

FROM debian:stable-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
		ca-certificates  \
        netbase \
        protobuf-compiler\
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y \
        && apt-get autoclean -y

# 使用轻量级的基础镜像

COPY --from=builder /src/bin /app

WORKDIR /app

EXPOSE 8000
EXPOSE 9000
VOLUME /data/conf

CMD ["./seckill", "-conf", "/data/conf"]
