## http的504错误

504错误代表网关超时 （Gateway timeout），是指服务器作为网关或代理，但是没有及时从上游服务器收到请求。
这通常意味着上游服务器已关闭（不响应网关 / 代理），而不是上游服务器和网关/代理在交换数据的协议上不一致。

首先，了解什么是网关。
网络的基本概念：
客户端:应用 C/S（客户端/服务器） B/S（浏览器/服务器）
服务器：为客户端提供服务、数据、资源的机器
请求：客户端向服务器索取数据
响应：服务器对客户端请求作出反应，一般是返回给客户端数据

在这之中，把nginx或Apache作为网关。一般服务的架构是：用PHP则是nginx+php的一系列进程，Apache+tomcat+JVM。
- nginx+php
- apache+tomcat+jvm

网关超时就与nginx或Apache配置的超时时间，和与php线程、java线程的响应时间有关。
以nginx与PHP为例：它的超时配置fastcgi_connect_timeout、fastcgi_send_timeout、fastcgi_read_timeout。
nginx将请求丢给PHP来处理，某个PHP的线程响应时间假如是10s，在10s内没有响应给nginx就报超时。
这时可以打开PHP慢日志记录，然后排查之。

另外，数据库的慢查询也会导致504 。nginx只要进程没有死，一般不是nginx的问题。
假如场景是：确定程序执行是正确的，比如向数据库插入大量数据，需要5分钟，nginx设置的超时时间是3分钟。
这时候可以将超时时间临时设置为大于5分钟。























