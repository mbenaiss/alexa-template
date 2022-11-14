FROM golang:alpine as build-stage
WORKDIR /project

RUN apk add build-base curl xz
RUN curl -OL https://github.com/upx/upx/releases/download/v3.96/upx-3.96-arm64_linux.tar.xz \
    &&  tar xvf upx-3.96-arm64_linux.tar.xz \
    && mv upx-3.96-arm64_linux/upx ./

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -ldflags "-s -w"  -o app *.go

RUN ./upx app

FROM alpine
WORKDIR /project

COPY --from=build-stage /project /project

ENTRYPOINT ["./app"]
