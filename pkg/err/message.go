package err

var MsgFlags = map[int]string{
	Success:  "ok",
	Error:    "fail",
	NotFound: "not found",
	BadRequest: "bad request",
}

// GetMsg get err information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[Error]
}
