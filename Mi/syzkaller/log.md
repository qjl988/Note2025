# O2

```sh
2025/02/24 18:22:12 serving rpc on tcp://34685
2025/02/24 18:22:12 serving http on http://10.189.145.186:56749
2025/02/24 18:22:13 executing adb [wait-for-device]
2025/02/24 18:22:13 adb returned
2025/02/24 18:22:13 executing adb [shell reboot]
2025/02/24 18:22:13 skipped 7 seeds
2025/02/24 18:22:20 adb returned
2025/02/24 18:22:31 executing adb [wait-for-device]
2025/02/24 18:23:27 adb returned
2025/02/24 18:23:27 executing adb [root]
2025/02/24 18:23:27 adb returned
2025/02/24 18:23:28 executing adb [wait-for-device]
2025/02/24 18:23:28 adb returned
2025/02/24 18:23:33 executing adb [shell pgrep systemui | wc -l]
2025/02/24 18:23:34 adb returned
2025/02/24 18:23:39 executing adb [shell pgrep systemui | wc -l]
2025/02/24 18:23:39 adb returned
2025/02/24 18:23:44 executing adb [shell pgrep systemui | wc -l]
2025/02/24 18:23:46 adb returned
2025/02/24 18:23:46 executing adb [root]
2025/02/24 18:23:46 adb returned
2025/02/24 18:23:46 root access granted
2025/02/24 18:23:46 executing adb [shell echo mtbf buck 1500 > /sys/class/xm_power/charger/charge_interface/iin_limit]
2025/02/24 18:23:46 adb returned
2025/02/24 18:23:46 charge limit set successfully
2025/02/24 18:23:46 boot completed
2025/02/24 18:23:46 executing adb [shell ls /sys/kernel/debug/kcov]
2025/02/24 18:23:46 adb returned
2025/02/24 18:23:46 failed to associate adb device 0A03C47243448540 with console: no unassociated console devices left
2025/02/24 18:23:46 falling back to 'adb shell dmesg -w'
2025/02/24 18:23:46 note: some bugs may be detected as 'lost connection to test machine' with no kernel output
2025/02/24 18:23:46 associating adb device 0A03C47243448540 with console adb
2025/02/24 18:23:46 executing adb [shell dumpsys battery | grep level:]
2025/02/24 18:23:46 adb returned
2025/02/24 18:23:46 device 0A03C47243448540: battery level 76%, OK
2025/02/24 18:23:46 executing adb [shell ls /data/syzkaller*]
2025/02/24 18:23:47 adb returned
2025/02/24 18:23:47 executing adb [shell find /data/syzkaller* -type l -exec unlink {} \; && rm -Rf /data/syzkaller*]
2025/02/24 18:23:48 adb returned
2025/02/24 18:23:48 executing adb [shell echo 0 > /proc/sys/kernel/kptr_restrict]
2025/02/24 18:23:48 adb returned
2025/02/24 18:23:48 executing adb [reverse tcp:10566 tcp:34685]
2025/02/24 18:23:48 adb returned
2025/02/24 18:23:48 executing adb [push /home/xiaomi/Documents/o1syzkaller/bin/linux_arm64/syz-executor /data/syz-executor]
2025/02/24 18:23:48 adb returned
2025/02/24 18:23:48 executing adb [shell chmod +x /data/syz-executor]
2025/02/24 18:23:49 adb returned
2025/02/24 18:23:49 starting: adb shell /data/syz-executor runner 0 127.0.0.1 10566
2025/02/24 18:23:49 runner 0 connected
2025/02/24 18:25:02 machine check:
disabled the following syscalls:
acct                                        : syscall acct is not present
fanotify_init                               : syscall fanotify_init is not present
fanotify_mark                               : syscall fanotify_mark is not present
get_mempolicy                               : syscall get_mempolicy is not present
kexec_load                                  : syscall kexec_load is not present
landlock_add_rule$LANDLOCK_RULE_NET_PORT    : syscall landlock_add_rule$LANDLOCK_RULE_NET_PORT is not present
landlock_add_rule$LANDLOCK_RULE_PATH_BENEATH: syscall landlock_add_rule$LANDLOCK_RULE_NET_PORT is not present
landlock_create_ruleset                     : syscall landlock_create_ruleset is not present
landlock_restrict_self                      : syscall landlock_restrict_self is not present
lookup_dcookie                              : syscall lookup_dcookie is not present
lsm_get_self_attr                           : syscall lsm_get_self_attr is not present
lsm_list_modules                            : syscall lsm_list_modules is not present
lsm_set_self_attr                           : syscall lsm_set_self_attr is not present
map_shadow_stack                            : syscall map_shadow_stack is not present
mbind                                       : syscall mbind is not present
migrate_pages                               : syscall migrate_pages is not present
mount$9p_fd                                 : /proc/filesystems does not contain 9p
mount$9p_rdma                               : /proc/filesystems does not contain 9p
mount$9p_tcp                                : /proc/filesystems does not contain 9p
mount$9p_unix                               : /proc/filesystems does not contain 9p
mount$9p_virtio                             : /proc/filesystems does not contain 9p
mount$9p_xen                                : /proc/filesystems does not contain 9p
mount$afs                                   : /proc/filesystems does not contain afs
mount$esdfs                                 : /proc/filesystems does not contain esdfs
mount$nfs                                   : /proc/filesystems does not contain nfs
mount$nfs4                                  : /proc/filesystems does not contain nfs4
mount$pvfs2                                 : /proc/filesystems does not contain pvfs2
move_pages                                  : syscall move_pages is not present
mq_getsetattr                               : syscall mq_getsetattr is not present
mq_notify                                   : syscall mq_notify is not present
mq_open                                     : syscall mq_open is not present
mq_timedreceive                             : syscall mq_timedreceive is not present
mq_timedsend                                : syscall mq_timedsend is not present
mq_unlink                                   : syscall mq_unlink is not present
msgctl$IPC_INFO                             : syscall msgctl$IPC_INFO is not present
msgctl$IPC_RMID                             : syscall msgctl$IPC_INFO is not present
msgctl$IPC_SET                              : syscall msgctl$IPC_INFO is not present
msgctl$IPC_STAT                             : syscall msgctl$IPC_INFO is not present
msgctl$MSG_INFO                             : syscall msgctl$IPC_INFO is not present
msgctl$MSG_STAT                             : syscall msgctl$IPC_INFO is not present
msgctl$MSG_STAT_ANY                         : syscall msgctl$IPC_INFO is not present
msgget                                      : syscall msgget is not present
msgget$private                              : syscall msgget is not present
msgrcv                                      : syscall msgrcv is not present
msgsnd                                      : syscall msgsnd is not present
name_to_handle_at                           : syscall name_to_handle_at is not present
open_by_handle_at                           : syscall open_by_handle_at is not present
openat$6lowpan_control                      : failed to open /sys/kernel/debug/bluetooth/6lowpan_control: no such file or directory
openat$6lowpan_enable                       : failed to open /sys/kernel/debug/bluetooth/6lowpan_enable: no such file or directory
openat$acpi_thermal_rel                     : failed to open /dev/acpi_thermal_rel: no such file or directory
openat$adsp1                                : failed to open /dev/adsp1: no such file or directory
openat$audio                                : failed to open /dev/audio: no such file or directory
openat$audio1                               : failed to open /dev/audio1: no such file or directory
openat$autofs                               : failed to open /dev/autofs: no such file or directory
openat$bifrost                              : failed to open /dev/bifrost: no such file or directory
openat$bsg                                  : failed to open /dev/bsg: no such file or directory
openat$btrfs_control                        : failed to open /dev/btrfs-control: no such file or directory
openat$cachefiles                           : failed to open /dev/cachefiles: no such file or directory
openat$camx                                 : failed to open /dev/v4l/by-path/platform-soc@0:qcom_cam-req-mgr-video-index0: no such file or directory
openat$capi20                               : failed to open /dev/capi20: no such file or directory
openat$cdrom                                : failed to open /dev/cdrom: no such file or directory
openat$cdrom1                               : failed to open /dev/cdrom1: no such file or directory
openat$cuse                                 : failed to open /dev/cuse: no such file or directory
openat$damon_attrs                          : failed to open /sys/kernel/debug/damon/attrs: no such file or directory
openat$damon_init_regions                   : failed to open /sys/kernel/debug/damon/init_regions: no such file or directory
openat$damon_kdamond_pid                    : failed to open /sys/kernel/debug/damon/kdamond_pid: no such file or directory
openat$damon_mk_contexts                    : failed to open /sys/kernel/debug/damon/mk_contexts: no such file or directory
openat$damon_monitor_on                     : failed to open /sys/kernel/debug/damon/monitor_on: no such file or directory
openat$damon_rm_contexts                    : failed to open /sys/kernel/debug/damon/rm_contexts: no such file or directory
openat$damon_schemes                        : failed to open /sys/kernel/debug/damon/schemes: no such file or directory
openat$damon_target_ids                     : failed to open /sys/kernel/debug/damon/target_ids: no such file or directory
openat$dlm_control                          : failed to open /dev/dlm-control: no such file or directory
openat$dlm_monitor                          : failed to open /dev/dlm-monitor: no such file or directory
openat$dlm_plock                            : failed to open /dev/dlm_plock: no such file or directory
openat$drirender128                         : failed to open /dev/dri/renderD128: no such file or directory
openat$dsp                                  : failed to open /dev/dsp: no such file or directory
openat$dsp1                                 : failed to open /dev/dsp1: no such file or directory
openat$fb0                                  : failed to open /dev/fb0: no such file or directory
openat$fb1                                  : failed to open /dev/fb1: no such file or directory
openat$hpet                                 : failed to open /dev/hpet: no such file or directory
openat$hwrng                                : failed to open /dev/hwrng: no such file or directory
openat$i915                                 : failed to open /dev/i915: no such file or directory
openat$img_rogue                            : failed to open /dev/img-rogue: no such file or directory
openat$iommufd                              : failed to open /dev/iommu: no such file or directory
openat$ipvs                                 : failed to open /proc/sys/net/ipv4/vs/sync_qlen_max: no such file or directory
openat$irnet                                : failed to open /dev/irnet: no such file or directory
openat$keychord                             : failed to open /dev/keychord: no such file or directory
openat$kvm                                  : failed to open /dev/kvm: no such file or directory
openat$lightnvm                             : failed to open /dev/lightnvm/control: no such file or directory
openat$md                                   : failed to open /dev/md0: no such file or directory
openat$mice                                 : failed to open /dev/input/mice: no such file or directory
openat$misdntimer                           : failed to open /dev/mISDNtimer: no such file or directory
openat$mixer                                : failed to open /dev/mixer: no such file or directory
openat$msm                                  : failed to open /dev/msm: no such file or directory
openat$nci                                  : failed to open /dev/virtual_nci: no such file or directory
openat$ndctl0                               : failed to open /dev/ndctl0: no such file or directory
openat$nmem0                                : failed to open /dev/nmem0: no such file or directory
openat$nullb                                : failed to open /dev/nullb0: no such file or directory
openat$nvme_fabrics                         : failed to open /dev/nvme-fabrics: no such file or directory
openat$nvram                                : failed to open /dev/nvram: no such file or directory
openat$ocfs2_control                        : failed to open /dev/ocfs2_control: no such file or directory
openat$pktcdvd                              : failed to open /dev/pktcdvd/control: no such file or directory
openat$pmem0                                : failed to open /dev/pmem0: no such file or directory
openat$ppp                                  : failed to open /dev/ppp: no such file or directory
openat$proc_capi20                          : failed to open /proc/capi/capi20: no such file or directory
openat$proc_capi20ncci                      : failed to open /proc/capi/capi20ncci: no such file or directory
openat$proc_mixer                           : failed to open /proc/asound/card0/oss_mixer: no such file or directory
openat$proc_reclaim                         : failed to open /proc/self/reclaim: no such file or directory
openat$ptp0                                 : failed to open /dev/ptp0: no such file or directory
openat$ptp1                                 : failed to open /dev/ptp1: no such file or directory
openat$qat_adf_ctl                          : failed to open /dev/qat_adf_ctl: no such file or directory
openat$qrtrtun                              : failed to open /dev/qrtr-tun: no such file or directory
openat$rdma_cm                              : failed to open /dev/infiniband/rdma_cm: no such file or directory
openat$sequencer                            : failed to open /dev/sequencer: no such file or directory
openat$sequencer2                           : failed to open /dev/sequencer2: no such file or directory
openat$sgx_provision                        : failed to open /dev/sgx_provision: no such file or directory
openat$smackfs_access                       : failed to open /sys/fs/smackfs/access: no such file or directory
openat$smackfs_ambient                      : failed to open /sys/fs/smackfs/ambient: no such file or directory
openat$smackfs_change_rule                  : failed to open /sys/fs/smackfs/change-rule: no such file or directory
openat$smackfs_cipso                        : failed to open /sys/fs/smackfs/cipso: no such file or directory
openat$smackfs_cipsonum                     : failed to open /sys/fs/smackfs/direct: no such file or directory
openat$smackfs_ipv6host                     : failed to open /sys/fs/smackfs/ipv6host: no such file or directory
openat$smackfs_load                         : failed to open /sys/fs/smackfs/load: no such file or directory
openat$smackfs_logging                      : failed to open /sys/fs/smackfs/logging: no such file or directory
openat$smackfs_netlabel                     : failed to open /sys/fs/smackfs/netlabel: no such file or directory
openat$smackfs_onlycap                      : failed to open /sys/fs/smackfs/onlycap: no such file or directory
openat$smackfs_ptrace                       : failed to open /sys/fs/smackfs/ptrace: no such file or directory
openat$smackfs_relabel_self                 : failed to open /sys/fs/smackfs/relabel-self: no such file or directory
openat$smackfs_revoke_subject               : failed to open /sys/fs/smackfs/revoke-subject: no such file or directory
openat$smackfs_syslog                       : failed to open /sys/fs/smackfs/syslog: no such file or directory
openat$smackfs_unconfined                   : failed to open /sys/fs/smackfs/unconfined: no such file or directory
openat$snapshot                             : failed to open /dev/snapshot: no such file or directory
openat$sndseq                               : failed to open /dev/snd/seq: no such file or directory
openat$sr                                   : failed to open /dev/sr0: no such file or directory
openat$sw_sync                              : failed to open /sys/kernel/debug/sync/sw_sync: no such file or directory
openat$sw_sync_info                         : failed to open /sys/kernel/debug/sync/info: no such file or directory
openat$sysctl                               : failed to open /sys/kernel/mm/ksm/run: no such file or directory
openat$tlk_device                           : failed to open /dev/tlk_device: no such file or directory
openat$trusty                               : failed to open /dev/trusty-ipc-dev0: no such file or directory
openat$trusty_avb                           : failed to open /dev/trusty-ipc-dev0: no such file or directory
openat$trusty_gatekeeper                    : failed to open /dev/trusty-ipc-dev0: no such file or directory
openat$trusty_hwkey                         : failed to open /dev/trusty-ipc-dev0: no such file or directory
openat$trusty_hwrng                         : failed to open /dev/trusty-ipc-dev0: no such file or directory
openat$trusty_km                            : failed to open /dev/trusty-ipc-dev0: no such file or directory
openat$trusty_km_secure                     : failed to open /dev/trusty-ipc-dev0: no such file or directory
openat$trusty_storage                       : failed to open /dev/trusty-ipc-dev0: no such file or directory
openat$tty                                  : failed to open /dev/tty: no such device or address
openat$ttyS3                                : failed to open /dev/ttyS3: no such file or directory
openat$ttyprintk                            : failed to open /dev/ttyprintk: no such file or directory
openat$ubi_ctrl                             : failed to open /dev/ubi_ctrl: no such file or directory
openat$udambuf                              : failed to open /dev/udmabuf: no such file or directory
openat$userio                               : failed to open /dev/userio: no such file or directory
openat$uverbs0                              : failed to open /dev/infiniband/uverbs0: no such file or directory
openat$vcs                                  : failed to open /dev/vcs: no such file or directory
openat$vcsa                                 : failed to open /dev/vcsa: no such file or directory
openat$vcsu                                 : failed to open /dev/vcsu: no such file or directory
openat$vfio                                 : failed to open /dev/vfio/vfio: no such file or directory
openat$vga_arbiter                          : failed to open /dev/vga_arbiter: no such file or directory
openat$vicodec0                             : failed to open /dev/video36: no such file or directory
openat$vicodec1                             : failed to open /dev/video37: no such file or directory
openat$vim2m                                : failed to open /dev/vim2m: no such file or directory
openat$vimc0                                : failed to open /dev/video0: invalid argument
openat$vmci                                 : failed to open /dev/vmci: no such file or directory
openat$vnet                                 : failed to open /dev/vhost-net: no such file or directory
openat$vtpm                                 : failed to open /dev/vtpmx: no such file or directory
openat$xenevtchn                            : failed to open /dev/xen/evtchn: no such file or directory
openat$yama_ptrace_scope                    : failed to open /proc/sys/kernel/yama/ptrace_scope: no such file or directory
openat$zygote                               : failed to open /dev/socket/zygote: no such device or address
pkey_alloc                                  : pkey_alloc(0x0, 0x0) failed: function not implemented
pkey_free                                   : syscall pkey_free is not present
pkey_mprotect                               : syscall pkey_mprotect is not present
rseq                                        : syscall rseq is not present
semget                                      : syscall semget is not present
semget$private                              : syscall semget is not present
semop                                       : syscall semop is not present
semtimedop                                  : syscall semtimedop is not present
set_mempolicy                               : syscall set_mempolicy is not present
set_mempolicy_home_node                     : syscall set_mempolicy_home_node is not present
shmat                                       : syscall shmat is not present
shmctl$IPC_INFO                             : syscall shmctl$IPC_INFO is not present
shmctl$IPC_RMID                             : syscall shmctl$IPC_INFO is not present
shmctl$IPC_SET                              : syscall shmctl$IPC_INFO is not present
shmctl$IPC_STAT                             : syscall shmctl$IPC_INFO is not present
shmctl$SHM_INFO                             : syscall shmctl$IPC_INFO is not present
shmctl$SHM_LOCK                             : syscall shmctl$IPC_INFO is not present
shmctl$SHM_STAT                             : syscall shmctl$IPC_INFO is not present
shmctl$SHM_STAT_ANY                         : syscall shmctl$IPC_INFO is not present
shmctl$SHM_UNLOCK                           : syscall shmctl$IPC_INFO is not present
shmdt                                       : syscall shmdt is not present
shmget                                      : syscall shmget is not present
shmget$private                              : syscall shmget is not present
socket$alg                                  : socket$alg(0x26, 0x5, 0x0) failed: address family not supported by protocol
socket$caif_seqpacket                       : socket$caif_seqpacket(0x25, 0x5, 0x0) failed: address family not supported by protocol
socket$caif_stream                          : socket$caif_stream(0x25, 0x1, 0x0) failed: address family not supported by protocol
socket$can_bcm                              : socket$can_bcm(0x1d, 0x2, 0x2) failed: address family not supported by protocol
socket$can_j1939                            : socket$can_j1939(0x1d, 0x2, 0x7) failed: address family not supported by protocol
socket$can_raw                              : socket$can_raw(0x1d, 0x3, 0x1) failed: address family not supported by protocol
socket$hf                                   : socket$hf(0x13, 0x2, 0x0) failed: address family not supported by protocol
socket$inet6_dccp                           : socket$inet6_dccp(0xa, 0x6, 0x0) failed: socket type not supported
socket$inet6_mptcp                          : socket$inet6_mptcp(0xa, 0x1, 0x106) failed: protocol not supported
socket$inet6_sctp                           : socket$inet6_sctp(0xa, 0x1, 0x84) failed: protocol not supported
socket$inet_dccp                            : socket$inet_dccp(0x2, 0x6, 0x0) failed: socket type not supported
socket$inet_mptcp                           : socket$inet_mptcp(0x2, 0x1, 0x106) failed: protocol not supported
socket$inet_sctp                            : socket$inet_sctp(0x2, 0x1, 0x84) failed: protocol not supported
socket$inet_smc                             : socket$inet_smc(0x2b, 0x1, 0x0) failed: address family not supported by protocol
socket$isdn                                 : socket$isdn(0x22, 0x3, 0x0) failed: address family not supported by protocol
socket$isdn_base                            : socket$isdn_base(0x22, 0x3, 0x0) failed: address family not supported by protocol
socket$kcm                                  : socket$kcm(0x29, 0x2, 0x0) failed: address family not supported by protocol
socket$l2tp                                 : socket$l2tp(0x2, 0x2, 0x73) failed: protocol not supported
socket$l2tp6                                : socket$l2tp6(0xa, 0x2, 0x73) failed: protocol not supported
socket$nl_crypto                            : socket$nl_crypto(0x10, 0x3, 0x15) failed: protocol not supported
socket$nl_rdma                              : socket$nl_rdma(0x10, 0x3, 0x14) failed: protocol not supported
socket$phonet                               : socket$phonet(0x23, 0x2, 0x1) failed: address family not supported by protocol
socket$phonet_pipe                          : socket$phonet_pipe(0x23, 0x5, 0x2) failed: address family not supported by protocol
socket$pppl2tp                              : socket$pppl2tp(0x18, 0x1, 0x1) failed: address family not supported by protocol
socket$pppoe                                : socket$pppoe(0x18, 0x1, 0x0) failed: address family not supported by protocol
socket$pptp                                 : socket$pptp(0x18, 0x1, 0x2) failed: address family not supported by protocol
socket$qrtr                                 : socket$qrtr(0x2a, 0x2, 0x0) failed: address family not supported by protocol
socket$rds                                  : socket$rds(0x15, 0x5, 0x0) failed: address family not supported by protocol
socket$rxrpc                                : socket$rxrpc(0x21, 0x2, 0x0) failed: address family not supported by protocol
socket$tipc                                 : socket$tipc(0x1e, 0x2, 0x0) failed: address family not supported by protocol
socket$vsock_dgram                          : socket$vsock_dgram(0x28, 0x2, 0x0) failed: no such device
socketpair$tipc                             : socket(0x1e, 0x2, 0x0) failed: address family not supported by protocol
syz_80211_inject_frame                      : root failed to open /sys/class/mac80211_hwsim/: no such file or directory
syz_80211_join_ibss                         : root failed to open /sys/class/mac80211_hwsim/: no such file or directory
syz_emit_vhci                               : root failed to open /dev/vhci: no such file or directory
syz_init_net_socket$802154_dgram            : syz_init_net_socket$802154_dgram(0x24, 0x2, 0x0) failed: address family not supported by protocol
syz_init_net_socket$802154_raw              : syz_init_net_socket$802154_raw(0x24, 0x3, 0x0) failed: address family not supported by protocol
syz_init_net_socket$ax25                    : syz_init_net_socket$ax25(0x3, 0x2, 0x0) failed: address family not supported by protocol
syz_init_net_socket$bt_bnep                 : syz_init_net_socket$bt_bnep(0x1f, 0x3, 0x4) failed: protocol not supported
syz_init_net_socket$bt_cmtp                 : syz_init_net_socket$bt_cmtp(0x1f, 0x3, 0x5) failed: protocol not supported
syz_init_net_socket$bt_hidp                 : syz_init_net_socket$bt_hidp(0x1f, 0x3, 0x6) failed: protocol not supported
syz_init_net_socket$bt_rfcomm               : syz_init_net_socket$bt_rfcomm(0x1f, 0x1, 0x3) failed: protocol not supported
syz_init_net_socket$llc                     : syz_init_net_socket$llc(0x1a, 0x1, 0x0) failed: address family not supported by protocol
syz_init_net_socket$netrom                  : syz_init_net_socket$netrom(0x6, 0x5, 0x0) failed: address family not supported by protocol
syz_init_net_socket$nfc_llcp                : syz_init_net_socket$nfc_llcp(0x27, 0x1, 0x1) failed: address family not supported by protocol
syz_init_net_socket$nfc_raw                 : syz_init_net_socket$nfc_raw(0x27, 0x3, 0x0) failed: address family not supported by protocol
syz_init_net_socket$nl_rdma                 : syz_init_net_socket$nl_rdma(0x10, 0x3, 0x14) failed: protocol not supported
syz_init_net_socket$rose                    : syz_init_net_socket$rose(0xb, 0x5, 0x0) failed: address family not supported by protocol
syz_init_net_socket$x25                     : syz_init_net_socket$x25(0x9, 0x5, 0x0) failed: address family not supported by protocol
syz_kvm_setup_cpu$ppc64                     : unsupported arch
syz_kvm_setup_cpu$x86                       : unsupported arch
syz_mount_image$adfs                        : /proc/filesystems does not contain adfs
syz_mount_image$affs                        : /proc/filesystems does not contain affs
syz_mount_image$bcachefs                    : /proc/filesystems does not contain bcachefs
syz_mount_image$befs                        : /proc/filesystems does not contain befs
syz_mount_image$bfs                         : /proc/filesystems does not contain bfs
syz_mount_image$btrfs                       : /proc/filesystems does not contain btrfs
syz_mount_image$cramfs                      : /proc/filesystems does not contain cramfs
syz_mount_image$efs                         : /proc/filesystems does not contain efs
syz_mount_image$gfs2                        : /proc/filesystems does not contain gfs2
syz_mount_image$gfs2meta                    : /proc/filesystems does not contain gfs2meta
syz_mount_image$hfs                         : /proc/filesystems does not contain hfs
syz_mount_image$hfsplus                     : /proc/filesystems does not contain hfsplus
syz_mount_image$hpfs                        : /proc/filesystems does not contain hpfs
syz_mount_image$iso9660                     : /proc/filesystems does not contain iso9660
syz_mount_image$jffs2                       : /proc/filesystems does not contain jffs2
syz_mount_image$jfs                         : /proc/filesystems does not contain jfs
syz_mount_image$minix                       : /proc/filesystems does not contain minix
syz_mount_image$nilfs2                      : /proc/filesystems does not contain nilfs2
syz_mount_image$ntfs                        : /proc/filesystems does not contain ntfs
syz_mount_image$ntfs3                       : /proc/filesystems does not contain ntfs3
syz_mount_image$ocfs2                       : /proc/filesystems does not contain ocfs2
syz_mount_image$omfs                        : /proc/filesystems does not contain omfs
syz_mount_image$qnx4                        : /proc/filesystems does not contain qnx4
syz_mount_image$qnx6                        : /proc/filesystems does not contain qnx6
syz_mount_image$reiserfs                    : /proc/filesystems does not contain reiserfs
syz_mount_image$romfs                       : /proc/filesystems does not contain romfs
syz_mount_image$squashfs                    : /proc/filesystems does not contain squashfs
syz_mount_image$sysv                        : /proc/filesystems does not contain sysv
syz_mount_image$ubifs                       : /proc/filesystems does not contain ubifs
syz_mount_image$udf                         : /proc/filesystems does not contain udf
syz_mount_image$ufs                         : /proc/filesystems does not contain ufs
syz_mount_image$v7                          : /proc/filesystems does not contain v7
syz_mount_image$vxfs                        : /proc/filesystems does not contain vxfs
syz_mount_image$xfs                         : /proc/filesystems does not contain xfs
syz_mount_image$zonefs                      : /proc/filesystems does not contain zonefs
syz_open_dev$MSR                            : failed to open /dev/cpu/#/msr: no such file or directory
syz_open_dev$admmidi                        : failed to open /dev/admmidi#: no such file or directory
syz_open_dev$amidi                          : failed to open /dev/amidi#: no such file or directory
syz_open_dev$audion                         : failed to open /dev/audio#: no such file or directory
syz_open_dev$cec                            : failed to open /dev/cec#: no such file or directory
syz_open_dev$dmmidi                         : failed to open /dev/dmmidi#: no such file or directory
syz_open_dev$dricontrol                     : failed to open /dev/dri/controlD#: no such file or directory
syz_open_dev$drirender                      : failed to open /dev/dri/renderD#: no such file or directory
syz_open_dev$floppy                         : failed to open /dev/fd#: no such file or directory
syz_open_dev$hiddev                         : failed to open /dev/usb/hiddev#: no such file or directory
syz_open_dev$hidraw                         : failed to open /dev/hidraw#: no such file or directory
syz_open_dev$ircomm                         : failed to open /dev/ircomm#: no such file or directory
syz_open_dev$loop                           : failed to open /dev/loop#: no such file or directory
syz_open_dev$midi                           : failed to open /dev/midi#: no such file or directory
syz_open_dev$mouse                          : failed to open /dev/input/mouse#: no such file or directory
syz_open_dev$ndb                            : failed to open /dev/nbd#: no such file or directory
syz_open_dev$radio                          : failed to open /dev/radio#: no such file or directory
syz_open_dev$sg                             : failed to open /dev/sg#: no such file or directory
syz_open_dev$sndhw                          : failed to open /dev/snd/hwC#D#: no such file or directory
syz_open_dev$sndmidi                        : failed to open /dev/snd/midiC#D#: no such file or directory
syz_open_dev$swradio                        : failed to open /dev/swradio#: no such file or directory
syz_open_dev$usbfs                          : failed to open /dev/bus/usb/00#/00#: no such file or directory
syz_open_dev$usbmon                         : failed to open /dev/usbmon#: no such file or directory
syz_open_dev$vbi                            : failed to open /dev/vbi#: no such file or directory
syz_open_dev$vcsa                           : failed to open /dev/vcsa#: no such file or directory
syz_open_dev$vcsn                           : failed to open /dev/vcs#: no such file or directory
syz_open_dev$vcsu                           : failed to open /dev/vcsu#: no such file or directory
syz_open_dev$video4linux                    : failed to open /dev/v4l-subdev#: no such file or directory
syz_open_dev$vivid                          : failed to open /dev/video#: no such file or directory
syz_pkey_set                                : pkey_alloc(0x0, 0x0) failed: function not implemented
syz_usb_connect                             : root failed to open /dev/raw-gadget: no such file or directory
syz_usb_connect$cdc_ecm                     : root failed to open /dev/raw-gadget: no such file or directory
syz_usb_connect$cdc_ncm                     : root failed to open /dev/raw-gadget: no such file or directory
syz_usb_connect$hid                         : root failed to open /dev/raw-gadget: no such file or directory
syz_usb_connect$printer                     : root failed to open /dev/raw-gadget: no such file or directory
syz_usb_connect$uac1                        : root failed to open /dev/raw-gadget: no such file or directory
syz_usb_connect_ath9k                       : root failed to open /dev/raw-gadget: no such file or directory
syz_usb_control_io                          : root failed to open /dev/raw-gadget: no such file or directory
syz_usb_control_io$cdc_ecm                  : root failed to open /dev/raw-gadget: no such file or directory
syz_usb_control_io$cdc_ncm                  : root failed to open /dev/raw-gadget: no such file or directory
syz_usb_control_io$hid                      : root failed to open /dev/raw-gadget: no such file or directory
syz_usb_control_io$printer                  : root failed to open /dev/raw-gadget: no such file or directory
syz_usb_control_io$uac1                     : root failed to open /dev/raw-gadget: no such file or directory
syz_usb_disconnect                          : root failed to open /dev/raw-gadget: no such file or directory
syz_usb_ep_read                             : root failed to open /dev/raw-gadget: no such file or directory
syz_usb_ep_write                            : root failed to open /dev/raw-gadget: no such file or directory
syz_usb_ep_write$ath9k_ep1                  : root failed to open /dev/raw-gadget: no such file or directory
syz_usb_ep_write$ath9k_ep2                  : root failed to open /dev/raw-gadget: no such file or directory
syz_usbip_server_init                       : failed to open /sys/devices/platform/vhci_hcd.0/attach: no such file or directory

transitively disabled the following syscalls (missing resource [creating syscalls]):
accept$alg                                  : sock_alg [socket$alg]
accept$ax25                                 : sock_ax25 [accept$ax25 accept4$ax25 syz_init_net_socket$ax25]
accept$netrom                               : sock_netrom [accept$netrom accept4$netrom syz_init_net_socket$netrom]
accept$nfc_llcp                             : sock_nfc_llcp [accept$nfc_llcp accept4$nfc_llcp syz_init_net_socket$nfc_llcp]
accept$phonet_pipe                          : sock_phonet_pipe [accept$phonet_pipe accept4$phonet_pipe socket$phonet_pipe]
accept4$alg                                 : sock_alg [socket$alg]
accept4$ax25                                : sock_ax25 [accept$ax25 accept4$ax25 syz_init_net_socket$ax25]
accept4$llc                                 : sock_llc [accept4$llc syz_init_net_socket$llc]
accept4$netrom                              : sock_netrom [accept$netrom accept4$netrom syz_init_net_socket$netrom]
accept4$nfc_llcp                            : sock_nfc_llcp [accept$nfc_llcp accept4$nfc_llcp syz_init_net_socket$nfc_llcp]
accept4$phonet_pipe                         : sock_phonet_pipe [accept$phonet_pipe accept4$phonet_pipe socket$phonet_pipe]
accept4$rose                                : sock_rose [accept4$rose syz_init_net_socket$rose]
accept4$tipc                                : sock_tipc [accept4$tipc socket$tipc socketpair$tipc]
accept4$x25                                 : sock_x25 [accept4$x25 syz_init_net_socket$x25]
bind                                        : nfc_dev_id [ioctl$IOCTL_GET_NCIDEV_IDX]
bind$802154_dgram                           : sock_802154_dgram [syz_init_net_socket$802154_dgram]
bind$802154_raw                             : sock_802154_raw [syz_init_net_socket$802154_raw]
bind$alg                                    : sock_alg [socket$alg]
bind$ax25                                   : sock_ax25 [accept$ax25 accept4$ax25 syz_init_net_socket$ax25]
bind$bt_rfcomm                              : sock_bt_rfcomm [syz_init_net_socket$bt_rfcomm]
bind$can_j1939                              : sock_can_j1939 [socket$can_j1939]
bind$can_raw                                : sock_can_raw [socket$can_raw]
bind$isdn                                   : sock_isdn [socket$isdn]
bind$isdn_base                              : sock_isdn_base [socket$isdn_base]
bind$l2tp                                   : sock_l2tp [socket$l2tp]
bind$l2tp6                                  : sock_l2tp6 [socket$l2tp6]
bind$llc                                    : sock_llc [accept4$llc syz_init_net_socket$llc]
bind$netrom                                 : sock_netrom [accept$netrom accept4$netrom syz_init_net_socket$netrom]
bind$nfc_llcp                               : sock_nfc_llcp [accept$nfc_llcp accept4$nfc_llcp syz_init_net_socket$nfc_llcp]
bind$phonet                                 : sock_phonet [accept$phonet_pipe accept4$phonet_pipe socket$phonet socket$phonet_pipe]
bind$pptp                                   : sock_pptp [socket$pptp]
bind$qrtr                                   : sock_qrtr [socket$qrtr]
bind$rds                                    : sock_rds [socket$rds]
bind$rose                                   : sock_rose [accept4$rose syz_init_net_socket$rose]
bind$rxrpc                                  : sock_rxrpc [socket$rxrpc]
bind$tipc                                   : sock_tipc [accept4$tipc socket$tipc socketpair$tipc]
bind$vsock_dgram                            : sock_vsock_dgram [socket$vsock_dgram]
bind$x25                                    : sock_x25 [accept4$x25 syz_init_net_socket$x25]
close$ibv_device                            : fd_rdma [openat$uverbs0]
connect                                     : nfc_dev_id [ioctl$IOCTL_GET_NCIDEV_IDX]
connect$802154_dgram                        : sock_802154_dgram [syz_init_net_socket$802154_dgram]
connect$ax25                                : sock_ax25 [accept$ax25 accept4$ax25 syz_init_net_socket$ax25]
connect$bt_rfcomm                           : sock_bt_rfcomm [syz_init_net_socket$bt_rfcomm]
connect$caif                                : sock_caif [socket$caif_seqpacket socket$caif_stream]
connect$can_bcm                             : sock_can_bcm [socket$can_bcm]
connect$can_j1939                           : sock_can_j1939 [socket$can_j1939]
connect$hf                                  : sock_hf [socket$hf]
connect$l2tp                                : sock_l2tp [socket$l2tp]
connect$l2tp6                               : sock_l2tp6 [socket$l2tp6]
connect$llc                                 : sock_llc [accept4$llc syz_init_net_socket$llc]
connect$netrom                              : sock_netrom [accept$netrom accept4$netrom syz_init_net_socket$netrom]
connect$nfc_llcp                            : sock_nfc_llcp [accept$nfc_llcp accept4$nfc_llcp syz_init_net_socket$nfc_llcp]
connect$nfc_raw                             : sock_nfc_raw [syz_init_net_socket$nfc_raw]
connect$phonet_pipe                         : sock_phonet_pipe [accept$phonet_pipe accept4$phonet_pipe socket$phonet_pipe]
connect$pppl2tp                             : sock_pppl2tp [socket$pppl2tp]
connect$pppoe                               : sock_pppoe [socket$pppoe]
connect$pptp                                : sock_pptp [socket$pptp]
connect$qrtr                                : sock_qrtr [socket$qrtr]
connect$rds                                 : sock_rds [socket$rds]
connect$rose                                : sock_rose [accept4$rose syz_init_net_socket$rose]
connect$rxrpc                               : sock_rxrpc [socket$rxrpc]
connect$tipc                                : sock_tipc [accept4$tipc socket$tipc socketpair$tipc]
connect$vsock_dgram                         : sock_vsock_dgram [socket$vsock_dgram]
connect$x25                                 : sock_x25 [accept4$x25 syz_init_net_socket$x25]
getpeername$ax25                            : sock_ax25 [accept$ax25 accept4$ax25 syz_init_net_socket$ax25]
getpeername$l2tp                            : sock_l2tp [socket$l2tp]
getpeername$l2tp6                           : sock_l2tp6 [socket$l2tp6]
getpeername$llc                             : sock_llc [accept4$llc syz_init_net_socket$llc]
getpeername$netrom                          : sock_netrom [accept$netrom accept4$netrom syz_init_net_socket$netrom]
getpeername$qrtr                            : sock_qrtr [socket$qrtr]
getpeername$tipc                            : sock_tipc [accept4$tipc socket$tipc socketpair$tipc]
getsockname$ax25                            : sock_ax25 [accept$ax25 accept4$ax25 syz_init_net_socket$ax25]
getsockname$l2tp                            : sock_l2tp [socket$l2tp]
getsockname$l2tp6                           : sock_l2tp6 [socket$l2tp6]
getsockname$llc                             : sock_llc [accept4$llc syz_init_net_socket$llc]
getsockname$netrom                          : sock_netrom [accept$netrom accept4$netrom syz_init_net_socket$netrom]
getsockname$qrtr                            : sock_qrtr [socket$qrtr]
getsockname$tipc                            : sock_tipc [accept4$tipc socket$tipc socketpair$tipc]
getsockopt$CAN_RAW_FD_FRAMES                : sock_can_raw [socket$can_raw]
getsockopt$CAN_RAW_FILTER                   : sock_can_raw [socket$can_raw]
getsockopt$CAN_RAW_JOIN_FILTERS             : sock_can_raw [socket$can_raw]
getsockopt$CAN_RAW_LOOPBACK                 : sock_can_raw [socket$can_raw]
getsockopt$CAN_RAW_RECV_OWN_MSGS            : sock_can_raw [socket$can_raw]
getsockopt$MISDN_TIME_STAMP                 : sock_isdn [socket$isdn]
getsockopt$PNPIPE_ENCAP                     : sock_phonet_pipe [accept$phonet_pipe accept4$phonet_pipe socket$phonet_pipe]
getsockopt$PNPIPE_HANDLE                    : sock_phonet_pipe [accept$phonet_pipe accept4$phonet_pipe socket$phonet_pipe]
getsockopt$PNPIPE_IFINDEX                   : sock_phonet_pipe [accept$phonet_pipe accept4$phonet_pipe socket$phonet_pipe]
getsockopt$PNPIPE_INITSTATE                 : sock_phonet_pipe [accept$phonet_pipe accept4$phonet_pipe socket$phonet_pipe]
getsockopt$SO_J1939_ERRQUEUE                : sock_can_j1939 [socket$can_j1939]
getsockopt$SO_J1939_PROMISC                 : sock_can_j1939 [socket$can_j1939]
getsockopt$SO_J1939_SEND_PRIO               : sock_can_j1939 [socket$can_j1939]
getsockopt$TIPC_CONN_TIMEOUT                : sock_tipc [accept4$tipc socket$tipc socketpair$tipc]
getsockopt$TIPC_DEST_DROPPABLE              : sock_tipc [accept4$tipc socket$tipc socketpair$tipc]
getsockopt$TIPC_GROUP_JOIN                  : sock_tipc [accept4$tipc socket$tipc socketpair$tipc]
getsockopt$TIPC_IMPORTANCE                  : sock_tipc [accept4$tipc socket$tipc socketpair$tipc]
getsockopt$TIPC_NODE_RECVQ_DEPTH            : sock_tipc [accept4$tipc socket$tipc socketpair$tipc]
getsockopt$TIPC_SOCK_RECVQ_DEPTH            : sock_tipc [accept4$tipc socket$tipc socketpair$tipc]
getsockopt$TIPC_SRC_DROPPABLE               : sock_tipc [accept4$tipc socket$tipc socketpair$tipc]
getsockopt$WPAN_SECURITY                    : sock_802154_dgram [syz_init_net_socket$802154_dgram]
getsockopt$WPAN_SECURITY_LEVEL              : sock_802154_dgram [syz_init_net_socket$802154_dgram]
getsockopt$WPAN_WANTACK                     : sock_802154_dgram [syz_init_net_socket$802154_dgram]
getsockopt$WPAN_WANTLQI                     : sock_802154_dgram [syz_init_net_socket$802154_dgram]
getsockopt$X25_QBITINCL                     : sock_x25 [accept4$x25 syz_init_net_socket$x25]
getsockopt$ax25_int                         : sock_ax25 [accept$ax25 accept4$ax25 syz_init_net_socket$ax25]
getsockopt$bt_rfcomm_RFCOMM_CONNINFO        : sock_bt_rfcomm [syz_init_net_socket$bt_rfcomm]
getsockopt$bt_rfcomm_RFCOMM_LM              : sock_bt_rfcomm [syz_init_net_socket$bt_rfcomm]
getsockopt$inet6_dccp_buf                   : sock_dccp6 [socket$inet6_dccp]
getsockopt$inet6_dccp_int                   : sock_dccp6 [socket$inet6_dccp]
getsockopt$inet6_mptcp_buf                  : sock_mptcp6 [socket$inet6_mptcp]
getsockopt$inet_dccp_buf                    : sock_dccp [socket$inet_dccp]
getsockopt$inet_dccp_int                    : sock_dccp [socket$inet_dccp]
getsockopt$inet_mptcp_buf                   : sock_mptcp [socket$inet_mptcp]
getsockopt$inet_sctp6_SCTP_ADAPTATION_LAYER : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_ASSOCINFO        : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_AUTH_ACTIVE_KEY  : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_AUTOCLOSE        : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_AUTO_ASCONF      : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_CONTEXT          : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_DEFAULT_PRINFO   : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_DEFAULT_SEND_PARAM: sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_DEFAULT_SNDINFO  : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_DELAYED_SACK     : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_DISABLE_FRAGMENTS: sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_ENABLE_STREAM_RESET: sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_EVENTS           : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_FRAGMENT_INTERLEAVE: sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_GET_ASSOC_ID_LIST: sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_GET_ASSOC_NUMBER : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_GET_ASSOC_STATS  : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_GET_LOCAL_ADDRS  : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_GET_PEER_ADDRS   : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_GET_PEER_ADDR_INFO: sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_HMAC_IDENT       : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_INITMSG          : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_I_WANT_MAPPED_V4_ADDR: sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_LOCAL_AUTH_CHUNKS: sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_MAXSEG           : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_MAX_BURST        : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_NODELAY          : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_PARTIAL_DELIVERY_POINT: sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_PEER_ADDR_PARAMS : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_PEER_ADDR_THLDS  : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_PEER_AUTH_CHUNKS : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_PRIMARY_ADDR     : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_PR_ASSOC_STATUS  : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_PR_STREAM_STATUS : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_PR_SUPPORTED     : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_RECONFIG_SUPPORTED: sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_RECVNXTINFO      : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_RECVRCVINFO      : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_RESET_STREAMS    : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_RTOINFO          : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_SOCKOPT_CONNECTX3: sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_SOCKOPT_PEELOFF  : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_STATUS           : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_STREAM_SCHEDULER : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_STREAM_SCHEDULER_VALUE: sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp_SCTP_ADAPTATION_LAYER  : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_ASSOCINFO         : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_AUTH_ACTIVE_KEY   : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_AUTOCLOSE         : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_AUTO_ASCONF       : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_CONTEXT           : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_DEFAULT_PRINFO    : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_DEFAULT_SEND_PARAM: sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_DEFAULT_SNDINFO   : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_DELAYED_SACK      : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_DISABLE_FRAGMENTS : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_ENABLE_STREAM_RESET: sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_EVENTS            : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_FRAGMENT_INTERLEAVE: sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_GET_ASSOC_ID_LIST : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_GET_ASSOC_NUMBER  : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_GET_ASSOC_STATS   : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_GET_LOCAL_ADDRS   : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_GET_PEER_ADDRS    : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_GET_PEER_ADDR_INFO: sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_HMAC_IDENT        : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_INITMSG           : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_I_WANT_MAPPED_V4_ADDR: sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_LOCAL_AUTH_CHUNKS : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_MAXSEG            : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_MAX_BURST         : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_NODELAY           : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_PARTIAL_DELIVERY_POINT: sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_PEER_ADDR_PARAMS  : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_PEER_ADDR_THLDS   : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_PEER_AUTH_CHUNKS  : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_PRIMARY_ADDR      : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_PR_ASSOC_STATUS   : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_PR_STREAM_STATUS  : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_PR_SUPPORTED      : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_RECONFIG_SUPPORTED: sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_RECVNXTINFO       : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_RECVRCVINFO       : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_RESET_STREAMS     : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_RTOINFO           : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_SOCKOPT_CONNECTX3 : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_SOCKOPT_PEELOFF   : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_STATUS            : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_STREAM_SCHEDULER  : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_STREAM_SCHEDULER_VALUE: sock_sctp [socket$inet_sctp]
getsockopt$kcm_KCM_RECV_DISABLE             : sock_kcm [socket$kcm]
getsockopt$llc_int                          : sock_llc [accept4$llc syz_init_net_socket$llc]
getsockopt$netrom_NETROM_IDLE               : sock_netrom [accept$netrom accept4$netrom syz_init_net_socket$netrom]
getsockopt$netrom_NETROM_N2                 : sock_netrom [accept$netrom accept4$netrom syz_init_net_socket$netrom]
getsockopt$netrom_NETROM_T1                 : sock_netrom [accept$netrom accept4$netrom syz_init_net_socket$netrom]
getsockopt$netrom_NETROM_T2                 : sock_netrom [accept$netrom accept4$netrom syz_init_net_socket$netrom]
getsockopt$netrom_NETROM_T4                 : sock_netrom [accept$netrom accept4$netrom syz_init_net_socket$netrom]
getsockopt$nfc_llcp                         : sock_nfc_llcp [accept$nfc_llcp accept4$nfc_llcp syz_init_net_socket$nfc_llcp]
getsockopt$rose                             : sock_rose [accept4$rose syz_init_net_socket$rose]
ioctl$ACPI_THERMAL_GET_ART                  : fd_acpi_thermal_rel [openat$acpi_thermal_rel]
ioctl$ACPI_THERMAL_GET_ART_COUNT            : fd_acpi_thermal_rel [openat$acpi_thermal_rel]
ioctl$ACPI_THERMAL_GET_ART_LEN              : fd_acpi_thermal_rel [openat$acpi_thermal_rel]
ioctl$ACPI_THERMAL_GET_TRT                  : fd_acpi_thermal_rel [openat$acpi_thermal_rel]
ioctl$ACPI_THERMAL_GET_TRT_COUNT            : fd_acpi_thermal_rel [openat$acpi_thermal_rel]
ioctl$ACPI_THERMAL_GET_TRT_LEN              : fd_acpi_thermal_rel [openat$acpi_thermal_rel]
ioctl$AUTOFS_DEV_IOCTL_ASKUMOUNT            : fd_autofs [openat$autofs]
ioctl$AUTOFS_DEV_IOCTL_CATATONIC            : fd_autofs [openat$autofs]
ioctl$AUTOFS_DEV_IOCTL_CLOSEMOUNT           : fd_autofs [openat$autofs]
ioctl$AUTOFS_DEV_IOCTL_EXPIRE               : fd_autofs [openat$autofs]
ioctl$AUTOFS_DEV_IOCTL_FAIL                 : fd_autofs [openat$autofs]
ioctl$AUTOFS_DEV_IOCTL_ISMOUNTPOINT         : fd_autofs [openat$autofs]
ioctl$AUTOFS_DEV_IOCTL_OPENMOUNT            : fd_autofs [openat$autofs]
ioctl$AUTOFS_DEV_IOCTL_PROTOSUBVER          : fd_autofs [openat$autofs]
ioctl$AUTOFS_DEV_IOCTL_PROTOVER             : fd_autofs [openat$autofs]
ioctl$AUTOFS_DEV_IOCTL_READY                : fd_autofs [openat$autofs]
ioctl$AUTOFS_DEV_IOCTL_REQUESTER            : fd_autofs [openat$autofs]
ioctl$AUTOFS_DEV_IOCTL_SETPIPEFD            : fd_autofs [openat$autofs]
ioctl$AUTOFS_DEV_IOCTL_TIMEOUT              : fd_autofs [openat$autofs]
ioctl$AUTOFS_DEV_IOCTL_VERSION              : fd_autofs [openat$autofs]
ioctl$BLKALIGNOFF                           : fd_block [openat$md openat$nullb openat$pmem0 syz_open_dev$loop syz_open_dev$ndb]
ioctl$BLKBSZGET                             : fd_block [openat$md openat$nullb openat$pmem0 syz_open_dev$loop syz_open_dev$ndb]
ioctl$BLKBSZSET                             : fd_block [openat$md openat$nullb openat$pmem0 syz_open_dev$loop syz_open_dev$ndb]
ioctl$BLKDISCARD                            : fd_block [openat$md openat$nullb openat$pmem0 syz_open_dev$loop syz_open_dev$ndb]
ioctl$BLKFLSBUF                             : fd_block [openat$md openat$nullb openat$pmem0 syz_open_dev$loop syz_open_dev$ndb]
ioctl$BLKFRASET                             : fd_block [openat$md openat$nullb openat$pmem0 syz_open_dev$loop syz_open_dev$ndb]
ioctl$BLKGETDISKSEQ                         : fd_block [openat$md openat$nullb openat$pmem0 syz_open_dev$loop syz_open_dev$ndb]
ioctl$BLKGETSIZE                            : fd_block [openat$md openat$nullb openat$pmem0 syz_open_dev$loop syz_open_dev$ndb]
ioctl$BLKGETSIZE64                          : fd_block [openat$md openat$nullb openat$pmem0 syz_open_dev$loop syz_open_dev$ndb]
ioctl$BLKIOMIN                              : fd_block [openat$md openat$nullb openat$pmem0 syz_open_dev$loop syz_open_dev$ndb]
ioctl$BLKIOOPT                              : fd_block [openat$md openat$nullb openat$pmem0 syz_open_dev$loop syz_open_dev$ndb]
ioctl$BLKPBSZGET                            : fd_block [openat$md openat$nullb openat$pmem0 syz_open_dev$loop syz_open_dev$ndb]
ioctl$BLKPG                                 : fd_block [openat$md openat$nullb openat$pmem0 syz_open_dev$loop syz_open_dev$ndb]
ioctl$BLKRAGET                              : fd_block [openat$md openat$nullb openat$pmem0 syz_open_dev$loop syz_open_dev$ndb]
ioctl$BLKREPORTZONE                         : fd_block [openat$md openat$nullb openat$pmem0 syz_open_dev$loop syz_open_dev$ndb]
ioctl$BLKRESETZONE                          : fd_block [openat$md openat$nullb openat$pmem0 syz_open_dev$loop syz_open_dev$ndb]
ioctl$BLKROGET                              : fd_block [openat$md openat$nullb openat$pmem0 syz_open_dev$loop syz_open_dev$ndb]
ioctl$BLKROSET                              : fd_block [openat$md openat$nullb openat$pmem0 syz_open_dev$loop syz_open_dev$ndb]
ioctl$BLKROTATIONAL                         : fd_block [openat$md openat$nullb openat$pmem0 syz_open_dev$loop syz_open_dev$ndb]
ioctl$BLKRRPART                             : fd_block [openat$md openat$nullb openat$pmem0 syz_open_dev$loop syz_open_dev$ndb]
ioctl$BLKSECDISCARD                         : fd_block [openat$md openat$nullb openat$pmem0 syz_open_dev$loop syz_open_dev$ndb]
ioctl$BLKSECTGET                            : fd_block [openat$md openat$nullb openat$pmem0 syz_open_dev$loop syz_open_dev$ndb]
ioctl$BLKTRACESETUP                         : fd_block_trace [openat$md openat$nullb openat$pmem0 ...]
ioctl$BLKTRACESTART                         : fd_block_trace [openat$md openat$nullb openat$pmem0 ...]
ioctl$BLKTRACESTOP                          : fd_block_trace [openat$md openat$nullb openat$pmem0 ...]
ioctl$BLKTRACETEARDOWN                      : fd_block_trace [openat$md openat$nullb openat$pmem0 ...]
ioctl$BLKZEROOUT                            : fd_block [openat$md openat$nullb openat$pmem0 syz_open_dev$loop syz_open_dev$ndb]
ioctl$CAM_REQ_MGR_ALLOC_BUF                 : fd_camx [openat$camx]
ioctl$CAM_REQ_MGR_CREATE_SESSION            : fd_camx [openat$camx]
ioctl$CAM_REQ_MGR_DESTROY_SESSION           : fd_camx [openat$camx]
ioctl$CAM_REQ_MGR_FLUSH_REQ                 : fd_camx [openat$camx]
ioctl$CAM_REQ_MGR_LINK                      : fd_camx [openat$camx]
ioctl$CAM_REQ_MGR_LINK_CONTROL              : fd_camx [openat$camx]
ioctl$CAM_REQ_MGR_LINK_V2                   : fd_camx [openat$camx]
ioctl$CAM_REQ_MGR_MAP_BUF                   : fd_camx [openat$camx]
ioctl$CAM_REQ_MGR_RELEASE_BUF               : fd_camx [openat$camx]
ioctl$CAM_REQ_MGR_REQUEST_DUMP              : fd_camx [openat$camx]
ioctl$CAM_REQ_MGR_SCHED_REQ                 : fd_camx [openat$camx]
ioctl$CAM_REQ_MGR_SYNC_MODE                 : fd_camx [openat$camx]
ioctl$CAM_REQ_MGR_UNLINK                    : fd_camx [openat$camx]
ioctl$CAPI_CLR_FLAGS                        : fd_capi20 [openat$capi20]
ioctl$CAPI_GET_ERRCODE                      : fd_capi20 [openat$capi20]
ioctl$CAPI_GET_FLAGS                        : fd_capi20 [openat$capi20]
ioctl$CAPI_GET_MANUFACTURER                 : fd_capi20 [openat$capi20]
ioctl$CAPI_GET_PROFILE                      : fd_capi20 [openat$capi20]
ioctl$CAPI_GET_SERIAL                       : fd_capi20 [openat$capi20]
ioctl$CAPI_INSTALLED                        : fd_capi20 [openat$capi20]
ioctl$CAPI_MANUFACTURER_CMD                 : fd_capi20 [openat$capi20]
ioctl$CAPI_NCCI_GETUNIT                     : fd_capi20 [openat$capi20]
ioctl$CAPI_NCCI_OPENCOUNT                   : fd_capi20 [openat$capi20]
ioctl$CAPI_REGISTER                         : fd_capi20 [openat$capi20]
ioctl$CAPI_SET_FLAGS                        : fd_capi20 [openat$capi20]
ioctl$CDROMCLOSETRAY                        : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROMEJECT                            : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROMEJECT_SW                         : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROMGETSPINDOWN                      : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROMMULTISESSION                     : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROMPAUSE                            : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROMPLAYBLK                          : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROMPLAYMSF                          : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROMPLAYTRKIND                       : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROMREADALL                          : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROMREADAUDIO                        : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROMREADCOOKED                       : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROMREADMODE1                        : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROMREADMODE2                        : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROMREADRAW                          : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROMREADTOCENTRY                     : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROMREADTOCHDR                       : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROMRESET                            : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROMRESUME                           : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROMSEEK                             : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROMSETSPINDOWN                      : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROMSTART                            : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROMSTOP                             : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROMSUBCHNL                          : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROMVOLCTRL                          : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROMVOLREAD                          : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROM_CHANGER_NSLOTS                  : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROM_CLEAR_OPTIONS                   : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROM_DEBUG                           : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROM_DISC_STATUS                     : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROM_GET_CAPABILITY                  : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROM_GET_MCN                         : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROM_LAST_WRITTEN                    : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROM_LOCKDOOR                        : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROM_MEDIA_CHANGED                   : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROM_NEXT_WRITABLE                   : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROM_SELECT_DISK                     : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROM_SELECT_SPEED                    : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROM_SEND_PACKET                     : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROM_SET_OPTIONS                     : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROM_TIMED_MEDIA_CHANGE              : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CEC_ADAP_G_CAPS                       : fd_cec [syz_open_dev$cec]
ioctl$CEC_ADAP_G_CONNECTOR_INFO             : fd_cec [syz_open_dev$cec]
ioctl$CEC_ADAP_G_LOG_ADDRS                  : fd_cec [syz_open_dev$cec]
ioctl$CEC_ADAP_G_PHYS_ADDR                  : fd_cec [syz_open_dev$cec]
ioctl$CEC_ADAP_S_LOG_ADDRS                  : fd_cec [syz_open_dev$cec]
ioctl$CEC_ADAP_S_PHYS_ADDR                  : fd_cec [syz_open_dev$cec]
ioctl$CEC_DQEVENT                           : fd_cec [syz_open_dev$cec]
ioctl$CEC_G_MODE                            : fd_cec [syz_open_dev$cec]
ioctl$CEC_RECEIVE                           : fd_cec [syz_open_dev$cec]
ioctl$CEC_S_MODE                            : fd_cec [syz_open_dev$cec]
ioctl$CEC_TRANSMIT                          : fd_cec [syz_open_dev$cec]
ioctl$CREATE_COUNTERS                       : fd_rdma [openat$uverbs0]
ioctl$DESTROY_COUNTERS                      : fd_rdma [openat$uverbs0]
ioctl$DRM_IOCTL_I915_GEM_BUSY               : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_GEM_CONTEXT_CREATE     : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_GEM_CONTEXT_DESTROY    : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_GEM_CONTEXT_GETPARAM   : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_GEM_CONTEXT_SETPARAM   : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_GEM_CREATE             : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_GEM_EXECBUFFER         : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_GEM_EXECBUFFER2        : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_GEM_EXECBUFFER2_WR     : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_GEM_GET_APERTURE       : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_GEM_GET_CACHING        : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_GEM_GET_TILING         : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_GEM_MADVISE            : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_GEM_MMAP               : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_GEM_MMAP_GTT           : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_GEM_MMAP_OFFSET        : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_GEM_PIN                : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_GEM_PREAD              : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_GEM_PWRITE             : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_GEM_SET_CACHING        : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_GEM_SET_DOMAIN         : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_GEM_SET_TILING         : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_GEM_SW_FINISH          : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_GEM_THROTTLE           : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_GEM_UNPIN              : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_GEM_USERPTR            : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_GEM_VM_CREATE          : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_GEM_VM_DESTROY         : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_GEM_WAIT               : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_GETPARAM               : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_GET_PIPE_FROM_CRTC_ID  : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_GET_RESET_STATS        : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_OVERLAY_ATTRS          : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_OVERLAY_PUT_IMAGE      : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_PERF_ADD_CONFIG        : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_PERF_OPEN              : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_PERF_REMOVE_CONFIG     : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_QUERY                  : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_REG_READ               : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_SET_SPRITE_COLORKEY    : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_MSM_GEM_CPU_FINI            : fd_msm [openat$msm]
ioctl$DRM_IOCTL_MSM_GEM_CPU_PREP            : fd_msm [openat$msm]
ioctl$DRM_IOCTL_MSM_GEM_INFO                : fd_msm [openat$msm]
ioctl$DRM_IOCTL_MSM_GEM_MADVISE             : fd_msm [openat$msm]
ioctl$DRM_IOCTL_MSM_GEM_NEW                 : fd_msm [openat$msm]
ioctl$DRM_IOCTL_MSM_GEM_SUBMIT              : fd_msm [openat$msm]
ioctl$DRM_IOCTL_MSM_GET_PARAM               : fd_msm [openat$msm]
ioctl$DRM_IOCTL_MSM_SET_PARAM               : fd_msm [openat$msm]
ioctl$DRM_IOCTL_MSM_SUBMITQUEUE_CLOSE       : fd_msm [openat$msm]
ioctl$DRM_IOCTL_MSM_SUBMITQUEUE_NEW         : fd_msm [openat$msm]
ioctl$DRM_IOCTL_MSM_SUBMITQUEUE_QUERY       : fd_msm [openat$msm]
ioctl$DRM_IOCTL_MSM_WAIT_FENCE              : fd_msm [openat$msm]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_CACHE_CACHEOPEXEC: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_CACHE_CACHEOPLOG: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_CACHE_CACHEOPQUEUE: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_CMM_DEVMEMINTACQUIREREMOTECTX: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_CMM_DEVMEMINTEXPORTCTX: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_CMM_DEVMEMINTUNEXPORTCTX: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_DEVICEMEMHISTORY_DEVICEMEMHISTORYMAP: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_DEVICEMEMHISTORY_DEVICEMEMHISTORYMAPVRANGE: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_DEVICEMEMHISTORY_DEVICEMEMHISTORYSPARSECHANGE: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_DEVICEMEMHISTORY_DEVICEMEMHISTORYUNMAP: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_DEVICEMEMHISTORY_DEVICEMEMHISTORYUNMAPVRANGE: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_DMABUF_PHYSMEMEXPORTDMABUF: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_DMABUF_PHYSMEMIMPORTDMABUF: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_DMABUF_PHYSMEMIMPORTSPARSEDMABUF: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_HTBUFFER_HTBCONTROL: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_HTBUFFER_HTBLOG: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_CHANGESPARSEMEM: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_DEVMEMFLUSHDEVSLCRANGE: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_DEVMEMGETFAULTADDRESS: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_DEVMEMINTCTXCREATE: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_DEVMEMINTCTXDESTROY: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_DEVMEMINTHEAPCREATE: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_DEVMEMINTHEAPDESTROY: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_DEVMEMINTMAPPAGES: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_DEVMEMINTMAPPMR: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_DEVMEMINTPIN: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_DEVMEMINTPINVALIDATE: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_DEVMEMINTREGISTERPFNOTIFYKM: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_DEVMEMINTRESERVERANGE: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_DEVMEMINTUNMAPPAGES: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_DEVMEMINTUNMAPPMR: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_DEVMEMINTUNPIN: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_DEVMEMINTUNPININVALIDATE: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_DEVMEMINTUNRESERVERANGE: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_DEVMEMINVALIDATEFBSCTABLE: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_DEVMEMISVDEVADDRVALID: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_GETMAXDEVMEMSIZE: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_HEAPCFGHEAPCONFIGCOUNT: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_HEAPCFGHEAPCONFIGNAME: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_HEAPCFGHEAPCOUNT: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_HEAPCFGHEAPDETAILS: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_PHYSMEMNEWRAMBACKEDLOCKEDPMR: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_PHYSMEMNEWRAMBACKEDPMR: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_PMREXPORTPMR: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_PMRGETUID: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_PMRIMPORTPMR: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_PMRLOCALIMPORTPMR: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_PMRMAKELOCALIMPORTHANDLE: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_PMRUNEXPORTPMR: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_PMRUNMAKELOCALIMPORTHANDLE: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_PMRUNREFPMR: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_PMRUNREFUNLOCKPMR: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_PVRSRVUPDATEOOMSTATS: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_PVRTL_TLACQUIREDATA: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_PVRTL_TLCLOSESTREAM: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_PVRTL_TLCOMMITSTREAM: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_PVRTL_TLDISCOVERSTREAMS: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_PVRTL_TLOPENSTREAM: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_PVRTL_TLRELEASEDATA: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_PVRTL_TLRESERVESTREAM: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_PVRTL_TLWRITEDATA: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXBREAKPOINT_RGXCLEARBREAKPOINT: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXBREAKPOINT_RGXDISABLEBREAKPOINT: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXBREAKPOINT_RGXENABLEBREAKPOINT: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXBREAKPOINT_RGXOVERALLOCATEBPREGISTERS: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXBREAKPOINT_RGXSETBREAKPOINT: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXCMP_RGXCREATECOMPUTECONTEXT: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXCMP_RGXDESTROYCOMPUTECONTEXT: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXCMP_RGXFLUSHCOMPUTEDATA: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXCMP_RGXGETLASTCOMPUTECONTEXTRESETREASON: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXCMP_RGXKICKCDM2: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXCMP_RGXNOTIFYCOMPUTEWRITEOFFSETUPDATE: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXCMP_RGXSETCOMPUTECONTEXTPRIORITY: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXCMP_RGXSETCOMPUTECONTEXTPROPERTY: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXFWDBG_RGXCURRENTTIME: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXFWDBG_RGXFWDEBUGDUMPFREELISTPAGELIST: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXFWDBG_RGXFWDEBUGPHRCONFIGURE: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXFWDBG_RGXFWDEBUGSETFWLOG: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXFWDBG_RGXFWDEBUGSETHCSDEADLINE: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXFWDBG_RGXFWDEBUGSETOSIDPRIORITY: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXFWDBG_RGXFWDEBUGSETOSNEWONLINESTATE: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXHWPERF_RGXCONFIGCUSTOMCOUNTERS: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXHWPERF_RGXCONFIGENABLEHWPERFCOUNTERS: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXHWPERF_RGXCTRLHWPERF: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXHWPERF_RGXCTRLHWPERFCOUNTERS: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXHWPERF_RGXGETHWPERFBVNCFEATUREFLAGS: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXKICKSYNC_RGXCREATEKICKSYNCCONTEXT: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXKICKSYNC_RGXDESTROYKICKSYNCCONTEXT: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXKICKSYNC_RGXKICKSYNC2: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXKICKSYNC_RGXSETKICKSYNCCONTEXTPROPERTY: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXREGCONFIG_RGXADDREGCONFIG: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXREGCONFIG_RGXCLEARREGCONFIG: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXREGCONFIG_RGXDISABLEREGCONFIG: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXREGCONFIG_RGXENABLEREGCONFIG: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXREGCONFIG_RGXSETREGCONFIGTYPE: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXSIGNALS_RGXNOTIFYSIGNALUPDATE: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXTA3D_RGXCREATEFREELIST: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXTA3D_RGXCREATEHWRTDATASET: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXTA3D_RGXCREATERENDERCONTEXT: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXTA3D_RGXCREATEZSBUFFER: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXTA3D_RGXDESTROYFREELIST: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXTA3D_RGXDESTROYHWRTDATASET: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXTA3D_RGXDESTROYRENDERCONTEXT: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXTA3D_RGXDESTROYZSBUFFER: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXTA3D_RGXGETLASTRENDERCONTEXTRESETREASON: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXTA3D_RGXKICKTA3D2: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXTA3D_RGXPOPULATEZSBUFFER: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXTA3D_RGXRENDERCONTEXTSTALLED: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXTA3D_RGXSETRENDERCONTEXTPRIORITY: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXTA3D_RGXSETRENDERCONTEXTPROPERTY: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXTA3D_RGXUNPOPULATEZSBUFFER: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXTQ2_RGXTDMCREATETRANSFERCONTEXT: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXTQ2_RGXTDMDESTROYTRANSFERCONTEXT: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXTQ2_RGXTDMGETSHAREDMEMORY: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXTQ2_RGXTDMNOTIFYWRITEOFFSETUPDATE: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXTQ2_RGXTDMRELEASESHAREDMEMORY: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXTQ2_RGXTDMSETTRANSFERCONTEXTPRIORITY: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXTQ2_RGXTDMSETTRANSFERCONTEXTPROPERTY: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXTQ2_RGXTDMSUBMITTRANSFER2: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXTQ_RGXCREATETRANSFERCONTEXT: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXTQ_RGXDESTROYTRANSFERCONTEXT: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXTQ_RGXSETTRANSFERCONTEXTPRIORITY: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXTQ_RGXSETTRANSFERCONTEXTPROPERTY: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXTQ_RGXSUBMITTRANSFER2: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_SRVCORE_ACQUIREGLOBALEVENTOBJECT: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_SRVCORE_ACQUIREINFOPAGE: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_SRVCORE_ALIGNMENTCHECK: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_SRVCORE_CONNECT: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_SRVCORE_DISCONNECT: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_SRVCORE_DUMPDEBUGINFO: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_SRVCORE_EVENTOBJECTCLOSE: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_SRVCORE_EVENTOBJECTOPEN: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_SRVCORE_EVENTOBJECTWAIT: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_SRVCORE_EVENTOBJECTWAITTIMEOUT: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_SRVCORE_FINDPROCESSMEMSTATS: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_SRVCORE_GETDEVCLOCKSPEED: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_SRVCORE_GETDEVICESTATUS: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_SRVCORE_GETMULTICOREINFO: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_SRVCORE_HWOPTIMEOUT: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_SRVCORE_RELEASEGLOBALEVENTOBJECT: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_SRVCORE_RELEASEINFOPAGE: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_SYNCTRACKING_SYNCRECORDADD: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_SYNCTRACKING_SYNCRECORDREMOVEBYHANDLE: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_SYNC_ALLOCSYNCPRIMITIVEBLOCK: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_SYNC_FREESYNCPRIMITIVEBLOCK: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_SYNC_SYNCALLOCEVENT: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_SYNC_SYNCCHECKPOINTSIGNALLEDPDUMPPOL: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_SYNC_SYNCFREEEVENT: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_SYNC_SYNCPRIMPDUMP: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_SYNC_SYNCPRIMPDUMPCBP: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_SYNC_SYNCPRIMPDUMPPOL: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_SYNC_SYNCPRIMPDUMPVALUE: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_SYNC_SYNCPRIMSET: fd_rogue [openat$img_rogue]
ioctl$DVD_AUTH                              : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$DVD_READ_STRUCT                       : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$DVD_WRITE_STRUCT                      : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$FBIOBLANK                             : fd_fb [openat$fb0 openat$fb1]
ioctl$FBIOGETCMAP                           : fd_fb [openat$fb0 openat$fb1]
ioctl$FBIOGET_CON2FBMAP                     : fd_fb [openat$fb0 openat$fb1]
ioctl$FBIOGET_FSCREENINFO                   : fd_fb [openat$fb0 openat$fb1]
ioctl$FBIOGET_VSCREENINFO                   : fd_fb [openat$fb0 openat$fb1]
ioctl$FBIOPAN_DISPLAY                       : fd_fb [openat$fb0 openat$fb1]
ioctl$FBIOPUTCMAP                           : fd_fb [openat$fb0 openat$fb1]
ioctl$FBIOPUT_CON2FBMAP                     : fd_fb [openat$fb0 openat$fb1]
ioctl$FBIOPUT_VSCREENINFO                   : fd_fb [openat$fb0 openat$fb1]
ioctl$FBIO_WAITFORVSYNC                     : fd_fb [openat$fb0 openat$fb1]
ioctl$FLOPPY_FDCLRPRM                       : fd_floppy [syz_open_dev$floppy]
ioctl$FLOPPY_FDDEFPRM                       : fd_floppy [syz_open_dev$floppy]
ioctl$FLOPPY_FDEJECT                        : fd_floppy [syz_open_dev$floppy]
ioctl$FLOPPY_FDFLUSH                        : fd_floppy [syz_open_dev$floppy]
ioctl$FLOPPY_FDFMTBEG                       : fd_floppy [syz_open_dev$floppy]
ioctl$FLOPPY_FDFMTEND                       : fd_floppy [syz_open_dev$floppy]
ioctl$FLOPPY_FDFMTTRK                       : fd_floppy [syz_open_dev$floppy]
ioctl$FLOPPY_FDGETDRVPRM                    : fd_floppy [syz_open_dev$floppy]
ioctl$FLOPPY_FDGETDRVSTAT                   : fd_floppy [syz_open_dev$floppy]
ioctl$FLOPPY_FDGETDRVTYP                    : fd_floppy [syz_open_dev$floppy]
ioctl$FLOPPY_FDGETFDCSTAT                   : fd_floppy [syz_open_dev$floppy]
ioctl$FLOPPY_FDGETMAXERRS                   : fd_floppy [syz_open_dev$floppy]
ioctl$FLOPPY_FDGETPRM                       : fd_floppy [syz_open_dev$floppy]
ioctl$FLOPPY_FDMSGOFF                       : fd_floppy [syz_open_dev$floppy]
ioctl$FLOPPY_FDMSGON                        : fd_floppy [syz_open_dev$floppy]
ioctl$FLOPPY_FDPOLLDRVSTAT                  : fd_floppy [syz_open_dev$floppy]
ioctl$FLOPPY_FDRAWCMD                       : fd_floppy [syz_open_dev$floppy]
ioctl$FLOPPY_FDRESET                        : fd_floppy [syz_open_dev$floppy]
ioctl$FLOPPY_FDSETDRVPRM                    : fd_floppy [syz_open_dev$floppy]
ioctl$FLOPPY_FDSETEMSGTRESH                 : fd_floppy [syz_open_dev$floppy]
ioctl$FLOPPY_FDSETMAXERRS                   : fd_floppy [syz_open_dev$floppy]
ioctl$FLOPPY_FDSETPRM                       : fd_floppy [syz_open_dev$floppy]
ioctl$FLOPPY_FDTWADDLE                      : fd_floppy [syz_open_dev$floppy]
ioctl$FLOPPY_FDWERRORCLR                    : fd_floppy [syz_open_dev$floppy]
ioctl$FLOPPY_FDWERRORGET                    : fd_floppy [syz_open_dev$floppy]
ioctl$HDIO_GETGEO                           : fd_block [openat$md openat$nullb openat$pmem0 syz_open_dev$loop syz_open_dev$ndb]
ioctl$HIDIOCAPPLICATION                     : fd_hiddev [syz_open_dev$hiddev]
ioctl$HIDIOCGCOLLECTIONINDEX                : fd_hiddev [syz_open_dev$hiddev]
ioctl$HIDIOCGCOLLECTIONINFO                 : fd_hiddev [syz_open_dev$hiddev]
ioctl$HIDIOCGDEVINFO                        : fd_hiddev [syz_open_dev$hiddev]
ioctl$HIDIOCGFEATURE                        : fd_hidraw [syz_open_dev$hidraw]
ioctl$HIDIOCGFIELDINFO                      : fd_hiddev [syz_open_dev$hiddev]
ioctl$HIDIOCGFLAG                           : fd_hiddev [syz_open_dev$hiddev]
ioctl$HIDIOCGNAME                           : fd_hiddev [syz_open_dev$hiddev]
ioctl$HIDIOCGPHYS                           : fd_hiddev [syz_open_dev$hiddev]
ioctl$HIDIOCGRAWINFO                        : fd_hidraw [syz_open_dev$hidraw]
ioctl$HIDIOCGRAWNAME                        : fd_hidraw [syz_open_dev$hidraw]
ioctl$HIDIOCGRAWPHYS                        : fd_hidraw [syz_open_dev$hidraw]
ioctl$HIDIOCGRDESC                          : fd_hidraw [syz_open_dev$hidraw]
ioctl$HIDIOCGRDESCSIZE                      : fd_hidraw [syz_open_dev$hidraw]
ioctl$HIDIOCGREPORT                         : fd_hiddev [syz_open_dev$hiddev]
ioctl$HIDIOCGREPORTINFO                     : fd_hiddev [syz_open_dev$hiddev]
ioctl$HIDIOCGSTRING                         : fd_hiddev [syz_open_dev$hiddev]
ioctl$HIDIOCGUCODE                          : fd_hiddev [syz_open_dev$hiddev]
ioctl$HIDIOCGUSAGE                          : fd_hiddev [syz_open_dev$hiddev]
ioctl$HIDIOCGUSAGES                         : fd_hiddev [syz_open_dev$hiddev]
ioctl$HIDIOCGVERSION                        : fd_hiddev [syz_open_dev$hiddev]
ioctl$HIDIOCINITREPORT                      : fd_hiddev [syz_open_dev$hiddev]
ioctl$HIDIOCSFEATURE                        : fd_hidraw [syz_open_dev$hidraw]
ioctl$HIDIOCSFLAG                           : fd_hiddev [syz_open_dev$hiddev]
ioctl$HIDIOCSREPORT                         : fd_hiddev [syz_open_dev$hiddev]
ioctl$HIDIOCSUSAGE                          : fd_hiddev [syz_open_dev$hiddev]
ioctl$HIDIOCSUSAGES                         : fd_hiddev [syz_open_dev$hiddev]
ioctl$IMADDTIMER                            : fd_misdntimer [openat$misdntimer]
ioctl$IMCLEAR_L2                            : sock_isdn [socket$isdn]
ioctl$IMCTRLREQ                             : sock_isdn [socket$isdn]
ioctl$IMDELTIMER                            : fd_misdntimer [openat$misdntimer]
ioctl$IMGETCOUNT                            : sock_isdn_base [socket$isdn_base]
ioctl$IMGETDEVINFO                          : sock_isdn_base [socket$isdn_base]
ioctl$IMGETVERSION                          : sock_isdn_base [socket$isdn_base]
ioctl$IMHOLD_L1                             : sock_isdn [socket$isdn]
ioctl$IMSETDEVNAME                          : sock_isdn_base [socket$isdn_base]
ioctl$IOCTL_CONFIG_SYS_RESOURCE_PARAMETERS  : fd_qat [openat$qat_adf_ctl]
ioctl$IOCTL_GET_NCIDEV_IDX                  : fd_nci [openat$nci]
ioctl$IOCTL_GET_NUM_DEVICES                 : fd_qat [openat$qat_adf_ctl]
ioctl$IOCTL_START_ACCEL_DEV                 : fd_qat [openat$qat_adf_ctl]
ioctl$IOCTL_STATUS_ACCEL_DEV                : fd_qat [openat$qat_adf_ctl]
ioctl$IOCTL_STOP_ACCEL_DEV                  : fd_qat [openat$qat_adf_ctl]
ioctl$IOCTL_VMCI_CTX_ADD_NOTIFICATION       : fd_vmci [openat$vmci]
ioctl$IOCTL_VMCI_CTX_GET_CPT_STATE          : fd_vmci [openat$vmci]
ioctl$IOCTL_VMCI_CTX_REMOVE_NOTIFICATION    : fd_vmci [openat$vmci]
ioctl$IOCTL_VMCI_CTX_SET_CPT_STATE          : fd_vmci [openat$vmci]
ioctl$IOCTL_VMCI_DATAGRAM_RECEIVE           : fd_vmci [openat$vmci]
ioctl$IOCTL_VMCI_DATAGRAM_SEND              : fd_vmci [openat$vmci]
ioctl$IOCTL_VMCI_GET_CONTEXT_ID             : fd_vmci [openat$vmci]
ioctl$IOCTL_VMCI_INIT_CONTEXT               : fd_vmci [openat$vmci]
ioctl$IOCTL_VMCI_NOTIFICATIONS_RECEIVE      : fd_vmci [openat$vmci]
ioctl$IOCTL_VMCI_NOTIFY_RESOURCE            : fd_vmci [openat$vmci]
ioctl$IOCTL_VMCI_QUEUEPAIR_ALLOC            : fd_vmci [openat$vmci]
ioctl$IOCTL_VMCI_QUEUEPAIR_DETACH           : fd_vmci [openat$vmci]
ioctl$IOCTL_VMCI_QUEUEPAIR_SETPF            : fd_vmci [openat$vmci]
ioctl$IOCTL_VMCI_QUEUEPAIR_SETVA            : fd_vmci [openat$vmci]
ioctl$IOCTL_VMCI_SET_NOTIFY                 : fd_vmci [openat$vmci]
ioctl$IOCTL_VMCI_VERSION                    : fd_vmci [openat$vmci]
ioctl$IOCTL_VMCI_VERSION2                   : fd_vmci [openat$vmci]
ioctl$IOC_PR_CLEAR                          : fd_block [openat$md openat$nullb openat$pmem0 syz_open_dev$loop syz_open_dev$ndb]
ioctl$IOC_PR_PREEMPT                        : fd_block [openat$md openat$nullb openat$pmem0 syz_open_dev$loop syz_open_dev$ndb]
ioctl$IOC_PR_PREEMPT_ABORT                  : fd_block [openat$md openat$nullb openat$pmem0 syz_open_dev$loop syz_open_dev$ndb]
ioctl$IOC_PR_REGISTER                       : fd_block [openat$md openat$nullb openat$pmem0 syz_open_dev$loop syz_open_dev$ndb]
ioctl$IOC_PR_RELEASE                        : fd_block [openat$md openat$nullb openat$pmem0 syz_open_dev$loop syz_open_dev$ndb]
ioctl$IOC_PR_RESERVE                        : fd_block [openat$md openat$nullb openat$pmem0 syz_open_dev$loop syz_open_dev$ndb]
ioctl$IOMMU_DESTROY$hwpt                    : fd_iommufd [openat$iommufd]
ioctl$IOMMU_DESTROY$ioas                    : fd_iommufd [openat$iommufd]
ioctl$IOMMU_DESTROY$stdev                   : fd_iommufd [openat$iommufd]
ioctl$IOMMU_GET_HW_INFO                     : fd_iommufd [openat$iommufd]
ioctl$IOMMU_HWPT_ALLOC$NONE                 : fd_iommufd [openat$iommufd]
ioctl$IOMMU_HWPT_ALLOC$TEST                 : fd_iommufd [openat$iommufd]
ioctl$IOMMU_HWPT_GET_DIRTY_BITMAP           : fd_iommufd [openat$iommufd]
ioctl$IOMMU_HWPT_INVALIDATE$TEST            : fd_iommufd [openat$iommufd]
ioctl$IOMMU_HWPT_SET_DIRTY_TRACKING         : fd_iommufd [openat$iommufd]
ioctl$IOMMU_IOAS_ALLOC                      : fd_iommufd [openat$iommufd]
ioctl$IOMMU_IOAS_ALLOW_IOVAS                : fd_iommufd [openat$iommufd]
ioctl$IOMMU_IOAS_COPY                       : fd_iommufd [openat$iommufd]
ioctl$IOMMU_IOAS_COPY$syz                   : fd_iommufd [openat$iommufd]
ioctl$IOMMU_IOAS_IOVA_RANGES                : fd_iommufd [openat$iommufd]
ioctl$IOMMU_IOAS_MAP                        : fd_iommufd [openat$iommufd]
ioctl$IOMMU_IOAS_MAP$PAGES                  : fd_iommufd [openat$iommufd]
ioctl$IOMMU_IOAS_UNMAP                      : fd_iommufd [openat$iommufd]
ioctl$IOMMU_IOAS_UNMAP$ALL                  : fd_iommufd [openat$iommufd]
ioctl$IOMMU_OPTION$IOMMU_OPTION_HUGE_PAGES  : fd_iommufd [openat$iommufd]
ioctl$IOMMU_OPTION$IOMMU_OPTION_RLIMIT_MODE : fd_iommufd [openat$iommufd]
ioctl$IOMMU_TEST_OP_ACCESS_PAGES            : fd_iommufd [openat$iommufd]
ioctl$IOMMU_TEST_OP_ACCESS_PAGES$syz        : fd_iommufd [openat$iommufd]
ioctl$IOMMU_TEST_OP_ACCESS_REPLACE_IOAS     : fd_iommufd [openat$iommufd]
ioctl$IOMMU_TEST_OP_ACCESS_RW               : fd_iommufd [openat$iommufd]
ioctl$IOMMU_TEST_OP_ACCESS_RW$syz           : fd_iommufd [openat$iommufd]
ioctl$IOMMU_TEST_OP_ADD_RESERVED            : fd_iommufd [openat$iommufd]
ioctl$IOMMU_TEST_OP_CREATE_ACCESS           : fd_iommufd [openat$iommufd]
ioctl$IOMMU_TEST_OP_DESTROY_ACCESS_PAGES    : fd_iommufd [openat$iommufd]
ioctl$IOMMU_TEST_OP_MD_CHECK_MAP            : fd_iommufd [openat$iommufd]
ioctl$IOMMU_TEST_OP_MD_CHECK_REFS           : fd_iommufd [openat$iommufd]
ioctl$IOMMU_TEST_OP_MOCK_DOMAIN             : fd_iommufd [openat$iommufd]
ioctl$IOMMU_TEST_OP_MOCK_DOMAIN_FLAGS       : fd_iommufd [openat$iommufd]
ioctl$IOMMU_TEST_OP_MOCK_DOMAIN_REPLACE     : fd_iommufd [openat$iommufd]
ioctl$IOMMU_TEST_OP_SET_TEMP_MEMORY_LIMIT   : fd_iommufd [openat$iommufd]
ioctl$IOMMU_VFIO_CHECK_EXTENSION            : fd_iommufd [openat$iommufd]
ioctl$IOMMU_VFIO_GET_API_VERSION            : fd_iommufd [openat$iommufd]
ioctl$IOMMU_VFIO_IOAS$CLEAR                 : fd_iommufd [openat$iommufd]
ioctl$IOMMU_VFIO_IOAS$GET                   : fd_iommufd [openat$iommufd]
ioctl$IOMMU_VFIO_IOAS$SET                   : fd_iommufd [openat$iommufd]
ioctl$IOMMU_VFIO_IOMMU_GET_INFO             : fd_iommufd [openat$iommufd]
ioctl$IOMMU_VFIO_IOMMU_MAP_DMA              : fd_iommufd [openat$iommufd]
ioctl$IOMMU_VFIO_IOMMU_UNMAP_DMA            : fd_iommufd [openat$iommufd]
ioctl$IOMMU_VFIO_SET_IOMMU                  : fd_iommufd [openat$iommufd]
ioctl$KVM_ARM_PREFERRED_TARGET              : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_ARM_SET_COUNTER_OFFSET            : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_ARM_SET_DEVICE_ADDR               : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_ARM_VCPU_FINALIZE                 : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_ARM_VCPU_INIT                     : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_ASSIGN_SET_MSIX_ENTRY             : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_ASSIGN_SET_MSIX_NR                : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_ARM_EAGER_SPLIT_CHUNK_SIZE    : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_ARM_INJECT_SERROR_ESR         : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_ARM_MTE                       : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_ARM_SYSTEM_SUSPEND            : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_ARM_USER_IRQ                  : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_DIRTY_LOG_RING                : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_DIRTY_LOG_RING_ACQ_REL        : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_DISABLE_QUIRKS                : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_DISABLE_QUIRKS2               : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_ENFORCE_PV_FEATURE_CPUID      : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_CAP_EXCEPTION_PAYLOAD             : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_EXIT_HYPERCALL                : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_EXIT_ON_EMULATION_FAILURE     : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_HALT_POLL                     : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_HYPERV_DIRECT_TLBFLUSH        : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_CAP_HYPERV_ENFORCE_CPUID          : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_CAP_HYPERV_ENLIGHTENED_VMCS       : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_CAP_HYPERV_SEND_IPI               : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_HYPERV_SYNIC                  : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_CAP_HYPERV_SYNIC2                 : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_CAP_HYPERV_TLBFLUSH               : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_HYPERV_VP_INDEX               : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_MANUAL_DIRTY_LOG_PROTECT2     : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_MAX_VCPU_ID                   : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_MEMORY_FAULT_INFO             : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_MSR_PLATFORM_INFO             : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_PMU_CAPABILITY                : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_PTP_KVM                       : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_SGX_ATTRIBUTE                 : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_SPLIT_IRQCHIP                 : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_STEAL_TIME                    : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_SYNC_REGS                     : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_CAP_VM_COPY_ENC_CONTEXT_FROM      : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_VM_DISABLE_NX_HUGE_PAGES      : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_VM_MOVE_ENC_CONTEXT_FROM      : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_VM_TYPES                      : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_X2APIC_API                    : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_X86_APIC_BUS_CYCLES_NS        : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_X86_BUS_LOCK_EXIT             : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_X86_DISABLE_EXITS             : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_X86_GUEST_MODE                : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_X86_NOTIFY_VMEXIT             : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_X86_USER_SPACE_MSR            : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_XEN_HVM                       : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CHECK_EXTENSION                   : fd_kvm [openat$kvm]
ioctl$KVM_CHECK_EXTENSION_VM                : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CLEAR_DIRTY_LOG                   : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CREATE_DEVICE                     : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CREATE_GUEST_MEMFD                : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CREATE_IRQCHIP                    : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CREATE_PIT2                       : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CREATE_VCPU                       : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CREATE_VM                         : fd_kvm [openat$kvm]
ioctl$KVM_DIRTY_TLB                         : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_GET_API_VERSION                   : fd_kvm [openat$kvm]
ioctl$KVM_GET_CLOCK                         : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_GET_DEVICE_ATTR                   : fd_kvmdev [ioctl$KVM_CREATE_DEVICE syz_kvm_vgic_v3_setup]
ioctl$KVM_GET_DEVICE_ATTR_vcpu              : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_GET_DEVICE_ATTR_vm                : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_GET_DIRTY_LOG                     : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_GET_FPU                           : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_GET_IRQCHIP                       : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_GET_MP_STATE                      : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_GET_NR_MMU_PAGES                  : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_GET_ONE_REG                       : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_GET_REGS                          : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_GET_REG_LIST                      : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_GET_SREGS                         : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_GET_TSC_KHZ                       : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_GET_VCPU_EVENTS                   : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_GET_VCPU_MMAP_SIZE                : fd_kvm [openat$kvm]
ioctl$KVM_HAS_DEVICE_ATTR                   : fd_kvmdev [ioctl$KVM_CREATE_DEVICE syz_kvm_vgic_v3_setup]
ioctl$KVM_HAS_DEVICE_ATTR_vcpu              : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_HAS_DEVICE_ATTR_vm                : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_INTERRUPT                         : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_IOEVENTFD                         : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_IRQFD                             : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_IRQ_LINE                          : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_IRQ_LINE_STATUS                   : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_KVMCLOCK_CTRL                     : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_NMI                               : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_PPC_ALLOCATE_HTAB                 : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_PRE_FAULT_MEMORY                  : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_REGISTER_COALESCED_MMIO           : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_REINJECT_CONTROL                  : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_RESET_DIRTY_RINGS                 : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_RUN                               : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_S390_VCPU_FAULT                   : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_SET_BOOT_CPU_ID                   : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_SET_CLOCK                         : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_SET_DEVICE_ATTR                   : fd_kvmdev [ioctl$KVM_CREATE_DEVICE syz_kvm_vgic_v3_setup]
ioctl$KVM_SET_DEVICE_ATTR_vcpu              : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_SET_DEVICE_ATTR_vm                : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_SET_FPU                           : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_SET_GSI_ROUTING                   : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_SET_GUEST_DEBUG                   : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_SET_IDENTITY_MAP_ADDR             : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_SET_IRQCHIP                       : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_SET_MEMORY_ATTRIBUTES             : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_SET_MP_STATE                      : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_SET_NR_MMU_PAGES                  : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_SET_ONE_REG                       : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_SET_REGS                          : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_SET_SIGNAL_MASK                   : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_SET_SREGS                         : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_SET_TSC_KHZ                       : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_SET_TSS_ADDR                      : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_SET_USER_MEMORY_REGION            : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_SET_USER_MEMORY_REGION2           : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_SET_VAPIC_ADDR                    : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_SET_VCPU_EVENTS                   : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_SIGNAL_MSI                        : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_SMI                               : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_TPR_ACCESS_REPORTING              : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_TRANSLATE                         : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_UNREGISTER_COALESCED_MMIO         : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_X86_GET_MCE_CAP_SUPPORTED         : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_X86_SETUP_MCE                     : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$LOOP_CHANGE_FD                        : fd_loop [syz_open_dev$loop]
ioctl$LOOP_CLR_FD                           : fd_loop [syz_open_dev$loop]
ioctl$LOOP_CONFIGURE                        : fd_loop [syz_open_dev$loop]
ioctl$LOOP_GET_STATUS                       : fd_loop [syz_open_dev$loop]
ioctl$LOOP_GET_STATUS64                     : fd_loop [syz_open_dev$loop]
ioctl$LOOP_SET_BLOCK_SIZE                   : fd_loop [syz_open_dev$loop]
ioctl$LOOP_SET_CAPACITY                     : fd_loop [syz_open_dev$loop]
ioctl$LOOP_SET_DIRECT_IO                    : fd_loop [syz_open_dev$loop]
ioctl$LOOP_SET_FD                           : fd_loop [syz_open_dev$loop]
ioctl$LOOP_SET_STATUS                       : fd_loop [syz_open_dev$loop]
ioctl$LOOP_SET_STATUS64                     : fd_loop [syz_open_dev$loop]
ioctl$MON_IOCG_STATS                        : fd_usbmon [syz_open_dev$usbmon]
ioctl$MON_IOCH_MFLUSH                       : fd_usbmon [syz_open_dev$usbmon]
ioctl$MON_IOCQ_RING_SIZE                    : fd_usbmon [syz_open_dev$usbmon]
ioctl$MON_IOCQ_URB_LEN                      : fd_usbmon [syz_open_dev$usbmon]
ioctl$MON_IOCT_RING_SIZE                    : fd_usbmon [syz_open_dev$usbmon]
ioctl$MON_IOCX_GET                          : fd_usbmon [syz_open_dev$usbmon]
ioctl$MON_IOCX_GETX                         : fd_usbmon [syz_open_dev$usbmon]
ioctl$MON_IOCX_MFETCH                       : fd_usbmon [syz_open_dev$usbmon]
ioctl$NBD_CLEAR_QUE                         : fd_nbd [syz_open_dev$ndb]
ioctl$NBD_CLEAR_SOCK                        : fd_nbd [syz_open_dev$ndb]
ioctl$NBD_DISCONNECT                        : fd_nbd [syz_open_dev$ndb]
ioctl$NBD_DO_IT                             : fd_nbd [syz_open_dev$ndb]
ioctl$NBD_PRINT_DEBUG                       : fd_nbd [syz_open_dev$ndb]
ioctl$NBD_SET_BLKSIZE                       : fd_nbd [syz_open_dev$ndb]
ioctl$NBD_SET_FLAGS                         : fd_nbd [syz_open_dev$ndb]
ioctl$NBD_SET_SIZE                          : fd_nbd [syz_open_dev$ndb]
ioctl$NBD_SET_SIZE_BLOCKS                   : fd_nbd [syz_open_dev$ndb]
ioctl$NBD_SET_SOCK                          : fd_nbd [syz_open_dev$ndb]
ioctl$NBD_SET_TIMEOUT                       : fd_nbd [syz_open_dev$ndb]
ioctl$PPPIOCATTACH                          : fd_ppp [openat$ppp]
ioctl$PPPIOCATTCHAN                         : fd_ppp [openat$ppp]
ioctl$PPPIOCBRIDGECHAN                      : fd_ppp [openat$ppp]
ioctl$PPPIOCCONNECT                         : fd_ppp [openat$ppp]
ioctl$PPPIOCDISCONN                         : fd_ppp [openat$ppp]
ioctl$PPPIOCGCHAN                           : sock_pppox [socket$pppl2tp socket$pppoe socket$pptp]
ioctl$PPPIOCGDEBUG                          : fd_ppp [openat$ppp]
ioctl$PPPIOCGFLAGS                          : sock_pppox [socket$pppl2tp socket$pppoe socket$pptp]
ioctl$PPPIOCGFLAGS1                         : fd_ppp [openat$ppp]
ioctl$PPPIOCGIDLE                           : fd_ppp [openat$ppp]
ioctl$PPPIOCGIDLE32                         : fd_ppp [openat$ppp]
ioctl$PPPIOCGIDLE64                         : fd_ppp [openat$ppp]
ioctl$PPPIOCGL2TPSTATS                      : sock_pppl2tp [socket$pppl2tp]
ioctl$PPPIOCGMRU                            : sock_pppox [socket$pppl2tp socket$pppoe socket$pptp]
ioctl$PPPIOCGNPMODE                         : fd_ppp [openat$ppp]
ioctl$PPPIOCGUNIT                           : fd_ppp [openat$ppp]
ioctl$PPPIOCNEWUNIT                         : fd_ppp [openat$ppp]
ioctl$PPPIOCSACTIVE                         : fd_ppp [openat$ppp]
ioctl$PPPIOCSCOMPRESS                       : fd_ppp [openat$ppp]
ioctl$PPPIOCSDEBUG                          : fd_ppp [openat$ppp]
ioctl$PPPIOCSFLAGS                          : sock_pppox [socket$pppl2tp socket$pppoe socket$pptp]
ioctl$PPPIOCSFLAGS1                         : fd_ppp [openat$ppp]
ioctl$PPPIOCSMAXCID                         : fd_ppp [openat$ppp]
ioctl$PPPIOCSMRRU                           : fd_ppp [openat$ppp]
ioctl$PPPIOCSMRU                            : sock_pppox [socket$pppl2tp socket$pppoe socket$pptp]
ioctl$PPPIOCSMRU1                           : fd_ppp [openat$ppp]
ioctl$PPPIOCSNPMODE                         : fd_ppp [openat$ppp]
ioctl$PPPIOCSPASS                           : fd_ppp [openat$ppp]
ioctl$PPPIOCUNBRIDGECHAN                    : fd_ppp [openat$ppp]
ioctl$PPPOEIOCDFWD                          : sock_pppoe [socket$pppoe]
ioctl$PPPOEIOCSFWD                          : sock_pppoe [socket$pppoe]
ioctl$PTP_CLOCK_GETCAPS                     : fd_ptp [openat$ptp0 openat$ptp1]
ioctl$PTP_ENABLE_PPS                        : fd_ptp [openat$ptp0 openat$ptp1]
ioctl$PTP_EXTTS_REQUEST                     : fd_ptp [openat$ptp0 openat$ptp1]
ioctl$PTP_EXTTS_REQUEST2                    : fd_ptp [openat$ptp0 openat$ptp1]
ioctl$PTP_PEROUT_REQUEST                    : fd_ptp [openat$ptp0 openat$ptp1]
ioctl$PTP_PEROUT_REQUEST2                   : fd_ptp [openat$ptp0 openat$ptp1]
ioctl$PTP_PIN_GETFUNC                       : fd_ptp [openat$ptp0 openat$ptp1]
ioctl$PTP_PIN_GETFUNC2                      : fd_ptp [openat$ptp0 openat$ptp1]
ioctl$PTP_PIN_SETFUNC                       : fd_ptp [openat$ptp0 openat$ptp1]
ioctl$PTP_PIN_SETFUNC2                      : fd_ptp [openat$ptp0 openat$ptp1]
ioctl$PTP_SYS_OFFSET                        : fd_ptp [openat$ptp0 openat$ptp1]
ioctl$PTP_SYS_OFFSET_EXTENDED               : fd_ptp [openat$ptp0 openat$ptp1]
ioctl$PTP_SYS_OFFSET_PRECISE                : fd_ptp [openat$ptp0 openat$ptp1]
ioctl$READ_COUNTERS                         : fd_rdma [openat$uverbs0]
ioctl$SCSI_IOCTL_BENCHMARK_COMMAND          : fd_sg [syz_open_dev$sg]
ioctl$SCSI_IOCTL_DOORLOCK                   : fd_sg [syz_open_dev$sg]
ioctl$SCSI_IOCTL_DOORUNLOCK                 : fd_sg [syz_open_dev$sg]
ioctl$SCSI_IOCTL_GET_BUS_NUMBER             : fd_sg [syz_open_dev$sg]
ioctl$SCSI_IOCTL_GET_IDLUN                  : fd_sg [syz_open_dev$sg]
ioctl$SCSI_IOCTL_GET_PCI                    : fd_sg [syz_open_dev$sg]
ioctl$SCSI_IOCTL_PROBE_HOST                 : fd_sg [syz_open_dev$sg]
ioctl$SCSI_IOCTL_SEND_COMMAND               : fd_sg [syz_open_dev$sg]
ioctl$SCSI_IOCTL_START_UNIT                 : fd_sg [syz_open_dev$sg]
ioctl$SCSI_IOCTL_STOP_UNIT                  : fd_sg [syz_open_dev$sg]
ioctl$SCSI_IOCTL_SYNC                       : fd_sg [syz_open_dev$sg]
ioctl$SCSI_IOCTL_TEST_UNIT_READY            : fd_sg [syz_open_dev$sg]
ioctl$SG_BLKSECTGET                         : fd_sg [syz_open_dev$sg]
ioctl$SG_BLKTRACESETUP                      : fd_sg [syz_open_dev$sg]
ioctl$SG_BLKTRACESTART                      : fd_sg [syz_open_dev$sg]
ioctl$SG_BLKTRACESTOP                       : fd_sg [syz_open_dev$sg]
ioctl$SG_BLKTRACETEARDOWN                   : fd_sg [syz_open_dev$sg]
ioctl$SG_EMULATED_HOST                      : fd_sg [syz_open_dev$sg]
ioctl$SG_GET_ACCESS_COUNT                   : fd_sg [syz_open_dev$sg]
ioctl$SG_GET_COMMAND_Q                      : fd_sg [syz_open_dev$sg]
ioctl$SG_GET_KEEP_ORPHAN                    : fd_sg [syz_open_dev$sg]
ioctl$SG_GET_LOW_DMA                        : fd_sg [syz_open_dev$sg]
ioctl$SG_GET_NUM_WAITING                    : fd_sg [syz_open_dev$sg]
ioctl$SG_GET_PACK_ID                        : fd_sg [syz_open_dev$sg]
ioctl$SG_GET_REQUEST_TABLE                  : fd_sg [syz_open_dev$sg]
ioctl$SG_GET_RESERVED_SIZE                  : fd_sg [syz_open_dev$sg]
ioctl$SG_GET_SCSI_ID                        : fd_sg [syz_open_dev$sg]
ioctl$SG_GET_SG_TABLESIZE                   : fd_sg [syz_open_dev$sg]
ioctl$SG_GET_TIMEOUT                        : fd_sg [syz_open_dev$sg]
ioctl$SG_GET_VERSION_NUM                    : fd_sg [syz_open_dev$sg]
ioctl$SG_IO                                 : fd_sg [syz_open_dev$sg]
ioctl$SG_NEXT_CMD_LEN                       : fd_sg [syz_open_dev$sg]
ioctl$SG_SCSI_RESET                         : fd_sg [syz_open_dev$sg]
ioctl$SG_SET_COMMAND_Q                      : fd_sg [syz_open_dev$sg]
ioctl$SG_SET_DEBUG                          : fd_sg [syz_open_dev$sg]
ioctl$SG_SET_FORCE_PACK_ID                  : fd_sg [syz_open_dev$sg]
ioctl$SG_SET_KEEP_ORPHAN                    : fd_sg [syz_open_dev$sg]
ioctl$SG_SET_RESERVED_SIZE                  : fd_sg [syz_open_dev$sg]
ioctl$SG_SET_TIMEOUT                        : fd_sg [syz_open_dev$sg]
ioctl$SIOCAX25ADDFWD                        : sock_ax25 [accept$ax25 accept4$ax25 syz_init_net_socket$ax25]
ioctl$SIOCAX25ADDUID                        : sock_ax25 [accept$ax25 accept4$ax25 syz_init_net_socket$ax25]
ioctl$SIOCAX25CTLCON                        : sock_ax25 [accept$ax25 accept4$ax25 syz_init_net_socket$ax25]
ioctl$SIOCAX25DELFWD                        : sock_ax25 [accept$ax25 accept4$ax25 syz_init_net_socket$ax25]
ioctl$SIOCAX25DELUID                        : sock_ax25 [accept$ax25 accept4$ax25 syz_init_net_socket$ax25]
ioctl$SIOCAX25GETINFO                       : sock_ax25 [accept$ax25 accept4$ax25 syz_init_net_socket$ax25]
ioctl$SIOCAX25GETINFOOLD                    : sock_ax25 [accept$ax25 accept4$ax25 syz_init_net_socket$ax25]
ioctl$SIOCAX25GETUID                        : sock_ax25 [accept$ax25 accept4$ax25 syz_init_net_socket$ax25]
ioctl$SIOCAX25NOUID                         : sock_ax25 [accept$ax25 accept4$ax25 syz_init_net_socket$ax25]
ioctl$SIOCAX25OPTRT                         : sock_ax25 [accept$ax25 accept4$ax25 syz_init_net_socket$ax25]
ioctl$SIOCGETLINKNAME                       : sock_tipc [accept4$tipc socket$tipc socketpair$tipc]
ioctl$SIOCGETNODEID                         : sock_tipc [accept4$tipc socket$tipc socketpair$tipc]
ioctl$SIOCGIFMTU                            : sock_pppl2tp [socket$pppl2tp]
ioctl$SIOCNRDECOBS                          : sock_netrom [accept$netrom accept4$netrom syz_init_net_socket$netrom]
ioctl$SIOCPNADDRESOURCE                     : sock_phonet [accept$phonet_pipe accept4$phonet_pipe socket$phonet socket$phonet_pipe]
ioctl$SIOCPNDELRESOURCE                     : sock_phonet_dgram [socket$phonet]
ioctl$SIOCPNENABLEPIPE                      : sock_phonet_pipe [accept$phonet_pipe accept4$phonet_pipe socket$phonet_pipe]
ioctl$SIOCPNGETOBJECT                       : sock_phonet [accept$phonet_pipe accept4$phonet_pipe socket$phonet socket$phonet_pipe]
ioctl$SIOCRSACCEPT                          : sock_rose [accept4$rose syz_init_net_socket$rose]
ioctl$SIOCRSGCAUSE                          : sock_rose [accept4$rose syz_init_net_socket$rose]
ioctl$SIOCRSGL2CALL                         : sock_rose [accept4$rose syz_init_net_socket$rose]
ioctl$SIOCRSSCAUSE                          : sock_rose [accept4$rose syz_init_net_socket$rose]
ioctl$SIOCRSSL2CALL                         : sock_rose [accept4$rose syz_init_net_socket$rose]
ioctl$SIOCSIFMTU                            : sock_pppl2tp [socket$pppl2tp]
ioctl$SIOCX25CALLACCPTAPPRV                 : sock_x25 [accept4$x25 syz_init_net_socket$x25]
ioctl$SIOCX25GCALLUSERDATA                  : sock_x25 [accept4$x25 syz_init_net_socket$x25]
ioctl$SIOCX25GCAUSEDIAG                     : sock_x25 [accept4$x25 syz_init_net_socket$x25]
ioctl$SIOCX25GDTEFACILITIES                 : sock_x25 [accept4$x25 syz_init_net_socket$x25]
ioctl$SIOCX25GFACILITIES                    : sock_x25 [accept4$x25 syz_init_net_socket$x25]
ioctl$SIOCX25GSUBSCRIP                      : sock_x25 [accept4$x25 syz_init_net_socket$x25]
ioctl$SIOCX25SCALLUSERDATA                  : sock_x25 [accept4$x25 syz_init_net_socket$x25]
ioctl$SIOCX25SCAUSEDIAG                     : sock_x25 [accept4$x25 syz_init_net_socket$x25]
ioctl$SIOCX25SCUDMATCHLEN                   : sock_x25 [accept4$x25 syz_init_net_socket$x25]
ioctl$SIOCX25SDTEFACILITIES                 : sock_x25 [accept4$x25 syz_init_net_socket$x25]
ioctl$SIOCX25SENDCALLACCPT                  : sock_x25 [accept4$x25 syz_init_net_socket$x25]
ioctl$SIOCX25SFACILITIES                    : sock_x25 [accept4$x25 syz_init_net_socket$x25]
ioctl$SIOCX25SSUBSCRIP                      : sock_x25 [accept4$x25 syz_init_net_socket$x25]
ioctl$SNAPSHOT_ALLOC_SWAP_PAGE              : fd_snapshot [openat$snapshot]
ioctl$SNAPSHOT_ATOMIC_RESTORE               : fd_snapshot [openat$snapshot]
ioctl$SNAPSHOT_AVAIL_SWAP_SIZE              : fd_snapshot [openat$snapshot]
ioctl$SNAPSHOT_CREATE_IMAGE                 : fd_snapshot [openat$snapshot]
ioctl$SNAPSHOT_FREE                         : fd_snapshot [openat$snapshot]
ioctl$SNAPSHOT_FREE_SWAP_PAGES              : fd_snapshot [openat$snapshot]
ioctl$SNAPSHOT_GET_IMAGE_SIZE               : fd_snapshot [openat$snapshot]
ioctl$SNAPSHOT_PLATFORM_SUPPORT             : fd_snapshot [openat$snapshot]
ioctl$SNAPSHOT_PREF_IMAGE_SIZE              : fd_snapshot [openat$snapshot]
ioctl$SNAPSHOT_S2RAM                        : fd_snapshot [openat$snapshot]
ioctl$SNAPSHOT_SET_SWAP_AREA                : fd_snapshot [openat$snapshot]
ioctl$SNAPSHOT_UNFREEZE                     : fd_snapshot [openat$snapshot]
ioctl$SNDCTL_DSP_CHANNELS                   : fd_dsp [openat$adsp1 openat$audio openat$audio1 openat$dsp openat$dsp1]
ioctl$SNDCTL_DSP_GETBLKSIZE                 : fd_dsp [openat$adsp1 openat$audio openat$audio1 openat$dsp openat$dsp1]
ioctl$SNDCTL_DSP_GETCAPS                    : fd_dsp [openat$adsp1 openat$audio openat$audio1 openat$dsp openat$dsp1]
ioctl$SNDCTL_DSP_GETFMTS                    : fd_dsp [openat$adsp1 openat$audio openat$audio1 openat$dsp openat$dsp1]
ioctl$SNDCTL_DSP_GETIPTR                    : fd_dsp [openat$adsp1 openat$audio openat$audio1 openat$dsp openat$dsp1]
ioctl$SNDCTL_DSP_GETISPACE                  : fd_dsp [openat$adsp1 openat$audio openat$audio1 openat$dsp openat$dsp1]
ioctl$SNDCTL_DSP_GETODELAY                  : fd_dsp [openat$adsp1 openat$audio openat$audio1 openat$dsp openat$dsp1]
ioctl$SNDCTL_DSP_GETOPTR                    : fd_dsp [openat$adsp1 openat$audio openat$audio1 openat$dsp openat$dsp1]
ioctl$SNDCTL_DSP_GETOSPACE                  : fd_dsp [openat$adsp1 openat$audio openat$audio1 openat$dsp openat$dsp1]
ioctl$SNDCTL_DSP_GETTRIGGER                 : fd_dsp [openat$adsp1 openat$audio openat$audio1 openat$dsp openat$dsp1]
ioctl$SNDCTL_DSP_NONBLOCK                   : fd_dsp [openat$adsp1 openat$audio openat$audio1 openat$dsp openat$dsp1]
ioctl$SNDCTL_DSP_POST                       : fd_dsp [openat$adsp1 openat$audio openat$audio1 openat$dsp openat$dsp1]
ioctl$SNDCTL_DSP_RESET                      : fd_dsp [openat$adsp1 openat$audio openat$audio1 openat$dsp openat$dsp1]
ioctl$SNDCTL_DSP_SETDUPLEX                  : fd_dsp [openat$adsp1 openat$audio openat$audio1 openat$dsp openat$dsp1]
ioctl$SNDCTL_DSP_SETFMT                     : fd_dsp [openat$adsp1 openat$audio openat$audio1 openat$dsp openat$dsp1]
ioctl$SNDCTL_DSP_SETFRAGMENT                : fd_dsp [openat$adsp1 openat$audio openat$audio1 openat$dsp openat$dsp1]
ioctl$SNDCTL_DSP_SETTRIGGER                 : fd_dsp [openat$adsp1 openat$audio openat$audio1 openat$dsp openat$dsp1]
ioctl$SNDCTL_DSP_SPEED                      : fd_dsp [openat$adsp1 openat$audio openat$audio1 openat$dsp openat$dsp1]
ioctl$SNDCTL_DSP_STEREO                     : fd_dsp [openat$adsp1 openat$audio openat$audio1 openat$dsp openat$dsp1]
ioctl$SNDCTL_DSP_SUBDIVIDE                  : fd_dsp [openat$adsp1 openat$audio openat$audio1 openat$dsp openat$dsp1]
ioctl$SNDCTL_DSP_SYNC                       : fd_dsp [openat$adsp1 openat$audio openat$audio1 openat$dsp openat$dsp1]
ioctl$SNDCTL_FM_4OP_ENABLE                  : fd_seq [openat$sequencer openat$sequencer2]
ioctl$SNDCTL_FM_LOAD_INSTR                  : fd_seq [openat$sequencer openat$sequencer2]
ioctl$SNDCTL_MIDI_INFO                      : fd_seq [openat$sequencer openat$sequencer2]
ioctl$SNDCTL_MIDI_PRETIME                   : fd_seq [openat$sequencer openat$sequencer2]
ioctl$SNDCTL_SEQ_CTRLRATE                   : fd_seq [openat$sequencer openat$sequencer2]
ioctl$SNDCTL_SEQ_GETINCOUNT                 : fd_seq [openat$sequencer openat$sequencer2]
ioctl$SNDCTL_SEQ_GETOUTCOUNT                : fd_seq [openat$sequencer openat$sequencer2]
ioctl$SNDCTL_SEQ_GETTIME                    : fd_seq [openat$sequencer openat$sequencer2]
ioctl$SNDCTL_SEQ_NRMIDIS                    : fd_seq [openat$sequencer openat$sequencer2]
ioctl$SNDCTL_SEQ_NRSYNTHS                   : fd_seq [openat$sequencer openat$sequencer2]
ioctl$SNDCTL_SEQ_OUTOFBAND                  : fd_seq [openat$sequencer openat$sequencer2]
ioctl$SNDCTL_SEQ_PANIC                      : fd_seq [openat$sequencer openat$sequencer2]
ioctl$SNDCTL_SEQ_RESET                      : fd_seq [openat$sequencer openat$sequencer2]
ioctl$SNDCTL_SEQ_RESETSAMPLES               : fd_seq [openat$sequencer openat$sequencer2]
ioctl$SNDCTL_SEQ_SYNC                       : fd_seq [openat$sequencer openat$sequencer2]
ioctl$SNDCTL_SEQ_TESTMIDI                   : fd_seq [openat$sequencer openat$sequencer2]
ioctl$SNDCTL_SEQ_THRESHOLD                  : fd_seq [openat$sequencer openat$sequencer2]
ioctl$SNDCTL_SYNTH_ID                       : fd_seq [openat$sequencer openat$sequencer2]
ioctl$SNDCTL_SYNTH_INFO                     : fd_seq [openat$sequencer openat$sequencer2]
ioctl$SNDCTL_SYNTH_MEMAVL                   : fd_seq [openat$sequencer openat$sequencer2]
ioctl$SNDCTL_TMR_CONTINUE                   : fd_seq [openat$sequencer openat$sequencer2]
ioctl$SNDCTL_TMR_METRONOME                  : fd_seq [openat$sequencer openat$sequencer2]
ioctl$SNDCTL_TMR_SELECT                     : fd_seq [openat$sequencer openat$sequencer2]
ioctl$SNDCTL_TMR_SOURCE                     : fd_seq [openat$sequencer openat$sequencer2]
ioctl$SNDCTL_TMR_START                      : fd_seq [openat$sequencer openat$sequencer2]
ioctl$SNDCTL_TMR_STOP                       : fd_seq [openat$sequencer openat$sequencer2]
ioctl$SNDCTL_TMR_TEMPO                      : fd_seq [openat$sequencer openat$sequencer2]
ioctl$SNDCTL_TMR_TIMEBASE                   : fd_seq [openat$sequencer openat$sequencer2]
ioctl$SNDRV_FIREWIRE_IOCTL_GET_INFO         : fd_snd_hw [syz_open_dev$sndhw]
ioctl$SNDRV_FIREWIRE_IOCTL_LOCK             : fd_snd_hw [syz_open_dev$sndhw]
ioctl$SNDRV_FIREWIRE_IOCTL_TASCAM_STATE     : fd_snd_hw [syz_open_dev$sndhw]
ioctl$SNDRV_FIREWIRE_IOCTL_UNLOCK           : fd_snd_hw [syz_open_dev$sndhw]
ioctl$SNDRV_HWDEP_IOCTL_DSP_LOAD            : fd_snd_hw [syz_open_dev$sndhw]
ioctl$SNDRV_HWDEP_IOCTL_DSP_STATUS          : fd_snd_hw [syz_open_dev$sndhw]
ioctl$SNDRV_HWDEP_IOCTL_INFO                : fd_snd_hw [syz_open_dev$sndhw]
ioctl$SNDRV_HWDEP_IOCTL_PVERSION            : fd_snd_hw [syz_open_dev$sndhw]
ioctl$SNDRV_RAWMIDI_IOCTL_DRAIN             : fd_midi [syz_open_dev$admmidi syz_open_dev$amidi syz_open_dev$dmmidi syz_open_dev$midi syz_open_dev$sndmidi]
ioctl$SNDRV_RAWMIDI_IOCTL_DROP              : fd_midi [syz_open_dev$admmidi syz_open_dev$amidi syz_open_dev$dmmidi syz_open_dev$midi syz_open_dev$sndmidi]
ioctl$SNDRV_RAWMIDI_IOCTL_INFO              : fd_midi [syz_open_dev$admmidi syz_open_dev$amidi syz_open_dev$dmmidi syz_open_dev$midi syz_open_dev$sndmidi]
ioctl$SNDRV_RAWMIDI_IOCTL_PARAMS            : fd_midi [syz_open_dev$admmidi syz_open_dev$amidi syz_open_dev$dmmidi syz_open_dev$midi syz_open_dev$sndmidi]
ioctl$SNDRV_RAWMIDI_IOCTL_PVERSION          : fd_midi [syz_open_dev$admmidi syz_open_dev$amidi syz_open_dev$dmmidi syz_open_dev$midi syz_open_dev$sndmidi]
ioctl$SNDRV_RAWMIDI_IOCTL_STATUS32          : fd_midi [syz_open_dev$admmidi syz_open_dev$amidi syz_open_dev$dmmidi syz_open_dev$midi syz_open_dev$sndmidi]
ioctl$SNDRV_RAWMIDI_IOCTL_STATUS64          : fd_midi [syz_open_dev$admmidi syz_open_dev$amidi syz_open_dev$dmmidi syz_open_dev$midi syz_open_dev$sndmidi]
ioctl$SNDRV_SEQ_IOCTL_CLIENT_ID             : fd_sndseq [openat$sndseq]
ioctl$SNDRV_SEQ_IOCTL_CREATE_PORT           : fd_sndseq [openat$sndseq]
ioctl$SNDRV_SEQ_IOCTL_CREATE_QUEUE          : fd_sndseq [openat$sndseq]
ioctl$SNDRV_SEQ_IOCTL_DELETE_PORT           : fd_sndseq [openat$sndseq]
ioctl$SNDRV_SEQ_IOCTL_DELETE_QUEUE          : fd_sndseq [openat$sndseq]
ioctl$SNDRV_SEQ_IOCTL_GET_CLIENT_INFO       : fd_sndseq [openat$sndseq]
ioctl$SNDRV_SEQ_IOCTL_GET_CLIENT_POOL       : fd_sndseq [openat$sndseq]
ioctl$SNDRV_SEQ_IOCTL_GET_NAMED_QUEUE       : fd_sndseq [openat$sndseq]
ioctl$SNDRV_SEQ_IOCTL_GET_PORT_INFO         : fd_sndseq [openat$sndseq]
ioctl$SNDRV_SEQ_IOCTL_GET_QUEUE_CLIENT      : fd_sndseq [openat$sndseq]
ioctl$SNDRV_SEQ_IOCTL_GET_QUEUE_INFO        : fd_sndseq [openat$sndseq]
ioctl$SNDRV_SEQ_IOCTL_GET_QUEUE_STATUS      : fd_sndseq [openat$sndseq]
ioctl$SNDRV_SEQ_IOCTL_GET_QUEUE_TEMPO       : fd_sndseq [openat$sndseq]
ioctl$SNDRV_SEQ_IOCTL_GET_QUEUE_TIMER       : fd_sndseq [openat$sndseq]
ioctl$SNDRV_SEQ_IOCTL_GET_SUBSCRIPTION      : fd_sndseq [openat$sndseq]
ioctl$SNDRV_SEQ_IOCTL_PVERSION              : fd_sndseq [openat$sndseq]
ioctl$SNDRV_SEQ_IOCTL_QUERY_NEXT_CLIENT     : fd_sndseq [openat$sndseq]
ioctl$SNDRV_SEQ_IOCTL_QUERY_NEXT_PORT       : fd_sndseq [openat$sndseq]
ioctl$SNDRV_SEQ_IOCTL_QUERY_SUBS            : fd_sndseq [openat$sndseq]
ioctl$SNDRV_SEQ_IOCTL_REMOVE_EVENTS         : fd_sndseq [openat$sndseq]
ioctl$SNDRV_SEQ_IOCTL_RUNNING_MODE          : fd_sndseq [openat$sndseq]
ioctl$SNDRV_SEQ_IOCTL_SET_CLIENT_INFO       : fd_sndseq [openat$sndseq]
ioctl$SNDRV_SEQ_IOCTL_SET_CLIENT_POOL       : fd_sndseq [openat$sndseq]
ioctl$SNDRV_SEQ_IOCTL_SET_PORT_INFO         : fd_sndseq [openat$sndseq]
ioctl$SNDRV_SEQ_IOCTL_SET_QUEUE_CLIENT      : fd_sndseq [openat$sndseq]
ioctl$SNDRV_SEQ_IOCTL_SET_QUEUE_INFO        : fd_sndseq [openat$sndseq]
ioctl$SNDRV_SEQ_IOCTL_SET_QUEUE_TEMPO       : fd_sndseq [openat$sndseq]
ioctl$SNDRV_SEQ_IOCTL_SET_QUEUE_TIMER       : fd_sndseq [openat$sndseq]
ioctl$SNDRV_SEQ_IOCTL_SUBSCRIBE_PORT        : fd_sndseq [openat$sndseq]
ioctl$SNDRV_SEQ_IOCTL_SYSTEM_INFO           : fd_sndseq [openat$sndseq]
ioctl$SNDRV_SEQ_IOCTL_UNSUBSCRIBE_PORT      : fd_sndseq [openat$sndseq]
ioctl$SOUND_MIXER_INFO                      : fd_mixer [openat$adsp1 openat$audio openat$audio1 ...]
ioctl$SOUND_MIXER_READ_CAPS                 : fd_mixer [openat$adsp1 openat$audio openat$audio1 ...]
ioctl$SOUND_MIXER_READ_DEVMASK              : fd_mixer [openat$adsp1 openat$audio openat$audio1 ...]
ioctl$SOUND_MIXER_READ_RECMASK              : fd_mixer [openat$adsp1 openat$audio openat$audio1 ...]
ioctl$SOUND_MIXER_READ_RECSRC               : fd_mixer [openat$adsp1 openat$audio openat$audio1 ...]
ioctl$SOUND_MIXER_READ_STEREODEVS           : fd_mixer [openat$adsp1 openat$audio openat$audio1 ...]
ioctl$SOUND_MIXER_READ_VOLUME               : fd_mixer [openat$adsp1 openat$audio openat$audio1 ...]
ioctl$SOUND_MIXER_WRITE_RECSRC              : fd_mixer [openat$adsp1 openat$audio openat$audio1 ...]
ioctl$SOUND_MIXER_WRITE_VOLUME              : fd_mixer [openat$adsp1 openat$audio openat$audio1 ...]
ioctl$SOUND_OLD_MIXER_INFO                  : fd_mixer [openat$adsp1 openat$audio openat$audio1 ...]
ioctl$SOUND_PCM_READ_BITS                   : fd_dsp [openat$adsp1 openat$audio openat$audio1 openat$dsp openat$dsp1]
ioctl$SOUND_PCM_READ_CHANNELS               : fd_dsp [openat$adsp1 openat$audio openat$audio1 openat$dsp openat$dsp1]
ioctl$SOUND_PCM_READ_RATE                   : fd_dsp [openat$adsp1 openat$audio openat$audio1 openat$dsp openat$dsp1]
ioctl$SW_SYNC_IOC_CREATE_FENCE              : fd_sw_sync [openat$sw_sync]
ioctl$SW_SYNC_IOC_INC                       : fd_sw_sync [openat$sw_sync]
ioctl$TE_IOCTL_CLOSE_CLIENT_SESSION         : fd_tlk [openat$tlk_device]
ioctl$TE_IOCTL_LAUNCH_OPERATION             : fd_tlk [openat$tlk_device]
ioctl$TE_IOCTL_OPEN_CLIENT_SESSION          : fd_tlk [openat$tlk_device]
ioctl$TE_IOCTL_SS_CMD                       : fd_tlk [openat$tlk_device]
ioctl$TIPC_IOC_CONNECT                      : fd_trusty [openat$trusty openat$trusty_avb openat$trusty_gatekeeper ...]
ioctl$TIPC_IOC_CONNECT_avb                  : fd_trusty_avb [openat$trusty_avb]
ioctl$TIPC_IOC_CONNECT_gatekeeper           : fd_trusty_gatekeeper [openat$trusty_gatekeeper]
ioctl$TIPC_IOC_CONNECT_hwkey                : fd_trusty_hwkey [openat$trusty_hwkey]
ioctl$TIPC_IOC_CONNECT_hwrng                : fd_trusty_hwrng [openat$trusty_hwrng]
ioctl$TIPC_IOC_CONNECT_keymaster_secure     : fd_trusty_km_secure [openat$trusty_km_secure]
ioctl$TIPC_IOC_CONNECT_km                   : fd_trusty_km [openat$trusty_km]
ioctl$TIPC_IOC_CONNECT_storage              : fd_trusty_storage [openat$trusty_storage]
ioctl$UDMABUF_CREATE                        : fd_udambuf [openat$udambuf]
ioctl$UDMABUF_CREATE_LIST                   : fd_udambuf [openat$udambuf]
ioctl$USBDEVFS_ALLOC_STREAMS                : fd_usbfs [syz_open_dev$usbfs]
ioctl$USBDEVFS_ALLOW_SUSPEND                : fd_usbfs [syz_open_dev$usbfs]
ioctl$USBDEVFS_BULK                         : fd_usbfs [syz_open_dev$usbfs]
ioctl$USBDEVFS_CLAIMINTERFACE               : fd_usbfs [syz_open_dev$usbfs]
ioctl$USBDEVFS_CLAIM_PORT                   : fd_usbfs [syz_open_dev$usbfs]
ioctl$USBDEVFS_CLEAR_HALT                   : fd_usbfs [syz_open_dev$usbfs]
ioctl$USBDEVFS_CONNECTINFO                  : fd_usbfs [syz_open_dev$usbfs]
ioctl$USBDEVFS_CONTROL                      : fd_usbfs [syz_open_dev$usbfs]
ioctl$USBDEVFS_DISCARDURB                   : fd_usbfs [syz_open_dev$usbfs]
ioctl$USBDEVFS_DISCONNECT_CLAIM             : fd_usbfs [syz_open_dev$usbfs]
ioctl$USBDEVFS_DISCSIGNAL                   : fd_usbfs [syz_open_dev$usbfs]
ioctl$USBDEVFS_DROP_PRIVILEGES              : fd_usbfs [syz_open_dev$usbfs]
ioctl$USBDEVFS_FORBID_SUSPEND               : fd_usbfs [syz_open_dev$usbfs]
ioctl$USBDEVFS_FREE_STREAMS                 : fd_usbfs [syz_open_dev$usbfs]
ioctl$USBDEVFS_GETDRIVER                    : fd_usbfs [syz_open_dev$usbfs]
ioctl$USBDEVFS_GET_CAPABILITIES             : fd_usbfs [syz_open_dev$usbfs]
ioctl$USBDEVFS_GET_SPEED                    : fd_usbfs [syz_open_dev$usbfs]
ioctl$USBDEVFS_IOCTL                        : fd_usbfs [syz_open_dev$usbfs]
ioctl$USBDEVFS_REAPURB                      : fd_usbfs [syz_open_dev$usbfs]
ioctl$USBDEVFS_REAPURBNDELAY                : fd_usbfs [syz_open_dev$usbfs]
ioctl$USBDEVFS_RELEASEINTERFACE             : fd_usbfs [syz_open_dev$usbfs]
ioctl$USBDEVFS_RELEASE_PORT                 : fd_usbfs [syz_open_dev$usbfs]
ioctl$USBDEVFS_RESET                        : fd_usbfs [syz_open_dev$usbfs]
ioctl$USBDEVFS_RESETEP                      : fd_usbfs [syz_open_dev$usbfs]
ioctl$USBDEVFS_SETCONFIGURATION             : fd_usbfs [syz_open_dev$usbfs]
ioctl$USBDEVFS_SETINTERFACE                 : fd_usbfs [syz_open_dev$usbfs]
ioctl$USBDEVFS_SUBMITURB                    : fd_usbfs [syz_open_dev$usbfs]
ioctl$USBDEVFS_WAIT_FOR_RESUME              : fd_usbfs [syz_open_dev$usbfs]
ioctl$VFIO_CHECK_EXTENSION                  : fd_vfio [openat$vfio]
ioctl$VFIO_GET_API_VERSION                  : fd_vfio [openat$vfio]
ioctl$VFIO_IOMMU_GET_INFO                   : fd_vfio [openat$vfio]
ioctl$VFIO_IOMMU_MAP_DMA                    : fd_vfio [openat$vfio]
ioctl$VFIO_IOMMU_UNMAP_DMA                  : fd_vfio [openat$vfio]
ioctl$VFIO_SET_IOMMU                        : fd_vfio [openat$vfio]
ioctl$VHOST_NET_SET_BACKEND                 : vhost_net [openat$vnet]
ioctl$VTPM_PROXY_IOC_NEW_DEV                : fd_vtpm [openat$vtpm]
ioctl$mixer_OSS_ALSAEMULVER                 : fd_mixer [openat$adsp1 openat$audio openat$audio1 ...]
ioctl$mixer_OSS_GETVERSION                  : fd_mixer [openat$adsp1 openat$audio openat$audio1 ...]
ioctl$sock_SIOCADDRT                        : nfc_dev_id [ioctl$IOCTL_GET_NCIDEV_IDX]
ioctl$sock_SIOCDELRT                        : nfc_dev_id [ioctl$IOCTL_GET_NCIDEV_IDX]
ioctl$sock_ax25_SIOCADDRT                   : sock_ax25 [accept$ax25 accept4$ax25 syz_init_net_socket$ax25]
ioctl$sock_ax25_SIOCDELRT                   : sock_ax25 [accept$ax25 accept4$ax25 syz_init_net_socket$ax25]
ioctl$sock_bt_bnep_BNEPCONNADD              : sock_bt_bnep [syz_init_net_socket$bt_bnep]
ioctl$sock_bt_bnep_BNEPCONNDEL              : sock_bt_bnep [syz_init_net_socket$bt_bnep]
ioctl$sock_bt_bnep_BNEPGETCONNINFO          : sock_bt_bnep [syz_init_net_socket$bt_bnep]
ioctl$sock_bt_bnep_BNEPGETCONNLIST          : sock_bt_bnep [syz_init_net_socket$bt_bnep]
ioctl$sock_bt_bnep_BNEPGETSUPPFEAT          : sock_bt_bnep [syz_init_net_socket$bt_bnep]
ioctl$sock_bt_cmtp_CMTPCONNADD              : sock_bt_cmtp [syz_init_net_socket$bt_cmtp]
ioctl$sock_bt_cmtp_CMTPCONNDEL              : sock_bt_cmtp [syz_init_net_socket$bt_cmtp]
ioctl$sock_bt_cmtp_CMTPGETCONNINFO          : sock_bt_cmtp [syz_init_net_socket$bt_cmtp]
ioctl$sock_bt_cmtp_CMTPGETCONNLIST          : sock_bt_cmtp [syz_init_net_socket$bt_cmtp]
ioctl$sock_bt_hidp_HIDPCONNADD              : sock_bt_hidp [syz_init_net_socket$bt_hidp]
ioctl$sock_bt_hidp_HIDPCONNDEL              : sock_bt_hidp [syz_init_net_socket$bt_hidp]
ioctl$sock_bt_hidp_HIDPGETCONNINFO          : sock_bt_hidp [syz_init_net_socket$bt_hidp]
ioctl$sock_bt_hidp_HIDPGETCONNLIST          : sock_bt_hidp [syz_init_net_socket$bt_hidp]
ioctl$sock_ifreq                            : nfc_dev_id [ioctl$IOCTL_GET_NCIDEV_IDX]
ioctl$sock_inet_sctp_SIOCINQ                : sock_sctp [socket$inet_sctp]
ioctl$sock_kcm_SIOCKCMATTACH                : sock_kcm [socket$kcm]
ioctl$sock_kcm_SIOCKCMCLONE                 : sock_kcm [socket$kcm]
ioctl$sock_kcm_SIOCKCMUNATTACH              : sock_kcm [socket$kcm]
ioctl$sock_netrom_SIOCADDRT                 : sock_netrom [accept$netrom accept4$netrom syz_init_net_socket$netrom]
ioctl$sock_netrom_SIOCDELRT                 : sock_netrom [accept$netrom accept4$netrom syz_init_net_socket$netrom]
ioctl$sock_qrtr_SIOCGIFADDR                 : sock_qrtr [socket$qrtr]
ioctl$sock_qrtr_TIOCINQ                     : sock_qrtr [socket$qrtr]
ioctl$sock_qrtr_TIOCOUTQ                    : sock_qrtr [socket$qrtr]
ioctl$sock_rose_SIOCADDRT                   : sock_rose [accept4$rose syz_init_net_socket$rose]
ioctl$sock_rose_SIOCDELRT                   : sock_rose [accept4$rose syz_init_net_socket$rose]
ioctl$sock_rose_SIOCRSCLRRT                 : sock_rose [accept4$rose syz_init_net_socket$rose]
ioctl$sock_x25_SIOCADDRT                    : sock_x25 [accept4$x25 syz_init_net_socket$x25]
ioctl$sock_x25_SIOCDELRT                    : sock_x25 [accept4$x25 syz_init_net_socket$x25]
mmap$DRM_I915                               : fd_i915 [openat$i915]
mmap$DRM_MSM                                : fd_msm [openat$msm]
mmap$KVM_VCPU                               : vcpu_mmap_size [ioctl$KVM_GET_VCPU_MMAP_SIZE]
mmap$dsp                                    : fd_dsp [openat$adsp1 openat$audio openat$audio1 openat$dsp openat$dsp1]
mmap$fb                                     : fd_fb [openat$fb0 openat$fb1]
mmap$qrtrtun                                : fd_qrtr_tun [openat$qrtrtun]
mmap$usbfs                                  : fd_usbfs [syz_open_dev$usbfs]
mmap$usbmon                                 : fd_usbmon [syz_open_dev$usbmon]
read$alg                                    : sock_algconn [accept$alg accept4$alg]
read$dsp                                    : fd_dsp [openat$adsp1 openat$audio openat$audio1 openat$dsp openat$dsp1]
read$fb                                     : fd_fb [openat$fb0 openat$fb1]
read$hiddev                                 : fd_hiddev [syz_open_dev$hiddev]
read$hidraw                                 : fd_hidraw [syz_open_dev$hidraw]
read$midi                                   : fd_midi [syz_open_dev$admmidi syz_open_dev$amidi syz_open_dev$dmmidi syz_open_dev$midi syz_open_dev$sndmidi]
read$msr                                    : fd_msr [syz_open_dev$MSR]
read$nci                                    : fd_nci [openat$nci]
read$proc_mixer                             : fd_proc_mixer [openat$proc_mixer]
read$ptp                                    : fd_ptp [openat$ptp0 openat$ptp1]
read$qrtrtun                                : fd_qrtr_tun [openat$qrtrtun]
read$sequencer                              : fd_seq [openat$sequencer openat$sequencer2]
read$smackfs_access                         : fd_smackfs_access [openat$smackfs_access]
read$smackfs_cipsonum                       : fd_smackfs_cipsonum [openat$smackfs_cipsonum]
read$smackfs_logging                        : fd_smackfs_logging [openat$smackfs_logging]
read$smackfs_ptrace                         : fd_smackfs_ptrace [openat$smackfs_ptrace]
read$snapshot                               : fd_snapshot [openat$snapshot]
read$sndhw                                  : fd_snd_hw [syz_open_dev$sndhw]
read$trusty                                 : fd_trusty [openat$trusty openat$trusty_avb openat$trusty_gatekeeper ...]
read$usbfs                                  : fd_usbfs [syz_open_dev$usbfs]
read$usbmon                                 : fd_usbmon [syz_open_dev$usbmon]
recvfrom$ax25                               : sock_ax25 [accept$ax25 accept4$ax25 syz_init_net_socket$ax25]
recvfrom$l2tp                               : sock_l2tp [socket$l2tp]
recvfrom$l2tp6                              : sock_l2tp6 [socket$l2tp6]
recvfrom$llc                                : sock_llc [accept4$llc syz_init_net_socket$llc]
recvfrom$netrom                             : sock_netrom [accept$netrom accept4$netrom syz_init_net_socket$netrom]
recvfrom$phonet                             : sock_phonet [accept$phonet_pipe accept4$phonet_pipe socket$phonet socket$phonet_pipe]
recvfrom$rose                               : sock_rose [accept4$rose syz_init_net_socket$rose]
recvfrom$rxrpc                              : sock_rxrpc [socket$rxrpc]
recvfrom$x25                                : sock_x25 [accept4$x25 syz_init_net_socket$x25]
recvmsg$can_bcm                             : sock_can_bcm [socket$can_bcm]
recvmsg$can_j1939                           : sock_can_j1939 [socket$can_j1939]
recvmsg$can_raw                             : sock_can_raw [socket$can_raw]
recvmsg$hf                                  : sock_hf [socket$hf]
recvmsg$kcm                                 : sock_kcm [socket$kcm]
recvmsg$qrtr                                : sock_qrtr [socket$qrtr]
semctl$GETALL                               : ipc_sem [semget semget$private]
semctl$GETNCNT                              : ipc_sem [semget semget$private]
semctl$GETPID                               : ipc_sem [semget semget$private]
semctl$GETVAL                               : ipc_sem [semget semget$private]
semctl$GETZCNT                              : ipc_sem [semget semget$private]
semctl$IPC_INFO                             : ipc_sem [semget semget$private]
semctl$IPC_RMID                             : ipc_sem [semget semget$private]
semctl$IPC_SET                              : ipc_sem [semget semget$private]
semctl$IPC_STAT                             : ipc_sem [semget semget$private]
semctl$SEM_INFO                             : ipc_sem [semget semget$private]
semctl$SEM_STAT                             : ipc_sem [semget semget$private]
semctl$SEM_STAT_ANY                         : ipc_sem [semget semget$private]
semctl$SETALL                               : ipc_sem [semget semget$private]
semctl$SETVAL                               : ipc_sem [semget semget$private]
sendmmsg$alg                                : sock_algconn [accept$alg accept4$alg]
sendmmsg$inet_sctp                          : sock_sctp [socket$inet_sctp]
sendmmsg$nfc_llcp                           : sock_nfc_llcp [accept$nfc_llcp accept4$nfc_llcp syz_init_net_socket$nfc_llcp]
sendmsg$802154_dgram                        : sock_802154_dgram [syz_init_net_socket$802154_dgram]
sendmsg$802154_raw                          : sock_802154_raw [syz_init_net_socket$802154_raw]
sendmsg$L2TP_CMD_NOOP                       : sock_l2tp [socket$l2tp]
sendmsg$L2TP_CMD_SESSION_CREATE             : sock_l2tp [socket$l2tp]
sendmsg$L2TP_CMD_SESSION_DELETE             : sock_l2tp [socket$l2tp]
sendmsg$L2TP_CMD_SESSION_GET                : sock_l2tp [socket$l2tp]
sendmsg$L2TP_CMD_SESSION_MODIFY             : sock_l2tp [socket$l2tp]
sendmsg$L2TP_CMD_TUNNEL_CREATE              : sock_l2tp [socket$l2tp]
sendmsg$L2TP_CMD_TUNNEL_DELETE              : sock_l2tp [socket$l2tp]
sendmsg$L2TP_CMD_TUNNEL_GET                 : sock_l2tp [socket$l2tp]
sendmsg$L2TP_CMD_TUNNEL_MODIFY              : sock_l2tp [socket$l2tp]
sendmsg$NFC_CMD_ACTIVATE_TARGET             : nfc_dev_id [ioctl$IOCTL_GET_NCIDEV_IDX]
sendmsg$NFC_CMD_DEACTIVATE_TARGET           : nfc_dev_id [ioctl$IOCTL_GET_NCIDEV_IDX]
sendmsg$NFC_CMD_DEP_LINK_DOWN               : nfc_dev_id [ioctl$IOCTL_GET_NCIDEV_IDX]
sendmsg$NFC_CMD_DEP_LINK_UP                 : nfc_dev_id [ioctl$IOCTL_GET_NCIDEV_IDX]
sendmsg$NFC_CMD_DEV_DOWN                    : nfc_dev_id [ioctl$IOCTL_GET_NCIDEV_IDX]
sendmsg$NFC_CMD_DEV_UP                      : nfc_dev_id [ioctl$IOCTL_GET_NCIDEV_IDX]
sendmsg$NFC_CMD_DISABLE_SE                  : nfc_dev_id [ioctl$IOCTL_GET_NCIDEV_IDX]
sendmsg$NFC_CMD_ENABLE_SE                   : nfc_dev_id [ioctl$IOCTL_GET_NCIDEV_IDX]
sendmsg$NFC_CMD_FW_DOWNLOAD                 : nfc_dev_id [ioctl$IOCTL_GET_NCIDEV_IDX]
sendmsg$NFC_CMD_GET_DEVICE                  : nfc_dev_id [ioctl$IOCTL_GET_NCIDEV_IDX]
sendmsg$NFC_CMD_LLC_GET_PARAMS              : nfc_dev_id [ioctl$IOCTL_GET_NCIDEV_IDX]
sendmsg$NFC_CMD_LLC_SDREQ                   : nfc_dev_id [ioctl$IOCTL_GET_NCIDEV_IDX]
sendmsg$NFC_CMD_LLC_SET_PARAMS              : nfc_dev_id [ioctl$IOCTL_GET_NCIDEV_IDX]
sendmsg$NFC_CMD_SE_IO                       : nfc_dev_id [ioctl$IOCTL_GET_NCIDEV_IDX]
sendmsg$NFC_CMD_START_POLL                  : nfc_dev_id [ioctl$IOCTL_GET_NCIDEV_IDX]
sendmsg$NFC_CMD_VENDOR                      : nfc_dev_id [ioctl$IOCTL_GET_NCIDEV_IDX]
sendmsg$RDMA_NLDEV_CMD_DELLINK              : sock_nl_rdma [socket$nl_rdma syz_init_net_socket$nl_rdma]
sendmsg$RDMA_NLDEV_CMD_GET                  : sock_nl_rdma [socket$nl_rdma syz_init_net_socket$nl_rdma]
sendmsg$RDMA_NLDEV_CMD_GET_CHARDEV          : sock_nl_rdma [socket$nl_rdma syz_init_net_socket$nl_rdma]
sendmsg$RDMA_NLDEV_CMD_NEWLINK              : sock_nl_rdma [socket$nl_rdma syz_init_net_socket$nl_rdma]
sendmsg$RDMA_NLDEV_CMD_PORT_GET             : sock_nl_rdma [socket$nl_rdma syz_init_net_socket$nl_rdma]
sendmsg$RDMA_NLDEV_CMD_RES_CM_ID_GET        : sock_nl_rdma [socket$nl_rdma syz_init_net_socket$nl_rdma]
sendmsg$RDMA_NLDEV_CMD_RES_CQ_GET           : sock_nl_rdma [socket$nl_rdma syz_init_net_socket$nl_rdma]
sendmsg$RDMA_NLDEV_CMD_RES_GET              : sock_nl_rdma [socket$nl_rdma syz_init_net_socket$nl_rdma]
sendmsg$RDMA_NLDEV_CMD_RES_MR_GET           : sock_nl_rdma [socket$nl_rdma syz_init_net_socket$nl_rdma]
sendmsg$RDMA_NLDEV_CMD_RES_PD_GET           : sock_nl_rdma [socket$nl_rdma syz_init_net_socket$nl_rdma]
sendmsg$RDMA_NLDEV_CMD_RES_QP_GET           : sock_nl_rdma [socket$nl_rdma syz_init_net_socket$nl_rdma]
sendmsg$RDMA_NLDEV_CMD_SET                  : sock_nl_rdma [socket$nl_rdma syz_init_net_socket$nl_rdma]
sendmsg$RDMA_NLDEV_CMD_STAT_DEL             : sock_nl_rdma [socket$nl_rdma syz_init_net_socket$nl_rdma]
sendmsg$RDMA_NLDEV_CMD_STAT_GET             : sock_nl_rdma [socket$nl_rdma syz_init_net_socket$nl_rdma]
sendmsg$RDMA_NLDEV_CMD_STAT_SET             : sock_nl_rdma [socket$nl_rdma syz_init_net_socket$nl_rdma]
sendmsg$RDMA_NLDEV_CMD_SYS_GET              : sock_nl_rdma [socket$nl_rdma syz_init_net_socket$nl_rdma]
sendmsg$RDMA_NLDEV_CMD_SYS_SET              : sock_nl_rdma [socket$nl_rdma syz_init_net_socket$nl_rdma]
sendmsg$alg                                 : sock_algconn [accept$alg accept4$alg]
sendmsg$can_bcm                             : sock_can_bcm [socket$can_bcm]
sendmsg$can_j1939                           : sock_can_j1939 [socket$can_j1939]
sendmsg$can_raw                             : sock_can_raw [socket$can_raw]
sendmsg$hf                                  : sock_hf [socket$hf]
sendmsg$inet_sctp                           : sock_sctp [socket$inet_sctp]
sendmsg$kcm                                 : sock_kcm [socket$kcm]
sendmsg$nfc_llcp                            : sock_nfc_llcp [accept$nfc_llcp accept4$nfc_llcp syz_init_net_socket$nfc_llcp]
sendmsg$nl_crypto                           : sock_nl_crypto [socket$nl_crypto]
sendmsg$qrtr                                : sock_qrtr [socket$qrtr]
sendmsg$rds                                 : sock_rds [socket$rds]
sendmsg$tipc                                : sock_tipc [accept4$tipc socket$tipc socketpair$tipc]
sendto$ax25                                 : sock_ax25 [accept$ax25 accept4$ax25 syz_init_net_socket$ax25]
sendto$isdn                                 : sock_isdn [socket$isdn]
sendto$l2tp                                 : sock_l2tp [socket$l2tp]
sendto$l2tp6                                : sock_l2tp6 [socket$l2tp6]
sendto$llc                                  : sock_llc [accept4$llc syz_init_net_socket$llc]
sendto$netrom                               : sock_netrom [accept$netrom accept4$netrom syz_init_net_socket$netrom]
sendto$phonet                               : sock_phonet [accept$phonet_pipe accept4$phonet_pipe socket$phonet socket$phonet_pipe]
sendto$rose                                 : sock_rose [accept4$rose syz_init_net_socket$rose]
sendto$rxrpc                                : sock_rxrpc [socket$rxrpc]
sendto$x25                                  : sock_x25 [accept4$x25 syz_init_net_socket$x25]
setsockopt$ALG_SET_AEAD_AUTHSIZE            : sock_alg [socket$alg]
setsockopt$ALG_SET_KEY                      : sock_alg [socket$alg]
setsockopt$CAIFSO_LINK_SELECT               : sock_caif [socket$caif_seqpacket socket$caif_stream]
setsockopt$CAIFSO_REQ_PARAM                 : sock_caif [socket$caif_seqpacket socket$caif_stream]
setsockopt$CAN_RAW_ERR_FILTER               : sock_can_raw [socket$can_raw]
setsockopt$CAN_RAW_FD_FRAMES                : sock_can_raw [socket$can_raw]
setsockopt$CAN_RAW_FILTER                   : sock_can_raw [socket$can_raw]
setsockopt$CAN_RAW_JOIN_FILTERS             : sock_can_raw [socket$can_raw]
setsockopt$CAN_RAW_LOOPBACK                 : sock_can_raw [socket$can_raw]
setsockopt$CAN_RAW_RECV_OWN_MSGS            : sock_can_raw [socket$can_raw]
setsockopt$MISDN_TIME_STAMP                 : sock_isdn [socket$isdn]
setsockopt$PNPIPE_ENCAP                     : sock_phonet_pipe [accept$phonet_pipe accept4$phonet_pipe socket$phonet_pipe]
setsockopt$PNPIPE_HANDLE                    : sock_phonet_pipe [accept$phonet_pipe accept4$phonet_pipe socket$phonet_pipe]
setsockopt$PNPIPE_INITSTATE                 : sock_phonet_pipe [accept$phonet_pipe accept4$phonet_pipe socket$phonet_pipe]
setsockopt$RDS_CANCEL_SENT_TO               : sock_rds [socket$rds]
setsockopt$RDS_CONG_MONITOR                 : sock_rds [socket$rds]
setsockopt$RDS_FREE_MR                      : sock_rds [socket$rds]
setsockopt$RDS_GET_MR                       : sock_rds [socket$rds]
setsockopt$RDS_GET_MR_FOR_DEST              : sock_rds [socket$rds]
setsockopt$RDS_RECVERR                      : sock_rds [socket$rds]
setsockopt$RXRPC_EXCLUSIVE_CONNECTION       : sock_rxrpc [socket$rxrpc]
setsockopt$RXRPC_MIN_SECURITY_LEVEL         : sock_rxrpc [socket$rxrpc]
setsockopt$RXRPC_SECURITY_KEY               : sock_rxrpc [socket$rxrpc]
setsockopt$RXRPC_SECURITY_KEYRING           : sock_rxrpc [socket$rxrpc]
setsockopt$RXRPC_UPGRADEABLE_SERVICE        : sock_rxrpc [socket$rxrpc]
setsockopt$SO_J1939_ERRQUEUE                : sock_can_j1939 [socket$can_j1939]
setsockopt$SO_J1939_FILTER                  : sock_can_j1939 [socket$can_j1939]
setsockopt$SO_J1939_PROMISC                 : sock_can_j1939 [socket$can_j1939]
setsockopt$SO_J1939_SEND_PRIO               : sock_can_j1939 [socket$can_j1939]
setsockopt$SO_RDS_MSG_RXPATH_LATENCY        : sock_rds [socket$rds]
setsockopt$SO_RDS_TRANSPORT                 : sock_rds [socket$rds]
setsockopt$TIPC_CONN_TIMEOUT                : sock_tipc [accept4$tipc socket$tipc socketpair$tipc]
setsockopt$TIPC_DEST_DROPPABLE              : sock_tipc [accept4$tipc socket$tipc socketpair$tipc]
setsockopt$TIPC_GROUP_JOIN                  : sock_tipc [accept4$tipc socket$tipc socketpair$tipc]
setsockopt$TIPC_GROUP_LEAVE                 : sock_tipc [accept4$tipc socket$tipc socketpair$tipc]
setsockopt$TIPC_IMPORTANCE                  : sock_tipc [accept4$tipc socket$tipc socketpair$tipc]
setsockopt$TIPC_MCAST_BROADCAST             : sock_tipc [accept4$tipc socket$tipc socketpair$tipc]
setsockopt$TIPC_MCAST_REPLICAST             : sock_tipc [accept4$tipc socket$tipc socketpair$tipc]
setsockopt$TIPC_SRC_DROPPABLE               : sock_tipc [accept4$tipc socket$tipc socketpair$tipc]
setsockopt$WPAN_SECURITY                    : sock_802154_dgram [syz_init_net_socket$802154_dgram]
setsockopt$WPAN_SECURITY_LEVEL              : sock_802154_dgram [syz_init_net_socket$802154_dgram]
setsockopt$WPAN_WANTACK                     : sock_802154_dgram [syz_init_net_socket$802154_dgram]
setsockopt$WPAN_WANTLQI                     : sock_802154_dgram [syz_init_net_socket$802154_dgram]
setsockopt$X25_QBITINCL                     : sock_x25 [accept4$x25 syz_init_net_socket$x25]
setsockopt$ax25_SO_BINDTODEVICE             : sock_ax25 [accept$ax25 accept4$ax25 syz_init_net_socket$ax25]
setsockopt$ax25_int                         : sock_ax25 [accept$ax25 accept4$ax25 syz_init_net_socket$ax25]
setsockopt$bt_rfcomm_RFCOMM_LM              : sock_bt_rfcomm [syz_init_net_socket$bt_rfcomm]
setsockopt$inet6_dccp_buf                   : sock_dccp6 [socket$inet6_dccp]
setsockopt$inet6_dccp_int                   : sock_dccp6 [socket$inet6_dccp]
setsockopt$inet_dccp_buf                    : sock_dccp [socket$inet_dccp]
setsockopt$inet_dccp_int                    : sock_dccp [socket$inet_dccp]
setsockopt$inet_sctp6_SCTP_ADAPTATION_LAYER : sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_ADD_STREAMS      : sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_ASSOCINFO        : sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_AUTH_ACTIVE_KEY  : sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_AUTH_CHUNK       : sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_AUTH_DEACTIVATE_KEY: sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_AUTH_DELETE_KEY  : sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_AUTH_KEY         : sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_AUTOCLOSE        : sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_AUTO_ASCONF      : sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_CONTEXT          : sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_DEFAULT_PRINFO   : sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_DEFAULT_SEND_PARAM: sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_DEFAULT_SNDINFO  : sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_DELAYED_SACK     : sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_DISABLE_FRAGMENTS: sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_ENABLE_STREAM_RESET: sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_EVENTS           : sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_FRAGMENT_INTERLEAVE: sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_HMAC_IDENT       : sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_INITMSG          : sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_I_WANT_MAPPED_V4_ADDR: sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_MAXSEG           : sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_MAX_BURST        : sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_NODELAY          : sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_PARTIAL_DELIVERY_POINT: sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_PEER_ADDR_PARAMS : sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_PEER_ADDR_THLDS  : sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_PRIMARY_ADDR     : sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_PR_SUPPORTED     : sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_RECONFIG_SUPPORTED: sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_RECVNXTINFO      : sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_RECVRCVINFO      : sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_RESET_ASSOC      : sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_RESET_STREAMS    : sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_RTOINFO          : sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_SET_PEER_PRIMARY_ADDR: sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_SOCKOPT_BINDX_ADD: sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_SOCKOPT_BINDX_REM: sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_SOCKOPT_CONNECTX : sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_SOCKOPT_CONNECTX_OLD: sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_STREAM_SCHEDULER : sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_STREAM_SCHEDULER_VALUE: sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp_SCTP_ADAPTATION_LAYER  : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_ADD_STREAMS       : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_ASSOCINFO         : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_AUTH_ACTIVE_KEY   : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_AUTH_CHUNK        : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_AUTH_DEACTIVATE_KEY: sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_AUTH_DELETE_KEY   : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_AUTH_KEY          : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_AUTOCLOSE         : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_AUTO_ASCONF       : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_CONTEXT           : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_DEFAULT_PRINFO    : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_DEFAULT_SEND_PARAM: sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_DEFAULT_SNDINFO   : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_DELAYED_SACK      : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_DISABLE_FRAGMENTS : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_ENABLE_STREAM_RESET: sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_EVENTS            : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_FRAGMENT_INTERLEAVE: sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_HMAC_IDENT        : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_INITMSG           : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_I_WANT_MAPPED_V4_ADDR: sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_MAXSEG            : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_MAX_BURST         : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_NODELAY           : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_PARTIAL_DELIVERY_POINT: sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_PEER_ADDR_PARAMS  : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_PEER_ADDR_THLDS   : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_PRIMARY_ADDR      : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_PR_SUPPORTED      : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_RECONFIG_SUPPORTED: sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_RECVNXTINFO       : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_RECVRCVINFO       : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_RESET_ASSOC       : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_RESET_STREAMS     : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_RTOINFO           : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_SET_PEER_PRIMARY_ADDR: sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_SOCKOPT_BINDX_ADD : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_SOCKOPT_BINDX_REM : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_SOCKOPT_CONNECTX  : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_SOCKOPT_CONNECTX_OLD: sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_STREAM_SCHEDULER  : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_STREAM_SCHEDULER_VALUE: sock_sctp [socket$inet_sctp]
setsockopt$kcm_KCM_RECV_DISABLE             : sock_kcm [socket$kcm]
setsockopt$llc_int                          : sock_llc [accept4$llc syz_init_net_socket$llc]
setsockopt$netrom_NETROM_IDLE               : sock_netrom [accept$netrom accept4$netrom syz_init_net_socket$netrom]
setsockopt$netrom_NETROM_N2                 : sock_netrom [accept$netrom accept4$netrom syz_init_net_socket$netrom]
setsockopt$netrom_NETROM_T1                 : sock_netrom [accept$netrom accept4$netrom syz_init_net_socket$netrom]
setsockopt$netrom_NETROM_T2                 : sock_netrom [accept$netrom accept4$netrom syz_init_net_socket$netrom]
setsockopt$netrom_NETROM_T4                 : sock_netrom [accept$netrom accept4$netrom syz_init_net_socket$netrom]
setsockopt$nfc_llcp_NFC_LLCP_MIUX           : sock_nfc_llcp [accept$nfc_llcp accept4$nfc_llcp syz_init_net_socket$nfc_llcp]
setsockopt$nfc_llcp_NFC_LLCP_RW             : sock_nfc_llcp [accept$nfc_llcp accept4$nfc_llcp syz_init_net_socket$nfc_llcp]
setsockopt$pppl2tp_PPPOL2TP_SO_DEBUG        : sock_pppl2tp [socket$pppl2tp]
setsockopt$pppl2tp_PPPOL2TP_SO_LNSMODE      : sock_pppl2tp [socket$pppl2tp]
setsockopt$pppl2tp_PPPOL2TP_SO_RECVSEQ      : sock_pppl2tp [socket$pppl2tp]
setsockopt$pppl2tp_PPPOL2TP_SO_REORDERTO    : sock_pppl2tp [socket$pppl2tp]
setsockopt$pppl2tp_PPPOL2TP_SO_SENDSEQ      : sock_pppl2tp [socket$pppl2tp]
setsockopt$rose                             : sock_rose [accept4$rose syz_init_net_socket$rose]
syz_io_uring_submit                         : nfc_dev_id [ioctl$IOCTL_GET_NCIDEV_IDX]
syz_kvm_add_vcpu                            : kvm_syz_vm [syz_kvm_setup_syzos_vm]
syz_kvm_setup_cpu$arm64                     : fd_kvmvm [ioctl$KVM_CREATE_VM]
syz_kvm_setup_syzos_vm                      : fd_kvmvm [ioctl$KVM_CREATE_VM]
syz_kvm_vgic_v3_setup                       : fd_kvmvm [ioctl$KVM_CREATE_VM]
syz_memcpy_off$KVM_EXIT_HYPERCALL           : kvm_run_ptr [mmap$KVM_VCPU]
syz_memcpy_off$KVM_EXIT_MMIO                : kvm_run_ptr [mmap$KVM_VCPU]
write$6lowpan_control                       : fd_6lowpan_control [openat$6lowpan_control]
write$6lowpan_enable                        : fd_6lowpan_enable [openat$6lowpan_enable]
write$ALLOC_MW                              : fd_rdma [openat$uverbs0]
write$ALLOC_PD                              : fd_rdma [openat$uverbs0]
write$ATTACH_MCAST                          : fd_rdma [openat$uverbs0]
write$CLOSE_XRCD                            : fd_rdma [openat$uverbs0]
write$CREATE_AH                             : fd_rdma [openat$uverbs0]
write$CREATE_COMP_CHANNEL                   : fd_rdma [openat$uverbs0]
write$CREATE_CQ                             : fd_rdma [openat$uverbs0]
write$CREATE_CQ_EX                          : fd_rdma [openat$uverbs0]
write$CREATE_FLOW                           : fd_rdma [openat$uverbs0]
write$CREATE_QP                             : fd_rdma [openat$uverbs0]
write$CREATE_RWQ_IND_TBL                    : fd_rdma [openat$uverbs0]
write$CREATE_SRQ                            : fd_rdma [openat$uverbs0]
write$CREATE_WQ                             : fd_rdma [openat$uverbs0]
write$DEALLOC_MW                            : fd_rdma [openat$uverbs0]
write$DEALLOC_PD                            : fd_rdma [openat$uverbs0]
write$DEREG_MR                              : fd_rdma [openat$uverbs0]
write$DESTROY_AH                            : fd_rdma [openat$uverbs0]
write$DESTROY_CQ                            : fd_rdma [openat$uverbs0]
write$DESTROY_FLOW                          : fd_rdma [openat$uverbs0]
write$DESTROY_QP                            : fd_rdma [openat$uverbs0]
write$DESTROY_RWQ_IND_TBL                   : fd_rdma [openat$uverbs0]
write$DESTROY_SRQ                           : fd_rdma [openat$uverbs0]
write$DESTROY_WQ                            : fd_rdma [openat$uverbs0]
write$DETACH_MCAST                          : fd_rdma [openat$uverbs0]
write$MLX5_ALLOC_PD                         : fd_rdma [openat$uverbs0]
write$MLX5_CREATE_CQ                        : fd_rdma [openat$uverbs0]
write$MLX5_CREATE_DV_QP                     : fd_rdma [openat$uverbs0]
write$MLX5_CREATE_QP                        : fd_rdma [openat$uverbs0]
write$MLX5_CREATE_SRQ                       : fd_rdma [openat$uverbs0]
write$MLX5_CREATE_WQ                        : fd_rdma [openat$uverbs0]
write$MLX5_GET_CONTEXT                      : fd_rdma [openat$uverbs0]
write$MLX5_MODIFY_WQ                        : fd_rdma [openat$uverbs0]
write$MODIFY_QP                             : fd_rdma [openat$uverbs0]
write$MODIFY_SRQ                            : fd_rdma [openat$uverbs0]
write$OPEN_XRCD                             : fd_rdma [openat$uverbs0]
write$POLL_CQ                               : fd_rdma [openat$uverbs0]
write$POST_RECV                             : fd_rdma [openat$uverbs0]
write$POST_SEND                             : fd_rdma [openat$uverbs0]
write$POST_SRQ_RECV                         : fd_rdma [openat$uverbs0]
write$QUERY_DEVICE_EX                       : fd_rdma [openat$uverbs0]
write$QUERY_PORT                            : fd_rdma [openat$uverbs0]
write$QUERY_QP                              : fd_rdma [openat$uverbs0]
write$QUERY_SRQ                             : fd_rdma [openat$uverbs0]
write$RDMA_USER_CM_CMD_ACCEPT               : fd_rdma_cm [openat$rdma_cm]
write$RDMA_USER_CM_CMD_BIND                 : fd_rdma_cm [openat$rdma_cm]
write$RDMA_USER_CM_CMD_BIND_IP              : fd_rdma_cm [openat$rdma_cm]
write$RDMA_USER_CM_CMD_CONNECT              : fd_rdma_cm [openat$rdma_cm]
write$RDMA_USER_CM_CMD_CREATE_ID            : fd_rdma_cm [openat$rdma_cm]
write$RDMA_USER_CM_CMD_DESTROY_ID           : fd_rdma_cm [openat$rdma_cm]
write$RDMA_USER_CM_CMD_DISCONNECT           : fd_rdma_cm [openat$rdma_cm]
write$RDMA_USER_CM_CMD_GET_EVENT            : fd_rdma_cm [openat$rdma_cm]
write$RDMA_USER_CM_CMD_INIT_QP_ATTR         : fd_rdma_cm [openat$rdma_cm]
write$RDMA_USER_CM_CMD_JOIN_IP_MCAST        : fd_rdma_cm [openat$rdma_cm]
write$RDMA_USER_CM_CMD_JOIN_MCAST           : fd_rdma_cm [openat$rdma_cm]
write$RDMA_USER_CM_CMD_LEAVE_MCAST          : fd_rdma_cm [openat$rdma_cm]
write$RDMA_USER_CM_CMD_LISTEN               : fd_rdma_cm [openat$rdma_cm]
write$RDMA_USER_CM_CMD_MIGRATE_ID           : fd_rdma_cm [openat$rdma_cm]
write$RDMA_USER_CM_CMD_NOTIFY               : fd_rdma_cm [openat$rdma_cm]
write$RDMA_USER_CM_CMD_QUERY                : fd_rdma_cm [openat$rdma_cm]
write$RDMA_USER_CM_CMD_QUERY_ROUTE          : fd_rdma_cm [openat$rdma_cm]
write$RDMA_USER_CM_CMD_REJECT               : fd_rdma_cm [openat$rdma_cm]
write$RDMA_USER_CM_CMD_RESOLVE_ADDR         : fd_rdma_cm [openat$rdma_cm]
write$RDMA_USER_CM_CMD_RESOLVE_IP           : fd_rdma_cm [openat$rdma_cm]
write$RDMA_USER_CM_CMD_RESOLVE_ROUTE        : fd_rdma_cm [openat$rdma_cm]
write$RDMA_USER_CM_CMD_SET_OPTION           : fd_rdma_cm [openat$rdma_cm]
write$REG_MR                                : fd_rdma [openat$uverbs0]
write$REQ_NOTIFY_CQ                         : fd_rdma [openat$uverbs0]
write$REREG_MR                              : fd_rdma [openat$uverbs0]
write$RESIZE_CQ                             : fd_rdma [openat$uverbs0]
write$USERIO_CMD_REGISTER                   : fd_userio [openat$userio]
write$USERIO_CMD_SEND_INTERRUPT             : fd_userio [openat$userio]
write$USERIO_CMD_SET_PORT_TYPE              : fd_userio [openat$userio]
write$capi20                                : fd_capi20 [openat$capi20]
write$capi20_data                           : fd_capi20 [openat$capi20]
write$damon_attrs                           : fd_damon_attrs [openat$damon_attrs]
write$damon_contexts                        : fd_damon_contexts [openat$damon_mk_contexts openat$damon_rm_contexts]
write$damon_init_regions                    : fd_damon_init_regions [openat$damon_init_regions]
write$damon_monitor_on                      : fd_damon_monitor_on [openat$damon_monitor_on]
write$damon_schemes                         : fd_damon_schemes [openat$damon_schemes]
write$damon_target_ids                      : fd_damon_target_ids [openat$damon_target_ids]
write$dsp                                   : fd_dsp [openat$adsp1 openat$audio openat$audio1 openat$dsp openat$dsp1]
write$fb                                    : fd_fb [openat$fb0 openat$fb1]
write$hidraw                                : fd_hidraw [syz_open_dev$hidraw]
write$midi                                  : fd_midi [syz_open_dev$admmidi syz_open_dev$amidi syz_open_dev$dmmidi syz_open_dev$midi syz_open_dev$sndmidi]
write$nci                                   : fd_nci [openat$nci]
write$ppp                                   : fd_ppp [openat$ppp]
write$proc_mixer                            : fd_proc_mixer [openat$proc_mixer]
write$proc_reclaim                          : fd_proc_reclaim [openat$proc_reclaim]
write$qrtrtun                               : fd_qrtr_tun [openat$qrtrtun]
write$sequencer                             : fd_seq [openat$sequencer openat$sequencer2]
write$smackfs_access                        : fd_smackfs_access [openat$smackfs_access]
write$smackfs_change_rule                   : fd_smackfs_change_rule [openat$smackfs_change_rule]
write$smackfs_cipso                         : fd_smackfs_cipso [openat$smackfs_cipso]
write$smackfs_cipsonum                      : fd_smackfs_cipsonum [openat$smackfs_cipsonum]
write$smackfs_ipv6host                      : fd_smackfs_ipv6host [openat$smackfs_ipv6host]
write$smackfs_label                         : fd_smackfs_label [openat$smackfs_ambient openat$smackfs_revoke_subject openat$smackfs_syslog openat$smackfs_unconfined]
write$smackfs_labels_list                   : fd_smackfs_labels_list [openat$smackfs_onlycap openat$smackfs_relabel_self]
write$smackfs_load                          : fd_smackfs_load [openat$smackfs_load]
write$smackfs_logging                       : fd_smackfs_logging [openat$smackfs_logging]
write$smackfs_netlabel                      : fd_smackfs_netlabel [openat$smackfs_netlabel]
write$smackfs_ptrace                        : fd_smackfs_ptrace [openat$smackfs_ptrace]
write$snapshot                              : fd_snapshot [openat$snapshot]
write$sndhw                                 : fd_snd_hw [syz_open_dev$sndhw]
write$sndhw_fireworks                       : fd_snd_hw [syz_open_dev$sndhw]
write$sndseq                                : fd_sndseq [openat$sndseq]
write$sysctl                                : fd_sysctl [openat$sysctl]
write$trusty                                : fd_trusty [openat$trusty openat$trusty_avb openat$trusty_gatekeeper ...]
write$trusty_avb                            : fd_trusty_avb [openat$trusty_avb]
write$trusty_gatekeeper                     : fd_trusty_gatekeeper [openat$trusty_gatekeeper]
write$trusty_hwkey                          : fd_trusty_hwkey [openat$trusty_hwkey]
write$trusty_hwrng                          : fd_trusty_hwrng [openat$trusty_hwrng]
write$trusty_km                             : fd_trusty_km [openat$trusty_km]
write$trusty_km_secure                      : fd_trusty_km_secure [openat$trusty_km_secure]
write$trusty_storage                        : fd_trusty_storage [openat$trusty_storage]
write$usbip_server                          : fd_usbip_server [syz_usbip_server_init]
write$vga_arbiter                           : fd_vga_arbiter [openat$vga_arbiter]
write$vhost_msg                             : vhost_net [openat$vnet]
write$vhost_msg_v2                          : vhost_net [openat$vnet]
write$yama_ptrace_scope                     : fd_yama_ptrace_scope [openat$yama_ptrace_scope]

failed to read the following files in the VM:
/sys/kernel/security/lsm                    : No such file or directory

BinFmtMisc              : enabled
Comparisons             : enabled
Coverage                : enabled
DelayKcovMmap           : enabled
DevlinkPCI              : PCI device 0000:00:10.0 is not available
ExtraCoverage           : enabled
Fault                   : CONFIG_FAULT_INJECTION is not enabled
KCSAN                   : write(/sys/kernel/debug/kcsan, on) failed
LRWPANEmulation         : netlink_query_family_id failed
Leak                    : enabled
NetDevices              : enabled
NetInjection            : enabled
NicVF                   : PCI device 0000:00:11.0 is not available
SandboxAndroid          : enabled
SandboxNamespace        : NOTFAIL: sandbox fork failed.  (errno 22: Invalid argument). . process exited with status 67.
SandboxNone             : enabled
SandboxSetuid           : enabled
Swap                    : enabled
USBEmulation            : failed to chmod /dev/raw-gadget
VhciInjection           : unshare(CLONE_NEWPID): 22. spawned loop pid 9065. NOTFAIL: open /dev/vhci failed.  (errno 2: No such file or directory).
WifiEmulation           : unshare(CLONE_NEWPID): 22. spawned loop pid 9002. unshare(CLONE_NEWIPC): 22. netlink: failed to get family id for MAC80211_HWSIM: No such file or directory. NOTFAIL: netlink_query_family_id failed.  (errno 2: No such file or directory).
syscalls                : 2544/5291


2025/02/24 18:25:03 corpus                  : 274 (274 seeds), 0 to be reminimized, 0 to be resmashed
2025/02/24 18:25:12 candidates=181 corpus=0 coverage=0 exec total=708 (234/min) pending=0 reproducing=0 
2025/02/24 18:25:22 candidates=115 corpus=0 coverage=0 exec total=773 (242/min) pending=0 reproducing=0 
2025/02/24 18:25:32 candidates=19 corpus=0 coverage=0 exec total=870 (259/min) pending=0 reproducing=0 
2025/02/24 18:25:42 candidates=0 corpus=7 coverage=1845 exec total=961 (273/min) pending=0 reproducing=0 
2025/02/24 18:25:52 candidates=0 corpus=15 coverage=3080 exec total=1150 (312/min) pending=0 reproducing=0 


```





# O1


```sh
2025/02/14 16:10:48 serving rpc on tcp://46787
2025/02/14 16:10:48 serving http on http://10.189.145.186:56741
2025/02/14 16:10:49 executing adb [wait-for-device]
2025/02/14 16:10:49 skipped 7 seeds
2025/02/14 16:11:51 adb returned
2025/02/14 16:11:51 executing adb [shell reboot]
2025/02/14 16:12:02 adb returned
2025/02/14 16:12:13 executing adb [wait-for-device]
2025/02/14 16:15:00 adb returned
2025/02/14 16:15:00 executing adb [root]
2025/02/14 16:15:03 adb returned
2025/02/14 16:15:04 executing adb [wait-for-device]
2025/02/14 16:15:04 adb returned
2025/02/14 16:15:09 executing adb [shell pgrep systemui | wc -l]
2025/02/14 16:15:10 adb returned

2025/02/14 16:15:50 adb returned
2025/02/14 16:15:50 executing adb [root]
2025/02/14 16:15:50 adb returned
2025/02/14 16:15:50 root access granted
2025/02/14 16:15:50 executing adb [shell echo mtbf buck 1500 > /sys/class/xm_power/charger/charge_interface/iin_limit]
2025/02/14 16:15:50 adb returned
2025/02/14 16:15:50 charge limit set successfully
2025/02/14 16:15:50 boot completed
2025/02/14 16:15:50 executing adb [shell ls /sys/kernel/debug/kcov]
2025/02/14 16:15:50 adb returned
2025/02/14 16:15:52 failed to associate adb device db47dfe2 with console: failed to open console file: permission denied
2025/02/14 16:15:52 falling back to 'adb shell dmesg -w'
2025/02/14 16:15:52 note: some bugs may be detected as 'lost connection to test machine' with no kernel output
2025/02/14 16:15:52 associating adb device db47dfe2 with console adb
2025/02/14 16:15:52 executing adb [shell dumpsys battery | grep level:]
2025/02/14 16:15:53 adb returned
2025/02/14 16:15:53 device db47dfe2: battery level 98%, OK
2025/02/14 16:15:53 executing adb [shell ls /data/syzkaller*]
2025/02/14 16:15:53 adb returned
2025/02/14 16:15:53 executing adb [shell find /data/syzkaller* -type l -exec unlink {} \; && rm -Rf /data/syzkaller*]
2025/02/14 16:15:54 adb returned
2025/02/14 16:15:54 executing adb [shell echo 0 > /proc/sys/kernel/kptr_restrict]
2025/02/14 16:15:55 adb returned
2025/02/14 16:15:55 executing adb [reverse tcp:17730 tcp:46787]
2025/02/14 16:15:55 adb returned
2025/02/14 16:15:55 executing adb [push /home/xiaomi/Documents/syzkaller/bin/linux_arm64/syz-executor /data/syz-executor]
2025/02/14 16:15:55 adb returned
2025/02/14 16:15:55 executing adb [shell chmod +x /data/syz-executor]
2025/02/14 16:15:56 adb returned
2025/02/14 16:15:56 starting: adb shell /data/syz-executor runner 0 127.0.0.1 17730
2025/02/14 16:15:57 runner 0 connected
2025/02/14 16:17:33 machine check:
disabled the following syscalls:
acct                                        : syscall acct is not present
fanotify_init                               : syscall fanotify_init is not present
fanotify_mark                               : syscall fanotify_mark is not present
get_mempolicy                               : syscall get_mempolicy is not present
kexec_load                                  : syscall kexec_load is not present
landlock_add_rule$LANDLOCK_RULE_NET_PORT    : syscall landlock_add_rule$LANDLOCK_RULE_NET_PORT is not present
landlock_add_rule$LANDLOCK_RULE_PATH_BENEATH: syscall landlock_add_rule$LANDLOCK_RULE_NET_PORT is not present
landlock_create_ruleset                     : syscall landlock_create_ruleset is not present
landlock_restrict_self                      : syscall landlock_restrict_self is not present
lookup_dcookie                              : syscall lookup_dcookie is not present
lsm_get_self_attr                           : syscall lsm_get_self_attr is not present
lsm_list_modules                            : syscall lsm_list_modules is not present
lsm_set_self_attr                           : syscall lsm_set_self_attr is not present
map_shadow_stack                            : syscall map_shadow_stack is not present
mbind                                       : syscall mbind is not present
migrate_pages                               : syscall migrate_pages is not present
mount$9p_fd                                 : /proc/filesystems does not contain 9p
mount$9p_rdma                               : /proc/filesystems does not contain 9p
mount$9p_tcp                                : /proc/filesystems does not contain 9p
mount$9p_unix                               : /proc/filesystems does not contain 9p
mount$9p_virtio                             : /proc/filesystems does not contain 9p
mount$9p_xen                                : /proc/filesystems does not contain 9p
mount$afs                                   : /proc/filesystems does not contain afs
mount$esdfs                                 : /proc/filesystems does not contain esdfs
mount$nfs                                   : /proc/filesystems does not contain nfs
mount$nfs4                                  : /proc/filesystems does not contain nfs4
mount$pvfs2                                 : /proc/filesystems does not contain pvfs2
move_pages                                  : syscall move_pages is not present
mq_getsetattr                               : syscall mq_getsetattr is not present
mq_notify                                   : syscall mq_notify is not present
mq_open                                     : syscall mq_open is not present
mq_timedreceive                             : syscall mq_timedreceive is not present
mq_timedsend                                : syscall mq_timedsend is not present
mq_unlink                                   : syscall mq_unlink is not present
msgctl$IPC_INFO                             : syscall msgctl$IPC_INFO is not present
msgctl$IPC_RMID                             : syscall msgctl$IPC_INFO is not present
msgctl$IPC_SET                              : syscall msgctl$IPC_INFO is not present
msgctl$IPC_STAT                             : syscall msgctl$IPC_INFO is not present
msgctl$MSG_INFO                             : syscall msgctl$IPC_INFO is not present
msgctl$MSG_STAT                             : syscall msgctl$IPC_INFO is not present
msgctl$MSG_STAT_ANY                         : syscall msgctl$IPC_INFO is not present
msgget                                      : syscall msgget is not present
msgget$private                              : syscall msgget is not present
msgrcv                                      : syscall msgrcv is not present
msgsnd                                      : syscall msgsnd is not present
name_to_handle_at                           : syscall name_to_handle_at is not present
open_by_handle_at                           : syscall open_by_handle_at is not present
openat$6lowpan_control                      : failed to open /sys/kernel/debug/bluetooth/6lowpan_control: no such file or directory
openat$6lowpan_enable                       : failed to open /sys/kernel/debug/bluetooth/6lowpan_enable: no such file or directory
openat$acpi_thermal_rel                     : failed to open /dev/acpi_thermal_rel: no such file or directory
openat$adsp1                                : failed to open /dev/adsp1: no such file or directory
openat$audio                                : failed to open /dev/audio: no such file or directory
openat$audio1                               : failed to open /dev/audio1: no such file or directory
openat$autofs                               : failed to open /dev/autofs: no such file or directory
openat$bifrost                              : failed to open /dev/bifrost: no such file or directory
openat$bsg                                  : failed to open /dev/bsg: no such file or directory
openat$btrfs_control                        : failed to open /dev/btrfs-control: no such file or directory
openat$cachefiles                           : failed to open /dev/cachefiles: no such file or directory
openat$camx                                 : failed to open /dev/v4l/by-path/platform-soc@0:qcom_cam-req-mgr-video-index0: no such file or directory
openat$capi20                               : failed to open /dev/capi20: no such file or directory
openat$cdrom                                : failed to open /dev/cdrom: no such file or directory
openat$cdrom1                               : failed to open /dev/cdrom1: no such file or directory
openat$cuse                                 : failed to open /dev/cuse: no such file or directory
openat$damon_attrs                          : failed to open /sys/kernel/debug/damon/attrs: no such file or directory
openat$damon_init_regions                   : failed to open /sys/kernel/debug/damon/init_regions: no such file or directory
openat$damon_kdamond_pid                    : failed to open /sys/kernel/debug/damon/kdamond_pid: no such file or directory
openat$damon_mk_contexts                    : failed to open /sys/kernel/debug/damon/mk_contexts: no such file or directory
openat$damon_monitor_on                     : failed to open /sys/kernel/debug/damon/monitor_on: no such file or directory
openat$damon_rm_contexts                    : failed to open /sys/kernel/debug/damon/rm_contexts: no such file or directory
openat$damon_schemes                        : failed to open /sys/kernel/debug/damon/schemes: no such file or directory
openat$damon_target_ids                     : failed to open /sys/kernel/debug/damon/target_ids: no such file or directory
openat$dlm_control                          : failed to open /dev/dlm-control: no such file or directory
openat$dlm_monitor                          : failed to open /dev/dlm-monitor: no such file or directory
openat$dlm_plock                            : failed to open /dev/dlm_plock: no such file or directory
openat$dsp                                  : failed to open /dev/dsp: no such file or directory
openat$dsp1                                 : failed to open /dev/dsp1: no such file or directory
openat$fb0                                  : failed to open /dev/fb0: no such file or directory
openat$fb1                                  : failed to open /dev/fb1: no such file or directory
openat$hpet                                 : failed to open /dev/hpet: no such file or directory
openat$hwrng                                : failed to open /dev/hwrng: no such file or directory
openat$i915                                 : failed to open /dev/i915: no such file or directory
openat$img_rogue                            : failed to open /dev/img-rogue: no such file or directory
openat$iommufd                              : failed to open /dev/iommu: no such file or directory
openat$ipvs                                 : failed to open /proc/sys/net/ipv4/vs/sync_qlen_max: no such file or directory
openat$irnet                                : failed to open /dev/irnet: no such file or directory
openat$keychord                             : failed to open /dev/keychord: no such file or directory
openat$kvm                                  : failed to open /dev/kvm: no such file or directory
openat$lightnvm                             : failed to open /dev/lightnvm/control: no such file or directory
openat$mali                                 : failed to open /dev/mali0: no such file or directory
openat$md                                   : failed to open /dev/md0: no such file or directory
openat$mice                                 : failed to open /dev/input/mice: no such file or directory
openat$misdntimer                           : failed to open /dev/mISDNtimer: no such file or directory
openat$mixer                                : failed to open /dev/mixer: no such file or directory
openat$msm                                  : failed to open /dev/msm: no such file or directory
openat$nci                                  : failed to open /dev/virtual_nci: no such file or directory
openat$ndctl0                               : failed to open /dev/ndctl0: no such file or directory
openat$nmem0                                : failed to open /dev/nmem0: no such file or directory
openat$nullb                                : failed to open /dev/nullb0: no such file or directory
openat$nvme_fabrics                         : failed to open /dev/nvme-fabrics: no such file or directory
openat$nvram                                : failed to open /dev/nvram: no such file or directory
openat$ocfs2_control                        : failed to open /dev/ocfs2_control: no such file or directory
openat$pktcdvd                              : failed to open /dev/pktcdvd/control: no such file or directory
openat$pmem0                                : failed to open /dev/pmem0: no such file or directory
openat$proc_capi20                          : failed to open /proc/capi/capi20: no such file or directory
openat$proc_capi20ncci                      : failed to open /proc/capi/capi20ncci: no such file or directory
openat$proc_mixer                           : failed to open /proc/asound/card0/oss_mixer: no such file or directory
openat$proc_reclaim                         : failed to open /proc/self/reclaim: no such file or directory
openat$ptp0                                 : failed to open /dev/ptp0: no such file or directory
openat$ptp1                                 : failed to open /dev/ptp1: no such file or directory
openat$qat_adf_ctl                          : failed to open /dev/qat_adf_ctl: no such file or directory
openat$rdma_cm                              : failed to open /dev/infiniband/rdma_cm: no such file or directory
openat$sequencer                            : failed to open /dev/sequencer: no such file or directory
openat$sequencer2                           : failed to open /dev/sequencer2: no such file or directory
openat$sgx_provision                        : failed to open /dev/sgx_provision: no such file or directory
openat$smackfs_access                       : failed to open /sys/fs/smackfs/access: no such file or directory
openat$smackfs_ambient                      : failed to open /sys/fs/smackfs/ambient: no such file or directory
openat$smackfs_change_rule                  : failed to open /sys/fs/smackfs/change-rule: no such file or directory
openat$smackfs_cipso                        : failed to open /sys/fs/smackfs/cipso: no such file or directory
openat$smackfs_cipsonum                     : failed to open /sys/fs/smackfs/direct: no such file or directory
openat$smackfs_ipv6host                     : failed to open /sys/fs/smackfs/ipv6host: no such file or directory
openat$smackfs_load                         : failed to open /sys/fs/smackfs/load: no such file or directory
openat$smackfs_logging                      : failed to open /sys/fs/smackfs/logging: no such file or directory
openat$smackfs_netlabel                     : failed to open /sys/fs/smackfs/netlabel: no such file or directory
openat$smackfs_onlycap                      : failed to open /sys/fs/smackfs/onlycap: no such file or directory
openat$smackfs_ptrace                       : failed to open /sys/fs/smackfs/ptrace: no such file or directory
openat$smackfs_relabel_self                 : failed to open /sys/fs/smackfs/relabel-self: no such file or directory
openat$smackfs_revoke_subject               : failed to open /sys/fs/smackfs/revoke-subject: no such file or directory
openat$smackfs_syslog                       : failed to open /sys/fs/smackfs/syslog: no such file or directory
openat$smackfs_unconfined                   : failed to open /sys/fs/smackfs/unconfined: no such file or directory
openat$snapshot                             : failed to open /dev/snapshot: no such file or directory
openat$sndseq                               : failed to open /dev/snd/seq: no such file or directory
openat$sr                                   : failed to open /dev/sr0: no such file or directory
openat$sw_sync                              : failed to open /sys/kernel/debug/sync/sw_sync: no such file or directory
openat$sw_sync_info                         : failed to open /sys/kernel/debug/sync/info: no such file or directory
openat$sysctl                               : failed to open /sys/kernel/mm/ksm/run: no such file or directory
openat$tlk_device                           : failed to open /dev/tlk_device: no such file or directory
openat$trusty                               : failed to open /dev/trusty-ipc-dev0: no such file or directory
openat$trusty_avb                           : failed to open /dev/trusty-ipc-dev0: no such file or directory
openat$trusty_gatekeeper                    : failed to open /dev/trusty-ipc-dev0: no such file or directory
openat$trusty_hwkey                         : failed to open /dev/trusty-ipc-dev0: no such file or directory
openat$trusty_hwrng                         : failed to open /dev/trusty-ipc-dev0: no such file or directory
openat$trusty_km                            : failed to open /dev/trusty-ipc-dev0: no such file or directory
openat$trusty_km_secure                     : failed to open /dev/trusty-ipc-dev0: no such file or directory
openat$trusty_storage                       : failed to open /dev/trusty-ipc-dev0: no such file or directory
openat$tty                                  : failed to open /dev/tty: no such device or address
openat$ttyS3                                : failed to open /dev/ttyS3: no such file or directory
openat$ttyprintk                            : failed to open /dev/ttyprintk: no such file or directory
openat$ubi_ctrl                             : failed to open /dev/ubi_ctrl: no such file or directory
openat$udambuf                              : failed to open /dev/udmabuf: no such file or directory
openat$userio                               : failed to open /dev/userio: no such file or directory
openat$uverbs0                              : failed to open /dev/infiniband/uverbs0: no such file or directory
openat$vcs                                  : failed to open /dev/vcs: no such file or directory
openat$vcsa                                 : failed to open /dev/vcsa: no such file or directory
openat$vcsu                                 : failed to open /dev/vcsu: no such file or directory
openat$vfio                                 : failed to open /dev/vfio/vfio: no such file or directory
openat$vga_arbiter                          : failed to open /dev/vga_arbiter: no such file or directory
openat$vicodec0                             : failed to open /dev/video36: no such file or directory
openat$vicodec1                             : failed to open /dev/video37: no such file or directory
openat$vim2m                                : failed to open /dev/vim2m: no such file or directory
openat$vimc0                                : failed to open /dev/video0: operation already in progress
openat$vimc1                                : failed to open /dev/video1: operation already in progress
openat$vimc2                                : failed to open /dev/video2: no such file or directory
openat$vmci                                 : failed to open /dev/vmci: no such file or directory
openat$vnet                                 : failed to open /dev/vhost-net: no such file or directory
openat$vtpm                                 : failed to open /dev/vtpmx: no such file or directory
openat$xenevtchn                            : failed to open /dev/xen/evtchn: no such file or directory
openat$yama_ptrace_scope                    : failed to open /proc/sys/kernel/yama/ptrace_scope: no such file or directory
openat$zygote                               : failed to open /dev/socket/zygote: no such device or address
pkey_alloc                                  : pkey_alloc(0x0, 0x0) failed: function not implemented
pkey_free                                   : syscall pkey_free is not present
pkey_mprotect                               : syscall pkey_mprotect is not present
rseq                                        : syscall rseq is not present
semget                                      : syscall semget is not present
semget$private                              : syscall semget is not present
semop                                       : syscall semop is not present
semtimedop                                  : syscall semtimedop is not present
set_mempolicy                               : syscall set_mempolicy is not present
set_mempolicy_home_node                     : syscall set_mempolicy_home_node is not present
shmat                                       : syscall shmat is not present
shmctl$IPC_INFO                             : syscall shmctl$IPC_INFO is not present
shmctl$IPC_RMID                             : syscall shmctl$IPC_INFO is not present
shmctl$IPC_SET                              : syscall shmctl$IPC_INFO is not present
shmctl$IPC_STAT                             : syscall shmctl$IPC_INFO is not present
shmctl$SHM_INFO                             : syscall shmctl$IPC_INFO is not present
shmctl$SHM_LOCK                             : syscall shmctl$IPC_INFO is not present
shmctl$SHM_STAT                             : syscall shmctl$IPC_INFO is not present
shmctl$SHM_STAT_ANY                         : syscall shmctl$IPC_INFO is not present
shmctl$SHM_UNLOCK                           : syscall shmctl$IPC_INFO is not present
shmdt                                       : syscall shmdt is not present
shmget                                      : syscall shmget is not present
shmget$private                              : syscall shmget is not present
socket$alg                                  : socket$alg(0x26, 0x5, 0x0) failed: address family not supported by protocol
socket$caif_seqpacket                       : socket$caif_seqpacket(0x25, 0x5, 0x0) failed: address family not supported by protocol
socket$caif_stream                          : socket$caif_stream(0x25, 0x1, 0x0) failed: address family not supported by protocol
socket$can_j1939                            : socket$can_j1939(0x1d, 0x2, 0x7) failed: protocol not supported
socket$hf                                   : socket$hf(0x13, 0x2, 0x0) failed: address family not supported by protocol
socket$inet6_dccp                           : socket$inet6_dccp(0xa, 0x6, 0x0) failed: socket type not supported
socket$inet6_mptcp                          : socket$inet6_mptcp(0xa, 0x1, 0x106) failed: protocol not supported
socket$inet6_sctp                           : socket$inet6_sctp(0xa, 0x1, 0x84) failed: protocol not supported
socket$inet_dccp                            : socket$inet_dccp(0x2, 0x6, 0x0) failed: socket type not supported
socket$inet_mptcp                           : socket$inet_mptcp(0x2, 0x1, 0x106) failed: protocol not supported
socket$inet_sctp                            : socket$inet_sctp(0x2, 0x1, 0x84) failed: protocol not supported
socket$inet_smc                             : socket$inet_smc(0x2b, 0x1, 0x0) failed: address family not supported by protocol
socket$isdn                                 : socket$isdn(0x22, 0x3, 0x0) failed: address family not supported by protocol
socket$isdn_base                            : socket$isdn_base(0x22, 0x3, 0x0) failed: address family not supported by protocol
socket$kcm                                  : socket$kcm(0x29, 0x2, 0x0) failed: address family not supported by protocol
socket$l2tp                                 : socket$l2tp(0x2, 0x2, 0x73) failed: protocol not supported
socket$l2tp6                                : socket$l2tp6(0xa, 0x2, 0x73) failed: protocol not supported
socket$nl_crypto                            : socket$nl_crypto(0x10, 0x3, 0x15) failed: protocol not supported
socket$nl_rdma                              : socket$nl_rdma(0x10, 0x3, 0x14) failed: protocol not supported
socket$phonet                               : socket$phonet(0x23, 0x2, 0x1) failed: address family not supported by protocol
socket$phonet_pipe                          : socket$phonet_pipe(0x23, 0x5, 0x2) failed: address family not supported by protocol
socket$pppoe                                : socket$pppoe(0x18, 0x1, 0x0) failed: protocol not supported
socket$rds                                  : socket$rds(0x15, 0x5, 0x0) failed: address family not supported by protocol
socket$rxrpc                                : socket$rxrpc(0x21, 0x2, 0x0) failed: address family not supported by protocol
socket$vsock_dgram                          : socket$vsock_dgram(0x28, 0x2, 0x0) failed: no such device
syz_80211_inject_frame                      : root failed to open /sys/class/mac80211_hwsim/: no such file or directory
syz_80211_join_ibss                         : root failed to open /sys/class/mac80211_hwsim/: no such file or directory
syz_emit_vhci                               : root failed to open /dev/vhci: no such file or directory
syz_init_net_socket$ax25                    : syz_init_net_socket$ax25(0x3, 0x2, 0x0) failed: address family not supported by protocol
syz_init_net_socket$bt_bnep                 : syz_init_net_socket$bt_bnep(0x1f, 0x3, 0x4) failed: protocol not supported
syz_init_net_socket$bt_cmtp                 : syz_init_net_socket$bt_cmtp(0x1f, 0x3, 0x5) failed: protocol not supported
syz_init_net_socket$llc                     : syz_init_net_socket$llc(0x1a, 0x1, 0x0) failed: address family not supported by protocol
syz_init_net_socket$netrom                  : syz_init_net_socket$netrom(0x6, 0x5, 0x0) failed: address family not supported by protocol
syz_init_net_socket$nl_rdma                 : syz_init_net_socket$nl_rdma(0x10, 0x3, 0x14) failed: protocol not supported
syz_init_net_socket$rose                    : syz_init_net_socket$rose(0xb, 0x5, 0x0) failed: address family not supported by protocol
syz_init_net_socket$x25                     : syz_init_net_socket$x25(0x9, 0x5, 0x0) failed: address family not supported by protocol
syz_kvm_setup_cpu$ppc64                     : unsupported arch
syz_kvm_setup_cpu$x86                       : unsupported arch
syz_mount_image$adfs                        : /proc/filesystems does not contain adfs
syz_mount_image$affs                        : /proc/filesystems does not contain affs
syz_mount_image$bcachefs                    : /proc/filesystems does not contain bcachefs
syz_mount_image$befs                        : /proc/filesystems does not contain befs
syz_mount_image$bfs                         : /proc/filesystems does not contain bfs
syz_mount_image$btrfs                       : /proc/filesystems does not contain btrfs
syz_mount_image$cramfs                      : /proc/filesystems does not contain cramfs
syz_mount_image$efs                         : /proc/filesystems does not contain efs
syz_mount_image$gfs2                        : /proc/filesystems does not contain gfs2
syz_mount_image$gfs2meta                    : /proc/filesystems does not contain gfs2meta
syz_mount_image$hfs                         : /proc/filesystems does not contain hfs
syz_mount_image$hfsplus                     : /proc/filesystems does not contain hfsplus
syz_mount_image$hpfs                        : /proc/filesystems does not contain hpfs
syz_mount_image$iso9660                     : /proc/filesystems does not contain iso9660
syz_mount_image$jffs2                       : /proc/filesystems does not contain jffs2
syz_mount_image$jfs                         : /proc/filesystems does not contain jfs
syz_mount_image$minix                       : /proc/filesystems does not contain minix
syz_mount_image$nilfs2                      : /proc/filesystems does not contain nilfs2
syz_mount_image$ntfs                        : /proc/filesystems does not contain ntfs
syz_mount_image$ntfs3                       : /proc/filesystems does not contain ntfs3
syz_mount_image$ocfs2                       : /proc/filesystems does not contain ocfs2
syz_mount_image$omfs                        : /proc/filesystems does not contain omfs
syz_mount_image$qnx4                        : /proc/filesystems does not contain qnx4
syz_mount_image$qnx6                        : /proc/filesystems does not contain qnx6
syz_mount_image$reiserfs                    : /proc/filesystems does not contain reiserfs
syz_mount_image$romfs                       : /proc/filesystems does not contain romfs
syz_mount_image$squashfs                    : /proc/filesystems does not contain squashfs
syz_mount_image$sysv                        : /proc/filesystems does not contain sysv
syz_mount_image$ubifs                       : /proc/filesystems does not contain ubifs
syz_mount_image$udf                         : /proc/filesystems does not contain udf
syz_mount_image$ufs                         : /proc/filesystems does not contain ufs
syz_mount_image$v7                          : /proc/filesystems does not contain v7
syz_mount_image$vxfs                        : /proc/filesystems does not contain vxfs
syz_mount_image$xfs                         : /proc/filesystems does not contain xfs
syz_mount_image$zonefs                      : /proc/filesystems does not contain zonefs
syz_open_dev$MSR                            : failed to open /dev/cpu/#/msr: no such file or directory
syz_open_dev$admmidi                        : failed to open /dev/admmidi#: no such file or directory
syz_open_dev$amidi                          : failed to open /dev/amidi#: no such file or directory
syz_open_dev$audion                         : failed to open /dev/audio#: no such file or directory
syz_open_dev$cec                            : failed to open /dev/cec#: no such file or directory
syz_open_dev$dmmidi                         : failed to open /dev/dmmidi#: no such file or directory
syz_open_dev$dricontrol                     : failed to open /dev/dri/controlD#: no such file or directory
syz_open_dev$drirender                      : failed to open /dev/dri/renderD#: no such file or directory
syz_open_dev$floppy                         : failed to open /dev/fd#: no such file or directory
syz_open_dev$hiddev                         : failed to open /dev/usb/hiddev#: no such file or directory
syz_open_dev$hidraw                         : failed to open /dev/hidraw#: no such file or directory
syz_open_dev$ircomm                         : failed to open /dev/ircomm#: no such file or directory
syz_open_dev$loop                           : failed to open /dev/loop#: no such file or directory
syz_open_dev$midi                           : failed to open /dev/midi#: no such file or directory
syz_open_dev$mouse                          : failed to open /dev/input/mouse#: no such file or directory
syz_open_dev$ndb                            : failed to open /dev/nbd#: no such file or directory
syz_open_dev$radio                          : failed to open /dev/radio#: no such file or directory
syz_open_dev$sndhw                          : failed to open /dev/snd/hwC#D#: no such file or directory
syz_open_dev$sndmidi                        : failed to open /dev/snd/midiC#D#: no such file or directory
syz_open_dev$sndpcmc                        : failed to open /dev/snd/pcmC#D#c: no such file or directory
syz_open_dev$swradio                        : failed to open /dev/swradio#: no such file or directory
syz_open_dev$usbfs                          : failed to open /dev/bus/usb/00#/00#: no such file or directory
syz_open_dev$usbmon                         : failed to open /dev/usbmon#: no such file or directory
syz_open_dev$vbi                            : failed to open /dev/vbi#: no such file or directory
syz_open_dev$vcsa                           : failed to open /dev/vcsa#: no such file or directory
syz_open_dev$vcsn                           : failed to open /dev/vcs#: no such file or directory
syz_open_dev$vcsu                           : failed to open /dev/vcsu#: no such file or directory
syz_open_dev$video                          : failed to open /dev/video#: operation already in progress
syz_open_dev$vim2m                          : failed to open /dev/video#: operation already in progress
syz_open_dev$vivid                          : failed to open /dev/video#: no such file or directory
syz_pkey_set                                : pkey_alloc(0x0, 0x0) failed: function not implemented
syz_usb_connect                             : root failed to open /dev/raw-gadget: no such file or directory
syz_usb_connect$cdc_ecm                     : root failed to open /dev/raw-gadget: no such file or directory
syz_usb_connect$cdc_ncm                     : root failed to open /dev/raw-gadget: no such file or directory
syz_usb_connect$hid                         : root failed to open /dev/raw-gadget: no such file or directory
syz_usb_connect$printer                     : root failed to open /dev/raw-gadget: no such file or directory
syz_usb_connect$uac1                        : root failed to open /dev/raw-gadget: no such file or directory
syz_usb_connect_ath9k                       : root failed to open /dev/raw-gadget: no such file or directory
syz_usb_control_io                          : root failed to open /dev/raw-gadget: no such file or directory
syz_usb_control_io$cdc_ecm                  : root failed to open /dev/raw-gadget: no such file or directory
syz_usb_control_io$cdc_ncm                  : root failed to open /dev/raw-gadget: no such file or directory
syz_usb_control_io$hid                      : root failed to open /dev/raw-gadget: no such file or directory
syz_usb_control_io$printer                  : root failed to open /dev/raw-gadget: no such file or directory
syz_usb_control_io$uac1                     : root failed to open /dev/raw-gadget: no such file or directory
syz_usb_disconnect                          : root failed to open /dev/raw-gadget: no such file or directory
syz_usb_ep_read                             : root failed to open /dev/raw-gadget: no such file or directory
syz_usb_ep_write                            : root failed to open /dev/raw-gadget: no such file or directory
syz_usb_ep_write$ath9k_ep1                  : root failed to open /dev/raw-gadget: no such file or directory
syz_usb_ep_write$ath9k_ep2                  : root failed to open /dev/raw-gadget: no such file or directory
syz_usbip_server_init                       : failed to open /sys/devices/platform/vhci_hcd.0/attach: no such file or directory

transitively disabled the following syscalls (missing resource [creating syscalls]):
accept$alg                                  : sock_alg [socket$alg]
accept$ax25                                 : sock_ax25 [accept$ax25 accept4$ax25 syz_init_net_socket$ax25]
accept$netrom                               : sock_netrom [accept$netrom accept4$netrom syz_init_net_socket$netrom]
accept$phonet_pipe                          : sock_phonet_pipe [accept$phonet_pipe accept4$phonet_pipe socket$phonet_pipe]
accept4$alg                                 : sock_alg [socket$alg]
accept4$ax25                                : sock_ax25 [accept$ax25 accept4$ax25 syz_init_net_socket$ax25]
accept4$llc                                 : sock_llc [accept4$llc syz_init_net_socket$llc]
accept4$netrom                              : sock_netrom [accept$netrom accept4$netrom syz_init_net_socket$netrom]
accept4$phonet_pipe                         : sock_phonet_pipe [accept$phonet_pipe accept4$phonet_pipe socket$phonet_pipe]
accept4$rose                                : sock_rose [accept4$rose syz_init_net_socket$rose]
accept4$x25                                 : sock_x25 [accept4$x25 syz_init_net_socket$x25]
bind                                        : nfc_dev_id [ioctl$IOCTL_GET_NCIDEV_IDX]
bind$alg                                    : sock_alg [socket$alg]
bind$ax25                                   : sock_ax25 [accept$ax25 accept4$ax25 syz_init_net_socket$ax25]
bind$can_j1939                              : sock_can_j1939 [socket$can_j1939]
bind$isdn                                   : sock_isdn [socket$isdn]
bind$isdn_base                              : sock_isdn_base [socket$isdn_base]
bind$l2tp                                   : sock_l2tp [socket$l2tp]
bind$l2tp6                                  : sock_l2tp6 [socket$l2tp6]
bind$llc                                    : sock_llc [accept4$llc syz_init_net_socket$llc]
bind$netrom                                 : sock_netrom [accept$netrom accept4$netrom syz_init_net_socket$netrom]
bind$nfc_llcp                               : nfc_dev_id [ioctl$IOCTL_GET_NCIDEV_IDX]
bind$phonet                                 : sock_phonet [accept$phonet_pipe accept4$phonet_pipe socket$phonet socket$phonet_pipe]
bind$rds                                    : sock_rds [socket$rds]
bind$rose                                   : sock_rose [accept4$rose syz_init_net_socket$rose]
bind$rxrpc                                  : sock_rxrpc [socket$rxrpc]
bind$vsock_dgram                            : sock_vsock_dgram [socket$vsock_dgram]
bind$x25                                    : sock_x25 [accept4$x25 syz_init_net_socket$x25]
close$ibv_device                            : fd_rdma [openat$uverbs0]
connect                                     : nfc_dev_id [ioctl$IOCTL_GET_NCIDEV_IDX]
connect$ax25                                : sock_ax25 [accept$ax25 accept4$ax25 syz_init_net_socket$ax25]
connect$caif                                : sock_caif [socket$caif_seqpacket socket$caif_stream]
connect$can_j1939                           : sock_can_j1939 [socket$can_j1939]
connect$hf                                  : sock_hf [socket$hf]
connect$l2tp                                : sock_l2tp [socket$l2tp]
connect$l2tp6                               : sock_l2tp6 [socket$l2tp6]
connect$llc                                 : sock_llc [accept4$llc syz_init_net_socket$llc]
connect$netrom                              : sock_netrom [accept$netrom accept4$netrom syz_init_net_socket$netrom]
connect$nfc_llcp                            : nfc_dev_id [ioctl$IOCTL_GET_NCIDEV_IDX]
connect$nfc_raw                             : nfc_dev_id [ioctl$IOCTL_GET_NCIDEV_IDX]
connect$phonet_pipe                         : sock_phonet_pipe [accept$phonet_pipe accept4$phonet_pipe socket$phonet_pipe]
connect$pppoe                               : sock_pppoe [socket$pppoe]
connect$rds                                 : sock_rds [socket$rds]
connect$rose                                : sock_rose [accept4$rose syz_init_net_socket$rose]
connect$rxrpc                               : sock_rxrpc [socket$rxrpc]
connect$vsock_dgram                         : sock_vsock_dgram [socket$vsock_dgram]
connect$x25                                 : sock_x25 [accept4$x25 syz_init_net_socket$x25]
getpeername$ax25                            : sock_ax25 [accept$ax25 accept4$ax25 syz_init_net_socket$ax25]
getpeername$l2tp                            : sock_l2tp [socket$l2tp]
getpeername$l2tp6                           : sock_l2tp6 [socket$l2tp6]
getpeername$llc                             : sock_llc [accept4$llc syz_init_net_socket$llc]
getpeername$netrom                          : sock_netrom [accept$netrom accept4$netrom syz_init_net_socket$netrom]
getsockname$ax25                            : sock_ax25 [accept$ax25 accept4$ax25 syz_init_net_socket$ax25]
getsockname$l2tp                            : sock_l2tp [socket$l2tp]
getsockname$l2tp6                           : sock_l2tp6 [socket$l2tp6]
getsockname$llc                             : sock_llc [accept4$llc syz_init_net_socket$llc]
getsockname$netrom                          : sock_netrom [accept$netrom accept4$netrom syz_init_net_socket$netrom]
getsockopt$MISDN_TIME_STAMP                 : sock_isdn [socket$isdn]
getsockopt$PNPIPE_ENCAP                     : sock_phonet_pipe [accept$phonet_pipe accept4$phonet_pipe socket$phonet_pipe]
getsockopt$PNPIPE_HANDLE                    : sock_phonet_pipe [accept$phonet_pipe accept4$phonet_pipe socket$phonet_pipe]
getsockopt$PNPIPE_IFINDEX                   : sock_phonet_pipe [accept$phonet_pipe accept4$phonet_pipe socket$phonet_pipe]
getsockopt$PNPIPE_INITSTATE                 : sock_phonet_pipe [accept$phonet_pipe accept4$phonet_pipe socket$phonet_pipe]
getsockopt$SO_J1939_ERRQUEUE                : sock_can_j1939 [socket$can_j1939]
getsockopt$SO_J1939_PROMISC                 : sock_can_j1939 [socket$can_j1939]
getsockopt$SO_J1939_SEND_PRIO               : sock_can_j1939 [socket$can_j1939]
getsockopt$X25_QBITINCL                     : sock_x25 [accept4$x25 syz_init_net_socket$x25]
getsockopt$ax25_int                         : sock_ax25 [accept$ax25 accept4$ax25 syz_init_net_socket$ax25]
getsockopt$inet6_dccp_buf                   : sock_dccp6 [socket$inet6_dccp]
getsockopt$inet6_dccp_int                   : sock_dccp6 [socket$inet6_dccp]
getsockopt$inet6_mptcp_buf                  : sock_mptcp6 [socket$inet6_mptcp]
getsockopt$inet_dccp_buf                    : sock_dccp [socket$inet_dccp]
getsockopt$inet_dccp_int                    : sock_dccp [socket$inet_dccp]
getsockopt$inet_mptcp_buf                   : sock_mptcp [socket$inet_mptcp]
getsockopt$inet_sctp6_SCTP_ADAPTATION_LAYER : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_ASSOCINFO        : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_AUTH_ACTIVE_KEY  : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_AUTOCLOSE        : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_AUTO_ASCONF      : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_CONTEXT          : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_DEFAULT_PRINFO   : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_DEFAULT_SEND_PARAM: sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_DEFAULT_SNDINFO  : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_DELAYED_SACK     : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_DISABLE_FRAGMENTS: sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_ENABLE_STREAM_RESET: sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_EVENTS           : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_FRAGMENT_INTERLEAVE: sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_GET_ASSOC_ID_LIST: sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_GET_ASSOC_NUMBER : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_GET_ASSOC_STATS  : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_GET_LOCAL_ADDRS  : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_GET_PEER_ADDRS   : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_GET_PEER_ADDR_INFO: sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_HMAC_IDENT       : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_INITMSG          : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_I_WANT_MAPPED_V4_ADDR: sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_LOCAL_AUTH_CHUNKS: sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_MAXSEG           : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_MAX_BURST        : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_NODELAY          : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_PARTIAL_DELIVERY_POINT: sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_PEER_ADDR_PARAMS : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_PEER_ADDR_THLDS  : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_PEER_AUTH_CHUNKS : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_PRIMARY_ADDR     : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_PR_ASSOC_STATUS  : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_PR_STREAM_STATUS : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_PR_SUPPORTED     : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_RECONFIG_SUPPORTED: sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_RECVNXTINFO      : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_RECVRCVINFO      : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_RESET_STREAMS    : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_RTOINFO          : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_SOCKOPT_CONNECTX3: sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_SOCKOPT_PEELOFF  : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_STATUS           : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_STREAM_SCHEDULER : sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp6_SCTP_STREAM_SCHEDULER_VALUE: sock_sctp6 [socket$inet6_sctp]
getsockopt$inet_sctp_SCTP_ADAPTATION_LAYER  : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_ASSOCINFO         : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_AUTH_ACTIVE_KEY   : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_AUTOCLOSE         : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_AUTO_ASCONF       : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_CONTEXT           : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_DEFAULT_PRINFO    : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_DEFAULT_SEND_PARAM: sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_DEFAULT_SNDINFO   : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_DELAYED_SACK      : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_DISABLE_FRAGMENTS : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_ENABLE_STREAM_RESET: sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_EVENTS            : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_FRAGMENT_INTERLEAVE: sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_GET_ASSOC_ID_LIST : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_GET_ASSOC_NUMBER  : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_GET_ASSOC_STATS   : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_GET_LOCAL_ADDRS   : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_GET_PEER_ADDRS    : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_GET_PEER_ADDR_INFO: sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_HMAC_IDENT        : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_INITMSG           : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_I_WANT_MAPPED_V4_ADDR: sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_LOCAL_AUTH_CHUNKS : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_MAXSEG            : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_MAX_BURST         : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_NODELAY           : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_PARTIAL_DELIVERY_POINT: sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_PEER_ADDR_PARAMS  : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_PEER_ADDR_THLDS   : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_PEER_AUTH_CHUNKS  : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_PRIMARY_ADDR      : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_PR_ASSOC_STATUS   : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_PR_STREAM_STATUS  : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_PR_SUPPORTED      : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_RECONFIG_SUPPORTED: sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_RECVNXTINFO       : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_RECVRCVINFO       : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_RESET_STREAMS     : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_RTOINFO           : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_SOCKOPT_CONNECTX3 : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_SOCKOPT_PEELOFF   : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_STATUS            : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_STREAM_SCHEDULER  : sock_sctp [socket$inet_sctp]
getsockopt$inet_sctp_SCTP_STREAM_SCHEDULER_VALUE: sock_sctp [socket$inet_sctp]
getsockopt$kcm_KCM_RECV_DISABLE             : sock_kcm [socket$kcm]
getsockopt$llc_int                          : sock_llc [accept4$llc syz_init_net_socket$llc]
getsockopt$netrom_NETROM_IDLE               : sock_netrom [accept$netrom accept4$netrom syz_init_net_socket$netrom]
getsockopt$netrom_NETROM_N2                 : sock_netrom [accept$netrom accept4$netrom syz_init_net_socket$netrom]
getsockopt$netrom_NETROM_T1                 : sock_netrom [accept$netrom accept4$netrom syz_init_net_socket$netrom]
getsockopt$netrom_NETROM_T2                 : sock_netrom [accept$netrom accept4$netrom syz_init_net_socket$netrom]
getsockopt$netrom_NETROM_T4                 : sock_netrom [accept$netrom accept4$netrom syz_init_net_socket$netrom]
getsockopt$rose                             : sock_rose [accept4$rose syz_init_net_socket$rose]
ioctl$ACPI_THERMAL_GET_ART                  : fd_acpi_thermal_rel [openat$acpi_thermal_rel]
ioctl$ACPI_THERMAL_GET_ART_COUNT            : fd_acpi_thermal_rel [openat$acpi_thermal_rel]
ioctl$ACPI_THERMAL_GET_ART_LEN              : fd_acpi_thermal_rel [openat$acpi_thermal_rel]
ioctl$ACPI_THERMAL_GET_TRT                  : fd_acpi_thermal_rel [openat$acpi_thermal_rel]
ioctl$ACPI_THERMAL_GET_TRT_COUNT            : fd_acpi_thermal_rel [openat$acpi_thermal_rel]
ioctl$ACPI_THERMAL_GET_TRT_LEN              : fd_acpi_thermal_rel [openat$acpi_thermal_rel]
ioctl$AUTOFS_DEV_IOCTL_ASKUMOUNT            : fd_autofs [openat$autofs]
ioctl$AUTOFS_DEV_IOCTL_CATATONIC            : fd_autofs [openat$autofs]
ioctl$AUTOFS_DEV_IOCTL_CLOSEMOUNT           : fd_autofs [openat$autofs]
ioctl$AUTOFS_DEV_IOCTL_EXPIRE               : fd_autofs [openat$autofs]
ioctl$AUTOFS_DEV_IOCTL_FAIL                 : fd_autofs [openat$autofs]
ioctl$AUTOFS_DEV_IOCTL_ISMOUNTPOINT         : fd_autofs [openat$autofs]
ioctl$AUTOFS_DEV_IOCTL_OPENMOUNT            : fd_autofs [openat$autofs]
ioctl$AUTOFS_DEV_IOCTL_PROTOSUBVER          : fd_autofs [openat$autofs]
ioctl$AUTOFS_DEV_IOCTL_PROTOVER             : fd_autofs [openat$autofs]
ioctl$AUTOFS_DEV_IOCTL_READY                : fd_autofs [openat$autofs]
ioctl$AUTOFS_DEV_IOCTL_REQUESTER            : fd_autofs [openat$autofs]
ioctl$AUTOFS_DEV_IOCTL_SETPIPEFD            : fd_autofs [openat$autofs]
ioctl$AUTOFS_DEV_IOCTL_TIMEOUT              : fd_autofs [openat$autofs]
ioctl$AUTOFS_DEV_IOCTL_VERSION              : fd_autofs [openat$autofs]
ioctl$BLKALIGNOFF                           : fd_block [openat$md openat$nullb openat$pmem0 syz_open_dev$loop syz_open_dev$ndb]
ioctl$BLKBSZGET                             : fd_block [openat$md openat$nullb openat$pmem0 syz_open_dev$loop syz_open_dev$ndb]
ioctl$BLKBSZSET                             : fd_block [openat$md openat$nullb openat$pmem0 syz_open_dev$loop syz_open_dev$ndb]
ioctl$BLKDISCARD                            : fd_block [openat$md openat$nullb openat$pmem0 syz_open_dev$loop syz_open_dev$ndb]
ioctl$BLKFLSBUF                             : fd_block [openat$md openat$nullb openat$pmem0 syz_open_dev$loop syz_open_dev$ndb]
ioctl$BLKFRASET                             : fd_block [openat$md openat$nullb openat$pmem0 syz_open_dev$loop syz_open_dev$ndb]
ioctl$BLKGETDISKSEQ                         : fd_block [openat$md openat$nullb openat$pmem0 syz_open_dev$loop syz_open_dev$ndb]
ioctl$BLKGETSIZE                            : fd_block [openat$md openat$nullb openat$pmem0 syz_open_dev$loop syz_open_dev$ndb]
ioctl$BLKGETSIZE64                          : fd_block [openat$md openat$nullb openat$pmem0 syz_open_dev$loop syz_open_dev$ndb]
ioctl$BLKIOMIN                              : fd_block [openat$md openat$nullb openat$pmem0 syz_open_dev$loop syz_open_dev$ndb]
ioctl$BLKIOOPT                              : fd_block [openat$md openat$nullb openat$pmem0 syz_open_dev$loop syz_open_dev$ndb]
ioctl$BLKPBSZGET                            : fd_block [openat$md openat$nullb openat$pmem0 syz_open_dev$loop syz_open_dev$ndb]
ioctl$BLKPG                                 : fd_block [openat$md openat$nullb openat$pmem0 syz_open_dev$loop syz_open_dev$ndb]
ioctl$BLKRAGET                              : fd_block [openat$md openat$nullb openat$pmem0 syz_open_dev$loop syz_open_dev$ndb]
ioctl$BLKREPORTZONE                         : fd_block [openat$md openat$nullb openat$pmem0 syz_open_dev$loop syz_open_dev$ndb]
ioctl$BLKRESETZONE                          : fd_block [openat$md openat$nullb openat$pmem0 syz_open_dev$loop syz_open_dev$ndb]
ioctl$BLKROGET                              : fd_block [openat$md openat$nullb openat$pmem0 syz_open_dev$loop syz_open_dev$ndb]
ioctl$BLKROSET                              : fd_block [openat$md openat$nullb openat$pmem0 syz_open_dev$loop syz_open_dev$ndb]
ioctl$BLKROTATIONAL                         : fd_block [openat$md openat$nullb openat$pmem0 syz_open_dev$loop syz_open_dev$ndb]
ioctl$BLKRRPART                             : fd_block [openat$md openat$nullb openat$pmem0 syz_open_dev$loop syz_open_dev$ndb]
ioctl$BLKSECDISCARD                         : fd_block [openat$md openat$nullb openat$pmem0 syz_open_dev$loop syz_open_dev$ndb]
ioctl$BLKSECTGET                            : fd_block [openat$md openat$nullb openat$pmem0 syz_open_dev$loop syz_open_dev$ndb]
ioctl$BLKZEROOUT                            : fd_block [openat$md openat$nullb openat$pmem0 syz_open_dev$loop syz_open_dev$ndb]
ioctl$CAM_REQ_MGR_ALLOC_BUF                 : fd_camx [openat$camx]
ioctl$CAM_REQ_MGR_CREATE_SESSION            : fd_camx [openat$camx]
ioctl$CAM_REQ_MGR_DESTROY_SESSION           : fd_camx [openat$camx]
ioctl$CAM_REQ_MGR_FLUSH_REQ                 : fd_camx [openat$camx]
ioctl$CAM_REQ_MGR_LINK                      : fd_camx [openat$camx]
ioctl$CAM_REQ_MGR_LINK_CONTROL              : fd_camx [openat$camx]
ioctl$CAM_REQ_MGR_LINK_V2                   : fd_camx [openat$camx]
ioctl$CAM_REQ_MGR_MAP_BUF                   : fd_camx [openat$camx]
ioctl$CAM_REQ_MGR_RELEASE_BUF               : fd_camx [openat$camx]
ioctl$CAM_REQ_MGR_REQUEST_DUMP              : fd_camx [openat$camx]
ioctl$CAM_REQ_MGR_SCHED_REQ                 : fd_camx [openat$camx]
ioctl$CAM_REQ_MGR_SYNC_MODE                 : fd_camx [openat$camx]
ioctl$CAM_REQ_MGR_UNLINK                    : fd_camx [openat$camx]
ioctl$CAPI_CLR_FLAGS                        : fd_capi20 [openat$capi20]
ioctl$CAPI_GET_ERRCODE                      : fd_capi20 [openat$capi20]
ioctl$CAPI_GET_FLAGS                        : fd_capi20 [openat$capi20]
ioctl$CAPI_GET_MANUFACTURER                 : fd_capi20 [openat$capi20]
ioctl$CAPI_GET_PROFILE                      : fd_capi20 [openat$capi20]
ioctl$CAPI_GET_SERIAL                       : fd_capi20 [openat$capi20]
ioctl$CAPI_INSTALLED                        : fd_capi20 [openat$capi20]
ioctl$CAPI_MANUFACTURER_CMD                 : fd_capi20 [openat$capi20]
ioctl$CAPI_NCCI_GETUNIT                     : fd_capi20 [openat$capi20]
ioctl$CAPI_NCCI_OPENCOUNT                   : fd_capi20 [openat$capi20]
ioctl$CAPI_REGISTER                         : fd_capi20 [openat$capi20]
ioctl$CAPI_SET_FLAGS                        : fd_capi20 [openat$capi20]
ioctl$CDROMCLOSETRAY                        : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROMEJECT                            : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROMEJECT_SW                         : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROMGETSPINDOWN                      : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROMMULTISESSION                     : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROMPAUSE                            : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROMPLAYBLK                          : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROMPLAYMSF                          : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROMPLAYTRKIND                       : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROMREADALL                          : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROMREADAUDIO                        : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROMREADCOOKED                       : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROMREADMODE1                        : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROMREADMODE2                        : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROMREADRAW                          : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROMREADTOCENTRY                     : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROMREADTOCHDR                       : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROMRESET                            : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROMRESUME                           : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROMSEEK                             : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROMSETSPINDOWN                      : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROMSTART                            : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROMSTOP                             : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROMSUBCHNL                          : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROMVOLCTRL                          : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROMVOLREAD                          : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROM_CHANGER_NSLOTS                  : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROM_CLEAR_OPTIONS                   : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROM_DEBUG                           : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROM_DISC_STATUS                     : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROM_GET_CAPABILITY                  : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROM_GET_MCN                         : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROM_LAST_WRITTEN                    : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROM_LOCKDOOR                        : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROM_MEDIA_CHANGED                   : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROM_NEXT_WRITABLE                   : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROM_SELECT_DISK                     : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROM_SELECT_SPEED                    : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROM_SEND_PACKET                     : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROM_SET_OPTIONS                     : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CDROM_TIMED_MEDIA_CHANGE              : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$CEC_ADAP_G_CAPS                       : fd_cec [syz_open_dev$cec]
ioctl$CEC_ADAP_G_CONNECTOR_INFO             : fd_cec [syz_open_dev$cec]
ioctl$CEC_ADAP_G_LOG_ADDRS                  : fd_cec [syz_open_dev$cec]
ioctl$CEC_ADAP_G_PHYS_ADDR                  : fd_cec [syz_open_dev$cec]
ioctl$CEC_ADAP_S_LOG_ADDRS                  : fd_cec [syz_open_dev$cec]
ioctl$CEC_ADAP_S_PHYS_ADDR                  : fd_cec [syz_open_dev$cec]
ioctl$CEC_DQEVENT                           : fd_cec [syz_open_dev$cec]
ioctl$CEC_G_MODE                            : fd_cec [syz_open_dev$cec]
ioctl$CEC_RECEIVE                           : fd_cec [syz_open_dev$cec]
ioctl$CEC_S_MODE                            : fd_cec [syz_open_dev$cec]
ioctl$CEC_TRANSMIT                          : fd_cec [syz_open_dev$cec]
ioctl$CREATE_COUNTERS                       : fd_rdma [openat$uverbs0]
ioctl$DESTROY_COUNTERS                      : fd_rdma [openat$uverbs0]
ioctl$DRM_IOCTL_I915_GEM_BUSY               : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_GEM_CONTEXT_CREATE     : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_GEM_CONTEXT_DESTROY    : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_GEM_CONTEXT_GETPARAM   : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_GEM_CONTEXT_SETPARAM   : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_GEM_CREATE             : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_GEM_EXECBUFFER         : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_GEM_EXECBUFFER2        : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_GEM_EXECBUFFER2_WR     : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_GEM_GET_APERTURE       : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_GEM_GET_CACHING        : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_GEM_GET_TILING         : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_GEM_MADVISE            : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_GEM_MMAP               : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_GEM_MMAP_GTT           : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_GEM_MMAP_OFFSET        : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_GEM_PIN                : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_GEM_PREAD              : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_GEM_PWRITE             : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_GEM_SET_CACHING        : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_GEM_SET_DOMAIN         : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_GEM_SET_TILING         : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_GEM_SW_FINISH          : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_GEM_THROTTLE           : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_GEM_UNPIN              : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_GEM_USERPTR            : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_GEM_VM_CREATE          : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_GEM_VM_DESTROY         : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_GEM_WAIT               : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_GETPARAM               : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_GET_PIPE_FROM_CRTC_ID  : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_GET_RESET_STATS        : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_OVERLAY_ATTRS          : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_OVERLAY_PUT_IMAGE      : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_PERF_ADD_CONFIG        : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_PERF_OPEN              : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_PERF_REMOVE_CONFIG     : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_QUERY                  : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_REG_READ               : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_I915_SET_SPRITE_COLORKEY    : fd_i915 [openat$i915]
ioctl$DRM_IOCTL_MSM_GEM_CPU_FINI            : fd_msm [openat$msm]
ioctl$DRM_IOCTL_MSM_GEM_CPU_PREP            : fd_msm [openat$msm]
ioctl$DRM_IOCTL_MSM_GEM_INFO                : fd_msm [openat$msm]
ioctl$DRM_IOCTL_MSM_GEM_MADVISE             : fd_msm [openat$msm]
ioctl$DRM_IOCTL_MSM_GEM_NEW                 : fd_msm [openat$msm]
ioctl$DRM_IOCTL_MSM_GEM_SUBMIT              : fd_msm [openat$msm]
ioctl$DRM_IOCTL_MSM_GET_PARAM               : fd_msm [openat$msm]
ioctl$DRM_IOCTL_MSM_SET_PARAM               : fd_msm [openat$msm]
ioctl$DRM_IOCTL_MSM_SUBMITQUEUE_CLOSE       : fd_msm [openat$msm]
ioctl$DRM_IOCTL_MSM_SUBMITQUEUE_NEW         : fd_msm [openat$msm]
ioctl$DRM_IOCTL_MSM_SUBMITQUEUE_QUERY       : fd_msm [openat$msm]
ioctl$DRM_IOCTL_MSM_WAIT_FENCE              : fd_msm [openat$msm]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_CACHE_CACHEOPEXEC: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_CACHE_CACHEOPLOG: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_CACHE_CACHEOPQUEUE: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_CMM_DEVMEMINTACQUIREREMOTECTX: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_CMM_DEVMEMINTEXPORTCTX: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_CMM_DEVMEMINTUNEXPORTCTX: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_DEVICEMEMHISTORY_DEVICEMEMHISTORYMAP: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_DEVICEMEMHISTORY_DEVICEMEMHISTORYMAPVRANGE: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_DEVICEMEMHISTORY_DEVICEMEMHISTORYSPARSECHANGE: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_DEVICEMEMHISTORY_DEVICEMEMHISTORYUNMAP: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_DEVICEMEMHISTORY_DEVICEMEMHISTORYUNMAPVRANGE: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_DMABUF_PHYSMEMEXPORTDMABUF: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_DMABUF_PHYSMEMIMPORTDMABUF: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_DMABUF_PHYSMEMIMPORTSPARSEDMABUF: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_HTBUFFER_HTBCONTROL: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_HTBUFFER_HTBLOG: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_CHANGESPARSEMEM: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_DEVMEMFLUSHDEVSLCRANGE: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_DEVMEMGETFAULTADDRESS: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_DEVMEMINTCTXCREATE: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_DEVMEMINTCTXDESTROY: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_DEVMEMINTHEAPCREATE: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_DEVMEMINTHEAPDESTROY: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_DEVMEMINTMAPPAGES: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_DEVMEMINTMAPPMR: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_DEVMEMINTPIN: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_DEVMEMINTPINVALIDATE: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_DEVMEMINTREGISTERPFNOTIFYKM: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_DEVMEMINTRESERVERANGE: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_DEVMEMINTUNMAPPAGES: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_DEVMEMINTUNMAPPMR: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_DEVMEMINTUNPIN: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_DEVMEMINTUNPININVALIDATE: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_DEVMEMINTUNRESERVERANGE: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_DEVMEMINVALIDATEFBSCTABLE: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_DEVMEMISVDEVADDRVALID: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_GETMAXDEVMEMSIZE: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_HEAPCFGHEAPCONFIGCOUNT: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_HEAPCFGHEAPCONFIGNAME: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_HEAPCFGHEAPCOUNT: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_HEAPCFGHEAPDETAILS: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_PHYSMEMNEWRAMBACKEDLOCKEDPMR: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_PHYSMEMNEWRAMBACKEDPMR: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_PMREXPORTPMR: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_PMRGETUID: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_PMRIMPORTPMR: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_PMRLOCALIMPORTPMR: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_PMRMAKELOCALIMPORTHANDLE: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_PMRUNEXPORTPMR: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_PMRUNMAKELOCALIMPORTHANDLE: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_PMRUNREFPMR: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_PMRUNREFUNLOCKPMR: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_MM_PVRSRVUPDATEOOMSTATS: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_PVRTL_TLACQUIREDATA: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_PVRTL_TLCLOSESTREAM: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_PVRTL_TLCOMMITSTREAM: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_PVRTL_TLDISCOVERSTREAMS: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_PVRTL_TLOPENSTREAM: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_PVRTL_TLRELEASEDATA: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_PVRTL_TLRESERVESTREAM: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_PVRTL_TLWRITEDATA: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXBREAKPOINT_RGXCLEARBREAKPOINT: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXBREAKPOINT_RGXDISABLEBREAKPOINT: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXBREAKPOINT_RGXENABLEBREAKPOINT: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXBREAKPOINT_RGXOVERALLOCATEBPREGISTERS: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXBREAKPOINT_RGXSETBREAKPOINT: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXCMP_RGXCREATECOMPUTECONTEXT: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXCMP_RGXDESTROYCOMPUTECONTEXT: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXCMP_RGXFLUSHCOMPUTEDATA: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXCMP_RGXGETLASTCOMPUTECONTEXTRESETREASON: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXCMP_RGXKICKCDM2: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXCMP_RGXNOTIFYCOMPUTEWRITEOFFSETUPDATE: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXCMP_RGXSETCOMPUTECONTEXTPRIORITY: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXCMP_RGXSETCOMPUTECONTEXTPROPERTY: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXFWDBG_RGXCURRENTTIME: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXFWDBG_RGXFWDEBUGDUMPFREELISTPAGELIST: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXFWDBG_RGXFWDEBUGPHRCONFIGURE: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXFWDBG_RGXFWDEBUGSETFWLOG: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXFWDBG_RGXFWDEBUGSETHCSDEADLINE: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXFWDBG_RGXFWDEBUGSETOSIDPRIORITY: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXFWDBG_RGXFWDEBUGSETOSNEWONLINESTATE: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXHWPERF_RGXCONFIGCUSTOMCOUNTERS: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXHWPERF_RGXCONFIGENABLEHWPERFCOUNTERS: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXHWPERF_RGXCTRLHWPERF: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXHWPERF_RGXCTRLHWPERFCOUNTERS: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXHWPERF_RGXGETHWPERFBVNCFEATUREFLAGS: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXKICKSYNC_RGXCREATEKICKSYNCCONTEXT: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXKICKSYNC_RGXDESTROYKICKSYNCCONTEXT: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXKICKSYNC_RGXKICKSYNC2: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXKICKSYNC_RGXSETKICKSYNCCONTEXTPROPERTY: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXREGCONFIG_RGXADDREGCONFIG: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXREGCONFIG_RGXCLEARREGCONFIG: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXREGCONFIG_RGXDISABLEREGCONFIG: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXREGCONFIG_RGXENABLEREGCONFIG: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXREGCONFIG_RGXSETREGCONFIGTYPE: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXSIGNALS_RGXNOTIFYSIGNALUPDATE: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXTA3D_RGXCREATEFREELIST: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXTA3D_RGXCREATEHWRTDATASET: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXTA3D_RGXCREATERENDERCONTEXT: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXTA3D_RGXCREATEZSBUFFER: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXTA3D_RGXDESTROYFREELIST: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXTA3D_RGXDESTROYHWRTDATASET: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXTA3D_RGXDESTROYRENDERCONTEXT: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXTA3D_RGXDESTROYZSBUFFER: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXTA3D_RGXGETLASTRENDERCONTEXTRESETREASON: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXTA3D_RGXKICKTA3D2: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXTA3D_RGXPOPULATEZSBUFFER: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXTA3D_RGXRENDERCONTEXTSTALLED: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXTA3D_RGXSETRENDERCONTEXTPRIORITY: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXTA3D_RGXSETRENDERCONTEXTPROPERTY: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXTA3D_RGXUNPOPULATEZSBUFFER: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXTQ2_RGXTDMCREATETRANSFERCONTEXT: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXTQ2_RGXTDMDESTROYTRANSFERCONTEXT: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXTQ2_RGXTDMGETSHAREDMEMORY: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXTQ2_RGXTDMNOTIFYWRITEOFFSETUPDATE: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXTQ2_RGXTDMRELEASESHAREDMEMORY: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXTQ2_RGXTDMSETTRANSFERCONTEXTPRIORITY: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXTQ2_RGXTDMSETTRANSFERCONTEXTPROPERTY: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXTQ2_RGXTDMSUBMITTRANSFER2: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXTQ_RGXCREATETRANSFERCONTEXT: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXTQ_RGXDESTROYTRANSFERCONTEXT: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXTQ_RGXSETTRANSFERCONTEXTPRIORITY: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXTQ_RGXSETTRANSFERCONTEXTPROPERTY: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_RGXTQ_RGXSUBMITTRANSFER2: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_SRVCORE_ACQUIREGLOBALEVENTOBJECT: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_SRVCORE_ACQUIREINFOPAGE: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_SRVCORE_ALIGNMENTCHECK: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_SRVCORE_CONNECT: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_SRVCORE_DISCONNECT: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_SRVCORE_DUMPDEBUGINFO: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_SRVCORE_EVENTOBJECTCLOSE: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_SRVCORE_EVENTOBJECTOPEN: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_SRVCORE_EVENTOBJECTWAIT: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_SRVCORE_EVENTOBJECTWAITTIMEOUT: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_SRVCORE_FINDPROCESSMEMSTATS: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_SRVCORE_GETDEVCLOCKSPEED: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_SRVCORE_GETDEVICESTATUS: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_SRVCORE_GETMULTICOREINFO: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_SRVCORE_HWOPTIMEOUT: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_SRVCORE_RELEASEGLOBALEVENTOBJECT: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_SRVCORE_RELEASEINFOPAGE: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_SYNCTRACKING_SYNCRECORDADD: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_SYNCTRACKING_SYNCRECORDREMOVEBYHANDLE: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_SYNC_ALLOCSYNCPRIMITIVEBLOCK: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_SYNC_FREESYNCPRIMITIVEBLOCK: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_SYNC_SYNCALLOCEVENT: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_SYNC_SYNCCHECKPOINTSIGNALLEDPDUMPPOL: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_SYNC_SYNCFREEEVENT: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_SYNC_SYNCPRIMPDUMP: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_SYNC_SYNCPRIMPDUMPCBP: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_SYNC_SYNCPRIMPDUMPPOL: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_SYNC_SYNCPRIMPDUMPVALUE: fd_rogue [openat$img_rogue]
ioctl$DRM_IOCTL_PVR_SRVKM_CMD_PVRSRV_BRIDGE_SYNC_SYNCPRIMSET: fd_rogue [openat$img_rogue]
ioctl$DVD_AUTH                              : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$DVD_READ_STRUCT                       : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$DVD_WRITE_STRUCT                      : fd_cdrom [openat$cdrom openat$cdrom1]
ioctl$FBIOBLANK                             : fd_fb [openat$fb0 openat$fb1]
ioctl$FBIOGETCMAP                           : fd_fb [openat$fb0 openat$fb1]
ioctl$FBIOGET_CON2FBMAP                     : fd_fb [openat$fb0 openat$fb1]
ioctl$FBIOGET_FSCREENINFO                   : fd_fb [openat$fb0 openat$fb1]
ioctl$FBIOGET_VSCREENINFO                   : fd_fb [openat$fb0 openat$fb1]
ioctl$FBIOPAN_DISPLAY                       : fd_fb [openat$fb0 openat$fb1]
ioctl$FBIOPUTCMAP                           : fd_fb [openat$fb0 openat$fb1]
ioctl$FBIOPUT_CON2FBMAP                     : fd_fb [openat$fb0 openat$fb1]
ioctl$FBIOPUT_VSCREENINFO                   : fd_fb [openat$fb0 openat$fb1]
ioctl$FBIO_WAITFORVSYNC                     : fd_fb [openat$fb0 openat$fb1]
ioctl$FLOPPY_FDCLRPRM                       : fd_floppy [syz_open_dev$floppy]
ioctl$FLOPPY_FDDEFPRM                       : fd_floppy [syz_open_dev$floppy]
ioctl$FLOPPY_FDEJECT                        : fd_floppy [syz_open_dev$floppy]
ioctl$FLOPPY_FDFLUSH                        : fd_floppy [syz_open_dev$floppy]
ioctl$FLOPPY_FDFMTBEG                       : fd_floppy [syz_open_dev$floppy]
ioctl$FLOPPY_FDFMTEND                       : fd_floppy [syz_open_dev$floppy]
ioctl$FLOPPY_FDFMTTRK                       : fd_floppy [syz_open_dev$floppy]
ioctl$FLOPPY_FDGETDRVPRM                    : fd_floppy [syz_open_dev$floppy]
ioctl$FLOPPY_FDGETDRVSTAT                   : fd_floppy [syz_open_dev$floppy]
ioctl$FLOPPY_FDGETDRVTYP                    : fd_floppy [syz_open_dev$floppy]
ioctl$FLOPPY_FDGETFDCSTAT                   : fd_floppy [syz_open_dev$floppy]
ioctl$FLOPPY_FDGETMAXERRS                   : fd_floppy [syz_open_dev$floppy]
ioctl$FLOPPY_FDGETPRM                       : fd_floppy [syz_open_dev$floppy]
ioctl$FLOPPY_FDMSGOFF                       : fd_floppy [syz_open_dev$floppy]
ioctl$FLOPPY_FDMSGON                        : fd_floppy [syz_open_dev$floppy]
ioctl$FLOPPY_FDPOLLDRVSTAT                  : fd_floppy [syz_open_dev$floppy]
ioctl$FLOPPY_FDRAWCMD                       : fd_floppy [syz_open_dev$floppy]
ioctl$FLOPPY_FDRESET                        : fd_floppy [syz_open_dev$floppy]
ioctl$FLOPPY_FDSETDRVPRM                    : fd_floppy [syz_open_dev$floppy]
ioctl$FLOPPY_FDSETEMSGTRESH                 : fd_floppy [syz_open_dev$floppy]
ioctl$FLOPPY_FDSETMAXERRS                   : fd_floppy [syz_open_dev$floppy]
ioctl$FLOPPY_FDSETPRM                       : fd_floppy [syz_open_dev$floppy]
ioctl$FLOPPY_FDTWADDLE                      : fd_floppy [syz_open_dev$floppy]
ioctl$FLOPPY_FDWERRORCLR                    : fd_floppy [syz_open_dev$floppy]
ioctl$FLOPPY_FDWERRORGET                    : fd_floppy [syz_open_dev$floppy]
ioctl$HDIO_GETGEO                           : fd_block [openat$md openat$nullb openat$pmem0 syz_open_dev$loop syz_open_dev$ndb]
ioctl$HIDIOCAPPLICATION                     : fd_hiddev [syz_open_dev$hiddev]
ioctl$HIDIOCGCOLLECTIONINDEX                : fd_hiddev [syz_open_dev$hiddev]
ioctl$HIDIOCGCOLLECTIONINFO                 : fd_hiddev [syz_open_dev$hiddev]
ioctl$HIDIOCGDEVINFO                        : fd_hiddev [syz_open_dev$hiddev]
ioctl$HIDIOCGFEATURE                        : fd_hidraw [syz_open_dev$hidraw]
ioctl$HIDIOCGFIELDINFO                      : fd_hiddev [syz_open_dev$hiddev]
ioctl$HIDIOCGFLAG                           : fd_hiddev [syz_open_dev$hiddev]
ioctl$HIDIOCGNAME                           : fd_hiddev [syz_open_dev$hiddev]
ioctl$HIDIOCGPHYS                           : fd_hiddev [syz_open_dev$hiddev]
ioctl$HIDIOCGRAWINFO                        : fd_hidraw [syz_open_dev$hidraw]
ioctl$HIDIOCGRAWNAME                        : fd_hidraw [syz_open_dev$hidraw]
ioctl$HIDIOCGRAWPHYS                        : fd_hidraw [syz_open_dev$hidraw]
ioctl$HIDIOCGRDESC                          : fd_hidraw [syz_open_dev$hidraw]
ioctl$HIDIOCGRDESCSIZE                      : fd_hidraw [syz_open_dev$hidraw]
ioctl$HIDIOCGREPORT                         : fd_hiddev [syz_open_dev$hiddev]
ioctl$HIDIOCGREPORTINFO                     : fd_hiddev [syz_open_dev$hiddev]
ioctl$HIDIOCGSTRING                         : fd_hiddev [syz_open_dev$hiddev]
ioctl$HIDIOCGUCODE                          : fd_hiddev [syz_open_dev$hiddev]
ioctl$HIDIOCGUSAGE                          : fd_hiddev [syz_open_dev$hiddev]
ioctl$HIDIOCGUSAGES                         : fd_hiddev [syz_open_dev$hiddev]
ioctl$HIDIOCGVERSION                        : fd_hiddev [syz_open_dev$hiddev]
ioctl$HIDIOCINITREPORT                      : fd_hiddev [syz_open_dev$hiddev]
ioctl$HIDIOCSFEATURE                        : fd_hidraw [syz_open_dev$hidraw]
ioctl$HIDIOCSFLAG                           : fd_hiddev [syz_open_dev$hiddev]
ioctl$HIDIOCSREPORT                         : fd_hiddev [syz_open_dev$hiddev]
ioctl$HIDIOCSUSAGE                          : fd_hiddev [syz_open_dev$hiddev]
ioctl$HIDIOCSUSAGES                         : fd_hiddev [syz_open_dev$hiddev]
ioctl$IMADDTIMER                            : fd_misdntimer [openat$misdntimer]
ioctl$IMCLEAR_L2                            : sock_isdn [socket$isdn]
ioctl$IMCTRLREQ                             : sock_isdn [socket$isdn]
ioctl$IMDELTIMER                            : fd_misdntimer [openat$misdntimer]
ioctl$IMGETCOUNT                            : sock_isdn_base [socket$isdn_base]
ioctl$IMGETDEVINFO                          : sock_isdn_base [socket$isdn_base]
ioctl$IMGETVERSION                          : sock_isdn_base [socket$isdn_base]
ioctl$IMHOLD_L1                             : sock_isdn [socket$isdn]
ioctl$IMSETDEVNAME                          : sock_isdn_base [socket$isdn_base]
ioctl$IOCTL_CONFIG_SYS_RESOURCE_PARAMETERS  : fd_qat [openat$qat_adf_ctl]
ioctl$IOCTL_GET_NCIDEV_IDX                  : fd_nci [openat$nci]
ioctl$IOCTL_GET_NUM_DEVICES                 : fd_qat [openat$qat_adf_ctl]
ioctl$IOCTL_START_ACCEL_DEV                 : fd_qat [openat$qat_adf_ctl]
ioctl$IOCTL_STATUS_ACCEL_DEV                : fd_qat [openat$qat_adf_ctl]
ioctl$IOCTL_STOP_ACCEL_DEV                  : fd_qat [openat$qat_adf_ctl]
ioctl$IOCTL_VMCI_CTX_ADD_NOTIFICATION       : fd_vmci [openat$vmci]
ioctl$IOCTL_VMCI_CTX_GET_CPT_STATE          : fd_vmci [openat$vmci]
ioctl$IOCTL_VMCI_CTX_REMOVE_NOTIFICATION    : fd_vmci [openat$vmci]
ioctl$IOCTL_VMCI_CTX_SET_CPT_STATE          : fd_vmci [openat$vmci]
ioctl$IOCTL_VMCI_DATAGRAM_RECEIVE           : fd_vmci [openat$vmci]
ioctl$IOCTL_VMCI_DATAGRAM_SEND              : fd_vmci [openat$vmci]
ioctl$IOCTL_VMCI_GET_CONTEXT_ID             : fd_vmci [openat$vmci]
ioctl$IOCTL_VMCI_INIT_CONTEXT               : fd_vmci [openat$vmci]
ioctl$IOCTL_VMCI_NOTIFICATIONS_RECEIVE      : fd_vmci [openat$vmci]
ioctl$IOCTL_VMCI_NOTIFY_RESOURCE            : fd_vmci [openat$vmci]
ioctl$IOCTL_VMCI_QUEUEPAIR_ALLOC            : fd_vmci [openat$vmci]
ioctl$IOCTL_VMCI_QUEUEPAIR_DETACH           : fd_vmci [openat$vmci]
ioctl$IOCTL_VMCI_QUEUEPAIR_SETPF            : fd_vmci [openat$vmci]
ioctl$IOCTL_VMCI_QUEUEPAIR_SETVA            : fd_vmci [openat$vmci]
ioctl$IOCTL_VMCI_SET_NOTIFY                 : fd_vmci [openat$vmci]
ioctl$IOCTL_VMCI_VERSION                    : fd_vmci [openat$vmci]
ioctl$IOCTL_VMCI_VERSION2                   : fd_vmci [openat$vmci]
ioctl$IOC_PR_CLEAR                          : fd_block [openat$md openat$nullb openat$pmem0 syz_open_dev$loop syz_open_dev$ndb]
ioctl$IOC_PR_PREEMPT                        : fd_block [openat$md openat$nullb openat$pmem0 syz_open_dev$loop syz_open_dev$ndb]
ioctl$IOC_PR_PREEMPT_ABORT                  : fd_block [openat$md openat$nullb openat$pmem0 syz_open_dev$loop syz_open_dev$ndb]
ioctl$IOC_PR_REGISTER                       : fd_block [openat$md openat$nullb openat$pmem0 syz_open_dev$loop syz_open_dev$ndb]
ioctl$IOC_PR_RELEASE                        : fd_block [openat$md openat$nullb openat$pmem0 syz_open_dev$loop syz_open_dev$ndb]
ioctl$IOC_PR_RESERVE                        : fd_block [openat$md openat$nullb openat$pmem0 syz_open_dev$loop syz_open_dev$ndb]
ioctl$IOMMU_DESTROY$hwpt                    : fd_iommufd [openat$iommufd]
ioctl$IOMMU_DESTROY$ioas                    : fd_iommufd [openat$iommufd]
ioctl$IOMMU_DESTROY$stdev                   : fd_iommufd [openat$iommufd]
ioctl$IOMMU_GET_HW_INFO                     : fd_iommufd [openat$iommufd]
ioctl$IOMMU_HWPT_ALLOC$NONE                 : fd_iommufd [openat$iommufd]
ioctl$IOMMU_HWPT_ALLOC$TEST                 : fd_iommufd [openat$iommufd]
ioctl$IOMMU_HWPT_GET_DIRTY_BITMAP           : fd_iommufd [openat$iommufd]
ioctl$IOMMU_HWPT_INVALIDATE$TEST            : fd_iommufd [openat$iommufd]
ioctl$IOMMU_HWPT_SET_DIRTY_TRACKING         : fd_iommufd [openat$iommufd]
ioctl$IOMMU_IOAS_ALLOC                      : fd_iommufd [openat$iommufd]
ioctl$IOMMU_IOAS_ALLOW_IOVAS                : fd_iommufd [openat$iommufd]
ioctl$IOMMU_IOAS_COPY                       : fd_iommufd [openat$iommufd]
ioctl$IOMMU_IOAS_COPY$syz                   : fd_iommufd [openat$iommufd]
ioctl$IOMMU_IOAS_IOVA_RANGES                : fd_iommufd [openat$iommufd]
ioctl$IOMMU_IOAS_MAP                        : fd_iommufd [openat$iommufd]
ioctl$IOMMU_IOAS_MAP$PAGES                  : fd_iommufd [openat$iommufd]
ioctl$IOMMU_IOAS_UNMAP                      : fd_iommufd [openat$iommufd]
ioctl$IOMMU_IOAS_UNMAP$ALL                  : fd_iommufd [openat$iommufd]
ioctl$IOMMU_OPTION$IOMMU_OPTION_HUGE_PAGES  : fd_iommufd [openat$iommufd]
ioctl$IOMMU_OPTION$IOMMU_OPTION_RLIMIT_MODE : fd_iommufd [openat$iommufd]
ioctl$IOMMU_TEST_OP_ACCESS_PAGES            : fd_iommufd [openat$iommufd]
ioctl$IOMMU_TEST_OP_ACCESS_PAGES$syz        : fd_iommufd [openat$iommufd]
ioctl$IOMMU_TEST_OP_ACCESS_REPLACE_IOAS     : fd_iommufd [openat$iommufd]
ioctl$IOMMU_TEST_OP_ACCESS_RW               : fd_iommufd [openat$iommufd]
ioctl$IOMMU_TEST_OP_ACCESS_RW$syz           : fd_iommufd [openat$iommufd]
ioctl$IOMMU_TEST_OP_ADD_RESERVED            : fd_iommufd [openat$iommufd]
ioctl$IOMMU_TEST_OP_CREATE_ACCESS           : fd_iommufd [openat$iommufd]
ioctl$IOMMU_TEST_OP_DESTROY_ACCESS_PAGES    : fd_iommufd [openat$iommufd]
ioctl$IOMMU_TEST_OP_MD_CHECK_MAP            : fd_iommufd [openat$iommufd]
ioctl$IOMMU_TEST_OP_MD_CHECK_REFS           : fd_iommufd [openat$iommufd]
ioctl$IOMMU_TEST_OP_MOCK_DOMAIN             : fd_iommufd [openat$iommufd]
ioctl$IOMMU_TEST_OP_MOCK_DOMAIN_FLAGS       : fd_iommufd [openat$iommufd]
ioctl$IOMMU_TEST_OP_MOCK_DOMAIN_REPLACE     : fd_iommufd [openat$iommufd]
ioctl$IOMMU_TEST_OP_SET_TEMP_MEMORY_LIMIT   : fd_iommufd [openat$iommufd]
ioctl$IOMMU_VFIO_CHECK_EXTENSION            : fd_iommufd [openat$iommufd]
ioctl$IOMMU_VFIO_GET_API_VERSION            : fd_iommufd [openat$iommufd]
ioctl$IOMMU_VFIO_IOAS$CLEAR                 : fd_iommufd [openat$iommufd]
ioctl$IOMMU_VFIO_IOAS$GET                   : fd_iommufd [openat$iommufd]
ioctl$IOMMU_VFIO_IOAS$SET                   : fd_iommufd [openat$iommufd]
ioctl$IOMMU_VFIO_IOMMU_GET_INFO             : fd_iommufd [openat$iommufd]
ioctl$IOMMU_VFIO_IOMMU_MAP_DMA              : fd_iommufd [openat$iommufd]
ioctl$IOMMU_VFIO_IOMMU_UNMAP_DMA            : fd_iommufd [openat$iommufd]
ioctl$IOMMU_VFIO_SET_IOMMU                  : fd_iommufd [openat$iommufd]
ioctl$KBASE_HWCNT_READER_CLEAR              : fd_hwcnt [ioctl$KBASE_IOCTL_HWCNT_READER_SETUP]
ioctl$KBASE_HWCNT_READER_DISABLE_EVENT      : fd_hwcnt [ioctl$KBASE_IOCTL_HWCNT_READER_SETUP]
ioctl$KBASE_HWCNT_READER_DUMP               : fd_hwcnt [ioctl$KBASE_IOCTL_HWCNT_READER_SETUP]
ioctl$KBASE_HWCNT_READER_ENABLE_EVENT       : fd_hwcnt [ioctl$KBASE_IOCTL_HWCNT_READER_SETUP]
ioctl$KBASE_HWCNT_READER_GET_API_VERSION    : fd_hwcnt [ioctl$KBASE_IOCTL_HWCNT_READER_SETUP]
ioctl$KBASE_HWCNT_READER_GET_API_VERSION_WITH_FEATURES: fd_hwcnt [ioctl$KBASE_IOCTL_HWCNT_READER_SETUP]
ioctl$KBASE_HWCNT_READER_GET_BUFFER         : fd_hwcnt [ioctl$KBASE_IOCTL_HWCNT_READER_SETUP]
ioctl$KBASE_HWCNT_READER_GET_BUFFER_SIZE    : fd_hwcnt [ioctl$KBASE_IOCTL_HWCNT_READER_SETUP]
ioctl$KBASE_HWCNT_READER_GET_BUFFER_WITH_CYCLES: fd_hwcnt [ioctl$KBASE_IOCTL_HWCNT_READER_SETUP]
ioctl$KBASE_HWCNT_READER_GET_HWVER          : fd_hwcnt [ioctl$KBASE_IOCTL_HWCNT_READER_SETUP]
ioctl$KBASE_HWCNT_READER_PUT_BUFFER         : fd_hwcnt [ioctl$KBASE_IOCTL_HWCNT_READER_SETUP]
ioctl$KBASE_HWCNT_READER_PUT_BUFFER_WITH_CYCLES: fd_hwcnt [ioctl$KBASE_IOCTL_HWCNT_READER_SETUP]
ioctl$KBASE_HWCNT_READER_SET_INTERVAL       : fd_hwcnt [ioctl$KBASE_IOCTL_HWCNT_READER_SETUP]
ioctl$KBASE_IOCTL_BUFFER_LIVENESS_UPDATE    : fd_bifrost [openat$bifrost openat$mali]
ioctl$KBASE_IOCTL_CONTEXT_PRIORITY_CHECK    : fd_bifrost [openat$bifrost openat$mali]
ioctl$KBASE_IOCTL_CS_CPU_QUEUE_DUMP         : fd_bifrost [openat$bifrost openat$mali]
ioctl$KBASE_IOCTL_CS_EVENT_SIGNAL           : fd_bifrost [openat$bifrost openat$mali]
ioctl$KBASE_IOCTL_CS_GET_GLB_IFACE          : fd_bifrost [openat$bifrost openat$mali]
ioctl$KBASE_IOCTL_CS_QUEUE_BIND             : fd_bifrost [openat$bifrost openat$mali]
ioctl$KBASE_IOCTL_CS_QUEUE_GROUP_CREATE     : fd_bifrost [openat$bifrost openat$mali]
ioctl$KBASE_IOCTL_CS_QUEUE_GROUP_CREATE_1_6 : fd_bifrost [openat$bifrost openat$mali]
ioctl$KBASE_IOCTL_CS_QUEUE_GROUP_TERMINATE  : fd_bifrost [openat$bifrost openat$mali]
ioctl$KBASE_IOCTL_CS_QUEUE_KICK             : fd_bifrost [openat$bifrost openat$mali]
ioctl$KBASE_IOCTL_CS_QUEUE_REGISTER         : fd_bifrost [openat$bifrost openat$mali]
ioctl$KBASE_IOCTL_CS_QUEUE_REGISTER_EX      : fd_bifrost [openat$bifrost openat$mali]
ioctl$KBASE_IOCTL_CS_QUEUE_TERMINATE        : fd_bifrost [openat$bifrost openat$mali]
ioctl$KBASE_IOCTL_CS_TILER_HEAP_INIT        : fd_bifrost [openat$bifrost openat$mali]
ioctl$KBASE_IOCTL_CS_TILER_HEAP_INIT_1_13   : fd_bifrost [openat$bifrost openat$mali]
ioctl$KBASE_IOCTL_CS_TILER_HEAP_TERM        : fd_bifrost [openat$bifrost openat$mali]
ioctl$KBASE_IOCTL_DISJOINT_QUERY            : fd_bifrost [openat$bifrost openat$mali]
ioctl$KBASE_IOCTL_FENCE_VALIDATE            : fd_bifrost [openat$bifrost openat$mali]
ioctl$KBASE_IOCTL_GET_CONTEXT_ID            : fd_bifrost [openat$bifrost openat$mali]
ioctl$KBASE_IOCTL_GET_CPU_GPU_TIMEINFO      : fd_bifrost [openat$bifrost openat$mali]
ioctl$KBASE_IOCTL_GET_DDK_VERSION           : fd_bifrost [openat$bifrost openat$mali]
ioctl$KBASE_IOCTL_GET_GPUPROPS              : fd_bifrost [openat$bifrost openat$mali]
ioctl$KBASE_IOCTL_HWCNT_CLEAR               : fd_bifrost [openat$bifrost openat$mali]
ioctl$KBASE_IOCTL_HWCNT_DUMP                : fd_bifrost [openat$bifrost openat$mali]
ioctl$KBASE_IOCTL_HWCNT_ENABLE              : fd_bifrost [openat$bifrost openat$mali]
ioctl$KBASE_IOCTL_HWCNT_READER_SETUP        : fd_bifrost [openat$bifrost openat$mali]
ioctl$KBASE_IOCTL_HWCNT_SET                 : fd_bifrost [openat$bifrost openat$mali]
ioctl$KBASE_IOCTL_JOB_SUBMIT                : fd_bifrost [openat$bifrost openat$mali]
ioctl$KBASE_IOCTL_KCPU_QUEUE_CREATE         : fd_bifrost [openat$bifrost openat$mali]
ioctl$KBASE_IOCTL_KCPU_QUEUE_DELETE         : fd_bifrost [openat$bifrost openat$mali]
ioctl$KBASE_IOCTL_KCPU_QUEUE_ENQUEUE        : fd_bifrost [openat$bifrost openat$mali]
ioctl$KBASE_IOCTL_KINSTR_PRFCNT_ENUM_INFO   : fd_bifrost [openat$bifrost openat$mali]
ioctl$KBASE_IOCTL_KINSTR_PRFCNT_SETUP       : fd_bifrost [openat$bifrost openat$mali]
ioctl$KBASE_IOCTL_MEM_ALIAS                 : fd_bifrost [openat$bifrost openat$mali]
ioctl$KBASE_IOCTL_MEM_ALLOC                 : fd_bifrost [openat$bifrost openat$mali]
ioctl$KBASE_IOCTL_MEM_ALLOC_EX              : fd_bifrost [openat$bifrost openat$mali]
ioctl$KBASE_IOCTL_MEM_COMMIT                : fd_bifrost [openat$bifrost openat$mali]
ioctl$KBASE_IOCTL_MEM_EXEC_INIT             : fd_bifrost [openat$bifrost openat$mali]
ioctl$KBASE_IOCTL_MEM_FIND_CPU_OFFSET       : fd_bifrost [openat$bifrost openat$mali]
ioctl$KBASE_IOCTL_MEM_FIND_GPU_START_AND_OFFSET: fd_bifrost [openat$bifrost openat$mali]
ioctl$KBASE_IOCTL_MEM_FLAGS_CHANGE          : fd_bifrost [openat$bifrost openat$mali]
ioctl$KBASE_IOCTL_MEM_FREE                  : fd_bifrost [openat$bifrost openat$mali]
ioctl$KBASE_IOCTL_MEM_IMPORT                : fd_bifrost [openat$bifrost openat$mali]
ioctl$KBASE_IOCTL_MEM_JIT_INIT              : fd_bifrost [openat$bifrost openat$mali]
ioctl$KBASE_IOCTL_MEM_JIT_INIT_10_2         : fd_bifrost [openat$bifrost openat$mali]
ioctl$KBASE_IOCTL_MEM_JIT_INIT_11_5         : fd_bifrost [openat$bifrost openat$mali]
ioctl$KBASE_IOCTL_MEM_PROFILE_ADD           : fd_bifrost [openat$bifrost openat$mali]
ioctl$KBASE_IOCTL_MEM_QUERY                 : fd_bifrost [openat$bifrost openat$mali]
ioctl$KBASE_IOCTL_MEM_SYNC                  : fd_bifrost [openat$bifrost openat$mali]
ioctl$KBASE_IOCTL_POST_TERM                 : fd_bifrost [openat$bifrost openat$mali]
ioctl$KBASE_IOCTL_READ_USER_PAGE            : fd_bifrost [openat$bifrost openat$mali]
ioctl$KBASE_IOCTL_SET_FLAGS                 : fd_bifrost [openat$bifrost openat$mali]
ioctl$KBASE_IOCTL_SET_LIMITED_CORE_COUNT    : fd_bifrost [openat$bifrost openat$mali]
ioctl$KBASE_IOCTL_SOFT_EVENT_UPDATE         : fd_bifrost [openat$bifrost openat$mali]
ioctl$KBASE_IOCTL_STICKY_RESOURCE_MAP       : fd_bifrost [openat$bifrost openat$mali]
ioctl$KBASE_IOCTL_STICKY_RESOURCE_UNMAP     : fd_bifrost [openat$bifrost openat$mali]
ioctl$KBASE_IOCTL_STREAM_CREATE             : fd_bifrost [openat$bifrost openat$mali]
ioctl$KBASE_IOCTL_TLSTREAM_ACQUIRE          : fd_bifrost [openat$bifrost openat$mali]
ioctl$KBASE_IOCTL_TLSTREAM_FLUSH            : fd_bifrost [openat$bifrost openat$mali]
ioctl$KBASE_IOCTL_VERSION_CHECK             : fd_bifrost [openat$bifrost openat$mali]
ioctl$KBASE_IOCTL_VERSION_CHECK_RESERVED    : fd_bifrost [openat$bifrost openat$mali]
ioctl$KVM_ARM_PREFERRED_TARGET              : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_ARM_SET_COUNTER_OFFSET            : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_ARM_SET_DEVICE_ADDR               : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_ARM_VCPU_FINALIZE                 : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_ARM_VCPU_INIT                     : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_ASSIGN_SET_MSIX_ENTRY             : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_ASSIGN_SET_MSIX_NR                : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_ARM_EAGER_SPLIT_CHUNK_SIZE    : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_ARM_INJECT_SERROR_ESR         : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_ARM_MTE                       : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_ARM_SYSTEM_SUSPEND            : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_ARM_USER_IRQ                  : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_DIRTY_LOG_RING                : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_DIRTY_LOG_RING_ACQ_REL        : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_DISABLE_QUIRKS                : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_DISABLE_QUIRKS2               : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_ENFORCE_PV_FEATURE_CPUID      : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_CAP_EXCEPTION_PAYLOAD             : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_EXIT_HYPERCALL                : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_EXIT_ON_EMULATION_FAILURE     : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_HALT_POLL                     : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_HYPERV_DIRECT_TLBFLUSH        : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_CAP_HYPERV_ENFORCE_CPUID          : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_CAP_HYPERV_ENLIGHTENED_VMCS       : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_CAP_HYPERV_SEND_IPI               : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_HYPERV_SYNIC                  : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_CAP_HYPERV_SYNIC2                 : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_CAP_HYPERV_TLBFLUSH               : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_HYPERV_VP_INDEX               : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_MANUAL_DIRTY_LOG_PROTECT2     : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_MAX_VCPU_ID                   : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_MEMORY_FAULT_INFO             : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_MSR_PLATFORM_INFO             : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_PMU_CAPABILITY                : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_PTP_KVM                       : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_SGX_ATTRIBUTE                 : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_SPLIT_IRQCHIP                 : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_STEAL_TIME                    : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_SYNC_REGS                     : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_CAP_VM_COPY_ENC_CONTEXT_FROM      : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_VM_DISABLE_NX_HUGE_PAGES      : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_VM_MOVE_ENC_CONTEXT_FROM      : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_VM_TYPES                      : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_X2APIC_API                    : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_X86_APIC_BUS_CYCLES_NS        : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_X86_BUS_LOCK_EXIT             : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_X86_DISABLE_EXITS             : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_X86_GUEST_MODE                : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_X86_NOTIFY_VMEXIT             : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_X86_USER_SPACE_MSR            : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CAP_XEN_HVM                       : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CHECK_EXTENSION                   : fd_kvm [openat$kvm]
ioctl$KVM_CHECK_EXTENSION_VM                : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CLEAR_DIRTY_LOG                   : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CREATE_DEVICE                     : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CREATE_GUEST_MEMFD                : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CREATE_IRQCHIP                    : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CREATE_PIT2                       : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CREATE_VCPU                       : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_CREATE_VM                         : fd_kvm [openat$kvm]
ioctl$KVM_DIRTY_TLB                         : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_GET_API_VERSION                   : fd_kvm [openat$kvm]
ioctl$KVM_GET_CLOCK                         : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_GET_DEVICE_ATTR                   : fd_kvmdev [ioctl$KVM_CREATE_DEVICE syz_kvm_vgic_v3_setup]
ioctl$KVM_GET_DEVICE_ATTR_vcpu              : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_GET_DEVICE_ATTR_vm                : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_GET_DIRTY_LOG                     : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_GET_FPU                           : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_GET_IRQCHIP                       : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_GET_MP_STATE                      : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_GET_NR_MMU_PAGES                  : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_GET_ONE_REG                       : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_GET_REGS                          : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_GET_REG_LIST                      : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_GET_SREGS                         : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_GET_TSC_KHZ                       : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_GET_VCPU_EVENTS                   : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_GET_VCPU_MMAP_SIZE                : fd_kvm [openat$kvm]
ioctl$KVM_HAS_DEVICE_ATTR                   : fd_kvmdev [ioctl$KVM_CREATE_DEVICE syz_kvm_vgic_v3_setup]
ioctl$KVM_HAS_DEVICE_ATTR_vcpu              : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_HAS_DEVICE_ATTR_vm                : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_INTERRUPT                         : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_IOEVENTFD                         : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_IRQFD                             : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_IRQ_LINE                          : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_IRQ_LINE_STATUS                   : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_KVMCLOCK_CTRL                     : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_NMI                               : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_PPC_ALLOCATE_HTAB                 : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_PRE_FAULT_MEMORY                  : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_REGISTER_COALESCED_MMIO           : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_REINJECT_CONTROL                  : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_RESET_DIRTY_RINGS                 : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_RUN                               : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_S390_VCPU_FAULT                   : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_SET_BOOT_CPU_ID                   : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_SET_CLOCK                         : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_SET_DEVICE_ATTR                   : fd_kvmdev [ioctl$KVM_CREATE_DEVICE syz_kvm_vgic_v3_setup]
ioctl$KVM_SET_DEVICE_ATTR_vcpu              : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_SET_DEVICE_ATTR_vm                : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_SET_FPU                           : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_SET_GSI_ROUTING                   : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_SET_GUEST_DEBUG                   : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_SET_IDENTITY_MAP_ADDR             : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_SET_IRQCHIP                       : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_SET_MEMORY_ATTRIBUTES             : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_SET_MP_STATE                      : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_SET_NR_MMU_PAGES                  : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_SET_ONE_REG                       : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_SET_REGS                          : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_SET_SIGNAL_MASK                   : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_SET_SREGS                         : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_SET_TSC_KHZ                       : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_SET_TSS_ADDR                      : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_SET_USER_MEMORY_REGION            : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_SET_USER_MEMORY_REGION2           : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_SET_VAPIC_ADDR                    : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_SET_VCPU_EVENTS                   : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_SIGNAL_MSI                        : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_SMI                               : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_TPR_ACCESS_REPORTING              : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_TRANSLATE                         : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$KVM_UNREGISTER_COALESCED_MMIO         : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_X86_GET_MCE_CAP_SUPPORTED         : fd_kvmvm [ioctl$KVM_CREATE_VM]
ioctl$KVM_X86_SETUP_MCE                     : fd_kvmcpu [ioctl$KVM_CREATE_VCPU syz_kvm_add_vcpu]
ioctl$LOOP_CHANGE_FD                        : fd_loop [syz_open_dev$loop]
ioctl$LOOP_CLR_FD                           : fd_loop [syz_open_dev$loop]
ioctl$LOOP_CONFIGURE                        : fd_loop [syz_open_dev$loop]
ioctl$LOOP_GET_STATUS                       : fd_loop [syz_open_dev$loop]
ioctl$LOOP_GET_STATUS64                     : fd_loop [syz_open_dev$loop]
ioctl$LOOP_SET_BLOCK_SIZE                   : fd_loop [syz_open_dev$loop]
ioctl$LOOP_SET_CAPACITY                     : fd_loop [syz_open_dev$loop]
ioctl$LOOP_SET_DIRECT_IO                    : fd_loop [syz_open_dev$loop]
ioctl$LOOP_SET_FD                           : fd_loop [syz_open_dev$loop]
ioctl$LOOP_SET_STATUS                       : fd_loop [syz_open_dev$loop]
ioctl$LOOP_SET_STATUS64                     : fd_loop [syz_open_dev$loop]
ioctl$MON_IOCG_STATS                        : fd_usbmon [syz_open_dev$usbmon]
ioctl$MON_IOCH_MFLUSH                       : fd_usbmon [syz_open_dev$usbmon]
ioctl$MON_IOCQ_RING_SIZE                    : fd_usbmon [syz_open_dev$usbmon]
ioctl$MON_IOCQ_URB_LEN                      : fd_usbmon [syz_open_dev$usbmon]
ioctl$MON_IOCT_RING_SIZE                    : fd_usbmon [syz_open_dev$usbmon]
ioctl$MON_IOCX_GET                          : fd_usbmon [syz_open_dev$usbmon]
ioctl$MON_IOCX_GETX                         : fd_usbmon [syz_open_dev$usbmon]
ioctl$MON_IOCX_MFETCH                       : fd_usbmon [syz_open_dev$usbmon]
ioctl$NBD_CLEAR_QUE                         : fd_nbd [syz_open_dev$ndb]
ioctl$NBD_CLEAR_SOCK                        : fd_nbd [syz_open_dev$ndb]
ioctl$NBD_DISCONNECT                        : fd_nbd [syz_open_dev$ndb]
ioctl$NBD_DO_IT                             : fd_nbd [syz_open_dev$ndb]
ioctl$NBD_PRINT_DEBUG                       : fd_nbd [syz_open_dev$ndb]
ioctl$NBD_SET_BLKSIZE                       : fd_nbd [syz_open_dev$ndb]
ioctl$NBD_SET_FLAGS                         : fd_nbd [syz_open_dev$ndb]
ioctl$NBD_SET_SIZE                          : fd_nbd [syz_open_dev$ndb]
ioctl$NBD_SET_SIZE_BLOCKS                   : fd_nbd [syz_open_dev$ndb]
ioctl$NBD_SET_SOCK                          : fd_nbd [syz_open_dev$ndb]
ioctl$NBD_SET_TIMEOUT                       : fd_nbd [syz_open_dev$ndb]
ioctl$PPPOEIOCDFWD                          : sock_pppoe [socket$pppoe]
ioctl$PPPOEIOCSFWD                          : sock_pppoe [socket$pppoe]
ioctl$PTP_CLOCK_GETCAPS                     : fd_ptp [openat$ptp0 openat$ptp1]
ioctl$PTP_ENABLE_PPS                        : fd_ptp [openat$ptp0 openat$ptp1]
ioctl$PTP_EXTTS_REQUEST                     : fd_ptp [openat$ptp0 openat$ptp1]
ioctl$PTP_EXTTS_REQUEST2                    : fd_ptp [openat$ptp0 openat$ptp1]
ioctl$PTP_PEROUT_REQUEST                    : fd_ptp [openat$ptp0 openat$ptp1]
ioctl$PTP_PEROUT_REQUEST2                   : fd_ptp [openat$ptp0 openat$ptp1]
ioctl$PTP_PIN_GETFUNC                       : fd_ptp [openat$ptp0 openat$ptp1]
ioctl$PTP_PIN_GETFUNC2                      : fd_ptp [openat$ptp0 openat$ptp1]
ioctl$PTP_PIN_SETFUNC                       : fd_ptp [openat$ptp0 openat$ptp1]
ioctl$PTP_PIN_SETFUNC2                      : fd_ptp [openat$ptp0 openat$ptp1]
ioctl$PTP_SYS_OFFSET                        : fd_ptp [openat$ptp0 openat$ptp1]
ioctl$PTP_SYS_OFFSET_EXTENDED               : fd_ptp [openat$ptp0 openat$ptp1]
ioctl$PTP_SYS_OFFSET_PRECISE                : fd_ptp [openat$ptp0 openat$ptp1]
ioctl$READ_COUNTERS                         : fd_rdma [openat$uverbs0]
ioctl$SIOCAX25ADDFWD                        : sock_ax25 [accept$ax25 accept4$ax25 syz_init_net_socket$ax25]
ioctl$SIOCAX25ADDUID                        : sock_ax25 [accept$ax25 accept4$ax25 syz_init_net_socket$ax25]
ioctl$SIOCAX25CTLCON                        : sock_ax25 [accept$ax25 accept4$ax25 syz_init_net_socket$ax25]
ioctl$SIOCAX25DELFWD                        : sock_ax25 [accept$ax25 accept4$ax25 syz_init_net_socket$ax25]
ioctl$SIOCAX25DELUID                        : sock_ax25 [accept$ax25 accept4$ax25 syz_init_net_socket$ax25]
ioctl$SIOCAX25GETINFO                       : sock_ax25 [accept$ax25 accept4$ax25 syz_init_net_socket$ax25]
ioctl$SIOCAX25GETINFOOLD                    : sock_ax25 [accept$ax25 accept4$ax25 syz_init_net_socket$ax25]
ioctl$SIOCAX25GETUID                        : sock_ax25 [accept$ax25 accept4$ax25 syz_init_net_socket$ax25]
ioctl$SIOCAX25NOUID                         : sock_ax25 [accept$ax25 accept4$ax25 syz_init_net_socket$ax25]
ioctl$SIOCAX25OPTRT                         : sock_ax25 [accept$ax25 accept4$ax25 syz_init_net_socket$ax25]
ioctl$SIOCNRDECOBS                          : sock_netrom [accept$netrom accept4$netrom syz_init_net_socket$netrom]
ioctl$SIOCPNADDRESOURCE                     : sock_phonet [accept$phonet_pipe accept4$phonet_pipe socket$phonet socket$phonet_pipe]
ioctl$SIOCPNDELRESOURCE                     : sock_phonet_dgram [socket$phonet]
ioctl$SIOCPNENABLEPIPE                      : sock_phonet_pipe [accept$phonet_pipe accept4$phonet_pipe socket$phonet_pipe]
ioctl$SIOCPNGETOBJECT                       : sock_phonet [accept$phonet_pipe accept4$phonet_pipe socket$phonet socket$phonet_pipe]
ioctl$SIOCRSACCEPT                          : sock_rose [accept4$rose syz_init_net_socket$rose]
ioctl$SIOCRSGCAUSE                          : sock_rose [accept4$rose syz_init_net_socket$rose]
ioctl$SIOCRSGL2CALL                         : sock_rose [accept4$rose syz_init_net_socket$rose]
ioctl$SIOCRSSCAUSE                          : sock_rose [accept4$rose syz_init_net_socket$rose]
ioctl$SIOCRSSL2CALL                         : sock_rose [accept4$rose syz_init_net_socket$rose]
ioctl$SIOCX25CALLACCPTAPPRV                 : sock_x25 [accept4$x25 syz_init_net_socket$x25]
ioctl$SIOCX25GCALLUSERDATA                  : sock_x25 [accept4$x25 syz_init_net_socket$x25]
ioctl$SIOCX25GCAUSEDIAG                     : sock_x25 [accept4$x25 syz_init_net_socket$x25]
ioctl$SIOCX25GDTEFACILITIES                 : sock_x25 [accept4$x25 syz_init_net_socket$x25]
ioctl$SIOCX25GFACILITIES                    : sock_x25 [accept4$x25 syz_init_net_socket$x25]
ioctl$SIOCX25GSUBSCRIP                      : sock_x25 [accept4$x25 syz_init_net_socket$x25]
ioctl$SIOCX25SCALLUSERDATA                  : sock_x25 [accept4$x25 syz_init_net_socket$x25]
ioctl$SIOCX25SCAUSEDIAG                     : sock_x25 [accept4$x25 syz_init_net_socket$x25]
ioctl$SIOCX25SCUDMATCHLEN                   : sock_x25 [accept4$x25 syz_init_net_socket$x25]
ioctl$SIOCX25SDTEFACILITIES                 : sock_x25 [accept4$x25 syz_init_net_socket$x25]
ioctl$SIOCX25SENDCALLACCPT                  : sock_x25 [accept4$x25 syz_init_net_socket$x25]
ioctl$SIOCX25SFACILITIES                    : sock_x25 [accept4$x25 syz_init_net_socket$x25]
ioctl$SIOCX25SSUBSCRIP                      : sock_x25 [accept4$x25 syz_init_net_socket$x25]
ioctl$SNAPSHOT_ALLOC_SWAP_PAGE              : fd_snapshot [openat$snapshot]
ioctl$SNAPSHOT_ATOMIC_RESTORE               : fd_snapshot [openat$snapshot]
ioctl$SNAPSHOT_AVAIL_SWAP_SIZE              : fd_snapshot [openat$snapshot]
ioctl$SNAPSHOT_CREATE_IMAGE                 : fd_snapshot [openat$snapshot]
ioctl$SNAPSHOT_FREE                         : fd_snapshot [openat$snapshot]
ioctl$SNAPSHOT_FREE_SWAP_PAGES              : fd_snapshot [openat$snapshot]
ioctl$SNAPSHOT_GET_IMAGE_SIZE               : fd_snapshot [openat$snapshot]
ioctl$SNAPSHOT_PLATFORM_SUPPORT             : fd_snapshot [openat$snapshot]
ioctl$SNAPSHOT_PREF_IMAGE_SIZE              : fd_snapshot [openat$snapshot]
ioctl$SNAPSHOT_S2RAM                        : fd_snapshot [openat$snapshot]
ioctl$SNAPSHOT_SET_SWAP_AREA                : fd_snapshot [openat$snapshot]
ioctl$SNAPSHOT_UNFREEZE                     : fd_snapshot [openat$snapshot]
ioctl$SNDCTL_DSP_CHANNELS                   : fd_dsp [openat$adsp1 openat$audio openat$audio1 openat$dsp openat$dsp1]
ioctl$SNDCTL_DSP_GETBLKSIZE                 : fd_dsp [openat$adsp1 openat$audio openat$audio1 openat$dsp openat$dsp1]
ioctl$SNDCTL_DSP_GETCAPS                    : fd_dsp [openat$adsp1 openat$audio openat$audio1 openat$dsp openat$dsp1]
ioctl$SNDCTL_DSP_GETFMTS                    : fd_dsp [openat$adsp1 openat$audio openat$audio1 openat$dsp openat$dsp1]
ioctl$SNDCTL_DSP_GETIPTR                    : fd_dsp [openat$adsp1 openat$audio openat$audio1 openat$dsp openat$dsp1]
ioctl$SNDCTL_DSP_GETISPACE                  : fd_dsp [openat$adsp1 openat$audio openat$audio1 openat$dsp openat$dsp1]
ioctl$SNDCTL_DSP_GETODELAY                  : fd_dsp [openat$adsp1 openat$audio openat$audio1 openat$dsp openat$dsp1]
ioctl$SNDCTL_DSP_GETOPTR                    : fd_dsp [openat$adsp1 openat$audio openat$audio1 openat$dsp openat$dsp1]
ioctl$SNDCTL_DSP_GETOSPACE                  : fd_dsp [openat$adsp1 openat$audio openat$audio1 openat$dsp openat$dsp1]
ioctl$SNDCTL_DSP_GETTRIGGER                 : fd_dsp [openat$adsp1 openat$audio openat$audio1 openat$dsp openat$dsp1]
ioctl$SNDCTL_DSP_NONBLOCK                   : fd_dsp [openat$adsp1 openat$audio openat$audio1 openat$dsp openat$dsp1]
ioctl$SNDCTL_DSP_POST                       : fd_dsp [openat$adsp1 openat$audio openat$audio1 openat$dsp openat$dsp1]
ioctl$SNDCTL_DSP_RESET                      : fd_dsp [openat$adsp1 openat$audio openat$audio1 openat$dsp openat$dsp1]
ioctl$SNDCTL_DSP_SETDUPLEX                  : fd_dsp [openat$adsp1 openat$audio openat$audio1 openat$dsp openat$dsp1]
ioctl$SNDCTL_DSP_SETFMT                     : fd_dsp [openat$adsp1 openat$audio openat$audio1 openat$dsp openat$dsp1]
ioctl$SNDCTL_DSP_SETFRAGMENT                : fd_dsp [openat$adsp1 openat$audio openat$audio1 openat$dsp openat$dsp1]
ioctl$SNDCTL_DSP_SETTRIGGER                 : fd_dsp [openat$adsp1 openat$audio openat$audio1 openat$dsp openat$dsp1]
ioctl$SNDCTL_DSP_SPEED                      : fd_dsp [openat$adsp1 openat$audio openat$audio1 openat$dsp openat$dsp1]
ioctl$SNDCTL_DSP_STEREO                     : fd_dsp [openat$adsp1 openat$audio openat$audio1 openat$dsp openat$dsp1]
ioctl$SNDCTL_DSP_SUBDIVIDE                  : fd_dsp [openat$adsp1 openat$audio openat$audio1 openat$dsp openat$dsp1]
ioctl$SNDCTL_DSP_SYNC                       : fd_dsp [openat$adsp1 openat$audio openat$audio1 openat$dsp openat$dsp1]
ioctl$SNDCTL_FM_4OP_ENABLE                  : fd_seq [openat$sequencer openat$sequencer2]
ioctl$SNDCTL_FM_LOAD_INSTR                  : fd_seq [openat$sequencer openat$sequencer2]
ioctl$SNDCTL_MIDI_INFO                      : fd_seq [openat$sequencer openat$sequencer2]
ioctl$SNDCTL_MIDI_PRETIME                   : fd_seq [openat$sequencer openat$sequencer2]
ioctl$SNDCTL_SEQ_CTRLRATE                   : fd_seq [openat$sequencer openat$sequencer2]
ioctl$SNDCTL_SEQ_GETINCOUNT                 : fd_seq [openat$sequencer openat$sequencer2]
ioctl$SNDCTL_SEQ_GETOUTCOUNT                : fd_seq [openat$sequencer openat$sequencer2]
ioctl$SNDCTL_SEQ_GETTIME                    : fd_seq [openat$sequencer openat$sequencer2]
ioctl$SNDCTL_SEQ_NRMIDIS                    : fd_seq [openat$sequencer openat$sequencer2]
ioctl$SNDCTL_SEQ_NRSYNTHS                   : fd_seq [openat$sequencer openat$sequencer2]
ioctl$SNDCTL_SEQ_OUTOFBAND                  : fd_seq [openat$sequencer openat$sequencer2]
ioctl$SNDCTL_SEQ_PANIC                      : fd_seq [openat$sequencer openat$sequencer2]
ioctl$SNDCTL_SEQ_RESET                      : fd_seq [openat$sequencer openat$sequencer2]
ioctl$SNDCTL_SEQ_RESETSAMPLES               : fd_seq [openat$sequencer openat$sequencer2]
ioctl$SNDCTL_SEQ_SYNC                       : fd_seq [openat$sequencer openat$sequencer2]
ioctl$SNDCTL_SEQ_TESTMIDI                   : fd_seq [openat$sequencer openat$sequencer2]
ioctl$SNDCTL_SEQ_THRESHOLD                  : fd_seq [openat$sequencer openat$sequencer2]
ioctl$SNDCTL_SYNTH_ID                       : fd_seq [openat$sequencer openat$sequencer2]
ioctl$SNDCTL_SYNTH_INFO                     : fd_seq [openat$sequencer openat$sequencer2]
ioctl$SNDCTL_SYNTH_MEMAVL                   : fd_seq [openat$sequencer openat$sequencer2]
ioctl$SNDCTL_TMR_CONTINUE                   : fd_seq [openat$sequencer openat$sequencer2]
ioctl$SNDCTL_TMR_METRONOME                  : fd_seq [openat$sequencer openat$sequencer2]
ioctl$SNDCTL_TMR_SELECT                     : fd_seq [openat$sequencer openat$sequencer2]
ioctl$SNDCTL_TMR_SOURCE                     : fd_seq [openat$sequencer openat$sequencer2]
ioctl$SNDCTL_TMR_START                      : fd_seq [openat$sequencer openat$sequencer2]
ioctl$SNDCTL_TMR_STOP                       : fd_seq [openat$sequencer openat$sequencer2]
ioctl$SNDCTL_TMR_TEMPO                      : fd_seq [openat$sequencer openat$sequencer2]
ioctl$SNDCTL_TMR_TIMEBASE                   : fd_seq [openat$sequencer openat$sequencer2]
ioctl$SNDRV_FIREWIRE_IOCTL_GET_INFO         : fd_snd_hw [syz_open_dev$sndhw]
ioctl$SNDRV_FIREWIRE_IOCTL_LOCK             : fd_snd_hw [syz_open_dev$sndhw]
ioctl$SNDRV_FIREWIRE_IOCTL_TASCAM_STATE     : fd_snd_hw [syz_open_dev$sndhw]
ioctl$SNDRV_FIREWIRE_IOCTL_UNLOCK           : fd_snd_hw [syz_open_dev$sndhw]
ioctl$SNDRV_HWDEP_IOCTL_DSP_LOAD            : fd_snd_hw [syz_open_dev$sndhw]
ioctl$SNDRV_HWDEP_IOCTL_DSP_STATUS          : fd_snd_hw [syz_open_dev$sndhw]
ioctl$SNDRV_HWDEP_IOCTL_INFO                : fd_snd_hw [syz_open_dev$sndhw]
ioctl$SNDRV_HWDEP_IOCTL_PVERSION            : fd_snd_hw [syz_open_dev$sndhw]
ioctl$SNDRV_RAWMIDI_IOCTL_DRAIN             : fd_midi [syz_open_dev$admmidi syz_open_dev$amidi syz_open_dev$dmmidi syz_open_dev$midi syz_open_dev$sndmidi]
ioctl$SNDRV_RAWMIDI_IOCTL_DROP              : fd_midi [syz_open_dev$admmidi syz_open_dev$amidi syz_open_dev$dmmidi syz_open_dev$midi syz_open_dev$sndmidi]
ioctl$SNDRV_RAWMIDI_IOCTL_INFO              : fd_midi [syz_open_dev$admmidi syz_open_dev$amidi syz_open_dev$dmmidi syz_open_dev$midi syz_open_dev$sndmidi]
ioctl$SNDRV_RAWMIDI_IOCTL_PARAMS            : fd_midi [syz_open_dev$admmidi syz_open_dev$amidi syz_open_dev$dmmidi syz_open_dev$midi syz_open_dev$sndmidi]
ioctl$SNDRV_RAWMIDI_IOCTL_PVERSION          : fd_midi [syz_open_dev$admmidi syz_open_dev$amidi syz_open_dev$dmmidi syz_open_dev$midi syz_open_dev$sndmidi]
ioctl$SNDRV_RAWMIDI_IOCTL_STATUS32          : fd_midi [syz_open_dev$admmidi syz_open_dev$amidi syz_open_dev$dmmidi syz_open_dev$midi syz_open_dev$sndmidi]
ioctl$SNDRV_RAWMIDI_IOCTL_STATUS64          : fd_midi [syz_open_dev$admmidi syz_open_dev$amidi syz_open_dev$dmmidi syz_open_dev$midi syz_open_dev$sndmidi]
ioctl$SNDRV_SEQ_IOCTL_CLIENT_ID             : fd_sndseq [openat$sndseq]
ioctl$SNDRV_SEQ_IOCTL_CREATE_PORT           : fd_sndseq [openat$sndseq]
ioctl$SNDRV_SEQ_IOCTL_CREATE_QUEUE          : fd_sndseq [openat$sndseq]
ioctl$SNDRV_SEQ_IOCTL_DELETE_PORT           : fd_sndseq [openat$sndseq]
ioctl$SNDRV_SEQ_IOCTL_DELETE_QUEUE          : fd_sndseq [openat$sndseq]
ioctl$SNDRV_SEQ_IOCTL_GET_CLIENT_INFO       : fd_sndseq [openat$sndseq]
ioctl$SNDRV_SEQ_IOCTL_GET_CLIENT_POOL       : fd_sndseq [openat$sndseq]
ioctl$SNDRV_SEQ_IOCTL_GET_NAMED_QUEUE       : fd_sndseq [openat$sndseq]
ioctl$SNDRV_SEQ_IOCTL_GET_PORT_INFO         : fd_sndseq [openat$sndseq]
ioctl$SNDRV_SEQ_IOCTL_GET_QUEUE_CLIENT      : fd_sndseq [openat$sndseq]
ioctl$SNDRV_SEQ_IOCTL_GET_QUEUE_INFO        : fd_sndseq [openat$sndseq]
ioctl$SNDRV_SEQ_IOCTL_GET_QUEUE_STATUS      : fd_sndseq [openat$sndseq]
ioctl$SNDRV_SEQ_IOCTL_GET_QUEUE_TEMPO       : fd_sndseq [openat$sndseq]
ioctl$SNDRV_SEQ_IOCTL_GET_QUEUE_TIMER       : fd_sndseq [openat$sndseq]
ioctl$SNDRV_SEQ_IOCTL_GET_SUBSCRIPTION      : fd_sndseq [openat$sndseq]
ioctl$SNDRV_SEQ_IOCTL_PVERSION              : fd_sndseq [openat$sndseq]
ioctl$SNDRV_SEQ_IOCTL_QUERY_NEXT_CLIENT     : fd_sndseq [openat$sndseq]
ioctl$SNDRV_SEQ_IOCTL_QUERY_NEXT_PORT       : fd_sndseq [openat$sndseq]
ioctl$SNDRV_SEQ_IOCTL_QUERY_SUBS            : fd_sndseq [openat$sndseq]
ioctl$SNDRV_SEQ_IOCTL_REMOVE_EVENTS         : fd_sndseq [openat$sndseq]
ioctl$SNDRV_SEQ_IOCTL_RUNNING_MODE          : fd_sndseq [openat$sndseq]
ioctl$SNDRV_SEQ_IOCTL_SET_CLIENT_INFO       : fd_sndseq [openat$sndseq]
ioctl$SNDRV_SEQ_IOCTL_SET_CLIENT_POOL       : fd_sndseq [openat$sndseq]
ioctl$SNDRV_SEQ_IOCTL_SET_PORT_INFO         : fd_sndseq [openat$sndseq]
ioctl$SNDRV_SEQ_IOCTL_SET_QUEUE_CLIENT      : fd_sndseq [openat$sndseq]
ioctl$SNDRV_SEQ_IOCTL_SET_QUEUE_INFO        : fd_sndseq [openat$sndseq]
ioctl$SNDRV_SEQ_IOCTL_SET_QUEUE_TEMPO       : fd_sndseq [openat$sndseq]
ioctl$SNDRV_SEQ_IOCTL_SET_QUEUE_TIMER       : fd_sndseq [openat$sndseq]
ioctl$SNDRV_SEQ_IOCTL_SUBSCRIBE_PORT        : fd_sndseq [openat$sndseq]
ioctl$SNDRV_SEQ_IOCTL_SYSTEM_INFO           : fd_sndseq [openat$sndseq]
ioctl$SNDRV_SEQ_IOCTL_UNSUBSCRIBE_PORT      : fd_sndseq [openat$sndseq]
ioctl$SOUND_MIXER_INFO                      : fd_mixer [openat$adsp1 openat$audio openat$audio1 ...]
ioctl$SOUND_MIXER_READ_CAPS                 : fd_mixer [openat$adsp1 openat$audio openat$audio1 ...]
ioctl$SOUND_MIXER_READ_DEVMASK              : fd_mixer [openat$adsp1 openat$audio openat$audio1 ...]
ioctl$SOUND_MIXER_READ_RECMASK              : fd_mixer [openat$adsp1 openat$audio openat$audio1 ...]
ioctl$SOUND_MIXER_READ_RECSRC               : fd_mixer [openat$adsp1 openat$audio openat$audio1 ...]
ioctl$SOUND_MIXER_READ_STEREODEVS           : fd_mixer [openat$adsp1 openat$audio openat$audio1 ...]
ioctl$SOUND_MIXER_READ_VOLUME               : fd_mixer [openat$adsp1 openat$audio openat$audio1 ...]
ioctl$SOUND_MIXER_WRITE_RECSRC              : fd_mixer [openat$adsp1 openat$audio openat$audio1 ...]
ioctl$SOUND_MIXER_WRITE_VOLUME              : fd_mixer [openat$adsp1 openat$audio openat$audio1 ...]
ioctl$SOUND_OLD_MIXER_INFO                  : fd_mixer [openat$adsp1 openat$audio openat$audio1 ...]
ioctl$SOUND_PCM_READ_BITS                   : fd_dsp [openat$adsp1 openat$audio openat$audio1 openat$dsp openat$dsp1]
ioctl$SOUND_PCM_READ_CHANNELS               : fd_dsp [openat$adsp1 openat$audio openat$audio1 openat$dsp openat$dsp1]
ioctl$SOUND_PCM_READ_RATE                   : fd_dsp [openat$adsp1 openat$audio openat$audio1 openat$dsp openat$dsp1]
ioctl$SW_SYNC_IOC_CREATE_FENCE              : fd_sw_sync [openat$sw_sync]
ioctl$SW_SYNC_IOC_INC                       : fd_sw_sync [openat$sw_sync]
ioctl$TE_IOCTL_CLOSE_CLIENT_SESSION         : fd_tlk [openat$tlk_device]
ioctl$TE_IOCTL_LAUNCH_OPERATION             : fd_tlk [openat$tlk_device]
ioctl$TE_IOCTL_OPEN_CLIENT_SESSION          : fd_tlk [openat$tlk_device]
ioctl$TE_IOCTL_SS_CMD                       : fd_tlk [openat$tlk_device]
ioctl$TIPC_IOC_CONNECT                      : fd_trusty [openat$trusty openat$trusty_avb openat$trusty_gatekeeper ...]
ioctl$TIPC_IOC_CONNECT_avb                  : fd_trusty_avb [openat$trusty_avb]
ioctl$TIPC_IOC_CONNECT_gatekeeper           : fd_trusty_gatekeeper [openat$trusty_gatekeeper]
ioctl$TIPC_IOC_CONNECT_hwkey                : fd_trusty_hwkey [openat$trusty_hwkey]
ioctl$TIPC_IOC_CONNECT_hwrng                : fd_trusty_hwrng [openat$trusty_hwrng]
ioctl$TIPC_IOC_CONNECT_keymaster_secure     : fd_trusty_km_secure [openat$trusty_km_secure]
ioctl$TIPC_IOC_CONNECT_km                   : fd_trusty_km [openat$trusty_km]
ioctl$TIPC_IOC_CONNECT_storage              : fd_trusty_storage [openat$trusty_storage]
ioctl$UDMABUF_CREATE                        : fd_udambuf [openat$udambuf]
ioctl$UDMABUF_CREATE_LIST                   : fd_udambuf [openat$udambuf]
ioctl$USBDEVFS_ALLOC_STREAMS                : fd_usbfs [syz_open_dev$usbfs]
ioctl$USBDEVFS_ALLOW_SUSPEND                : fd_usbfs [syz_open_dev$usbfs]
ioctl$USBDEVFS_BULK                         : fd_usbfs [syz_open_dev$usbfs]
ioctl$USBDEVFS_CLAIMINTERFACE               : fd_usbfs [syz_open_dev$usbfs]
ioctl$USBDEVFS_CLAIM_PORT                   : fd_usbfs [syz_open_dev$usbfs]
ioctl$USBDEVFS_CLEAR_HALT                   : fd_usbfs [syz_open_dev$usbfs]
ioctl$USBDEVFS_CONNECTINFO                  : fd_usbfs [syz_open_dev$usbfs]
ioctl$USBDEVFS_CONTROL                      : fd_usbfs [syz_open_dev$usbfs]
ioctl$USBDEVFS_DISCARDURB                   : fd_usbfs [syz_open_dev$usbfs]
ioctl$USBDEVFS_DISCONNECT_CLAIM             : fd_usbfs [syz_open_dev$usbfs]
ioctl$USBDEVFS_DISCSIGNAL                   : fd_usbfs [syz_open_dev$usbfs]
ioctl$USBDEVFS_DROP_PRIVILEGES              : fd_usbfs [syz_open_dev$usbfs]
ioctl$USBDEVFS_FORBID_SUSPEND               : fd_usbfs [syz_open_dev$usbfs]
ioctl$USBDEVFS_FREE_STREAMS                 : fd_usbfs [syz_open_dev$usbfs]
ioctl$USBDEVFS_GETDRIVER                    : fd_usbfs [syz_open_dev$usbfs]
ioctl$USBDEVFS_GET_CAPABILITIES             : fd_usbfs [syz_open_dev$usbfs]
ioctl$USBDEVFS_GET_SPEED                    : fd_usbfs [syz_open_dev$usbfs]
ioctl$USBDEVFS_IOCTL                        : fd_usbfs [syz_open_dev$usbfs]
ioctl$USBDEVFS_REAPURB                      : fd_usbfs [syz_open_dev$usbfs]
ioctl$USBDEVFS_REAPURBNDELAY                : fd_usbfs [syz_open_dev$usbfs]
ioctl$USBDEVFS_RELEASEINTERFACE             : fd_usbfs [syz_open_dev$usbfs]
ioctl$USBDEVFS_RELEASE_PORT                 : fd_usbfs [syz_open_dev$usbfs]
ioctl$USBDEVFS_RESET                        : fd_usbfs [syz_open_dev$usbfs]
ioctl$USBDEVFS_RESETEP                      : fd_usbfs [syz_open_dev$usbfs]
ioctl$USBDEVFS_SETCONFIGURATION             : fd_usbfs [syz_open_dev$usbfs]
ioctl$USBDEVFS_SETINTERFACE                 : fd_usbfs [syz_open_dev$usbfs]
ioctl$USBDEVFS_SUBMITURB                    : fd_usbfs [syz_open_dev$usbfs]
ioctl$USBDEVFS_WAIT_FOR_RESUME              : fd_usbfs [syz_open_dev$usbfs]
ioctl$VFIO_CHECK_EXTENSION                  : fd_vfio [openat$vfio]
ioctl$VFIO_GET_API_VERSION                  : fd_vfio [openat$vfio]
ioctl$VFIO_IOMMU_GET_INFO                   : fd_vfio [openat$vfio]
ioctl$VFIO_IOMMU_MAP_DMA                    : fd_vfio [openat$vfio]
ioctl$VFIO_IOMMU_UNMAP_DMA                  : fd_vfio [openat$vfio]
ioctl$VFIO_SET_IOMMU                        : fd_vfio [openat$vfio]
ioctl$VHOST_NET_SET_BACKEND                 : vhost_net [openat$vnet]
ioctl$VTPM_PROXY_IOC_NEW_DEV                : fd_vtpm [openat$vtpm]
ioctl$mixer_OSS_ALSAEMULVER                 : fd_mixer [openat$adsp1 openat$audio openat$audio1 ...]
ioctl$mixer_OSS_GETVERSION                  : fd_mixer [openat$adsp1 openat$audio openat$audio1 ...]
ioctl$sock_SIOCADDRT                        : nfc_dev_id [ioctl$IOCTL_GET_NCIDEV_IDX]
ioctl$sock_SIOCDELRT                        : nfc_dev_id [ioctl$IOCTL_GET_NCIDEV_IDX]
ioctl$sock_ax25_SIOCADDRT                   : sock_ax25 [accept$ax25 accept4$ax25 syz_init_net_socket$ax25]
ioctl$sock_ax25_SIOCDELRT                   : sock_ax25 [accept$ax25 accept4$ax25 syz_init_net_socket$ax25]
ioctl$sock_bt_bnep_BNEPCONNADD              : sock_bt_bnep [syz_init_net_socket$bt_bnep]
ioctl$sock_bt_bnep_BNEPCONNDEL              : sock_bt_bnep [syz_init_net_socket$bt_bnep]
ioctl$sock_bt_bnep_BNEPGETCONNINFO          : sock_bt_bnep [syz_init_net_socket$bt_bnep]
ioctl$sock_bt_bnep_BNEPGETCONNLIST          : sock_bt_bnep [syz_init_net_socket$bt_bnep]
ioctl$sock_bt_bnep_BNEPGETSUPPFEAT          : sock_bt_bnep [syz_init_net_socket$bt_bnep]
ioctl$sock_bt_cmtp_CMTPCONNADD              : sock_bt_cmtp [syz_init_net_socket$bt_cmtp]
ioctl$sock_bt_cmtp_CMTPCONNDEL              : sock_bt_cmtp [syz_init_net_socket$bt_cmtp]
ioctl$sock_bt_cmtp_CMTPGETCONNINFO          : sock_bt_cmtp [syz_init_net_socket$bt_cmtp]
ioctl$sock_bt_cmtp_CMTPGETCONNLIST          : sock_bt_cmtp [syz_init_net_socket$bt_cmtp]
ioctl$sock_ifreq                            : nfc_dev_id [ioctl$IOCTL_GET_NCIDEV_IDX]
ioctl$sock_inet_sctp_SIOCINQ                : sock_sctp [socket$inet_sctp]
ioctl$sock_kcm_SIOCKCMATTACH                : sock_kcm [socket$kcm]
ioctl$sock_kcm_SIOCKCMCLONE                 : sock_kcm [socket$kcm]
ioctl$sock_kcm_SIOCKCMUNATTACH              : sock_kcm [socket$kcm]
ioctl$sock_netrom_SIOCADDRT                 : sock_netrom [accept$netrom accept4$netrom syz_init_net_socket$netrom]
ioctl$sock_netrom_SIOCDELRT                 : sock_netrom [accept$netrom accept4$netrom syz_init_net_socket$netrom]
ioctl$sock_rose_SIOCADDRT                   : sock_rose [accept4$rose syz_init_net_socket$rose]
ioctl$sock_rose_SIOCDELRT                   : sock_rose [accept4$rose syz_init_net_socket$rose]
ioctl$sock_rose_SIOCRSCLRRT                 : sock_rose [accept4$rose syz_init_net_socket$rose]
ioctl$sock_x25_SIOCADDRT                    : sock_x25 [accept4$x25 syz_init_net_socket$x25]
ioctl$sock_x25_SIOCDELRT                    : sock_x25 [accept4$x25 syz_init_net_socket$x25]
ioctl$vim2m_VIDIOC_CREATE_BUFS              : fd_vim2m [openat$vim2m syz_open_dev$vim2m]
ioctl$vim2m_VIDIOC_DQBUF                    : fd_vim2m [openat$vim2m syz_open_dev$vim2m]
ioctl$vim2m_VIDIOC_ENUM_FMT                 : fd_vim2m [openat$vim2m syz_open_dev$vim2m]
ioctl$vim2m_VIDIOC_ENUM_FRAMESIZES          : fd_vim2m [openat$vim2m syz_open_dev$vim2m]
ioctl$vim2m_VIDIOC_EXPBUF                   : fd_vim2m [openat$vim2m syz_open_dev$vim2m]
ioctl$vim2m_VIDIOC_G_FMT                    : fd_vim2m [openat$vim2m syz_open_dev$vim2m]
ioctl$vim2m_VIDIOC_PREPARE_BUF              : fd_vim2m [openat$vim2m syz_open_dev$vim2m]
ioctl$vim2m_VIDIOC_QBUF                     : fd_vim2m [openat$vim2m syz_open_dev$vim2m]
ioctl$vim2m_VIDIOC_QUERYBUF                 : fd_vim2m [openat$vim2m syz_open_dev$vim2m]
ioctl$vim2m_VIDIOC_QUERYCAP                 : fd_vim2m [openat$vim2m syz_open_dev$vim2m]
ioctl$vim2m_VIDIOC_REQBUFS                  : fd_vim2m [openat$vim2m syz_open_dev$vim2m]
ioctl$vim2m_VIDIOC_STREAMOFF                : fd_vim2m [openat$vim2m syz_open_dev$vim2m]
ioctl$vim2m_VIDIOC_STREAMON                 : fd_vim2m [openat$vim2m syz_open_dev$vim2m]
ioctl$vim2m_VIDIOC_S_CTRL                   : fd_vim2m [openat$vim2m syz_open_dev$vim2m]
ioctl$vim2m_VIDIOC_S_FMT                    : fd_vim2m [openat$vim2m syz_open_dev$vim2m]
ioctl$vim2m_VIDIOC_TRY_FMT                  : fd_vim2m [openat$vim2m syz_open_dev$vim2m]
mmap$DRM_I915                               : fd_i915 [openat$i915]
mmap$DRM_MSM                                : fd_msm [openat$msm]
mmap$KVM_VCPU                               : vcpu_mmap_size [ioctl$KVM_GET_VCPU_MMAP_SIZE]
mmap$bifrost                                : fd_bifrost [openat$bifrost openat$mali]
mmap$dsp                                    : fd_dsp [openat$adsp1 openat$audio openat$audio1 openat$dsp openat$dsp1]
mmap$fb                                     : fd_fb [openat$fb0 openat$fb1]
mmap$usbfs                                  : fd_usbfs [syz_open_dev$usbfs]
mmap$usbmon                                 : fd_usbmon [syz_open_dev$usbmon]
read$alg                                    : sock_algconn [accept$alg accept4$alg]
read$dsp                                    : fd_dsp [openat$adsp1 openat$audio openat$audio1 openat$dsp openat$dsp1]
read$fb                                     : fd_fb [openat$fb0 openat$fb1]
read$hiddev                                 : fd_hiddev [syz_open_dev$hiddev]
read$hidraw                                 : fd_hidraw [syz_open_dev$hidraw]
read$midi                                   : fd_midi [syz_open_dev$admmidi syz_open_dev$amidi syz_open_dev$dmmidi syz_open_dev$midi syz_open_dev$sndmidi]
read$msr                                    : fd_msr [syz_open_dev$MSR]
read$nci                                    : fd_nci [openat$nci]
read$proc_mixer                             : fd_proc_mixer [openat$proc_mixer]
read$ptp                                    : fd_ptp [openat$ptp0 openat$ptp1]
read$sequencer                              : fd_seq [openat$sequencer openat$sequencer2]
read$smackfs_access                         : fd_smackfs_access [openat$smackfs_access]
read$smackfs_cipsonum                       : fd_smackfs_cipsonum [openat$smackfs_cipsonum]
read$smackfs_logging                        : fd_smackfs_logging [openat$smackfs_logging]
read$smackfs_ptrace                         : fd_smackfs_ptrace [openat$smackfs_ptrace]
read$snapshot                               : fd_snapshot [openat$snapshot]
read$sndhw                                  : fd_snd_hw [syz_open_dev$sndhw]
read$trusty                                 : fd_trusty [openat$trusty openat$trusty_avb openat$trusty_gatekeeper ...]
read$usbfs                                  : fd_usbfs [syz_open_dev$usbfs]
read$usbmon                                 : fd_usbmon [syz_open_dev$usbmon]
recvfrom$ax25                               : sock_ax25 [accept$ax25 accept4$ax25 syz_init_net_socket$ax25]
recvfrom$l2tp                               : sock_l2tp [socket$l2tp]
recvfrom$l2tp6                              : sock_l2tp6 [socket$l2tp6]
recvfrom$llc                                : sock_llc [accept4$llc syz_init_net_socket$llc]
recvfrom$netrom                             : sock_netrom [accept$netrom accept4$netrom syz_init_net_socket$netrom]
recvfrom$phonet                             : sock_phonet [accept$phonet_pipe accept4$phonet_pipe socket$phonet socket$phonet_pipe]
recvfrom$rose                               : sock_rose [accept4$rose syz_init_net_socket$rose]
recvfrom$rxrpc                              : sock_rxrpc [socket$rxrpc]
recvfrom$x25                                : sock_x25 [accept4$x25 syz_init_net_socket$x25]
recvmsg$can_j1939                           : sock_can_j1939 [socket$can_j1939]
recvmsg$hf                                  : sock_hf [socket$hf]
recvmsg$kcm                                 : sock_kcm [socket$kcm]
semctl$GETALL                               : ipc_sem [semget semget$private]
semctl$GETNCNT                              : ipc_sem [semget semget$private]
semctl$GETPID                               : ipc_sem [semget semget$private]
semctl$GETVAL                               : ipc_sem [semget semget$private]
semctl$GETZCNT                              : ipc_sem [semget semget$private]
semctl$IPC_INFO                             : ipc_sem [semget semget$private]
semctl$IPC_RMID                             : ipc_sem [semget semget$private]
semctl$IPC_SET                              : ipc_sem [semget semget$private]
semctl$IPC_STAT                             : ipc_sem [semget semget$private]
semctl$SEM_INFO                             : ipc_sem [semget semget$private]
semctl$SEM_STAT                             : ipc_sem [semget semget$private]
semctl$SEM_STAT_ANY                         : ipc_sem [semget semget$private]
semctl$SETALL                               : ipc_sem [semget semget$private]
semctl$SETVAL                               : ipc_sem [semget semget$private]
sendmmsg$alg                                : sock_algconn [accept$alg accept4$alg]
sendmmsg$inet_sctp                          : sock_sctp [socket$inet_sctp]
sendmmsg$nfc_llcp                           : nfc_dev_id [ioctl$IOCTL_GET_NCIDEV_IDX]
sendmsg$L2TP_CMD_NOOP                       : sock_l2tp [socket$l2tp]
sendmsg$L2TP_CMD_SESSION_CREATE             : sock_l2tp [socket$l2tp]
sendmsg$L2TP_CMD_SESSION_DELETE             : sock_l2tp [socket$l2tp]
sendmsg$L2TP_CMD_SESSION_GET                : sock_l2tp [socket$l2tp]
sendmsg$L2TP_CMD_SESSION_MODIFY             : sock_l2tp [socket$l2tp]
sendmsg$L2TP_CMD_TUNNEL_CREATE              : sock_l2tp [socket$l2tp]
sendmsg$L2TP_CMD_TUNNEL_DELETE              : sock_l2tp [socket$l2tp]
sendmsg$L2TP_CMD_TUNNEL_GET                 : sock_l2tp [socket$l2tp]
sendmsg$L2TP_CMD_TUNNEL_MODIFY              : sock_l2tp [socket$l2tp]
sendmsg$NFC_CMD_ACTIVATE_TARGET             : nfc_dev_id [ioctl$IOCTL_GET_NCIDEV_IDX]
sendmsg$NFC_CMD_DEACTIVATE_TARGET           : nfc_dev_id [ioctl$IOCTL_GET_NCIDEV_IDX]
sendmsg$NFC_CMD_DEP_LINK_DOWN               : nfc_dev_id [ioctl$IOCTL_GET_NCIDEV_IDX]
sendmsg$NFC_CMD_DEP_LINK_UP                 : nfc_dev_id [ioctl$IOCTL_GET_NCIDEV_IDX]
sendmsg$NFC_CMD_DEV_DOWN                    : nfc_dev_id [ioctl$IOCTL_GET_NCIDEV_IDX]
sendmsg$NFC_CMD_DEV_UP                      : nfc_dev_id [ioctl$IOCTL_GET_NCIDEV_IDX]
sendmsg$NFC_CMD_DISABLE_SE                  : nfc_dev_id [ioctl$IOCTL_GET_NCIDEV_IDX]
sendmsg$NFC_CMD_ENABLE_SE                   : nfc_dev_id [ioctl$IOCTL_GET_NCIDEV_IDX]
sendmsg$NFC_CMD_FW_DOWNLOAD                 : nfc_dev_id [ioctl$IOCTL_GET_NCIDEV_IDX]
sendmsg$NFC_CMD_GET_DEVICE                  : nfc_dev_id [ioctl$IOCTL_GET_NCIDEV_IDX]
sendmsg$NFC_CMD_LLC_GET_PARAMS              : nfc_dev_id [ioctl$IOCTL_GET_NCIDEV_IDX]
sendmsg$NFC_CMD_LLC_SDREQ                   : nfc_dev_id [ioctl$IOCTL_GET_NCIDEV_IDX]
sendmsg$NFC_CMD_LLC_SET_PARAMS              : nfc_dev_id [ioctl$IOCTL_GET_NCIDEV_IDX]
sendmsg$NFC_CMD_SE_IO                       : nfc_dev_id [ioctl$IOCTL_GET_NCIDEV_IDX]
sendmsg$NFC_CMD_START_POLL                  : nfc_dev_id [ioctl$IOCTL_GET_NCIDEV_IDX]
sendmsg$NFC_CMD_VENDOR                      : nfc_dev_id [ioctl$IOCTL_GET_NCIDEV_IDX]
sendmsg$RDMA_NLDEV_CMD_DELLINK              : sock_nl_rdma [socket$nl_rdma syz_init_net_socket$nl_rdma]
sendmsg$RDMA_NLDEV_CMD_GET                  : sock_nl_rdma [socket$nl_rdma syz_init_net_socket$nl_rdma]
sendmsg$RDMA_NLDEV_CMD_GET_CHARDEV          : sock_nl_rdma [socket$nl_rdma syz_init_net_socket$nl_rdma]
sendmsg$RDMA_NLDEV_CMD_NEWLINK              : sock_nl_rdma [socket$nl_rdma syz_init_net_socket$nl_rdma]
sendmsg$RDMA_NLDEV_CMD_PORT_GET             : sock_nl_rdma [socket$nl_rdma syz_init_net_socket$nl_rdma]
sendmsg$RDMA_NLDEV_CMD_RES_CM_ID_GET        : sock_nl_rdma [socket$nl_rdma syz_init_net_socket$nl_rdma]
sendmsg$RDMA_NLDEV_CMD_RES_CQ_GET           : sock_nl_rdma [socket$nl_rdma syz_init_net_socket$nl_rdma]
sendmsg$RDMA_NLDEV_CMD_RES_GET              : sock_nl_rdma [socket$nl_rdma syz_init_net_socket$nl_rdma]
sendmsg$RDMA_NLDEV_CMD_RES_MR_GET           : sock_nl_rdma [socket$nl_rdma syz_init_net_socket$nl_rdma]
sendmsg$RDMA_NLDEV_CMD_RES_PD_GET           : sock_nl_rdma [socket$nl_rdma syz_init_net_socket$nl_rdma]
sendmsg$RDMA_NLDEV_CMD_RES_QP_GET           : sock_nl_rdma [socket$nl_rdma syz_init_net_socket$nl_rdma]
sendmsg$RDMA_NLDEV_CMD_SET                  : sock_nl_rdma [socket$nl_rdma syz_init_net_socket$nl_rdma]
sendmsg$RDMA_NLDEV_CMD_STAT_DEL             : sock_nl_rdma [socket$nl_rdma syz_init_net_socket$nl_rdma]
sendmsg$RDMA_NLDEV_CMD_STAT_GET             : sock_nl_rdma [socket$nl_rdma syz_init_net_socket$nl_rdma]
sendmsg$RDMA_NLDEV_CMD_STAT_SET             : sock_nl_rdma [socket$nl_rdma syz_init_net_socket$nl_rdma]
sendmsg$RDMA_NLDEV_CMD_SYS_GET              : sock_nl_rdma [socket$nl_rdma syz_init_net_socket$nl_rdma]
sendmsg$RDMA_NLDEV_CMD_SYS_SET              : sock_nl_rdma [socket$nl_rdma syz_init_net_socket$nl_rdma]
sendmsg$alg                                 : sock_algconn [accept$alg accept4$alg]
sendmsg$can_j1939                           : sock_can_j1939 [socket$can_j1939]
sendmsg$hf                                  : sock_hf [socket$hf]
sendmsg$inet_sctp                           : sock_sctp [socket$inet_sctp]
sendmsg$kcm                                 : sock_kcm [socket$kcm]
sendmsg$nfc_llcp                            : nfc_dev_id [ioctl$IOCTL_GET_NCIDEV_IDX]
sendmsg$nl_crypto                           : sock_nl_crypto [socket$nl_crypto]
sendmsg$rds                                 : sock_rds [socket$rds]
sendto$ax25                                 : sock_ax25 [accept$ax25 accept4$ax25 syz_init_net_socket$ax25]
sendto$isdn                                 : sock_isdn [socket$isdn]
sendto$l2tp                                 : sock_l2tp [socket$l2tp]
sendto$l2tp6                                : sock_l2tp6 [socket$l2tp6]
sendto$llc                                  : sock_llc [accept4$llc syz_init_net_socket$llc]
sendto$netrom                               : sock_netrom [accept$netrom accept4$netrom syz_init_net_socket$netrom]
sendto$phonet                               : sock_phonet [accept$phonet_pipe accept4$phonet_pipe socket$phonet socket$phonet_pipe]
sendto$rose                                 : sock_rose [accept4$rose syz_init_net_socket$rose]
sendto$rxrpc                                : sock_rxrpc [socket$rxrpc]
sendto$x25                                  : sock_x25 [accept4$x25 syz_init_net_socket$x25]
setsockopt$ALG_SET_AEAD_AUTHSIZE            : sock_alg [socket$alg]
setsockopt$ALG_SET_KEY                      : sock_alg [socket$alg]
setsockopt$CAIFSO_LINK_SELECT               : sock_caif [socket$caif_seqpacket socket$caif_stream]
setsockopt$CAIFSO_REQ_PARAM                 : sock_caif [socket$caif_seqpacket socket$caif_stream]
setsockopt$MISDN_TIME_STAMP                 : sock_isdn [socket$isdn]
setsockopt$PNPIPE_ENCAP                     : sock_phonet_pipe [accept$phonet_pipe accept4$phonet_pipe socket$phonet_pipe]
setsockopt$PNPIPE_HANDLE                    : sock_phonet_pipe [accept$phonet_pipe accept4$phonet_pipe socket$phonet_pipe]
setsockopt$PNPIPE_INITSTATE                 : sock_phonet_pipe [accept$phonet_pipe accept4$phonet_pipe socket$phonet_pipe]
setsockopt$RDS_CANCEL_SENT_TO               : sock_rds [socket$rds]
setsockopt$RDS_CONG_MONITOR                 : sock_rds [socket$rds]
setsockopt$RDS_FREE_MR                      : sock_rds [socket$rds]
setsockopt$RDS_GET_MR                       : sock_rds [socket$rds]
setsockopt$RDS_GET_MR_FOR_DEST              : sock_rds [socket$rds]
setsockopt$RDS_RECVERR                      : sock_rds [socket$rds]
setsockopt$RXRPC_EXCLUSIVE_CONNECTION       : sock_rxrpc [socket$rxrpc]
setsockopt$RXRPC_MIN_SECURITY_LEVEL         : sock_rxrpc [socket$rxrpc]
setsockopt$RXRPC_SECURITY_KEY               : sock_rxrpc [socket$rxrpc]
setsockopt$RXRPC_SECURITY_KEYRING           : sock_rxrpc [socket$rxrpc]
setsockopt$RXRPC_UPGRADEABLE_SERVICE        : sock_rxrpc [socket$rxrpc]
setsockopt$SO_J1939_ERRQUEUE                : sock_can_j1939 [socket$can_j1939]
setsockopt$SO_J1939_FILTER                  : sock_can_j1939 [socket$can_j1939]
setsockopt$SO_J1939_PROMISC                 : sock_can_j1939 [socket$can_j1939]
setsockopt$SO_J1939_SEND_PRIO               : sock_can_j1939 [socket$can_j1939]
setsockopt$SO_RDS_MSG_RXPATH_LATENCY        : sock_rds [socket$rds]
setsockopt$SO_RDS_TRANSPORT                 : sock_rds [socket$rds]
setsockopt$X25_QBITINCL                     : sock_x25 [accept4$x25 syz_init_net_socket$x25]
setsockopt$ax25_SO_BINDTODEVICE             : sock_ax25 [accept$ax25 accept4$ax25 syz_init_net_socket$ax25]
setsockopt$ax25_int                         : sock_ax25 [accept$ax25 accept4$ax25 syz_init_net_socket$ax25]
setsockopt$inet6_dccp_buf                   : sock_dccp6 [socket$inet6_dccp]
setsockopt$inet6_dccp_int                   : sock_dccp6 [socket$inet6_dccp]
setsockopt$inet_dccp_buf                    : sock_dccp [socket$inet_dccp]
setsockopt$inet_dccp_int                    : sock_dccp [socket$inet_dccp]
setsockopt$inet_sctp6_SCTP_ADAPTATION_LAYER : sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_ADD_STREAMS      : sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_ASSOCINFO        : sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_AUTH_ACTIVE_KEY  : sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_AUTH_CHUNK       : sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_AUTH_DEACTIVATE_KEY: sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_AUTH_DELETE_KEY  : sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_AUTH_KEY         : sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_AUTOCLOSE        : sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_AUTO_ASCONF      : sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_CONTEXT          : sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_DEFAULT_PRINFO   : sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_DEFAULT_SEND_PARAM: sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_DEFAULT_SNDINFO  : sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_DELAYED_SACK     : sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_DISABLE_FRAGMENTS: sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_ENABLE_STREAM_RESET: sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_EVENTS           : sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_FRAGMENT_INTERLEAVE: sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_HMAC_IDENT       : sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_INITMSG          : sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_I_WANT_MAPPED_V4_ADDR: sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_MAXSEG           : sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_MAX_BURST        : sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_NODELAY          : sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_PARTIAL_DELIVERY_POINT: sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_PEER_ADDR_PARAMS : sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_PEER_ADDR_THLDS  : sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_PRIMARY_ADDR     : sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_PR_SUPPORTED     : sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_RECONFIG_SUPPORTED: sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_RECVNXTINFO      : sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_RECVRCVINFO      : sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_RESET_ASSOC      : sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_RESET_STREAMS    : sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_RTOINFO          : sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_SET_PEER_PRIMARY_ADDR: sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_SOCKOPT_BINDX_ADD: sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_SOCKOPT_BINDX_REM: sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_SOCKOPT_CONNECTX : sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_SOCKOPT_CONNECTX_OLD: sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_STREAM_SCHEDULER : sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp6_SCTP_STREAM_SCHEDULER_VALUE: sock_sctp6 [socket$inet6_sctp]
setsockopt$inet_sctp_SCTP_ADAPTATION_LAYER  : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_ADD_STREAMS       : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_ASSOCINFO         : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_AUTH_ACTIVE_KEY   : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_AUTH_CHUNK        : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_AUTH_DEACTIVATE_KEY: sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_AUTH_DELETE_KEY   : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_AUTH_KEY          : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_AUTOCLOSE         : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_AUTO_ASCONF       : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_CONTEXT           : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_DEFAULT_PRINFO    : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_DEFAULT_SEND_PARAM: sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_DEFAULT_SNDINFO   : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_DELAYED_SACK      : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_DISABLE_FRAGMENTS : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_ENABLE_STREAM_RESET: sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_EVENTS            : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_FRAGMENT_INTERLEAVE: sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_HMAC_IDENT        : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_INITMSG           : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_I_WANT_MAPPED_V4_ADDR: sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_MAXSEG            : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_MAX_BURST         : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_NODELAY           : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_PARTIAL_DELIVERY_POINT: sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_PEER_ADDR_PARAMS  : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_PEER_ADDR_THLDS   : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_PRIMARY_ADDR      : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_PR_SUPPORTED      : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_RECONFIG_SUPPORTED: sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_RECVNXTINFO       : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_RECVRCVINFO       : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_RESET_ASSOC       : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_RESET_STREAMS     : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_RTOINFO           : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_SET_PEER_PRIMARY_ADDR: sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_SOCKOPT_BINDX_ADD : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_SOCKOPT_BINDX_REM : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_SOCKOPT_CONNECTX  : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_SOCKOPT_CONNECTX_OLD: sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_STREAM_SCHEDULER  : sock_sctp [socket$inet_sctp]
setsockopt$inet_sctp_SCTP_STREAM_SCHEDULER_VALUE: sock_sctp [socket$inet_sctp]
setsockopt$kcm_KCM_RECV_DISABLE             : sock_kcm [socket$kcm]
setsockopt$llc_int                          : sock_llc [accept4$llc syz_init_net_socket$llc]
setsockopt$netrom_NETROM_IDLE               : sock_netrom [accept$netrom accept4$netrom syz_init_net_socket$netrom]
setsockopt$netrom_NETROM_N2                 : sock_netrom [accept$netrom accept4$netrom syz_init_net_socket$netrom]
setsockopt$netrom_NETROM_T1                 : sock_netrom [accept$netrom accept4$netrom syz_init_net_socket$netrom]
setsockopt$netrom_NETROM_T2                 : sock_netrom [accept$netrom accept4$netrom syz_init_net_socket$netrom]
setsockopt$netrom_NETROM_T4                 : sock_netrom [accept$netrom accept4$netrom syz_init_net_socket$netrom]
setsockopt$rose                             : sock_rose [accept4$rose syz_init_net_socket$rose]
syz_io_uring_submit                         : nfc_dev_id [ioctl$IOCTL_GET_NCIDEV_IDX]
syz_kvm_add_vcpu                            : kvm_syz_vm [syz_kvm_setup_syzos_vm]
syz_kvm_setup_cpu$arm64                     : fd_kvmvm [ioctl$KVM_CREATE_VM]
syz_kvm_setup_syzos_vm                      : fd_kvmvm [ioctl$KVM_CREATE_VM]
syz_kvm_vgic_v3_setup                       : fd_kvmvm [ioctl$KVM_CREATE_VM]
syz_memcpy_off$KVM_EXIT_HYPERCALL           : kvm_run_ptr [mmap$KVM_VCPU]
syz_memcpy_off$KVM_EXIT_MMIO                : kvm_run_ptr [mmap$KVM_VCPU]
write$6lowpan_control                       : fd_6lowpan_control [openat$6lowpan_control]
write$6lowpan_enable                        : fd_6lowpan_enable [openat$6lowpan_enable]
write$ALLOC_MW                              : fd_rdma [openat$uverbs0]
write$ALLOC_PD                              : fd_rdma [openat$uverbs0]
write$ATTACH_MCAST                          : fd_rdma [openat$uverbs0]
write$CLOSE_XRCD                            : fd_rdma [openat$uverbs0]
write$CREATE_AH                             : fd_rdma [openat$uverbs0]
write$CREATE_COMP_CHANNEL                   : fd_rdma [openat$uverbs0]
write$CREATE_CQ                             : fd_rdma [openat$uverbs0]
write$CREATE_CQ_EX                          : fd_rdma [openat$uverbs0]
write$CREATE_FLOW                           : fd_rdma [openat$uverbs0]
write$CREATE_QP                             : fd_rdma [openat$uverbs0]
write$CREATE_RWQ_IND_TBL                    : fd_rdma [openat$uverbs0]
write$CREATE_SRQ                            : fd_rdma [openat$uverbs0]
write$CREATE_WQ                             : fd_rdma [openat$uverbs0]
write$DEALLOC_MW                            : fd_rdma [openat$uverbs0]
write$DEALLOC_PD                            : fd_rdma [openat$uverbs0]
write$DEREG_MR                              : fd_rdma [openat$uverbs0]
write$DESTROY_AH                            : fd_rdma [openat$uverbs0]
write$DESTROY_CQ                            : fd_rdma [openat$uverbs0]
write$DESTROY_FLOW                          : fd_rdma [openat$uverbs0]
write$DESTROY_QP                            : fd_rdma [openat$uverbs0]
write$DESTROY_RWQ_IND_TBL                   : fd_rdma [openat$uverbs0]
write$DESTROY_SRQ                           : fd_rdma [openat$uverbs0]
write$DESTROY_WQ                            : fd_rdma [openat$uverbs0]
write$DETACH_MCAST                          : fd_rdma [openat$uverbs0]
write$MLX5_ALLOC_PD                         : fd_rdma [openat$uverbs0]
write$MLX5_CREATE_CQ                        : fd_rdma [openat$uverbs0]
write$MLX5_CREATE_DV_QP                     : fd_rdma [openat$uverbs0]
write$MLX5_CREATE_QP                        : fd_rdma [openat$uverbs0]
write$MLX5_CREATE_SRQ                       : fd_rdma [openat$uverbs0]
write$MLX5_CREATE_WQ                        : fd_rdma [openat$uverbs0]
write$MLX5_GET_CONTEXT                      : fd_rdma [openat$uverbs0]
write$MLX5_MODIFY_WQ                        : fd_rdma [openat$uverbs0]
write$MODIFY_QP                             : fd_rdma [openat$uverbs0]
write$MODIFY_SRQ                            : fd_rdma [openat$uverbs0]
write$OPEN_XRCD                             : fd_rdma [openat$uverbs0]
write$POLL_CQ                               : fd_rdma [openat$uverbs0]
write$POST_RECV                             : fd_rdma [openat$uverbs0]
write$POST_SEND                             : fd_rdma [openat$uverbs0]
write$POST_SRQ_RECV                         : fd_rdma [openat$uverbs0]
write$QUERY_DEVICE_EX                       : fd_rdma [openat$uverbs0]
write$QUERY_PORT                            : fd_rdma [openat$uverbs0]
write$QUERY_QP                              : fd_rdma [openat$uverbs0]
write$QUERY_SRQ                             : fd_rdma [openat$uverbs0]
write$RDMA_USER_CM_CMD_ACCEPT               : fd_rdma_cm [openat$rdma_cm]
write$RDMA_USER_CM_CMD_BIND                 : fd_rdma_cm [openat$rdma_cm]
write$RDMA_USER_CM_CMD_BIND_IP              : fd_rdma_cm [openat$rdma_cm]
write$RDMA_USER_CM_CMD_CONNECT              : fd_rdma_cm [openat$rdma_cm]
write$RDMA_USER_CM_CMD_CREATE_ID            : fd_rdma_cm [openat$rdma_cm]
write$RDMA_USER_CM_CMD_DESTROY_ID           : fd_rdma_cm [openat$rdma_cm]
write$RDMA_USER_CM_CMD_DISCONNECT           : fd_rdma_cm [openat$rdma_cm]
write$RDMA_USER_CM_CMD_GET_EVENT            : fd_rdma_cm [openat$rdma_cm]
write$RDMA_USER_CM_CMD_INIT_QP_ATTR         : fd_rdma_cm [openat$rdma_cm]
write$RDMA_USER_CM_CMD_JOIN_IP_MCAST        : fd_rdma_cm [openat$rdma_cm]
write$RDMA_USER_CM_CMD_JOIN_MCAST           : fd_rdma_cm [openat$rdma_cm]
write$RDMA_USER_CM_CMD_LEAVE_MCAST          : fd_rdma_cm [openat$rdma_cm]
write$RDMA_USER_CM_CMD_LISTEN               : fd_rdma_cm [openat$rdma_cm]
write$RDMA_USER_CM_CMD_MIGRATE_ID           : fd_rdma_cm [openat$rdma_cm]
write$RDMA_USER_CM_CMD_NOTIFY               : fd_rdma_cm [openat$rdma_cm]
write$RDMA_USER_CM_CMD_QUERY                : fd_rdma_cm [openat$rdma_cm]
write$RDMA_USER_CM_CMD_QUERY_ROUTE          : fd_rdma_cm [openat$rdma_cm]
write$RDMA_USER_CM_CMD_REJECT               : fd_rdma_cm [openat$rdma_cm]
write$RDMA_USER_CM_CMD_RESOLVE_ADDR         : fd_rdma_cm [openat$rdma_cm]
write$RDMA_USER_CM_CMD_RESOLVE_IP           : fd_rdma_cm [openat$rdma_cm]
write$RDMA_USER_CM_CMD_RESOLVE_ROUTE        : fd_rdma_cm [openat$rdma_cm]
write$RDMA_USER_CM_CMD_SET_OPTION           : fd_rdma_cm [openat$rdma_cm]
write$REG_MR                                : fd_rdma [openat$uverbs0]
write$REQ_NOTIFY_CQ                         : fd_rdma [openat$uverbs0]
write$REREG_MR                              : fd_rdma [openat$uverbs0]
write$RESIZE_CQ                             : fd_rdma [openat$uverbs0]
write$USERIO_CMD_REGISTER                   : fd_userio [openat$userio]
write$USERIO_CMD_SEND_INTERRUPT             : fd_userio [openat$userio]
write$USERIO_CMD_SET_PORT_TYPE              : fd_userio [openat$userio]
write$capi20                                : fd_capi20 [openat$capi20]
write$capi20_data                           : fd_capi20 [openat$capi20]
write$damon_attrs                           : fd_damon_attrs [openat$damon_attrs]
write$damon_contexts                        : fd_damon_contexts [openat$damon_mk_contexts openat$damon_rm_contexts]
write$damon_init_regions                    : fd_damon_init_regions [openat$damon_init_regions]
write$damon_monitor_on                      : fd_damon_monitor_on [openat$damon_monitor_on]
write$damon_schemes                         : fd_damon_schemes [openat$damon_schemes]
write$damon_target_ids                      : fd_damon_target_ids [openat$damon_target_ids]
write$dsp                                   : fd_dsp [openat$adsp1 openat$audio openat$audio1 openat$dsp openat$dsp1]
write$fb                                    : fd_fb [openat$fb0 openat$fb1]
write$hidraw                                : fd_hidraw [syz_open_dev$hidraw]
write$midi                                  : fd_midi [syz_open_dev$admmidi syz_open_dev$amidi syz_open_dev$dmmidi syz_open_dev$midi syz_open_dev$sndmidi]
write$nci                                   : fd_nci [openat$nci]
write$proc_mixer                            : fd_proc_mixer [openat$proc_mixer]
write$proc_reclaim                          : fd_proc_reclaim [openat$proc_reclaim]
write$sequencer                             : fd_seq [openat$sequencer openat$sequencer2]
write$smackfs_access                        : fd_smackfs_access [openat$smackfs_access]
write$smackfs_change_rule                   : fd_smackfs_change_rule [openat$smackfs_change_rule]
write$smackfs_cipso                         : fd_smackfs_cipso [openat$smackfs_cipso]
write$smackfs_cipsonum                      : fd_smackfs_cipsonum [openat$smackfs_cipsonum]
write$smackfs_ipv6host                      : fd_smackfs_ipv6host [openat$smackfs_ipv6host]
write$smackfs_label                         : fd_smackfs_label [openat$smackfs_ambient openat$smackfs_revoke_subject openat$smackfs_syslog openat$smackfs_unconfined]
write$smackfs_labels_list                   : fd_smackfs_labels_list [openat$smackfs_onlycap openat$smackfs_relabel_self]
write$smackfs_load                          : fd_smackfs_load [openat$smackfs_load]
write$smackfs_logging                       : fd_smackfs_logging [openat$smackfs_logging]
write$smackfs_netlabel                      : fd_smackfs_netlabel [openat$smackfs_netlabel]
write$smackfs_ptrace                        : fd_smackfs_ptrace [openat$smackfs_ptrace]
write$snapshot                              : fd_snapshot [openat$snapshot]
write$sndhw                                 : fd_snd_hw [syz_open_dev$sndhw]
write$sndhw_fireworks                       : fd_snd_hw [syz_open_dev$sndhw]
write$sndseq                                : fd_sndseq [openat$sndseq]
write$sysctl                                : fd_sysctl [openat$sysctl]
write$trusty                                : fd_trusty [openat$trusty openat$trusty_avb openat$trusty_gatekeeper ...]
write$trusty_avb                            : fd_trusty_avb [openat$trusty_avb]
write$trusty_gatekeeper                     : fd_trusty_gatekeeper [openat$trusty_gatekeeper]
write$trusty_hwkey                          : fd_trusty_hwkey [openat$trusty_hwkey]
write$trusty_hwrng                          : fd_trusty_hwrng [openat$trusty_hwrng]
write$trusty_km                             : fd_trusty_km [openat$trusty_km]
write$trusty_km_secure                      : fd_trusty_km_secure [openat$trusty_km_secure]
write$trusty_storage                        : fd_trusty_storage [openat$trusty_storage]
write$usbip_server                          : fd_usbip_server [syz_usbip_server_init]
write$vga_arbiter                           : fd_vga_arbiter [openat$vga_arbiter]
write$vhost_msg                             : vhost_net [openat$vnet]
write$vhost_msg_v2                          : vhost_net [openat$vnet]
write$yama_ptrace_scope                     : fd_yama_ptrace_scope [openat$yama_ptrace_scope]

failed to read the following files in the VM:
/sys/kernel/security/lsm                    : No such file or directory

BinFmtMisc              : enabled
Comparisons             : enabled
Coverage                : enabled
DelayKcovMmap           : enabled
DevlinkPCI              : PCI device 0000:00:10.0 is not available
ExtraCoverage           : enabled
Fault                   : CONFIG_FAULT_INJECTION is not enabled
KCSAN                   : write(/sys/kernel/debug/kcsan, on) failed
LRWPANEmulation         : NL802154_CMD_SET_SHORT_ADDR failed
Leak                    : failed to write(kmemleak, "scan=off")
NetDevices              : enabled
NetInjection            : enabled
NicVF                   : PCI device 0000:00:11.0 is not available
SandboxAndroid          : enabled
SandboxNamespace        : NOTFAIL: sandbox fork failed.  (errno 22: Invalid argument). . process exited with status 67.
SandboxNone             : enabled
SandboxSetuid           : enabled
Swap                    : enabled
USBEmulation            : failed to chmod /dev/raw-gadget
VhciInjection           : unshare(CLONE_NEWPID): 22. spawned loop pid 10429. NOTFAIL: open /dev/vhci failed.  (errno 2: No such file or directory).
WifiEmulation           : unshare(CLONE_NEWPID): 22. spawned loop pid 10802. unshare(CLONE_NEWIPC): 22. netlink: failed to get family id for MAC80211_HWSIM: No such file or directory. NOTFAIL: netlink_query_family_id failed.  (errno 2: No such file or directory).
syscalls                : 2627/5291


2025/02/14 16:17:33 corpus                  : 274 (274 seeds), 0 to be reminimized, 0 to be resmashed
2025/02/14 16:17:38 candidates=255 corpus=0 coverage=0 exec total=635 (92/min) pending=0 reproducing=0 
2025/02/14 16:17:48 candidates=188 corpus=0 coverage=0 exec total=702 (100/min) pending=0 reproducing=0 


```

