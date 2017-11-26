FROM nginx
RUN apt-get update
COPY ./dist /usr/share/nginx/html
COPY ch-client.conf /etc/nginx/conf.d/default.conf
EXPOSE 80
