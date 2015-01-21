docker run -d gogin /src/server_ubuntu
docker run -d -p 80:80 gogin /usr/local/nginx/sbin/nginx -c /src/nginx.conf
