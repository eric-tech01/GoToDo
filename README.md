# Scan App Using Wails

## 联系人

```
carlzfq@163.com
```

## 环境基础

node
golang
wails

## 构建开发

```
#如果 frontend/wailsjs 目录不存在，先执行如下命令
wails generate module

#如果首次打包提示 dist目录不存在，先执行下面命令
cd frontend
npm i
npm run build

#检测
wails doctor

#调试
wails dev

#调试遇到错误： failed to find Vite server URL，解决办法：通过指定 frontenddevserverurl 参数
wails dev -frontenddevserverurl "http://localhost:8080"

#打包 .exe
wails build
```

## 设计思路

1. 系统初始化：系统启动，初始化一个账户、初始化相关系统配置；
2. 登录：分为无账号登录 and 鉴权登录
   - 无账号登录：系统配置 autologin，则系统初始化启动不进入登录界面，直接进入主界面；且屏蔽相关登录操作事件；
   - 鉴权登录：需要输入账号密码进行登录；
