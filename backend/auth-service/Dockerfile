# syntax=docker/dockerfile:1

FROM golang:1.19

WORKDIR /

COPY ./ ./
RUN go mod download
RUN chmod 0777 /run.sh


ENTRYPOINT ["/run.sh"]