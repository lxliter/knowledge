## Java中如何停止一个线程
Java中有以下三种方法可以终止正在运行的线程：
使用退出标志，使线程正常退出，也就是当 run() 方法完成后线程中止。
这种方法需要在循环中检查标志位是否为 true，如果为 false，则跳出循环，结束线程。

使用 stop() 方法强行终止线程，但是不推荐使用这个方法，该方法已被弃用。
***这个方法会导致一些清理性的工作得不到完成，如文件，数据库等的关闭，以及数据不一致的问题***。

使用 interrupt() 方法中断线程。
这个方法会在当前线程中打一个停止的标记，并不是真的停止线程。
因此需要在线程中判断是否被中断，并增加相应的中断处理代码。
如果线程在 sleep() 或 wait() 等操作时被中断，会抛出 InterruptedException 异常。

### 使用标记位中止线程
使用退出标志，使线程正常退出，也就是当 run() 方法完成后线程中止，是一种比较简单而安全的方法。
这种方法需要在循环中检查标志位是否为 true，如果为 false，则跳出循环，结束线程。
这样可以保证线程的资源正确释放，不会导致数据不一致或其他异常问题。

例如，下面的代码展示了一个使用退出标志的线程类：
```java
public class ServerThread extends Thread {
    // volatile修饰符用来保证其它线程读取的总是该变量的最新的值
    public volatile boolean exit = false;
    @Override
    public void run() {
        ServerSocket serverSocket = new ServerSocket(8080);
        while (!exit) {
            serverSocket.accept(); // 阻塞等待客户端消息
            // do something
        }
    }
}
```
在主方法中，可以通过修改标志位来控制线程的退出：
```
public static void main(String[] args) {
    ServerThread t = new ServerThread();
    t.start();
    //do something else
    t.exit = true; //修改标志位，退出线程
}
```

这种方法的优点是简单易懂，缺点是需要在循环中不断检查标志位，可能会影响性能。
另外，如果线程在 sleep() 或 wait() 等操作时被设置为退出标志，它也不会立即响应，
而是要等到阻塞状态结束后才能检查标志位并退出。

### 使用 stop() 方法强行终止线程
使用 stop() 方法强行终止线程，是一种不推荐使用的方法，因为它会导致一些严重的问题。
stop() 方法会立即终止线程，不管它是否在执行一些重要的操作，
如***关闭文件***，***释放锁***，***更新数据库***等。
这样会导致***资源泄露***，***数据不一致***，或者其他异常错误。

stop() 方法会立即释放该线程所持有的所有的锁，
导致数据得不到同步，出现数据不一致的问题。
例如，如果一个线程在修改一个对象的两个属性时被 stop() 了，那么可能只修改了一个属性，而另一个属性还是原来的值。
这样就造成了对象的状态不一致。
例如，下面的代码展示了一个使用 stop() 方法的线程类：
```java
public class MyThread extends Thread {
    @Override
    public void run() {
        try {
            // 文件流操作
            FileWriter fw = new FileWriter("test.txt");
            fw.write("Hello, world!");
            Thread.sleep(1000); // 模拟耗时操作
            fw.close(); // 关闭文件
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}
```
在主方法中，可以通过调用 stop() 方法来强行终止线程：
```
public static void main(String[] args) {
    MyThread t = new MyThread();
    t.start();
    //do something else
    t.stop(); //强行终止线程
}
```

这种方法的缺点是很明显的，如果在关闭文件之前调用了 stop() 方法，那么文件就不会被正确关闭，可能会造成数据丢失或损坏。
而且，stop() 方法会抛出 ThreadDeath 异常，如果没有捕获处理这个异常，那么它会向上层传递，可能会影响其他线程或程序的正常运行 。
因此，***使用 stop() 方法强行终止线程是一种非常危险而不负责任的做法，应该尽量避免使用***。

### 使用interrupt() 方法中断线程
Thread.interrupt()它能帮助我们在一个线程中断另一个线程。
尽管它被命名为“interrupt”，但实际上它并不会立即停止一个线程的执行，而是设置一个中断标志，表示这个线程已经被中断。
它的具体行为取决于被中断线程当前的状态以及如何响应中断。
***interrupt是Thread对象一个内部字段，用来表示它的中断状态***。
***这个字段是由Java虚拟机（JVM）管理的，对应用程序代码是不可见的***。

以下是有关Thread.interrupt()的一些重要事项：
对于非阻塞状态的线程：***如果线程处于运行状态，并且没有执行任何阻塞操作***，那么调用interrupt()方法只会设置线程的中断状态，并不会影响线程的继续执行。
线程需要自己检查这个中断状态，并决定是否停止执行。
常见的检查方式包括调用Thread.interrupted()（这会清除中断状态）或者Thread.currentThread().isInterrupted()（不会清除中断状态）。
- Thread.interrupted()（这会清除中断状态）
- Thread.currentThread().isInterrupted()（不会清除中断状态）

对于阻塞状态的线程：如果线程处于阻塞状态，如调用了
- Object.wait(), 
- Thread.join()或者
- Thread.sleep()方法，那么线程会立即抛出InterruptedException，并且清除中断状态。

对于已经停止的线程：如果线程已经停止，那么调用interrupt()方法不会有任何影响。

interrupt()方法为我们提供了一种通用的、协作式的线程停止机制。
它允许被中断的线程决定如何处理中断请求，可以立即停止，也可以忽略中断，或者继续执行一段时间然后再停止。

```
interrupt()方法为我们提供了一种通用的、协作式的线程停止机制。
它允许被中断的线程决定如何处理中断请求，
可以立即停止，也可以忽略中断，或者继续执行一段时间后再停止。
```

以下是一个使用interrupt()方法的例子：
```
Thread t = new Thread(() -> {
    while (!Thread.currentThread().isInterrupted()) {
        // 执行任务
    }
});

t.start();

// 在另一个线程中中断t线程
t.interrupt();
```
这个例子中，线程t会一直执行，直到它的中断状态被设置。
这是通过检查Thread.currentThread().isInterrupted()实现的。
当t.interrupt()被调用时，线程t的中断状态被设置，因此线程将退出循环并结束执行。

需要注意的是，如果线程在响应中断时需要执行一些清理工作，或者需要抛出一个异常来通知上游代码，
那么就需要在捕获InterruptedException后，手动再次设置中断标志。
这是因为当InterruptedException被抛出时，中断状态会被清除。例如：

```
while (!Thread.currentThread().isInterrupted()) {
    try {
        // 执行可能抛出InterruptedException的任务
        Thread.sleep(1000);
    } catch (InterruptedException e) {
        // 捕获InterruptedException后，再次设置中断标志
        Thread.currentThread().interrupt();
    }
}
```
