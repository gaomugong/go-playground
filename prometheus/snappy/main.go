package main

import (
	"fmt"

	"github.com/golang/snappy"
)

// 简介
// 该包的目标并不是最大化的压缩比例，也不是和其他压缩库兼容；相反，snappy算法的目标是在合理的压缩率下尽可能的提高压缩速度。
//
// 例如，与zlib的最快压缩模式相比，snappy依然比其快了一个数量级，但产生的压缩文件要比zip的大20%到100%。
//
// 特性
// snappy压缩算法具有以下特性：
//
// 快速：压缩速度大概在250MB/秒及更快的速度进行压缩。
// 稳定：在过去的几年中，Snappy在Google的生产环境中压缩并解压缩了数P字节（petabytes）的数据。Snappy位流格式是稳定的，不会在版本之间发生变化
// 健壮性：Snappy解压缩器设计为不会因遇到损坏或恶意输入而崩溃
// 性能
// Snappy的目标是快速。在64位模式下，一个Corei7处理器的单核上，其压缩速度约为250MB/秒或更快，解压缩速度约为500MB/秒或更快。
// （这些数字是在我们的基准测试套件中最慢的输入情况下得出的；其他输入会快得多。）在我们的测试中，
// Snappy通常比同一级别的算法（如LZO、LZF、QuickLZ等）更快，同时实现了类似的压缩率。

func main() {
	src := `{"Ag(T+D)":{"instID":"Ag(T+D)","name":"test","last":"4141","upDown":"21","upDownRate":"0.51","quoteDate":"20170328","quoteTime":"22:34:29"},"Au(T+D)":{"instID":"Au(T+D)","name":"黄金延期","last":"280.55","upDown":"0.88","upDownRate":"0.31","quoteDate":"20170328","quoteTime":"22:34:15"},"mAu(T+D)":{"instID":"mAu(T+D)","name":"Mini黄金延期","last":"280.5","upDown":"0.7","upDownRate":"0.25","quoteDate":"20170328","quoteTime":"22:34:10"}}`
	dst := snappy.Encode(nil, []byte(src))
	// before compression len:446
	// after compression len:220
	fmt.Printf("before compression len:%d\n", len(src))
	fmt.Printf("after compression len:%d\n", len(dst))

	src1, _ := snappy.Decode(nil, dst)
	// uncompressed len: 446
	fmt.Println("uncompressed len:", len(src1))
}
