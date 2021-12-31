package utils

import (
	"github.com/tus/tusd/pkg/handler"
	"regexp"
	"strings"
)

var (
	reExtractFileID  = regexp.MustCompile(`([^/]+)\/?$`)
	reForwardedHost  = regexp.MustCompile(`host=([^;]+)`)
	reForwardedProto = regexp.MustCompile(`proto=(https?)`)
	reMimeType       = regexp.MustCompile(`^[a-z]+\/[a-z0-9\-\+\.]+$`)
)



// Parse the Upload-Concat header, e.g.
// Upload-Concat: partial
// Upload-Concat: final;http://tus.io/files/a /files/b/
func ParseConcat(header string) (isPartial bool, isFinal bool, partialUploads []string, err error) {
	if len(header) == 0 {
		return
	}

	if header == "partial" {
		isPartial = true
		return
	}

	l := len("final;")
	if strings.HasPrefix(header, "final;") && len(header) > l {
		isFinal = true

		list := strings.Split(header[l:], " ")
		for _, value := range list {
			value := strings.TrimSpace(value)
			if value == "" {
				continue
			}

			id, extractErr := extractIDFromPath(value)
			if extractErr != nil {
				err = extractErr
				return
			}

			partialUploads = append(partialUploads, id)
		}
	}

	// If no valid partial upload ids are extracted this is not a final upload.
	if len(partialUploads) == 0 {
		isFinal = false
		err = handler.ErrInvalidConcat
	}

	return
}

// extractIDFromPath pulls the last segment from the url provided
func extractIDFromPath(url string) (string, error) {
	result := reExtractFileID.FindStringSubmatch(url)
	if len(result) != 2 {
		return "",  handler.ErrNotFound
	}
	return result[1], nil
}
