package response

type healthz struct {
	ServerStatus   string `json:"server_status"`
	DatabaseStatus string `json:"database_status"`
}
