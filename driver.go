package main

import (
	"fmt"
	// "io"

	// factor "github.com/video-descriptor-matching/compact/factor"
	//vidio "github.com/video-descriptor-matching/vidio"
)

func main() {
	// w, h := vidio.GetVideoSize("data/smiling.mp4")
	var fileName string = "data/smiling.mp4"

	fmt.Println(fileName)
	// _, pw1 := io.Pipe()
	// pr2, pw2 := io.Pipe()

	// done1 := vidio.ReadVideo(fileName, pw1)
	// process(pr1, pw2, w, h)
	// _ = vidio.ReadVideo(fileName, pw1)
	// n, err := io.ReadFull(reader, buf)

	// fmt.Printf("Width: %d, Height: %d", w, h)
}

// func process(reader io.ReadCloser, writer io.WriteCloser, w, h int) {
// 	go func() {
// 		frameSize := w * h * 3
// 		buf := make([]byte, frameSize, frameSize)
// 		for {
// 			n, err := io.ReadFull(reader, buf)
// 			if n == 0 || err == io.EOF {
// 				_ = writer.Close()
// 				return
// 			} else if n != frameSize || err != nil {
// 				panic(fmt.Sprintf("read error: %d, %s", n, err))
// 			}
// 			for i := range buf {
// 				buf[i] = buf[i] / 3
// 			}
// 			n, err = writer.Write(buf)
// 			if n != frameSize || err != nil {
// 				panic(fmt.Sprintf("write error: %d, %s", n, err))
// 			}
// 		}
// 	}()
// 	return
// }