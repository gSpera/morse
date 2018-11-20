package morse

//ConverterOption is a function that modifies a Converter
//The main use of ConvertOption is inside NewConverter
type ConverterOption func(Converter) Converter

//WithHandler sets the handler for the Converter
func WithHandler(handler ErrorHandler) ConverterOption {
	return func(c Converter) Converter {
		c.Handling = handler
		return c
	}
}

//WithLowercaseHandling sets if the Converter may convert to uppercase before checking inside the EncodingMap
func WithLowercaseHandling(lowercaseHandling bool) ConverterOption {
	return func(c Converter) Converter {
		c.convertToUpper = lowercaseHandling
		return c
	}
}

//WithTrailingSeparator sets if the Converter may trail the charSeparator
func WithTrailingSeparator(trailingSpace bool) ConverterOption {
	return func(c Converter) Converter {
		c.trailingSeparator = trailingSpace
		return c
	}
}
