// Package v1 provides an implementation for the (page.Page) interface, as well
// as a function to load such a page from a given byte slice. v1 pages consist
// of headers and data, and NO trailer. Header fields are fixed, every header
// field has a fixed offset and memory area. Variable header fields are not
// supported. Cell types are stored in the cells.
//
// Assuming that data is given as a byte slice called data.
//
//  data = ...
//  p, err := v1.Load(data)
//  val, err := p.Header(page.HeaderVersion)
//  version := val.(uint16)
package v1
