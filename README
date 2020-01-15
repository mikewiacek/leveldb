This is a fork of github.com/golang/leveldb. I've pulled out the subdirectory record package and replaced
the rest of repository with that forked copy. It was then modified to change the wire format of records to
support 2MiB chunks (instead of 32KiB chunks) and add the ability to append to one of these record files.
record_test.go was updated to reflect these changes as well.

-- Mike Wiacek 15 Jan 2020

------------

# WARNING: This is an incomplete work-in-progress.

## It is not ready for production use. Some features aren't implemented yet. Documentation is missing.

The LevelDB key-value database in the Go programming language.

To download and install from source:
$ go get github.com/golang/leveldb

Unless otherwise noted, the LevelDB-Go source files are distributed
under the BSD-style license found in the LICENSE file.
