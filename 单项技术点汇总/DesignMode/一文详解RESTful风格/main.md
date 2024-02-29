## 一文详解RESTful风格

RESTful风格是一种***基于HTTP协议设计Web API的软件架构风格***，由Roy Fielding在2000年提出。
它强调使用***HTTP动词来表示对资源的操作***（GET、POST、PUT、PATCH、DELETE等），***并通过URI表示资源的唯一标识符***。
```
RESTful风格是一种基于HTTP协议设计Web API的软件架构风格
它强调使用
- HTTP动词来表示对资源的操作（GET、POST、PUT、PATCH、DELETE等）
- 并通过URI表示资源的唯一标识符
```

### 一、RESTful API的设计原则
RESTful API的设计遵循以下几个原则：
- 1. 基于资源：***将数据和功能抽象成资源***，并通过URI来唯一标识资源。例如，一个用户资源可以通过URL“/users/{id}”来访问，其中“{id}”表示该用户的唯一标识符。
- 2. 使用HTTP动词：使用HTTP动词来表示对资源的操作，如GET（获取资源）、POST（创建资源）、PUT（更新资源）和DELETE（删除资源）等。
- 3. 无状态：每个请求都包含足够的信息来完成请求，服务器不需要保存任何上下文信息。
- 4. 统一接口：使用统一的接口来简化客户端与服务器之间的交互，包括***资源标识符***、***资源操作***和***响应消息***的格式。
- 5. 可缓存性：客户端可以缓存响应，以提高性能和减少网络流量。
- 6. 分层系统：将系统分为多个层次，每个层次处理特定的功能。

RESTful风格的API设计具有良好的可读性、易用性和可扩展性，
广泛应用于Web应用程序和移动应用程序的API设计中。

### 二、使用到的注解
#### （1）@RequestMapping
- 类型 方法注解
- 位置 SpringMVC控制器方法定义上方
- 作用 设置当前控制器方法请求访问路径
- 范例
```
@RequestMapping(value = "/users", method = RequestMethod.GET)
@ResponseBody
public String save()
{
    System.out.println("save user");
    return " '{'module': 'user save' }' ";
}
```
- 属性
- value 请求访问路径
- method http请求动作，标准动作（GET、POST、PUT、DELETE）

#### （2）@PathVariable
- 类型 形参注解
- 位置 SpringMVC控制器方法形参定义前面
- 作用 绑定路径参数与处理器方法形参间的关系，要求路径参数名与形参名一一对应
- 范例
```
@RequestMapping(value = "users/{id}", method = RequestMethod.DELETE)
@ResponseBody
public String delete(@PathVariable Integer id) // PathVariable 路径参数 id对应路径中的id
{
    System.out.println("delete user");
    return "'{'module': 'user delete'}'";
}
```

#### （3）@RestControl
- 类型 类注解
- 位置 基于SpringMVC的RESTful开发控制器类定义上方
- 作用 设置当前控制器类为RESTful风格，等同于 @Controller 与 @ResponseBody两个注解的组合功能
- 范例
```
@RestController
public class UserController
{
    @RequestMapping(value = "/users",method = RequestMethod.GET)
    public String save()
    {
         System.out.println("save user");
         return " '{'module': 'user save' }' ";
    }
```

#### （4）@GetMapping @PostMapping @PutMapping @DeleteMapping
- 类型 方法注解
- 位置 基于SpringMVC的RESTful开发控制器方法定义上方
- 作用 设置当前控制器方法请求访问路径与请求动作，每种对应一个请求动作
- 范例
```
@RestController
@RequestMapping("/users") // 下面的每个控制器方法的请求路径都有前缀 /users
public class UserController
{
 @GetMapping("/{id}")
 public String getById(@PathVariable Integer id)
    {
        return "getById";
    }
}
```

#### （5）@RequestBody @RequestParam @PathVariable
区别
RequestParam 用于接收URL地址传参或表单传参
RequestBody 用于接收JSON数据
PathVariable 用于接收路径参数，使用 {参数名} 描述路径参数

应用
后期开发中，发送请求参数超过1个时，以JSON格式为主，所以@RequestBody应用较广泛
如果发送非JSON格式数据，选用 @RequestParam 接收请求参数
当参数数量只有一个时，或为数字时，可以采用 @PathVariable接收请求路径变量，通常传递id值

### 三、综合案例
这里提供一个简单的Java示例，用于实现一个基本的RESTful API。
假设我们正在开发一个学生管理系统，需要使用RESTful API来实现对学生资源的增删改查操作。
首先，我们需要定义一个表示学生信息的Java类：
```java
public class Student {
    private int id;
    private String name;
    private int age;
    public Student() { }
    public Student(int id, String name, int age) {
        this.id = id;
        this.name = name;
        this.age = age;
    }
    // Getters and setters
}
```

然后，我们需要创建一个控制器类来处理客户端请求：
```java
@RestController
@RequestMapping("/students")
public class StudentController {
    // Mock data - replace with database queries later
    private static List<Student> students = new ArrayList<>(Arrays.asList(
            new Student(1, "Alice", 20),
            new Student(2, "Bob", 21),
            new Student(3, "Charlie", 22)
    ));
    // GET /students - get all students
    @GetMapping("")
    public List<Student> getAllStudents() {
        return students;
    }
    // GET /students/{id} - get a student by id
    @GetMapping("/{id}")
    public Student getStudentById(@PathVariable int id) {
        for (Student s : students) {
            if (s.getId() == id) {
                return s;
            }
        }
        return null;  // Return null if student not found
    }
    // POST /students - create a new student
    @PostMapping("")
    public ResponseEntity<String> createStudent(@RequestBody Student student) {
        students.add(student);
        return ResponseEntity.status(HttpStatus.CREATED).build();
    }
    // PUT /students/{id} - update an existing student
    @PutMapping("/{id}")
    public ResponseEntity<String> updateStudent(@PathVariable int id, @RequestBody Student updatedStudent) {
        for (int i = 0; i < students.size(); i++) {
            if (students.get(i).getId() == id) {
                students.set(i, updatedStudent);
                return ResponseEntity.status(HttpStatus.OK).build();
            }
        }
        return ResponseEntity.status(HttpStatus.NOT_FOUND).build();  // Return 404 if student not found
    }
    // DELETE /students/{id} - delete a student by id
    @DeleteMapping("/{id}")
    public ResponseEntity<String> deleteStudentById(@PathVariable int id) {
        for (int i = 0; i < students.size(); i++) {
            if (students.get(i).getId() == id) {
                students.remove(i);
                return ResponseEntity.status(HttpStatus.OK).build();
            }
        }
        return ResponseEntity.status(HttpStatus.NOT_FOUND).build();  // Return 404 if student not found
    }
}
```

这个控制器类中定义了四个HTTP方法，分别处理对学生资源的不同操作。
我们使用Spring Boot框架和Spring MVC模块来实现RESTful API，
并使用注解来定义路由和请求处理逻辑。
最后，我们需要在应用程序的入口点（如Spring Boot的main方法）中启动应用程序：
```
@SpringBootApplication
public class StudentManagementSystemApplication {
    public static void main(String[] args) {
        SpringApplication.run(StudentManagementSystemApplication.class, args);
    }
}
```
这样，我们就创建了一个简单的RESTful API，
可以通过发送HTTP请求来执行学生管理系统的基本操作。

