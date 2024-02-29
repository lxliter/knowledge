## mybatis+shardingJdbc实现数据库读写分离和分库分表

### 文章目录
- 一、原理介绍
- 二、环境准备
  - 2.1 数据库环境
  - 2.2 开发环境
  - 2.2.1 pom.xml
  - 2.2.2 建表语句
- 三、主要代码
  - 3.1 实体
  - 3.2 Mapper
  - 3.3 Controller
- 四、配置
  - 4.1 读写分离的配置
  - 4.2 多库多表的配置
  - 4.3 单库多表的配置
  - 4.4 根据自定义类配置分片规则

### 一、原理介绍
下面这篇讲的很完整就不赘述了
MySQL数据库的读写分离、分库分表

### 二、环境准备
#### 2.1 数据库环境
读写分离必须依赖数据的主从复制，这篇博客有详细的过程
mysql主从安装

#### 2.2 开发环境
本案例使用springboot2.6.4+mybatis3.6.7+sharding-jdbc4.1.1

##### 2.2.1 pom.xml
```xml
<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd">
    <modelVersion>4.0.0</modelVersion>
    <parent>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-starter-parent</artifactId>
        <version>2.6.4</version>
        <relativePath/> <!-- lookup parent from repository -->
    </parent>
    <groupId>com.yex</groupId>
    <artifactId>sharding-jdbc</artifactId>
    <version>0.0.1-SNAPSHOT</version>
    <name>sharding-jdbc</name>
    <description>Demo project for Spring Boot</description>
    <properties>
        <java.version>1.8</java.version>
        <sharding-sphere.version>4.1.1</sharding-sphere.version>
    </properties>
    <dependencies>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-web</artifactId>
        </dependency>
        <dependency>
            <groupId>org.mybatis.spring.boot</groupId>
            <artifactId>mybatis-spring-boot-starter</artifactId>
            <version>2.2.2</version>
        </dependency>
        <dependency>
            <groupId>mysql</groupId>
            <artifactId>mysql-connector-java</artifactId>
            <scope>runtime</scope>
        </dependency>
        <!--数据库连接池，不能使用druid-spring-boot-starter，会报错
        Property 'sqlSessionFactory' or 'sqlSessionTemplate' are required-->
        <dependency>
            <groupId>com.alibaba</groupId>
            <artifactId>druid</artifactId>
            <version>1.1.22</version>
        </dependency>
        <dependency>
            <groupId>org.projectlombok</groupId>
            <artifactId>lombok</artifactId>
            <optional>true</optional>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-test</artifactId>
            <scope>test</scope>
        </dependency>
        <!--分库分表工具-->
        <dependency>
            <groupId>org.apache.shardingsphere</groupId>
            <artifactId>sharding-jdbc-spring-boot-starter</artifactId>
            <version>${sharding-sphere.version}</version>
        </dependency>
        <dependency>
            <groupId>org.apache.shardingsphere</groupId>
            <artifactId>sharding-core-common</artifactId>
            <version>${sharding-sphere.version}</version>
        </dependency>
    </dependencies>
    <build>
        <plugins>
            <plugin>
                <groupId>org.springframework.boot</groupId>
                <artifactId>spring-boot-maven-plugin</artifactId>
                <configuration>
                    <excludes>
                        <exclude>
                            <groupId>org.projectlombok</groupId>
                            <artifactId>lombok</artifactId>
                        </exclude>
                    </excludes>
                </configuration>
            </plugin>
        </plugins>
    </build>
</project>
```

##### 2.2.2 建表语句
```sql
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `birthday` date NULL DEFAULT NULL,
  `sex` bit(1) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
```

### 三、主要代码
service不需要

#### 3.1 实体
```java
package com.yex.shardingjdbc.entity;
import lombok.Data;
import java.util.Date;

@Data
public class User {
    private Long id;

    private String name;

    private String password;

    private int sex;

    private Date birthday;
}
```

#### 3.2 Mapper
```java
package com.yex.shardingjdbc.mapper;

import com.yex.shardingjdbc.entity.User;
import org.apache.ibatis.annotations.Insert;
import org.apache.ibatis.annotations.Param;
import org.apache.ibatis.annotations.Select;

import java.util.List;

public interface UserMapper {
    @Insert("insert into user(id,name,password,sex,birthday) values(#{user.id},#{user.name},#{user.password},#{user.sex},#{user.birthday})")
    void addUser(@Param("user") User user);

    @Select("select * from user")
    List<User> findUsers();
}
```

#### 3.3 Controller
```java
package com.yex.shardingjdbc.controller;

import com.yex.shardingjdbc.entity.User;
import com.yex.shardingjdbc.mapper.UserMapper;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import javax.annotation.Resource;
import java.util.Date;
import java.util.List;
import java.util.Random;

@RestController
@RequestMapping("/user")
public class UserController {
    @Resource
    private UserMapper userMapper;

    @GetMapping("/save")
    public User addUser() {
        User user = new User();
        user.setId(System.currentTimeMillis());
        user.setName("user" + new Random().nextInt());
        user.setPassword("123456");
        user.setSex(1);
        user.setBirthday(new Date());
        userMapper.addUser(user);
        return user;
    }

    @GetMapping("/list")
    public List<User> listUser() {
        return userMapper.findUsers();
    }
}
```

### 四、配置

#### 4.1 读写分离的配置
```yaml
server:
  port: 8085
#整合mybatis的配置
mybatis:
  mapper-locations: classpath:mapper/*.xml

spring:
  main:
    allow-bean-definition-overriding: true
  shardingsphere:
    #参数配置，显示sql
    props:
      sql:
        show: true
    datasource:
      # 给每个数据源取别名
      names: ds0,ds1
      # 给每个数据源配置数据库连接信息
      ds0:
        type: com.alibaba.druid.pool.DruidDataSource
        driver-class-name: com.mysql.cj.jdbc.Driver
        url: jdbc:mysql://192.168.136.133:3307/db01?useUnicode=true&useSSL=false&characterEncoding=utf8&serverTimezone=Asia/Shanghai
        username: root
        password: root
        maxPoolSize: 100
        minPoolSize: 5
      ds1:
        type: com.alibaba.druid.pool.DruidDataSource
        driver-class-name: com.mysql.cj.jdbc.Driver
        url: jdbc:mysql://192.168.136.133:3308/db01?useUnicode=true&useSSL=false&characterEncoding=utf8&serverTimezone=Asia/Shanghai
        username: root
        password: root
        maxPoolSize: 100
        minPoolSize: 5
    # 配置默认数据源ds0
    sharding:
      # 默认数据源，用于写，不配置写库，会把两个数据源都当作从库，增删改会报错
      default-data-source-name: ds0
    # 配置数据源的读写分离，必须要配置数据库的主从复制
    masterslave:
      # 配置主从名称，任意取名
      name: ms
      # 配置主库master节点，负责数据的写入
      master-data-source-name: ds0
      # 配置从库slave节点
      slave-data-source-names: ds1
      # 配置从库的负载均衡策略，轮询
      load-balance-algorithm-type: round_robin
```

#### 4.2 多库多表的配置
需要在每个库建立两张表fsd_user0、fsd_user1
```yaml
server:
  port: 8085
#整合mybatis的配置
mybatis:
  mapper-locations: classpath:mapper/*.xml

spring:
  main:
    allow-bean-definition-overriding: true
  shardingsphere:
    #参数配置，显示sql
    props:
      sql:
        show: true
    datasource:
      # 给每个数据源取别名
      names: ds0,ds1
      # 给每个数据源配置数据库连接信息
      ds0:
        type: com.alibaba.druid.pool.DruidDataSource
        driver-class-name: com.mysql.cj.jdbc.Driver
        url: jdbc:mysql://192.168.136.133:3307/db01?useUnicode=true&useSSL=false&characterEncoding=utf8&serverTimezone=Asia/Shanghai
        username: root
        password: root
        maxPoolSize: 100
        minPoolSize: 5
      ds1:
        type: com.alibaba.druid.pool.DruidDataSource
        driver-class-name: com.mysql.cj.jdbc.Driver
        url: jdbc:mysql://192.168.136.133:3308/db01?useUnicode=true&useSSL=false&characterEncoding=utf8&serverTimezone=Asia/Shanghai
        username: root
        password: root
        maxPoolSize: 100
        minPoolSize: 5
    sharding:
      #配置分表规则
      tables:
        #逻辑表名
        fsd_user:
          #数据节点：数据源$->{0..N}.表名$->{0..N}
          actual-data-nodes: ds$->{0..1}.fsd_user$->{0..1}
          #数据库分片策略，也就是什么样的数据放到哪个库中
          database-strategy:
            inline:
              sharding-column: id                  # 分片字段
              algorithm-expression: ds$->{id%2} # 分片算法表达式
          #表分片策略
          table-strategy:
            inline:
              sharding-column: sex                 # 分片字段
              algorithm-expression: fsd_user$->{sex}

```

> 分库分表支持将数据分片到多个库的多个表，同时还可以主从复制。
> 当数据分到主库master.user01，从库slave.user01就会复制一份，
> 其他数据可能直接分片到slave.user01，查询的时候就会有重复数据。
> 因此要先关闭数据库主从复制功能，同时也不建议将一张表分成多库多表，单库多表即可。
> 如果表特别多，可以根据业务拆分多个微服务使用不同数据库，达到分库的目的。

#### 4.3 单库多表的配置
单库多表，再配合主从复制的读写分离，这样的配置是很合理的
```yaml
server:
  port: 8085
#整合mybatis的配置
mybatis:
  mapper-locations: classpath:mapper/*.xml
spring:
  main:
    allow-bean-definition-overriding: true
  shardingsphere:
    #参数配置，显示sql
    props:
      sql:
        show: true
    datasource:
      # 给每个数据源取别名
      names: ds0,ds1
      # 给每个数据源配置数据库连接信息
      ds0:
        type: com.alibaba.druid.pool.DruidDataSource
        driver-class-name: com.mysql.cj.jdbc.Driver
        url: jdbc:mysql://192.168.136.133:3307/db01?useUnicode=true&useSSL=false&characterEncoding=utf8&serverTimezone=Asia/Shanghai
        username: root
        password: root
        maxPoolSize: 100
        minPoolSize: 5
      ds1:
        type: com.alibaba.druid.pool.DruidDataSource
        driver-class-name: com.mysql.cj.jdbc.Driver
        url: jdbc:mysql://192.168.136.133:3308/db01?useUnicode=true&useSSL=false&characterEncoding=utf8&serverTimezone=Asia/Shanghai
        username: root
        password: root
        maxPoolSize: 100
        minPoolSize: 5
    sharding:
      #配置分表规则
      tables:
        #逻辑表名
        fsd_user:
          #数据节点：数据源$->{0..N}.表名$->{0..N}
          actual-data-nodes: ds0.fsd_user$->{0..1}
          #表分片策略
          table-strategy:
            inline:
              sharding-column: sex                 # 分片字段
              algorithm-expression: fsd_user$->{sex}

```

#### 4.4 根据自定义类配置分片规则
这里用单库多表测试分片
```yaml
server:
  port: 8085
#整合mybatis的配置
mybatis:
  mapper-locations: classpath:mapper/*.xml

spring:
  main:
    allow-bean-definition-overriding: true
  shardingsphere:
    #参数配置，显示sql
    props:
      sql:
        show: true
    datasource:
      # 给每个数据源取别名
      names: ds0,ds1
      # 给每个数据源配置数据库连接信息
      ds0:
        type: com.alibaba.druid.pool.DruidDataSource
        driver-class-name: com.mysql.cj.jdbc.Driver
        url: jdbc:mysql://192.168.136.133:3307/db01?useUnicode=true&useSSL=false&characterEncoding=utf8&serverTimezone=Asia/Shanghai
        username: root
        password: root
        maxPoolSize: 100
        minPoolSize: 5
      ds1:
        type: com.alibaba.druid.pool.DruidDataSource
        driver-class-name: com.mysql.cj.jdbc.Driver
        url: jdbc:mysql://192.168.136.133:3308/db01?useUnicode=true&useSSL=false&characterEncoding=utf8&serverTimezone=Asia/Shanghai
        username: root
        password: root
        maxPoolSize: 100
        minPoolSize: 5
    sharding:
      #配置分表规则
      tables:
        #逻辑表名
        fsd_user:
          # 自定义id算法
          key-generator:
            type: SNOWFLAKE
            column: id
          #数据节点：数据源$->{0..N}.表名$->{0..N}
          actual-data-nodes: ds0.fsd_user$->{0..1}
          #表分片策略
          table-strategy:
            # 根据自定义类
            standard:
              sharding-column: birthday                 # 分片字段
              preciseAlgorithmClassName: com.yex.shardingjdbc.algorithm.BirthdayAlgorithm


```
自定义类代码
```java
package com.yex.shardingjdbc.algorithm;

import org.apache.shardingsphere.api.sharding.standard.PreciseShardingAlgorithm;
import org.apache.shardingsphere.api.sharding.standard.PreciseShardingValue;

import java.util.Calendar;
import java.util.Collection;
import java.util.Date;
import java.util.Iterator;


public class BirthdayAlgorithm implements PreciseShardingAlgorithm<Date> {
    @Override
    public String doSharding(Collection<String> collection, PreciseShardingValue<Date> preciseShardingValue) {
        Calendar calendar = Calendar.getInstance();
        Date date = preciseShardingValue.getValue();
        Iterator<String> iterator =collection.iterator();
        calendar.setTime(date);
        //生日单数和双数在不同的表
        if(calendar.get(Calendar.DAY_OF_MONTH)%2==0){
            return iterator.next();
        }else {
            iterator.next();
            return iterator.next();
        }
    }
}
```










