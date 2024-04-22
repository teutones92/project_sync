package apis

import (
	auth "app/authentication"
	comments "app/db_connection/tables/comments_crud"
	project "app/db_connection/tables/project_crud"
	user_role "app/db_connection/tables/roles_crud"
	task "app/db_connection/tables/task_crud"
	taskStatus "app/db_connection/tables/task_status_crud"
	team_member "app/db_connection/tables/team_member_crud"
	user_contacts "app/db_connection/tables/user_contacts_crud"
	user "app/db_connection/tables/user_crud"
	"log"

	"net/http"
)

var Host = "localhost"
var Port = ":8080"

// Start the server
func StartServer() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, Welcome to Project Sync!"))
	})
	log.Printf("Server running on %s%s", Host, Port)
	mux := http.NewServeMux()
	authApis(mux)
	usersApis(mux)
	userRolesApis(mux)
	projectsApis(mux)
	taskStatusApis(mux)
	taskApis(mux)
	teamMembersApis(mux)
	commentsApis(mux)
	userContactsApis(mux)
	err := http.ListenAndServe(Port, nil)
	return err
}

// Auth API
func authApis(mux *http.ServeMux) {
	mux.HandleFunc("/auth/signup", auth.SignUp)
	mux.HandleFunc("/auth/login", auth.LogIn)
	mux.HandleFunc("/auth/logout", auth.LogOut)
	mux.HandleFunc("/auth/delete_account", auth.DeleteAccount)
}

// Users API
func usersApis(mux *http.ServeMux) error {
	mux.HandleFunc("/users/read", user.ReadUserAPI)
	mux.HandleFunc("/users/update", user.UpdateUserAPI)
	mux.HandleFunc("/users/change_password", user.ChangePasswordAPI)
	err := http.ListenAndServe(Port, mux)
	return err
}

// User Roles
func userRolesApis(mux *http.ServeMux) {
	// mux.HandleFunc("/user_roles/create", user_role.CreateUserRoleAPI)
	mux.HandleFunc("/user_roles/read", user_role.ReadUserRoleAPI)
	// mux.HandleFunc("/user_roles/update", user_role.UpdateUserRoleAPI)
	// mux.HandleFunc("/user_roles/delete", user_role.DeleteUserRoleAPI)
}

// Projects API
func projectsApis(mux *http.ServeMux) {
	mux.HandleFunc("/projects/create", project.CreateProject)
	mux.HandleFunc("/projects/read", project.ReadProjectByID)
	mux.HandleFunc("/projects/update", project.UpdateProject)
	mux.HandleFunc("/projects/delete", project.DeleteProject)
}

// Task Status API
func taskStatusApis(mux *http.ServeMux) {
	mux.HandleFunc("/task_status/create", taskStatus.CreateTaskStatusAPI)
	mux.HandleFunc("/task_status/read", taskStatus.ReadTaskStatusByProjectIDApi)
	mux.HandleFunc("/task_status/update", taskStatus.DeleteTaskStatusByProjectIdAndUserIdAPI)
	mux.HandleFunc("/task_status/delete", taskStatus.DeleteTaskStatusByProjectIdAndUserIdAPI)
}

// Task API
func taskApis(mux *http.ServeMux) {
	mux.HandleFunc("/tasks/create", task.CreateTaskAPI)
	mux.HandleFunc("/tasks/read", task.ReadTaskByProjectIDAndStatusIdAPI)
	mux.HandleFunc("/tasks/update", task.UpdateTaskAPI)
	mux.HandleFunc("/tasks/delete", task.DeleteTaskAPI)
}

// Team Members API
func teamMembersApis(mux *http.ServeMux) {
	mux.HandleFunc("/team_members/create", team_member.CreateTeamMemberAPI)
	mux.HandleFunc("/team_members/read", team_member.ReadTeamMembersAPI)
	mux.HandleFunc("/team_members/update", team_member.UpdateTeamMemberAPI)
	mux.HandleFunc("/team_members/delete", team_member.DeleteTeamMemberAPI)
}

// Comments API
func commentsApis(mux *http.ServeMux) {
	mux.HandleFunc("/comments/create", comments.CreateCommentAPI)
	mux.HandleFunc("/comments/read", comments.ReadCommentsAPI)
	mux.HandleFunc("/comments/update", comments.UpdateCommentAPI)
	mux.HandleFunc("/comments/delete", comments.DeleteCommentAPI)
}

// UserContacts API
func userContactsApis(mux *http.ServeMux) {
	mux.HandleFunc("/user_contacts/create", user_contacts.CreateUserContactAPI)
	mux.HandleFunc("/user_contacts/read", user_contacts.ReadUserContactByUserIdAPI)
	mux.HandleFunc("/user_contacts/update", user_contacts.UpdateUserContactAPI)
	mux.HandleFunc("/user_contacts/delete", user_contacts.DeleteUserContactAPI)
}
