## 协程间的三种通信方式

golang中要想灵活运用协程来解决问题，协程间的通信一定要掌握。
这里列举了3种方式。sync.waitgroup 其实也应该掌握

### 1、channel
```go
package main

import (
   "fmt"
   "time"
)

func main() {
	// 创建管道
    ch := make(chan string)
    // 向管道中发送数据
    go sendData(ch)
	// 从管道中获取数据
    go getData(ch)
    // 睡眠
    time.Sleep(1e9)
}

func sendData(ch chan string){
	// 向管道ch中写入golang
    ch <- "golang"
}

func getData(ch chan string){
	// 从管道中获取数据，如果没有数据，则阻塞
    fmt.Println(<- ch)
}
```

### 2、context
```
// 使用 Context 控制多个 goroutine
func contextPart4() {
    ctx, cancel := context.WithCancel(context.Background())
    go watch(ctx, "task 1")
    go watch(ctx, "task 2")
    go watch(ctx, "task 3")
 
    time.Sleep(10 * time.Second)
    log.Println("可以通知任务结束")
    // 当调用cancel时，发出done事件
    cancel()
    time.Sleep(5 * time.Second)
}
 
func watch(ctx context.Context, name string) {
    // selector+channel->epoll事件轮询
    for {
        select {
        case <-ctx.Done():
            log.Println(name + " 任务即将要退出了...")
            return
        default:
            log.Println(name + " goroutine 继续处理任务中...")
            time.Sleep(2 * time.Second)
        }
    }
}
```

### sync.cond
```
// 定义锁
var locker sync.Mutex 
// 包装锁对象->NewCond(l Locker)里面定义的是一个接口,拥有lock和unlock方法。
// 看到sync.Mutex的方法,func (m *Mutex) Lock(),可以看到是指针有这两个方法,所以应该传递的是指针 
var cond = sync.NewCond(&locker) 
 
func main() { 
    // 启动多个协程 
    for i := 0; i < 10; i++ { 
        // go 方法启动协程
        go func(x int) { 
            cond.L.Lock()          // 获取锁 
            defer cond.L.Unlock()  // 释放锁 
           
            cond.Wait()            // 等待通知，阻塞当前 goroutine 
           
            // 通知到来的时候, cond.Wait()就会结束阻塞, do something. 这里仅打印 
            fmt.Println(x) 
        }(i) 
    } 
   
    time.Sleep(time.Second * 1) // 睡眠 1 秒，等待所有 goroutine 进入 Wait 阻塞状态 
    fmt.Println("Signal...") 
    cond.Signal()               // 1 秒后下发一个通知给已经获取锁的 goroutine 
   
    time.Sleep(time.Second * 1) 
    fmt.Println("Signal...") 
    cond.Signal()               // 1 秒后下发下一个通知给已经获取锁的 goroutine 
   
    time.Sleep(time.Second * 1) 
    cond.Broadcast()            // 1 秒后下发广播给所有等待的goroutine 
    fmt.Println("Broadcast...") 
    time.Sleep(time.Second * 1) // 等待所有 goroutine 执行完毕 
} 
```
























