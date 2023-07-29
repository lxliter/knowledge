### MySQL8.0新特性——支持原子DDL语句
MySQL 8.0开始支持原子数据定义语言（DDL）语句。<br>
此功能称为原子DDL。原子DDL语句将与DDL操作关联的数据字典更新，存储引擎操作和二进制日志写入组合到单个原子事务中。<br>
即使服务器在操作期间暂停，也会提交事务，并将适用的更改保留到数据字典，存储引擎和二进制日志，或者回滚事务。<br>
通过在MySQL 8.0中引入MySQL数据字典，可以实现Atomic DDL。<br>
在早期的MySQL版本中，元数据存储在元数据文件，非事务性表和存储引擎特定的字典中，这需要中间提交。<br>
MySQL数据字典提供的集中式事务元数据存储消除了这一障碍，使得将DDL语句操作重组为原子事务成为可能。<br>

#### 1、支持的DDL语句
原子DDL功能支持表和非表DDL语句。与表相关的DDL操作需要存储引擎支持，而非表DDL操作则不需要。目前，只有InnoDB存储引擎支持原子DDL。<br>
①：受支持的表DDL语句包括 CREATE，ALTER和 DROP对数据库，表，表和索引，以及语句 TRUNCATE TABLE声明。<br>
②：支持的非表DDL语句包括：<br>
CREATE和DROP 语句，以及（如果适用）ALTER 存储程序，触发器，视图和用户定义函数（UDF）的语句。 <br>
账户管理语句： CREATE，ALTER， DROP，如果适用， RENAME报表用户和角色，以及GRANT 和REVOKE报表。<br>

##### 1.1、原子DDL功能不支持以下语句：
①：涉及除存储引擎之外的存储引擎的与表相关的DDL语句InnoDB。<br>
②：INSTALL PLUGIN和 UNINSTALL PLUGIN 陈述。<br>
③：INSTALL COMPONENT和 UNINSTALL COMPONENT 陈述。<br>
④：CREATE SERVER， ALTER SERVER和 DROP SERVER语句。<br>

#### 2、原子DDL特性：
①：元数据更新，二进制日志写入和存储引擎操作（如果适用）将合并为单个事务。<br>
②：在DDL操作期间，SQL层没有中间提交。<br>
③：在适用的情况下：<br>
数据字典，程序，事件和UDF高速缓存的状态与DDL操作的状态一致，这意味着更新高速缓存以反映DDL操作是成功完成还是回滚。<br>
DDL操作中涉及的存储引擎方法不执行中间提交，并且存储引擎将自身注册为DDL事务的一部分。<br>
存储引擎支持DDL操作的重做和回滚，这在DDL操作的 Post-DDL阶段执行。<br>
④：DDL操作的可见行为是原子的，这会更改某些DDL语句的行为<br>

注意：<br>
原子或其他DDL语句隐式结束当前会话中处于活动状态的任何事务，就好像您COMMIT在执行语句之前完成了一样。<br>
这意味着DDL语句不能在另一个事务中，在事务控制语句中执行 START TRANSACTION ... COMMIT，或者与同一事务中的其他语句结合使用。<br>

#### 3、DDL语句行为的变化
如果所有命名表都使用原子DDL支持的存储引擎，则操作是完全原子的。该语句要么成功删除所有表，要么回滚。<br>
DROP TABLE如果命名表不存在，并且未进行任何更改（无论存储引擎如何），则会失败并显示错误。如下所示：<br>
```
mysql> CREATE TABLE t1 (c1 INT);
mysql> DROP TABLE t1, t2;
ERROR 1051 (42S02): Unknown table 'test.t2'
mysql> SHOW TABLES;
+----------------+
| Tables_in_test |
+----------------+
| t1
+----------------+
```

在引入原子DDL之前， DROP TABLE虽然会报错误表不存在，但是存在的表会被执行成功，如下：<br>
```
mysql> CREATE TABLE t1 (c1 INT);
mysql> DROP TABLE t1, t2;
ERROR 1051 (42S02): Unknown table 'test.t2'
mysql> SHOW TABLES;
Empty set (0.00 sec)
```

注意：<br>
由于行为的这种变化，DROP TABLE会在 MySQL 5.7主服务器上的部分完成 语句在MySQL 8.0从服务器上复制时失败。<br>
要避免此故障情形，请在DROP TABLE语句中使用IF EXISTS语法以防止对不存在的表发生错误<br>

#### 4、存储引擎支持：目前只有innodb存储引擎支持原子DDL
目前，只有InnoDB存储引擎支持原子DDL。不支持原子DDL的存储引擎免于DDL原子性。涉及豁免存储引擎的DDL操作仍然能够引入操作中断或仅部分完成时可能发生的不一致。
要支持重做和回滚DDL操作， InnoDB请将DDL日志写入 mysql.innodb_ddl_log表，该表是驻留在mysql.ibd数据字典表空间中的隐藏数据字典表 。
要mysql.innodb_ddl_log在DDL操作期间查看写入表的DDL日志 ，请启用 innodb_print_ddl_logs 配置选项。
注意：
mysql.innodb_ddl_log无论innodb_flush_log_at_trx_commit 设置多少，对表的 更改的重做日志 都会立即刷新到磁盘 。
立即刷新重做日志可以避免DDL操作修改数据文件的情况，但是mysql.innodb_ddl_log由这些操作产生的对表的更改的重做日志 不会持久保存到磁盘。
这种情况可能会在回滚或恢复期间导致错误。
InnoDB存储引擎分阶段执行DDL操作。DDL操作 ALTER TABLE可以在Commit阶段之前多次执行 Prepare和Perform阶段：
准备：创建所需对象并将DDL日志写入 mysql.innodb_ddl_log表中。DDL日志定义了如何前滚和回滚DDL操作。
执行：执行DDL操作。例如，为CREATE TABLE操作执行创建例程。
提交：更新数据字典并提交数据字典事务。
Post-DDL：重播并从mysql.innodb_ddl_log表中删除DDL日志。为了确保可以安全地执行回滚而不引入不一致性，在最后阶段执行文件操作，例如重命名或删除数据文件。这一阶段还从删除的动态元数据 mysql.innodb_dynamic_metadata的数据字典表DROP TABLE，TRUNCATE TABLE和该重建表其他DDL操作。
注意：
无论事务是提交还是回滚， DDL日志都会在Post-DDL阶段重播并从表中删除 。mysql.innodb_ddl_log如果服务器在DDL操作期间暂停，则DDL日志应仅保留在表中。在这种情况下，DDL日志将在恢复后重播并删除。
在恢复情况下，可以在重新启动服务器时提交或回滚DDL事务。如果在重做日志和二进制日志中存在在DDL操作的提交阶段期间执行的数据字典事务，则 该操作被视为成功并且前滚。否则，在InnoDB重放数据字典重做日志时回滚不完整的数据字典事务 ，并回滚DDL事务。

#### 5、查看DDL日志：
InnoDB将DDL日志写入 mysql.innodb_ddl_log表以支持重做和回滚DDL操作。
该 mysql.innodb_ddl_log表是隐藏在mysql.ibd数据字典表空间中的隐藏数据字典表 。
与其他隐藏数据字典表一样，mysql.innodb_ddl_log在非调试版本的MySQL中无法直接访问该表。


















