package compact

import (

	vidio "github.com/video-descriptor-matching/vidio"
	cv "gocv.io/x/gocv"
)

type ONMFCompactDescriptor struct {
	Left *cv.Mat
	Right *cv.Mat
}



func NewONMFCompactDescriptor(left, right *cv.Mat) *ONMFCompactDescriptor {
	return &ONMFCompactDescriptor{left, right}
}

func NewONMFCompactDescriptorFromGroupOfPictures(group_of_pictures *vidio.GroupOfPictures) *ONMFCompactDescriptor {

	// return NewONMFCompactDescriptor(vidio.NewMat(), vidio.NewMat())
}
// func main() {
// 	// Load the image
// 	img := gocv.LoadImage("../../testdata/lena.jpg")

// 	// Create a descriptor
// 	desc := NewSIFT()

// 	// Extract keypoints and descriptors
// 	keypoints, descriptors := desc.Compute(img)

// 	// Show the keypoints
// 	gocv.ShowKeypoints("SIFT", img, keypoints)
// }
