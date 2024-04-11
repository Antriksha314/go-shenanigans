package fileutiles

import "os"

func WriteTextFile(filename string, content string) error {
	return os.WriteFile(filename, []byte(content), 0644)
}
