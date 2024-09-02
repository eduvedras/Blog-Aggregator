FROM --platform=linux/amd64 debian:stable-slim

RUN apt-get update && apt-get install -y ca-certificates
#ENV PORT="8080"
#ENV CONN_STRING="postgres://postgres:postgres@localhost:5432/blogator?sslmode=disable"

ADD blog-aggregator /usr/bin/blog-aggregator

CMD ["blog-aggregator"]
