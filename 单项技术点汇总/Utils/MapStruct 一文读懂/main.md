## MapStruct 一文读懂

### 目录
- 1、MapStruct 简介
  - 1.1 MapStruct Maven引入
- 2、MapStruct 基础操作
  - 2.1 MapStruct 基本映射
  - 2.2 MapStruct 指定默认值
  - 2.3 MapStruct 表达式
  - 2.4 MapStruct 时间格式
  - 2.5 MapStruct 数字格式
- 3、MapStruct 组合映射
  - 3.1 多参数源映射
  - 3.2 使用其他参数值
  - 3.3 嵌套映射
  - 3.4 逆映射
  - 3.4 继承映射
  - 3.5 共享映射
  - 3.6 自定义方法
    - 3.6.1 自定义类型转换方法
    - 3.6.2 使用@Qualifier
    - 3.6.3  使用@namd
- 4、MapStruct 集合映射、Map映射、枚举映射和Stream 流映射
  - 4.1 集合映射
  - 4.2 Map映射
  - 4.3 枚举映射
  - 4.4 Stream流映射
- 5、MapStruct 其他
  - 5.1 MapStruct 异常处理
  - 5.2 MapStruct 自定义映射
  - 5.3 MapStruct Null值处理
- 6、MapStruct 常见注解总结

### 1、MapStruct 简介
MapStruct是基于JSR269规范的一个***Java注解处理器***，***用于为Java Bean生成类型安全且高性能的映射***。
它基于编译阶段生成get/set代码，此实现过程中没有反射，不会造成额外的性能损失。

#### 1.1 MapStruct Maven引入 
在pom.xml文件添加MapStruct 相关依赖
```xml
<!--mapStruct依赖 高性能对象映射-->
<!--mapstruct核心-->
<dependency>
    <groupId>org.mapstruct</groupId>
    <artifactId>mapstruct</artifactId>
    <version>1.5.0.Beta1</version>
</dependency>
<!--mapstruct编译-->
<dependency>
    <groupId>org.mapstruct</groupId>
    <artifactId>mapstruct-processor</artifactId>
    <version>1.5.0.Beta1</version>
</dependency>
```

### 2、MapStruct 基础操作 

#### 2.1 MapStruct 基本映射
创建MapStruct 映射步骤总结：
- 1.添加MapStruct jar包依赖
- 2.新增接口或抽象类，并且使用org.mapstruct.Mapper注解标签修饰。
- 3.添加自定义转换方法

示例：创建Person 和PersonDto 类定义，通过MapStruct 实现Person 类实例对象转换PersonDto类实例对象。

```java
package com.zzg.mapstruct.entity;
import lombok.Data;
import java.math.BigDecimal;
import java.util.Date;
@Data
public class Person {
    String describe;
    private String id;
    private String name;
    private int age;
    private BigDecimal source;
    private double height;
    private Date createTime;
}
```

```java
package com.zzg.mapstruct.entity;
import lombok.Data;
@Data
public class PersonDTO {
    String describe;
    private Long id;
    private String personName;
    private String age;
    private String source;
    private String height;
}
```

```java
package com.zzg.mapstruct.mapper;
 
import com.zzg.mapstruct.entity.Person;
import com.zzg.mapstruct.entity.PersonDTO;
import org.mapstruct.Mapper;
import org.mapstruct.Mapping;
import org.mapstruct.Mappings;
import org.mapstruct.factory.Mappers;
 
/**
 * mapstruct 工具类定义步骤：
 * 1、添加MapStruct jar包依赖
 * 2、新增接口或抽象类，并且使用org.mapstruct.Mapper注解标签修饰。
 * 3、添加自定义转换方法
 */
@Mapper
public interface PersonMapper {
    PersonMapper INSTANCE = Mappers.getMapper(PersonMapper.class);
 
    PersonDTO convert(Person person);
}
```

测试：
```java
package com.zzg.mapstruct.test;
 
import com.zzg.mapstruct.entity.Person;
import com.zzg.mapstruct.entity.PersonDTO;
import com.zzg.mapstruct.mapper.PersonMapper;
import java.math.BigDecimal;
import java.text.ParseException;
import java.text.SimpleDateFormat;
 
public class OneTest {
    public static void main(String[] args) throws ParseException {
        SimpleDateFormat format = new SimpleDateFormat("yyyy-MM-dd");
 
        Person person = Person.builder().age(31).createTime(format.parse("1991-12-20")).id("1").describe("Java 开发")
                .height(180L).name("在奋斗的大道上").source(new BigDecimal(10000)).build();
 
        PersonDTO personDTO = PersonMapper.INSTANCE.convert(person);
        System.out.println(personDTO);
    }
}
```

执行截图：
```
PersonDTO(describe=Java 开发, id=1, personName=null, age=31, source=10000, height=180.0)
```

优化：Person 类实例对象 转换PersonDto 类实例对象发现 Person.name 属性无法与PersonDTO.personName  
属性无法实现一一对应，可以通过@Mapper注解标签实现不同属性名称转换。
```java
package com.zzg.mapstruct.mapper;
 
import com.zzg.mapstruct.entity.Person;
import com.zzg.mapstruct.entity.PersonDTO;
import org.mapstruct.Mapper;
import org.mapstruct.Mapping;
import org.mapstruct.Mappings;
import org.mapstruct.factory.Mappers;
 
/**
 * mapstruct 工具类定义步骤：
 * 1、添加MapStruct jar包依赖
 * 2、新增接口或抽象类，并且使用org.mapstruct.Mapper注解标签修饰。
 * 3、添加自定义转换方法
 */
@Mapper
public interface PersonMapper {
    PersonMapper INSTANCE = Mappers.getMapper(PersonMapper.class);
 
    @Mapping(source = "name", target="personName")
    PersonDTO convert(Person person);
}
```

```
public static void main(String[] args) throws ParseException {
    SimpleDateFormat format = new SimpleDateFormat("yyyy-MM-dd");

    Person person = Person.builder().age(31).createTime(format.parse("1991-12-20")).id("1").describe("Java 开发")
            .height(180L).name("在奋斗的大道上").source(new BigDecimal(10000)).build();

    PersonDTO personDTO = PersonMapper.INSTANCE.convert(person);
    System.out.println(personDTO);
}
```

执行效果：
```
PersonDTO(describe=Java 开发, id=1, personName=在奋斗的大道上, age=31, source=10000, height=180.0)
```
MapStruct 转换总结：
- 当一个属性与其目标实体对应的名称相同时，它将被隐式映射。
- 当属性在目标实体中具有不同的名称时，可以通过@Mapping注释指定其名称。

温馨提示：
```
如果映射的对象field name不一样，通过 @Mapping 指定。
如果忽略映射字段加@Mapping#ignore() = true 
```

####  2.2 MapStruct 指定默认值
功能要求：Person 类实例对象 转换PersonDTO 类实例对象时，将describe属性值设置为:"默认值"。
在@Mapping注解标签中必须添加***target属性，source属性，默认值使用defaultValue属性设置***。
```java
package com.zzg.mapstruct.mapper;
 
import com.zzg.mapstruct.entity.Person;
import com.zzg.mapstruct.entity.PersonDTO;
import org.mapstruct.Mapper;
import org.mapstruct.Mapping;
import org.mapstruct.Mappings;
import org.mapstruct.factory.Mappers;
 
/**
 * mapstruct 工具类定义步骤：
 * 1、添加MapStruct jar包依赖
 * 2、新增接口或抽象类，并且使用org.mapstruct.Mapper注解标签修饰。
 * 3、添加自定义转换方法
 */
@Mapper
public interface PersonMapper {
    PersonMapper INSTANCE = Mappers.getMapper(PersonMapper.class);
 
    @Mapping(source = "name", target="personName")
    @Mapping(source = "describe", target = "describe", defaultValue = "默认值")
    PersonDTO convert(Person person);
}
```



