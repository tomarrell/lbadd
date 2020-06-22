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
* `tables` is a pointer cell which points to a page, that contains pointers to
  all tables that are stored in this database. The format of the table pages is explained in the next section.

### Table pages
Table pages do not directly hold data of a table. Instead, they hold pointers to
pages, that do, i.e. the index and data page. Table pages do however hold
information about the table schema. The schema information is a single record
that is to be interpreted as a schema (<span style="color:red;">**TODO:
schemas**</span>).

The keys of the three values, index page, data page and schema are as follows.

* `schema` is a record cell containing the schema information about this table.
  That is, columns, column types, references, triggers etc. How the schema
  information is to be interpreted, is explained [here](#)(<span
  style="color:red;">**TODO: schemas**</span>).
* `index` is a pointer cell pointing to the index page of this table. The index
  page contains pages that are used as nodes in a btree. See more
  [here](#index-pages)
* `data` is a pointer cell pointing to the data page of this table. See more
  [here](#data-pages)

### Index pages

### Data pages