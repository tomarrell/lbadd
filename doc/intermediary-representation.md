# Commands

Given that we parse SQL input into an AST, we somehow have to process that produced AST.
However, working with an AST for execution has multiple drawbacks.
1. ASTs carry a lot of information, in our case that would be string values for every token, even keyword tokens, punctuation, comments etc.
   Copying all that information in memory can slow down the process.
2. ASTs are different for every dialect of SQL.
   That means, if we want to switch our supported dialect (which is SQLite as of 2020-05), we basically have to re-write 70% of the project (scanner, parser, executor).
3. ASTs are difficult to work with.
   They have lots of nullable fields, and it's very difficult to see what grammar production rule was applied, meaning that the null checking can't be simplified.

Because of these reasons, we introduce an intermediary representation.

## Intermediary representation
> Me: I don't need this when I'm out of Uni<br>
> Also me: *building a database query tree*

The intermediary representation was heavily inspired by the QIR proposed in [this paper](https://arxiv.org/pdf/1607.04197.pdf).
It is close to the way of writing relational algebraic expressions, which allows for great flexibility in optimization and understandability.
A lot of people learn relational algebra in university, which makes our IR friendly for those interested in taking a look.

The IR is generated from our AST through a compiler.
During compilation, a lot of unnecessary information gets lost, so it's impossible to reconstruct the AST from our IR.
This is not important though, because all the important parts are kept.
We just discard things like token positions, keyword values and comments.

The finished IR is close to relational algebra, however it is a type structure rather than a string or unreadable byte code.
For now, the IR is not fully defined.
However, basic functionality exists.
A quick note on the IR documentation: Whenever it says that a struct "does anything", it doesn't do it, it just is a command for the executor to do that thing.
For example, if the documentation on `Scan` says "scans a table", what it really means is, that if the executor encounters that instruction, it has to scan that table.

### Basic operations

The IR is a set of commands that can instruct our executor.
This is because "IR" and "commands" are synonyms in this document.

Commands can be parameterized and configured.
Configuration is written in brackets, e.g. `Command[config1=value1,config2=value2]()`.
Parameters are written in parenthesis, e.g. `Command[](param1,param2)`.
Parameters are always evaluated before the command, configuration is evaluated while the command is run.

The configuration and parameters are just separated for sake of readability.
They are fields in the same struct, somewhere in our `command` package.
configuration generally influences the parameterization.
As an example, see the `Project` command.
The input list is a list of datasets, but the columns configuration influences, which part of the input is considered.

#### `Scan`
The most basic command is `Scan`.
It scans a table and returns all datasets in that table.
Notice the previous sentence actually means: "It tells the executor to consider all datasets in that table".
Please be aware that `Scan` does not actually loads all datasets from the indicated table **into memory**, it just tells the executor to consider them all.
The executor can perform necessary optimizations to not load all datasets.

<details>
<summary><code>Scan[table=Person]()</code></summary>

For the sake of ease, we will keep using this example throughout this document.

| Name | Age |
|---|---|
| Peter | 19 |
| Sandra | 43 |
| Elsa | 65 |
| Frederic | 21 |
| Serious | 36 |
| Severus | 38 |
| Sam | 22 |

</details>

#### `Select`
Another very basic command is `Select`, which is the equivalent to the relationally algebraic "select" (&sigma;).
`Select` is configured with a filter, which is the subscript to the relational algebra equivalent, and parameterized with the input.
It returns a table with datasets that all satisfy the filter expression.
This implies, that the expression must be able to be evaluated to a boolean value.

| Relational algebra | Command |
|---|---|
| &sigma;<sub>age&ge;25</sub>(Person) | <code>Select[filter=age&ge;25]\(Scan[table=Person]\(\)\)</code> |

In this case, the expression is valid, since age&ge;25 can be evaluated to a boolean value.
An invalid expression would be `"twelve"`.
However, <code>1&ne;2</code> is valid, though it is constant and may be optimized away.

<details>
<summary><code>Select[filter=age&ge;25](Scan[table=Person]())</code></summary>

| Name | Age |
|---|---|
| Sandra | 43 |
| Elsa | 65 |
| Serious | 36 |
| Severus | 38 |

</details>

#### `Project`
The third basic command is `Project`.
The relationally algebraic equivalent also is "project" (&Pi;)
Simply put, it limits what columns are considered.
The `Project` command is parameterized with a list of input datasets and returns the datasets in the same order as it received them, except that it removed all columns that should not be considered.

`Project` is configured with a list of columns that the input is filtered with.
This implies, that the expressions must be able to be evaluated to literal values.

| Relational algebra | Command |
|---|---|
| &Pi;<sub>Name,Age</sub>(Person) | <code>Project[cols=Name,Age]\(Scan[table=Person]\(\)\)</code> |

With `Project`, a few opportunities for optimization open up, for example two projections that cancel each other out.
This is, where the `Empty` command comes into play.
More on that command is further down below.
In the following example, first, the inner projection removes all columns that are not `Age`, then, the outer projection removes all columns that are not `Name`, yielding an empty list.

&Pi;<sub>Name</sub>(&Pi;<sub>Age</sub>(Person))

<details>
<summary><code>Project[cols=Age](Scan[table=Person]())</code></summary>

| Age |
|---|
| 19 |
| 43 |
| 65 |
| 21 |
| 36 |
| 38 |
| 22 |

</details>

#### `Rename`
We don't have a rename (&rho; in relational algebra) in our commands.
We just keep the aliases in the same place where we keep the original names, because we need to represent the aliases in the result table passed to the user.
The renames and references however can be handled by the compiler.
