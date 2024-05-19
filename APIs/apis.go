package apis

import (
	auth "app/authentication"
	comments "app/db_connection/tables/comments_crud"
	priority "app/db_connection/tables/priority_crud"
	project "app/db_connection/tables/project_crud"
	project_tag "app/db_connection/tables/project_tags_crud"
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
var apiVersion = "/v1"

// Start the server
func StartServer() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, Welcome to Project Sync!"))
	})
	log.Printf("Server running on %s%s", Host, Port)
	mux := http.NewServeMux()
	authApis(mux)
	userApis(mux)
	userRolesApis(mux)
	projectsApis(mux)
	projectTagsApis(mux)
	priorityApis(mux)
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
	mux.HandleFunc(apiVersion+"/auth/signup", auth.SignUp)
	mux.HandleFunc(apiVersion+"/auth/login", auth.LogIn)
	mux.HandleFunc(apiVersion+"/auth/logout", auth.LogOut)
	mux.HandleFunc(apiVersion+"/auth/delete_account", auth.DeleteAccount)
}

// Users API
func userApis(mux *http.ServeMux) error {
	mux.HandleFunc(apiVersion+"/user/read", user.ReadUserAPI)
	mux.HandleFunc(apiVersion+"/user/update", user.UpdateUserAPI)
	mux.HandleFunc(apiVersion+"/user/change_password", user.ChangePasswordAPI)

	err := http.ListenAndServe(Port, mux)
	return err
}

// User Roles
func userRolesApis(mux *http.ServeMux) {
	// mux.HandleFunc("/user_roles/create", user_role.CreateUserRoleAPI)
	mux.HandleFunc(apiVersion+"/user_roles/read", user_role.ReadUserRoleAPI)
	// mux.HandleFunc("/user_roles/update", user_role.UpdateUserRoleAPI)
	// mux.HandleFunc("/user_roles/delete", user_role.DeleteUserRoleAPI)
}

// Projects API
func projectsApis(mux *http.ServeMux) {
	mux.HandleFunc(apiVersion+"/projects/create", project.CreateProject)
	mux.HandleFunc(apiVersion+"/projects/read", project.ReadProjectByID)
	mux.HandleFunc(apiVersion+"/projects/update", project.UpdateProject)
	mux.HandleFunc(apiVersion+"/projects/delete", project.DeleteProject)
}

// Project Tags API
func projectTagsApis(mux *http.ServeMux) {
	mux.HandleFunc(apiVersion+"/project_tags/create", project_tag.CreateProjectTagAPI)
	mux.HandleFunc(apiVersion+"/project_tags/read", project_tag.ReadProjectTagAPI)
	// mux.HandleFunc(
	//
	// 	"/project_tags/update", project_tags_crud.UpdateProjectTagAPI)
	mux.HandleFunc(apiVersion+"/project_tags/delete", project_tag.DeleteProjectTagAPI)
}

// Priority API
func priorityApis(mux *http.ServeMux) {
	mux.HandleFunc(apiVersion+"/priority/read", priority.ReadPriorityAPI)
}

// Task Status API
func taskStatusApis(mux *http.ServeMux) {
	mux.HandleFunc(apiVersion+"/task_status/create", taskStatus.CreateTaskStatusAPI)
	mux.HandleFunc(apiVersion+"/task_status/read", taskStatus.ReadTaskStatusByProjectIDApi)
	mux.HandleFunc(apiVersion+"/task_status/update", taskStatus.DeleteTaskStatusByProjectIdAndUserIdAPI)
	mux.HandleFunc(apiVersion+"/task_status/delete", taskStatus.DeleteTaskStatusByProjectIdAndUserIdAPI)
}

// Task API
func taskApis(mux *http.ServeMux) {
	mux.HandleFunc(apiVersion+"/tasks/create", task.CreateTaskAPI)
	mux.HandleFunc(apiVersion+"/tasks/read", task.ReadTaskByProjectIDAndStatusIdAPI)
	mux.HandleFunc(apiVersion+"/tasks/update", task.UpdateTaskAPI)
	mux.HandleFunc(apiVersion+"/tasks/delete", task.DeleteTaskAPI)
}

// Team Members API
func teamMembersApis(mux *http.ServeMux) {
	mux.HandleFunc(apiVersion+"/team_members/create", team_member.CreateTeamMemberAPI)
	mux.HandleFunc(apiVersion+"/team_members/read", team_member.ReadTeamMembersAPI)
	mux.HandleFunc(apiVersion+"/team_members/update", team_member.UpdateTeamMemberAPI)
	mux.HandleFunc(apiVersion+"/team_members/delete", team_member.DeleteTeamMemberAPI)
}

// Comments API
func commentsApis(mux *http.ServeMux) {
	mux.HandleFunc(apiVersion+"/comments/create", comments.CreateCommentAPI)
	mux.HandleFunc(apiVersion+"/comments/read", comments.ReadCommentsAPI)
	mux.HandleFunc(apiVersion+"/comments/update", comments.UpdateCommentAPI)
	mux.HandleFunc(apiVersion+"/comments/delete", comments.DeleteCommentAPI)
}

// UserContacts API
func userContactsApis(mux *http.ServeMux) {
	mux.HandleFunc(apiVersion+"/user_contacts/create", user_contacts.CreateUserContactAPI)
	mux.HandleFunc(apiVersion+"/user_contacts/read", user_contacts.ReadUserContactByUserIdAPI)
	mux.HandleFunc(apiVersion+"/user_contacts/update", user_contacts.UpdateUserContactAPI)
	mux.HandleFunc(apiVersion+"/user_contacts/delete", user_contacts.DeleteUserContactAPI)
}
