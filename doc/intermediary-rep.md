# Intermediary Representation Format

LBADD uses it's own intermediary representation in order to decouple the SQL AST the execution layer. This has a couple of benefits. Primarily:
- Easier to add specific optimisations in the future
- Reducing the complexity of each component
- Allow for easier support of multiple SQL variants

We also have the benefit of being able to build a REPL to be able to interact with the IR.

## Format
The IR is made up of a limited set of commands. These commands can be found in [command.go](../command.go).

The column names must include valid *ASCII* characters, and the column types must exist in [column.go](../column.go).

Each command has a unique set of parameters. The grammar for each is outlined below.


#### Create Table
```
is_nullable ::= true | false
col ::= col " " col
  | <column_name> " " <column_type> " " is_nullable

expr ::= "create table" <table_name> col
```

---
**TODO**: insert, select, delete, ...
