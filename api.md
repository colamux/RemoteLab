# 登录状态 采用全局中间件


# / 
主页
登录与否用中间件判断
1. 未登录：实验室间接 右上角登录 注册按钮
2. 登录： 主页(包含设备申请)

# /{username}
个人中心，课程成绩

# source
管理、申请

## api 实验人员 
|url|method||request|response|
|-|-|-|-|-|
|/login|POST|登录认证|username、password|token
|/signup|POST|注册|username、password、tel|
|/api/apply|POST|申请设备|type|result
|/api/getscore|GET|查看分数||score(json)

## api 教师
|url|method||request|response|
|-|-|-|-|-|
|/api/getcourse|GET|查看设备列表|
|/api/setcourse|POST|

## api 管理人员
|url|method||request|response|
|-|-|-|-|-|
|/api/getdev|GET|查看设备列表|
|/api/setdev|POST|
