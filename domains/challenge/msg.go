package challenge

type msg struct {
	Failed   string
	NotFound string
}

var Msg = msg{
	Failed:   "challenge_failed",
	NotFound: "challenge_not_found",
}
