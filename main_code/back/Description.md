# 后端使用说明


## 配置文件:

位置:`/main_code/back/src/SMS_config/conf.toml`

内容:
```toml
[DB]
Address = "localhost"           # 数据库所在的位置(ip)
User = "YuanCheng_user"         # 数据库用户名
Password  = "_Hexagram_user"    # 数据库用户密码
DBName = "test_python"          # 数据库名

[Web]
Port = ":1221"                  # 使用的接口,注意':'不能少
```


## 源代码:
位置:`/main_code/back/src`

源文件:
- [main.go 程序入口](src/main.go)
- [service.go 服务模型](src/service.go)
    - [config.go 读取配置文件](src/config.go)
    - [database.go 初始化数据库](src/database.go)
    - [router.go 定义并开启服务](src/router.go)
        - [Allsystem.go 获取所有事务](src/AllSystem.go)
        - [DailySystem.go 操作每日任务](src/DailySystem.go)
        - [EventsSystem.go 操作长期事务](src/EventsSystem.go)
- [error.go 处理程序上的错误](src/error.go)