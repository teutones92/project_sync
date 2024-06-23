package models

type AllData struct {
	User       User
	UserRole   UserRole
	Project    Project
	Task       Task
	TeamMember TeamMember
	Comment    Comment
	Session    Session
	Status     StatusCode
}

// StatusCode represent the structure of a status code
type StatusCode struct {
	StatusCode        int    `json:"status_code"`
	StatusCodeMessage string `json:"status_code_message"`
}

// Error implements error.
func (s StatusCode) Error() string {
	panic("unimplemented")
}

// User represent the structure of a user
type User struct {
	ID               int    `json:"id"`
	Username         string `json:"username"`
	Email            string `json:"email"`
	DoB              string `json:"DoB"`
	PhoneNumber      string `json:"phone_number"`
	CountryCode      string `json:"country_code"`
	CountryPhoneCode string `json:"country_phone_code"`
	LangCode         string `json:"lang_code"`
	PasswordHash     string `json:"password_hash"`
	Password         string `json:"password"`
	UserAvatar       string `json:"user_avatar_path"`
	DarkMode         *bool  `json:"dark_mode"`
}

// UserRole represent the structure of a user role
type UserRole struct {
	ID              int    `json:"id"`
	RoleName        string `json:"role_name"`
	RoleDescription string `json:"role_description"`
}

// Project represent the structure of a project
type Project struct {
	ID            int    `json:"id"`
	ProjectName   string `json:"project_name"`
	Description   string `json:"description"`
	StartDate     string `json:"start_date"`
	EndDate       string `json:"end_date"`
	ProjectLeadID int    `json:"project_lead_id"`
	ImagePath     string `json:"image_path"`
}

// ProjectTag represent the structure of a project tag
type ProjectTag struct {
	ID        int    `json:"id"`
	ProjectID *int   `json:"project_id"`
	TagName   string `json:"tag_name"`
}

// Priority represent the structure of a project priority
type Priority struct {
	ID                  int    `json:"id"`
	PriorityName        string `json:"priority_name"`
	PriorityDescription string `json:"priority_description"`
	PriorityRGBColor    string `json:"priority_rgb_color"`
}

// Task represent the structure of a task
type Task struct {
	ID           int      `json:"id"`
	ProjectID    int      `json:"project_id"`
	TaskName     string   `json:"task_name"`
	Description  string   `json:"description"`
	StatusID     int      `json:"status_id"`
	Priority     string   `json:"priority"`
	AssignedUser int      `json:"assigned_user"`
	Deadline     string   `json:"deadline"`
	ImagePath    []string `json:"image_path"`
}

// TeamMember represent the structure of a team member
type TeamMember struct {
	ID        int `json:"id"`
	ProjectID int `json:"project_id"`
	UserID    int `json:"user_id"`
	RoleID    int `json:"role_id"`
}

// Comment represent the structure of a comment
type Comment struct {
	ID          int    `json:"id"`
	ProjectID   int    `json:"project_id"`
	TaskID      int    `json:"task_id"`
	UserID      int    `json:"user_id"`
	Timestamp   string `json:"timestamp"`
	CommentText string `json:"comment_text"`
}

// Session represent the structure of a session
type Session struct {
	ID                    int    `json:"id"`
	UserID                int    `json:"user_id"`
	Token                 string `json:"token"`
	LastActivityTimestamp string `json:"last_activity_timestamp"`
	ExpirationMinutes     int    `json:"expiration_minutes"`
}

// TaskStatus represent the structure of a task status
type TaskStatus struct {
	ID                int    `json:"id"`
	ProjectId         int    `json:"project_id"`
	UserID            int    `json:"user_id"`
	StatusName        string `json:"status_name"`
	StatusDescription string `json:"status_description"`
}

// UserContact represent the structure of a user contact
type UserContact struct {
	ID           int    `json:"id"`
	UserID       int    `json:"user_id"`
	ContactName  string `json:"contact_name"`
	ContactEmail string `json:"contact_email"`
}
