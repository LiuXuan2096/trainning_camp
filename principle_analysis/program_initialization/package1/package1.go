package package1

import (
	"fmt"
	"training_camp/principle_analysis/program_initialization/package2"
	"training_camp/principle_analysis/program_initialization/utils"
)

var V1 = utils.TraceLog("init package1 value1", package2.Value1+10)
var V2 = utils.TraceLog("init package1 value2", package2.Value2+10)

func init() {
	fmt.Println("init func in package1")
}
