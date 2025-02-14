#!/bin/bash

# 获取当前的用户名
USERNAME=$(whoami)

# 获取当前时间，格式：yyyy-mm-dd HH:MM:SS
CURRENT_TIME=$(date "+%Y-%m-%d %H:%M:%S")

# 创建提交信息
COMMIT_MESSAGE="${USERNAME} ${CURRENT_TIME}"

# 提交更改
git add .
git commit -m "$COMMIT_MESSAGE"
git push