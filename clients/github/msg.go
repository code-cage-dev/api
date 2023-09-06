package github

type msg struct {
	StateInvalid       string
	CodeExchangeFailed string
	GetUserInfoFailed  string
}

var Msg = msg{
	StateInvalid:       "github_state_invalid",
	CodeExchangeFailed: "github_code_exchange_failed",
	GetUserInfoFailed:  "github_get_user_info_failed",
}
