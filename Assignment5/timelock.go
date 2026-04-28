package main



import (
	"fmt"
	"time"

)
const(
	Epoch = "2017 01 01 00 00 00"
	SystemTime = "2017 03 23 18 02 06"
)

func TimeTake(x string)(string, string, string, string, string, string){
	var Year string
	var Month string
	var Day string
	var Hour string
	var Minute string
	var Second string
	for i := 0; i < 19; i++{
		if i < 4 {
			Year += string(x[i])
		}else if (i > 4 && i < 7){
			Month += string(x[i])
		}else if (i > 7 && i < 10){
			Day += string(x[i])
		}else if (i > 10 && i < 13){
			Hour += string(x[i])
		}else if (i > 13 && i < 16){
			Minute += string(x[i])
		}else if (i > 16 && i <= 19){
			Second += string(x[i])
		}

	}

	return Year, Month, Day, Hour, Minute, Second

}
			
func main(){
	local, err := time.LoadLocation("America/Chicago")
		if err != nil{
			fmt.Println("ISSUE WITH TIMEZONE LOAD")
		}

	//final := SystemTime.Sub(Epoch).Seconds()

	_ = local
	//_ = final
	_ = Epoch
	
	var1, var2, var3, var4, var5, var6 := TimeTake(SystemTime)
	fmt.Println(var1)
	fmt.Println(var2)
	fmt.Println(var3)
	fmt.Println(var4)
	fmt.Println(var5)
	fmt.Println(var6)
	


}
