//处理error
package dealError

func DealError(s error)  {
	if s != nil{
		//fmt.Println(s)
		panic(s)
	}
}
