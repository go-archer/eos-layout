# API项目布局

## 目录结构

```bash
.
├── README.md
├── cmd -- 应用入口
│   └── server -- http服务入口
│       ├── config.toml
│       ├── main.go
│       ├── wire.go
│       └── wire_gen.go
├── go.mod
├── go.sum
├── internal -- 内部代码
│   ├── config -- 配置解析
│   │   └── config.go
│   ├── dto -- 数据传输对象
│   │   ├── request -- 请求数据对象
│   │   │   └── area.go
│   │   └── response -- 响应数据对象
│   │       └── area.go
│   ├── handler -- 请求处理，调用业务逻辑服务（service），返回http响应
│   │   ├── area.go
│   │   └── handler.go
│   ├── middleware -- 中间件，用于处理请求和响应，如日志、签名等
│   │   ├── cors.go
│   │   └── logger.go
│   ├── model -- 数据模型，定义业务逻辑需要等数据结构
│   │   └── area.go
│   ├── repository -- 数据访问对象，封装数据库操作
│   │   ├── area.go
│   │   └── repository.go
│   ├── router -- 路由层，根据版本号统一管理路由配置
│   │   └── v1
│   │       └── v1.go
│   ├── server
│   │   └── http.go
│   ├── service -- 业务逻辑层，嗲用数据访问层（repository）
│   │   ├── area.go
│   │   └── service.go
│   └── status -- 状态码
│       └── status.go
├── pkg -- 公共代码
│   ├── http
│   │   └── http.go
│   ├── log
│   │   ├── log.go
│   │   └── option.go
│   ├── md5
│   │   └── md5.go
│   └── uuid
│       └── uuid.go
├── scripts -- 脚本文件，用于部署和其他自动化任务
└── test -- 测试代码


```

## wire

```bash
# 安装
go install github.com/google/wire/cmd/wire@latest

# 执行
cd cmd/
wire
```

## 请求验证使用说明

### 自定义验证

```bash
# 结构体
func customFunc(fl verifier.FieldLevel) bool {
    if fl.Field().String() == "invalid" {
        return false
    }
    return true
}
validate.RegisterValidation("custom tag name", customFunc)

# 注意：使用与现有函数相同的标记名称将覆盖现有函数
```

### 字段说明

| 字段                          | 说明                                                                                                                                                                     |
| ----------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| -                             | 忽略字段，告诉验证跳过这个struct字段                                                                                                                                     |
| &#124;                        | 'or'运算符，允许使用和接受多个验证器                                                                                                                                     |
| structonly                    | 当遇到嵌套结构的字段并包含此标志时，仅验证结构体，不验证任何结构体字段                                                                                                   |
| nostructlevel                 | 与structonly标记相同，但不会运行任何结构级别验证                                                                                                                         |
| omitempty                     | 允许条件验证                                                                                                                                                             |
| dive                          | 告诉验证者潜入切片，数组或映射，并使用后面的验证标记验证切片，数组或映射的该级别                                                                                         |
| required                      | 验证该值不是数据类型的默认零值。数字不为０，字符串不为 " ", slices, maps, pointers, interfaces, channels and functions不为nil                                            |
| isdefault                     | 验证该值是默认值                                                                                                                                                         |
| len=10                        | 对于数字，长度将确保该值等于给定的参数。对于字符串，它会检查字符串长度是否与字符数完全相同。对于切片，数组和map，验证元素个数                                            |
| max=10                        | 对于数字，max将确保该值小于或等于给定的参数。对于字符串，它会检查字符串长度是否最多为该字符数。对于切片，数组和map，验证元素个数                                         |
| min=10                        | 与max相反                                                                                                                                                                |
| eq=10                         | 对于字符串和数字，eq将确保该值等于给定的参数。对于切片，数组和map，验证元素个数                                                                                          |
| ne=10                         | 与eq相反                                                                                                                                                                 |
| oneof=red green (oneof=5 7 9) | 对于字符串，整数和uint，oneof将确保该值是参数中的值之一。参数应该是由空格分隔的值列表。值可以是字符串或数字                                                              |
| gt=10                         | 大于，对于数字，确保该值大于给定的参数。对于字符串，它会检查字符串长度是否大于该字符数。对于切片，数组和map，它会验证元素个数。对于time.Time确保时间值大于time.Now.UTC() |
| gte=10                        | 大于等于，对于time.Time确保时间值大于或等于time.Now.UTC()                                                                                                                |
| lt=10                         | 小于，对于time.Time确保时间值小于time.Now.UTC()                                                                                                                          |
| lte=10                        | 小于等于，对于time.Time确保时间值小于等于time.Now.UTC()                                                                                                                  |
| unique                        | 对于数组和切片，unique将确保没有重复项。对于map，unique将确保没有重复值                                                                                                  |
| alpha                         | 验证字符串值是否仅包含ASCII字母字符                                                                                                                                      |
| alphanum                      | 验证字符串值是否仅包含ASCII字母数字字符                                                                                                                                  |
| alphaunicode                  | 验证字符串值是否仅包含unicode字符                                                                                                                                        |
| alphanumunicode               | 验证字符串值是否仅包含unicode字母数字字符                                                                                                                                |
| numeric                       | 验证字符串值是否包含基本数值。基本排除指数等...对于整数或浮点数，返回true                                                                                                |
| hexadecimal                   | 验证字符串值是否包含有效的十六进制                                                                                                                                       |
| hexcolor                      | 验证字符串值包含有效的十六进制颜色，包括＃标签                                                                                                                           |
| rgb                           | 验证字符串值是否包含有效的rgb颜色                                                                                                                                        |
| rgba                          | 验证字符串值是否包含有效的rgba颜色                                                                                                                                       |
| hsl                           | 验证字符串值是否包含有效的hsl颜色                                                                                                                                        |
| hsla                          | 验证字符串值是否包含有效的hsla颜色                                                                                                                                       |
| email                         | 验证字符串值包含有效的电子邮件                                                                                                                                           |
| file                          | 验证字符串值是否包含有效的文件路径，并且该文件存在于计算机上                                                                                                             |
| url                           | 验证字符串值是否包含有效的url接受golang请求uri接受的任何url，但必须包含一个模式，例如http://或rtmp://                                                                    |
| uri                           | 验证了字符串值包含有效的uri。接受uri接受的golang请求的任何uri                                                                                                            |
| base64                        | 验证字符串值是否包含有效的base64值                                                                                                                                       |
| base64url                     | 根据RFC4648规范验证字符串值是否包含有效的base64 URL安全值                                                                                                                |
| btc_addr                      | 验证字符串值是否包含有效的比特币地址                                                                                                                                     |
| btc_addr_bech32               | 验证了字符串值包含bip-0173定义的有效比特币Bech32地址                                                                                                                     |
| eth_addr                      | 验证字符串值是否包含有效的以太坊地址                                                                                                                                     |
| contains=@                    | 验证字符串值是否包含子字符串值                                                                                                                                           |
| containsany=!@#?              | 验证字符串值是否包含子字符串值中的任何Unicode code points                                                                                                                |
| containsrune=@                | 验证字符串值是否包含提供的符文值                                                                                                                                         |
| excludes=@                    | 验证字符串值不包含子字符串值                                                                                                                                             |
| excludesall=!@#?              | 将验证字符串值在子字符串值中是否包含任何Unicode code points                                                                                                              |
| excludesrune=@                | 验证字符串值是否包含提供的符文值                                                                                                                                         |
| isbn                          | 验证字符串值是否包含有效的isbn10或isbn13值。                                                                                                                             |
| isbn10                        | 验证字符串值是否包含有效的isbn10值。                                                                                                                                     |
| isbn13                        | 验证字符串值是否包含有效的isbn13值。                                                                                                                                     |
| uuid                          | 验证字符串值是否包含有效的UUID。                                                                                                                                         |
| uuid3                         | 验证字符串值是否包含有效的版本3 UUID。                                                                                                                                   |
| uuid4                         | 验证字符串值是否包含有效的版本4 UUID。                                                                                                                                   |
| uuid5                         | 验证字符串值是否包含有效的版本5 UUID。                                                                                                                                   |
| ascii                         | 验证字符串值是否仅包含ASCII字符。注意：如果字符串为空，则验证为true                                                                                                      |
| printascii                    | 验证字符串值是否仅包含可打印的ASCII字符。注意：如果字符串为空，则验证为true。                                                                                            |
| multibyte                     | 验证字符串值是否包含一个或多个多字节字符。注意：如果字符串为空，则验证为true                                                                                             |
| datauri                       | 验证字符串值是否包含有效的DataURI。注意：这也将验证数据部分是否有效base64                                                                                                |
| latitude                      | 验证字符串值是否包含有效的纬度。                                                                                                                                         |
| longitude                     | 验证字符串值是否包含有效经度。                                                                                                                                           |
| ssn                           | 验证字符串值是否包含有效的美国社会安全号码。                                                                                                                             |
| ip                            | 验证字符串值是否包含有效的IP地址                                                                                                                                         |
| ipv4                          | 验证字符串值是否包含有效的v4 IP地址                                                                                                                                      |
| ipv6                          | 验证字符串值是否包含有效的v6 IP地址                                                                                                                                      |
| cidr                          | 验证字符串值是否包含有效的CIDR地址                                                                                                                                       |
| cidrv4                        | 验证字符串值是否包含有效的v4 CIDR地址                                                                                                                                    |
| cidrv5                        | 验证字符串值是否包含有效的v5 CIDR地址                                                                                                                                    |
| tcp_addr                      | 验证字符串值是否包含有效的可解析TCP地址                                                                                                                                  |
| tcp4_addr                     | 验证字符串值是否包含有效的可解析v4 TCP地址                                                                                                                               |
| tcp6_addr                     | 验证字符串值是否包含有效的可解析v6 TCP地址                                                                                                                               |
| udp_addr                      | 验证字符串值是否包含有效的可解析UDP地址                                                                                                                                  |
| udp4_addr                     | 验证字符串值是否包含有效的可解析v4 UDP地址                                                                                                                               |
| udp6_addr                     | 验证字符串值是否包含有效的可解析v6 UDP地址                                                                                                                               |
| ip_addr                       | 验证字符串值是否包含有效的可解析IP地址                                                                                                                                   |
| ip4_addr                      | 验证字符串值是否包含有效的可解析v4 IP地址                                                                                                                                |
| ip6_addr                      | 验证字符串值是否包含有效的可解析v6 IP地址                                                                                                                                |
| unix_addr                     | 验证字符串值是否包含有效的Unix地址                                                                                                                                       |
| mac                           | 验证字符串值是否包含有效的MAC地址。注意：有关可接受的格式和类型，请参阅Go的ParseMAC: http://golang.org/src/net/mac.go?s=866:918#L29                                      |
| hostname                      | 根据RFC 952 https://tools.ietf.org/html/rfc952验证字符串值是否为有效主机名                                                                                               |
| fqdn                          | 验证字符串值是否包含有效的FQDN (完全合格的有效域名)，Full Qualified Domain Name (FQDN)                                                                                   |
| html                          | 验证字符串值是否为HTML元素标记，包括https://developer.mozilla.org/en-US/docs/Web/HTML/Element中描述的标记。                                                              |
| html_encoded                  | 验证字符串值是十进制或十六进制格式的正确字符引用                                                                                                                         |
| url_encoded                   | 这验证了根据https://tools.ietf.org/html/rfc3986#section-2.1对字符串值进行了百分比编码（URL编码）                                                                         |

## HTTP 常用状态码
| <div style="width:48px">状态码</div> | 状态码英文名称                  | 中文描述                                                                                                             |
| :----------------------------------- | :------------------------------ | :------------------------------------------------------------------------------------------------------------------- |
| 100                                  | Continue                        | (继续) 请求者应当继续提出请求。 服务器返回此代码表示已收到请求的第一部分，正在等待其余部分                           |
| 101                                  | Switching Protocols             | (切换协议) 请求者已要求服务器切换协议，服务器已确认并准备切换                                                        |
|                                      |                                 |                                                                                                                      |
| 200                                  | OK                              | (成功) 服务器已成功处理了请求。 通常，这表示服务器提供了请求的网页                                                   |
| 201                                  | Created                         | (已创建) 请求成功并且服务器创建了新的资源                                                                            |
| 202                                  | Accepted                        | (已接受) 服务器已接受请求，但尚未处理                                                                                |
| 203                                  | Non-Authoritative Information   | (非授权信息) 服务器已成功处理了请求，但返回的信息可能来自另一来源                                                    |
| 204                                  | No Content                      | (无内容) 服务器成功处理了请求，但没有返回任何内容                                                                    |
| 205                                  | Reset Content                   | (重置内容) 服务器成功处理了请求，但没有返回任何内容                                                                  |
| 206                                  | Partial Content                 | (部分内容) 服务器成功处理了部分 GET 请求                                                                             |
|                                      |                                 |                                                                                                                      |
| 300                                  | Multiple Choices                | (多种选择) 针对请求，服务器可执行多种操作。 服务器可根据请求者 (user agent) 选择一项操作，或提供操作列表供请求者选择 |
| 301                                  | Moved Permanently               | (永久移动) 请求的网页已永久移动到新位置。 服务器返回此响应(对 GET 或 HEAD 请求的响应)时，会自动将请求者转到新位置    |
| 302                                  | Found                           | (临时移动) 服务器目前从不同位置的网页响应请求，但请求者应继续使用原有位置来进行以后的请求                            |
| 303                                  | See Other                       | (查看其他位置) 请求者应当对不同的位置使用单独的 GET 请求来检索响应时，服务器返回此代码                               |
| 304                                  | Not Modified                    | (未修改) 自从上次请求后，请求的网页未修改过。 服务器返回此响应时，不会返回网页内容                                   |
| 305                                  | Use Proxy                       | (使用代理) 请求者只能使用代理访问请求的网页。 如果服务器返回此响应，还表示请求者应使用代理                           |
| 307                                  | Temporary Redirect              | (临时重定向) 服务器目前从不同位置的网页响应请求，但请求者应继续使用原有位置来进行以后的请求                          |
|                                      |                                 |                                                                                                                      |
| 400                                  | Bad Request                     | (错误请求) 服务器不理解请求的语法                                                                                    |
| 401                                  | Unauthorized                    | (未授权) 请求要求身份验证。 对于需要登录的网页，服务器可能返回此响应                                                 |
| 403                                  | Forbidden                       | (禁止) 服务器拒绝请求                                                                                                |
| 404                                  | Not Found                       | (未找到) 服务器找不到请求的网页                                                                                      |
| 405                                  | Method Not Allowed              | (方法禁用) 禁用请求中指定的方法                                                                                      |
| 406                                  | Not Acceptable                  | (不接受) 无法使用请求的内容特性响应请求的网页                                                                        |
| 407                                  | Proxy Authentication Required   | (需要代理授权) 此状态代码与 401(未授权)类似，但指定请求者应当授权使用代理                                            |
| 408                                  | Request Time-out                | (请求超时) 服务器等候请求时发生超时                                                                                  |
| 409                                  | Conflict                        | (冲突) 服务器在完成请求时发生冲突。 服务器必须在响应中包含有关冲突的信息                                             |
| 410                                  | Gone                            | (已删除) 如果请求的资源已永久删除，服务器就会返回此响应                                                              |
| 411                                  | Length Required                 | (需要有效长度) 服务器不接受不含有效内容长度标头字段的请求                                                            |
| 412                                  | Precondition Failed             | (未满足前提条件) 服务器未满足请求者在请求中设置的其中一个前提条件                                                    |
| 413                                  | Request Entity Too Large        | (请求实体过大) 服务器无法处理请求，因为请求实体过大，超出服务器的处理能力                                            |
| 414                                  | Request-URI Too Large           | (请求的 URI 过长) 请求的 URI(通常为网址)过长，服务器无法处理                                                         |
| 415                                  | Unsupported Media Type          | (不支持的媒体类型) 请求的格式不受请求页面的支持                                                                      |
| 416                                  | Requested range not satisfiable | (请求范围不符合要求) 如果页面无法提供请求的范围，则服务器会返回此状态代码                                            |
| 417                                  | Expectation Failed              | (未满足期望值) 服务器未满足"期望"请求标头字段的要求                                                                  |
|                                      |                                 |                                                                                                                      |
| 500                                  | Internal Server Error           | (服务器内部错误) 服务器遇到错误，无法完成请求                                                                        |
| 501                                  | Not Implemented                 | (尚未实施) 服务器不具备完成请求的功能。 例如，服务器无法识别请求方法时可能会返回此代码                               |
| 502                                  | Bad Gateway                     | (错误网关) 服务器作为网关或代理，从上游服务器收到无效响应                                                            |
| 503                                  | Service Unavailable             | (服务不可用) 服务器目前无法使用(由于超载或停机维护)。 通常，这只是暂时状态                                           |
| 504                                  | Gateway Time-out                | (网关超时) 服务器作为网关或代理，但是没有及时从上游服务器收到请求                                                    |
| 505                                  | HTTP Version not supported      | (HTTP 版本不受支持) 服务器不支持请求中所用的 HTTP 协议版本                                                           |