# golang beego postgresql swagger restful api demo
	   bee   :1.4.1
	   beego :1.6.0
	   Go    :go version go1.5rc1 darwin/amd64
# 先执行提供的数据库脚本,无误后进行以下步骤 
>### psql -U [leonard]用户名 -d [demo]数据库  < db.sql
![](https://github.com/leonardyp/go-postgresql-demo/blob/master/static/img/db.png)  

1,将项目 放到gopath/src下

2,cd 项目目录

3,数据库配置在conf 下 请更改 相应user和db即可

4,在项目根目录下:
  
	bee run -gendoc=true
5,打开浏览器
# 访问http://ip:8080/api  
*api都可以在上面查看测试*

![](https://github.com/leonardyp/go-postgresql-demo/blob/master/static/img/api.png)

-------
# 附个人资料:
[github地址](https://github.com/leonardyp)

[博客](http://leonardyp.github.io/)