# go-IM

教程作者刘丹冰，这是视频地址[8 小时转职 Golang 工程师(如果你想低成本学习 Go 语言)](https://www.bilibili.com/video/BV1gf4y1r79E)

感谢作者

## v1

这是基础服务，分为两个文件，一个 main 和 server，我感觉这个还是用对象思维，不过 go 是将函数模拟成对象，用方法接收者方式为 struct 增加方法

这里也是如此

首先一个 server 的 struct

```go
type Server struct {
	Ip   string
	Port int
}
```

引入了 net 的库，三个方法，一个是新建 server，类比是构造函数，Start 是启动函数，net 监听 tcp 地址和端口，用 for 无限循环执行接受和 handler 函数

handler 函数的动作就是打印一条链接成功消息

还有 main

其实只是两个动作，一个新建 server 和开始执行新建 server 的 start 方法，不过没注意的是，这虽然是两个文件其实是在一个 package 下，也就是 build 后在一个空间之中

还是很简单的

## v2

这个版本加上了上线提醒，也就是一个用户上线后，通知所有用户并发送消息，这个我第一反应是 qq 的敲门声

这次新建了一 user 文件，用于描述用户

一个 structural 来描述元数据，怎么感觉这个就是私有变量，用户名，地址，c 信道和一个 conn 网络连接接口

绑定一个方法，ListenMessage，将信道 chan 的数据发送出去

这里的信道 chan 有点大道至简的感觉

剩下就是改造 server

首先是在 server 的 structural 里增加一个用户列表的 map，lock 锁，message 的 chan 信道

用户列表用于遍历用户发送消息，lock 互斥锁保证同一时间只有一个 go 程读取数据，message 的 chan 信道在不同 go 程之间传递消息

函数 Handler 改造，每次上线一个便新建一个 user，将用户添加到 map 中，这个动作是在互斥锁中进行，上线消息传递给 message 这个 chan 信道，这个实现在 BroadCast 函数中，简单来说就是生成一个字符串，传递给 message

最后是将消息发送给所有用户的 ListenMessage 函数，for 无限循环，读取 message 信道消息，遍历用户 map 列表，将消息发送到 user 的 c 信道，这个动作也是在互斥锁中实现

最后在 start 函数中增加 go 程-监听消息并发送消息的 ListenMessage 函数，让其在后台一直循环读取

for 循环，go 程开启，chan 信道传数据，go 果然大道至简

## v3

这个版本增加了一个读取用户消息的功能，也就是在 server 的 handler 中增加一个 go 程，应该是属于匿名函数，设置一个切片，存储字节，for 无限循环，从 conn 中读取流后，将进行两个判断，一个是字节为 0，一个是否读取错误，最后将流转化为 string 后通过 BroadCast 进行广播

## v4

这个版本是将一些关于聊天的业务封装到 user 中，三个功能，上线添加到 user 列表，下线将剔除 list 并广播消息，广播消息

其实主要是将 server 指针传到用户里，我确认这就是面向对象思维，用户的数据和行为放在一起，可是，go 不是为了面向对象设计的啊

## v5

将用户列表信息发送到当前咨询的用户窗口，也就是在 DoMessage 函数中增加一个判断，当用户输入 who 时候，遍历用户列表 mpa，将用户姓名发送到当前用户的 conn，也就是链接，具体实现是判断，后面用互斥锁包裹读取服务器的用户列表值，user 新增一个函数，将消息发送给自己，这里调用的是 conn

## v6

for 无限循环的话，是指会读取一次，也就是消息通过信道传递到用户，只会发生一次，这背后的逻辑是什么呢

添加了修改用户名，这一块主要是在 DoMessage 中增加一个 else if，检测出现指定的消息，rename|后将进入设置用户名模式，将 rename|后的字符串设置为新的用户名，这里还有个逻辑，就是判断该用户名是不是已被使用，如果是，发送消息通知，否则进入修改用户名逻辑，也还是在互斥锁中，先删除 server 中用户 map 中的当前用户，在用新的用户名作为键与值对应

## v7

添加了超时踢人，这个功能是一个人如果规定时间内没有发言就强制下线，关闭链接，而他的实现方式是监听一个变量，也就是 isLive，这是一个信道变量，可以在 go 程中传递状态，在接受客户端消息发送时设置为 true，而有一个 select 阻塞，select 语句使一个 Go 程可以等待多个通信操作，两个 case，一个当 isLive 时不做任何动作只是更新 select，第二个是定时器，一旦触发，就会发送给客户端消息，并消耗 user 的信道 C，关闭链接，return 退出 Handel，准确说是退出当前 go 程的 handle，这里让我最疑惑的是 select，也就是他是一个等待信道来触发，然后执行相关的操作

```go
package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 50; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
```

用这段代码解释下，这里是两个函数，一个斐波那契数列，一个是 main 主函数

调用逻辑是这样的，首先 main 主函数两个信道变量，c 与 quit，一个 go 程，异步自动执行，一个斐波那契数列函数，我理解是这个函数会在 go 程异步执行时执行，也就是不等待，直接开始，传入两个信道变量

斐波那契数列函数里设置两个变量，作为初始变量，一个 for 无限循环，将执行一个 select 阻塞，执行分支，便是第一个，将初始变量传递给信道变量中，而此时主函数开始打印信道 c，循环结束后将执行传递 0 到 quit，select 执行第二个 case，打印 quit 并退出 select

这里我不明白的是这个逻辑是如何实现的，我可以理解 for 无限循环，select 一直后台监视 case，也就是信道变量，循环打印时出发了读取 c 信道变量，而 0 传递到 quit，则是触发了第二个 case，只能说 go 的思维和我之前的写的并不一致

## v8

私聊功能，这个只是在 DoMessage 函数中再次新增一个处理特殊消息的函数，类似改名，读取到特定的字符串，to|名字|消息，将这个字符串分割，按照姓名查找对应的用户，将消息发送到对方的端口

## v9
