## 简述一下GMP模型

https://zhuanlan.zhihu.com/p/261057034

### 来自混元的答案：

Golang 的 GPM 模型是指 Goroutine（协程）、Processor（处理器）和 Machine（机器）三个元素组成的模型。
这个模型是为了实现高并发、高性能的程序设计而提出的。下面详细解释这三个元素：

Goroutine（协程）：Goroutine 是 Golang 中并发执行的基本单位，它是一种轻量级的线程。每个 Goroutine 都有自己的栈空间和执行函数。
Goroutine 之间的切换和调度成本比操作系统线程要低得多。使用关键字 go 可以创建一个新的 Goroutine。
Processor（处理器）：Processor 是 Golang 调度器中负责调度 Goroutine 的执行的元素。
它负责从全局的 Goroutine 队列中取出 Goroutine 分配给各个 Machine 执行。
Processor 的数量是可以配置的，默认情况下，Golang 会根据系统的 CPU 核心数来设置 Processor 的数量。
Machine（机器）：Machine 是指真正执行 Goroutine 的实体，它与操作系统线程密切相关。
每个 Machine 都会关联一个操作系统线程，并在该线程上执行分配给它的 Goroutine。
Machine 的数量是由操作系统决定的，通常与系统的 CPU 核心数相同。

### 来自GPT的答案：

GPM 模型是 Go 语言中的一种协程调度模型，它主要包含三个部分：

G：G（Goroutine）是 Go 语言中的协程，它类似于线程，但是比线程更轻量级。每个 G 都有一个对应的栈，以及一些与调度相关的信息。
P：P（Processor）是处理器的缩写，它是协程的执行者。每个 P 都有一个对应的本地队列（Local Run Queue），用于存储等待执行的 G。P 的个数是固定的，默认情况下等于 CPU 核心数。
M：M（Machine）是线程的缩写，它是操作系统调度的基本单位。每个 M 都有一个对应的 G，用于执行 Go 代码。M 与 P 的关系是一对多的关系，即一个 M 可以对应多个 P，但一个 P 只能对应一个 M。

GPM 模型的基本工作流程如下：

当一个 G 准备好之后，它会被放入一个全局队列（Global Run Queue）中。
当一个 P 空闲时，它会从全局队列中获取一个 G，并将其放入本地队列中。
当一个 P 的本地队列为空时，它会从其他 P 的本地队列中窃取一半的 G，以保证负载均衡。
当一个 G 阻塞时，它会被放入一个对应的等待队列中。
当一个 G 被唤醒时，它会被放回全局队列中，等待被调度。