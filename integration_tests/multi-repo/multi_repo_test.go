package multi_repo

import (
	"context"
	"github.com/metrumresearchgroup/command"
	. "github.com/metrumresearchgroup/pkgr/testhelper"
	"github.com/sebdah/goldie/v2"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func setupMultiRepoTest(t *testing.T) {
	err := os.RemoveAll("test-library")
	if err != nil {
		t.Fatalf("failed to remove test library at beginning of test, %s", err)
	}
	err = os.RemoveAll("test-cache")
	if err != nil {
		t.Fatalf("failed to remove test cache at beginning of test, %s", err)
	}
}

// Test IDs
const(
	multiRepoE2ETest1 = "MRPO-E2E-001"
	multiRepoE2ETest2 = "MRPO-E2E-002"
	multiRepoE2ETest3 = "MRPO-E2E-003"
	multiRepoE2ETest4 = "MRPO-E2E-004"
	multiRepoE2ETest5 = "MRPO-E2E-005"
)

// Golden file names
const (
	multiRepoPlan = "multi-repo-plan"
	multiRepoInstallation = "multi-repo-installed-packages"
)

func TestMultiRepoInstall(t *testing.T) {
	t.Run(MakeTestName(multiRepoE2ETest1, "pkgr plan takes packages from both local and remote repos in the order listed in pkgr.yml"), func(t *testing.T) {
		setupMultiRepoTest(t)

		ctx := context.TODO()
		planCmd := command.New()

		capture, err := planCmd.Run(ctx, "pkgr", "plan", "--loglevel=debug", "--logjson")
		if err != nil {
			t.Fatalf("error occurred when installing packages: %s", err)
		}

		pkgRepoSetLogs := CollectPkgRepoSetLogs(t, capture)

		// This line is in here to explicitly test the order repositories are listed in matters.
		// LOCALREPO has R6, 2.4.1, REMOTEREPO has 2.5.0
		assert.True(t, pkgRepoSetLogs.Contains("R6", "2.4.1", "LOCALREPO", "user_defined"), "expected R6 2.4.1 to be installed because repo containing this version was listed first.")

		// Check repositories set correctly
		g := goldie.New(t)
		g.Assert(t, multiRepoPlan, pkgRepoSetLogs.ToBytes())
	})

	t.Run(MakeTestName(multiRepoE2ETest2, "pkgr install can install packages from multiple repositories"), func(t *testing.T) {

		ctx := context.TODO()
		installCmd := command.New()
		rScriptCmd := command.New(command.WithDir("Rscripts"))
		_, err := installCmd.Run(ctx, "pkgr", "install", "--loglevel=debug", "--logjson")
		if err != nil {
			t.Fatalf("error occurred when installing packages: %s", err)
		}

		rScriptOutputBytes, err := rScriptCmd.Run(ctx, "Rscript", "--quiet", "install_test.R")
		//t.Log(string(rScriptOutputBytes.Output))
		if err != nil {
			t.Fatalf("error occurred while detecting installed packages: %s", err)
		}

		g := goldie.New(t)
		g.Assert(t, multiRepoInstallation, rScriptOutputBytes.Output)
	})

}