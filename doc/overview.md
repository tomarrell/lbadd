# Component Overview

Each component in the database serves an isolated purpose by design. This is to make the database as modular and easier to work on. The primary components that are developed, or currently under development are listed below, with a description of their purpose towards the overall functionality of the database.

## Parser

The parser is responsible for reading and understanding the structure the of SQL queries which are sent to the database.
The output is a structured representation of the query known as an abstract syntax tree (AST).

## Codegen

LBADD executes queries in a virtual machine like environment. This machine has a specific and defined representation known as the intermediary representation (IR), which can be understood as a UST. The IR is generated from an AST. This allows for better separation between the execution and query layers of the database and also makes multi-level optimization more easy.

The codegen step in particular takes an abstract syntax tree and transforms it into the intermediary representation which can be executed by the virtual machine.

## Virtual Machine/Executor

The virtual machine (aka. executor) takes the intermediary representation, and performs the steps necessary to fulfill the commands. This includes interacting with the various storage components, aggregation, transformation, and ultimately returning the result of the executed queries in the expected format.

## Storage Frontend

The storage frontend handles the representation of the database in memory, as well as the interaction with the lower level components, in order to make sure that information can be retrieved efficiently from the database.

## Storage Backend

The storage backend handles the interaction with the persistent storage (disk), as well as paging.

## Consensus (TBD)

The consensus component handles the communication with other nodes in the cluster, making sure they are in agreement as to the current state of the database. This also handles the service discovery, leader election etc.
