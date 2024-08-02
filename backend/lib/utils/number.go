package utils

import (
	"fmt"
	"math/rand"
)

// 生成min~max之间的随机float64,随机数保留n位小数
func RandomFloat64(min, max float64, n int) string {
	// 生成0~1之间的随机数，区间范围：[0.0,1.0)
	randomFloat := min + rand.Float64()*(max-min)
	//format成两位小数
	formattedString := fmt.Sprintf("%.2f", randomFloat)
	return formattedString
}
