
- 测试命令： 

```
curl -X GET 'http://localhost:8080/someJSON'
curl -X GET 'http://localhost:8080/moreJSON'
curl -X GET 'http://localhost:8080/someXML'
curl -X GET 'http://localhost:8080/someYAML'
curl -X GET 'http://localhost:8080/someTOML'
curl -X GET 'http://localhost:8080/someProtoBuf'
curl -X GET 'http://localhost:8080/SecureJSON'
curl -X GET 'http://localhost:8080/JSON'
curl -X GET 'http://localhost:8080/JSONP'
curl -X GET 'http://localhost:8080/JSONP?callback=sayHello'
curl -X GET 'http://localhost:8080/AsciiJSON'
curl -X GET 'http://localhost:8080/PureJSON'
```

- JSONP示例：
    - 参考：https://www.flysnow.org/2020/01/01/golang-gin-jsonp-and-hijacking
    - 浏览器输入：http://localhost:8080/
    - JavaScript脚本调用：curl -X GET 'http://localhost:8080/JSONP?callback=sayHello'
    - 该接口将callback提供的参数作为函数名，数据作为函数参数，构造待执行的函数语句，如下所示：
        ```
        $ curl -v -X GET 'http://localhost:8080/JSONP?callback=sayHello'
        Note: Unnecessary use of -X or --request, GET is already inferred.
        *   Trying 127.0.0.1:8080...
        * Connected to localhost (127.0.0.1) port 8080 (#0)
        > GET /JSONP?callback=sayHello HTTP/1.1
        > Host: localhost:8080
        > User-Agent: curl/7.79.1
        > Accept: */*
        >
        * Mark bundle as not supporting multiuse
        < HTTP/1.1 200 OK
        < Content-Type: application/javascript; charset=utf-8
        < Date: Wed, 30 Nov 2022 03:08:55 GMT
        < Content-Length: 24
        <
        * Connection #0 to host localhost left intact
        sayHello({"foo":"bar"});
        ```
    - 返回的Content-Type为：application/javascript，Body为一个函数：sayHello({"foo":"bar"});
    - 浏览器执行该函数，出现弹框：{"foo":"bar"}

