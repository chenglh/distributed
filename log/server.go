package log

// 因为需要使用到标准库中的 log，使用别名
import (
	"io"
	stdlog "log"
	"net/http"
	"os"
)

var log *stdlog.Logger

// string的别名
type fileLog string

// 这里是实现 fileLog变量中绑定 Write方法
func (fl fileLog) Write(data []byte) (int, error) {
	// 0600 是读写权限
	fWrite, err := os.OpenFile(string(fl), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		return 0, err
	}
	defer fWrite.Close()

	return fWrite.Write(data)
}

// 这里使用go的标准库日志New方法，继承和实现标准库中的日志类
func Run(destination string) {
	// 第一个参数是：类型转换, string(目录文件名，路径)
	// 第二个参数是：日志数据前缀
	// 第三个参数是：日志长格式时间值(标准库中定义的常量)
	log = stdlog.New(fileLog(destination), "go ", stdlog.LstdFlags)
}

func RegisterHandlers() {
	// http的handle模块（一般也称为钩子模块），通过高级语言的匿名函数很容易实现这种内嵌功能的handle
	// 就是平时使用的路由功能，长驻内存
	http.HandleFunc("/log", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		// 针对Post方法写日志
		case http.MethodPost:
			//msg, err := ioutil.ReadAll(r.Body)这个方法要被删除了
			msg, err := io.ReadAll(r.Body)
			if err != nil || len(msg) == 0 {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			//对于请求 /log 的路径上的 post请求，把请求参数写入日志里边
			write(string(msg))
		// 其他的方法不允许访问
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	})
}

func write(message string) {
	log.Printf("%v\n", message)
}
