### 阿里大佬：DDD 领域层，该如何设计？

#### 说在前面
在40岁老架构师 尼恩的读者交流群(50+)中，最近有小伙伴拿到了一线互联网企业如阿里、滴滴、极兔、有赞、希音、百度、网易、美团的面试资格，
遇到很多很重要的面试题：
```
谈谈你的DDD落地经验？
谈谈你对DDD的理解？
如何保证RPC代码不会腐烂，升级能力强?
```
最近有小伙伴在字节，又遇到了相关的面试题。小伙伴懵了， 他从来没有用过DDD，面挂了。
关于DDD，尼恩之前给大家梳理过一篇很全的文章： 阿里一面：谈一下你对DDD的理解？2W字，帮你实现DDD自由

但是尼恩的文章， 太过理论化，不适合刚入门的人员。所以，尼恩也在不断的为大家找更好的学习资料。
前段时间，尼恩在阿里的技术公众号上看到了一篇文章《殷浩详解DDD：领域层设计规范》 作者是阿里 技术大佬殷浩，
非常适合于初学者入门，同时也足够的有深度。

美中不足的是， 殷浩那篇文章的行文风格，对初学者不太友好， 尼恩刚开始看的时候，也比较晦涩。
于是，尼恩在读的过程中，把那些晦涩的内容，给大家用尼恩的语言， 浅化了一下， 这样大家更容易懂。

本着技术学习、技术交流的目的，这里，把尼恩修改过的 《殷浩详解DDD：领域层设计规范》，通过尼恩的公众号《技术自由圈》发布出来。
```
特别声明，由于没有殷浩同学的联系方式，这里没有找殷浩的授权，
如果殷浩同学或者阿里技术公众号不同意我的修改，不同意我的发布， 我即刻从《技术自由圈》公众号扯下来。
```

另外， 文章也特别长， 我也特别准备了PDF版本。
如果需要尼恩修改过的PDF版本，也可以通过《技术自由圈》公众号找到尼恩来获取。

#### 本文目录
- 说在前面
- 本文说明
- 尼恩总结：DDD的本质和最终收益
- 第4篇 - 领域层设计规范
  - 1. 前置知识：领域模型 Model定义、分类、生命周期
    - 1.1 Model分类
    - 1.2 Model的生命周期
  - 2. 业务场景：龙与魔法的世界      
    - 2.1 背景和规则             
    - 2.2 OOP实现         
    - 2.3 OOP代码的设计缺陷
      - 缺陷一：编程语言的强类型无法承载业务规则
      - 缺陷二：对象继承导致代码强依赖父类逻辑，违反开闭原则Open-Closed Principle（OCP）           
      - 缺陷三：多对象行为类似，导致代码重复
  - 3. Entity-Component-System（ECS）架构简介    
    - 3.1 ECS介绍        
    - 3.2 ECS架构分析      
    - 3.3 ECS的缺陷      
  - 4. 基于DDD架构的龙与魔法的世界      
    - 4.1 领域对象
      - 实体类            
      - 值对象的组件化        
    - 4.2 装备行为
    - 4.3 攻击行为
    - 4.4 单元测试
    - 4.5 移动系统
  - 5. DDD领域层的一些设计规范    
    - 5.1 实体类（Entity）         
      - 原则1：创建即一致
      - 原则2：尽量避免public setter
      - 原则3：通过聚合根保证父子实体的一致性
      - 原则4：不可以强依赖其他聚合根实体或领域服务
      - 原则5：任何实体的行为只能直接影响到本实体（和其子实体）
    - 5.2 领域服务（Domain Service）
      - 类型1：单对象策略型
      - 类型2：跨对象事务型
      - 类型3：通用组件型
    - 5.3 策略对象（Domain Policy）
  - 6. 使用事件处理领域模型副作用     
    - 6.1 领域事件介绍
    - 6.2 领域事件实现
    - 6.3 目前领域事件的缺陷和展望
    - 6.4 领域事件总结
  - 7. 领域层设计规范总结     
    - 7.1 Entity设计规范总结
    - 7.2 Domain Service 设计规范总结
    - 7.3 Application Service 设计规范总结
    - 7.4 其它规范总结
- 未完待续，尼恩说在最后
- 部分历史案例

#### 本文说明
本文是 《阿里DDD大佬：从0到1，带大家精通DDD》系列的第二篇
本文是 《从0到1，带大家精通DDD》系列的第3篇， 第1、2、3篇的链接地址是：
《阿里DDD大佬：从0到1，带大家精通DDD》
《阿里大佬：DDD 落地两大步骤，以及Repository核心模式》
《阿里面试：让代码不腐烂，DDD是怎么做的？》

另外，尼恩会结合一个工业级的DDD实操项目，
在第34章视频《DDD的顶奢面经》中，给大家彻底介绍一下DDD的实操、COLA 框架、DDD的面试题。

#### 尼恩总结：DDD的本质和最终收益
在正式开始第4篇之前，尼恩说一下自己对DDD的 亲身体验、和深入思考。
DDD的本质：
- 大大提升 核心代码  业务纯度

老的mvc架构，代码中紧紧的耦合着特定ORM框架、特定DB存储、特定的缓存、特定的事务框架、特定中间件，特定对外依赖解耦， 很多很多。
总之就是 业务和技术紧密耦合，代码的 业务纯度低， 导致软件“固化”， 没法做快速扩展和升级。

- 大大提升 代码工程  测维扩 能力

DDD进行了多个层次的解耦，包括 持久层的DB解耦，外依第三方赖的隔离解耦，大大提升了  可测试度、可维护度、可扩展度

- 更大限度  积累 业务领域模型 资产

由于spring mvc 模式下， 代码的业务纯度不高， 导致尼恩的曾经一个项目，10年多时间， 
衍生出  50多个不同的版本，推导重来5次，付出巨大的 时间成本、经济成本

##### DDD的收益：
- 极大的降低升级的工作量
- 极大的降低推到重来的风险
- 极大提升代码的核心代码业务纯度，积累更多的代码资产

不用DDD的反面案例，尼恩曾经见过一个项目：
- 10年多时间， 衍生出  50多个不同的版本， 每个版本80%的功能相同，但是代码各种冲突，没有合并
- 10年多时间，经历过至少 5次推倒重来， 基本换一个领导，恨不得推导重来一次， 感觉老的版本都是不行，只有自己设计的才好
- 5次推倒重来，每次都是 风风火火/加班到进ICU， 投入了大量的人力/财力。其实大多是重复投入、重复建设
- 可谓,  一将不才累死三军， 所以， 一个优秀的架构师，对一个项目来说是多么的重要

#### 第4篇 - 领域层设计规范
在一个DDD架构设计中，领域层的设计 非常重要，
领域层 不仅仅 会直接影响整个架构的代码结构，而且会影响  应用层、基础设施层的设计。
领域层设计又是有挑战的任务，
特别是在一个业务逻辑相对复杂应用中，每一个业务规则是应该放在Entity、ValueObject 还是 DomainService是值得用心思考的，
既要避免未来的扩展性差，又要避免过度设计导致复杂性。
```
特别是在一个业务逻辑相对复杂应用中，每一个业务规则是应该放在Entity、ValueObject还是DomainService是值得用心思考的，
既要避免未来的扩展性差，又要避免过度设计导致复杂性
```
今天我用一个相对轻松易懂的领域做一个案例演示，但在实际业务应用中，无论是交易、营销还是互动，都可以用类似的逻辑来实现。

##### 1. 前置知识：领域模型 Model定义、分类、生命周期
Model（模型）：承载着***业务的属性和具体的行为***，是业务表达的方式，是DDD的内核。
- Model（模型）是一个类中有属性、属性有Get/Set方法，
- 并且业务的行为（Action）操作也是在模型类中（充血模型）
- 模型分为***Entity***、***Value Object***、***Service***这三种类型

###### 1.1 Model分类
- Entity (实体)
  - 有特定的标识，标识着这个Model在系统中全局唯一
  - 内部值可以是变化的，可能存在生命周期 (***比如订单对象，状态值是连续变化的***)
  - ***有状态的Value Object***
- Value Object （值对象）
  - 内部值是不变的，不存在生命周期 (***比如地址对象不存在生命周期***)
  - ***无状态对象***
- Service （服务）
  - 无状态对象
  - ***当一个属性或行为放在Entity、Value Object中模棱两可或不合适的时候就需要以Service的形式来呈现***

三种模型复杂度：Service > Entity > ValueObject，优先选择简单模型

###### 1.2 Model的生命周期
- Factory （工厂）：用来创建Model，以及帮助Repository (数据源)注入到Model中
- Aggregate （聚合根）：封装Model，一个Mode中可能包含其他Model（类似一个对象中包含其他对象的引用，实际概念更复杂）
  - ***聚合是用来封装真正的不变性，而不是简单的将对象组合在一起***；
  - 聚合应尽量设计的小；
  - ***聚合之间的关联通过ID，而不是对象引用***；
  - 聚合内强一致性，聚合之间最终一致性。
- Repository (数据源)：
  - 数据源的访问网关层
  - Model通过Repository来对接不同的数据源
    
了解了领域模型 Model定义、分类、生命周期之后，咱们来进行  DDD的 领域层设计规范介绍。

##### 2. 业务场景：龙与魔法的世界
这里 找一个轻松的案例 ，一个龙与魔法的游戏世界的（极简）规则，看看 如何用代码实现

###### 2.1 背景和规则
基础配置如下：
- 玩家（Player）
战士（Fighter）、法师（Mage）、龙骑（Dragoon）

- 怪物（Monster）
兽人（Orc）、精灵（Elf）、龙（Dragon），怪物有血量

- 武器（Weapon）
是剑（Sword）、法杖（Staff），武器有攻击力

玩家可以装备一个武器，武器有攻击的类型， 攻击类型可以是物理（0），魔法（1），冰（2）等，攻击类型决定伤害类型。
攻击规则如下：
- 兽人对物理攻击伤害减半
- 精灵对魔法攻击伤害减半
- 龙对物理和魔法攻击免疫，除非玩家是龙骑，则伤害加倍

###### 2.2 OOP实现
通过 Object-Oriented Programming 面向对象的方式进行设计，
一通过类的继承关系（此处省略部分非核心代码） 实现上面的case。

玩家的接口和类如下：
```java
public abstract class Player {
      Weapon weapon;
}
public class Fighter extends Player {}
public class Mage extends Player {}
public class Dragoon extends Player {}
```

怪物的接口和类如下：
```java
public abstract class Monster {
    Long health;
}
public class Orc extends Monster {}
public class Elf extends Monster {}
public class Dragoon extends Monster {}
```

武器的接口和类如下：
```java
public abstract class Weapon {
    int damage;
    int damageType; // 0 - physical, 1 - fire, 2 - ice etc.
}
public class Sword extends Weapon {}
public class Staff extends Weapon {}
```

而实现规则代码如下：
```java
public class Player {
    public void attack(Monster monster) {
        monster.receiveDamageBy(weapon, this);
    }
}
public class Monster {
  public void receiveDamageBy(Weapon weapon, Player player) {
    this.health -= weapon.getDamage(); // 基础规则
  }
}
//兽人（Orc） 对物理攻击伤害减半

public class Orc extends Monster {
  @Override
  public void receiveDamageBy(Weapon weapon, Player player) {
    if (weapon.getDamageType() == 0) {
      this.setHealth(this.getHealth() - weapon.getDamage() / 2); // Orc的物理防御规则
    } else {
      super.receiveDamageBy(weapon, player);
    }
  }
}

//龙（Dragon）对物理和魔法攻击免疫，除非玩家是龙骑，则伤害加倍
public class Dragon extends Monster {
  @Override
  public void receiveDamageBy(Weapon weapon, Player player) {
    if (player instanceof Dragoon) {
      this.setHealth(this.getHealth() - weapon.getDamage() * 2); // 龙骑伤害规则
    }
    // else no damage, 龙免疫力规则
  }
}
```

然后跑几个单测：
```java
public class BattleTest {
    @Test
    @DisplayName("Dragon is immune to attacks")
    public void testDragonImmunity() {
      // 玩家
      Fighter fighter = new Fighter("Hero");
      //武器， 剑（Sword）
      Sword sword = new Sword("Excalibur", 10);
      //给玩家设计武器
      fighter.setWeapon(sword);
      Dragon dragon = new Dragon("Dragon", 100L);
      // When
      fighter.attack(dragon);
      // Then
      assertThat(dragon.getHealth()).isEqualTo(100);
    }

    @Test
    @DisplayName("Dragoon attack dragon doubles damage")
    public void testDragoonSpecial() {
      // Given
      Dragoon dragoon = new Dragoon("Dragoon");
      Sword sword = new Sword("Excalibur", 10);
      dragoon.setWeapon(sword);
      Dragon dragon = new Dragon("Dragon", 100L);
      // When
      dragoon.attack(dragon);
      // Then
      assertThat(dragon.getHealth()).isEqualTo(100 - 10 * 2);
    }

    @Test
    @DisplayName("Orc should receive half damage from physical weapons")
    public void testFighterOrc() {
      // Given
      Fighter fighter = new Fighter("Hero");
      Sword sword = new Sword("Excalibur", 10);
      fighter.setWeapon(sword);
      Orc orc = new Orc("Orc", 100L);
      // When
      fighter.attack(orc);
      // Then
      assertThat(orc.getHealth()).isEqualTo(100 - 10 / 2);
    }

    @Test
    @DisplayName("Orc receive full damage from magic attacks")
    public void testMageOrc() {
      // Given
      Mage mage = new Mage("Mage");
      Staff staff = new Staff("Fire Staff", 10);
      mage.setWeapon(staff);
      Orc orc = new Orc("Orc", 100L);
      // When
      mage.attack(orc);
      // Then
      assertThat(orc.getHealth()).isEqualTo(100 - 10);
    }
}
```
以上代码和单测都比较简单，不做多余的解释了。

###### 2.3 OOP代码的设计缺陷
缺陷一：编程语言的强类型无法承载业务规则
以上的OOP代码可以跑得通，但是，现在要进行扩展，我们加一个限制条件：
- 战士只能装备剑
- 法师只能装备法杖

这个规则在Java语言里无法通过强类型来实现，
虽然Java有Variable Hiding（或者C#的new class variable），但实际上只是在子类上加了一个新变量，所以会导致以下的问题：
```
public abstract class Player {
      Weapon weapon;
}

//  战士（Fighter）  只能装备剑 
@Data
public class Fighter extends Player {
    private Sword weapon;
}

@Test
public void testEquip() {
    Fighter fighter = new Fighter("Hero");
   // 剑（Sword）
    Sword sword = new Sword("Sword", 10);
    fighter.setWeapon(sword);

    //  法杖（Staff）
    Staff staff = new Staff("Staff", 10);
    fighter.setWeapon(staff);

    assertThat(fighter.getWeapon()).isInstanceOf(Staff.class); // 错误了
}
```

在最后，虽然代码感觉是setWeapon(Staff)，但实际上只修改了父类的变量，
并没有修改子类的变量，所以实际不生效，也不抛异常，但结果是错的。

当然，可以在父类限制setter为protected：
```
@Data
public abstract class Player {
    @Setter(AccessLevel.PROTECTED)
    private Weapon weapon;
}

@Test
public void testCastEquip() {
    Fighter fighter = new Fighter("Hero");

    Sword sword = new Sword("Sword", 10);
    fighter.setWeapon(sword);

    Player player = fighter;
    Staff staff = new Staff("Staff", 10);
    player.setWeapon(staff); // 编译不过，但从API层面上应该开放可用
}
```

在父类限制setter为protected， 有一个很大的坏处：限制了父类的API，极大的降低了灵活性，
同时也违背了Liskov substitution principle，即一个父类必须要cast成子类才能使用

现在还要进行扩展，我们又加一个限制条件：
- 战士和法师都能装备匕首（dagger）

这下问题更大，之前写的强类型代码都废了，需要重构。

###### 缺陷二：对象继承导致代码强依赖父类逻辑，违反开闭原则Open-Closed Principle（OCP）
开闭原则（OCP）规定“对象应该对于扩展开放，对于修改封闭“，
继承虽然可以通过子类扩展新的行为，但因为子类可能直接依赖父类的实现，导致一个变更可能会影响所有对象。

在这个例子里，如果增加任意一种类型的玩家、怪物或武器，或增加一种规则，都有可能需要修改从父类到子类的所有方法。
比如，如果要增加一个武器类型：狙击枪，能够无视所有防御一击必杀，需要修改的代码包括：
- Weapon
- Player和所有的子类（是否能装备某个武器的判断）
- Monster和所有的子类（伤害计算逻辑）

```java
public class Monster {
    public void receiveDamageBy(Weapon weapon, Player player) {
     
        this.health -= weapon.getDamage(); // 老的基础规则
      
        if (Weapon instanceof Gun) { // 新的逻辑
            this.setHealth(0);
        }
    }
}
public class Dragon extends Monster {
    public void receiveDamageBy(Weapon weapon, Player player) {
        if (Weapon instanceof Gun) { // 新的逻辑
            super.receiveDamageBy(weapon, player);
        }
        // 老的逻辑省略
    }
}
```

在一个复杂的软件中为什么会建议“尽量”不要违背OCP？
最核心的原因就是：***一个现有逻辑的变更，可能会影响一些原有的代码，导致一些无法预见的影响***。
这个风险只能通过完整的单元测试覆盖来保障，但在实际开发中很难保障单测的覆盖率。
OCP的原则能尽可能的规避这种风险，当新的行为只能通过新的字段/方法来实现时，老代码的行为自然不会变。

继承虽然能Open for extension，但很难做到Closed for modification。
```
继承虽然能open for extension，但很难做到Closed for modification
```
所以今天解决OCP的主要方法是通过***Composition-over-inheritance***，即通过组合来做到扩展性，而不是通过继承。
这就是常常说的 组合优于继承。

但是，尽管如此，在这个例子里，其实业务规则的逻辑到底应该写在哪里，其实是有异议的：
当我们去看一个对象和另一个对象之间的交互时，到底是Player去攻击Monster，还是Monster被Player攻击？
```
Player.attack(monster) 还是 Monster.receiveDamage(Weapon, Player)？
```
目前的代码主要将逻辑写在Monster的类中，主要考虑是Monster会受伤降低Health，
但如果是Player拿着一把双刃剑会同时伤害自己呢？是不是发现写在Monster类里也有问题？
代码写在哪里的原则是什么？

###### 缺陷三：多对象行为类似，导致代码重复
当我们有不同的对象，但又有相同或类似的行为时，OOP会不可避免的导致代码的重复。
在这个例子里，如果我们去增加一个“可移动”的行为，需要在Player和Monster类中都增加类似的逻辑：
```java
public abstract class Player {
    int x;
    int y;
    void move(int targetX, int targetY) {
        // logic
    }
}

public abstract class Monster {
    int x;
    int y;
    void move(int targetX, int targetY) {
        // logic
    }
}
```
一个可能的解法是有个通用的父类：
```java
public abstract class Movable {
    int x;
    int y;
    void move(int targetX, int targetY) {
        // logic
    }
}

public abstract class Player extends Movable{};
public abstract class Monster extends Movable{};
```
但如果再增加一个跳跃能力Jumpable呢？
但如果再增加一个跑步能力Runnable呢？
如果Player可以Move和Jump，Monster可以Move和Run，怎么处理继承关系？
要知道Java（以及绝大部分语言）是不支持多父类继承的，所以只能通过重复代码来实现。

> 问题总结

在这个案例里虽然从直觉来看OOP的逻辑很简单，但如果你的业务比较复杂，未来会有大量的业务规则变更时，
简单的OOP代码会在后期变成复杂的一团浆糊，逻辑分散在各地，缺少全局视角，各种规则的叠加会触发bug。
有没有感觉似曾相识？
对的，***电商体系里的优惠、交易等链路经常会碰到类似的坑***。
而这类问题的核心本质在于：
- 业务规则的归属到底是对象的“行为”还是独立的”规则对象“？
- 业务规则之间的关系如何处理？
- 通用“行为”应该如何复用和维护？

在讲DDD的解法前，我们先去看看一套游戏里最近比较火的架构设计，
Entity-Component-System（ECS）是如何实现的。

##### 3. Entity-Component-System（ECS）架构简介
###### 3.1 ECS介绍
ECS架构模式是其实是一个很老的游戏架构设计，最早应该能追溯到《地牢围攻》的组件化设计，
但最近因为Unity的加入而开始变得流行（比如《守望先锋》就是用的ECS）。

要很快的理解ECS架构的价值，我们需要理解一个游戏代码的核心问题：
- 性能：游戏必须要实现一个高的渲染率（60FPS），也就是说整个游戏世界需要在1/60s（大概16ms）内完整更新一次（包括物理引擎、游戏状态、渲染、AI等）。
在一个游戏中，通常有大量的（万级、十万级）游戏对象需要更新状态，除了渲染可以依赖GPU之外，其他的逻辑都需要由CPU完成，甚至绝大部分只能由单线程完成，
导致绝大部分时间复杂场景下CPU（主要是内存到CPU的带宽）会成为瓶颈。
在CPU单核速度几乎不再增加的时代，如何能让CPU处理的效率提升，是提升游戏性能的核心。

- 代码组织：当我们用传统OOP的模式进行游戏开发时，很容易就会陷入代码组织上的问题，最终导致代码难以阅读，维护和优化。

- 可扩展性：这个跟上一条类似，但更多的是游戏的特性导致：需要快速更新，加入新的元素。
一个游戏的架构需要能通过低代码、甚至0代码的方式增加游戏元素，从而通过快速更新而留住用户。
如果每次变更都需要开发新的代码，测试，然后让用户重新下载客户端，可想而知这种游戏很难在现在的竞争环境下活下来。

而ECS架构能很好的解决上面的几个问题，ECS架构主要分为：
- Entity：用来代表任何一个游戏对象，但是在ECS里一个Entity最重要的仅仅是他的EntityID，一个Entity里包含多个Component
- Component：是真正的数据，ECS架构把一个个的实体对象拆分为更加细化的组件，比如位置、素材、状态等，也就是说一个Entity实际上只是一个Bag of Components。
- System（或者ComponentSystem，组件系统）：是真正的行为，一个游戏里可以有很多个不同的组件系统，每个组件系统都只负责一件事，
可以依次处理大量的相同组件，而不需要去理解具体的Entity。

所以一个ComponentSystem理论上可以有更加高效的组件处理效率，甚至可以实现并行处理，从而提升CPU利用率。

ECS的一些核心性能优化包括将同类型组件放在同一个Array中，然后Entity仅保留到各自组件的pointer，
这样能更好的利用CPU的缓存，减少数据的加载成本，以及SIMD的优化等。

```
public class Entity {
  public Vector position; // 此处Vector是一个Component, 指向的是MovementSystem.list里的一个
}

public class MovementSystem {
  List< Vector> list;

  // System的行为
  public void update(float delta) {
    for(Vector pos : list) { // 这个loop直接走了CPU缓存，性能很高，同时可以用SIMD优化
      pos.x = pos.x + delta;
      pos.y = pos.y + delta;
    }
  }
}

@Test
public void test() {
  MovementSystem system = new MovementSystem();
  system.list = new List<>() { new Vector(0, 0) };
  Entity entity = new Entity(list.get(0));
  system.update(0.1);
  assertTrue(entity.position.x == 0.1);
}
```
由于本文不是讲解ECS架构的，感兴趣的同学可以搜索Entity-Component-System或者看看Unity的ECS文档等。

###### 3.2 ECS架构分析
重新回来分析ECS，其实它的本源还是几个很老的概念：
- 组件化
在软件系统里，我们通常将复杂的大系统拆分为独立的组件，来降低复杂度。
比如网页里通过前端组件化降低重复开发成本，微服务架构通过服务和数据库的拆分降低服务复杂度和系统影响面等。
但是ECS架构把这个走到了极致，即每个对象内部都实现了组件化。
通过将一个游戏对象的数据和行为拆分为多个组件和组件系统，能实现组件的高度复用性，降低重复开发成本。

- 行为抽离
这个在游戏系统里有个比较明显的优势。
如果按照OOP的方式，一个游戏对象里可能会包括移动代码、战斗代码、渲染代码、AI代码等，如果都放在一个类里会很长，且很难去维护。
通过将通用逻辑抽离出来为单独的System类，可以明显提升代码的可读性。
另一个好处则是抽离了一些和对象代码无关的依赖，比如上文的delta，这个delta如果是放在Entity的update方法，则需要作为入参注入，而放在System里则可以统一管理。
在介绍组合优于继承的时候有个遗留问题，到底是应该Player.attack(monster) 还是 Monster.receiveDamage(Weapon, Player)？
这个问题在ECS里这个问题就变的很简单，放在CombatSystem里就可以了。

- 数据驱动
即一个对象的行为不是写死的而是通过其参数决定，通过参数的动态修改，就可以快速改变一个对象的具体行为。
在ECS的游戏架构里，通过给Entity注册相应的Component，以及改变Component的具体参数的组合，就可以改变一个对象的行为和玩法，
比如创建一个水壶+爆炸属性就变成了“爆炸水壶”、给一个自行车加上风魔法就变成了飞车等。
在有些Rougelike游戏中，可能有超过1万件不同类型、不同功能的物品，如果这些不同功能的物品都去单独写代码，
可能永远都写不完，但是通过数据驱动+组件化架构，所有物品的配置最终就是一张表，修改也极其简单。
这个也是组合胜于继承原则的一次体现。

###### 3.3 ECS的缺陷
虽然ECS在游戏界已经开始崭露头角，我发现ECS架构目前还没有在哪个大型商业应用中被使用过。
原因可能很多，包括ECS比较新大家还不了解、缺少商业成熟可用的框架、程序员们还不够能适应从写逻辑脚本到写组件的思维转变等，
但我认为其最大的一个问题是ECS为了提升性能，强调了数据/状态（State）和行为（Behaviour）分离，并且为了降低GC成本，直接操作数据，走到了一个极端。
而在商业应用中，数据的正确性、一致性和健壮性应该是最高的优先级，而性能只是锦上添花的东西，所以ECS很难在商业场景里带来特别大的好处。
但这不代表我们不能借鉴一些ECS的突破性思维，包括组件化、跨对象行为的抽离、以及数据驱动模式，而这些在DDD里也能很好的用起来。

#### 4. 基于DDD架构的龙与魔法的世界
##### 4.1 领域对象
回到我们原来的问题域上面，我们从***领域层拆分一下各种对象***：

- 实体类

在DDD里，***实体类包含ID和内部状态***，在这个案例里实体类包含Player、Monster和Weapon。
Weapon之所以被设计成实体类，是因为两把同名的Weapon应该可以同时存在，所以必须要有ID来区分，
同时未来也可以预期Weapon会包含一些状态，比如升级、临时的buff、耐久等。
```java
public class Player implements Movable {
    private PlayerId id;
    private String name;
    private PlayerClass playerClass; // enum
    private WeaponId weaponId; // （Note 1）
    private Transform position = Transform.ORIGIN;
    private Vector velocity = Vector.ZERO;
}

public class Monster implements Movable {
    private MonsterId id;
    private MonsterClass monsterClass; // enum
    private Health health;
    private Transform position = Transform.ORIGIN;
    private Vector velocity = Vector.ZERO;
}

public class Weapon {
    private WeaponId id;
    private String name;
    private WeaponType weaponType; // enum
    private int damage;
    private int damageType; // 0 - physical, 1 - fire, 2 - ice
}
```

在这个简单的案例里，我们可以利用enum的PlayerClass、MonsterClass来代替继承关系，
后续也可以利用Type Object 设计模式来做到数据驱动。

```
Note 1: 因为 Weapon 是实体类，但是Weapon能独立存在，
Player不是聚合根，所以Player只能保存WeaponId，而不能直接指向Weapon。
```

- 值对象的组件化

在前面的ECS架构里，有个MovementSystem的概念是可以复用的，
虽然不应该直接去操作Component或者继承通用的父类，
但是可以通过接口的方式对领域对象做组件化处理：

```java
public interface Movable {
    // 相当于组件
    Transform getPosition();
    Vector getVelocity();

    // 行为
    void moveTo(long x, long y);
    void startMove(long velX, long velY);
    void stopMove();
    boolean isMoving();
}

// 具体实现
public class Player implements Movable {
    public void moveTo(long x, long y) {
        this.position = new Transform(x, y);
    }

    public void startMove(long velocityX, long velocityY) {
        this.velocity = new Vector(velocityX, velocityY);
    }

    public void stopMove() {
        this.velocity = Vector.ZERO;
    }

    @Override
    public boolean isMoving() {
        return this.velocity.getX() != 0 || this.velocity.getY() != 0;
    }
}

@Value
public class Transform {
    public static final Transform ORIGIN = new Transform(0, 0);
    long x;
    long y;
}

@Value
public class Vector {
    public static final Vector ZERO = new Vector(0, 0);
    long x;
    long y;
}
```
注意两点：
- Movable的接口没有Setter。
- 单个Entity的规则是不能直接变更其属性，必须通过Entity的方法去对内部状态做变更。这样能保证数据的一致性。
- 抽象Movable的好处是如同ECS一样，一些特别通用的行为（如在大地图里移动）可以通过统一的System代码去处理，避免了重复劳动。

###### 4.2 装备行为
因为我们已经不会用Player的子类来决定什么样的Weapon可以装备，所以这段逻辑应该被拆分到一个单独的类里。
这种类在DDD里被叫做领域服务（Domain Service）。
```java
public interface EquipmentService {
    boolean canEquip(Player player, Weapon weapon);
}
```
在DDD里，一个Entity不应该直接依赖另一个Entity或服务，也就是说以下的代码是错误的：
```java
public class Player {
    @Autowired
    EquipmentService equipmentService; // BAD: 不可以直接依赖

    public void equip(Weapon weapon) {
       // ...
    }
}
```
这里的问题是Entity只能保留自己的状态（或非聚合根的对象）。
任何其他的对象，无论是否通过依赖注入的方式弄进来，都会破坏Entity的Invariance，并且还难以单测。
正确的引用方式是通过方法参数引入（Double Dispatch）：
```java
public class Player {
    public void equip(Weapon weapon, EquipmentService equipmentService) {
        if (equipmentService.canEquip(this, weapon)) {
            this.weaponId = weapon.getId();
        } else {
            throw new IllegalArgumentException("Cannot Equip: " + weapon);
        }
    }
}
```
在这里，无论是Weapon还是EquipmentService都是通过方法参数传入，确保不会污染Player的自有状态。
Double Dispatch是一个使用Domain Service经常会用到的方法，类似于调用反转。
```
Double Dispatch是一个使用Domain Service经常会用到的方法，类似于调用反转
```
然后在EquipmentService里实现相关的逻辑判断，这里我们用了另一个常用的Strategy（或者叫Policy）设计模式：
```java
public class EquipmentServiceImpl implements EquipmentService {
    private EquipmentManager equipmentManager; 

    @Override
    public boolean canEquip(Player player, Weapon weapon) {
        return equipmentManager.canEquip(player, weapon);
    }
}

// 策略优先级管理
public class EquipmentManager {
    private static final List< EquipmentPolicy> POLICIES = new ArrayList<>();
    static {
        POLICIES.add(new FighterEquipmentPolicy());
        POLICIES.add(new MageEquipmentPolicy());
        POLICIES.add(new DragoonEquipmentPolicy());
        POLICIES.add(new DefaultEquipmentPolicy());
    }

    public boolean canEquip(Player player, Weapon weapon) {
        for (EquipmentPolicy policy : POLICIES) {
            if (!policy.canApply(player, weapon)) {
                continue;
            }
            return policy.canEquip(player, weapon);
        }
        return false;
    }
}

// 策略案例
public class FighterEquipmentPolicy implements EquipmentPolicy {

    @Override
    public boolean canApply(Player player, Weapon weapon) {
        return player.getPlayerClass() == PlayerClass.Fighter;
    }

    /**
     * Fighter能装备Sword和Dagger
     */
    @Override
    public boolean canEquip(Player player, Weapon weapon) {
        return weapon.getWeaponType() == WeaponType.Sword
                || weapon.getWeaponType() == WeaponType.Dagger;
    }
}

// 其他策略省略，见源码
```
***这样设计的最大好处是未来的规则增加只需要添加新的Policy类，而不需要去改变原有的类***。

###### 4.3 攻击行为
在上文中曾经有提起过，到底应该是Player.attack(Monster)还是Monster.receiveDamage(Weapon, Player)？
在DDD里，因为这个行为可能会影响到Player、Monster和Weapon，所以属于跨实体的业务逻辑。
在这种情况下需要通过一个第三方的领域服务（Domain Service）来完成。

```java
public interface CombatService {
    void performAttack(Player player, Monster monster);
}

public class CombatServiceImpl implements CombatService {
    private WeaponRepository weaponRepository;
    private DamageManager damageManager;

    @Override
    public void performAttack(Player player, Monster monster) {
        Weapon weapon = weaponRepository.find(player.getWeaponId());
        int damage = damageManager.calculateDamage(player, weapon, monster);
        if (damage > 0) {
            monster.takeDamage(damage); // （Note 1）在领域服务里变更Monster
        }
        // 省略掉Player和Weapon可能受到的影响
    }
}
```

同样的在这个案例里，可以通过Strategy设计模式来解决damage的计算问题：
```java
// 策略优先级管理
public class DamageManager {
    private static final List< DamagePolicy> POLICIES = new ArrayList<>();
    static {
        POLICIES.add(new DragoonPolicy());
        POLICIES.add(new DragonImmunityPolicy());
        POLICIES.add(new OrcResistancePolicy());
        POLICIES.add(new ElfResistancePolicy());
        POLICIES.add(new PhysicalDamagePolicy());
        POLICIES.add(new DefaultDamagePolicy());
    }

    public int calculateDamage(Player player, Weapon weapon, Monster monster) {
        for (DamagePolicy policy : POLICIES) {
            if (!policy.canApply(player, weapon, monster)) {
                continue;
            }
            return policy.calculateDamage(player, weapon, monster);
        }
        return 0;
    }
}

// 策略案例
public class DragoonPolicy implements DamagePolicy {
    public int calculateDamage(Player player, Weapon weapon, Monster monster) {
        return weapon.getDamage() * 2;
    }
    @Override
    public boolean canApply(Player player, Weapon weapon, Monster monster) {
        return player.getPlayerClass() == PlayerClass.Dragoon &&
                monster.getMonsterClass() == MonsterClass.Dragon;
    }
}
```
特别需要注意的是这里的CombatService领域服务和3.2的EquipmentService领域服务，虽然都是领域服务，但实质上有很大的差异。
上文的EquipmentService更多的是提供只读策略，且只会影响单个对象，所以可以在Player.equip方法上通过参数注入。
但是CombatService有可能会影响多个对象，所以不能直接通过参数注入的方式调用。

###### 4.4 单元测试
```
@Test
@DisplayName("Dragoon attack dragon doubles damage")
public void testDragoonSpecial() {
    // Given
    Player dragoon = playerFactory.createPlayer(PlayerClass.Dragoon, "Dart");
    Weapon sword = weaponFactory.createWeaponFromPrototype(swordProto, "Soul Eater", 60);
    ((WeaponRepositoryMock)weaponRepository).cache(sword);
    dragoon.equip(sword, equipmentService);
    Monster dragon = monsterFactory.createMonster(MonsterClass.Dragon, 100);

    // When
    combatService.performAttack(dragoon, dragon);

    // Then
    assertThat(dragon.getHealth()).isEqualTo(Health.ZERO);
    assertThat(dragon.isAlive()).isFalse();
}

@Test
@DisplayName("Orc should receive half damage from physical weapons")
public void testFighterOrc() {
    // Given
    Player fighter = playerFactory.createPlayer(PlayerClass.Fighter, "MyFighter");
    Weapon sword = weaponFactory.createWeaponFromPrototype(swordProto, "My Sword");
    ((WeaponRepositoryMock)weaponRepository).cache(sword);
    fighter.equip(sword, equipmentService);
    Monster orc = monsterFactory.createMonster(MonsterClass.Orc, 100);

    // When
    combatService.performAttack(fighter, orc);

    // Then
    assertThat(orc.getHealth()).isEqualTo(Health.of(100 - 10 / 2));
}
```
具体的代码比较简单，解释省略。

###### 4.5 移动系统
最后还有一种Domain Service，通过组件化，我们其实可以实现ECS一样的System，来降低一些重复性的代码：
```java
public class MovementSystem {

    private static final long X_FENCE_MIN = -100;
    private static final long X_FENCE_MAX = 100;
    private static final long Y_FENCE_MIN = -100;
    private static final long Y_FENCE_MAX = 100;

    private List< Movable> entities = new ArrayList<>();

    public void register(Movable movable) {
        entities.add(movable);
    }

    public void update() {
        for (Movable entity : entities) {
            if (!entity.isMoving()) {
                continue;
            }

            Transform old = entity.getPosition();
            Vector vel = entity.getVelocity();
            long newX = Math.max(Math.min(old.getX() + vel.getX(), X_FENCE_MAX), X_FENCE_MIN);
            long newY = Math.max(Math.min(old.getY() + vel.getY(), Y_FENCE_MAX), Y_FENCE_MIN);
            entity.moveTo(newX, newY);
        }
    }
}
```
单测：
```
@Test
@DisplayName("Moving player and monster at the same time")
public void testMovement() {
    // Given
    Player fighter = playerFactory.createPlayer(PlayerClass.Fighter, "MyFighter");
    fighter.moveTo(2, 5);
    fighter.startMove(1, 0);

    Monster orc = monsterFactory.createMonster(MonsterClass.Orc, 100);
    orc.moveTo(10, 5);
    orc.startMove(-1, 0);

    movementSystem.register(fighter);
    movementSystem.register(orc);

    // When
    movementSystem.update();

    // Then
    assertThat(fighter.getPosition().getX()).isEqualTo(2 + 1);
    assertThat(orc.getPosition().getX()).isEqualTo(10 - 1);
}
```
在这里MovementSystem就是一个相对独立的Domain Service，
通过对Movable的组件化，实现了类似代码的集中化、以及一些通用依赖/配置的中心化（如X、Y边界等）

#### 5. DDD领域层的一些设计规范
上面我主要针对同一个例子对比了OOP、ECS和DDD的3种实现，比较如下：
- 基于继承关系的OOP代码：
OOP的代码最好写，也最容易理解，所有的规则代码都写在对象里，但是当领域规则变得越来越复杂时，其结构会限制它的发展。
新的规则有可能会导致代码的整体重构。

- 基于组件化的ECS代码：
ECS代码有最高的灵活性、可复用性、及性能，但极具弱化了实体类的内聚，所有的业务逻辑都写在了服务里，
会导致业务的一致性无法保障，对商业系统会有较大的影响。

- 基于领域对象 + 领域服务的DDD架构：
DDD的规则其实最复杂，同时要考虑到实体类的内聚和保证不变性（Invariants），
也要考虑跨对象规则代码的归属，甚至要考虑到具体领域服务的调用方式，理解成本比较高。

所以，尽量通过一些设计规范，来降低DDD领域层的设计成本。

##### 5.1 实体类（Entity）
大多数DDD架构的核心都是实体类，实体类包含了一两块大的内容
- 领域里的状态
- 以及对状态操作

Entity最重要的设计原则是保证实体的一致性，也就是说要确保无论外部怎么操作，
一个实体内部的属性都不能出现相互冲突，状态不一致的情况。
```
数据库系统中的数据一致性是指数据库中的数据始终保持有效、准确、完整、可靠和可用的特性。
实体的数据一致性，可以套用这个概念，是指实体的数据始终保持有效、准确、完整、可靠和可用的特性。
```

所以几个设计原则如下：
- 创建即一致
- 尽量避免public setter
- 原则3：通过聚合根保证父子实体的一致性
- 原则4：不可以强依赖其他聚合根实体或领域服务
- 原则5：任何实体的行为只能直接影响到本实体（和其子实体）

- 原则1：创建即一致
- 原则2：尽量避免public setter
- 原则3：通过聚合根保证父子实体的一致性
- 原则4：不可以强依赖其他聚合根实体或领域服务
- 原则5：任何实体的行为只能直接影响到本实体（和其子实体）

##### 5.2 领域服务（Domain Service）
- 类型1：单对象策略型
- 类型3：通用组件型
这种类型的领域服务更像ECS里的System，提供了组件化的行为，但本身又不直接绑死在一种实体类上。

##### 5.3 策略对象（Domain Policy）

#### 6. 使用事件处理领域模型副作用
在上文中，有一种类型的领域规则被刻意忽略了，那就是”领域模型副作用“。
一般的领域模型副作用发生在核心领域模型状态变更后，同步或者异步对另一个对象的影响或行为。
```
一般的领域模型副作用发生在核心领域模型状态变更后，同步或者异步对另一个对象的影响或行为
```
在这个案例里，我们可以增加一个副作用规则：
```
当Monster的生命值降为0后，给Player奖励经验值
```
这种问题有很多种解法，比如直接把副作用写在CombatService里：
```java
public class CombatService {
    public void performAttack(Player player, Monster monster) {
        // ...
        monster.takeDamage(damage);
        if (!monster.isAlive()) {
            player.receiveExp(10); // 收到经验
        }
    }
}
```
但是这样写的问题是：很快CombatService的代码就会变得很复杂，比如我们再加一个副作用：
```
当Player的exp达到100时，升一级
```
这时我们的代码就会变成：
```java
public class CombatService {
    public void performAttack(Player player, Monster monster) {
        // ...
        monster.takeDamage(damage);
        if (!monster.isAlive()) {
            player.receiveExp(10); // 收到经验
            if (player.canLevelUp()) {
                player.levelUp(); // 升级
            }
        }
    }
}
```
如果再加上“升级后奖励XXX”呢？
“更新XXX排行”呢？
依此类推，后续这种代码将无法维护。
所以我们需要介绍一下领域层最后一个概念：领域事件（Domain Event）。

##### 6.1 领域事件介绍
***领域事件是一个通知机制***
领域事件是一个在领域里发生了某些事后，希望领域里其他对象能够感知到的通知机制。
在上面的案例里，代码之所以会越来越复杂，其根本的原因是反应代码（比如升级）直接和上面的事件触发条件（比如收到经验）直接耦合，而且这种耦合性是隐性的。
领域事件的好处就是将这种隐性的副作用“显性化”，通过一个显性的事件，将事件触发和事件处理解耦，最终起到代码更清晰、扩展性更好的目的。
所以，领域事件是在DDD里，比较推荐使用的跨实体“副作用”传播机制。

##### 6.2 领域事件实现
##### 6.3 目前领域事件的缺陷和展望
##### 6.4 领域事件总结
领域事件 是 领域模型副作用的处理方法
一般的 副作用发生在核心领域模型状态变更后，同步或者异步对另一个对象的影响或行为。
领域事件介绍
- 领域事件是一个在领域里发生了某些事后，希望领域里其他对象能够感知到的通知机制。
- 领域事件将隐性的副作用“显性化”，通过一个显性的事件，将事件触发和事件处理解耦，最终起到代码更清晰、扩展性更好的目的。

领域事件实现
- 领域事件通常是立即执行的、在同一个进程内、可能是同步或异步。可通过一个EventBus来实现进程内的通知机制。
- 缺陷：领域事件的很好的实施依赖***EventBus***、***Dispatcher***、***Invoker***这些属于框架级别的支持。
但因为Entity不能直接依赖外部对象，所以EventBus目前只能是一个全局的Singleton，导致Entity对象无法被完整单测覆盖全。

#### 7. 领域层设计规范总结
在真实的业务逻辑里，我们的领域模型或多或少的都有一定的“特殊性”，
***如果100%的要符合DDD规范可能会比较累***，
所以最主要的是梳理一个对象行为的影响面，然后作出设计决策，即：
- 是仅影响单一对象还是多个对象
- 规则未来的拓展性、灵活性
- 性能要求
- 副作用的处理，等等

当然，很多时候一个好的设计是多种因素的取舍，需要大家有一定的积累，真正理解每个架构背后的逻辑和优缺点。
一个好的架构师不是有一个正确答案，而是能从多个方案中选出一个最平衡的方案。

##### 7.1 Entity设计规范总结
- Entity最重要的设计原则是保证实体的一致性，也就是说要确保无论外部怎么操作，
一个实体内部的属性都不能出现相互冲突，状态不一致的情况。

- 原则1：创建即一致
  - constructor参数要包含所有必要属性，或者在constructor里有合理的默认值。
  - 使用Factory模式来降低调用方复杂度

- 原则2：尽量避免public setter
  - 因为set单一参数会导致状态不一致的情况；
  - @ Setter(AccessLevel.PRIVATE) // 确保不生成public setter

- 原则3：通过聚合根保证主子实体的一致性，主实体会包含子实体，主实体就起到聚合根的作用，即：
  - 子实体不能单独存在，只能通过聚合根的方法获取到。任何外部的对象都不能直接保留子实体的引用
  - 子实体没有独立的Repository，不可以单独保存和取出，必须要通过聚合根的Repository实例化
  - 子实体可以单独修改自身状态，但是多个子实体之间的状态一致性需要聚合根来保障

- 原则4：不可以强依赖其他聚合根实体或领域服务。对外部对象的依赖性会直接导致实体无法被单测；
以及***一个实体无法保证外部实体变更后不会影响本实体的一致性和正确性***。正确的对外部依赖的方法有两种：
  - 只保存外部实体的ID：强烈建议使用强类型的ID对象，而不是Long型ID。 强类型的ID对象不单单能自我包含验证代码，保证ID值的正确性，同时还能确保各种入参不会因为参数顺序变化而出bug。
  - 针对于“无副作用”的外部依赖，通过方法入参的方式传入。

- 原则5：任何实体的行为只能直接影响到本实体（和其子实体）
- 原则6：***实体的充血模型不包含持久化逻辑***

##### 7.2 Domain Service 设计规范总结
- 领域服务一般分为单对象策略型、跨对象事务型、通用组件型三种。
- 单对象策略型：主要面向的是单个实体对象的变更，但涉及到多个领域对象或外部依赖的一些规则。
实体应该通过方法入参的方式传入这种领域服务，然后通过Double Dispatch来反转调用领域服务的方法。
- 跨对象事务型：当一个行为会直接修改多个实体时，不能再通过单一实体的方法作处理，而必须直接使用领域服务的方法来做操作。
- 通用组件型：与ECS里的System类似，***提供了组件化的行为，但本身又不直接绑死在一种实体类上***。
- 让Domain Service与Repository打交道，而不是让领域模型Entity与Repository打交道，因为我们想保持领域模型的独立性，
不与任何其他层的代码（Repository）或开发框架（比如Spring、MyBatis）耦合在一起，
将流程性的代码逻辑（比如从DB中取数据、映射数据）与领域模型的业务逻辑解耦，让领域模型更加可复用。
- Domain Service类负责一些非功能性及与三方系统交互的工作。比如幂等、事务、发邮件、发消息、记录日志、调用其他系统的RPC接口等，都可以放到Domain Service类中。
```
Domain Service类负责一些非功能性及与三方系统交互的工作。
比如幂等、事务、发邮件、发消息、记录日志、调用其他系统的RPC接口等，
都可以放到Domain Service类中。
```

##### 7.3 Application Service 设计规范总结
- Application Service 是业务流程的封装，不处理业务逻辑，即不要有if/else分支逻辑、不要有任何计算、一些数据的转化可以交给其他对象来做。
- 常用的ApplicationService“套路”：
  - 准备数据：包括从外部服务或持久化源取出相对应的Entity、VO以及外部服务返回的DTO。
  - 执行操作：包括新对象的创建、赋值，以及调用领域对象的方法对其进行操作。需要注意的是这个时候通常都是纯内存操作，非持久化。
  - 持久化：将操作结果持久化，或操作外部系统产生相应的影响，包括发消息等异步操作。

##### 7.4 其它规范总结
- Interface层：
  - 职责：主要负责承接网络协议的转化、Session管理等
  - 接口数量：避免所谓的统一API，不必人为限制接口类的数量，每个/每类业务对应一套接口即可，接口参数应该符合业务需求，避免大而全的入参
  - 接口出参：统一返回Result
  - 异常处理：***应该捕捉所有异常，避免异常信息的泄漏***。***可以通过AOP统一处理，避免代码里有大量重复代码***。
  
- Application层：
  - 入参：具像化Command、Query、Event对象作为ApplicationService的入参，唯一可以的例外是单ID查询的场景。
  - CQE的语意化：CQE对象有语意，不同用例之间语意不同，即使参数一样也要避免复用。
  - 入参校验：基础校验通过Bean Validation api解决。Spring Validation自带Validation的AOP，也可以自己写AOP。
  - 出参：统一返回DTO，而不是Entity或DO。
  - DTO转化：用DTO Assembler负责Entity/VO到DTO的转化。
  - 异常处理：不统一捕捉异常，可以随意抛异常。
- 部分Infra层：
  - 用ACL防腐层将外部依赖转化为内部代码，隔离外部的影响

关于 Interface层、Application层的设计规范，下一个文章，再详细介绍。
