package cache

import (
	"fmt"
)

func Key(prefix string, checksum string) string {
	return fmt.Sprintf("%s__%s", prefix, checksum)
}
