## Java网络编程之BIO(Socket)-

```
现在流行NIO网络编程，比较火的框架有Netty和Mina,这个地方我实现传统Socket编程，
每一个请求，都会为之创建一个线程来进行处理操作，在Socket数据传输中，用到了PrintWriter，
需要注意println()和write()两个方法的区别，println()是带有回车符号的写数据，而write()没有回车符的些数据，
需要手动加上回车符，不然消息发送不过去，因为回车符是消息结束的标志，没有回车符，就认为消息没读完。
```

### 概念简介
Java对BIO、NIO、AIO的支持：
- 1、Java BIO ： 同步并阻塞，服务器实现模式为一个连接一个线程，即客户端有连接请求时服务器端就需要启动一个线程进行处理，
如果这个连接不做任何事情会造成不必要的线程开销，当然可以通过线程池机制改善。

- 2、Java NIO ： 同步非阻塞，服务器实现模式为一个请求一个线程，即客户端发送的连接请求都会注册到多路复用器上，
多路复用器轮询到连接有I/O请求时才启动一个线程进行处理。

- 3、Java AIO(NIO.2) ： 异步非阻塞，服务器实现模式为一个有效请求一个线程，客户端的I/O请求都是由OS先完成了再通知服务器应用去启动线程进行处理

BIO、NIO、AIO适用场景分析:
- 1、BIO方式适用于连接数目比较小且固定的架构，这种方式对服务器资源要求比较高，并发局限于应用中，JDK1.4以前的唯一选择，但程序直观简单易理解。
- 2、NIO方式适用于连接数目多且连接比较短（轻操作）的架构，比如聊天服务器，并发局限于应用中，编程比较复杂，JDK1.4开始支持。
- 3、AIO方式使用于连接数目多且连接比较长（重操作）的架构，比如相册服务器，充分调用OS参与并发操作，编程比较复杂，JDK7开始支持。

### 案例
服务端
服务端中，定义了一个SocketHandler，用于处理消息和数据的，每进来一个消息，就给他创建一个线程，这种传统的方式，效率比较的低下
```java
package yellowcong.socket;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.OutputStreamWriter;
import java.io.PrintWriter;
import java.net.ServerSocket;
import java.net.Socket;

/**
 * 创建日期:2017年10月6日 <br/>
 * 创建用户:yellowcong <br/>
 * 功能描述:
 */
public class Server {

    //端口
    public static final Integer PORT = 8080;

    //ip地址
    public static final String ADDRESS = "127.0.0.1";

    /**
     * 创建日期:2017年10月6日<br/>
     * 创建用户:yellowcong<br/>
     * 功能描述:
     * 
     * @param args
     * @throws Exception
     */
    public static void main(String[] args) throws Exception {

        // 创建Socket的链接
        ServerSocket server = new ServerSocket(PORT);

        System.out.println("Socket服务器端,启动");

        // 接受数据, 线程会阻塞到这个地方，这里创建的socket和客户端socket形成一对
        Socket socket = server.accept();

        System.out.println("连接上了");

        new Thread(new SocketHandler(socket)).start();
    }
}
```

消息处理
消息的读取用到了BufferedReader，消息的写用到了PrintWriter ,写消息的时候，建议使用 println()方法，如果非要用write方法，就需要在消息后面添加回车符，
不然数据传输过去后，读取方法的时候，BufferedReader中readLine()需要读取到回车符，才可以结束。

```java
package yellowcong.socket;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.OutputStreamWriter;
import java.io.PrintWriter;
import java.net.Socket;

/**
 * 创建日期:2017年10月6日 <br/>
 * 创建用户:yellowcong <br/>
 * 功能描述:  
 */
public class SocketHandler implements Runnable{
    private Socket socket ;
    public SocketHandler(Socket socket) {
        super();
        this.socket = socket;
    }

    public void run() {
        //获取写的数据,自动刷新
        BufferedReader in = null;
        //写数据
        PrintWriter out = null;

        try {
            //获取写的数据,自动刷新
            in = new BufferedReader(new InputStreamReader(socket.getInputStream()));
            //写数据， 第二个参数，表示自动flush
            out = new PrintWriter(new OutputStreamWriter(socket.getOutputStream()),true);

            String content = null;

            while(true){
                content = in.readLine();
                if(content == null){
                    break;
                }
                System.out.println("服务器收到:\t"+content);
                out.println("逗比连接上了");
            }
        } catch (IOException e) {
            // TODO Auto-generated catch block
            e.printStackTrace();
        }finally{
            try {
                if(in != null){
                    in.close();
                }
            } catch (IOException e) {
                // TODO Auto-generated catch block
                e.printStackTrace();
            }

            if(out != null){
                out.close();
            }

            socket = null;
        }
    }

}
```




























