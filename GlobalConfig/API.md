# API List

response格式
```bat
{
    "code": int,
    "msg": string,
    "data": {...}
}
```

*Global Host (Backend)*: http://localhost:1926


## 用户user

### **注册**

```bat
[POST]
{host} /user/register
[param]
body:
{
    "Name": "string",               // 10位以内的字母或数字组合且必须以字母开头
    "Email": "string",              // 邮箱校验
    "Pwd": "string",                // 输入密码：密码不可短于8位，且必须包含字母数字特殊符号
    "Pwd_confirm": "string",        // 确认密码 
    "SecQ": "string",               // 密保问题
    "SecA": "string"                // 密保答案
}
```
*response*

空字段
```bat
[response]
{
    code: -200,
    message: "Name字段不能为空",               // Name == nil
    data: {nil}
}
{
    code: -201,
    message: "Email字段不能为空",               // Email == nil
    data: {nil}
}
{
    code: -202,
    message: "密码不能为空",                    // Pwd == nil
    data: {nil}
}
{
    code: -203,
    message: "请确认密码",                      // Pwd_confirm == nil
    data: {nil}
}
{
    code: -204,
    message: "请输入密保问题",                  // SecQ == nil
    data: {nil}
}
{
    code: -205,
    message: "请输入密保答案",                  // SecA == nil
    data: {nil}
}
```
校验
```bat
[response]
{
    code: -100,
    message: "...(Name要求提示)",               // Name字段校验失败
    data: {Name: string}
}
{
    code: -101,
    message: "...(Email要求提示)",              // Email字段校验失败
    data: {Email: string}
}
{
    code: -102,
    message: "两次输入的密码不一致",
    data: {nil}
}
{
    code: -103,
    message: "用户名或邮箱已被使用",
    data: {Name: string, Email: string}
}
```
注册成功
```bat
[response]
{
    code: 000,
    message: "注册成功，请返回登录",
    data: {Name: string, Email: string, Id: string}
    // Id: Hash code
}
```

### **登录**

```bat
[POST]
{host} /users/login
[param]
body:
{
    "N_E": "string",               // 输入用户名或邮箱
    "Pwd": "string"
}
```

校验
```bat
{
    code: -100,
    message: "...(Name要求提示)",               // Name字段校验失败
    data: {Name: string}
}
{
    code: -101,
    message: "...(Email要求提示)",              // Email字段校验失败
    data: {Email: string}
}
```
密码错误
```bat
[response]
{
    code: -300,
    message: "密码错误",
    data: {nil}
}
```
登录成功
```bat
[response]
{
    code: 001,
    message: "欢迎使用",
    data: {nil}
}
```
### **登出**
```bat
[POST]
{host} /users/logout
[param]
body:{}
```
已退出账户
```bat
[response]
{
    code: 004,
    message: "您已退出账户",
    data: {nil}
}
```




### **获取用户信息**

获取完整信息
```bat
[GET]
{host}  /user/account/get/all
[param]
body: {}
```
response
```bat
[response]
{
    code: -999,
    message: "请先登录",
    data: {nil}
}
{
    code: 999,
    message: "用户信息",
    data: {Name:string, Email:string, SecQ:string, Id:string}
}
```

获取密保问题
```bat
[GET]
{host}  /user/account/get/secq
[param]
body: {}
```
response
```bat
[response]
{
    code: -999,
    message: "请先登录",
    data: {nil}
}
{
    code: 998,
    message: "密保问题",
    data: {SecQ:string}
}
```



### **修改密码(原密码)**

```bat
[POST]
{host}  /user/account/password/pwd
[param]
body:
{
    "N_E": "string"             // 输入用户名或邮箱
    "Pwd_ori": "string"         // 原始密码
    "Pwd_new": "string"         // 新密码
}.
```
失败
```bat
[response]
{
    code: -104,
    message: "用户不存在",
    data: {nil}
}
{
    code: -105,
    message: "原始密码错误",
    data: {nil}
}
```
成功
```bat
[response]
{
    code: 002,
    message: "密码修改成功",
    data: {nil}
}
```


### **修改密码（密保）**

```bat
[POST]
{host}  /user/account/password/sec
[param]
body:
{
    "N_E": "string"             // 输入用户名或邮箱
    "SecA_now": "string"        // 密保问题答案 -> 这里先展示得到的密保问题(/user/account/get/secq)
    "Pwd_new": "string"         // 新密码
}
```
失败
```bat
[response]
{
    code: -104,
    message: "用户不存在",
    data: {nil}
}
{
    coda: -106,
    message: "密保问题错误",
    data: {SecA_now}
}
```
成功
```bat
[response]
{
    code: 003,
    message: "密码修改成功",
    data: {nil}
}
```

## 绘本

（文本和图片内容分别点击按键上传，分别response）

### **输入描述文本**
```bat
[GET]
{host} /picbook/text/input_text
[param]
body:
{
    "Text": "string"        // 前端实时统计输入文本字数（不超过300字）
}
```

字数超限
```bat
[response]
{
    code: -301,
    message: "输入不应超过300字",
    data: {le(Text):int}
}
```
有效输入
```bat
[response]
{
    code: 300,
    message: "accepted text",       // 前端无需显示
    data: {nil}
}
```



### **上传描述文本文件**
```bat
[POST]
{host} /picbook/text/input_file
[param]
body:
{
    "Text_file": "multipart.file"     // 不超过30KB
}
```

文件大小超限
```bat
[response]
{
    code: -302,
    message: "上传文本文件不应超过30KB",
    data: {Text_file.size:int}
}
```
有效输入
```bat
[response]
{
    code: 301,
    message: "accepted text file",      // 前端无需显示
    data: {nil}
}
```


### **使用文本标签**
（用勾选形式，勾选后进入该模式展示文本标签，前端返回一个0-1串？）

```bat
[POST]
{host} /picbook/text/input_label
[param]
body:
{
    "Label_List": [] bool
}
```


无选中的label
```bat
[response]
{
    code: -303,
    message: "请选择需要的基础文本标签",        // 后端全零布尔数组输入
    data: {nil}
}
```
有效label选择
```bat
[response]
{
    code: 302,
    message: "label accepted",      // 前端无需显示
    data: {nil}
}
```



### 注入风格的图片
```bat
[POST]
{host} /picbook/image/style_img/upload
[param]
body:
{
    "Style_img": "image/png | image/jpeg"
}
```

图片存储过大
```bat
[response]
{
    code: -304,
    message: "图像文件不得超过20M",
    data: {nil}
}
```
图片尺寸过大
```bat
[response]
{
    code: -305,
    message: "图像像素不得超过1024*1024",
    data: {nil}
}
```
有效Style上传
```bat
[response]
{
    code: 303,
    message: "style accepted",      // 前端无需显示
    data: {nil}
}
```




### 手绘草图
（可以的话在网页上直接开放用鼠标光标绘制草图，再做一个url）
```bat
[POST]
{host} /picbook/image/sketch_img/upload
[param]
body:
{
    "Sketch_img": "image/png | image/jpeg"
}
```
图片存储过大
```bat
[response]
{
    code: -304,
    message: "图像文件不得超过20M",
    data: {nil}
}
```
图片尺寸过大
```bat
[response]
{
    code: -305,
    message: "图像像素不得超过1024*1024",
    data: {nil}
}
```
有效Sketch上传
```bat
[response]
{
    code: 304,
    message: "sketch accepted",      // 前端无需显示
    data: {nil}
}
```





```bat
[POST]
{host}
[param]
body:
{

}
```



```
[response]
{

}
```

