### FAT和qpst

- 橙屏定屏命令：
  - 橙屏定屏：使能该命令后，橙屏不会自动抓dump，需要人工处理。 `fastboot oem ramdump qpst`
  - 橙屏不定屏： 橙屏抓dump到dump分区，手机自动重启。这种情况是手机默认行为。 `fastboot oem ramdump fat`