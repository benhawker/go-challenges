//Given a single array of real values, each of which represents the stock value of a company
//after an arbitrary period of time, find the best buy price and its corresponding best sell
//price (buy low, sell high)//
//To illustrate with an example, let's take the stock ticker of Company Z
//
//55.39, 109.23, 48.29, 81.59, 105.53, 94.45, 12.24
//Important to note is the fact that the array is "sorted" temporally -
//i.e. as time passes, values are appended to the right end of the array.
//Thus, our buy value will be (has to be) to the left of our sell value
//
//in the above example, the ideal solution is to buy at 48.29 and sell at 105.53
