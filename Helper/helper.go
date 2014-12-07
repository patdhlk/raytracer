package Helper

import (
	"bufio"
	"image"
	"image/png"
	"log"
	"os"
)

type Helper struct {
}

func (h *Helper) ImageWriter(filename string, m image.Image) {
	log.Println("Saving Image")
	file, err := os.Create(filename)
	if err != nil {
		log.Println("Error Creating Imagefile ", err)
		os.Exit(1)
	}
	w := bufio.NewWriter(file)
	png.Encode(w, m)
	w.Flush()
	log.Println("image writing finished")
	file.Close()
}
