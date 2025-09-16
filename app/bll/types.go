package bll

type Result struct {
	Result bool   `json:"result"`
	Msg    string `json:"msg"`
	Data   any    `json:"data"`
}

func NilDataErr(err error) Result {
	if err != nil {
		return Err(err)
	}
	return Result{
		Result: true,
		Msg:    "success",
		Data:   nil,
	}
}

func Err(err error) Result {
	return Result{
		Result: false,
		Msg:    err.Error(),
		Data:   nil,
	}
}

func DataErr(data any, err error) Result {
	if err != nil {
		return Err(err)
	}
	return Result{
		Result: true,
		Msg:    "success",
		Data:   data,
	}
}
