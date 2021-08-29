package reqo

// GetUser GET "/user" request object
type GetUser struct {
	Username string `json:"username"`
	ID       uint   `json:"id"`
}

// PutUser PUT "/user" request object
type PutUser struct {
	Gender   int64  `json:"gender"`
	Nickname string `json:"nickname"`
	Mail     string `json:"mail"`
}

// PostUser POST "/user" request object
type PostUser struct {
	Username string `json:"username"`
	Passwd   string `json:"passwd"`
	Gender   int64  `json:"gender"`
	Nickname string `json:"nickname"`
	Mail     string `json:"mail"`
}

// POST "/login" 接收客户端的登录请求接口
type PostLogin struct {
	Username string `json:"username"`
	Passwd   string `json:"passwd"`
}
