package todoapi

type apiLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type apiLoginResponse struct {
	Token string `json:"token"`
}

type apiResponse[T any] struct {
	Data T `json:"data"`
}

type apiTodo struct {
	ID   string `json:"_id,omitempty"`
	Name string `json:"description"`
	Done bool   `json:"completed"`
}
