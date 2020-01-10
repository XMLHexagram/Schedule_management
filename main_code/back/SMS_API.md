# SMS_API
###### Schedule Manage System

约定:
事务->长期事务和短期事务

baseURL:`localhost:1221`


#### 总事务获取 

##### GET `/all/affairs` 获取所有事务

Success(200)
```json
{
    "data": [
        {
            "id": 7,
            "title": "PHP留言板",
            "deadline": "1111-11-11 11:11:11",
            "extra": "",
            "created_at": "2019-10-20 13:26:50"
        },
        {
            "id": 8,
            "title": "英语课介绍用PPT",
            "deadline": "1111-11-11 11:11:11",
            "extra": "",
            "created_at": "2019-10-20 13:27:48"
        },
    ],
    "error": 0,
    "msg": "success"
}
```

##### GET `/all/dailyAffairs` 获取所有每日任务

Success(200)
```json
{
    "data": [
        {
            "ID": 9,
            "CreatedAt": "0001-01-01T00:00:00Z",
            "UpdatedAt": "0001-01-01T00:00:00Z",
            "DeletedAt": null,
            "Title": "英语单词打卡",
            "Extra": ""
        },
    ],
    "error": 0,
    "msg": "success"
}
```


#### 操作事务

##### POST `/opera/add` 增

Payload:
```json
{
	"title":"test", //required
	"deadline":"2000-01-01T00:00:00Z",
	"extra":"extra"
}
```

Success(200):
```json
{
    "data": "",
    "error": 0,
    "msg": "success"
}
```

##### DELETE `/opera` 删

Query:

* `id` 事务编号

Success(200)
```json
{
    "data": "",
    "error": 0,
    "msg": "success"
}
```

##### PUT `/opera` 改

Query:

* `id` 事务编号

Payload:
```json
{
	"title":"test111", //required
	"deadline":"3000-01-01T00:00:00Z",
	"extra":"extra!!!"
}
```

Success(200)
```json
{
    "data": "",
    "error": 0,
    "msg": "success"
}
```


#### 操作每日任务

##### POST `/operaDaily/add` 增

Payload:
```json
{
	"title":"test", //required
	"extra":"extra"
}
```

Success(200)
```json
{
    "data": "",
    "error": 0,
    "msg": "success"
}
```

##### DELETE `/operaDaily` 删

Query:

* `id` 每日任务编号

Success(200):
```json
{
    "data": "",
    "error": 0,
    "msg": "success"
}
```

##### PUT `/operaDaily` 改

Query:

* `id` 每日任务编号

```json
{
	"title":"test111", //required
	"extra":"extra!!!"
}
```

Success(200):
```json
{
    "data": 20000,
    "error": 0,
    "msg": "success"
}
```

|错误码对照表|含义|msg|
---|---|---|
|40000|json格式错误|`Wrong Format Of JSON`|
|40400|传入的参数无法解析|`Unable To Parse arameters`|
|40410|访问的数据不存在|`Not Found`|
|50000|合法数据无法插入数据库|`Can't Insert Into Database`|



