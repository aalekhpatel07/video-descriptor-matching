package vidio

import (

	// "fmt"
	"encoding/json"

	// "fmt"
	"io"
	"strconv"
	"strings"

	uuid "github.com/google/uuid"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	cv "gocv.io/x/gocv"
)

func (video *Video) FPS() float64 {

	num_den := strings.Split(video.FramerateString, "/")
	var denominator int64
	numerator, err := strconv.ParseInt(num_den[0], 10, 64)
	if err != nil {
		panic(err)
	}
	denominator, err = strconv.ParseInt(num_den[1], 10, 64)
	if err != nil {
		panic(err)
	}
	ratio := float64(numerator) / float64(denominator)
	return ratio
}

type Frame struct {
	Image		*cv.Mat
}

type GroupedVideoFrame struct {
	Frame			*Frame 				`json:"frame"`
	GroupOfPictures *GroupOfPictures 	`json:"group_of_pictures"`
	IndexInGroup	int 				`json:"index"`
}

type Video struct {
	FramerateString string 		`json:"avg_frame_rate"`
	NumberOfFrames 	int			`json:"nb_frames"`
	Duration        float64		`json:"duration"`
	ID				uuid.UUID	`json:"id"`
	Width 			int 		`json:"width"`
	Height 			int 		`json:"height"`
}

type GroupOfPictures struct {
	NumberOfPictures int	`json:"nb_pictures"`
	SourceVideo		 *Video	`json:"source_video"`
	ID				 uuid.UUID	`json:"id"`
	Frames			 *cv.Mat	`json:"frames"`
}

type VideoFrameGrouper struct {
	VideoID			uuid.UUID			`json:"video_id"`
	GroupOfPictures	[]*GroupOfPictures	`json:"group_of_pictures"`
}

func (grouper *VideoFrameGrouper) GetGroup(index int) *GroupOfPictures {
	return grouper.GroupOfPictures[index]
}


func (frame *Frame) GetDescriptors(algorithm *cv.SIFT) cv.Mat {
	_, descriptors := algorithm.DetectAndCompute(*(frame.Image), cv.NewMat())
	return descriptors
}



func GetFrames(filename string) <-chan []byte {
	// w, h := getVideoSize(filename)	
	frames := make(chan []byte)
	// buff := make([]byte, w * h * 3, w * h * 3)
	_, wr := io.Pipe()

	err := ffmpeg.
				Input(filename).
				Output("pipe:").
				WithOutput(wr).
				Run()
	if err != nil {
		panic(err)
	}

	return frames
	// for {
	// 	n, err := io.ReadFull()
	// }

}

// func ReadVideo (filename string, writer io.WriteCloser) <-chan error {
// 	w, h := getVideoSize(filename)
// 	frameSize := w * h * 3
// 	buf := make([]byte, frameSize, frameSize)

// 	for {
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

// 	done := make(chan error)

// 	go func() {
// 		err := ffmpeg.
// 				Input(filename).
// 				Output("pipe:", ffmpeg.KwArgs{"format": "rawvideo", "pix_fmt": "rgb24"}).
// 				WithOutput(writer).
// 				Run()
// 		fmt.Println("Done!")
// 		_ = writer.Close()
// 		done <- err
// 		close(done)
// 	}()
// 	return done
// }


func getVideoSize(fileName string) (int, int) {
	data, err := ffmpeg.Probe(fileName)
	if err != nil {
		panic(err)
	}
	type VideoInfo struct {
		Streams []struct {
			CodecType string `json:"codec_type"`
			Width     int
			Height    int
		} `json:"streams"`
	}
	vInfo := &VideoInfo{}
	err = json.Unmarshal([]byte(data), vInfo)
	if err != nil {
		panic(err)
	}
	for _, s := range vInfo.Streams {
		if s.CodecType == "video" {
			return s.Width, s.Height
		}
	}
	return 0, 0
}