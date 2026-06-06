// Given an array prices where prices[i] is the price of a given stock on the ith day.
// Maximize your profit by choosing a single day to buy one stock and choosing a different
// day in the future to sell that stock.
// Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

package maximiseprofit

func MaximiseProfit(prices []int) int {
	var lowestPrice, largestDiff int

	for i, price := range prices {
		if i == 0 {
			lowestPrice = price
		}

		if price < lowestPrice {
			lowestPrice = price
		}

		if i != 0 && price-lowestPrice > largestDiff {
			largestDiff = price - lowestPrice
		}
	}

	return largestDiff
}
