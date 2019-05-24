package util

func CheckConnectError(err error){
	if err != nil{
		panic(err)
	}
}