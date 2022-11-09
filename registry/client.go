package registry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// 服务注册函数，主要是给registryservice发送一个post请求。
func RegisterService(r Registration) error {
	// 组装数据
	buffer := new(bytes.Buffer)
	enc := json.NewEncoder(buffer)
	err := enc.Encode(r) // 对r进行编码
	if err != nil {
		return err
	}

	// 发送请求,第三个参数是：io.Reader 类型
	res, err := http.Post(ServicesURL, "application/json", buffer)
	if err != nil {
		return err
	}

	// 获取请求结果状态码
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to register service. Registry service "+"response with code %v", res.StatusCode)
	}

	return nil
}

// 注销服务函数，发送一个 delete 请求的方法，http包中没有该方法，需要自己创建一个 delete请求
func ShutdownService(url string) error {
	// 创建一个 http请求
	req, err := http.NewRequest(http.MethodDelete, ServicesURL, bytes.NewBuffer([]byte(url)))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "text/plain")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to deregister service. %v, with code %v", url, res.StatusCode)
	}

	return nil
}
