## 手动将本地jar导入到maven库

有时候我们在使用maven管理项目的时候，会出现无法导入jar的情况，
或者说pom.xml中的信息，maven无法全部从远程仓库中拉取到本地，
这样我们在编译项目的时候就无法通过，出现编译错误等问题。

解决的方法有很多，***可以通过网上下载相应的jar包，然后在maven中配置路径，指向jar包位置***，
也可以直接将下载的jar导入到我们本地的maven库中，这里记录下自己是第二种方法操作步骤。

以ik分词jar包为例，从网上下载相应的jar，放到D:\Users\Downloads\IKAnalyzer\IKAnalyzer\IKAnalyzer3.2.0Stable_bin目录下

```shell
mvn install:install-file -Dfile=("jar包的位置") 
-DgroupId=groupId("分组") 
-DartifactId=artifactId("jar名称") 
-Dversion=version("版本号") 
-Dpackaging=jar
```
































