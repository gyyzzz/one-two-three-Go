# Go语言自学笔记

| 日期       | 内容                  |
| ---------- | --------------------- |
| 2025.07.03 | 译者序、前言、1.1 1.2 |
| 2025.07.04 | 1.3                   |
| 2025.07.07 | 1.4、1.5              |
| 2025.07.09 | 1.6、1.7              |
| 2025.07.15 | 2                     |

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

---

*“Go语言有时候被描述为“类C语言”，或者是“21世纪的C语言”。”*

*“正如[Rob Pike](http://genius.cat-v.org/rob-pike/)所说，“软件的复杂性是乘法级相关的”，通过增加一个部分的复杂性来修复问题通常将慢慢地增加其他部分的复杂性。通过增加功能、选项和配置是修复问题的最快的途径，但是这很容易让人忘记简洁的内涵，即从长远来看，简洁依然是好软件的关键因素。”*

---

本人学生阶段疏于学习，水平有限，学习过程中的一些奇奇怪怪的疑问会记录成Q&A的形式，A一般是从AI那得到的答案。

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

#### 1.3.查找重复的行

本小结以查找重复行为案例，提及了多种读取方式（标准输入、文件逐行读取、一次性读取）、`map`、`Printf`、参数传递、错误处理等内容，有一个练习题。

从标准输入读取内容

```go
// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
  //创建一个字符串到整数的 map，记录每行文本出现的次数。
    counts := make(map[string]int)
  //创建一个 Scanner 对象，从标准输入读取文本（逐行读取）。
    input := bufio.NewScanner(os.Stdin)
  /*input.Scan() 会读取下一行内容，直到遇到 EOF（结束输入）。
		input.Text() 返回当前行的字符串。
		每读取一行，就将该行作为键加入 map 并递增次数。
	*/
    for input.Scan() {
        counts[input.Text()]++
    }
    // NOTE: ignoring potential errors from input.Err()
  
    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}

```

从标准输入读取`bufio.NewScanner(os.Stdin)`

*内置函数 `make` 创建空 `map`*

*基于 `range` 的循环，并在 `counts` 这个 `map` 上迭代。跟之前类似，每次迭代得到两个结果，键和其在 `map` 中对应的值。**`map` 的迭代顺序并不确定**，从实践来看，该顺序随机，每次运行都会变化。这种设计是有意为之的，*

*类似于 C 或其它语言里的 `printf` 函数，`fmt.Printf` 函数对一些表达式产生格式化输出。*

*`Printf` 有一大堆这种转换，Go程序员称之为动词（verb）。下面的表格虽然远不是完整的规范，但展示了可用的很多特性：*

```text
%d          十进制整数
%x, %o, %b  十六进制，八进制，二进制整数。
%f, %g, %e  浮点数： 3.141593 3.141592653589793 3.141593e+00
%t          布尔：true或false
%c          字符（rune） (Unicode码点)
%s          字符串
%q          带双引号的字符串"abc"或带单引号的字符'c'
%v          变量的自然形式（natural format）
%T          变量的类型
%%          字面上的百分号标志（无操作数）
制表符\t和换行符\n
```

从文件中读取

```go
// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    counts := make(map[string]int)
    files := os.Args[1:]
    if len(files) == 0 {
        countLines(os.Stdin, counts)
    } else {
        for _, arg := range files {
            f, err := os.Open(arg)
            if err != nil {
                fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
                continue
            }
            countLines(f, counts)
            f.Close()
        }
    }
    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}
	
func countLines(f *os.File, counts map[string]int) {
    input := bufio.NewScanner(f)
    for input.Scan() {
        counts[input.Text()]++
    }
    // NOTE: ignoring potential errors from input.Err()
}

```

*`os.Open` 函数返回两个值。第一个值是被打开的文件（`*os.File`），其后被 `Scanner` 读取*。PS:使用逐行读取，**内存占用更低**（比 `ReadFile` 更安全）

*`os.Open` 返回的第二个值是内置 `error` 类型的值。如果 `err` 等于内置值`nil`（译注：相当于其它语言里的 `NULL`），那么文件被成功打开。*

*进入错误处理流程后，`continue` 语句直接跳到 `for` 循环的下个迭代开始执行。*

*函数和包级别的变量（package-level entities）可以任意顺序声明，并不影响其被调用。*

*`map` 是一个由 `make` 函数创建的数据结构的引用。`map` 作为参数传递给某函数时，该函数接收这个引用的一份拷贝（copy，或译为副本），被调用函数对 `map` 底层数据结构的任何修改，调用者函数都可以通过持有的 `map` 引用看到。在我们的例子中，`countLines` 函数向 `counts` 插入的值，也会被 `main` 函数看到。（译注：类似于 C++ 里的引用传递，实际上指针是另一个指针了，但内部存的值指向同一块内存）*



前两个例子基于“流”的模式读取数据，下面将全部数据一次性读入内存

```go
package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "strings"
)

func main() {
    counts := make(map[string]int)
    for _, filename := range os.Args[1:] {
        data, err := ioutil.ReadFile(filename)
        if err != nil {
            fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
            continue
        }
      //转换为string后按换行符分割
        for _, line := range strings.Split(string(data), "\n") {
            counts[line]++
        }
    }
    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}

```



*简化，只读指定文件，不读标准输入。其次，由于行计数代码只在一处用到，故将其移回 `main` 函数。*

*`ReadFile` 函数返回一个字节切片（byte slice），必须把它转换为 `string`，才能用 `strings.Split` 分割。*

*实现上，`bufio.Scanner`、`ioutil.ReadFile` 和 `ioutil.WriteFile` 都使用 `*os.File` 的 `Read` 和 `Write` 方法，但是，大多数程序员很少需要直接调用那些低级（lower-level）函数。高级（higher-level）函数，像 `bufio` 和 `io/ioutil` 包中所提供的那些，用起来要容易点。*

**练习 1.4：** 修改 `dup2`，出现重复的行时打印文件名称。

```go
// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// map[filename][line]count
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines("<stdin>", os.Stdin, counts)
	} else {
		for _, filename := range files {
			f, err := os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(filename, f, counts)
			f.Close()
		}
	}
	for filename, lineMap := range counts {
		for line, n := range lineMap {
			if n > 1 {
				fmt.Printf("%s\t%d\t%s\n", filename, n, line)
			}
		}
	}
}

func countLines(filename string, f *os.File, counts map[string]map[string]int) {
	// 安全初始化 map 的标准写法，用于处理嵌套 map 的情况
	if counts[filename] == nil {
		counts[filename] = make(map[string]int)
	}

	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[filename][line]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
```

思路：

首先要修改counts的类型为map[string]map[string]int，

其次修改countLines函数，使其接收三个参数，并在函数中增加安全初始化嵌套map的情况（`counts` 是一个 map，它的值本身又是一个 map。在第一次访问 `counts[filename]` 时，（第二层）默认是 `nil`，**不能直接赋值**，Go 会报运行时错误）

修改for循环，处理嵌套map的情况

#### 1.4. GIF动画

```go
// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
    "image"
    "image/color"
    "image/gif"
    "io"
    "math"
    "math/rand"
    "os"
    "time"
)
/*
color.Color 是 Go 标准库中的一个接口，用于表示“一个颜色”。
部分	含义
var palette	声明一个变量，名字叫 palette
[]color.Color	类型是：color.Color 接口的切片（即可以存多个颜色）
{color.White, color.Black}	切片的初始值是白色和黑色两个颜色
*/
var palette = []color.Color{color.White, color.Black}

const (
    whiteIndex = 0 // first color in palette
    blackIndex = 1 // next color in palette
)

func main() {
    // The sequence of images is deterministic unless we seed
    // the pseudo-random number generator using the current time.
    // Thanks to Randall McPherson for pointing out the omission.
    rand.Seed(time.Now().UTC().UnixNano())
    lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
    const (
        cycles  = 5     // number of complete x oscillator revolutions
        res     = 0.001 // angular resolution
        size    = 100   // image canvas covers [-size..+size]
        nframes = 64    // number of animation frames
        delay   = 8     // delay between frames in 10ms units
    )

    freq := rand.Float64() * 3.0 // relative frequency of y oscillator
    anim := gif.GIF{LoopCount: nframes}
    phase := 0.0 // phase difference
    for i := 0; i < nframes; i++ {
        rect := image.Rect(0, 0, 2*size+1, 2*size+1)
        img := image.NewPaletted(rect, palette)
        for t := 0.0; t < cycles*2*math.Pi; t += res {
            x := math.Sin(t)
            y := math.Sin(t*freq + phase)
            img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
                blackIndex)
        }
        phase += 0.1
        anim.Delay = append(anim.Delay, delay)
        anim.Image = append(anim.Image, img)
    }
    gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}


```

在包的导入路径包含多个单词时，如“image/color”，可以只用最后一个单词表示->color.White

在 Go 语言中，`const` 用于声明**常量**（constant），即在编译时就确定了值，且运行时不可更改。





**练习 1.5：** 修改前面的Lissajous程序里的调色板，由黑色改为绿色。我们可以用`color.RGBA{0xRR, 0xGG, 0xBB, 0xff}`来得到`#RRGGBB`这个色值，三个十六进制的字符串分别代表红、绿、蓝像素。

```
var palette = []color.Color{color.White, color.RGBA{0x00, 0xff, 0x00, 0xff}} // 绿色
```

**练习 1.6：** 修改Lissajous程序，修改其调色板来生成更丰富的颜色，然后修改SetColorIndex的第三个参数，看看显示结果吧。

```go
var palette = []color.Color{
    color.White,                   // index 0：背景色
    color.RGBA{255, 0, 0, 255},    // index 1：红
    color.RGBA{255, 165, 0, 255},  // index 2：橙
    color.RGBA{255, 255, 0, 255},  // index 3：黄
    color.RGBA{0, 255, 0, 255},    // index 4：绿
    color.RGBA{0, 0, 255, 255},    // index 5：蓝
}
.........
.........
const (
	whiteIndex = 0 // first color in palette
)
........
........
colorIndex := uint8(i%(len(palette)-1) + 1)

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex)
		}


```

首先定义了调色板包含一个背景色和五个颜色，这时 `len(palette) = 6`，其中：

- `palette[0]` 是背景白色
- `palette[1]`~`palette[5]` 是彩色线条颜色



colorIndex := uint8(i%(len(palette)-1) + 1)用于随机：

`len(palette) - 1 = 5`：去掉背景色，只用彩色部分

`i % 5`：确保帧编号在 0~4 之间循环（避免超出索引）

`+1`：跳过背景色 `palette[0]`，确保颜色索引从 1 开始

uint8：SetColorIndex需要的类型

#### 1.5.获取url

*为了最简单地展示基于HTTP获取信息的方式，下面给出一个示例程序fetch，这个程序将获取对应的url，并将其源文本打印出来；这个例子的灵感来源于curl工具。当然，curl提供的功能更为复杂丰富，这里只编写最简单的样例。这个样例之后还会多次被用到。*

```go
// Fetch prints the content found at a URL.
package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
)

func main() {
    for _, url := range os.Args[1:] {
        resp, err := http.Get(url)
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
            os.Exit(1)
        }
        b, err := io.ReadAll(resp.Body)
        resp.Body.Close()
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
            os.Exit(1)
        }
        fmt.Printf("%s", b)
    }
}

```



**练习 1.7：** 函数调用io.Copy(dst, src)会从src中读取内容，并将读到的结果写入到dst中，使用这个函数替代掉例子中的ioutil.ReadAll来拷贝响应结构体到os.Stdout，避免申请一个缓冲区（例子中的b）来存储。记得处理io.Copy返回结果中的错误。

```go
// Fetch prints the content found at a URL.
package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
)

func main() {
    for _, url := range os.Args[1:] {
        resp, err := http.Get(url)
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
            os.Exit(1)
        }
      
       defer resp.Body.Close()
      
        _, err = io.Copy(os.Stdout, resp.Body)
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: Copy faild %s: %v\n", url, err)
            os.Exit(1)
        }
    }
}

```

`defer resp.Body.Close()`：`defer` 是 Go 的延迟执行机制，在函数结束时自动执行。

`resp.Body` 是一个 `io.ReadCloser`，它是连接服务器的数据流通道。你**必须手动关闭它**，否则连接会一直占用内存和资源。

`io.Copy`有两个返回值：返回的字节数、err，我们只关注返回的错误，使用_忽略第一个返回值。

**练习 1.8：** 修改fetch这个范例，如果输入的url参数没有 `http://` 前缀的话，为这个url加上该前缀。你可能会用到strings.HasPrefix这个函数。

```
 if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
            url = "http://" + url
        }
```

`func HasPrefix(s, prefix string) bool` 判断s是否以prefix开头

**练习 1.9：** 修改fetch打印出HTTP协议的状态码，可以从resp.Status变量得到该状态码。

```
        //打印状态码，不需要在异常流程里打印，因为err时resp == nil
        fmt.Fprintf(os.Stdout, "status code: %s\n\n", resp.Status)
```

完整代码

```
// Fetch prints the content found at a URL.
package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "strings"
)

func main() {
    for _, url := range os.Args[1:] {
    	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
    		url = "http://" + url
    	}
        resp, err := http.Get(url)
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
            os.Exit(1)
        }
      
        defer resp.Body.Close()
        //打印状态码，不需要在异常流程里打印，因为err时resp == nil
        fmt.Fprintf(os.Stdout, "status code: %s\n\n", resp.Status)

        _, err = io.Copy(os.Stdout, resp.Body)
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: Copy faild %s: %v\n", url, err)
            os.Exit(1)
        }
        
    }
}

```

#### 1.6.并发获取多个url

*Go语言最有意思并且最新奇的特性就是对并发编程的支持。并发编程是一个大话题，在第八章和第九章中会专门讲到。这里我们只浅尝辄止地来体验一下Go语言里的goroutine和channel。*

```go
// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "os"
    "time"
)

func main() {
    start := time.Now()
    ch := make(chan string)
    for _, url := range os.Args[1:] {
        go fetch(url, ch) // start a goroutine
    }
    for range os.Args[1:] {
        fmt.Println(<-ch) // receive from channel ch
    }
    fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
    start := time.Now()
    resp, err := http.Get(url)
    if err != nil {
        ch <- fmt.Sprint(err) // send to channel ch
        return
    }
    nbytes, err := io.Copy(ioutil.Discard, resp.Body)
    resp.Body.Close() // don't leak resources
    if err != nil {
        ch <- fmt.Sprintf("while reading %s: %v", url, err)
        return
    }
    secs := time.Since(start).Seconds()
    ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

```

`go fetch(url, ch) // start a goroutine`

启动一个新的goroutine执行fetch(url, ch)函数，每个URL都在独立的线程中请求，main线程不会阻塞，会继续执行。

*goroutine是一种函数的并发执行方式，而channel是用来在goroutine之间进行参数传递。main函数本身也运行在一个goroutine中*

` ch := make(chan string)`

创建一个channel（通道），让goroutine把结果传回主goroutine

`ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)`

fetch函数把格式化好的字符串发送到通道ch，这个过程会**阻塞发送方**，直到有接收方（main 函数）从 `ch` 中读取。

    for range os.Args[1:] {
        fmt.Println(<-ch) // receive from channel ch
    }

主线程依次接收并打印

总体流程

```
main()
│
├── 创建 channel ch
│
├── 启动 N 个 goroutine 并发 fetch(url, ch)
│      └── 每个 fetch 执行 http.Get
│      └── 完成后通过 ch <- result 把结果发回主线程
│
├── 主线程通过 <-ch 读取每个结果（共 N 次）
│
└── 打印总耗时

```

**练习 1.10：** 找一个数据量比较大的网站，用本小节中的程序调研网站的缓存策略，对每个URL执行两遍请求，查看两次时间是否有较大的差别，并且每次获取到的响应内容是否一致，修改本节中的程序，将响应结果输出到文件，以便于进行对比。

```go
// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "time"
    "strings"
)

func main() {
    start := time.Now()
    ch := make(chan string)

    for _, url := range os.Args[1:] {
        go fetch(url, ch) // start a goroutine
    }
    for range os.Args[1:] {
        fmt.Println(<-ch) // receive from channel ch
    }

    fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<-- string) {
    start := time.Now()
    resp, err := http.Get(url)
    if err != nil {
        ch <- fmt.Sprint(err) // send to channel ch
        return
    }
    filename := sanitizeFilename(url) + ".html"
    file, err := os.Create(filename)
    if err != nil {
        ch <- fmt.Sprintf("create file error for %s: %v", url, err)
        return
    }
    defer file.Close()

    // 将 body 内容写入文件，同时统计大小
    nbytes, err := io.Copy(file, resp.Body)
    if err != nil {
        ch <- fmt.Sprintf("write error for %s: %v", url, err)
        return
    }

    secs := time.Since(start).Seconds()
    ch <- fmt.Sprintf("%.2fs  %7d  %s  -> saved to %s", secs, nbytes, url, filename)
}

// sanitizeFilename 将 URL 简化为合法文件名
func sanitizeFilename(url string) string {
    url = strings.TrimPrefix(url, "http://")
    url = strings.TrimPrefix(url, "https://")
    url = strings.ReplaceAll(url, "/", "_")
    return url
}

```

主要修改fetch函数，将Body内容写入文件

```go
❯ go run ./fetchall.go https://golang-china.github.io/gopl-zh/
0.39s    30164  https://golang-china.github.io/gopl-zh/  -> saved to golang-china.github.io_gopl-zh_.html
0.39s elapsed
❯ mv golang-china.github.io_gopl-zh_.html a.html
❯ go run ./fetchall.go https://golang-china.github.io/gopl-zh/
0.22s    30164  https://golang-china.github.io/gopl-zh/  -> saved to golang-china.github.io_gopl-zh_.html
0.22s elapsed
❯ diff a.html golang-china.github.io_gopl-zh_.html
```

**练习 1.11：** 在fetchall中尝试使用长一些的参数列表，比如使用在alexa.com的上百万网站里排名靠前的。如果一个网站没有回应，程序将采取怎样的行为？（Section8.9 描述了在这种情况下的应对机制）。

```go
 client := http.Client{
        Timeout: 5 * time.Second,
    }

    resp, err := client.Get(url)
```

创建一个客户端，设置超时时间、

限制并发数量、记录日志（后续再尝试）

---

Q：本小结用到的几个print有什么区别？

A：

| 函数      | 用途                      | 是否格式化 | 是否换行 | 是否返回字符串 |
| --------- | ------------------------- | ---------- | -------- | -------------- |
| `Sprintf` | 构造字符串，不打印        | ✅          | ❌        | ✅              |
| `Printf`  | 格式化并打印              | ✅          | ❌        | ❌              |
| `Println` | 直接打印（自动空格+换行） | ❌          | ✅        | ❌              |

---

#### 1.7.web服务

**练习 1.12：** 修改Lissajour服务，从URL读取变量，比如你可以访问 http://localhost:8000/?cycles=20 这个URL，这样访问可以将程序里的cycles默认的5修改为20。字符串转换为数字可以调用strconv.Atoi函数。你可以在godoc里查看strconv.Atoi的详细说明。

```go
// Server1 is a minimal "echo" server.
package main

import (
    "image"
    "image/color"
    "image/gif"
    "io"
    "math"
    "math/rand"
	  "net/http"
    "log"
    "fmt"
    "strconv"
)


var palette = []color.Color{
    color.White,                        // 背景色
    color.RGBA{0xFF, 0x00, 0x00, 0xFF}, // 红
    color.RGBA{0xFF, 0xA5, 0x00, 0xFF}, // 橙
    color.RGBA{0xFF, 0xFF, 0x00, 0xFF}, // 黄
    color.RGBA{0x00, 0xFF, 0x00, 0xFF}, // 绿
    color.RGBA{0x00, 0x00, 0xFF, 0xFF}, // 蓝
    color.RGBA{0x4B, 0x00, 0x82, 0xFF}, // 靛
    color.RGBA{0x8B, 0x00, 0xFF, 0xFF}, // 紫
}

const (
    whiteIndex = 0 // first color in palette
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    r.ParseForm() // 解析URL中的参数
    cycles := 5
    //获取cycel值
    if val := r.Form.Get("cycles"); val != "" {
    //字符串转换in t
            if parsed, err := strconv.Atoi(val); err == nil {
                cycles = parsed
            }
        }

    lissajous(w, cycles)
})
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//修改函数签名
func lissajous(out io.Writer, cycles int) {
    const (
        res     = 0.001 // angular resolution
        size    = 100   // image canvas covers [-size..+size]
        nframes = 64    // number of animation frames
        delay   = 8     // delay between frames in 10ms units
    )

    freq := rand.Float64() * 3.0 // relative frequency of y oscillator
    anim := gif.GIF{LoopCount: nframes}
    phase := 0.0 // phase difference
    for i := 0; i < nframes; i++ {
        rect := image.Rect(0, 0, 2*size+1, 2*size+1)
        img := image.NewPaletted(rect, palette)

        colorIndex := uint8(i%(len(palette)-1) + 1)
				//进行浮点数计算，转换cycles类型
        for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
            x := math.Sin(t)
            y := math.Sin(t*freq + phase)
            img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex)
        }
        phase += 0.1
        anim.Delay = append(anim.Delay, delay)
        anim.Image = append(anim.Image, img)
    }
    gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

```

#### 1.8.本章要点

*本章对Go语言做了一些介绍，Go语言很多方面在有限的篇幅中无法覆盖到。本节会把没有讲到的内容也做一些简单的介绍，这样读者在读到完整的内容之前，可以有个简单的印象。*

***控制流：** 在本章我们只介绍了if控制和for，但是没有提到switch多路选择。这里是一个简单的switch的例子：*

```
switch coinflip() {
case "heads":
    heads++
case "tails":
    tails++
default:
    fmt.Println("landed on edge!")
}
```

*在每一个函数之前写一个说明函数行为的注释也是一个好习惯。这些惯例很重要，因为这些内容会被像godoc这样的工具检测到，并且在执行命令时显示这些注释*

*多行注释可以用 `/* ... */` 来包裹，和其它大多数语言一样。*

### 2.程序结构

*Go语言和其他编程语言一样，一个大的程序是由很多小的基础构件组成的。变量保存值，简单的加法和减法运算被组合成较复杂的表达式。基础类型被聚合为数组或结构体等更复杂的数据结构。然后使用if和for之类的控制语句来组织和控制表达式的执行流程。然后多个语句被组织到一个个函数中，以便代码的隔离和复用。函数以源文件和包的方式被组织。*

*我们已经在前面章节的例子中看到了很多例子。在本章中，我们将深入讨论Go程序基础结构方面的一些细节。每个示例程序都是刻意写的简单，这样我们可以减少复杂的算法或数据结构等不相关的问题带来的干扰，从而可以专注于Go语言本身的学习。*

#### 2.1.命名

*Go语言中的函数名、变量名、常量名、类型名、语句标号和包名等所有的命名，都遵循一个简单的命名规则：一个名字必须以一个字母（Unicode字母）或下划线开头，后面可以跟任意数量的字母、数字或下划线。大写字母和小写字母是不同的：heapSort和Heapsort是两个不同的名字。*

类似if和switch的关键字有25个；关键字不能用于自定义名字，只能在特定语法结构中使用。

```
break      default       func     interface   select
case       defer         go       map         struct
chan       else          goto     package     switch
const      fallthrough   if       range       type
continue   for           import   return      var
```

此外，还有大约30多个预定义的名字，比如int和true等，主要对应内建的常量、类型和函数。

```
内建常量: true false iota nil

内建类型: int int8 int16 int32 int64
          uint uint8 uint16 uint32 uint64 uintptr
          float32 float64 complex128 complex64
          bool byte rune string error

内建函数: make len cap new append copy close delete
          complex real imag
          panic recover
```

*这些内部预先定义的名字并不是关键字，你可以在定义中重新使用它们。*



*如果一个名字是在函数**内部定义**，那么它就只在函数内部有效。如果是在函数**外部定义**，那么将在当前包的所有文件中都可以访问。名字的**开头字母的大小写决定了名字在包外的可见性**。**例如fmt包的Printf函数就是导出的，可以在fmt包外部访问。包本身的名字一般总是用小写字母*。

导出举例：

```go
// 在 mypkg 包中定义的内容
package mypkg

// 导出类型（外部可见）
type Person struct {
    Name string // 导出字段
    age  int    // 未导出字段（只能在 mypkg 内访问）
}

// 导出函数
func SayHello() {
    // ...
}

// 未导出函数（小写字母开头，只能在 mypkg 包内部使用）
func secretFunction() {
    // ...
}

```

```go
import "mypkg"

func main() {
    p := mypkg.Person{}  // ✅ 可以访问导出的类型
    p.Name = "Alice"     // ✅ 可以访问导出的字段
    // p.age = 30         // ❌ 编译错误，age 是未导出的字段
    mypkg.SayHello()     // ✅ 导出的函数
    // mypkg.secretFunction() // ❌ 编译错误，未导出的函数
}

```



*尽量使用短小的名字，对于局部变量尤其是这样*

*如果一个名字的作用域比较大，生命周期也比较长，那么用长的名字将会更有意义。*

*推荐使用 **驼峰式** 命名，而像ASCII和HTML这样的缩略词则避免使用大小写混合的写法，它们可能被称为htmlEscape、HTMLEscape或escapeHTML，但不会是escapeHtml。*

#### 2.2.声明

Go语言主要有四种类型的声明语句：var、const、type和func，分别对应变量、常量、类型和函数实体对象的声明。

*一个Go语言编写的程序对应一个或多个以.go为文件后缀名的源文件。每个源文件中以**包的声明语句开始**，说明该源文件是属于哪个包。包声明语句之后是i**mport语句导入依赖的其它包**，然后是**包一级的类型**、**变量、常量、函数的声明语句**，**包一级的各种类型的声明语句的顺序无关紧要*

```
// Boiling prints the boiling point of water.
package main

import "fmt"

const boilingF = 212.0

func main() {
    var f = boilingF
    var c = (f - 32) * 5 / 9
    fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
    // Output:
    // boiling point = 212°F or 100°C
}

```

*其中**常量boilingF是在包一级范围声明**语句声明的，然后**f和c两个变量是在main函数内部声明**的声明语句声明的。在包一级声明语句声明的名字可在整个包对应的每个源文件中访问，而不是仅仅在其声明语句所在的源文件中访问。*

例如：

a.go

```
// Boiling prints the boiling point of water.
package main

import "fmt"

const boilingF = 212

func main() {
		fmt.Println("Boiling point in Fahrenheit:", boilingF)
    printCelsius()

}
```

b.go

```
package main

import "fmt"

func printCelsius() {
		c := (boilingF - 32) * 5 / 9 // 可以访问boilingF
    fmt.Println("Boiling point in Celsius:", c)
}
```

执行

```
❯ go mod init example.com/myapp

go: creating new go.mod: module example.com/myapp
go: to add module requirements and sums:
        go mod tidy
        
❯ go run .
Boiling point in Fahrenheit: 212
Boiling point in Celsius: 100
```



函数声明的格式

```
func 函数名(参数列表) 返回值类型 {
    // 函数体
}

如
func fToC(f float64) float64 {
    return (f - 32) * 5 / 9
}
```

一个函数的声明由一个函数名字、参数列表、一个可选的返回值列表和包含函数定义的函数体组成。如果函数没有返回值，那么返回值列表是省略的。执行函数从函数的第一个语句开始，依次顺序执行直到遇到return返回语句，如果没有返回语句则是执行到函数末尾，然后返回到函数调用者。

#### 2.3.变量

```
var 变量名字 类型 = 表达式
```

*零值初始化：数值类型变量对应的零值是0，布尔类型变量对应的零值是false，字符串类型对应的零值是空字符串，接口或引用类型（包括slice、指针、map、chan和函数）变量对应的零值是nil。数组或结构体等聚合类型对应的零值是每个元素或字段都是对应该类型的零值。*

可以在一个声明语句中同时声明一组变量

```
var i, j, k int                 // int, int, int
var b, f, s = true, 2.3, "four" // bool, float64, string
```

一组变量也可以通过调用一个函数，由函数返回的多个返回值初始化：

```Go
var f, err = os.Open(name) // os.Open returns a file and an error
```

##### 2.3.1.简短变量声明

*以“名字 := 表达式”形式声明变量，变量的类型根据表达式来自动推导。*

*var形式的声明语句往往是用于需要显式指定变量类型的地方，或者因为变量稍后会被重新赋值而初始值无关紧要的地方*

```
i := 100                  // an int
var boiling float64 = 100 // a float64
var names []string
var err error
var p Point
```

*简短变量声明语句也可以用来声明和初始化一组变量：*

```go
i, j := 0, 1
```

*这种同时声明多个变量的方式应该限制只在可以提高代码可读性的地方使用，比如for语句的循环的初始化语句部分。*

简短变量声明语句中必须至少要声明一个新的变量，下面的代码将不能编译通过：

```Go
f, err := os.Open(infile)
// ...
f, err := os.Create(outfile) // compile error: no new variables
```

##### 2.3.2. 指针

如果用“var x int”声明语句声明一个x变量，那么&x表达式（取x变量的内存地址）将产生一个指向该整数变量的指针，指针对应的数据类型是`*int`，指针被称之为“指向int类型的指针”。如果指针名字为p，那么可以说“p指针指向变量x”，或者说“p指针保存了x变量的内存地址”。同时`*p`表达式对应p指针指向的变量的值。一般`*p`表达式读取指针指向的变量的值，这里为int类型的值，同时因为`*p`对应一个变量，所以该表达式也可以出现在赋值语句的左边，表示更新指针所指向的变量的值。

```Go
x := 1
p := &x         // p, of type *int, points to x
fmt.Println(*p) // "1"
*p = 2          // equivalent to x = 2
fmt.Println(x)  // "2"
```

##### 2.3.3. new函数

另一个创建变量的方法是调用内建的new函数。表达式new(T)将创建一个T类型的匿名变量，初始化为T类型的零值，然后返回变量地址，返回的指针类型为`*T`。

```Go
p := new(int)   // p, *int 类型, 指向匿名的 int 变量
fmt.Println(*p) // "0"
*p = 2          // 设置 int 匿名变量的值为 2
fmt.Println(*p) // "2"
```

##### 2.3.4. 变量的生命周期

*变量的生命周期指的是在程序运行期间变量有效存在的时间段。对于在包一级声明的变量来说，它们的生命周期和整个程序的运行周期是一致的。而相比之下，局部变量的生命周期则是动态的：每次从创建一个新变量的声明语句开始，直到该变量不再被引用为止，然后变量的存储空间可能被回收。函数的参数变量和返回值变量都是局部变量。它们在函数每次被调用的时候创建。*

*一个变量的有效周期只取决于是否可达，一个循环迭代内部的局部变量的生命周期可能超出其局部作用域。同时，局部变量可能在函数返回之后依然存在。*

*Go编译器会自动选择在栈上还是在堆上分配局部变量的存储空间，*

🔹**1. 栈和堆的区别**

- **栈（stack）**：快速分配/释放，但生命周期仅限于函数调用期间。
- **堆（heap）**：内存持续时间更长，但分配/回收较慢，需要垃圾回收器（GC）处理。

------

🔹**2. `var` 和 `new` 并不决定内存位置**

很多语言（比如 C/C++）中，你用 `new` 就会在堆上分配，用局部变量就会在栈上分配。但 **Go 不这样** —— 它通过编译器分析代码来决定变量到底是在堆上还是栈上分配。

这意味着：
 ✅ 你写 `var x int`，它**可能在堆上**
 ✅ 你写 `y := new(int)`，它**可能在栈上**

这个决定 **不由你代码的语法决定，而是由编译器对变量是否“逃逸”进行判断决定的**。

------

🔹**3. 什么是“逃逸”？**

> “逃逸”指的是一个变量在函数执行完毕后，依然可能被函数外部的代码访问。

如果一个变量在函数返回之后还要继续用，**它就不能放在栈上**，因为函数返回时，栈空间就会被释放。这样就必须把它放在堆上。

------

🔹**具体例子分析：**

🧩 示例 1：

```
go


复制编辑
var global *int

func f() {
    var x int
    x = 1
    global = &x
}
```

- `x` 是 `f` 函数内部的变量。
- 但我们把 `x` 的地址 `&x` 存到了全局变量 `global` 中。
- 所以 **即使函数 `f` 返回了，`global` 依然可以访问 `x` 的值**。
- 所以 `x` 的生命周期超出了函数作用域，也就是“逃逸”了。
- 编译器就会 **把 `x` 放到堆上**。

------

🧩 示例 2：

```
go


复制编辑
func g() {
    y := new(int)
    *y = 1
}
```

- 这里用了 `new(int)`，看起来是“堆分配”，但这并不是关键。
- `y` 是局部变量，`*y` 也只在函数 `g` 中使用。
- `g` 返回后，没人再用这个变量了，它不会“逃逸”。
- 所以编译器可能 **优化成栈上分配**，而不是放到堆上。

#### 2.4.赋值

使用赋值语句可以更新一个变量的值，最简单的赋值语句是将要被赋值的变量放在=的左边，新值的表达式放在=的右边。

```Go
x = 1                       // 命名变量的赋值
*p = true                   // 通过指针间接赋值
person.name = "bob"         // 结构体字段赋值
count[x] = count[x] * scale 或 count[x] *= scale
// 数组、slice或map的元素赋值
```

数值变量也可以支持`++`递增和`--`递减语句（译注：自增和自减是语句，而不是表达式

```
v := 1
v++    // 等价方式 v = v + 1；v 变成 2
v--    // 等价方式 v = v - 1；v 变成 1
```

##### 2.4.1. 元组赋值

元组赋值是另一种形式的赋值语句，它允许同时更新多个变量的值。例如我们可以这样交换两个变量的值：

```go
x, y = y, x

a[i], a[j] = a[j], a[i]
```

有些表达式会产生多个值，比如调用一个有多个返回值的函数。

```go
f, err = os.Open("foo.txt") // function call returns two values
```

和变量声明一样，我们可以用下划线空白标识符`_`来丢弃不需要的值。

```go
_, err = io.Copy(dst, src) // 丢弃字节数
_, ok = x.(T)              // 只检测类型，忽略具体值
```

#### 2.5.类型

*在任何程序中都会存在一些变量有着相同的内部结构，但是却表示完全不同的概念。例如，一个int类型的变量可以用来表示一个循环的迭代索引、或者一个时间戳、或者一个文件描述符、或者一个月份；一个float64类型的变量可以用来表示每秒移动几米的速度、或者是不同温度单位下的温度；一个字符串可以用来表示一个密码或者一个颜色的名称。*

```
type 类型名字 底层类型
```



```go
// Package tempconv performs Celsius and Fahrenheit temperature computations.
package tempconv

import "fmt"

//定义了两个类型，底层都是float64，它们是不同的数据类型，因此它们不可以被相互比较或混在一个表达式运算
type Celsius float64    // 摄氏温度
type Fahrenheit float64 // 华氏温度

const (
    AbsoluteZeroC Celsius = -273.15 // 绝对零度
    FreezingC     Celsius = 0       // 结冰点温度
    BoilingC      Celsius = 100     // 沸水温度
)

/*Celsius(t)和Fahrenheit(t)是类型转换操作，它们并不是函数调用，类型转换不会改变值本身，但是会使它们的语义发生变化。
CToF和FToC两个函数则是对不同温度单位下的温度进行换算，它们会返回不同的值。*/
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

```



*对于每一个类型T，都有一个对应的类型转换操作T(x)，用于将x转为T类型（译注：如果T是指针类型，可能会需要用小括弧包装T，比如`(*int)(0)`）。只有当两个类型的底层基础类型相同时，才允许这种转型操作，或者是两者都是指向相同底层结构的指针类型，这些转换只改变类型而不会影响值本身。*



下面的声明语句，Celsius类型的参数c出现在了函数名的前面，表示声明的是Celsius类型的一个名叫String的方法，该方法返回该类型对象c带着°C温度单位的字符串：

```Go
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }
```

#### 2.6.包和文件



**练习 2.1：** 写一个通用的单位转换程序，用类似cf程序的方式从命令行读取参数，如果缺省的话则是从标准输入读取参数，然后做类似Celsius和Fahrenheit的单位转换，长度单位可以对应英尺和米，重量单位可以对应磅和公斤等。

目录结构：

```
❯ tree
.
├── go.mod
├── main.go
└── tempconv
    ├── conv.go
    ├── go.mod
    └── tempconv.go
```

```
go mod init local/tempconv
```

go.mod

```
module local/myapp

go 1.24.4

replace local/tempconv => ./tempconv

require local/tempconv v0.0.0-00010101000000-000000000000 // indirect
```

main.go

```
package main

import (
	"fmt"
	"local/tempconv" // 注意这里的模块名必须与你 go.mod 中的一致
)

func main() {
	k := tempconv.Kelvin(273.15)
	c := tempconv.KToC(k)

	fmt.Println("k:", k)       // 输出：273.15K
	fmt.Println("k to c :", c) // 输出：0°C
}

```

conv.go

```go
package tempconv

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func CToK(c Celsius) Kelvin { return Kelvin(c + 273.15) }

func KToC(k Kelvin) Celsius { return Celsius(k - 273.15) }

```

tempconv.go

```go
// Package tempconv performs Celsius and Fahrenheit conversions.
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
	AbsoluteZeroK Kelvin  = 0
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%gK", k) }

```

**练习 2.2：** 写一个通用的单位转换程序，用类似cf程序的方式从命令行读取参数，如果缺省的话则是从标准输入读取参数，然后做类似Celsius和Fahrenheit的单位转换，长度单位可以对应英尺和米，重量单位可以对应磅和公斤等。

```go
// 从输入参数进行重量转换
package main

import (
	"fmt"
	"os"
	"strconv"
)

type ounce float64
type gram float64

func (o ounce) String() string { return fmt.Sprintf("%g oz", o) }
func (g gram) String() string  { return fmt.Sprintf("%g g", g) }

func OToG(o ounce) gram { return gram(o * 28.3495) }

func GToO(g gram) ounce { return ounce(g / 28.3495) }

func main() {
	for _, args := range os.Args[1:] {
		t, err := strconv.ParseFloat(args, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "weightconv: %v\n", err)
			os.Exit(1)
		}
		o := ounce(t)
		g := OToG(o)
		fmt.Printf("%s = %s\n", o, g)
		g2 := GToO(g)
		fmt.Printf("%s = %s\n", g, g2)
	}
}

```

