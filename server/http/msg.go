package http

type messages struct {
	Ok           string
	Unauthorized string
}

var Messages = messages{
	Ok:           "HTTP_OK",
	Unauthorized: "HTTP_UNAUTHORIZED",
}
