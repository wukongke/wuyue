# 无阅--章节模块

	提供无阅--章节模块的接口描述。


## 0x00 <span id = "main">目录</span>

|路径|方法|描述|
|:---|:---|:---|
|[/app/api/chapters](#chapter-list)|get|章节信息列表|
|[/app/api/chapters/:chapterNo](#chapter-get)|get|获取章节信息|


## 0x01 请求方式
###### HTTP请求

|环境|HTTP 请求地址|
|:---|:---|
|环境|http://127.0.0.1/app

##  <span id = "chapter-list">get /app/api/chapters</span>
####  [返回目录](#main)

######  描述:  小说信息列表

###### query请求参数:

|参数|类型|是否必填|描述|示例值|
|:---|:---:|:---:|:---|:---|
|page|number|N|default: 1|-|
|limit|number|N|default: 15|-|
|sort|string|N|排序|-title,name|
|bookNo|int|N|小说编号|8634|
|search|string|N|章节名称、章节编号||

###### 响应参数:

|参数|类型|是否必填|描述|示例值|
|:---|:---:|:---:|:---|:---|
|code|int|Y|返回码|0 = 正常<br>其他值表示错误或者异常|
|message|int|Y|返回错误信息|-|
|data|object|Y|[章节结构](#chapter-structure)| 详细见下面请求示例|

###### 响应示例:

```
{
    "code": 0, 
    "msg": "操作成功",
    "data": {
        "limit": 3, 
        "page": 1, 
        "pages": 396, 
        "total": 1187,
        "docs": [
            {
                "bookNo": "8634", 
                "chapterNo": "9858714", 
                "content": "    新国三十三年春，生了很多事情 （全文终）(本章完) ", 
                "createdAt": 1501555353, 
                "id": "597fda4eac454d131cb04f2c", 
                "status": 1, 
                "title": "第1185章 神隐之路", 
                "updatedAt": 1501555353, 
                "url": "http://www.zwdu.com/book/8634/9858714.html"
            }, 
            {
                "bookNo": "8634", 
                "chapterNo": "9832706", 
                "content": "    王之策的唇角露出一抹自嘲的笑容，眼神有些伤感。 (本章完) ", 
                "createdAt": 1501555352, 
                "id": "597fda4dac454d131cb04f29", 
                "status": 1, 
                "title": "第1182章 黑袍之死", 
                "updatedAt": 1501555352, 
                "url": "http://www.zwdu.com/book/8634/9832706.html"
            }, 
            {
                "bookNo": "8634", 
                "chapterNo": "9840413", 
                "content": "    从墓地爬出来的怪人是黑袍。  这是回答陈长生刚才的那句话。     因为有些突然，陈长生怔了怔，才做出回答。     “是的。”     (本章完) ", 
                "createdAt": 1501555352, 
                "id": "597fda4dac454d131cb04f2a", 
                "status": 1, 
                "title": "第1183章 天凉好个秋", 
                "updatedAt": 1501555352, 
                "url": "http://www.zwdu.com/book/8634/9840413.html"
            }
        ],
    }, 
}

```

##  <span id = "chapter-get">get /app/api/chapters/:chapterNo</span>
####  [返回目录](#main)

######  描述: 获取章节信息

###### query请求参数:

|参数|类型|是否必填|描述|示例值|
|:---|:---:|:---:|:---|:---|

###### 响应参数:

|参数|类型|是否必填|描述|示例值|
|:---|:---:|:---:|:---|:---|
|code|int|Y|返回码|0 = 正常<br>其他值表示错误或者异常|
|message|int|Y|返回错误信息|-|
|data|object|Y|[章节结构](#chapter-structure)| 详细见下面请求示例|

###### 响应示例:

```
{
  code: 0,
  message: "操作成功",
  data: {
    "bookNo": "8634", 
    "chapterNo": "9858714", 
    "content": "    新国三十三年春，生了很多事情 （全文终）(本章完) ", 
    "createdAt": 1501555353, 
    "id": "597fda4eac454d131cb04f2c", 
    "status": 1, 
    "title": "第1185章 神隐之路", 
    "updatedAt": 1501555353, 
    "url": "http://www.zwdu.com/book/8634/9858714.html"
  }, 
}
```