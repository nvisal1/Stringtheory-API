package shared

type SecureUser struct {
	Id string
	Username string
	Name string
	Email string
}

type User struct {
	SecureUser
	Password string
}

type Course struct {
	Id string
	Name string
	Description string
}

type Lesson struct {
	Id string
	Name string
	Order int
	Description string
	CourseId string
	HasNext bool
}

type Exercise struct {
	Id string
	Name string
	Order int
	Notes []string
	Description string
	LessonId string
	HasNext bool
}




