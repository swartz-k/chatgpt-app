# ChatGPT app
ChatGPT 后端 + 小程序端    

## 使用指南
Golang 实现 session 的方式接入 ChatGPT 服务，详细逻辑可以参考 `pkg/chat` 部分

### session 来源
登陆后，打开调试，复制名为 `__Secure-next-auth.session-token` 的Cookie 的值即可
![image](docs/images/session-token.png)


### 服务端
1. 复制并修改配置文件，主要是获取 session    
`cp config.json.example config.json`

2. 启动服务    
`make local-run`

### 小程序端
1. 启动服务    
`yarn && yarn dev:weapp`

2. 发布版本，需要修改 `app/config/prod.js`，指定 `env.URL`