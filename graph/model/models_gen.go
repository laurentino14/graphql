// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Course struct {
	ID            string        `json:"id"`
	Title         string        `json:"title"`
	Slug          string        `json:"slug"`
	Description   *string       `json:"description"`
	CreatedAt     string        `json:"created_at"`
	UpdatedAt     *string       `json:"updated_at"`
	AuthorID      []*User       `json:"author_id"`
	EnrollmentsID []*Enrollment `json:"enrollments_id"`
	ModulesID     []*Module     `json:"modules_id"`
}

type Enrollment struct {
	ID       string    `json:"id"`
	UserID   string    `json:"user_id"`
	CourseID string    `json:"course_id"`
	User     []*User   `json:"user"`
	Course   []*Course `json:"course"`
}

type Lesson struct {
	ID     string    `json:"id"`
	Title  string    `json:"title"`
	Slug   string    `json:"slug"`
	Link   *string   `json:"link"`
	Module []*Module `json:"module"`
	Course []*Course `json:"course"`
}

type Module struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Slug        string    `json:"slug"`
	Lessons     []*Lesson `json:"lessons"`
	Course      []*Course `json:"course"`
}

type NewCourse struct {
	ID            string    `json:"id"`
	Title         string    `json:"title"`
	Slug          string    `json:"slug"`
	Description   *string   `json:"description"`
	CreatedAt     string    `json:"created_at"`
	UpdatedAt     *string   `json:"updated_at"`
	AuthorID      []*string `json:"author_id"`
	EnrollmentsID []*string `json:"enrollments_id"`
	ModulesID     []*string `json:"modules_id"`
}

type NewEnrollment struct {
	ID       string    `json:"id"`
	UserID   []*string `json:"user_id"`
	CourseID []*string `json:"course_id"`
}

type NewLesson struct {
	ID       string    `json:"id"`
	Title    string    `json:"title"`
	Slug     string    `json:"slug"`
	Link     string    `json:"link"`
	ModuleID []*string `json:"module_id"`
	CourseID []*string `json:"course_id"`
}

type NewModule struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Slug        string    `json:"slug"`
	LessonsID   []*string `json:"lessons_id"`
	CourseID    []*string `json:"course_id"`
}

type NewUser struct {
	ID        *string `json:"id"`
	Firstname string  `json:"firstname"`
	Lastname  string  `json:"lastname"`
	Email     string  `json:"email"`
	Password  string  `json:"password"`
	Cellphone string  `json:"cellphone"`
	BirthDate *string `json:"birth_date"`
	TokenUser *string `json:"token_user"`
}

type User struct {
	ID          string        `json:"id"`
	Firstname   string        `json:"firstname"`
	Lastname    string        `json:"lastname"`
	Email       string        `json:"email"`
	Password    string        `json:"password"`
	Cellphone   string        `json:"cellphone"`
	BirthDate   *string       `json:"birth_date"`
	TokenUser   *string       `json:"token_user"`
	AuthorOf    []*Course     `json:"author_of"`
	Enrollments []*Enrollment `json:"enrollments"`
}
