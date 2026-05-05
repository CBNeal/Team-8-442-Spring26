package main



import (
	"fmt"
	"time"
	"strconv"
	//"os"
	"crypto/md5"
	"encoding/hex"

)


func MD5(seconds string) string{
	hash := md5.New()
	hash.Write([]byte(seconds))
	return hex.EncodeToString(hash.Sum(nil))
}

			
func main(){

	local := time.Local

	SystemTimeTest := time.Date(2015, time.Month(01), 01, 00, 01, 00, 00, local)
	var EY, EM, ED, EH, EMI, ES int 
	fmt.Scan(&EY, &EM, &ED, &EH, &EMI, &ES )

	// You need to cast EM as a Month variable, Ive named it EpochMonthFinal
	EMF := time.Month(EM)
	EpochTime := time.Date(EY, EMF, ED, EH, EMI, ES, 0, local)

	diff := SystemTimeTest.Sub(EpochTime)
	diffSeconds := (int(diff.Seconds()) / 60) * 60

	InterimSeconds := MD5(MD5(strconv.Itoa(diffSeconds)))

	charcount := 0
	intcount := 0
	finalhash := ""

	for i := 0; i < len(InterimSeconds); i++ {
			if (InterimSeconds[i] >= 'a' && InterimSeconds[i] <= 'z') || (InterimSeconds[i] >= 'A' && InterimSeconds[i] <= 'Z') {
			if charcount < 2{
			charcount++
			finalhash += string(InterimSeconds[i])
			}
		}
	}
	for i := len(InterimSeconds) - 1; i >= 0; i-- {
		if InterimSeconds[i] >= '0' && InterimSeconds[i] <= '9' {
			if intcount < 2{
			intcount++
			finalhash += string(InterimSeconds[i])
			}
	}
}

	fmt.Println(finalhash)
}
	


