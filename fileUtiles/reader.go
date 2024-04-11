package fileutiles

import "os"

func ReadTextFile(filename string) (string, error) {
	constent, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	} else {
		return string(constent), nil
	}
}
