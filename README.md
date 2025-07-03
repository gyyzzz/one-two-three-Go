# Go语言自学笔记

| 日期       | 内容                                                | 链接        |
| ---------- | --------------------------------------------------- | ----------- |
| 2025.07.03 | 记录本文初稿、学习（译者序、前言、入门1.1 1.2章节） | [1](#jump1) |

## 一、为什么要学

  中国有句古话，三思而后行。英文就是one, two ,three , Go.所以三思后要学Go（开个玩笑）。

以下内容源自互联网：

  Golang —— 让你用写Python代码的开发效率编写出C++代码的性能。

  Go语言(或Golang)是Golanguage的简称，Go是Google的Ken Thompson，Rob Pike以及Robert Griesemer开发的一种静态强类型、编译并发型语言。

  学习Go语言的主要原因是其在并发编程、系统编程和云原生领域的高效性和简洁性。Go语言特别适合构建高性能的网络服务、分布式系统、云计算基础设施和容器化应用。此外，Go语言的学习曲线相对平缓，对于有一定编程经验的开发者来说，上手比较容易。

   此外在工作中接触到的许多优秀的云原生应用如Docker，Kubernetes，etcd，Prometheus等都使用Go语言开发，更引发了我的兴趣，所以定一个目标，希望能学习掌握Go语言，并尝试在自己喜欢的基于go开发的项目中提交贡献。

## 二、学习记录

后续学习基于[《The Go Programming Language》（Go语言圣经中文版）](https://gopl-zh.github.io/index.html)进行，记录个人觉得书中关键和有意思的内容、代码练习等。

### **前言**

后续参考引用会尽量表明，*“斜体”*代表直接引用书中内容，其他引自互联网的内容会特别标注。

本人学生阶段疏于学习，水平有限，学习过程中的一些奇奇怪怪的疑问会记录成Q&A的形式，A一般是从AI那得到的答案。



*“Go语言有时候被描述为“类C语言”，或者是“21世纪的C语言”。”*

*“正如[Rob Pike](http://genius.cat-v.org/rob-pike/)所说，“软件的复杂性是乘法级相关的”，通过增加一个部分的复杂性来修复问题通常将慢慢地增加其他部分的复杂性。通过增加功能、选项和配置是修复问题的最快的途径，但是这很容易让人忘记简洁的内涵，即从长远来看，简洁依然是好软件的关键因素。”*



---

Q：666

A：😄 你太客气了！

---

**环境安装**

https://go.dev/doc/install, 选择对应版本安装

```shell
❯ go version
go version go1.24.4 darwin/arm64
```

后面的标题对应书中标题

### 1.<span id="jump1">入门</span>

#### 1.1.Hello, World

“*我们以现已成为传统的“hello world”案例来开始吧，这个例子首次出现于 1978 年出版的 C 语言圣经 《The C Programming Language》*”

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, world")
}
```

```
❯ vi helloworld.go
❯ go run helloworld.go
Hello, world
```

---

Q：go run helloworld.go 运行速度慢，同样的写一个.py脚本执行速度要快n倍？

A：关键区别：Go 是编译型语言，Python 是解释型语言。py直接调用解释器执行 `.py` 文件，不需要编译，几乎是**即点即运行**。`go run`实际上是两步过程：先 **临时编译** 成二进制文件（放在临时目录下，如 `/tmp`）然后再执行这个临时二进制文件。

```
❯ date;go run helloworld.go;date
2025年 7月 3日 星期四 11时35分15秒 CST
Hello, world
2025年 7月 3日 星期四 11时35分16秒 CST
❯ go build helloworld.go
❯ ls
helloworld    helloworld.go
❯ date;./helloworld;date
2025年 7月 3日 星期四 11时35分39秒 CST
Hello, world
2025年 7月 3日 星期四 11时35分39秒 CST
```

可以看到运行编译后的二进制文件会快很多

---



*Go 语言的代码通过**包**（package）组织，包类似于其它语言里的库（libraries）或者模块（modules）。一个包由位于单个目录下的一个或多个 `.go` 源代码文件组成，目录定义包的作用。每个源文件都以一条 `package` 声明语句开始，这个例子里就是 `package main`，表示该文件属于哪个包，紧跟着一系列导入（import）的包，之后是存储在这个文件里的程序语句。*

*Go 语言不需要在语句或者声明的末尾添加分号，除非一行上有多条语句。*

*Go 语言在代码格式上采取了很强硬的态度。`gofmt`工具把代码格式化为标准格式*

---

Q：go get/install 访问proxy.golang.org超时的解决方式：

```
go env -w GOPROXY=https://goproxy.cn
```

```
❯ go install  golang.org/x/tools/cmd/goimports@latest
go: downloading golang.org/x/tools v0.34.0
go: downloading golang.org/x/mod v0.25.0
go: downloading golang.org/x/sync v0.15.0
```

---



#### 1.2.命令行参数



```go
// Echo2 prints its command-line arguments.
package main

import (
    "fmt"
    "os"
)

func main() {
    s, sep := "", ""
    for _, arg := range os.Args[1:] {
        s += sep + arg
        sep = " "
    }
    fmt.Println(s)
}

```

*符号 `:=` 是 短变量声明（short variable declaration）的一部分，这是定义一个或多个变量并根据它们的初始值为这些变量赋予适当类型的语句。*

*Go 语言只有 `for` 循环这一种循环语句。`for` 循环有多种形式，其中一种如下所示：*

```go
for initialization; condition; post {
    // zero or more statements
}
```



**练习 1.1：** 修改 `echo` 程序，使其能够打印 `os.Args[0]`，即被执行命令本身的名字。

**练习 1.2：** 修改 `echo` 程序，使其打印每个参数的索引和值，每个一行。

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args {
		if i == 0 {
			fmt.Println("os.Args[0]= ", arg)
		}
		fmt.Println(i, arg)
	}
}
```

**练习 1.3：** 做实验测量潜在低效的版本和使用了 `strings.Join` 的版本的运行时间差异。

```go
// compare.
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	now := time.Now()
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	end := time.Now()
	fmt.Println("cost time:", end.Sub(now))

	n_now := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	n_end := time.Now()
	fmt.Println("Join funcation cost time:", n_end.Sub(n_now))
}
```

