package gitfs_test

import (
	"os"
	"path/filepath"
	"testing"

	gitfs "github.com/lkingland/go-git-fs-test"
)

// TestModes ensures that the modes are as expected.
func TestModes(t *testing.T) {

	// Calcualte file URL to ./testdata/repository.git
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	url := filepath.Join(cwd, "testdata", "repository.git")

	// Clone the URL, returning the Fileystem of the worktree
	fs, err := gitfs.Clone(url)
	if err != nil {
		t.Fatal(err)
	}

	// Enumerate each file and its expected mode.
	// Git presumes either 644 or 744, only caring about the executable bit
	// and leaving the rest to umask.  What we are concerned with is how
	// directories are represented, as they are not directly included in a git
	// repo.
	tests := []struct {
		Path string
		Perm uint32
		Dir  bool
	}{
		{Path: "file", Perm: 0644},
		{Path: "dir-a/file", Perm: 0644},
		{Path: "dir-b/file", Perm: 0644},
		{Path: "dir-b/executable", Perm: 0755},
		{Path: "dir-b", Perm: 0755},
		{Path: "dir-a", Perm: 0755},
	}

	// Note that .Perm() are used to only consider the least-signifigant 9 and
	// thus not have to consider the directory bit.
	for _, test := range tests {
		file, err := fs.Stat(test.Path)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%04o repository/%v", file.Mode().Perm(), test.Path)
		if file.Mode().Perm() != os.FileMode(test.Perm) {
			t.Fatalf("expected 'repository/%v' to have mode %04o, got %04o", test.Path, test.Perm, file.Mode().Perm())
		}
	}

}
