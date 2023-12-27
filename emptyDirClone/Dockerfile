# syntax=docker/dockerfile:1
FROM golang:1.21 as builder
ARG VERSION

WORKDIR /app

RUN --mount=source=go.mod,target=go.mod \
    --mount=source=go.sum,target=go.sum \
    --mount=type=cache,target=/go/pkg \
    go mod download -x

RUN --mount=target=. \
    --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux \
    go build -ldflags "-s -w -X main.version=${VERSION}" -o /emptydirclone main.go

FROM gcr.io/distroless/static-debian12:debug as debug
COPY --from=builder /emptydirclone /emptydirclone
ENTRYPOINT ["/emptydirclone"]

FROM gcr.io/distroless/static-debian12 as final
COPY --from=builder /emptydirclone /emptydirclone
ENTRYPOINT ["/emptydirclone"]
