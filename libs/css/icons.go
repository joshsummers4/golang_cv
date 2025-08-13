package css

import "sort"

// Important these must be correctly sorted alphabetically
var Icons = []string{
	"add",
	"check",
	"check_circle",
	"developer_mode_tv",
	"email",
	"error",
	"home",
	"key",
	"phone",
	"share",
}

func init() {
	sort.Strings(Icons)
}
