# Glados-checkin
Glados自动签到脚本，通过配置文件配置Cookie,每日执行时间,超时时间

### 使用方式

1.

```shell
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build .

#将得到的二进制文件上传至linux服务器，config.yml文件需要置于同级目录
```

2.

```yaml
#编辑yml文件

cookie: xxxx:xxxxxxx; xxx:xxxx #cookie
execTime: 20:00:00             #每日签到时间
timeout: 10					   #超时时间，单位为秒
```

3.

运行脚本

```shell
chmod 777  glados-checkin #赋予权限

nohup glados-checkin &   #启动脚本

ps aux |grep glados-checkin #查看运行状态
```

