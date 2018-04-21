package lib

func ShowError(code int) string {
	codeErrors := map[int]string{
		101: "The file doesn't exist",
		102: "The CSV file has a line with the wrong number of fields",
		201: "The connection with the db failed",
		202: "The scan of the row failed",
		203: "The read of the row failed",
	}
	result, ok := codeErrors[code]
	if ok {
		return result
	}
	return "Code error not found"
}
