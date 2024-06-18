package extract

import "bytes"

// BetweeMarkers extracts the bytes between two byte slice markers exclusively.
// If no start marker is found, the whole data is returned.
// If no end marker is found, the data from the start marker until the end of the data is returned.
func BetweenMarkers(data, startMarker, endMarker []byte) []byte {
	start := bytes.Index(data, startMarker)
	if start == -1 {
		return data
	}

	start += len(startMarker)

	end := bytes.Index(data[start:], endMarker)
	if end == -1 {
		return data[start:]
	}

	end += start

	return data[start:end]
}
