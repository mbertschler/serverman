package pkg

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

var (
	buildHash string
	buildDate string

	buildTime time.Time
)

func init() {
	i, _ := strconv.Atoi(buildDate)
	buildTime = time.Unix(int64(i), 0)
}

// VersionString returns build information as a string.
func VersionString() string {
	if len(buildHash) > 10 {
		return fmt.Sprintf(`
	Commit:     %s
	Date:       %s
	Go version: %s %s/%s`,
			buildHash[:10],
			buildTime,
			runtime.Version(), runtime.GOOS, runtime.GOARCH,
		)
	}

	return "compiled without buildinfo"
}
