# builder image
FROM golang:1.13-alpine3.10 as builder

ENV CGO_ENABLED 0
ENV GO111MODULE on
WORKDIR /go/src/github.com/linki/wanchain-cli
COPY . .
RUN go build -o /bin/wanchain-cli -v -ldflags "-w -s"

# final image
FROM alpine:3.10

RUN apk --no-cache add ca-certificates
COPY --from=builder /bin/wanchain-cli /bin/wanchain-cli

USER 65534
ENTRYPOINT ["/bin/wanchain-cli"]
