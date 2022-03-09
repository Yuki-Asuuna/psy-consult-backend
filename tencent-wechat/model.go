package tencent_wechat

//属性	        类型	          说明
//openid	    string	     用户唯一标识
//session_key	string	     会话密钥
//unionid	    string	     用户在开放平台的唯一标识符，若当前小程序已绑定到微信开放平台帐号下会返回，详见 UnionID 机制说明。
//errcode	    number	     错误码
//errmsg	    string	     错误信息

type WeChatLoginResponse struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
	ErrCode    int64  `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

//errcode 的合法值
//
//值	说明	最低版本
//-1	系统繁忙，此时请开发者稍候再试
//0	请求成功
//40029	code 无效
//45011	频率限制，每个用户每分钟100次
//40226	高风险等级用户，小程序登录拦截 。风险等级详见用户安全解方案
