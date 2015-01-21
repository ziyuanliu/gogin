#GOGIN
####A combination of golang's Gin webframe work and Nginx
***

*easily* scalable multi/general purpose messaging platform that places a focus
on low latency, concurrent connections, and messages/sec.

[NGINX user summit '14'](https://www.youtube.com/watch?v=yL4Q7D4ynxU)

[Disqus Pycon '13](https://www.youtube.com/watch?v=5A5Iw9z6z2s)

A great amount of research and inspiration came from Disqus' architecture/application
talks. According to their presentations, the whole entire cluster includes 7 servers (2 to manage queues from the main server, 5 to communicate (comet) with clients). However, by assumption, the real time notification system at Disqus sounds looks very basic in terms of just relaying new information (comments) without refreshing -- one step above AJAX-polls. For us, it starts with a handshake. For Disqus, it seems like browser compatibility is the number one concern and thus long-polling is the mode of transportation.  
***
###Tools
####[Gin](https://github.com/gin-gonic/gin) - public endpoint written in Go
####[NGINX](http://nginx.org/) - Server
####[NGINX-Push-Stream-module](https://github.com/wandenberg/nginx-push-stream-module) - "Comet made easy and really scalable"
***

###Installation

*see* Dockerfile

***
###Process

The main concern with our setup is security. We want to isolate our users from the server channels. As per the creator of the Push Stream modules [advice](https://groups.google.com/forum/#!topic/nginxpushstream/rEPdXl3vNpA), we will set the subscription and publishing endpoints to internal/deny all (allow localhost) settings. The subscription goes as follows

Client -> Subscribe/ws -> Gin Server will authenticate by header (talk to main server to confirm auth)  -> main server will return all the channels to be subscribed -> 200 the handshake and X-Accel-Redirect the response to the internal subscription endpoint

The publishing side of things goes as follows:

Client -> sends message with auth (api key/session key) and object -> Gin Server authenticates and authorizes send -> main server is notified and sends a notification as usual

Communication with main server?  internal HTTP requests or message queues
***
###NOTE

Socket.IO and other high-level implementation libs will not work with this implementation, such implementations are mainly used in combination
with its official server implementations with the official websocket.


HTML5 websocket is used in place for framework-agnostic frontends. Websocket is around 90% global adoption
