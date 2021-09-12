FROM golang:1.15.2-alpine3.12 as build
ARG go_name=bysykkel
ENV RUN_NAME=$go_name
ENV CGO_ENABLED=0

WORKDIR /go/src/app

COPY . .

RUN go test .
RUN go run .
RUN go build -o $go_name

RUN echo -e "#!/bin/sh \n/go/src/app/${RUN_NAME}" > ./entrypoint.sh
RUN chmod +x ./entrypoint.sh

ENTRYPOINT ["/go/src/app/entrypoint.sh"]