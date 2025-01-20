### 基本操作

### git基本配置

```git
git config --global user.name 用户名
git config --global user.eamil 邮箱
```
### git设置代理
```git
git config --global -e
```

### git提交流程

```git
git add --all #提交所有文件
git commit -m "Initial commit" #放到缓存区
git push -u origin main #提交到github仓库
```

### git配置代理

```
//全局代理
git config --global http.proxy http://127.0.0.1:10809
git config --global https.proxy https://127.0.0.1:10809

//只对github.com使用代理，其他仓库不走代理
git config --global http.https://github.com.proxy socks5://127.0.0.1:1080
git config --global https.https://github.com.proxy socks5://127.0.0.1:1080
```

