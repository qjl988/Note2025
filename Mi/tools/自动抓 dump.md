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



