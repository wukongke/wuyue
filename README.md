
## Echo, Go, Vue : 无阅——最全面的阅读平台。

> 技术栈：golang + echo + mgo + vue

> 预览地址： [http://47.94.11.23/wuyue](http://47.94.11.23/wuyue)

> api地址：[https://github.com/wukongke/wuyue](https://github.com/wukongke/wuyue)

> 爬虫地址：[https://github.com/wukongke/wuspider](https://github.com/wukongke/wuspider)


### 项目运行

```
1.环境安装
  配置好GOROOT和GOPATH
  cd src
  mkdir work-codes
  git clone https://code.aliyun.com/wukc/wuyue.git
  cd wuyue/app
  glide install

2. 运行项目
  启动mongodb:  mongod
  启动项目： bee run  或者 go build 

  运行测试： go test -v

3. 访问项目：
  浏览器输入： localhost:1323

4. 测试接口：
  postman: 
    post 请求接收json数据

5. 前端编译：
  cd view
  cnpm i
  npm run build 后生成dist文件夹

6. 前后端结合： 
  将编译后的dist文件， 复制到app下， 并将名称改为public
  cd view 
  cp -R dist ../app && cd ../app && mv dist public

```

### 项目部署

```
docker + caddy

```
