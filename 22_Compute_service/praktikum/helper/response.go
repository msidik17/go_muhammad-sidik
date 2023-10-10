package helper

func ErrorResponse(Massage string) map[string]any{
	return map[string]any{
		"status":  "Failed",
		"message": Massage,
	}
}

func SuccessResponse(Massage string) map[string]any{
	return map[string]any{
		"status":  "Success",
		"message": Massage,
	}
}

func SuccessWithDataResponse(Massage string, Data any) map[string]any{
	return map[string]any{
		"status":  "Success",
		"message": Massage,
		"data":	Data,
	}
}