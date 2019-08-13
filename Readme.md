# What this tool does

Pulls the dataset from a cryptocurrency API to get live data such as Bitcoin. This tool runs the MACD technical analysis on that dataset to generate a histogram. Using the historgram, the tool can generate slopes with denoising functions.

TLDR; find a coin you want to invest in? Run this technical analysis tool to find out if the coin is trending in a profitable direction!

# Quick setup

Configure GO 11 and run `source ~/.bashrc`. Build and install with `make all`.

Start the tool with `coinsys fetch`, `coinsys start`, then visit `http://localhost:8080` to check out a cool visual of the MACD indicator taking shape!

# In progress

Only developing with the testing branch. Currently developing a mongodb api wrapper.

`coinsys test`

# Future implementations

Will produce a graph that shows least squares interpolating line. This way, we can reduce the noise from the MACD histogram slopes to produce more accurate trends.
