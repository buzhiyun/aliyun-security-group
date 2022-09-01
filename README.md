# 管理阿里云的安全组



#### 数据库中三张表

- firewall_records  （防火墙记录）

- security_groups   （安全组列表）

- businesses        （业务线）



#### 参数

###### aliyun-security-group:
    -check    检查现有记录DNS解析地址是否有效
    -migrate  初始化数据表结构
    -update   检查失效IP，并更新 （会自动先执行check）



#### 配置文件

###### config.yaml

```yaml
aliyun:  # 阿里云 accessKey 和 accessSecret
  key: "aliyun_key"
  secret: "aliyun_secret"

mysql:   # 数据库配置
  dsn: "user:passwd@tcp(mysql.address:3306)/dbname?charset=utf8&parseTime=True&loc=Local"


handler:  # 变更处理器 
  template_dir: "/host_template_dir"  # hosts模板文件路径,路径下面是  {handlerName}/hosts
  handlers:
    uat-win:    # handlerName
      type: "hosts"   # handlerType ,目前可选的是 hosts|command 。hosts 会更新模板hosts，再执行command 给推送到主机 ;command 直接执行，不会有其他处理
      args: ["true"]  # true 表示是windows 的模板文件, hosts文件换行是CRLF
      cmd: # 执行的命令列表
        - command: "ansible"
          args: ["uat_*", "-m", "win_copy", "-a", "src=/host_template_dir/uat-win/hosts dest=C:/Windows/System32/Drivers/etc/hosts", "-f", "100"]

    uat-linux:
      type: "command"
      cmd:
        - command: "ssh"
          args: ["septnet@127.0.0.1", "sudo /update_scripts_dir/updateHosts"]

```

