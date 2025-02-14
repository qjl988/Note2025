```sh
2025/02/14 10:40:11 serving rpc on tcp://40947
2025/02/14 10:40:11 serving http on http://10.189.131.210:56741
2025/02/14 10:40:12 executing adb [wait-for-device]
2025/02/14 10:40:13 skipped 7 seeds
2025/02/14 10:41:42 adb returned
2025/02/14 10:41:42 executing adb [shell reboot]
2025/02/14 10:41:54 adb returned
2025/02/14 10:42:05 executing adb [wait-for-device]
2025/02/14 10:45:48 adb returned
2025/02/14 10:45:48 executing adb [root]
2025/02/14 10:45:52 adb returned
2025/02/14 10:45:53 executing adb [wait-for-device]
2025/02/14 10:45:53 adb returned
2025/02/14 10:45:58 executing adb [shell pgrep systemui | wc -l]
2025/02/14 10:46:00 adb returned
2025/02/14 10:47:02 boot completed
2025/02/14 10:47:02 executing adb [shell ls /sys/kernel/debug/kcov]
2025/02/14 10:47:03 adb returned
2025/02/14 10:47:04 failed to associate adb device db47dfe2 with console: failed to open console file: permission denied
2025/02/14 10:47:04 falling back to 'adb shell dmesg -w'
2025/02/14 10:47:04 note: some bugs may be detected as 'lost connection to test machine' with no kernel output
2025/02/14 10:47:04 associating adb device db47dfe2 with console adb
2025/02/14 10:47:04 executing adb [shell dumpsys battery | grep level:]
2025/02/14 10:47:06 adb returned
2025/02/14 10:47:06 device db47dfe2: battery level 77%, OK
2025/02/14 10:47:06 executing adb [shell ls /data/syzkaller*]
2025/02/14 10:47:07 adb returned
2025/02/14 10:47:07 executing adb [shell find /data/syzkaller* -type l -exec unlink {} \; && rm -Rf /data/syzkaller*]
2025/02/14 10:47:08 adb returned
2025/02/14 10:47:08 executing adb [shell echo 0 > /proc/sys/kernel/kptr_restrict]
2025/02/14 10:47:09 adb returned
2025/02/14 10:47:09 executing adb [reverse tcp:17280 tcp:40947]
2025/02/14 10:47:09 adb returned
2025/02/14 10:47:09 executing adb [push /home/xiaomi/Documents/syzkaller/bin/linux_arm64/syz-executor /data/syz-executor]
2025/02/14 10:47:10 adb returned
2025/02/14 10:47:10 executing adb [shell chmod +x /data/syz-executor]
2025/02/14 10:47:11 adb returned
2025/02/14 10:47:11 starting: adb shell /data/syz-executor runner 0 127.0.0.1 17280
2025/02/14 10:47:12 runner 0 connected
```

