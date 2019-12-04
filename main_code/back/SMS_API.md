# SMS API
##### Schedule Manage System

baseURL:`localhost:1221`


### GET `/allAffairs` 获取所有事务
exampleURL:`localhost:1221/allAffairs`
Success(200):
```json
{
  "data": [
    {
      "ID": 5, //主键
      "CreatedAt": "0001-01-01T00:00:00Z", //创建时间
      "UpdatedAt": "2019-12-03T11:25:56+08:00", //更新时间
      "DeletedAt": null, //空
      "Title": "C语言总结", //事件主体
      "Deadline": "1111-11-11T00:00:02+08:00", //截至日期(死线)
      "Extra": "sssss" //格外补充
    },
    {
      "ID": 8, //同上
      "CreatedAt": "0001-01-01T00:00:00Z",
      "UpdatedAt": "0001-01-01T00:00:00Z",
      "DeletedAt": null,
      "Title": "英语课介绍用PPT",
      "Deadline": "1111-11-11T11:11:11+08:00",
      "Extra": ""
    },
  ],
  "error": 0,
  "msg": "success"
}
```


### GET `/opera/find/:id` 通过ID获取事务
exampleURL:`localhost:1221/opera/find/5`
```json
{
    "data": {
        "ID": 5, //返回主键值与ID相同的数据
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "2019-12-03T11:25:56+08:00",
        "DeletedAt": null,
        "Title": "C语言总结",
        "Deadline": "1111-11-11T00:00:02+08:00",
        "Extra": "sssss"
    },
    "error": 0,
    "msg": "success"
}
```

事务不存在(传入的ID值不存在)(300)
```json
{
    "error": 30000,
    "msg": "affair not exist"
}
```

URL格式错误(300)
```json
{
    "error": 30000,
    "msg": "wrong format"
}
```


### GET `/dailyEvents` 获取每日事务
exampleURL:`localhost:1221/dailyEvents`
Success(200)
```json
{
    "data": [
        {
            "Title": "英语单词打卡",
            "Extra": ""
        },
        {
            "Title": "每天的作业整理",
            "Extra": ""
        }
    ],
    "error": 0,
    "msg": "success"
}
```


### POST `/opera/add` 添加事务
exampleURL:`localhost:1221/opera/add`
Payload:
```json
    {
        "Title":"EampleTitle",
        "Extra":"ExtraExample",
        "Deadline":"1111-11-11 11-11-11" //1111-11-11 11-11-11默认没有deadline
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


### DELETE `/opera/:id` 删除事务
exampleURL:`localhost:1221/opera/5`

Success(200)
```json
{
    "data": "",
    "error": 0,
    "msg": "success"
}
```

事务不存在(300)
```json
{
    "error": 30000,
    "msg": "affair doesn't exist"
}
```

URL格式错误(300)
```json
{
    "error": 30000,
    "msg": "wrong format"
}
```

### PUT `/opera/:id` 修改事务
exampleURL:`localhost:1221/opera/65`

Payload
```json
    {
        "Title":"Title",
        "Extra":"Extra",
        "Deadline":"1111-11-11 11-11-11"
    }
```

Success(200)
```json
{
    "data": 20000,
    "error": 0,
    "msg": "success"
}
```

事务不存在(300)
```json
{
    "error": 30000,
    "msg": "affair doesn't exist"
}
```

URL格式错误(300)
```json
{
    "error": 30000,
    "msg": "wrong format"
}
```