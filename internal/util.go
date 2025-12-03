package internal

import "os"

func Check_error(e error) {
	if e != nil {
		panic(e)
	}
}

// Reads the entire input and returns a string of the input
func Read_file(file_path string) string {
	input_data, err := os.ReadFile(file_path)
	Check_error(err)
	// fmt.Print(string(input_data))
	return string(input_data)
}
