package registry

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

/** 用于查询哪些服务注册成功 */
// 服务的相关参数
const ServerPort = ":3000"
const ServicesURL = "http://localhost" + ServerPort + "/services"

// 私有结构体
type registry struct {
	registration []Registration
	mutex        *sync.Mutex //保证并发时是安全的(读写时)
}

// 实例化
var reg = registry{
	registration: make([]Registration, 0),
	mutex:        new(sync.Mutex),
}

// 添加到集合中
func (r *registry) add(reg Registration) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.registration = append(r.registration, reg)

	return nil
}

// 删除指定服务
func (r *registry) remove(url string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for idx, _ := range r.registration {
		if r.registration[idx].ServiceURL == url {
			r.mutex.Lock()
			reg.registration = append(reg.registration[:idx], reg.registration[idx+1:]...)
			r.mutex.Unlock()

			return nil
		}
	}

	return fmt.Errorf("Service at URL %s not found", url)
}

// 主要用于实现 ServeHTTP()
type RegistrationService struct{}

func (s *RegistrationService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("接收到http请求")

	switch r.Method {
	case http.MethodPost:
		// 解析json请求参数
		desc := json.NewDecoder(r.Body)

		// 也可以 json.NewDecoder(r.Body).Decode(&registration) 连惯操作
		var registration Registration
		if err := desc.Decode(&registration); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// 打印解析出来的内容
		log.Printf("Adding service:%v with URL:%s\n", registration.ServiceName, registration.ServiceURL)

		// 注册到服务集合中,如果失败，返回错误请求
		if err := reg.add(registration); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	case http.MethodDelete:
		payload, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		log.Printf("Removing service at URL:%s\n", string(payload))
		err = reg.remove(string(payload))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
