package registry

// 定义
type Registration struct {
	ServiceName ServiceName
	ServiceURL  string
}

// string 类型别名
type ServiceName string

// 服务名称常量
const (
	LogService = ServiceName("LogService")
)
