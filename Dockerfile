FROM golang:1.14-alpine

RUN apk add --no-cache git openssl ca-certificates gcc musl-dev

COPY .  /go/src/github.com/qorpress/qorpress-auth-example
WORKDIR /go/src/github.com/qorpress/qorpress-auth-example

RUN cd /go/src/github.com/qorpress/qorpress-auth-example \
	&& go get github.com/Masterminds/glide \
	&& go get -v

RUN go build -v

# Container configuration
EXPOSE 4000

CMD ["/go/src/github.com/qorpress/qorpress-auth-example/qorpress-auth-example"]
