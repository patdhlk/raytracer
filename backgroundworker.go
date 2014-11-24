package raytracer

//worker for parallel room scanning
import (
	"runtime"
)

type backgroundworker struct {
	cores int
}

func getNumbersOfCores() int {
	return runtime.NumCPU()
}

func NewBackgroundWorker() *backgroundworker {
	numCores := getNumbersOfCores()
	b := backgroundworker{numCores}
	return &b
}

//TODO: do raytracing parallel
