package util

// FrequencyToRGB converts a light frequency in THz to an RGB hex string.
func FrequencyToRGB(frequency float64) string {
	wavelength := 299792.458 / frequency // Convert frequency (THz) to wavelength (nm)
	return WavelengthToRGB(wavelength)
}
