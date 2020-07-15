package cmd

import (
	"fmt"
	"image"
	"log"
	"path/filepath"

	"github.com/disintegration/imaging"
	"github.com/spf13/cobra"
)

func getFilter(filter string) imaging.ResampleFilter {
	switch filter {
	case "Linear":
		return imaging.Linear
	case "Hermite":
		return imaging.Hermite
	case "MitchellNetravali":
		return imaging.MitchellNetravali
	case "CatmullRom":
		return imaging.CatmullRom
	case "BSpline":
		return imaging.BSpline
	case "Gaussian":
		return imaging.Gaussian
	case "Bartlett":
		return imaging.Bartlett
	case "Lanczos":
		return imaging.Lanczos
	case "Hann":
		return imaging.Hann
	case "Hamming":
		return imaging.Hamming
	case "Blackman":
		return imaging.Blackman
	case "Welch":
		return imaging.Welch
	case "Cosine":
		return imaging.Cosine
	default:
		return imaging.Gaussian

	}
}

func saveFile(origin string, dst image.Image, cmd *cobra.Command) error {
	err := imaging.Save(dst, fmt.Sprintf("./%s.%s", getFileNameWithoutExt(origin), cmd.Flags().Lookup("extention").Value))

	if err != nil {
		log.Fatalf("Failed to save image: %v", err)
		return err
	}

	return nil
}

func saveMultiFile(fn func(filePath string) error, multiFilePath []string) error {
	for _, s := range multiFilePath {
		fn(s)
	}
	return nil
}

func getFileNameWithoutExt(path string) string {
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}
