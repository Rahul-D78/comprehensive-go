package concepts

import (
	"context"
	"fmt"
	"runtime"
)

type Config struct {
	osInfoClient OsInfoProvider
}

// Switch without a condition is same as `switch true`
// This is a clean way of writing long if-then-else chains
func (c Config) SystemInfo() {
	os := runtime.GOOS
	switch {
	case os == "darwin":
		fmt.Println("Go runs on Darwin")
	case os == "linux":
		fmt.Println("Go runs on Linux")
	default:
		fmt.Println("Go runs on Unknown os")
	}
	ctx := context.Background()
	c.osInfoClient.GetOsCpuArch(ctx, "localhost")
}
