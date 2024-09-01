package ioutil

import (
	"fmt"
	"os"
)

func RemoveFolder(folderPath string) {

	err := os.RemoveAll(folderPath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Folder, %s removed successfully. \n", folderPath)
}
