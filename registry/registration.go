package registry

// 定义注册服务内容
type Registration struct {
	ServiceName ServiceName
	ServiceURL  string

	//需要拓展，先启动依赖项
	RequiredServices []ServiceName //如 grade 需要 log服务先启动
	ServiceUpdateURL string        //服务的更新(增加或减少)
}

// string 类型别名
type ServiceName string

// 服务名称常量
const (
	GradeService = ServiceName("GradeService")
	LogService   = ServiceName("LogService")
)
