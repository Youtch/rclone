package protondrive_test

import (
	"testing"

	"github.com/Youtch/rclone/backend/protondrive"
	"github.com/Youtch/rclone/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	fstests.Run(t, &fstests.Opt{
		RemoteName: "TestProtonDrive:",
		NilObject:  (*protondrive.Object)(nil),
	})
}
