//处理error
package dealError

import "fmt"

func DealError(s error)  {
	if s != nil{
		fmt.Println(s)
	}
}
