

# 处理 web 程序的输入与输出

## 静态文件服务以及实现JS访问

键入localhost:8080，得到Index.html以及附带的js文件：

![](https://raw.githubusercontent.com/ZhangZekun/images/master/cloudService_homework/hw4/1.PNG)

注意到，index.html文件如下:

~~~html
<html>
    <head>
    <script type="text/javascript" src="./js/jquery.js"></script>
    <script type="text/javascript" src="./js/jquery.query-2.1.7.js"></script>
    <script src="http://code.jquery.com/jquery-latest.js"></script>
      <script src="js/hello.js"></script>
    </head>
    <body>
      Sample Go Web Application!!
          <div>
              <p class="greeting-id">The ID is </p>
              <p class="greeting-content">The content is </p>
              <p >Please type In localhost:8080/login  to log in and then get the table </p>
          </div>
    </body>
    </html>
~~~

hello.js文件如下：

~~~~javascript
$(document).ready(function() {
    $.ajax({
        url: "/api/test"
    }).then(function(data) {
       $('.greeting-id').append(data.id);
       $('.greeting-content').append(data.content);
    });
});
~~~~

hello.js在html渲染完毕后，访问/api/test，并将获得到的数据填充到html中。

这实现了js访问。

## 表单处理以及模板输出

表单输入用户名和密码。成功登陆后，输出用户名，以及一个表格，该表格是利用模板输出的。

![](https://raw.githubusercontent.com/ZhangZekun/images/master/cloudService_homework/hw4/2.PNG)

![](https://raw.githubusercontent.com/ZhangZekun/images/master/cloudService_homework/hw4/3.PNG)

模板如下:

~~~html
<html>
    <body>
    
    <h>hello {{.Username}}</h>
    <p>this is the table</p>
    <table border="1">
      <tr>
        <th>姓名</th>
        <th>手机号码</th>
      </tr>
      <tr>
        <td>{{.Username1}}</td>
        <td>{{.Phone1}}</td>
      </tr>
      <tr>
        <td>{{.Username2}}</td>
        <td>{{.Phone2}}</td>
      </tr>
    </table>
    
    </body>
</html>
~~~

## 中间件设计

这里我设计了一个简单的中间件，在路由中间件被调用前，先调用我的中间件。我的中间件把request中的req.Form["password"]末尾再添加两个值，'"wwwwww", "eeeeeeee"'。可以再终端看到，对req的修改被成功地反应了出来：

![](https://raw.githubusercontent.com/ZhangZekun/images/master/cloudService_homework/hw4/4.PNG)



## 对/api/unknown或者/api/XXXX， 只要不是/api/login都会报501错误

![](https://raw.githubusercontent.com/ZhangZekun/images/master/cloudService_homework/hw4/5.PNG)





