package grade

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type studentsHandler struct{}

func RegisterHandlers() {
	handler := new(studentsHandler)
	http.Handle("/students", handler)
	http.Handle("/students/", handler) //这个后面还可以带参数
}

// /students
// /students/{id}
// /students/{id}/grades
func (s studentsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 提取参数，这里不做校验
	pathSegments := strings.Split(r.URL.Path, "/")

	// 根据拆分的长度来匹配
	switch len(pathSegments) {
	case 2:
		s.getAll(w, r)
	case 3:
		id, err := strconv.Atoi(pathSegments[2])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		s.getOne(w, r, id)
	case 4:
		id, err := strconv.Atoi(pathSegments[2])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		s.addGrade(w, r, id)
	default:
		w.WriteHeader(http.StatusNotFound)
		return
	}
}

// 添加用户
func (s studentsHandler) addGrade(w http.ResponseWriter, r *http.Request, id int) {
	studentsMutex.Lock()
	defer studentsMutex.Unlock()

	student, _ := s.getStudentByID(id)
	if student != nil {
		data, _ := s.toJSON(student)
		w.Header().Add("content-Type", "application/json")
		w.Write(data)
		log.Println("当前用户已经存在,")
		return
	}

	// 定义一个解码器
	var newStudent Student
	desc := json.NewDecoder(r.Body)
	err := desc.Decode(&newStudent)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("序列化json出错", err)
		return
	}

	students = append(students, newStudent)
	w.WriteHeader(http.StatusCreated)
	data, err := s.toJSON(newStudent)
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Add("content-Type", "application/json")
	w.Write(data)

	return
}

// 获取所有用户数据
func (s studentsHandler) getAll(w http.ResponseWriter, r *http.Request) {
	studentsMutex.Lock()
	defer studentsMutex.Unlock()

	data, err := s.toJSON(students)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("序列化json出错", err)
		return
	}
	w.Header().Add("content-Type", "application/json")
	w.Write(data)

	return
}

// 根据ID查询用户
func (s studentsHandler) getStudentByID(id int) (*Student, error) {
	for _, student := range students {
		if student.ID == id {
			return &student, nil
		}
	}

	return nil, fmt.Errorf("Student ID = %v not found", id)
}

// 查询一个用户数据
func (s studentsHandler) getOne(w http.ResponseWriter, r *http.Request, id int) {
	studentsMutex.Lock()
	defer studentsMutex.Unlock()

	student, err := s.getStudentByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("获取数据失败,", err)
		return
	}

	data, err := s.toJSON(student)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("序列化json出错", err)
		return
	}
	w.Header().Add("content-Type", "application/json")
	w.Write(data)

	return
}

// 序列化数据
func (s studentsHandler) toJSON(obj interface{}) ([]byte, error) {
	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	if err := enc.Encode(obj); err != nil {
		return nil, fmt.Errorf("Failed to serialize")
	}
	return b.Bytes(), nil
}
