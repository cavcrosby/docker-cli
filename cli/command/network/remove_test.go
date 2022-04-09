package network

import (
	"context"
	"io/ioutil"
	"testing"

	"github.com/docker/cli/internal/test"
	"github.com/pkg/errors"
	"gotest.tools/v3/assert"
)

func TestNetworkRemoveForce(t *testing.T) {
	fakeCli := test.NewFakeCli(&fakeClient{
		networkRemoveFunc: func(ctx context.Context, networkID string) error {
			if networkID == "foo" {
				return errors.Errorf("")
			}
			return nil
		},
	})

	cmd := newRemoveCommand(fakeCli)
	cmd.SetOut(ioutil.Discard)

	cmd.SetArgs([]string{"foo"})
	// 'network rm' currently only returns a generic error string that only states the
	// non-zero exit code.
	assert.ErrorContains(t, cmd.Execute(), "Code: 1")

	cmd.SetArgs([]string{"--force", "foo"})
	assert.NilError(t, cmd.Execute())
}
