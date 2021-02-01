# builder image
FROM golang:1.14.13-alpine3.11 as builder

ENV CGO_ENABLED 0
RUN apk --no-cache add git
WORKDIR /wanchain-cli
COPY . /wanchain-cli
RUN go build -o /bin/wanchain-cli -v \
  -ldflags "-X github.com/linki/wanchain-cli/cmd.version=$(git describe --tags --always --dirty) -w -s"
RUN /bin/wanchain-cli version

# final image
FROM alpine:3.13.1

RUN apk --no-cache add ca-certificates
COPY --from=builder /bin/wanchain-cli /bin/wanchain-cli

USER 65534
ENTRYPOINT ["/bin/wanchain-cli"]
