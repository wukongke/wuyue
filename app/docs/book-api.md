# 无阅--书籍模块

	提供无阅--书籍模块的接口描述。


## 0x00 <span id = "main">目录</span>

|路径|方法|描述|
|:---|:---|:---|
|[/app/api/books](#book-list)|get|小说信息列表|
|[/app/api/books/:bookNo](#book-get)|get|获取小说信息|
|[/app/api/logs](#log-list)|get|日志列表|


## 0x01 请求方式
###### HTTP请求

|环境|HTTP 请求地址|
|:---|:---|
|环境|http://127.0.0.1/app

##  <span id = "book-list">get /app/api/books</span>
####  [返回目录](#main)

######  描述:  小说信息列表

###### query请求参数:

|参数|类型|是否必填|描述|示例值|
|:---|:---:|:---:|:---|:---|
|page|number|N|default: 1|-|
|limit|number|N|default: 15|-|
|sort|string|N|排序|-title,name|
|search|string|N|小说名称、小说编号||

###### 响应参数:

|参数|类型|是否必填|描述|示例值|
|:---|:---:|:---:|:---|:---|
|code|int|Y|返回码|0 = 正常<br>其他值表示错误或者异常|
|message|int|Y|返回错误信息|-|
|data|object|Y|[小说结构](#book-structure)| 详细见下面请求示例|

###### 响应示例:

```
{
    "code": 0, 
    "message": "操作成功", 
    "data": {
        "limit": 1, 
        "page": 1, 
        "pages": 34, 
        "total": 34,
        "docs": [
            {
                "id": "598056fdac454d131cb0bf78", 
                "accountId": "598056fdac454d131cb0bf77", 
                "bookNo": "10026", 
                "name": "逆天邪神", 
                "image": "http://www.zwdu.com/files/article/image/10/10026/10026s.jpg", 
                "intro": " 天毒之珠，承邪神之血，修逆天之力，一代邪神，君临天下！ 各位书友要是觉得《逆天邪神》还不错的话请不要忘记向您QQ群和微博里的朋友推荐哦！ ", 
                "lastChapter": "第888章 苓儿，苓儿（上）", 
                "likes": [
                    "新奉系", 
                    "韩娱之梦", 
                    "蛊门", 
                    "重生之山村传奇", 
                    "上古世纪之诺亚之战", 
                    "重生潜入梦", 
                    "电影位面大冒险", 
                    "铁血大明", 
                    "天下", 
                    "驱魔人"
                ], 
                "loveCount": 1000, //收藏人数
                "rate": 5, 
                "serialize": 0, // 0:连载中， 1：完结
                "status": 1, 
                "type": "玄幻", 
                "url": "http://www.zwdu.com/book/10026", 
                "wordCount": 100.5,
                "createdAt": 1501602269,
                "updatedAt": 1501625690, 
            }
        ], 
    }
}

```

##  <span id = "book-get">get /app/api/books/:bookNo</span>
####  [返回目录](#main)

######  描述: 获取小说信息

###### query请求参数:

|参数|类型|是否必填|描述|示例值|
|:---|:---:|:---:|:---|:---|

###### 响应参数:

|参数|类型|是否必填|描述|示例值|
|:---|:---:|:---:|:---|:---|
|code|int|Y|返回码|0 = 正常<br>其他值表示错误或者异常|
|message|int|Y|返回错误信息|-|
|data|object|Y|[书籍结构](#book-structure)| 详细见下面请求示例|

###### 响应示例:

```
{
  code: 0,
  message: "操作成功",
  data: {
    id: "597fd79aac454d131cb04ab4",
    accountId: "597fe292ac454d131cb05098",
    bookNo: "8634",
    createdAt: 1501554851,
    image: "http://www.zwdu.com/files/article/image/8/8634/8634s.jpg",
    intro: " 太始元年，有神石自太空飞来，分散落在人间，其中落在东土大陆的神石，上面镌刻着奇怪的图腾，人因观其图腾而悟道，后立国教。数千年后，十四岁的少年孤儿陈长生，为治病改命离开自己的师父，带着一纸婚约来到神都，从而开启了一个逆天强者的崛起征程。 各位书友要是觉得《择天记》还不错的话请不要忘记向您QQ群和微博里的朋友推荐哦！",
    lastChapter: "第1185章 神隐之路",
    likes: [
      "来自秦朝的你",
      "盗梦无限",
      "仙王不朽",
      "合体双修",
      "女修宗门男掌教",
      "吕氏外戚",
      "武神空间",
      "异界逍遥系统",
      "寻烟",
      "电竞英雄"
    ],
    loveCount: 1000,
    name: "择天记",
    rate: 5,
    serialize: 0,
    status: 1,
    type: "玄幻",
    updatedAt: 1493928302,
    url: "http://www.zwdu.com/book/8634",
    wordCount: 100.5
  }
}
```