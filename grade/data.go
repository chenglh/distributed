package grade

func init() {
	students = []Student{
		{
			ID:        1,
			FirstName: "张",
			LastName:  "小三",
			Grades: []Grade{
				{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 85,
				},
				{
					Title: "Final Exam",
					Type:  GradeExam,
					Score: 94,
				},
				{
					Title: "Test",
					Type:  GradeTest,
					Score: 90,
				},
			},
		},
		{
			ID:        2,
			FirstName: "李",
			LastName:  "大明",
			Grades: []Grade{
				{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 80,
				},
				{
					Title: "Final Exam",
					Type:  GradeExam,
					Score: 78,
				},
				{
					Title: "Test",
					Type:  GradeTest,
					Score: 96,
				},
			},
		},
		{
			ID:        3,
			FirstName: "赵",
			LastName:  "得华",
			Grades: []Grade{
				{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 88,
				},
				{
					Title: "Final Exam",
					Type:  GradeExam,
					Score: 90,
				},
			},
		},
	}
}
