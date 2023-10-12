## MySQL锁等待超时

***一次代码报错的分析：MySQLTransactionRollbackException: Lock wait timeout exceeded; try restarting transaction***

- MySQLTransactionRollbackException
- Lock wait timeout exceeded; try restarting transaction

### 1. 报错信息

```
### Cause: com.mysql.cj.jdbc.exceptions.MySQLTransactionRollbackException: Lock wait timeout exceeded; 
try restarting transaction; SQL []; Lock wait timeout exceeded; 
try restarting transaction; 
nested exception is com.mysql.cj.jdbc.exceptions.MySQLTransactionRollbackException: 
Lock wait timeout exceeded; try restarting transaction
```

### 2. 报错原因

- 1.两个事务之间出现死锁，导致另外一个事物超时
- 2.某一种表频繁被锁表，导致其他事务无法拿到锁，导致事物超时

### 3. 问题过程

- 1.查看数据库当前的进程，看一下有无正在执行的慢SQL记录线程。

```
mysql> show  processlist;
```

- 2.查看当前的事务

```
当前运行的所有事务：
mysql> SELECT * FROM information_schema.INNODB_TRX;
```

```
当前出现的锁
mysql> SELECT * FROM information_schema.INNODB_LOCKs;
```

```
锁等待的对应关系
mysql> SELECT * FROM information_schema.INNODB_LOCK_waits;
```

- 3.解释：***看事务表INNODB_TRX，里面是否有正在锁定的事务线程***，看看ID是否在show processlist里面的sleep线程中，
  如果是，就证明这个sleep的线程事务一直没有commit或者rollback而是卡住了，我们需要手动kill掉。
  搜索的结果是在事务表发现了很多任务，这时候最好都kill掉。

> 查看事务表INNODB_TRX，里面是否有正在锁定的事务线程
> 看看ID是否在show processlist里面的sleep线程中
> 如果是，就证明这个sleep的线程事务一直没有commit或者rollback，而是卡主了，需要手动kill掉
> 搜索的结果是在事务表发现了很多任务，这时候最好都kill掉

- 4.批量删除事务表中的事务
  我这里用的方法是：通过information_schema.processlist表中的连接信息生成需要处理掉的MySQL连接的语句临时文件，
  然后执行临时文件中生成的指令。

```
mysql>  select concat('KILL ',id,';') from information_schema.processlist where user='cms_bokong';
+------------------------+
| concat('KILL ',id,';') |
+------------------------+
| KILL 10508;            |
| KILL 10521;            |
| KILL 10297;            |
+------------------------+
18 rows in set (0.00 sec)
```

当然结果不可能只有3个，这里我只是举例子。
参考链接上是建议导出到一个文本，然后执行文本。而我是直接copy到记事本处理掉 ‘|’，粘贴到命令行执行了。都可以。
kill掉以后再执行SELECT * FROM information_schema.INNODB_TRX; 就是空了。
这时候系统就正常了
故障排查

```
mysql都是autocommit配置
mysql> select @@autocommit;
+--------------+
| @@autocommit |
+--------------+
|            0 |
+--------------+
1 row in set (0.00 sec)
```

参考：MySQL事务锁问题-Lock wait timeout exceeded

### 4. 问题解决

经过了解业务信息，一次性执行500次的UPDATE操作，一次UPDATE耗时100毫秒，导致表被锁50秒。
为被更新操作增加索引之后，此问题解决，执行速度加快。

### show processlist 详解

#### 一、show processlist 简介

show processlist是显示用户正在运行的线程，需要注意的是，除了root用户能看到所有正在运行的线程外，
其他用户都只能看到自己正在运行的线程，看不到其它用户正在运行的线程。
除非单独给这个用户赋予了PROCESS权限。

- 除非单独给这个用户赋予了PROCESS权限

通常我们通过top检查发现mysqlCPU或者iowait过高 那么解决这些问题 都离不开show processlist查询当前mysql有些线程正在运行,
然后分析其中的参数,找出那些有问题的线程,该kill的kill,该优化的优化!

- 通过top检查发现MySQL CPU或者IOWait过高
- Show processlist查询当前mysql有哪些线程正在运行
- 然后分析其中的参数，找出那些有问题的线程，该kill的kill，该优化的优化

注意: show processlist只显示前100条 我们可以通过show full processlist 显示全部。

- show processlist只显示前100条
- show full processlist显示全部
- root用户，可以看到全部线程运行情况

```
mysql> show processlist ;
+---------+-------------+---------------------+--------------------+---------+------+--------------+------------------------------------------------------------------------------------------------------+
| Id      | User        | Host                | db                 | Command | Time | State        | Info                                                                                                 |
+---------+-------------+---------------------+--------------------+---------+------+--------------+------------------------------------------------------------------------------------------------------+
| 9801429 | lis         | 10.41.7.10:57962    | lis                | Query   |    8 | Sending data | SELECT r.RgtantMobile, r.RgtantName, r.RgtNo, ag.SumGetMoney, ag.EnterAccDate, ag.BankAccNo, ( SELEC |
| 9802020 | ruihua      | 10.41.5.6:37543     | sales_org          | Sleep   |  292 |              | NULL                                                                                                 |
| 9802070 | lis         | 10.41.7.10:58998    | lis                | Query   |    8 | Sending data | select distinct d.contno,e.phone,d.appntidtype,d.appntidno,d.appntname,d.appntsex ,d.AppntBirthday,( |
| 9802084 | evoiceadmin | 10.41.8.8:41868     | evoicerh           | Sleep   |   57 |              | NULL                                                                                                 |
| 9802201 | root        | 10.41.100.3:38976   | NULL               | Query   |    0 | init         | show processlist                                                                                     |
+---------+-------------+---------------------+--------------------+---------+------+--------------+------------------------------------------------------------------------------------------------------+
148 rows in set (0.00 sec)
```

普通的lis用户只能看到自己的

```
mysql> show processlist ;
+---------+-------------+---------------------+--------------------+---------+------+--------------+------------------------------------------------------------------------------------------------------+
| Id      | User        | Host                | db                 | Command | Time | State        | Info                                                                                                 |
+---------+-------------+---------------------+--------------------+---------+------+--------------+------------------------------------------------------------------------------------------------------+
| 9801429 | lis         | 10.41.7.10:57962    | lis                | Query   |    8 | Sending data | SELECT r.RgtantMobile, r.RgtantName, r.RgtNo, ag.SumGetMoney, ag.EnterAccDate, ag.BankAccNo, ( SELEC |
| 9802070 | lis         | 10.41.7.10:58998    | lis                | Query   |    8 | Sending data | select distinct d.contno,e.phone,d.appntidtype,d.appntidno,d.appntname,d.appntsex ,d.AppntBirthday,( |
+---------+-------------+---------------------+--------------------+---------+------+--------------+------------------------------------------------------------------------------------------------------+
2 rows in set (0.00 sec)
```

单独给普通的activiti用户授PROCESS权限，（授权后需要退出重新登录）
show processlist显示的信息都是来自MySQL系统库information_schema中的processlist表。所以使用下面的查询语句可以获得相同的结果：

```
select * from information_schema.processlist
```

- show processlist显示的信息都是来自MySQL系统库information_schema中的processlist表




