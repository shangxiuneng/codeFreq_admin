## codeFreq网站的后端
技术栈
golang + hertz

接口的返回规范
```
{
    "code": 5000,  // 错误码 返回0代表成功
    "msg": "",  // 错误信息 可供前端开发人员排查
    "prompt": "", // 提示信息 用户可见
    "data": {  // 数据信息

    }
}
```