package defaultvalue

import "strconv"

func SetDafaultValuePagination(limitStr string, pageStr string) (int, int) {
	var (
		err   error
		limit int
		page  int
	)

	// handling page
	limit, err = strconv.Atoi(limitStr)
	if err != nil {
		limit = 25
	}

	page, err = strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}

	return limit, page
}
