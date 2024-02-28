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

print r.URL.Query().Get("token")
```