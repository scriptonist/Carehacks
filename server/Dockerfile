FROM ubuntu
RUN apt-get update
WORKDIR /app
COPY ./server /app
CMD ["/app/server","-listen","0.0.0.0:80"]
EXPOSE 80