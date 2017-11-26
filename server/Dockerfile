FROM ubuntu
RUN apt-get update
RUN apt-get install -y ca-certificates
WORKDIR /app
COPY ./api /app/
CMD ["/app/api","-listen","0.0.0.0:80"]
EXPOSE 80
