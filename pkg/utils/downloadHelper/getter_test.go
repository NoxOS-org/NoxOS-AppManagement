package downloadHelper

import (
	"os"
	"path/filepath"
	"testing"

	"go.uber.org/goleak"
	"gotest.tools/v3/assert"
)

func TestDownload(t *testing.T) {
	defer goleak.VerifyNone(t, goleak.IgnoreTopFunction("go.opencensus.io/stats/view.(*worker).start")) // https://github.com/census-instrumentation/opencensus-go/issues/1191

	srcDir := t.TempDir()
	srcFile := filepath.Join(srcDir, "fixture.txt")
	assert.NilError(t, os.WriteFile(srcFile, []byte("fixture"), 0o600))

	dst := t.TempDir()

	err := Download(srcFile, dst)
	assert.NilError(t, err)

	_, statErr := os.Stat(filepath.Join(dst, "fixture.txt"))
	assert.NilError(t, statErr)
}
