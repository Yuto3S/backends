Following https://gowebexamples.com tutorial

```
dlv debug hello.go
list main.renderHomePage
break ./hello.go:13

(dlv) locals
test_int_for_debug_dlv = 5
test_debug = 4

(dlv) args
writer = net/http.ResponseWriter(*net/http.response) 0x14000157530
request = ("*net/http.Request")(0x1400014ac60)

print request.URL
call request.URL.Query()


```

SETUP MYSQL DOCKER (MacOS M1)
```
docker pull --platform linux/x86_64 mysql
docker run --platform linux/amd64 --name mysqlgo -p 3306:3306 -e MYSQL_ROOT_PASSWORD=test-go-mysql -d mysql
docker exec -it #DOCKER_ID bash
bash-4.4# mysql -u root -p
mysql> CREATE DATABASE testgo;
Query OK, 1 row affected (0.05 sec)

mysql> USE testgo;
Database changed
```