### adb 设备列表打印

```sh
# 获取所有已连接的设备序列号
get_connected_devices() {
    devices=$(adb devices | grep -w 'device' | awk '{print $1}')
    echo "$devices"
}
```

