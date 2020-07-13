# File Format
This document describes the v1.x file format of a database.

## Terms
This section will quickly describe the terms, that will be used throughout this
document.

A **database** is a single file, that holds all information of a single
database.

## Format
The database is a single file. Its size is a multiple of the page size, which is
16K or 16384 bytes for v1.x. The file consists of pages only, meaning there is
no fixed size header, only a header page (and maybe overflow pages).

### Header page
The page with the **ID 0** is the header page of the whole database. It holds
values for the following keys. The keys are given as strings, the actual key
bytes are the UTF-8 encoding of that string.

* `pageCount` is a record cell whose entry is an 8 byte big endian unsigned
  integer, representing the amount of pages stored in the file.
* `config` is a pointer cell which points to a page, that contains configuration
  parameters for this database.
* `tables` is a pointer cell which points to a page, that contains pointers to
  all tables that are stored in this database. The format of the table pages is
  explained in the next section.

### Table pages
Table pages do not directly hold data of a table. Instead, they hold pointers to
pages, that do, i.e. the index and data page. Table pages do however hold
information about the table data definition. The data definition information is
a single record that is to be interpreted as a data definition (<span
style="color:red;">**TODO: data definitions**</span>).

The keys of the three values, index page, data page and schema are as follows.

* `datadefinition` is a record cell containing the schema information about this
  table. That is, columns, column types, references, triggers etc. How the
  schema information is to be interpreted, is explained
  [here](#data-definition).
* `index` is a pointer cell pointing to the index page of this table. The index
  page points to pages that are an actual index in the table. See more
  [here](#index-pages)
* `data` is a pointer cell pointing to the data page of this table. See more
  [here](#data-pages)

### Index page

### Data pages
A data page stores plain record in a cell. Cell values are the full records,
cell keys are the RID of the record. The RID (row-ID) is an 8 byte unsigned
integer, which may not be reused for other records, even if a record was
deleted. The only scenario where an RID may be re-used is, when a record is
deleted from a page, while it is also being written into the same page or
another page (i.e. on move only) (this means, that the RID is not actually
re-used, just kept when moving or re-writing a cell). This can happen, if the
size of the record grows, and the cell has to be re-written. The cell keys aka.
RIDs are referenced by cells from the index pages. A full table scan is
performed by obtaining all cells in the data page and checking their records.

### Data definition
A data definition follows the following format (everything encoded in big
endian).

* 2 bytes `uint16` the amount of columns
* for each column
  * 2 bytes `uint16` frame for the column name
  * name bytes
  * 1 byte `bool` that is 0 if the table is **NOT**, and 1 if the column is
    nullable
  * 2 bytes `uint16` frame for the type name
  * type name bytes