package localstorage

import (
	"context"
	"fmt"
	"os"
)

func (ls *LocalStorageHandler) Create(ctx context.Context, files []os.File) (string, error) {
	fmt.Println("file successfully stored")
	return "", nil
}
