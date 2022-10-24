## 跳板机隧道工具


#### 详情看`ssh_rw.yaml`


```yaml
sshrw: v1
config:
  tunnel: "jump@fuckserver.slb-bibibi.com.cn:9102"
  password: "aaaaa!1100"
  destinations: ['123.111.111.124:8088->localhost:2293','192.2.1.131:8084->localhost:2222']

```


##### release 


```bash

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o sshrw-linux-x86_64 ./
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build  -o sshrw-windows-x86_64 ./

```


