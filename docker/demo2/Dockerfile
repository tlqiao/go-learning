FROM golang:1.17 AS builder
ENV GO111MODULE=off \
	CGO_ENABLED=0 \
	GOOS=linux \
	GOARCH=amd64
WORKDIR /build
COPY . .
RUN go build -o httpServer

FROM scratch
COPY --from=builder /build/httpServer /
EXPOSE 9099
CMD ["/httpServer"]
