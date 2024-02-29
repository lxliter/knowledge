## spring自动注入接口的多个实现类（结合策略设计模式）

在使用spring开发的时候，有时候会出现一个接口多个实现类的情况，但是有没有有时候有这样一种情况，
就是你的逻辑代码里面还不知道你需要使用哪个实现类，就是比如说：你去按摩，按摩店里面有几种会员打折,
比如有，vip、svip和普通用户，那按摩店里面是不是需要对这些会员定义好不同的折扣，然后根据每个用户不同的会员计算出不同的消费情况

虽然这样的情况，对于一般人来说，第一眼肯定就是说，直接加 if else 去判断就可以了
这样做，对于实现功能而言，肯定是没问题，如果以后这个按摩店又增加一种会员，那你是不是又要去修改你的逻辑代码去在加一个 if else ,
***这样就违反了系统架构设计的开闭原则，这样写if else  也使你的代码看起来不优雅***。

所以在代码里面，我们可以先定义一个DiscountStrategy接口类
```java
public interface DiscountStrategy {
    public String getType();
    public double disCount(double fee);
}
```

然后在写他的几个实现类
普通用户实现类
```java
@Service
public class NormalDisCountService implements  DiscountStrategy {
    public String getType(){
        return "normal";
    }
    public double disCount(double fee){
        return fee * 1;
    }
}
```
会员实现类
```java
public class VipDisCountService  implements  DiscountStrategy{
    public String getType(){
        return "vip";
    }
    public double disCount(double fee){
        return fee * 0.8;
    }
}
```
svip超级会员实现类
```java
@Service
public class SVipDisCountService  implements  DiscountStrategy {
    public String getType(){
        return "svip";
    }
    public double disCount(double fee){
        return fee * 0.5;
    }
}
```
然后当一个用户进来消费的时候，根据你当前的身份去打折扣
定义一个map集合，然后把所有的实现类都放入到这个集合中，然后根据当前的会员类型去进行不同的操作
```java
@Service
public class DisCountStrategyService {
    Map<String,DiscountStrategy> discountStrategyMap = new HashMap<>();
    // 构造函数，如果你是集合接口对象，那么久会把spring容器中所有关于该接口的子类，全部抓出来放入到集合中
    @Authwired 
    public DisCountStrageService(List<DiscountStrategy> discountStrategys){
        for (DiscountStrategy discountStrategy: discountStrategys) {
            discountStrategyMap.put(discountStrategy.getType(),discountStrategy);
        }
    }
 
    public double disCount(String type,Double fee){
        DiscountStrategy discountStrategy =discountStrategyMap.get(type);
        return discountStrategy.disCount(fee);
    }
}
```
测试类
```java
@RunWith(SpringRunner.class)
@SpringBootTest
public class MzySpringModeApplicationTests {
    @Autowired
    OrderService orderService;
    @Autowired
    DisCountStrageService disCountStrageService;
    @Test
    public void contextLoads() {
        //orderService.saveOrder();
        double vipresult = disCountStrageService.disCount("vip",100d);
        double svipresult = disCountStrageService.disCount("svip",100d);
        double normalresult = disCountStrageService.disCount("normal",100d);
        System.out.println(vipresult);
        System.out.println(svipresult);
        System.out.println(normalresult);
    }
}
```
其实这就是java设计模式的策略模式，只不过就是用构造函数注入到list集合中
```
使用构造函数的方式注入list集合中
```
就算以后按摩店继续增加了一种会员体系，比如黑卡用户，只需要去写一个接口的实现类，
而不需要去修改以前的代码，这也没有违反系统架构的开闭原则（对修改关闭，对扩展开放）
```
系统架构的设计原则：
对修改关闭 & 对扩展开放
```
注意：上面这种方式只能适用于springboot项目，对于以前的那种springmvc模式不支持，
会报错，因为以前的方式注入是通过配置文件去注入的，需要去修改一个applicationContext的配置文件






