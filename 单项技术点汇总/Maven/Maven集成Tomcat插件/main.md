## Maven集成Tomcat插件 

### 目录
- 类似插件及版本区别：
- 本地运行，***启动嵌入式tomcat***:
  - 错误一：
  - 错误二：
  - Idea运行调试：
  - vscode运行调试：
- 远程部署：
  - 项目中的pom.xml配置：
  - Tomcat中的tomcat-users.xml配置：
  - Maven中的settings.xml配置：
  - 注意事项：
- Tomcat插件命令：
- 参考：

### 类似插件及版本区别
Maven Tomcat插件现在主要有两个版本，tomcat-maven-plugin和tomcat7-maven-plugin，使用方式基本相同。
tomcat-maven-plugin 插件官网：http://mojo.codehaus.org/tomcat-maven-plugin/plugin-info.html。
tomcat7-maven-plugin 插件官网：http://tomcat.apache.org/maven-plugin.html。

tomcat-maven-plugin这个插件是老版本，不知道是被apache收购还是怎么的，现在已经停用，
命令是mvn tomcat:run，而且该插件应该也不支持tomcat7

Apache内部这个插件现在也有两个版本，分别是tomcat6,tomcat7

tomcat6:
```xml
<plugin>
	<groupId>org.apache.tomcat.maven</groupId>
	<artifactId>tomcat6-maven-plugin</artifactId>
	<version>2.2</version>
	<configuration>
		<url>http://127.0.0.1:8080/manager</url>
		<server>tomcat</server>
		<username>admin</username> 
		<password>admin</password>
		<path>/dev_web</path><!--WEB应用上下文路径-->
		<contextReloadable>true</contextReloadable>
	</configuration>
</plugin>
```

tomcat7:
```xml
<plugin>
	<groupId>org.apache.tomcat.maven</groupId>
	<artifactId>tomcat7-maven-plugin</artifactId>
	<version>2.2</version>
	<configuration>
		<url>http://127.0.0.1:8080/manager</url>
		<server>tomcat</server>
		<username>admin</username> 
		<password>admin</password>
		<path>/dev_web</path><!--WEB应用上下文路径-->
		<contextReloadable>true</contextReloadable>
	</configuration>
</plugin>
```

下面的篇幅中就只讨论tomcat7
本地运行，启动嵌入式tomcat:
```xml
<plugin>   
    <groupId>org.apache.tomcat.maven</groupId>   
    <artifactId>tomcat7-maven-plugin</artifactId>   
    <version>2.2</version>   
    <configuration>      
        <hostName>localhost</hostName>        <!--   Default: localhost -->  
        <port>8080</port>                     <!-- 启动端口 Default:8080 --> 
        <path>/</path>   <!-- 访问应用路径  Default: /${project.artifactId}-->  
        <uriEncoding>UTF-8</uriEncoding>      <!-- uri编码 Default: ISO-8859-1 -->
    </configuration>
</plugin>
```

如果在启动运行过程中报异常：
错误一：
```
Unknown default host [localhost] for connector [Connector[HTTP/1.1-8083]]
```

***那么把hostName改成localhost即可***。

错误二：
```
[ERROR] Failed to execute goal org.apache.tomcat.maven:tomcat7-maven-plugin:2.2:run (default-cli) on project springdemo-list: 
Could not start Tomcat: Failed to start component [StandardServer[-1]]: Failed to start component [StandardService[Tomcat]]: 
Failed to start component [StandardEngine[Tomcat]]: ***A child container failed during start*** -> [Help 1]
[ERROR]
[ERROR] To see the full stack trace of the errors, re-run Maven with the -e switch.
[ERROR] Re-run Maven using the -X switch to enable full debug logging.
[ERROR]
[ERROR] For more information about the errors and possible solutions, please read the following articles:
[ERROR] [Help 1] http://cwiki.apache.org/confluence/display/MAVEN/MojoExecutionException
```

那么一般就是
```xml
<dependency>
    <groupId>javax.servlet</groupId>
    <artifactId>javax.servlet-api</artifactId>
    <version>3.1.0</version>
    <scope>provided</scope>
</dependency>
```

这个依赖包添加scope为provided就可以

```
意思是这个servlet-api的依赖包只在编译和测试时使用而不在运行时使用；
因为web容器自身一般都会带这些依赖包，故配置上scope。
假如不配置此项，启动tomcat时出现上述的异常，个人认为是由于我们自己在pom.xml引入的依赖跟web容器自己的一些依赖包冲突导致。
```

Idea运行调试：
这种内嵌tomcat方式启动项目，直接命令操作即可
但是如果想要调试，就必须使用编辑器的maven插件，

```
比如idea,直接在Run/Debug Configuration->Maven->Commandline中输入 clean tomcat7:run 即可
上述方式调试，页面修改可以直接显示，后台代码可以使用Jrebel热部署
```

vscode运行调试：
launch.json中添加配置：

```json
{
    "type": "java",
    "name": "Debug (Attach)",
    "request": "attach",
    "hostName": "localhost",
    "port": 8000
}
```

命令启动：
```shell
mvnDebug -DskipTests tomcat7:run -Pirm-web -Pdev
```

再启动vscode中的启动按钮

### 远程部署：
项目中的pom.xml配置：
```
<!-- Tomcat 自动部署 plugin -->
<plugin>
    <groupId>org.apache.tomcat.maven</groupId>
    <artifactId>tomcat7-maven-plugin</artifactId>
    <version>2.2</version>
    <configuration>
        <!-- 对应的 tomcat manager的接口-->
        <url>http://127.0.0.1:8080/manager/text</url>
        <!-- setting.xml 的server id-->
        <server>tomcat</server>
        <!-- tomcat-user.xml 的 username -->
        <username>admin</username>
        <!-- tomcat-user,xml 的 password -->
        <password>admin</password>
        <!-- web项目的项目名称-->
        <path>/${project.artifactId}</path>
        <uriEncoding>UTF-8</uriEncoding>
        <!-- 若tomcat项目中已存在，且使"mvn tomcat7:deploy"命令必须要设置下面的代码 -->
        <!-- 更新项目时，仅需要执行"mvn tomcat7:redeploy"命令即可 -->
        <!-- 上述命令无论服务器是tomcat7、8或9，均是使用"mvn tomcat7:deploy"或"mvn tomcat7:redeploy" -->
        <update>true</update>
    </configuration>
</plugin>
```

```
url：打包好的包通过这个url上传到tomcat处
path：这里如果设置/，默认就是ROOT,最好设置为项目名称，这样可以在一个端口下部署多个项目
```

Tomcat中的tomcat-users.xml配置：
```
<role rolename="admin"/>
<role rolename="manager-script"/>
<role rolename="manager-gui"/>
<role rolename="manager-jmx"/>
<role rolename="manager-status"/>
<role rolename="admin-gui"/>
<user username="admin" password="admin" roles="manager-gui,admin,manager-jmx,manager-script,manager-status,admin-gui" />
```

注意tomcat一定要重启才会生效
Maven中的settings.xml配置：
```xml
<server>
	<id>tomcat</id>
	<username>admin</username>
	<password>admin</password>
</server>
```
注意事项：
tomcat一定要启动才能部署项目
Tomcat插件命令：
```
tomcat7:deploy --部署web war包
tomcat7:redeploy --重新部署web war包
tomcat7:undeploy --停止该项目运行，并删除部署的war包
tomcat7:run --启动嵌入式tomcat ，并运行当前项目
tomcat7:exec-war --创建一个可执行的jar文件，允许使用java -jar mywebapp.jar 运行web项目
tomcat7:help --在tomcat7-maven-plugin显示帮助信息
```


