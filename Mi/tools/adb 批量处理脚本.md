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
        ("wait-for-device", "等待设备连接"),
        ("root", "获取 ROOT 权限"),
        ("shell setprop persist.vendor.ssr.restart_level ALL_DISABLE", "关闭子系统重启"),
        ("shell settings put global systemui_keyguard_whitelist -1266261377", "设置系统 UI 锁屏白名单"),
        ("shell settings put global systemui_float_whitelist -196939687", "设置浮动窗口白名单"),
        ("shell settings put global miui_permission_check 3", "设置 MIUI 权限检查"),
        ("shell settings put global miui_account_login_check 0", "关闭 MIUI 账号登录检查"),
        ("shell settings put global device_provisioned 1", "设备已配置"),
        ("shell settings put global miui_pre_version_otaprovision 140", "设置 MIUI 预版本 OTA 配置"),
        ("shell settings put secure auto_download 0", "禁用自动下载"),
        ("shell settings put secure auto_update 0", "禁用自动更新"),
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

