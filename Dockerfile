FROM golang:1.14-alpine AS builder

RUN apk add --no-cache git openssl ca-certificates

COPY .  /go/src/github.com/qorpress/qorpress-auth-example
WORKDIR /go/src/github.com/qorpress/qorpress-auth-example

RUN cd /go/src/github.com/qorpress/qorpress-auth-example \
	&& go get github.com/Masterminds/glide \
	&& go get -v \
 	&& go build -v

FROM alpine:3.11 AS runtime

ARG TINI_VERSION=${TINI_VERSION:-"0.18.0"}
ARG TINI_ARCH=${TINI_ARCH:-"amd64"}

# Install tini to /usr/local/sbin
ADD https://github.com/krallin/tini/releases/download/v${TINI_VERSION}/tini-muslc-${TINI_ARCH} /usr/local/sbin/tini

# Install runtime dependencies & create runtime user
RUN apk --no-cache --no-progress add ca-certificates git libssh2 openssl \
 && chmod +x /usr/local/sbin/tini && mkdir -p /opt \
 && adduser -D qor -h /opt/qor -s /bin/sh \
 && su qor -c 'cd /opt/qor; mkdir -p bin config data'

# Switch to user context
USER qor
WORKDIR /opt/qor

COPY --from=builder /go/src/github.com/qorpress/qorpress-auth-example/qorpress-auth-example /opt/qor/bin/qorpress-auth-example
ENV PATH $PATH:/opt/qor/bin

# Container configuration
EXPOSE 9000

ENTRYPOINT ["tini", "-g", "--"]
CMD ["/opt/qor/bin/qorpress-auth-example"]
