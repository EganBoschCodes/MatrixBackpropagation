package datasets

import (
	"fmt"

	"github.com/EganBoschCodes/lossless/utils"
)

func IsCorrect(output []float64, target []float64) {
	fmt.Printf("Output: %.2f\nTarget: %.2f\nCorrect: %t\n\n", output, target, utils.GetMaxIndex(output) == FromOneHot(target))
}
