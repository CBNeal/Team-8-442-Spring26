package main



import (
	"fmt"
	"time"
	"strconv"
	"os"
	"crypto/md5"
	"encoding/hex"

)


func MD5(seconds string) string{
	hash := md5.New()
	hash.Write([]byte(seconds))
	return hex.EncodeToString(hash.Sum(nil))
}


			
func main(){

	local, err := time.LoadLocation("America/Chicago")
		if err != nil{
			fmt.Println("ISSUE WITH TIMEZONE LOAD")
		}
	SystemTimeTest := time.Date(2017, time.Month(03), 23, 18, 02, 06, 00, local)
	//var IEY, IEM, IED, IEH, IEMI, IES = TimeTake(Epoch)
	IEY := os.Args[1]
	IEM := os.Args[2]
	IED := os.Args[3]
	IEH := os.Args[4]
	IEMI := os.Args[5]
	IES := os.Args[6] 


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
	EpochTime := time.Date(EY, EMF, ED, EH, EMI, ES, 0, local)
	fmt.Println(EpochTime)

	diff := SystemTimeTest.Sub(EpochTime)
	diffSeconds := int(diff.Seconds())
	fmt.Println(diffSeconds)

	InterimSeconds := MD5(MD5(strconv.Itoa(diffSeconds)))
	fmt.Println(InterimSeconds)


}
	


