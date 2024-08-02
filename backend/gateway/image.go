package gateway

import (
	"context"
	"image"
	"io"
	"log"
	"os"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type ImageApp struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewImageApp() *ImageApp {
	return &ImageApp{}
}
func (a *ImageApp) Startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *ImageApp) Hello() string {
	return "just a hello"
}
func (a *ImageApp) ReadImage() []byte {

	f, err := os.OpenFile("C:\\Users\\fuqiang\\Downloads\\a.jpg", os.O_RDONLY, 0)
	if err != nil {
		log.Println(err)
		return nil
	}
	bs, err := io.ReadAll(f)
	log.Println(err)
	return bs
}

func ReadImageFile(filename string) (image.Image, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func (a *ImageApp) OpenFileDlg() string {
	opts := runtime.OpenDialogOptions{}
	filePath, err := runtime.OpenFileDialog(a.ctx, opts)
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	log.Println(filePath)
	return filePath
}
