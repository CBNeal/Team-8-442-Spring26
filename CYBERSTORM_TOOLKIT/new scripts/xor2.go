package main

//libraries
import (
	"image"
	"image/png"
	"io"
	"os"
)

// buffer for inputs
const BUF_SIZE = 4096

func processPixels(file image.Image) []byte {
	var pixels []byte
	MaxX, MaxY := file.Bounds().Max.X, file.Bounds().Max.Y

	for x := 0; x < MaxX; x++ {
		for y := 0; y < MaxY; y++ {
			r, g, b, a := file.At(x, y).RGBA()
			pixels = append(pixels, byte(r/257), byte(g/257), byte(b/257), byte(a/257))
		}
	}
	return pixels
}

// main function
func main() {
	//Open key file
	keyFile, err := os.Open("binary.txt")
	if err != nil {
		os.Stderr.WriteString("Cannot open key\n")
		os.Exit(1)
	}
	//closes keyfile after program runs
	defer keyFile.Close()

	//Read key into memory
	key, err := io.ReadAll(keyFile)
	//checks for empty key or invalid key
	if err != nil || len(key) == 0 {
		os.Stderr.WriteString("Invalid key\n")
		os.Exit(1)
	}
	//stores the length of the key
	keyLen := len(key)
	//error check for empty key
	if keyLen == 0 {
		os.Stderr.WriteString("Empty key\n")
		os.Exit(1)
	}
	var keyIndex int

	//FROM HERE ON OUT EVERYTHING IS DIFFERENT

	//open image file
	file, err := os.Open("image.png")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	cat, err := png.Decode(file)
	if err != nil {
		panic(err)
	}

	//break image down into stream of binary values
	list := processPixels(cat)

	for i := 0; i < len(list); i += BUF_SIZE {
		end := i + BUF_SIZE
		if end > len(list) {
			end = len(list)
		}
		chunk := list[i:end]
		for j := range chunk {
			chunk[j] ^= key[keyIndex]
			keyIndex = (keyIndex + 1) % keyLen
		}

		//write result to stdout in terminal
		_, writeErr := os.Stdout.Write(chunk)
		if writeErr != nil {
			os.Stderr.WriteString("Error writing output\n")
			os.Exit(1)
		}
	}
}
