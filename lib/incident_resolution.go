package lib


import t "time"
//import f "fmt"



func DTimeDiff(startDate string, endDate string) float64{ //get time difference in days

	input1 := startDate	
	input2 := endDate
	layout := "2006-01-02"
	t1, _ := t.Parse(layout, input1)
	t2, _ := t.Parse(layout, input2)
	//f.Println("Parameter1: ",startDate)
	//f.Println("Parameter2: ",endDate)
	//f.Println("Date 1: ",t1)
	//f.Println("Date 2: ",t2)
	days := t2.Sub(t1).Hours() / 24
	return days
}

	func HTimeDiff(startDate string, endDate string) float64{ //get time difference in days

		input1 := startDate	
		input2 := endDate
		layout := "2006-01-02"
		t1, _ := t.Parse(layout, input1)
		t2, _ := t.Parse(layout, input2)
		//f.Println("Parameter1: ",startDate)
		//f.Println("Parameter2: ",endDate)
		//f.Println("Date 1: ",t1)
		//f.Println("Date 2: ",t2)
		hours := t2.Sub(t1).Hours() 
		return hours
}

func MTimeDiff(startDate string, endDate string) int{ //get time difference in Months go lang

	input1 := startDate	
	input2 := endDate
	layout := "2006-01-02"
	t1, _ := t.Parse(layout, input1)
	t2, _ := t.Parse(layout, input2)
	tempdate:=t1
	var x int
	for x:=0;tempdate.Before(t2);x++{
		tempdate = t1.AddDate(0,0,x)
			
		}
	return x //return the months between 2 dates in go lang
}

