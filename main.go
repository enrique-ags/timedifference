package main

import f "fmt"
import "github.com/oblicore/lib"

func main(){

	var initialDate string
	var endDate string
	

	//lib.CaptureName("01/01/2019","02/01/2019")
	f.Println("Capture start Date: ")
	f.Scanln(&initialDate)
	f.Println("Capture end Date")
	f.Scanln(&endDate)
	f.Println("Your inputs: Start Date" , initialDate,  " End Date: " ,  endDate)
	days:=lib.DTimeDiff(initialDate,endDate)
	hours:=lib.HTimeDiff(initialDate,endDate)
	months:=lib.MTimeDiff(initialDate,endDate)
	f.Println("Difference in days: " , days) 
	f.Println("Difference in Hours: " , hours) 
	f.Println("Difference in Months: " , months) 




}