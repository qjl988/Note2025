**系统调用描述**

syzkaller 使用声明式描述的系统调用接口来操作程序（即一系列系统调用）。以下是描述的一个（希望是自解释的）示例摘录：

```c
open(file filename, flags flags[open_flags], mode flags[open_mode]) fd

read(fd fd, buf buffer[out], count len[buf])

close(fd fd)

open_mode = S_IRUSR, S_IWUSR, S_IXUSR, S_IRGRP, S_IWGRP, S_IXGRP, S_IROTH, S_IWOTH, S_IXOTH
```

这些描述存储在 `sys/$OS/*.txt` 文件中。例如，可以参阅 [sys/linux/dev_snd_midi.txt](https://github.com/google/syzkaller/blob/master/sys/linux/dev_snd_midi.txt) 文件了解 Linux MIDI 接口的描述。

关于描述语法的更正式说明可以在 [这里](syscall_descriptions_syntax.md) 找到。

**程序**

翻译后的描述被用于生成、变异、执行、最小化、序列化和反序列化程序。程序是一个包含具体参数值的系统调用序列。以下是一个程序的文本表示示例：

```c
r0 = open(&(0x7f0000000000)="./file0", 0x3, 0x9)

read(r0, &(0x7f0000000000), 42)

close(r0)
```

在实际操作中，syzkaller 使用包含 Call 和 Arg 值的内存中类似抽象语法树（AST）的表示形式，定义于 [prog/prog.go](https://github.com/google/syzkaller/blob/master/prog/prog.go)。该表示形式被用于[分析](https://github.com/google/syzkaller/blob/master/prog/analysis.go)、[生成](https://github.com/google/syzkaller/blob/master/prog/rand.go)、[变异](https://github.com/google/syzkaller/blob/master/prog/mutation.go)、[最小化](https://github.com/google/syzkaller/blob/master/prog/minimization.go)、[验证](https://github.com/google/syzkaller/blob/master/prog/validation.go) 等操作。

这种内存表示可以被[转换](https://github.com/google/syzkaller/blob/master/prog/encoding.go)为文本形式以便存储到磁盘中的语料库、供人类查看等。

还有另一种程序的[二进制表示](https://github.com/google/syzkaller/blob/master/prog/decodeexec.go)（称为 exec），这种表示更简单，不包含丰富的类型信息（不可逆），用于程序通过 [executor](https://github.com/google/syzkaller/blob/master/executor/executor.cc) 实际执行（解释）。

**描述新的系统调用**

本节说明如何扩展 syzkaller 以允许模糊测试更多的内核接口。这对于提出新系统调用的内核开发者尤其有用。

**当前流程**

目前所有的系统调用描述都是手动编写的。有一个[公开的问题](https://github.com/google/syzkaller/issues/590) 提供了对此过程的部分帮助并正在进行工作，但我们尚未实现完全自动化生成描述的方法。

一个辅助工具 [headerparser](headerparser_usage.md) 可以从头文件中自动生成部分描述。Visual Studio Code 还提供了 [syz-lang 扩展](https://marketplace.visualstudio.com/items?itemName=AndreyArtemiev.syzlang-extension&ssr=false#overview)，用于语法高亮。

**开始步骤**

要启用对新的内核接口的模糊测试：

**研究接口**

找出使用该接口所需的系统调用。有时除了源代码外没有其他资源，但以下方法可能会有所帮助：

- 搜索互联网有关接口名称或某些唯一常量的信息。

- 在内核的 Documentation/ 目录中查找。

- 搜索内核的 tools/testing/ 目录。

- 查看源代码中的大型注释块。
  
- 使用 git blame 或 git log 查找添加该接口的提交并阅读提交描述。
  
- 阅读已知使用该接口的库或应用程序的源代码或跟踪它们。

**添加描述**

使用[语法文档](syscall_descriptions_syntax.md)和[现有描述]()作为示例，将接口的声明式描述添加到适当的文件中：

- sys/linux/<subsystem>.txt 文件包含特定内核子系统的系统调用，例如 [bpf.txt]() 或 [socket.txt]()。
- [sys/linux/sys.txt]() 包含更通用的系统调用描述。
- 一个全新的子系统可以作为一个新的 sys/linux/<new>.txt 文件添加。
- 如果子系统描述分布在多个文件中，请在每个文件名中加上子系统名称作为前缀（例如，使用 dev_*.txt 描述 /dev/ 设备，使用 socket_*.txt 描述套接字等）。

**运行命令**

```c
make extract TARGETOS=linux SOURCEDIR=$KSRC

make generate

make
```

**运行 syzkaller**

使用[覆盖率](coverage.md)信息页面确保 syzkaller 已访问新增接口。

**描述技巧和常见问题解答**

**系统调用、结构体、字段、标志名称**

尽量使用现有的内核名称，不要另创新名称。这种做法带来的好处包括：

1. 名称与内核接口一致，便于搜索相关源代码。

2. 允许通过 [syz-check]() 进行描述的静态检查（如标志遗漏或字段拼写错误）。

**测试描述**

为了确保描述的准确性，建议在添加新的描述后，通过 syz-manager 的内核代码覆盖率报告进行验证。运行测试的详细步骤和调试方式可以参阅文档末尾的测试部分。

