package response

type healthz struct {
	serverStatus   string `json:"server_status"`
	databaseStatus string `json:"database_status"`
}
