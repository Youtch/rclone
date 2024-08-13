//go:build darwin && !cmount

package nfsmount

import (
	"testing"

	"github.com/Youtch/rclone/vfs/vfscommon"
	"github.com/Youtch/rclone/vfs/vfstest"
)

func TestMount(t *testing.T) {
	vfstest.RunTests(t, false, vfscommon.CacheModeMinimal, false, mount)
}
