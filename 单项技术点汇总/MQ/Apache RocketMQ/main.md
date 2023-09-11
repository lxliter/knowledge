## 为什么选择RocketMQ

### 为什么 RocketMQ

在阿里孕育 RocketMQ 的雏形时期，我们将其用于异步通信、搜索、社交网络活动流、数据管道，贸易流程中。
随着我们的贸易业务吞吐量的上升，源自我们的消息传递集群的压力也变得紧迫。

- 异步
- 削峰
- 解耦

根据我们的研究，随着***队列和虚拟主题***使用的增加，ActiveMQ IO模块达到了一个瓶颈。
我们尽力通过节流、断路器或降级来解决这个问题，但效果并不理想。
于是我们尝试了流行的消息传递解决方案Kafka。
不幸的是，Kafka不能满足我们的要求，其尤其表现在***低延迟和高可靠性***方面，详见下文。
在这种情况下，我们决定发明一个新的消息传递引擎来处理更广泛的消息用例，
覆盖从传统的pub/sub场景到高容量的实时零误差的交易系统。

Apache RocketMQ 自诞生以来，因其架构简单、业务功能丰富、具备极强可扩展性等特点被众多企业开发者以及云厂商广泛采用。
历经十余年的大规模场景打磨，RocketMQ 已经成为业内共识的***金融级可靠业务消息首选方案***，
被广泛应用于互联网、大数据、移动互联网、物联网等领域的业务场景。

提示
下表显示了RocketMQ、ActiveMQ和Kafka之间的比较

### RocketMQ vs. ActiveMQ vs. Kafka

|Messaging Product|Client SDK|Protocol and Specification|Ordered Message |Ordered Message|Scheduled Message|Batched
Message|BroadCast Message |Message Filter|Server Triggered Redelivery|Message Storage|Message Retroactive[追溯的]|Message
Priority|High Availability and Failover|Message Track|Configuration|Management and Operation Tools|
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
































