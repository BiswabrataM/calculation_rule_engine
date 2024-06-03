package main

import (
	eng "calculation_rule_engine/eng"
	"fmt"

	"github.com/tidwall/gjson"
)

func main() {
	inputJSON := eng.GetInputJSON()
	rates := eng.GetRates()
	promotions := eng.GetPromotions()

	rate_applied_input_json := eng.ExecuteRateRules(inputJSON, rates)

	calc_input_json := eng.CalculatePrice(rate_applied_input_json)
	total_price := gjson.Get(calc_input_json, "total_price").Float()
	fmt.Printf("\n\nTotal Price %2f\n\n", total_price)

	for _, promo := range promotions {
		is_eligible, cuopon_applied_input_json := eng.ExecuteCuoponRules(calc_input_json, promo)

		if is_eligible {
			cuopon_applied_total_price := gjson.Get(cuopon_applied_input_json, "total_price").Float()

			if total_price == cuopon_applied_total_price {
				cuopon_calc_json := eng.CalculatePrice(cuopon_applied_input_json)
				cuopon_calc_total_price := gjson.Get(cuopon_calc_json, "total_price").Float()
				price_diff := total_price - cuopon_calc_total_price
				fmt.Printf("\nPromo applied: %2f - %2f PriceDiff: %2f [ label: %s ]", total_price, cuopon_calc_total_price, price_diff, promo.Label)

			} else {
				price_diff := total_price - cuopon_applied_total_price
				fmt.Printf("\nPromo applied: %2f - %2f PriceDiff: %2f [ label: %s ]", total_price, cuopon_applied_total_price, price_diff, promo.Label)
			}
		}

	}

}
