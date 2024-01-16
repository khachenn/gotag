## Installation

Download gotag by using:
```sh
go install github.com/khachenn/gotag@latest
```
## Usage

Here is an example
### Increment Major
```sh
gotag release --major # example: v1.1.2 to v2.0.0

# OR

gotag r -m
```
### Increment Minor
```sh
gotag release --minor # example: v1.1.2 to v1.2.0

# OR

gotag r -n
```

### Increment Patch
```sh
gotag release --patch # example: v1.1.2 to v1.1.3

# OR

gotag r -p
```

### Show latest version
```sh
gotag latest
```

## Author

Khachenn

## License

Licensed under the MIT License - see the [LICENSE][1] file for details.

[1]: https://github.com/khachenn/gtag/blob/main/LICENSE