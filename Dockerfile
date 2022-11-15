ARG VERSION=1.17

FROM golang:${VERSION}-alpine As build

ENV GO111MODULE=on

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh vim

WORKDIR /var/www

COPY . ./

RUN go mod download

RUN go build

CMD ["/var/www/Cinnox-Homework", "config -p config.yaml"]
