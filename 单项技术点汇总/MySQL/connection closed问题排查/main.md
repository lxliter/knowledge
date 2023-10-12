## Cause: java.sql.SQLException: connection closed问题排查、解决

connection closed ***获取到的连接已经失效，导致抛出异常***：
message:com.framework.smart.admin.common.exception.AdminExceptionHandler.handleException:76 - 【exception】：

org.springframework.jdbc.UncategorizedSQLException:

- connection closed
- 获取到的连接已经失效，导致抛出异常

Error querying database. Cause: java.sql.SQLException: connection closed

The error may exist in URL 
[jar:file:/usr/local/smart-admin/smart-admin.jar!/BOOT-INF/lib/smart-admin-service-0.0.1-SNAPSHOT.jar!/mapper/MmChanceMapper.xml]

The error may involve com.framework.smart.admin.service.mapper.MmChanceMapper.selectChanceReport_COUNT

The error occurred while executing a query

SQL: SELECT count(0) FROM mm_chance WHERE state BETWEEN 0 AND 1 AND is_deleted = 0

Cause: java.sql.SQLException: connection closed

; uncategorized SQLException; SQL state [null]; error code [0]; connection closed; nested exception is java.sql.SQLException: connection closed
at org.springframework.jdbc.support.AbstractFallbackSQLExceptionTranslator.translate(AbstractFallbackSQLExceptionTranslator.java:89)
at org.springframework.jdbc.support.AbstractFallbackSQLExceptionTranslator.translate(AbstractFallbackSQLExceptionTranslator.java:81)
at org.springframework.jdbc.support.AbstractFallbackSQLExceptionTranslator.translate(AbstractFallbackSQLExceptionTranslator.java:81)

```
Caused by: java.sql.SQLException: connection closed
at com.alibaba.druid.pool.DruidPooledConnection.checkStateInternal(DruidPooledConnection.java:1163)
at com.alibaba.druid.pool.DruidPooledConnection.checkState(DruidPooledConnection.java:1154)
at com.alibaba.druid.pool.DruidPooledConnection.prepareStatement(DruidPooledConnection.java:337)
```

### 1.testOnBorrow含义
testOnBorrow：如果为true（默认为false），当应用向连接池申请连接时，连接池会判断这条连接是否是可用的。
testOnBorrow=false可能导致问题
假如连接池中的连接被数据库关闭了，应用通过连接池getConnection时，都可能获取到这些不可用的连接，且这些连接如果不被其他线程回收的话；
它们不会被连接池废除，也不会重新被创建，占用了连接池的名额，
项目如果是服务端，数据库链接被关闭，客户端调用服务端就会出现大量的timeout，
***客户端设置了超时时间，会主动断开，服务端就会出现close_wait***。

连接池如何判断连接是否有效的？
常用数据库：使用${DBNAME}ValidConnectionChecker进行判断，比如Mysql数据库，
使用MySqlValidConnectionChecker的isValidConnection进行判断

其他数据库：则使用validationQuery判断
验证不通过则会直接关闭连接，并重新从连接池获取下一条连接。

综上：
<1>.testOnBorrow能够确保我们每次都能获取到可用的连接，但是如果设置为true，则每次获取连接时候都要到数据库验证连接有效性，
这在高并发的时候会造成性能下降，可以将testOnBorrow设置成false，testWhileIdle设置成true这样能获得比较好的性能。
<2>.testOnBorrow和testOnReturn在生产环境一般是不开启的，主要是性能考虑。
失效连接主要通过testWhileIdle保证，如果获取到了不可用的数据库连接，一般由应用处理异常。

TestOnBorrow什么时候会用到？
这个参数主要在DruidDataSource的getConnection方法中用到

连接池是如何判断连接是否有效的？
如果是常用的数据库，则使用${DBNAME}ValidConnectionChecker进行判断，比如Mysql数据库，使用MySqlValidConnectionChecker的isValidConnection进行判断；
如果是其他数据库，则使用validationQuery判断；

如果验证不通过怎么办？
验证不通过则会直接关闭该连接，并重新从连接池获取下一条连接；

获取到连接后：在DruidPooledConnection中

testOnBorrow能够确保我们每次都能获取到可用的连接，但如果设置成true，则每次获取连接的时候都要到数据库验证连接有效性，
这在高并发的时候会造成性能下降，可以将testOnBorrow设成false，testWhileIdle设置成true这样能获得比较好的性能。

validationQuery是什么意思？
validationQuery：Druid用来测试连接是否可用的SQL语句,默认值每种数据库都不相同：

```sql
Mysql:SELECT 1;
SQLSERVER:SELECT 1;
ORACLE:SELECT 'x' FROM DUAL;
PostGresql:SELECT 'x';
```

validationQuery什么时候会起作用？
当Druid遇到testWhileIdle，testOnBorrow，testOnReturn时，就会验证连接的有效性，验证规则如下：
如果有相关数据库的ValidConnectionChecker，则使用ValidConnectionChecker验证（Druid提供常用数据库的ValidConnectionChecker，
包括MSSQLValidConnectionChecker，MySqlValidConnectionChecker，OracleValidConnectionChecker，PGValidConnectionChecker）；

如果没有ValidConnectionChecker，则直接使用validationQuery验证；

ValidConnectionChecker是如何验证的？
MySqlValidConnectionChecker会使用Mysql独有的ping方式进行验证，其他数据库其实也都是使用validationQuery进行验证

MySqlValidConnectionChecker验证方式如上图；不同数据库的默认值不同；
如果是Mysql数据库，则validationQuery不会起作用，Mysql会使用ping的方式验证；




















