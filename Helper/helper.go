//Package Helper implements neccessary helper functions
package Helper

import (
	"bufio"
	"image"
	"image/png"
	"log"
	"math"
	"os"
)

type Helper struct {
}

//writes the image to the disk,
// used by the Scene.Render function
func (h *Helper) ImageWriter(filename string, m image.Image) error {
	log.Println("Saving Image")
	file, err := os.Create(filename)
	if err != nil {
		return err
		//log.Println("Error Creating Imagefile ", err)
		//os.Exit(1)
	}
	w := bufio.NewWriter(file)
	png.Encode(w, m)
	w.Flush()
	log.Println("image writing finished")
	file.Close()
	return nil
}

//rounding function
func Round(val float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= 0.5 {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	return round / pow
}
