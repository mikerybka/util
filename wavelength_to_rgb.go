package util

import (
	"fmt"
	"math"
)

func WavelengthToRGB(wavelength float64) string {
	var r, g, b float64

	switch {
	case wavelength >= 380 && wavelength <= 440:
		r = -1 * (wavelength - 440) / (440 - 380)
		g = 0.0
		b = 1.0
	case wavelength >= 440 && wavelength <= 490:
		r = 0.0
		g = (wavelength - 440) / (490 - 440)
		b = 1.0
	case wavelength >= 490 && wavelength <= 510:
		r = 0.0
		g = 1.0
		b = -1 * (wavelength - 510) / (510 - 490)
	case wavelength >= 510 && wavelength <= 580:
		r = (wavelength - 510) / (580 - 510)
		g = 1.0
		b = 0.0
	case wavelength >= 580 && wavelength <= 645:
		r = 1.0
		g = -1 * (wavelength - 645) / (645 - 580)
		b = 0.0
	case wavelength >= 645 && wavelength <= 780:
		r = 1.0
		g = 0.0
		b = 0.0
	}

	// Let the intensity fall off near the vision limits
	// if wavelength > 700 {
	// 	intensityCorrection := 0.3 + 0.7*(780-wavelength)/(780-700)
	// 	r *= intensityCorrection
	// 	g *= intensityCorrection
	// 	b *= intensityCorrection
	// } else if wavelength < 420 {
	// 	intensityCorrection := 0.3 + 0.7*(wavelength-380)/(420-380)
	// 	r *= intensityCorrection
	// 	g *= intensityCorrection
	// 	b *= intensityCorrection
	// }
	return fmt.Sprintf("#%02x%02x%02x", int(math.Round(r*255)), int(math.Round(g*255)), int(math.Round(b*255)))
}
