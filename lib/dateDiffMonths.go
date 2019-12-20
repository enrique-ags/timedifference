package lib
import t "time"
import f "fmt"
func main(){
	var initialDate, endDate string
	layout := "2006-01-02"
	f.Println("Capture start Date: ")
	f.Scanln(&initialDate)
	f.Println("Capture end Date")
	f.Scanln(&endDate)
	f.Println("Your inputs: Start Date" , initialDate,  " End Date: " ,  endDate)
	t1, _ := t.Parse(layout, initialDate)
	t2, _ := t.Parse(layout, endDate)
	//var x int
	tempdate:=t1
	if tempdate.Before(t2){
	 f.Println("Date 1 is lower than date 2")
	}else{
	f.Println("Date 1 is greater than date 2")
	}
	for x:=0;tempdate.Before(t2);x++{
	tempdate = t1.AddDate(0,0,x)
	f.Println(tempdate)
	
	}
}