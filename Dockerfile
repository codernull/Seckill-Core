FROM golang:1.20 AS builder


# #RUN go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
# #RUN kratos new server
# #RUN cd server
# #RUN kratos proto add api/server/server.proto
# #RUN kratos proto client api/server/server.proto
# #RUN kratos proto server api/server/server.proto -t internal/service
# RUN go generate ./...
# RUN go build -o ./bin/ ./...
# RUN ./bin/server -conf ./configs

# #RUN go get github.com/google/wire/cmd/wire
# #RUN cd cmd/server
# #RUN wire
# #COPY . .
# RUN make init
# RUN pwd 
# RUN make api
# RUN make all

COPY . /src
WORKDIR /src

RUN make generate

RUN GOPROXY=https://goproxy.cn make build





FROM debian:stable-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
		ca-certificates  \
        netbase \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y

COPY --from=builder /src/bin /app

WORKDIR /app

EXPOSE 8000
EXPOSE 9000
VOLUME /data/conf



CMD ["./server", "-conf", "/data/conf"]
