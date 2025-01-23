# Program syntax

> https://github.com/google/syzkaller/blob/master/docs/program_syntax.md

Syzkaller 使用一种紧凑的领域特定语言（DSL）来记录执行的程序、测试其代码以及在语料库中持久化程序。本文简要描述了相关语法。一些有用的信息也可以在[现有示例]()和程序[反序列化代码]()中找到。

结合执行选项，该 DSL 提供了 syz-executor 运行程序所需的一切。

例如，以下是一个程序：

```c
r0 = syz_open_dev$loop(&(0x7f00000011c0), 0x0, 0x0)

r1 = openat$6lowpan_control(0xffffffffffffff9c, &(0x7f00000000c0), 0x2, 0x0)

ioctl$LOOP_SET_FD(r0, 0x4c00, r1)
```

每行描述一个特定的系统调用，其中前两个调用将结果保存到临时变量 r0 和 r1 中，这些变量被传递给第三个调用。

```c
line = assignment | call

assignment = variable " = " call

call = syscall-name "(" [arg ["," arg]*] ")" ["(" [call-prop ["," call-prop*] ")"]

arg = "nil" | "AUTO" | const-arg | resource-arg | result-arg | pointer-arg | string-arg | struct-arg | array-arg | union-arg

const-arg = "0x" hex-integer

resource-arg = variable ["/" hex-integer] ["+" hex-integer]

result-arg = "<" variable "=>" arg

pointer-arg = "&" pointer-arg-addr ["=ANY"] "=" arg

pointer-arg-addr = "AUTO" | "(" pointer-addr ["/" region-size] ")"

string-arg = "'" escaped-string "'" | "\"" escaped-string "\"" | "\"$" escaped-string "\""

struct-arg = "{" [arg ["," arg]*] "}"

array-arg = "[" [arg ["," arg]*] "]"

union-arg = "@" field-name ["=" arg]

call-prop = prop-name ": " prop-value

variable = "r" dec-integer

pointer-addr = hex-integer

region-size = hex-integer
```

程序中也可以包含空行和注释。

```
# 获取一个文件句柄
r0 = openat(0xffffffffffffff9c, &AUTO='./file1\x00', 0x42, 0x1ff)
# 执行写入操作
write(r0, &AUTO="01010101", 0x4)
```

**内存管理**

内存管理由 Syzkaller 自行完成。它会分配必要大小的虚拟内存区域，并设置指针参数的最终值。

通过使用 AUTO 关键字，程序可以将数据存储的全部控制权交给 Syzkaller。例如，当参数必须通过引用传递，但其值的确切位置并不重要时，这种方式非常方便。

```
r1 = syz_genetlink_get_family_id$nl80211(&AUTO='nl80211\x00', 0xffffffffffffffff)

ioctl$sock_SIOCGIFINDEX_80211(r0, 0x8933, &AUTO={'wlan0\x00', <r2=>0x0})
```

此外，某些数据可以“锚定”到特定地址。这在需要多个调用共享同一区域的情况下尤为重要。在这种情况下，指针地址必须设置为 0x7f0000000000 偏移量。实际执行前，Syzkaller 会将指针调整到实际 mmap 区域的起始位置。

**调用属性**

调用属性指定有关如何执行特定调用的额外信息。程序中的每个调用都有自己的一组调用属性。如果未提供属性，Syzkaller 会采用默认值。

目前，Syzkaller 支持以下调用属性。

**故障注入**

语法：fail_nth: N。

接受一个整数（十进制）参数 N。如果参数为非负值，将在第 N 次调用时注入故障。

```
r0 = openat$6lowpan_control(0xffffffffffffff9c, &(0x7f00000000c0), 0x2, 0x0)

ioctl$LOOP_SET_FD(r0, 0x4c00, r0) (fail_nth: 5)
```

**异步执行**

语法：async。

指示 syz-executor 不等待调用完成，而是立即继续下一个调用。

```
r0 = openat(0xffffffffffffff9c, &AUTO='./file1\x00', 0x42, 0x1ff)

write(r0, &AUTO="01010101", 0x4) (async)

read(r0, &AUTO=""/4, 0x4)

close(r0)
```

设置 async 标志时需注意以下几点：

1. 这些程序应仅在线程模式下执行（即必须为 syz-executor 传递 -threaded 标志）。

2. 每个 async 调用在单独的线程中执行，并且线程数量有限（kMaxThreads = 16）。

3. 如果一个 async 调用生成了资源，需注意某些其他调用可能会以其为输入；如果资源生成调用尚未完成，syz-executor 将仅传递 0。