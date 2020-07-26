package cmd

import (
	"fmt"
	"image"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

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

	savePath := cmd.Flags().Lookup("save").Value
	extention := cmd.Flags().Lookup("extention").Value

	var err error

	if saveExt := getExt(savePath.String()); !(saveExt == "" || saveExt == ".") {
		err = imaging.Save(dst, fmt.Sprintf("%s", savePath))
	} else {
		err = imaging.Save(dst, fmt.Sprintf("%s/%s.%s", savePath, getFileNameWithoutExt(origin), extention))
	}

	if err != nil {
		log.Fatalf("Failed to save image: %v", err)
		return err
	}

	return nil
}

func saveMultiFile(fn func(filePath string) error, multiFilePath []string) error {
	//TODO: need to change multi thread processing
	for _, s := range multiFilePath {
		fn(s)
	}
	return nil
}

func getExt(path string) string {
	return filepath.Ext(path)
}

func getFileNameWithoutExt(path string) string {
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}

func dirwalk(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, dirwalk(filepath.Join(dir, file.Name()))...)
			continue
		}
		if contains([]string{".png", ".jpg", ".jpeg", ".bmp", ".gif", ".tiff", ".tif"}, strings.ToLower(filepath.Ext(file.Name()))) {
			paths = append(paths, filepath.Join(dir, file.Name()))
		}
	}

	return paths
}

func contains(array []string, e string) bool {
	for _, v := range array {
		if e == v {
			return true
		}
	}
	return false
}
