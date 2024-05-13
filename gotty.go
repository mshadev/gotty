package gotty

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

var (
	byteUnits   = []string{"B", "kB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"}
	bibyteUnits = []string{"B", "KiB", "MiB", "GiB", "TiB", "PiB", "EiB", "ZiB", "YiB"}
	bitUnits    = []string{"b", "kbit", "Mbit", "Gbit", "Tbit", "Pbit", "Ebit", "Zbit", "Ybit"}
	bibitUnits  = []string{"b", "kibit", "Mibit", "Gibit", "Tibit", "Pibit", "Eibit", "Zibit", "Yibit"}
)

// Options represents the options for formatting the byte size.
type Options struct {
	Bits                  bool
	Binary                bool
	Space                 bool
	Signed                bool
	Locale                string
	MinimumFractionDigits int
	MaximumFractionDigits int
}

// Format formats the given byte size as a human-readable string.
// If the options parameter is nil, default options will be used.
func Format(byteSize float64, options *Options) (string, error) {
	if options == nil {
		options = &Options{}
	}

	if !isFinite(byteSize) {
		return "", errors.New(fmt.Sprintf("expected a finite number, got %T: %v", byteSize, byteSize))
	}

	units := byteUnits
	if options.Bits {
		if options.Binary {
			units = bibitUnits
		} else {
			units = bitUnits
		}
	} else if options.Binary {
		units = bibyteUnits
	}

	separator := ""
	if options.Space {
		separator = " "
	}

	if options.Signed && byteSize == 0 {
		return "0" + separator + units[0], nil
	}

	isNegative := byteSize < 0
	prefix := ""
	if isNegative {
		prefix = "-"
		byteSize = -byteSize
	} else if options.Signed {
		prefix = "+"
	}

	localeOptions := make(map[string]interface{})
	if options.MinimumFractionDigits != 0 {
		localeOptions["minimumFractionDigits"] = options.MinimumFractionDigits
	}
	if options.MaximumFractionDigits != 0 {
		localeOptions["maximumFractionDigits"] = options.MaximumFractionDigits
	}

	if byteSize < 1 {
		numberString := formatNumber(byteSize, options.Locale, localeOptions)
		return prefix + numberString + separator + units[0], nil
	}

	var exponent int
	if options.Binary {
		exponent = int(math.Min(math.Floor(math.Log(byteSize)/math.Log(1024)), float64(len(units)-1)))
		byteSize /= math.Pow(1024, float64(exponent))
	} else {
		exponent = int(math.Min(math.Floor(math.Log10(byteSize)/3), float64(len(units)-1)))
		byteSize /= math.Pow(1000, float64(exponent))
	}

	if len(localeOptions) == 0 {
		byteSize = roundToPrecision(byteSize, 3)
	}

	numberString := formatNumber(byteSize, options.Locale, localeOptions)
	unit := units[exponent]

	return prefix + numberString + separator + unit, nil
}

// isFinite checks if a given number is finite.
func isFinite(number float64) bool {
	return !math.IsInf(number, 0) && !math.IsNaN(number)
}

// roundToPrecision rounds a given number to a specified precision.
func roundToPrecision(number float64, precision int) float64 {
	scale := math.Pow(10, float64(precision))
	return math.Round(number*scale) / scale
}

// formatNumber formats a given number according to the specified locale and options.
func formatNumber(number float64, locale string, options map[string]interface{}) string {
	return strconv.FormatFloat(number, 'f', -1, 64)
}
