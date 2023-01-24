
- 参考：https://github.com/gin-gonic/gin/issues/742
- 测试命令： 

```
# only bind query
$ curl -X GET "localhost:8085/testing?name=eason&address=xyz"

# only bind query string, ignore form data
$ curl -X POST "localhost:8085/testing?name=eason&address=xyz" --data 'name=ignore&address=ignore' -H "Content-Type:application/x-www-form-urlencoded"
```

