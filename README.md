# go-grayscale

This is a simple go program to convert images to a grayscale format.

## Usage

With the weighted sum method

```console
$ go run main.go ./examples/winxp.png coeff
```

With the averaging method

```console
$ go run main.go ./examples/winxp.png avg
```


## Methods

 - averaging the RGB components
 - weighted sum of the RGB components `Y = 0.299 * R + 0.587 * G + 0.114 * B`

## Examples

Original image

![](./examples/winxp.png)


Grayscale using the weighted sum method

![](./examples/winxp-grayscaled-coeff.png)

Grayscale using the averaging method

![](./examples/winxp-grayscaled-avg.png)

## Resources

- https://en.wikipedia.org/wiki/Grayscale
- https://stackoverflow.com/questions/42516203/converting-rgba-image-to-grayscale-golang

