# o2sdump

```py
# version v2.0.1223.2, refs to https://xiaomi.f.mioffice.cn/docx/doxk4gjWUhkDAreoLhnZkMauXdf
version = 'v2.0.1223.2'


def get_fastboot_cmds(dump_path):
    return [
        "getvar version",
        f'oem memdump 0xa0000000:0x100000 "{dump_path}/pstore.bin"',
        f'oem memdump 0x43480000:0x800000 "{dump_path}/mdr.bin"',
        f'oem memdump 0x43480408:0x1F0 "{dump_path}/mdump_head.bin"',
        f'oem memdump 0x43481400:0x40000 "{dump_path}/fastboot_log.bin"',
        f'oem flashdump oops 0x0:0x1000000 "{dump_path}/oops.bin"',
        f'oem flashdump logfs 0x0:0x800000 "{dump_path}/logfs.bin"',
        "continue",
    ]


def get_adb_cmds(dump_path):
    return [
        {'method': 'pull', 'src': '/sys/kernel/tracing', 'dst': f'{dump_path}/'},
        {'method': 'pull', 'src': '/sys/kernel/debug/tracing', 'dst': f'{dump_path}/'},
        {'method': 'shell', 'cmd': 'screencap -p /sdcard/screen.png', 'dst': None},
        {'method': 'pull', 'src': '/sdcard/screen.png', 'dst': f'{dump_path}/'},
        {'method': 'pull', 'src': '/sdcard/DCIM/ScreenRecorder/', 'dst': f'{dump_path}/'},
        {'method': 'pull', 'src': '/sdcard/DCIM/Screenshots/', 'dst': f'{dump_path}/'},
        {'method': 'pull', 'src': '/data/local/log/', 'dst': f'{dump_path}/'},
        {'method': 'pull', 'src': '/data/local/traces/', 'dst': f'{dump_path}/'},
        {'method': 'pull', 'src': '/data/vendor/diag/', 'dst': f'{dump_path}/'},
        {'method': 'pull', 'src': '/data/miuilog/', 'dst': f'{dump_path}/'},
        {'method': 'pull', 'src': '/data/xring/acpu/kmsg/', 'dst': f'{dump_path}/'},
        {'method': 'pull', 'src': '/data/anr', 'dst': f'{dump_path}/'},
        {'method': 'pull', 'src': '/data/misc/logd', 'dst': f'{dump_path}/'},
        {'method': 'pull', 'src': '/data/xring_logs', 'dst': f'{dump_path}/'},
        {'method': 'pull', 'src': '/data/tombstones', 'dst': f'{dump_path}/'},
        {'method': 'pull', 'src': '/sys/kernel/debug/binder', 'dst': f'{dump_path}/'},
        {'method': 'pull', 'src': '/data/system/dropbox', 'dst': f'{dump_path}/'},
        {'method': 'pull', 'src': '/data/vendor/xring_logs/', 'dst': f'{dump_path}/'},
        {'method': 'pull', 'src': '/data/vendor/mitee/', 'dst': f'{dump_path}/'},
        {'method': 'pull', 'src': '/data/vendor/audio/', 'dst': f'{dump_path}/'},
        {'method': 'pull', 'src': '/data/vendor/logs/npu_logs', 'dst': f'{dump_path}/'},
        {'method': 'pull', 'src': '/data/vendor/logs/vdsp_logs', 'dst': f'{dump_path}/'},
        {'method': 'shell', 'cmd': 'rm -rf /data/vendor/offlinelog', 'dst': None},
        {'method': 'shell', 'cmd': 'rm -rf /data/vendor/EdrLog', 'dst': None},
        {'method': 'shell', 'cmd': 'cp -rf /data/vendor/camera/offlinelog /data/vendor/offlinelog', 'dst': None},
        {'method': 'shell', 'cmd': 'cp -rf /data/vendor/camera/EdrLog /data/vendor/EdrLog', 'dst': None},
        {'method': 'pull', 'src': '/data/vendor/offlinelog', 'dst': f'{dump_path}/'},
        {'method': 'pull', 'src': '/data/vendor/EdrLog', 'dst': f'{dump_path}/'},
        {'method': 'pull', 'src': '/data/vendor/ramdump', 'dst': f'{dump_path}/vendor/'},
        {'method': 'pull', 'src': '/data/vendor/mdl_coll', 'dst': f'{dump_path}/'},
        {'method': 'pull', 'src': '/data/vendor/sap_log', 'dst': f'{dump_path}/'},
        {'method': 'pull', 'src': '/data/data/screen.png', 'dst': f'{dump_path}/'},
        {'method': 'pull', 'src': '/sdcard/DCIM/Camera', 'dst': f'{dump_path}/'},
        {'method': 'pull', 'src': '/data/debuglogger', 'dst': f'{dump_path}/'},
        {'method': 'pull', 'src': '/data/ramdump', 'dst': f'{dump_path}/'},
        {'method': 'pull', 'src': '/data/logs/android_logs/kmsg', 'dst': f'{dump_path}/android/'},
        {'method': 'pull', 'src': '/dev/binderfs/binder_logs', 'dst': f'{dump_path}/'},
        {'method': 'pull', 'src': '/data/vendor/wlan_logs/', 'dst': f'{dump_path}/wcn_dump/'},
        {'method': 'pull', 'src': '/vendor/firmware/', 'dst': f'{dump_path}/wcn_dump/'},
        {'method': 'pull', 'src': '/data/connsyslog', 'dst': f'{dump_path}/wcn_dump/'},
        {'method': 'pull', 'src': '/data/log_wifi_temp', 'dst': f'{dump_path}/wcn_dump/'},
        {'method': 'pull', 'src': '/data/vendor/connsyslog/', 'dst': f'{dump_path}/wcn_dump/vendor/'},
        {'method': 'pull', 'src': '/data/vendor/thinmd/aee_exp', 'dst': f'{dump_path}/'},
        {'method': 'pull', 'src': '/mnt/vendor/nvdata/md/', 'dst': f'{dump_path}/nvram_log/'},
        {'method': 'pull', 'src': '/mnt/vendor/protect_f/md/', 'dst': f'{dump_path}/nvram_log/p_f/'},
        {'method': 'pull', 'src': '/mnt/vendor/protect_s/md/', 'dst': f'{dump_path}/nvram_log/p_s/'},
        {'method': 'pull', 'src': '/data/vendor/md_rfs', 'dst': f'{dump_path}/nvram_log/md_rfs'},
        {'method': 'shell', 'cmd': 'dmesg', 'dst': f'{dump_path}/dmesg.txt'},
        {'method': 'shell', 'cmd': 'mount', 'dst': f'{dump_path}/mount.log'},
        {'method': 'shell', 'cmd': 'cat /proc/tmi_log0', 'dst': f'{dump_path}/tmi_log.txt'},
        {'method': 'shell', 'cmd': 'dumpsys media.camera', 'dst': f'{dump_path}/media.camera.coredump'},
        {'method': 'shell', 'cmd': 'setprop persist.log.tag V', 'dst': None},
        {'method': 'shell', 'cmd': 'getprop ro.build.fingerprint', 'dst': f'{dump_path}/version.txt'},
        {'method': 'shell', 'cmd': 'dumpsys SurfaceFlinger', 'dst': f'{dump_path}/SurfaceFlinger.txt'},
        {'method': 'shell', 'cmd': 'dumpsys window -a', 'dst': f'{dump_path}/window.txt'},
        {'method': 'shell', 'cmd': 'dumpsys activity activities', 'dst': f'{dump_path}/activities.txt'},
        {'method': 'shell', 'cmd': 'dumpsys activity top', 'dst': f'{dump_path}/top_activity.txt'},
        {'method': 'shell', 'cmd': 'dumpsys activity containers', 'dst': f'{dump_path}/containers.txt'},
        {'method': 'shell', 'cmd': 'dumpsys activity input', 'dst': f'{dump_path}/input.txt'},
        {'method': 'shell', 'cmd': 'dumpsys battery', 'dst': f'{dump_path}/battery.info'},
        {'method': 'shell', 'cmd': 'ps -ef', 'dst': f'{dump_path}/ps.txt'},
        {'method': 'shell', 'cmd': 'cat /proc/meminfo', 'dst': f'{dump_path}/memory.txt'},
        {'method': 'shell', 'cmd': 'free -h', 'dst': f'{dump_path}/free_memory.txt'},
        {'method': 'shell', 'cmd': 'vmstat', 'dst': f'{dump_path}/free_memory.txt'},
        {'method': 'shell', 'cmd': 'top -n 1', 'dst': f'{dump_path}/free_memory.txt'},
        {'method': 'shell', 'cmd': 'cat /proc/tmi_log0', 'dst': f'{dump_path}/tmi_log.txt'},
        {'method': 'exec-out', 'cmd': 'dumpsys window --proto', 'dst': f'{dump_path}/window_dump.winscope'},
        {'method': 'exec-out', 'cmd': 'dumpsys SurfaceFlinger --proto', 'dst': f'{dump_path}/SurfaceFlinger_dump.winscope'},
        {'method': 'shell', 'cmd': 'mount -t debugfs debugfs /sys/kernel/debug', 'dst': None},
        {'method': 'shell', 'cmd': 'cat /sys/kernel/debug/wakeup_sources', 'dst': f'{dump_path}/wakeup_sources.txt'},
        {'method': 'shell', 'cmd': 'dumpsys suspend_control_internal', 'dst': f'{dump_path}/suspend_control_internal.txt'},
        {'method': 'shell', 'cmd': 'dumpsys power', 'dst': f'{dump_path}/power.txt'},
    ]

```





# 手机中自动抓 dump

### FAT和qpst

- 橙屏定屏命令：
  - 橙屏定屏：使能该命令后，橙屏不会自动抓dump，需要人工处理。

```shell
fastboot oem ramdump qpst
```

  - 橙屏不定屏： 橙屏抓dump到dump分区，手机自动重启。这种情况是手机默认行为。 

  ```shell
fastboot oem ramdump fat
  ```

- 暗码修改

```shell
*#*#27274#*#*
```




# ubuntu中自动抓dump

1. 获取文件、安装
   1. git clone  https://github.com/bkerler/edl
   2. cd edl
   3. git submodule update --init --recursive
   4. pip3 install -r requirements.txt
2. Linux (Debian/Ubuntu/Mint/etc)

```bash
# Debian/Ubuntu/Mint/etc
sudo apt install adb fastboot python3-dev python3-pip liblzma-dev git
sudo apt purge modemmanager
# Fedora/CentOS/etc
sudo dnf install adb fastboot python3-devel python3-pip xz-devel git
# Arch/Manjaro/etc
sudo pacman -S android-tools python python-pip git xz
sudo pacman -R modemmanager

sudo systemctl stop ModemManager
sudo systemctl disable ModemManager
sudo apt purge ModemManager


git clone https://github.com/bkerler/edl.git
cd edl
git submodule update --init --recursive
chmod +x ./install-linux-edl-drivers.sh
bash ./install-linux-edl-drivers.sh
python3 setup.py build
sudo python3 setup.py install
```

3. dump文件获取

```bash
Edl memorydump
```

4. 备注

在Ubuntu上的conda py39环境下运行



