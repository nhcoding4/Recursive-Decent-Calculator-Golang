Similar to my recursive decent parser in Ruby. Seems to do everything correctly (matches my system calculator with a bunch of inputs). Builds a tree from input in order to correctly bind tokens in order of precedence:

Highest >> Lowest

1) Base Numbers / Parenthesis

2) Multiplcation / Division

3) Addition / Subtraction.

Nothing else is supposed (it's very simple). 
