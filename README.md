# go-git in-memory file mode tests

Tests permissions of files and directories fully in-memory using
the go-git library.

Illustrates that a directory is returned as having permissions 644
unles it contains an executible file, in which case it is 755.

## Replicating Results

Run `go test -v`

## Implementaiton Notes

Includeded in `./testdata` is a bare git repository `repository` which contains
a root file and two subdirectories.  The first, `dir-a` contains only normal
files.  `dir-b` contains both a normal file and an executable file.

This structure is used when illustrating that `dir-a` is represented in the
billy.Filesystem as having perm 644, while `dir-b` reports `755`.

This test repository is cloned in-memory using a local `file://` URL.

