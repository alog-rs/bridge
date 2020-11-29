FROM golang:alpine AS builder

RUN apk add --no-cache git

LABEL maintainer="Chris Frank <dev@alog.rs>"

WORKDIR /bridge

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN mkdir -p build

RUN GIT_TAG=$(git describe --tags --always --long --dirty) && \
    CGO_ENABLED=0 \
    go build --ldflags "\
    -X github.com/alog-rs/bridge/internal/helpers.Version=$GIT_TAG" \
    -o ./build/bridge ./main.go

FROM scratch

COPY --from=builder ./bridge/build/bridge .

CMD ["./bridge"]