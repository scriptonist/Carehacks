FROM ubuntu
RUN apt-get update
WORKDIR /app
COPY ./server /app
EXPOSE 9080
CMD ["/app/server"]