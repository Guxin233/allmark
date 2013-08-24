// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package html

import (
	"github.com/andreaskoch/allmark/repository"
	"strings"
)

func Convert(item *repository.Item) string {

	// assign the raw markdown content for the add-ins to work on
	convertedContent := strings.TrimSpace(strings.Join(item.RawContent, "\n"))

	// render markdown extensions
	convertedContent = renderImageGalleries(item, convertedContent)
	convertedContent = renderFileLinks(item, convertedContent)
	convertedContent = renderCSVTables(item, convertedContent)
	convertedContent = renderPDFs(item, convertedContent)
	convertedContent = renderVideos(item, convertedContent)
	convertedContent = renderAudio(item, convertedContent)

	// render markdown
	convertedContent = renderMarkdown(item, convertedContent)

	return convertedContent
}
