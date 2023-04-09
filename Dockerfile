FROM golang:1.19-alpine

RUN mkdir /app

WORKDIR /app

COPY docker-entrypoint.sh /docker-entrypoint.sh
RUN chmod +x /docker-entrypoint.sh

ADD . /app

RUN go build -o main .

CMD ["/docker-entrypoint.sh"]