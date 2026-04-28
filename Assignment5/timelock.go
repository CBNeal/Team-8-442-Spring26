package main



import (
	"fmt"
	"time"
	"strconv"
	"bufio"
	"os"

)
const(
	Epoch = "2017 01 01 00 00 00"
	//SystemTime = "2017 03 23 18 02 06"
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

func TurnToTime(Time string)(time.Time){
// Parse the two strings into actual time Variables
	// Interim values that are still strings
	var IEY, IEM, IED, IEH, IEMI, IES = TimeTake(Time)


	// Parse into ints

	EY, err := strconv.Atoi(IEY)
	if err != nil{
		return}
	EM, err := strconv.Atoi(IEM)
	if err != nil{
		return}
	ED, err := strconv.Atoi(IED)
	if err != nil{
		return}
	EH, err := strconv.Atoi(IEH)
	if err != nil{
		return}
	EMI, err := strconv.Atoi(IEMI)
	if err != nil{
		return}
	ES, err := strconv.Atoi(IES)
	if err != nil{
		return}
	// You need to cast EM as a Month variable, Ive named it EpochMonthFinal
	EMF := time.Month(EM)
	EpochTime := time.Date(EY, EM, ED, EH, EMI, ES, 0, local)
	return EpochTime
}

	
			
func main(){
	local, err := time.LoadLocation("America/Chicago")
		if err != nil{
			fmt.Println("ISSUE WITH TIMEZONE LOAD")
		}

	// Parse the two strings into actual time Variables
	// Interim values that are still strings
	var IEY, IEM, IED, IEH, IEMI, IES = TimeTake(Epoch)


	// Parse into ints

	EY, err := strconv.Atoi(IEY)
	if err != nil{
		return}
	EM, err := strconv.Atoi(IEM)
	if err != nil{
		return}
	ED, err := strconv.Atoi(IED)
	if err != nil{
		return}
	EH, err := strconv.Atoi(IEH)
	if err != nil{
		return}
	EMI, err := strconv.Atoi(IEMI)
	if err != nil{
		return}
	ES, err := strconv.Atoi(IES)
	if err != nil{
		return}
	// You need to cast EM as a Month variable, Ive named it EpochMonthFinal
	EMF := time.Month(EM)
	EpochTime := time.Date(EY, EM, ED, EH, EMI, ES, 0, local)


	


}
