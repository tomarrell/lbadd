package page

// Loader describes a component that is used to load pages from files. The
// manager can infer the exact location of each page from its page ID. Depending
// on the implementation, the manager might have to be fed with the whole
// database.
type Loader interface {
	// Load retrieves the page with the given ID from secondary storage.
	Load(id ID) (Page, error)
}
