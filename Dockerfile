FROM golang:1.14 AS builder

ENV GO111MODULE=on

RUN mkdir k4ever
WORKDIR /k4ever
COPY go.mod go.sum ./
RUN go mod download && \
    go get -u github.com/gobuffalo/packr/packr && \
    go get -u github.com/go-swagger/go-swagger/cmd/swagger

COPY . .

RUN go generate && \
    packr && \
    CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags "-w -extldflags -static -X github.com/freitagsrunde/k4ever-backend/internal/context.GitCommit=$(git rev-parse HEAD) -X github.com/freitagsrunde/k4ever-backend/internal/context.GitBranch=$(git rev-parse --abbrev-ref HEAD) -X github.com/freitagsrunde/k4ever-backend/internal/context.BuildTime=$(date -u --iso-8601=seconds ) -X github.com/freitagsrunde/k4ever-backend/internal/context.version=0.0.1" -o /go/bin/k4ever

FROM alpine:3.5

COPY --from=builder /go/bin/k4ever /go/bin/k4ever

RUN chmod +x /go/bin/k4ever && apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

ENTRYPOINT ["/go/bin/k4ever"]



