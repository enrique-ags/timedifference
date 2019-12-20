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

func MTimeDiff(startDate string, endDate string) float64{ //get time difference in Months

	input1 := startDate	
	input2 := endDate
	layout := "2006-01-02"
	t1, _ := t.Parse(layout, input1)
	t2, _ := t.Parse(layout, input2)
	//f.Println("Parameter1: ",startDate)
	//f.Println("Parameter2: ",endDate)
	//f.Println("Date 1: ",t1)
	//f.Println("Date 2: ",t2)
	months := t2.Sub(t1).Hours() /24
	return months
}