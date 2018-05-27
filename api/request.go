package api

type SignUpRequest struct {
	Id int32 			//Device ID
	Name string			//Device Name
	Address	string		//Device address
	Capabilities string	//Device capabilities
	Services string   //List of services  the device will use
	Uuid string	   		//Registered user of the device
}

type DBRequest struct {
	Address string
}