//go:build noselfupdate

package selfupdate

import (
	"github.com/Youtch/rclone/lib/buildinfo"
)

func init() {
	buildinfo.Tags = append(buildinfo.Tags, "noselfupdate")
}
