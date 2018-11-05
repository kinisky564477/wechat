package wx

// 微信事件, 暂时实现部分
// 其余的如认证，年审等暂未体现
const (
	// EventSubscribe 关注事件
	EventSubscribe = "subscribe"

	// EventUnsubscribe 取消关注
	EventUnsubscribe = "unsubscribe"

	// EventScan 扫码
	EventScan = "SCAN"

	// EventScanSubscribe 未关注扫码
	// 注: 官方事件类型为 subscribe 与 关注重复
	EventScanSubscribe = "subscribe"

	// EventLoction 上报位置
	EventLoction = "LOCATION"

	// EventClick 菜单点击
	EventClick = "CLICK"
)
