package helper

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

func GetFilter(filter string) imaging.ResampleFilter {
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

func SaveFile(origin string, dst image.Image, cmd *cobra.Command) error {

	savePath := cmd.Flags().Lookup("save").Value
	extention := cmd.Flags().Lookup("extention").Value

	var err error

	if saveExt := GetExt(savePath.String()); !(saveExt == "" || saveExt == ".") {
		err = imaging.Save(dst, fmt.Sprintf("%s", savePath))
	} else {
		err = imaging.Save(dst, fmt.Sprintf("%s/%s.%s", savePath, GetFileNameWithoutExt(origin), extention))
	}

	if err != nil {
		log.Fatalf("Failed to save image: %v", err)
		return err
	}

	return nil
}

func SaveMultiFile(fn func(filePath string) error, multiFilePath []string) error {
	//TODO: need to change multi thread processing
	for _, s := range multiFilePath {
		fn(s)
	}
	return nil
}

func GetExt(path string) string {
	return filepath.Ext(path)
}

func GetFileNameWithoutExt(path string) string {
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}

func Dirwalk(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, Dirwalk(filepath.Join(dir, file.Name()))...)
			continue
		}
		if Contains([]string{".png", ".jpg", ".jpeg", ".bmp", ".gif", ".tiff", ".tif"}, strings.ToLower(filepath.Ext(file.Name()))) {
			paths = append(paths, filepath.Join(dir, file.Name()))
		}
	}

	return paths
}

func Contains(array []string, e string) bool {
	for _, v := range array {
		if e == v {
			return true
		}
	}
	return false
}
