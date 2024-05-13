# Gotty

Gotty is a Go library for formatting byte sizes in a human-readable format. 
It provides a simple and flexible way to convert byte sizes to various units such as bytes, kilobytes, megabytes, gigabytes, and more.

## Installation

To install the Gotty library, use the following command:

```
go get github.com/mshadev/gotty
```

## Usage
### Basic Usage

```go
package main

import (
    "fmt"
    "github.com/mshadev/gotty"
)

func main() {
    byteSize := 1024.0 * 1024.0 * 2.5 // 2.5 MB

    formattedSize, err := gotty.Format(byteSize, nil)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println(formattedSize) // Output: 2.5 MB
}
```

### Customizing Options

```go
package main

import (
    "fmt"
    "github.com/mshadev/gotty"
)

func main() {
    byteSize := 1024.0 * 1024.0 * 2.5 // 2.5 MiB

    options := &gotty.Options{
        Binary: true,
        Space:  false,
    }

    formattedSize, err := gotty.Format(byteSize, options)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println(formattedSize) // Output: 2.5MiB
}
```

### Using Bit Units

```go
package main

import (
	"fmt"
	"github.com/mshadev/gotty"
)

func main() {
	byteSize := 1024.0 * 1024.0 * 2.5 * 8 // 20 Mbit

	options := &gotty.Options{
		Bits: true,
	}

	formattedSize, err := gotty.Format(byteSize, options)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(formattedSize) // Output: 20 Mbit
}
```

## Options

The `Format` function accepts an optional `Options` struct that allows you to customize the formatting behavior. Here are the available options:

- `Bits`: If set to `true`, the output will be in bit units (e.g., Mbit) instead of byte units (e.g., MB).
- `Binary`: If set to `true`, the output will use binary prefixes (e.g., MiB) instead of decimal prefixes (e.g., MB).
- `Space`: If set to `true`, a space will be added between the number and the unit (e.g., "2.5 MB" instead of "2.5MB").
- `Signed`: If set to `true`, a sign (+ or -) will be added for positive or negative numbers.
- `Locale`: Specifies the locale to use for number formatting (not currently implemented).
- `MinimumFractionDigits`: Specifies the minimum number of fraction digits to include in the output.
- `MaximumFractionDigits`: Specifies the maximum number of fraction digits to include in the output.

If no options are provided or set to `nil`, the default options will be used.

## License
This library is licensed under the [MIT License](LICENSE).

## Contributing
Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.