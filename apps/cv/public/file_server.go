package public

import (
	"embed"

	shared "github.com/joshsummers4/golang_cv/libs/features/public"
)

//go:embed *.css
var embedded embed.FS

func init() {
	shared.AddETags(embedded)
}

var FileServer = shared.CreateFileServer(embedded)
