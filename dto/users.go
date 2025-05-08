package dto

type RegisterUsersRequest struct {
	Name 			string			`json:"name"`
	Email			string			`json:"email"`
	Password		string			`json:"password"`
	RoleID			uint			`json:"role_id"`
}

type LoginUsersRequest struct {
	Email			string			`json:"email"`
	Password		string			`json:"password"`
}
