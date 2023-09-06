package user

type msg struct {
	Failed   string
	NotFound string
}

var Msg = msg{
	Failed:   "user_failed",
	NotFound: "user_not_found",
}
