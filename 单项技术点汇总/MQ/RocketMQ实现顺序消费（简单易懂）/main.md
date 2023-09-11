## RocketMQ实现顺序消费（简单易懂）

### 前提
- 1.搭好SpringBoot集成RocketMQ环境代码。
- 2.使用Docker安装RocketMQ以及它的控制台。

### 需求背景
用订单进行分区有序的示例。一个订单的顺序流程是：创建、付款、推送、完成。
订单号相同的消息会被先后发送到同一个队列中，消费时，同一个OrderId获取到的肯定是同一个队列。

### 顺序消费的原理图解析
在默认的情况下消息发送会采取Round Robin轮询方式把消息发送到不同的queue(分区队列)；
而消费消息的时候从多个queue上拉取消息，这种情况发送和消费是不能保证顺序。
但是如果控制发送的顺序消息只依次发送到同一个queue中，消费的时候只从这个queue上依次拉取，则就保证了顺序。
当发送和消费参与的queue只有一个，则是全局有序；如果多个queue参与，
则为分区有序，即相对每个queue，消息都是有序的。

### 生产者代码
```
@Test
public void sendOrderlyMsg() {
    //顺序消息
    //选择器规则构建
    rocketMQTemplate.setMessageQueueSelector(new MessageQueueSelector() {
        @Override
        public MessageQueue select(List<MessageQueue> list, org.apache.rocketmq.common.message.Message message, Object o) {
            String i = ((String) o);
            Long  index = Long.parseLong(i);
            int  i1 = (int )(index % list.size());
            return list.get(i1);
        }
    });
    //注意这个消息用springboot包的message
    List<OrderStep> orderSteps = OrderUtil.buildOrders();
        for (OrderStep orderStep : orderSteps) {
            Message msg = MessageBuilder.withPayload(orderStep.toString()).build();
            rocketMQTemplate.sendOneWayOrderly("OrderlybootTopic",orderStep,String.valueOf(orderStep.getOrderId()));
        }
    }
}    
```

### 消费者代码
```
/**
 * 顺序消费
 */
@Component
@RocketMQMessageListener(consumerGroup = "Orderly-Consumer", topic = "OrderlybootTopic",
        consumeMode = ConsumeMode.ORDERLY)

public class OrderlyConsumer implements RocketMQListener<MessageExt> {
    @Override
    public void onMessage(MessageExt message) {
        System.out.println("线程"+Thread.currentThread()+"内容为:"
                + new String(message.getBody())+
                "队列序号:"+message.getQueueId());
    }
}
```

### 成功实现顺序消费
消息进同一队列，一个队列对应一个消费者，按照创建→付款→推送→完成的步骤进行消费。




















