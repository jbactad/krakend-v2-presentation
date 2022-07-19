FROM golang:1.17.11-alpine3.15 as build

WORKDIR /src

RUN apk add make gcc musl-dev binutils-gold
COPY go.mod go.sum ./
RUN go mod download

COPY plugins ./plugins

RUN go build -buildmode=plugin -o ./bin/headerlogger.so ./plugins/headerlogger
RUN go build -buildmode=plugin -o ./bin/s3.so ./plugins/s3
RUN go build -buildmode=plugin -o ./bin/authtoapikey.so ./plugins/authtoapikey

FROM devopsfaith/krakend:2.0.5

COPY --from=build /src/bin/ /opt/krakend/plugins/





