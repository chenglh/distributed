package grade

import (
	"sync"
)

type GradeType string

type Student struct {
	ID        int
	FirstName string
	LastName  string
	Grades    []Grade //考试内容类型
}

// 计算平均成绩(3次考试)
func (s *Student) Average() float32 {
	var result float32
	for _, grade := range s.Grades {
		result += grade.Score
	}

	return result / 3
}

type Students []Student

// 装载测试数据
var (
	students      Students
	studentsMutex sync.Mutex //加锁，并发安全
)

type Grade struct {
	Title string
	Type  GradeType
	Score float32
}

const (
	GradeQuiz = GradeType("Quiz") //竞赛
	GradeTest = GradeType("Test") //测试
	GradeExam = GradeType("Exam") //大考
)
