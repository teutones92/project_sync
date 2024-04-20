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
	StatusCodeMessage string `json:"status_code_name"`
}

// User represent the structure of a user
type User struct {
	UserID       int    `json:"user_id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
	Password     string `json:"password"`
	UserAvatar   string `json:"user_avatar_path"`
}

// UserRole represent the structure of a user role
type UserRole struct {
	RoleID          int    `json:"role_id"`
	RoleName        string `json:"role_name"`
	RoleDescription string `json:"role_description"`
}

// Project represent the structure of a project
type Project struct {
	ProjectID     int    `json:"project_id"`
	ProjectName   string `json:"project_name"`
	Description   string `json:"description"`
	StartDate     string `json:"start_date"`
	EndDate       string `json:"end_date"`
	ProjectLeadID int    `json:"project_lead_id"`
}

// Task represent the structure of a task
type Task struct {
	TaskID       int    `json:"task_id"`
	ProjectID    int    `json:"project_id"`
	TaskName     string `json:"task_name"`
	Description  string `json:"description"`
	Status       string `json:"status"`
	Priority     string `json:"priority"`
	AssignedUser int    `json:"assigned_user"`
	Deadline     string `json:"deadline"`
}

// TeamMember represent the structure of a team member
type TeamMember struct {
	TeamMemberID int `json:"team_member_id"`
	ProjectID    int `json:"project_id"`
	UserID       int `json:"user_id"`
	RoleID       int `json:"role_id"`
}

// Comment represent the structure of a comment
type Comment struct {
	CommentID   int    `json:"comment_id"`
	ProjectID   int    `json:"project_id"`
	TaskID      int    `json:"task_id"`
	UserID      int    `json:"user_id"`
	Timestamp   string `json:"timestamp"`
	CommentText string `json:"comment_text"`
}

// Session represent the structure of a session
type Session struct {
	SessionID             int    `json:"session_id"`
	UserID                int    `json:"user_id"`
	Token                 string `json:"token"`
	LastActivityTimestamp string `json:"last_activity_timestamp"`
	ExpirationMinutes     int    `json:"expiration_minutes"`
}
