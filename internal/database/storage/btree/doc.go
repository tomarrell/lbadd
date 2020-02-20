// Package btree contains the btree struct, which is used as the primary data store of
// the database.
//
// The btree supports 3 primary operations:
// - get: given a key, retrieve the corresponding entry
// - put: given a key and a value, create an entry in the btree
// - remove: given a key, remove the corresponding entry in the tree if it
// exists
package btree
