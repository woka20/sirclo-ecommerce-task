FROM golang:1.16-alpine

WORKDIR /go/src/github.com/woka20/sirclo-ecommerce-task/order/

# Add the source code
ENV SRC_DIR=/go/src/github.com/woka20/sirclo-ecommerce-task/order/

# COPY ./go.mod ./
# COPY ./go.sum ./
# RUN go mod download

ADD . $SRC_DIR

COPY *.go ./

# RUN go build -o /docker-gs-ping
RUN apk update && apk add --no-cache $BUILD_PACKAGES \
  && apk del $BUILD_PACKAGES

EXPOSE 3002

COPY .env $SRC_DIR.env

RUN go mod init github.com/woka20/sirclo-ecommerce-task/order
RUN go mod tidy

RUN cd $SRC_DIR; CGO_ENABLED=0 GOOS=linux go build -ldflags '-w -s' -a -o main

ENTRYPOINT ["./main"]