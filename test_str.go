package main

import (
	"fmt"
	"os"
)

func main() {
	var deviceid = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 123, 124, 125, 126}
	var eventtype = [...]int{1, 3, 5, 7, 8, 13, 15, 16, 17, 18, 20, 25}
	var etime int
	etime = 1661961600000 //+300000
	file, err := os.OpenFile("insert.sql", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	for i := 0; i < len(deviceid); i++ {
		for j := 1; j < 9; j++ {
			for k := 1; k < len(eventtype); k++ {
				etime += 250000
				str := fmt.Sprintf("INSERT INTO public.pecalarmextend(actiongroupt,alarmstatus,buildingid,channelid,code1,code2,confirmeventstatus,convtype,description,"+
					"deviceid,devicetype,devmodelid,eventbyte,eventtime,floorid,malarmstatus,manualfilterflag,peceventid,peceventlevel,peceventtype,roomid,ruleid,stationflag,"+
					"stationid) VALUES (0,1,NULL,5,1069,0,1,100,'测试事件%d',%d,'linesegment',63,32771,%d,4,NULL,0,-5051379875167811654,2,%d,1,0,0,1);\n", k+12*(j-1)+12*8*i, deviceid[i], etime, eventtype[k])
				file.WriteString(str)
			}
		}
	}
}
