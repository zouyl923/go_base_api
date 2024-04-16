package errx

var en = map[int]string{
	Success: "operate ok！",
	Error:   "operate failed！",
}

func GetEnMessage(code int) string {
	msg, ok := en[code]
	if !ok {
		msg = en[Error]
	}
	return msg
}
