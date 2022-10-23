package main

import (
	"fmt"
	"os"
)

func main() {
	var deviceid = [...]int{502497, 502498, 502499, 502500, 502501, 502502, 502503, 502504, 502505, 502506, 502507, 502508, 502509, 502510, 502511, 502512, 502513, 502514, 502515, 502516,
		502517, 502518, 502519, 502520, 502521, 502522, 502523, 502524, 502525, 502526, 502527, 502528, 502529, 502530, 502531, 502532, 502533, 502534, 502535, 502536, 502537, 502538, 502539,
		502540, 502541, 502542, 502543, 502544, 502545, 502546, 502547, 502548, 502549, 502550, 502551, 502552, 502553, 502554, 502555, 502556, 502557, 502558, 502559, 502560, 502561, 502562,
		502563, 502564, 502565, 502566, 502567, 502568, 502569, 502570, 502571, 502572, 502573, 502574, 502575, 502576, 502577, 502578, 502579, 502580, 502581, 502582, 502583, 502584, 502585,
		502586, 502587, 502588, 502589, 502590, 502591, 502592, 502593, 502594, 502595, 502596}
	var etime int
	etime = 1664553600000 //+9000000

	file, err := os.OpenFile("insert.sql", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	for i := 0; i < len(deviceid); i++ {
		for j := 1; j < 6; j++ {
			etime += 3000000
			str := fmt.Sprintf("INSERT INTO public.pecalarmextend(alarmstatus,buildingid,channelid,code1,code2,confirmeventstatus,convtype,description,"+
				"deviceid,devicetype,devmodelid,eventbyte,eventtime,floorid,malarmstatus,manualfilterflag,peceventid,peceventlevel,peceventtype,roomid,ruleid,stationflag,"+
				"stationid) VALUES (1,NULL,5,1069,0,1,100,'测试事件%d',%d,'linesegment',63,32771,%d,4,NULL,0,-5051379875167811654,2,1,1,0,0,1);\n", i*5+j, deviceid[i], etime)
			file.WriteString(str)
		}
	}
}
