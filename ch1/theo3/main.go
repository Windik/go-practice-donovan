package theo3

import (
	"fmt"
	"strings"
	"os"
)

func main() {
    fmt.Print(strings.Join(os.Args[1:], " "))
}