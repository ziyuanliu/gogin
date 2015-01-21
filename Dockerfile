FROM ubuntu:trusty

#Add the files to src
ADD nginx-push-stream-module /src/nginx-push-stream-module
ADD nginx /src/nginx
ADD . /src

#UPDATE
RUN apt-get update
RUN apt-get install -y git
RUN apt-get -y install python-software-properties  # 12.04
RUN add-apt-repository ppa:duh/golang
RUN apt-get update
RUN apt-get -y install golang

#SET ENV
ENV GOPATH /root/goApps/
ENV PATH /root/goApps/bin:$PATH

RUN \
  dpkg --get-selections | grep -v deinstall | awk '{print $1}' | sort > /tmp/initial-packages && \
  DEBIAN_FRONTEND=noninteractive apt-get -y build-dep nginx && \
  cd /src/nginx && \
  ./configure --add-module=../nginx-push-stream-module && \
  make && \
  cd /src/nginx && \
  make install && \
  dpkg --get-selections | grep -v deinstall | awk '{print $1}' | sort > /tmp/final-packages && \
  DEBIAN_FRONT_END=noninteractive apt-get -y purge `comm -13 /tmp/initial-packages /tmp/final-packages`


#FETCH Packages
ADD gin /root/goApps/src/github.com/gin-gonic/gin
RUN cd /src; go build -o server server.go;

#Set the commands
CMD /usr/local/nginx/sbin/nginx -g 'daemon off;'
CMD /src/server

EXPOSE 80
#EXPOSE 443
