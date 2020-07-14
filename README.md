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

| command                       | after processing                                                                     |
| ----------------------------- | ------------------------------------------------------------------------------------ |
| `pictar adjust`               | ![](https://github.com/disintegration/imaging/raw/master/testdata/flowers_small.png) |
| `pictar gray /path/to/file`   | ![](https://github.com/disintegration/imaging/raw/master/testdata/flowers_small.png) |
| `pictar invert /path/to/file` | ![](https://github.com/disintegration/imaging/raw/master/testdata/flowers_small.png) |

## API

```sh
pictar server
```

### Connect to S3
