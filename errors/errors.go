package errors

func ParamsError(msg string, err error) {
	if err != nil {
		panic(err)
	}
}

func HandleError(msg string, err error) {
	if err != nil {
		panic(err)
	}
}
