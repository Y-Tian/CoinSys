# What this tool does

# Quick setup

Configure GO 11 and run `source ~/.bashrc`. Build and install with `make all`.

Start the tool with `coinsys fetch`, `coinsys start`, then visit `http://localhost:8080` to check out a cool visual of the MACD indicator taking shape!

# In progress

Only developing with the testing branch. Currently developing a mongodb api wrapper.

`coinsys test`

# Future implementations

Will produce a graph that shows least squares interpolating line. This way, we can reduce the noise from the MACD histogram slopes to produce more accurate trends.
