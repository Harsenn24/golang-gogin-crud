package helper

import "strconv"

func StringToInT(data string) (int, error) {
	result, err := strconv.Atoi(data)
	if err != nil {
		return 0, err
	}

	return result, nil
}
