# imgen

By uisng `imgen` you can generate unique numbered images in a single command. It can be used to generate test images where you need to use many images but uniquely identifiable like testing sample product/items, testing different banners, placeholders, etc.

By [h4rdw1r3](https://github.com/h4rdw1r3)


## Usage

```
Usage of imgen:
  -h int
        Height of the image (default 200)
  -w int
      Width of the image (default 400)
  -t string
        Title of the image (default "test")
  -n int
      Number of images (default 5)
 
```
## Example

```
$ imgen -w 400 -h 150 -n 2 -t "test"
```
