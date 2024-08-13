//go:build linux

package mount

import (
	"testing"

	"github.com/Youtch/rclone/vfs/vfscommon"
	"github.com/Youtch/rclone/vfs/vfstest"
)

func TestMount(t *testing.T) {
	vfstest.RunTests(t, false, vfscommon.CacheModeOff, true, mount)
}
