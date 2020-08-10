## 跳板机隧道工具


#### 详情看`ssh_rw.yaml`


```yaml
sshrw: v1
config:
  tunnel: "admin@jumpserver.stc-xxxxc.com.cn:6593"
  password: "aaaaa!1100"
  destinations: ['172.22.0.13:8088->localhost:2293','172.22.0.13:8084->localhost:2222']

```

