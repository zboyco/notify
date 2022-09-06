### 1. wxpusher回调

1. route definition

- Url: /v0/wxpusher/callback
- Method: POST
- Request: `WxPusherCallbackRequest`
- Response: `-`

2. request definition



```golang
type WxPusherCallbackRequest struct {
	Action string `json:"action"` // 操作
	Data WxPusherCallbackRequestData `json:"data"` // 数据
}

type WxPusherCallbackRequestData struct {
	AppID int64 `json:"appId"` // 应用ID
	AppName string `json:"appName"` // 应用名称
	UID string `json:"uid"` // 用户ID
	Source string `json:"source,omitempty"` // 订阅来源
	Time int64 `json:"time"` // 消息时间
	Extra string `json:"extra,omitempty"` // 额外数据
	Content string `json:"content,omitempty"` // 消息内容
}
```


3. response definition


### 2. 创建通知

1. route definition

- Url: /v0/notifies
- Method: POST
- Request: `NotifyCreateRequest`
- Response: `-`

2. request definition



```golang
type NotifyCreateRequest struct {
Auth
Notify
}

type Auth struct {
	Token string `header:"token"` // 请求Token
}

type Notify struct {
	ID uint `json:"id,optional"` // 消息ID
	WechatUserID string `json:"wechatUserID,optional"` // 微信用户ID
	Topic string `json:"topic,optional"` // 主题
	Title string `json:"title"` // 标题
	Content string `json:"content,optional"` // 内容
	Loop bool `json:"loop"` // 是否循环
	StartAt int64 `json:"startAt"` // 开始时间
	EndAt int64 `json:"endAt,optional"` // 结束时间
	Spec string `json:"spec,optional"` // Cron表达式（循环有效）
	NotifyCount int64 `json:"notifyCount,optional"` // 通知次数
	LastNotifyAt int64 `json:"lastNotifyAt,optional"` // 最后通知时间
}
```


3. response definition


### 3. 获取通知列表

1. route definition

- Url: /v0/notifies
- Method: GET
- Request: `NotifyListRequest`
- Response: `NotifyListResponse`

2. request definition



```golang
type NotifyListRequest struct {
Auth
Pager
}

type Auth struct {
	Token string `header:"token"` // 请求Token
}

type Pager struct {
	Limit int `form:"limit,default=10"` // 每页数量
	Offset int `form:"offset,default=0"` // 偏移量
}
```


3. response definition



```golang
type NotifyListResponse struct {
	Total int64 `json:"total"` // 总数
	Data []*Notify `json:"data"` // 数据
}
```

### 4. 获取通知详情

1. route definition

- Url: /v0/notifies/:notifyID
- Method: GET
- Request: `NotifyByIDRequest`
- Response: `Notify`

2. request definition



```golang
type NotifyByIDRequest struct {
Auth
	NotifyID uint `path:"notifyID"` // 消息ID
}

type Auth struct {
	Token string `header:"token"` // 请求Token
}
```


3. response definition



```golang
type Notify struct {
	ID uint `json:"id,optional"` // 消息ID
	WechatUserID string `json:"wechatUserID,optional"` // 微信用户ID
	Topic string `json:"topic,optional"` // 主题
	Title string `json:"title"` // 标题
	Content string `json:"content,optional"` // 内容
	Loop bool `json:"loop"` // 是否循环
	StartAt int64 `json:"startAt"` // 开始时间
	EndAt int64 `json:"endAt,optional"` // 结束时间
	Spec string `json:"spec,optional"` // Cron表达式（循环有效）
	NotifyCount int64 `json:"notifyCount,optional"` // 通知次数
	LastNotifyAt int64 `json:"lastNotifyAt,optional"` // 最后通知时间
}
```

### 5. 删除通知

1. route definition

- Url: /v0/notifies/:notifyID
- Method: DELETE
- Request: `NotifyByIDRequest`
- Response: `-`

2. request definition



```golang
type NotifyByIDRequest struct {
Auth
	NotifyID uint `path:"notifyID"` // 消息ID
}

type Auth struct {
	Token string `header:"token"` // 请求Token
}
```


3. response definition


