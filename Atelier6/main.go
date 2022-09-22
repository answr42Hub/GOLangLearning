package main

import (
	"fmt"
	"github.com/nfnt/resize"
	"image/jpeg"
	"image/png"
	"log"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func main() {

	//Number 1
	//numOneA()
	//numOneB()

	//Number 2
	/*
		files := getFiles()
		var c chan string = make(chan string)
		numTwoA(files)
		numTwoB(files, c)
	*/
}

func numOneA() {
	for i := 0; i < 10; i++ {
		go fmt.Println(randPrime())
	}
	var input string
	fmt.Scanln(&input)
}

func numOneB() {
	var c chan string = make(chan string)
	go pinger(c)
	go printer(c)
	var input string
	fmt.Scanln(&input)
}

func numTwoA(files []string) {
	for i, file := range files {
		resizeImage(file, strconv.Itoa(i))
	}
}

func numTwoB(files []string, c chan string) {
	for i, file := range files {
		go resizeRoutine(file, strconv.Itoa(i), c)
	}
	count := len(files)
	for count > 0 {
		msg := <-c
		if msg == "done" {
			count--
		}
	}
}

func resizeImage(img string, index string) {
	// open "test.jpg"
	file, err := os.Open(img)
	if err != nil {
		log.Fatal(err)
	}

	// decode jpeg into image.Image
	ext := filepath.Ext(img)

	if ext == ".jpg" || ext == ".jpeg" {
		img, err := jpeg.Decode(file)
		if err != nil {
			log.Fatal(err)
		}
		file.Close()
		// resize to width 1000 using Lanczos resampling
		// and preserve aspect ratio
		m := resize.Resize(80, 0, img, resize.Lanczos3)

		out, err := os.Create("resized_img/test_resized_" + index + ".jpg")
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()

		// write new image to file
		jpeg.Encode(out, m, nil)
	} else if ext == ".png" {
		img, err := png.Decode(file)
		if err != nil {
			log.Fatal(err)
		}
		file.Close()
		// resize to width 1000 using Lanczos resampling
		// and preserve aspect ratio
		m := resize.Resize(80, 0, img, resize.Lanczos3)

		out, err := os.Create("resized_img/test_resized" + index + ".png")
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()

		// write new image to file
		png.Encode(out, m)
	}
}

func resizeRoutine(img string, index string, c chan string) {
	resizeImage(img, index)
	c <- "done"
}

func getFiles() []string {
	var files []string

	root := "./img"

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		if err != nil {

			fmt.Println(err)
			return nil
		}

		if !info.IsDir() && filepath.Ext(path) == ".jpg" || filepath.Ext(path) == ".png" || filepath.Ext(path) == ".jpeg" {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return files
}

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- strconv.Itoa(randPrime())
	}
}

func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func randPrime() int {
	var num int
	found := false
	rand.Seed(time.Now().UnixNano())
	for !found {
		num = rand.Intn(999) + 1
		if isPrime(num) {
			found = true
		}
	}
	return num
}

func isPrime(num int) bool {
	count := 0
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			count++
			break
		}
	}
	return count == 0 && num != 1
}
