# lark-web-hook

## 说明
作为webhook服务转换grafana至lark群机器人webhook的监控请求

## 使用方法

```shell
cd /opt
git clone https://github.com/ukiyowood/lark-web-hook.git
cd lark-web-hook
export LARK_WEBHOOK_URI="https://open.larksuite.com/open-apis/bot/v2/hook/xxx"
nohup ./webhook-linux 2>&1 &
```