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

Because I'm still learning `go` and I would really like to make this project 
somewhat decent, I rescheduled development of this algorithm. I will make it 
definetly, but I'm still experimenting with the stack I would use.

_Expected time of development: Before Christmas 2025_

### Thoughts

Many implementations of `wave function collapse` just take an input image and
generate an output image. I would like to make visual representation of the 
algorithm (whole generating process) with _probably_ `raylib`.

At first I wanted to make the `tiled` (tileset with rules) model wfc, but I will focus 
on the `bitmap` one instead (sample image of pixels -> output image of pixels).
