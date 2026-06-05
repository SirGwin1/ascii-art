package functions

import(
	"fmt"
)

func PrintResult(result string)error{
	if result == ""{
		return fmt.Errorf("Nothing to Print")
	}
	fmt.Print(result)
	return nil
}