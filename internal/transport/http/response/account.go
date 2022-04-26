package response

type (
	Login struct {
		Token  string `json:"token"`
		Status int    `json:"status"`
	}

	Ping struct {
		Message string `json:"msg"`
	}

	Error struct {
		Error  string `json:"error"`
		Status int    `json:"status"`
	}
)
