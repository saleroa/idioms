# **_功能实现_**


基础功能

● 用户注册
●  用户登录
●  发布问题
●  回答问题
●  发布文章
●  评论回答和评论文章
●  我的（个人信息栏）
●  用户信息更改（昵称，密码等）
●  我的收藏

加分项

● 用户状态保存使用 JWT         
●  搜索功能（搜索问题或文章）                
● 关注功能
● 用户密码加盐加密
● 点赞功能（赞同问题、文章、评论）         






# **_接口实现_**


//在查看问题或者想法的回答，评论之前，需要后端给前端传一段比较简单的数据，做成用户滑动浏览的界面，
如果点击某一个，再携带着后端传过来的数据再进行下一步的请求


## 用户层

注册 POST  /user/register

| 必选  | username | string |
|-----|----------|--------|
| 必选  | password | string |

登录    POST /user/login
 
| 必选  | username | string |
|-----|----------|--------|
| 必选  | password | string |

修改密码 PUT /user/changepass

| 必选  | newpasss | string |
|-----|----------|--------|

完善用户信息 POST /user/insertinfo

|   | nicknam   | string |
|---|-----------|--------|
|   | signature | string |
|   | educate   |   string     |
|   | gender    |   string     |
|   | reside    |    string    |
|   | sector    |    string    |


修改用户信息 PUT /user/changeinfo

|   | nicknam   | string |
|---|-----------|--------|
|   | signature | string |
|   | educate   |   string |
|   | gender    |   string |
|   | reside    | string |
|   | sector    |    string |

查看用户信息 GET /user/getinfo


## 问题层

提出问题 POST /question/create

| 必选  | content | string |
|-----|---------|--------|

回答问题 POST /question/answer
**_//需要先传入一个json， Aid struct {
Aid int `json:"aid"`
}_**

| 必选  | content | string |
|-----|---------|--------|

评论回答 POST /question/remark
**_//需要先传入一个json， Aid struct {
Aid int `json:"aid"`
}_**

| 必选  | content | string |
|-----|---------|--------|

关键字搜索问题 GET /question/searchnbykey
**_//返回的是创作者的id和问题的id以及问题内容，用来做浏览页面，
点击后返回问题的id给后端，用id查询，即下一个接口_**

| 必选  | key | string |
|-----|-----|--------|

根据问题的id查询问题 GET /question/searchbyqid
**_//需要先传入一个json， Aid struct {
Aid int `json:"aid"`
}_**
*_*//返回的是问题，回答，评论的复合体，可以用来做深入浏览问题的页面**_


##  文章层

写文章 POST /article/cerate

| 必选  | content | string |
|-----|---------|--------|

评论文章 POST /article/remark
**_//在这个之前需要先返回一个json，Arid struct {
Arid int `json:"arid"`
}_**

| 必选  | content | string |
|-----|---------|--------|
 
关键字搜索文章 GET /article/searchbykey
**_//返回的是文章的内容+文章的id，创作者的id，用来做浏览页面，点击后返回文章的id，并执行下一个api_**

| 必选  | key | string |
|-----|-----|--------|
 
根据文章的id查询文章 GET /article/searchbyarid
**_//在这个之前需要先返回一个json，Arid struct {
Arid int `json:"arid"`
}_**
_**//返回的是文章和评论的复合体，用于深度浏览的界面**_


## 收藏关注

关注用户 POST /attention/follow
**_//关注之前需要点进去用户界面，返回给后端一个json，包含该用户的id_**

收藏问题 POST /attention/question
//收藏之前需要点进去，获取以下数据，再传回后端json， Question struct {
Qid     int    `json:"qid"`
Uid     int    `json:"uid"`
Content string `json:"content"`
}

收藏文章 POST /attention/article
//收藏之前需要先点进去，获取以下数据，再传回后端json， Article struct {
Arid    int    `json:"arid"`
Uid     int    `json:"uid"`
Content string `json:"content"`
}

查看收藏的文章  GET /attention/getarticle

查看收藏的问题  GET /attention/getquestion

根据关注的用户推荐问题  GET /attention/articlebyfollow

根据关注的用户推荐想法文章  GET /attention/questionbyfollow

## 点赞层
**_//在进行一下几个造作之前都要先传一个jsontype Info1 struct {
Kind   string `json:"kind"`
Creid  int    `json:"creid"`
}**_

查看自己是否点赞 GET /givelike/islike

点赞文章/问题   GET /givelike/likeset

获取该创作的点赞数 GET /givelike/likecount

