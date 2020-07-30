package server

import (
	"github.com/disintegration/imaging"
)

func ProcessImage(filePath string) {
	// TODO: Read config file
	src, _ := imaging.Open(filePath)

	dst := imaging.Invert(src)

	imaging.Save(dst, filePath)

}
