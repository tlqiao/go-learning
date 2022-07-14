FROM golang:1.17 as builder
ENV GO111MODULE=off \
	CGO_ENABLED=0 \
	GOOS=linux \
	GOARCH=amd64
WORKDIR /build
ENV GOPROXY=https://goproxy.cn
COPY . .
RUN go build -o server

FROM scratch as runner
COPY --from=builder /build/server /opt/app/
CMD ["/opt/app/server"]
