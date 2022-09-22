# About

A quick-and-dirty tic-tac-toe engine to teach myself Go.

Would you like to play a game?

# Running

go run .

# Bugs

- We can cut down the computation by AT LEAST a factor of 8, by
  considering 4 rotations and one mirror image flip have identical
  outcomes.  (I don't THINK we get another x2 from the second axis flip,
  bc I believe that 2 rotations + flip around vert axis = flip around
  horiz axis).

- We should cache all computations at the start and then refer to
  them.  No need to go to depth 9 on our first move, then to depth 7
  on the next, then to depth 5 ... just store the depth 9 research.


- it would be fun to take params for board size (e.g. 4x4) and victory condition (3 in a row?  4?)

