1. 通过 syzkaller 进行 fuzz 测试，使用如下配置进行测试

```
git clone https://github.com/google/syzkaller
cd syzkaller
make -j1 TARGETOS=linux TARGETARCH=arm64
```

