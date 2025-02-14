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



### O1提高hub充电电流

```shell
adb shell "echo mtbf buck 1500 > /sys/class/xm_power/charger/charge_interface/iin_limit"
```



### 关闭子系统重启命令

```Bash
 adb root
 adb shell
 
 adb shell setprop persist.vendor.ssr.restart_level  ALL_ENABLE   //开启所有子系统SSR
 adb shell setprop persist.vendor.ssr.restart_level  ALL_DISABLE   //关闭所有子系统SSR
 adb shell setprop persist.vendor.ssr.restart_level  adsp,wlan  //开启adsp，wlan子系统SSR
 
adb shell getprop persist.vendor.ssr.restart_level    //查询SSR开启情况
```



### 查看DP（debugpolicy）是否开启

https://security.pt.xiaomi.com/#/dp/flashDebugPolicy

```sh
检查dp命令：adb shell getprop ro.boot.dp，高通和MTK新项目 0XB为开启，MTK老项目 2为开启
```

