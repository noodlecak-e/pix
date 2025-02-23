package files

import (
	"context"
	"os"
)

type IFileIO interface {
	Create(ctx context.Context, files []os.File) (string, error)
	// Get(ctx context.Context, filepath string)
	// Update()
	// Delete()
}
