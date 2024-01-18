# syntax=docker/dockerfile:1
ARG GO_VERSION=1.21
ARG VERSION
FROM --platform=$BUILDPLATFORM golang:${GO_VERSION} as builder

WORKDIR /src

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=go.sum,target=go.sum \
    --mount=type=bind,source=go.mod,target=go.mod \
    go mod download -x

ARG TARGETARCH

RUN --mount=type=bind,target=. \
    --mount=type=cache,target=/go/pkg/mod \
    CGO_ENABLED=0 GOARCH=$TARGETARCH \
    go build -ldflags "-s -w -X main.version=${VERSION}" -o /emptydirclone main.go

FROM alpine
# Add util-linux to get a new version of losetup.
RUN apk update && apk upgrade && apk add util-linux coreutils
COPY --from=builder /emptydirclone /emptydirclone
ENTRYPOINT ["/emptydirclone"]
