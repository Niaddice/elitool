package path

import (
	"fmt"
	"os"
	"testing"
)

func TestPath(t *testing.T) {
	fmt.Println("getTmpDir（当前系统临时目录） = ", os.TempDir())
	fmt.Println("getCurrentAbPathByExecutable（仅支持go build） = ", getCurrentAbPathByExecutable())
	fmt.Println("getCurrentAbPathByCaller（仅支持go run） = ", getCurrentAbPathByCaller())
	fmt.Println("getCurrentAbPath（最终方案-全兼容） = ", getCurrentAbPath())
}
