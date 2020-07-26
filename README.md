# pictar

Image processing tool made by Golang. It is designed to be used as both an API and a CLI. pictar wraps the [imaging](https://github.com/disintegration/imaging) library as a CLI tool and API server.

## Installation

```sh
go install github.com/yellow-high5/pictar
```

## Supported Format

- JPEG(jpg, jpeg)
- PNG(png)
- GIF(gif)
- TIFF(tiff, tif)
- BMP(bmp)

## CLI

## Processing

| command                                | image                                |
| -------------------------------------- | ------------------------------------ |
| original                               | ![](./example/result/original.png)   |
| `pictar adjust --saturation 90 [PATH]` | ![](./example/result/adjust.png)     |
| `pictar blur --sigma 3.5 [PATH]`       | ![](./example/result/blur.png)       |
| `pictar crop 64 64 [PATH]`             | ![](./example/result/crop.png)       |
| `pictar fit 64 96 [PATH]`              | ![](./example/result/fit.png)        |
| `pictar flip horizon [PATH]`           | ![](./example/result/flip.png)       |
| `pictar gray [PATH]`                   | ![](./example/result/gray.png)       |
| `pictar invert [PATH]`                 | ![](./example/result/invert.png)     |
| `pictar resize 64 96 [PATH]`           | ![](./example/result/resize.png)     |
| `pictar rotate 90 [PATH]`              | ![](./example/result/rotate.png)     |
| `pictar sharpen --sigma 3.5 [PATH]`    | ![](./example/result/sharpen.png)    |
| `pictar thumbnail 64 64 [PATH]`        | ![](./example/result/thumbnail.png)  |
| `pictar tranpose [PATH]`               | ![](./example/result/transpose.png)  |
| `pictar transverse [PATH]`             | ![](./example/result/transverse.png) |

## Option

- -c (config) ... config file (default "config.yaml")
- -e (extension) ... specifies the extension of the output file (default "png")
- -f (filter) ... specifies a resampling filter to be used for image resizing (default "Gaussian")
- -s (save path) ... file save destination path after image processing (default ".")
- -D (target directory) ... whether to process images as directories
