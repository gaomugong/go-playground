// Package http包主要是讲http.Handler和http.HandleFunc的实现抽取出来
// 展示了在函数类型上实现接口的一种用途，核心在于 AddRunner 时，可以直接对
// 业务函数 f 作 RunnerFunc(f) 的转换，业务侧只需要实现 RunnerFunc 类型的
// 方法，而无需关注 Runner 接口的实现
package http

import (
	"context"
	"log"
)

// Input Run参数
type Input struct{}

// Result Run结果
type Result struct{}

// Runner 只有一个Run方法的接口
type Runner interface {
	Run(ctx context.Context, input *Input, result *Result)
}

// RunnerFunc 定义一个符合上述接口方法定义的函数类型，并使其实现该接口
type RunnerFunc func(ctx context.Context, input *Input, result *Result)

// Run 是函数类型RunnerFunc实现 Runner, 逻辑上只是执行其本身，即在闭包中执行 r
func (r RunnerFunc) Run(ctx context.Context, input *Input, result *Result) {
	r(ctx, input, result)
}

//type RunnerWithInputResult struct {
//	Runner
//	Input
//	Result
//}

// Task 定义一个任务对象，负责执行预注册好的批处理任务
type Task struct {
	name string
	// TODO: 这里可以进一步优化，比如：
	//runners map[string]RunnerWithInputResult
	runners map[string]Runner
}

// AddRunner 添加子任务
func (t *Task) AddRunner(name string, runner Runner) {
	t.runners[name] = runner
}

// Run 批量执行任务
func (t *Task) Run() {
	// TODO: 这里可以简化Run的定义，只传入context，参数在 AddRunner 时就加入到 runners 中
	for name, runner := range t.runners {
		log.Printf("[%s] started\n", name)
		runner.Run(context.Background(), &Input{}, &Result{})
	}
}
