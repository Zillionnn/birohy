package main

import(
	"fmt"
	"time"
)

var bodyState map[string]string
bodyState = make(map[string]string)
bodyState["HIGH"]= "HIGH PERIOD"
bodyState["LOW"]= "LOW PERIOD"
bodyState["CRITICAL"]= "WARNING, IN CRITICAL DAY!"


func getDaysFromBirth(birthday string, date string) int64{
	m_birthday,_ := time.Parse("01/02/2006", birthday);
	m_date, _ :=time.Parse("01/02/2006",date);
	duration := (m_date.Unix() - m_birthday.Unix()) / 60/60/24
	return	duration
}

func formatState(type string, days int) string{
	s:=''
	if type=="phsical" {
		if days<12 {
			s+= bodyState["HIGH"]
		} else if days ==12 {
			s+=bodyState["CRITICAL"]
		} else {
			s+=bodyState["LOW"]
		}
	}
	return s
}

func main()  {
	duration := getDaysFromBirth("08/07/1993", "11/21/2019")
	phiscal := duration  % 23
	emotion := duration %28
	inte:=duration%33

	fmt.Println(formatState("phiscal", phiscal));
}