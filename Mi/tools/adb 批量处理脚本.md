# 运行f和a

```sh
#!/bin/bash

# 执行 ADB 或 fastboot 命令并打印执行信息
run_command() {
    local command=$1
    local description=$2
    local device_serial=$3

    if [ -n "$device_serial" ]; then
        command="adb -s $device_serial $command"
    fi

    echo "[$device_serial] 执行: $description..."
    result=$(eval "$command" 2>&1)
    if [ $? -eq 0 ]; then
        echo "[$device_serial] $description 完成: ${result:-'无输出'}"
    else
        echo "[$device_serial] 执行 $description 时出错: $result"
    fi
}

# 获取所有已连接的设备序列号
get_connected_devices() {
    devices=$(adb devices | grep -w 'device' | awk '{print $1}')
    echo "$devices"
}

# 配置设备
configure_device() {
    local device_serial=$1

    # 亮屏
    run_command "shell input keyevent 26" "亮屏" "$device_serial"
    sleep 2

    # 检查设备配置状态
    run_command "shell settings get global device_provisioned" "检查设备配置状态" "$device_serial"

    # 进入 fastboot 模式并重启
    run_command "reboot bootloader" "进入 fastboot 模式" "$device_serial"
    sleep 2
    run_command "fastboot reboot" "重启设备" "$device_serial"
    echo "[$device_serial] Android 设备设置已更新并重启。\n"
}

# 主程序
devices=$(get_connected_devices)

if [ -z "$devices" ]; then
    echo "未检测到任何连接的 Android 设备。"
else
    device_count=$(echo "$devices" | wc -l)
    echo "检测到 $device_count 台设备: $devices"
    for device in $devices; do
        echo -e "\n开始配置设备 [$device]..."
        configure_device "$device"
    done
    echo "所有设备已配置完成。"
fi
```









# wangsheng


```py
import subprocess
import time
import threading

# 用于控制打印的锁，避免多线程输出混乱
print_lock = threading.Lock()


def run_adb_command(command, description, device_serial=None):
    """运行 ADB 命令，并打印执行信息。"""
    if device_serial:
        command = f"adb -s {device_serial} {command}"
    with print_lock:
        print(f"[{device_serial}] 执行: {description}...")
    try:
        result = subprocess.run(command, shell=True, capture_output=True, text=True)
        output = result.stdout.strip()
        with print_lock:
            print(f"[{device_serial}] {description} 完成: {output if output else '无输出'}")
    except Exception as e:
        with print_lock:
            print(f"[{device_serial}] 执行 {description} 时出错: {e}")
        return None


def get_connected_devices():
    """获取所有已连接的设备序列号列表"""
    while True:
        result = subprocess.run("adb devices", shell=True, capture_output=True, text=True)
        lines = result.stdout.strip().split("\n")[1:]  # 跳过第一行
        devices = [line.split("\t")[0] for line in lines if "device" in line]
        if devices:
            return devices
        time.sleep(5)  # 等待设备连接


def configure_device(device_serial, commands):
    """对指定设备执行多个配置操作"""
    for cmd, des in commands:
        run_adb_command(cmd, des, device_serial)
    with print_lock:
        print(f"[{device_serial}] 所有配置操作完成。\n")


def process_device(device, commands):
    """对单个设备执行多个配置操作"""
    print(f"\n开始配置设备 [{device}]...")
    configure_device(device, commands)


if __name__ == "__main__":
    devices = get_connected_devices()
    print(f"检测到 {len(devices)} 个设备: {devices}")

    # 定义要执行的多个命令和描述
    commands = [
        ("reboot", "系统重启"),
    ]

    # 创建并启动线程
    threads = []
    for device in devices:
        thread = threading.Thread(
            target=process_device,
            args=(device, commands)
        )
        threads.append(thread)
        thread.start()

    # 等待所有线程完成
    for thread in threads:
        thread.join()

    print("所有设备已配置完成。")
```

