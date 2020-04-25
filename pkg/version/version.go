package version

import (
	"fmt"
	"strconv"
	"time"
)

var (
	buildHash      string
	buildDate      string
	buildGoVersion string

	buildTime time.Time
)

func init() {
	i, _ := strconv.Atoi(buildDate)
	buildTime = time.Unix(int64(i), 0)
}

// String returns build information as a string.
func String() string {
	if len(buildHash) > 10 {
		return fmt.Sprintf(`
	Commit:     %s
	Date:       %s
	Go version: %s`,
			buildHash[:10],
			buildTime,
			buildGoVersion,
		)
	}

	return "compiled without buildinfo"
}
