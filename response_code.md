# Response Code

api端点设计中httpcode码并不能完全表达资源访问的状态，比如token失效，一般都是401错误，但是401还表示资源访问认证失败，有歧义，所以响应要入自己的状态码 一般四位 5开头

## 自定义状态码字典

5010 access_token失效
5011 refresh_token失效
