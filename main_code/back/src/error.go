//处理error
package main

func DealError(s error)  {
	if s != nil{
		//fmt.Println(s)
		panic(s)
	}
}
