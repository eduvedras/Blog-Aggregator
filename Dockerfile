FROM debian:stable-slim

ENV PORT="8080"
ENV CONN_STRING="postgres://postgres:postgres@localhost:5432/blogator?sslmode=disable"

# COPY source destination
#COPY blog-aggregator /bin/blog-aggregator

#CMD ["/bin/blog-aggregator"]
ADD blog-aggregator /usr/bin/blog-aggregator

CMD ["blog-aggregator"]
