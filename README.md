# Wave function collapse

Wave function collapse is the algorithm in which a system, 
by interacting with its environment (using measurements), 
is transformed from a superposition of possible states
into a one singular state with a well-defined value.

Algorithm initializes map in a unobserved state, where each cell 
is in superposition of all possibilities that are allowed by the 
neighbouring cells (all possibilities). Each iteration starts 
with a collapse of a cell with lowest the entropy (amount of available 
states) followed by propagation of contrains 
imposed by initial conditions to other cells.
This continues in a cycle until all cells are collapsed into a single state.

## Status

In development...



