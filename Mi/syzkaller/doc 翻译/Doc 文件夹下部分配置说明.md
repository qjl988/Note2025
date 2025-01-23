# db
> https://github.com/google/syzkaller/blob/master/docs/db.md
# 常用操作
**参数说明**

```txt
  -arch string
    	target arch
  -os string
    	target OS
  -version uint
    	database version
  -vv int
    	verbosity
```

```bash
>./bin/syz-db -os=linux -arch=arm64 bench ./workdir/corpus.db                                                                        
>allocs 778 MB (1 M), next GC 118 MB, sys heap 158 MB, live allocs 8 MB (0 M), time 200.564668ms
                       
corpus size: 826

program size: min=1 avg=2 max=30 10%=1 50%=2 90%=4
```
**输出说明**
- **allocs**: 778 MB，是程序分配的内存总量，括号中的 (1 M) 表示该内存分配中涉及了 1 百万个分配操作。
- **next GC**: 118 MB，是触发下一次垃圾回收（Garbage Collection）所需的阈值。
- **sys heap**: 158 MB，指系统分配的堆内存的大小。
- **live allocs**: 8 MB (0 M)，表示当前存活的内存分配量，总内存大小为 8 MB，分配操作数为 0 百万次。
- **time**: 200.564668ms，运行该基准测试所花费的时间。
- **corpus size**: 数据库中的样本总数是 826 个。这些样本可能代表着一组测试用例或执行的系统调用集。
- **min**: 数据库中最小程序大小是 1（可能是表示系统调用的最少步骤数）。

- **avg**: 平均程序大小为 2。

- **max**: 最大程序大小为 30。

- **百分位分布**:
  - **10%**: 10% 的程序大小是 1。
  - **50%**: 一半（中位数）的程序大小是 2。
  - **90%**: 90% 的程序大小是 4。


这段输出是 syz-db 运行基准测试（bench）命令的结果，主要用于：

1. **评估数据库的性能**：通过内存分配和垃圾回收信息来衡量效率。

2. **分析语料库分布**：检查存储的测试用例（系统调用程序）的数量和复杂度。


此信息有助于开发者优化 syzkaller 的测试用例生成和管理策略，尤其是在资源分配和语料覆盖方面。



