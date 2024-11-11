package res

type ResMapKV struct {
	HttpCode int
	Message  string
}

var ResMap map[int]ResMapKV = map[int]ResMapKV{
	SUCCESS: {
		HttpCode: 200,
		Message:  "success",
	},
	FAIL: {
		HttpCode: 500,
		Message:  "fail",
	},
}
