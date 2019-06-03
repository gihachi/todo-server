package main

import(
	"testing"
	"net/http"
	"bytes"
	"io/ioutil"
	"encoding/json"
	"strconv"
)

var baseUrl string = "http://localhost:8080"

func TestPostEvent(t *testing.T){
	
	var deadline string = `2019-06-11T14:00:00+09:00`
	var title    string = `submit report`
	var memo     string = `must submit`

	jsonStr := `{"deadline":"`+deadline+`","title":"`+title+`","memo":"`+memo+`"}`
	var postUrl string = baseUrl + "/api/v1/event"

	req,err := http.NewRequest(
		"POST",
		postUrl,
		bytes.NewBuffer([]byte(jsonStr)),
	)

	if err != nil{
		t.Fatalf(err.Error())
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil{
		t.Fatalf(err.Error())
	}
	defer resp.Body.Close()

	body,err := ioutil.ReadAll(resp.Body)

	if(resp.StatusCode != 200){
		t.Fatalf("status code illegal" + string(resp.StatusCode))
	}

	var data map[string]interface{}
	json.Unmarshal(body,&data)

	if(data["status"] != "success"){
		t.Fatalf("deadline is illegal : "+string(body))
	}
	if(data["message"] != "registered"){
		t.Fatalf("title is illegal : "+string(body))
	}
}

func TestGetEvents(t *testing.T){

	var getUrl string = baseUrl+"/api/v1/event"

	resp, err := http.Get(getUrl)
	if err != nil{
		t.Fatalf(err.Error())
	}
	defer resp.Body.Close()

	if(resp.StatusCode != 200){
		t.Fatalf("statusCode illigal : "+string(resp.StatusCode))
	}

	body,_ := ioutil.ReadAll(resp.Body)
	var events EventArr
	if err := json.Unmarshal(body, &events); err != nil{
		t.Fatalf("fail cast json to EventArr")
	}
}


func TestGetOneEvent(t *testing.T){

	var getUrl string = baseUrl+"/api/v1/event/1"
	resp, err := http.Get(getUrl)
	if err != nil{
		t.Fatalf(err.Error())
	}
	defer resp.Body.Close()

	if(resp.StatusCode != 200){
		t.Fatalf("statusCode illigal : "+string(resp.StatusCode))
	}
	body,_ := ioutil.ReadAll(resp.Body)
	var event Event
	if err := json.Unmarshal(body, &event); err != nil{
		t.Fatalf("fail cast json to EventArr")
	}
}

type EventArr struct{
	Events []Event `json:"events"`
}

type Event struct{
	ID int `json:"id"`
	Title string `json:"title"`
	Deadline string `json:"deadline"`
	Memo string `json:"memo"`
}

func (event *Event) String() string{
	return "{"+strconv.Itoa(event.ID)+","+event.Title+","+event.Deadline+","+event.Memo+"}"
}

func (eventArr *EventArr) String() string{
	var arrString string = ""
	for _,event := range eventArr.Events{
		arrString += event.String()
	}
	return arrString
}