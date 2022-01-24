package response_error

type Error struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details"`
}

var (
	ParseRequest    = Error{Code: "parse_request", Message: "can't parse request"}
	SendResponse    = Error{Code: "send_response", Message: "can't create or send code"}
	ValidationError = Error{Code: "validation_error"}

	TournamentNotFound = Error{Code: "tournament_not_found", Message: "tournament not found"}

	Unknown  = Error{Code: "unknown", Message: "unknown error"}
	Internal = Error{Code: "internal", Message: "internal error"}
)
