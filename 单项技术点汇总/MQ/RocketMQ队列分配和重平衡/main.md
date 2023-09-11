## RocketMQ队列分配和重平衡
接着上节来说，Broker端如何进行分配消息的？
同时当新增或者删除消费者时，如果进行重平衡，被其他消费者分配后如何处理？

### 1、Consumer启动
负载均衡消息队列，并分配当前Consumer可以消费的MessageQueue。
RebalanceService是以守护线程启动，每 waitInterval=20s 调用一次doRebalance()进行分配

```g
public MQClientInstance(ClientConfig clientConfig, int instanceIndex, String clientId, RPCHook rpcHook) {
    this.rebalanceService = new RebalanceService(this);
}
public void start() throws MQClientException {
    synchronized (this) {
        switch (this.serviceState) {
            case CREATE_JUST:
                this.rebalanceService.start();
            default:
               break;    
    }
}    
public void run() {
    while (!this.isStopped()) {
        this.waitForRunning(waitInterval);
        this.mqClientFactory.doRebalance();
    }
}
```

### 2、Topic分配消息队列
#### 2.1 BROADCASTING：针对广播模式，分配 Topic 对应的所有消息队列。
```
case BROADCASTING: {
    Set<MessageQueue> mqSet = this.topicSubscribeInfoTable.get(topic);
    if (mqSet != null) {
        boolean changed = this.updateProcessQueueTableInRebalance(topic, mqSet, isOrder);
        if (changed) {
            this.messageQueueChanged(topic, mqSet, mqSet);
        }
}  
```

#### 2.2 CLUSTERING： 分配 Topic 对应的部分消息队列。
```
// RebalanceImpl# rebalanceByTopic
case CLUSTERING: {
    Set<MessageQueue> mqSet = this.topicSubscribeInfoTable.get(topic);
    List<String> cidAll = this.mQClientFactory.findConsumerIdList(topic, consumerGroup);

    if (mqSet != null && cidAll != null) {
        List<MessageQueue> mqAll = new ArrayList<MessageQueue>();
        mqAll.addAll(mqSet);
        Collections.sort(mqAll);
        Collections.sort(cidAll);
        AllocateMessageQueueStrategy strategy = this.allocateMessageQueueStrategy;
        List<MessageQueue> allocateResult = null;
        try {
            allocateResult = strategy.allocate(this.consumerGroup,this.mQClientFactory.getClientId(),mqAll,cidAll);
        } 
    }      
}
```

##### 1.获取MessageQueue列表 mqAll 和ConsumerId列表 cidAll 并进行排序。
##### 2.获取consumer配置的分配策略：具体怎么分配可以在consumer端进行设置，
默认为 AllocateMessageQueueAveragely平均分配策略.
RocketMQ提供一共包括如下策略，可以根据实际场景去设置。怎么设置则可参考如下代码：
```
DefaultMQPushConsumer consumer = new DefaultMQPushConsumer("arch-rocketmq-consumer");
final AllocateMachineRoomNearby.MachineRoomResolver machineRoomResolver =  new AllocateMachineRoomNearby.MachineRoomResolver() {
    @Override public String brokerDeployIn(MessageQueue messageQueue) {
        return messageQueue.getBrokerName().split("-")[0];
    }
    @Override public String consumerDeployIn(String clientID) {
        return clientID.split("-")[0];
    }
};
consumer.setAllocateMessageQueueStrategy(new AllocateMachineRoomNearby(new AllocateMessageQueueAveragely(), machineRoomResolver);
```

- 平均分配策略(默认)：AllocateMessageQueueAveragely
- 环形分配策略：AllocateMessageQueueAveragelyByCircle
- 手动配置分配策略：AllocateMessageQueueByConfig
- 机房分配策略：AllocateMessageQueueByMachineRoom
- 一致性哈希分配策略：AllocateMessageQueueConsistentHash

下面以 AllocateMessageQueueAveragely 在这里举一个队列分配机制的示例
1、当一个topic 有4个消息队列(q1,q2,q3,q4) ，如果有2个消费者 c1,c2。按照平均分配策略算法： c1->q1,q3; c2->q2,q4
2、当一个topic 有4个消息队列(q1,q2,q3,q4) ，如果有4个消费者 c1,c2,c3,c4。按照平均分配策略算法: c1->q1; c2->q2;c3->q3; c4->q4;
3、当一个topic 有4个消息队列(q1,q2,q3,q4) ，如果有5个消费者 c1,c2,c3,c4,c5。按照平均分配策略算法: c1->q1; c2->q2;c3->q3; c4->q4; 
其中c5不会被分配到MessageQueue队列。这时候可以考虑扩展messageQueue队列数。

##### 3.分配队列时，判断 topic 对应的消息队列返回是否有变化
updateProcessQueueTableInRebalance() 当分配队列时，更新 Topic 对应的消息队列，并返回是否有变更。
- 对于不存在分配的消息队列 mqSet 的消息队列 processQueueTable(消费者数量减少)。
- 对于队列拉取超时，即 当前时间 - 最后一次拉取消息时间 > 120s ( 120s 可配置)进行移除。
- 增加 不在processQueueTable && 存在于mqSet 里的消息队列。对于该consumer 分配到多的MessageQueue，需要组装 pullRequestList 并调用 executePullRequestImmediately() 将拉取消息的请求放入 pullRequestQueue 进行下一次长轮询拉取。

```
RebalanceImpl# rebalanceByTopic
boolean changed = this.updateProcessQueueTableInRebalance(topic, allocateResultSet, isOrder);
if (changed) {
       this.messageQueueChanged(topic, mqSet, allocateResultSet);
}

private boolean updateProcessQueueTableInRebalance(final String topic, final Set<MessageQueue> mqSet,
   final boolean isOrder) {
   boolean changed = false;

   Iterator<Entry<MessageQueue, ProcessQueue>> it = this.processQueueTable.entrySet().iterator();
   while (it.hasNext()) {
       Entry<MessageQueue, ProcessQueue> next = it.next();
       MessageQueue mq = next.getKey();
       ProcessQueue pq = next.getValue();

       if (mq.getTopic().equals(topic)) {
           // 消费者有下线的
           if (!mqSet.contains(mq)) {
               pq.setDropped(true);
               if (this.removeUnnecessaryMessageQueue(mq, pq)) {
                   it.remove();changed = true;
               }
           } else if (pq.isPullExpired()) {  // 拉取请求超时
               switch (this.consumeType()) {
                   case CONSUME_ACTIVELY:
                       break;
                   case CONSUME_PASSIVELY:
                       pq.setDropped(true);
                       if (this.removeUnnecessaryMessageQueue(mq, pq)) {
                           it.remove();
                           changed = true;
                       } break;
                   default: break;
               }
           }
       }
   }
   List<PullRequest> pullRequestList = new ArrayList<PullRequest>();
   for (MessageQueue mq : mqSet) {  // 增加 不在processQueueTable && 存在于mqSet 里的消息队列。
       if (!this.processQueueTable.containsKey(mq)) {
           this.removeDirtyOffset(mq);
           ProcessQueue pq = new ProcessQueue();
           long nextOffset = this.computePullFromWhere(mq);
           if (nextOffset >= 0) { 
               ProcessQueue pre = this.processQueueTable.putIfAbsent(mq, pq);
               PullRequest pullRequest = new PullRequest();
               pullRequest.setConsumerGroup(consumerGroup);pullRequest.setNextOffset(nextOffset);
               pullRequest.setMessageQueue(mq); pullRequest.setProcessQueue(pq); pullRequestList.add(pullRequest);
               pullRequestList.
           } 
       }
   }
   this.dispatchPullRequest(pullRequestList);  //进行下一次长轮询 
   return changed;
}     
```

针对上面 1和2 情况，当消费者下线或者 拉取请求超时。
需要移除不需要的消息队列相关的信息，并返回成功。其中还涉及到 顺序消息，后续在讨论。

```
public boolean removeUnnecessaryMessageQueue(MessageQueue mq, ProcessQueue pq) {
   this.defaultMQPushConsumerImpl.getOffsetStore().persist(mq);
   this.defaultMQPushConsumerImpl.getOffsetStore().removeOffset(mq);
   return true;
}
```

### 总结：
#### 1、根据上述源码分析可以看出RocketMQ 触发 重平衡条件？
当有新的消费者加入消费组；消费组成员中其中一个下线或者异常；消费者拉取请求超时导致重平衡；

#### 2、重平衡后会导致消息重复消费吗？
如果一个队列在重平衡前分配给了消费者c1，那c1在处理消息，但还没提交位点，然后重平衡后分给c2，
会从消息队列中持久化的进度开始消费，从而导致没有被持久化的位点再次被消费。
所以消费端做好幂等操作是很重要的。顺序消息则另说。

#### 3、下图为整个 重平衡线程 运行的整个流程，可以跟踪查看。



















