package err

var MsgFlags = map[int]string{
	SUCCESS: "ok",
	ERROR:   "fail",
}

// GetMsg get err information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
