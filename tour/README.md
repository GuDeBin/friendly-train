# Tour of Go

[Tour of Go](https://go.dev/learn/)

## 这是什么

这是 go 官网上的一个教程，将 go 基础知识讲解一遍，确实很全很有用，这里把知识点记录下

## 包

每个 Go 程序都是由包构成的。

程序从 main 包开始运行。

本程序通过导入路径 "fmt" 和 "math/rand" 来使用这两个包。

按照约定，包名与导入路径的最后一个元素一致。例如，"math/rand" 包中的源码均以 package rand 语句开始。

_注意：_ 此程序的运行环境是固定的，因此 rand.Intn 总是会返回相同的数字。 （要得到不同的数字，需为生成器提供不同的种子数，参见 rand.Seed。 练习场中的时间为常量，因此你需要用其它的值作为种子数。）

代码示例

```go
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}

```

## 导入

此代码用圆括号组合了导入，这是“分组”形式的导入语句。

当然你也可以编写多个导入语句，例如：

```go
import "fmt"
import "math"
```

不过使用分组导入语句是更好的形式。

示例代码

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}

```

## 导出名

在 Go 中，如果一个名字以大写字母开头，那么它就是已导出的。例如，Pizza 就是个已导出名，Pi 也同样，它导出自 math 包。

pizza 和 pi 并未以大写字母开头，所以它们是未导出的。

在导入一个包时，你只能引用其中已导出的名字。任何“未导出”的名字在该包外均无法访问。

执行代码，观察错误输出。

然后将 math.pi 改名为 math.Pi 再试着执行一次。

示例代码

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi)
}
```
