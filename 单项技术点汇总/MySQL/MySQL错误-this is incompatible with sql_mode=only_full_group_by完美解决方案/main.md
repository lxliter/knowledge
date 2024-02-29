## MySQL错误-this is incompatible with sql_mode=only_full_group_by完美解决方案

### 项目场景：
有时候，遇到数据库重复数据，需要将数据进行分组，并取出其中一条来展示，这时就需要用到group by语句。
但是，如果mysql是高版本，当执行group by时，select的字段不属于group by的字段的话，sql语句就会报错。报错信息如下：
```
Expression #1 of SELECT list is not in GROUP BY clause and contains
nonaggregated column ‘数据库名.表名.字段名’ which is not functionally dependent
on columns in GROUP BY clause; this is incompatible with
sql_mode=only_full_group_by
```

### 问题描述：
#### 1.表结构
```sql
CREATE TABLE `t_iov_help_feedback` (
  `ID` INT(11) NOT NULL  AUTO_INCREMENT COMMENT '主键ID',
  `USER_ID` INT(255) DEFAULT NULL  COMMENT '用户ID',
  `problems` VARCHAR(255) DEFAULT NULL COMMENT '问题描述',
  `last_updated_date` DATETIME DEFAULT NULL COMMENT '最后更新时间',
  PRIMARY KEY (`ID`)
) ENGINE=INNODB DEFAULT CHARSET=utf8;
```

#### 2.表数据
![img.png](imgs/table_data.png)

#### 3.sql语句
- 1）查询group by的字段（正常）
只有分组字段
```sql
SELECT USER_ID  FROM  t_iov_help_feedback GROUP BY USER_ID;
```
![img.png](imgs/group_by_normal.png)

聚合函数（其它字段）+分组字段
![img.png](imgs/group_by_normal1.png)

- 2）查询非group by的字段（报错）
![img.png](imgs/group_by_abnormal.png)

报错什么意思呢？
一句话概括：“错误代码1055与sql_mode = only_full_group_by不兼容”
翻译：
```
“错误代码：1055。SELECT列表的表达式＃1不在GROUP BY子句中，
并且包含非聚合列’test.t_iov_help_feedback.ID’，它在功能上不依赖于GROUP BY子句中的列; 
这与sql_mode = only_full_group_by不兼容”
```

### 原因分析：
```
一、原理层面
这个错误发生在mysql 5.7.5 版本及以上版本会出现的问题：
mysql 5.7.5版本以上默认的sql配置是:sql_mode=“ONLY_FULL_GROUP_BY”，这个配置严格执行了"SQL92标准"。
很多从5.6升级到5.7时，为了语法兼容，大部分都会选择调整sql_mode，使其保持跟5.6一致，为了尽量兼容程序。
```

```
二、sql层面
在sql执行时，出现该原因，简单来说就是：
由于开启了ONLY_FULL_GROUP_BY的设置，如果select 的字段不在 group by 中，
并且select 的字段未使用聚合函数（SUM,AVG,MAX,MIN等）的话，那么这条sql查询是被mysql认为非法的，会报错误…
```
验证是否此原因：

- 1.查询数据库版本的语句
```sql
SELECT VERSION();
```
可以看到，我这里数据库版本是：8.0.16，大于5.7.5了

- 2.查看sql_mode的语句
```sql
select @@GLOBAL.sql_mode;
```
```
ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION
```
可以看到，***sql_mode开启了only_full_group_by属性***

### 解决方案：
#### 解决方案一：使用函数ANY_VALUE()包含报错字段
将上述报错语句改成：
```
SELECT ANY_VALUE(ID),USER_ID,ANY_VALUE(problems),ANY_VALUE(last_updated_date) FROM  t_iov_help_feedback GROUP BY USER_ID;
```
可以看到，结果能正常查询了，根据需要自己改查询字段的别名就行。

ANY_VALUE()函数说明：
```
MySQL有any_value(field)函数，它主要的作用就是抑制ONLY_FULL_GROUP_BY值被拒绝。
这样sql语句不管是在ONLY_FULL_GROUP_BY模式关闭状态还是在开启模式都可以正常执行，不被mysql拒绝。
any_value()会选择被分到同一组的数据里第一条数据的指定列值作为返回数据。
官方有介绍，地址：https://dev.mysql.com/doc/refman/5.7/en/miscellaneous-functions.html#function_any-value
```

#### 解决方案二：通过sql语句暂时性修改sql_mode
去掉ONLY_FULL_GROUP_BY，重新设置值
```sql
SET @@global.sql_mode ='STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';
```

上面是改变了全局sql_mode，对于新建的数据库有效。对于已存在的数据库，则需要在对应的数据库下执行：
```sql
SET sql_mode ='STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';
```

问题：
```
重启mysql数据库服务之后，ONLY_FULL_GROUP_BY还会出现，所以这只是暂时性的。
```

备注：
网上有些朋友提供的sql语句如下：
```
set @@GLOBAL.sql_mode=STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION
```
但是却执行不了，报sql语法错误：
这时只需要加上单引号即可：
但是，添加了单引号仍然报错：
这里说sql_mode不能设置NO_AUTO_CREATE_USER这个值，那直接去掉这个值就行了呗，也就是上面我提供的值。

#### 解决方案三：通过配置文件永久修改sql_mode
mysql安装在服务器上和安装在本地，修改配置文件的方式有点区别。
- 1、Linux下修改配置文件
1）登录进入MySQL
使用命令 mysql -u username -p 进行登陆，然后输入密码，输入SQL：
```sql
show variables like '%sql_mode';
```

2）编辑my.cnf文件
文件地址一般在：***/etc/my.cnf***，***/etc/mysql/my.cnf***
使用vim命令编辑文件，不知道vim命令怎么使用的，可以参考我的另外篇文章：Linux中使用vi工具进行文本编辑
找到sql-mode的位置，去掉ONLY_FULL_GROUP_BY
然后重启MySQL；
有的my.cnf中可能没有sql-mode，需要追加：
```
sql-mode=STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION
```
注意要加入到[mysqld]下面，如加入到其他地方，重启后也不生效，具体的如下图：

3）修改成功后重启MySQL服务
```
service mysql restart
```

重启好后，再登录mysql，输入SQL：show variables like '%sql_mode''; 
如果没有ONLY_FULL_GROUP_BY，就说明已经成功了。

如果还不行，那么只保留STRICT_TRANS_TABLES,NO_ENGINE_SUBSTITUTION 即可
追加内容为：
```
sql-mode=STRICT_TRANS_TABLES,NO_ENGINE_SUBSTITUTION
```

- 2、window下修改配置文件
1）找到mysql安装目录，用记事本直接打开my.ini文件
2）编辑my.ini文件，在[mysql]标签下追加内容
3）重启mysql 服务
备注：
网上有些提供了sql_mode的值，却导致重启mysql服务启动不了
这时，只需要将sql_mode 值中 “NO_AUTO_CREATE_USER” 这个属性去掉即可。