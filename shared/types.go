package shared

type SecureUser struct {
	Username string
	Name string
	Email string
}

type User struct {
	Username string
	Name string
	Email string
	Password string
}

type Course struct {
	ID string
	Name string
	Description string
	LessonsURI string
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




