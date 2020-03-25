// Package securefs implements an afero.Fs that reads a file's content into
// memory when opening or creating it. All data is held encrypted by using
// github.com/awnumar/memguard. File writes need to be synced manually with the
// disk contents by calling file.Sync().
package securefs
