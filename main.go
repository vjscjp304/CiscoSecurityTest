package main

import (
	
	"fmt"
	"net/http"
	"crypto/tls"
	"encoding/json"
	"os"
	"strconv"
	"strings"
	marathon "github.com/gambol99/go-marathon"
)


type  App struct{
	
	App marathon.Application `json:"app"`
}

func main() {
	
	baseUrl := "https://shipped-tx3-control-01.tx3.shipped-cisco.com/marathon/v2/"
	taskApi := "tasks"
	applicationApi := "apps"
	host := "shipped-tx3-worker-005"
	port := 31866
	appid := "/234334"
	
	tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }
    client := &http.Client{Transport: tr}
	
	argCount := len(os.Args[1:])
	fmt.Println("***********************************************")
	if (argCount >=2){
		host = os.Args[1]
		port,_ = strconv.Atoi(os.Args[2])

	req, err := http.NewRequest("GET", baseUrl + taskApi, nil)
	req.SetBasicAuth("synthetic-mon","VpYdy5abudqkk3Ts")
	resp, err := client.Do(req)
	if err != nil {
	    fmt.Printf("Error : %s", err)
	}
	
	flag:=false
	decoder := json.NewDecoder(resp.Body)
	var test marathon.Tasks
	decoder.Decode(&test)
	for _,t := range test.Tasks {
		if t.Host == host{
			for _,p := range t.Ports{
				if port == p{
					req, err = http.NewRequest("GET", baseUrl+applicationApi+t.AppID, nil)
					req.SetBasicAuth("synthetic-mon","VpYdy5abudqkk3Ts")
					resp, err = client.Do(req)
					if err != nil {
					    fmt.Printf("Error : %s", err)
					}
					decoder1 := json.NewDecoder(resp.Body)
					var applicant App
					decoder1.Decode(&applicant)
					fmt.Println("App Id : "+applicant.App.ID)
					for k,v :=range *applicant.App.Labels{
						fmt.Println(k+" : ",v)
					}
					flag=true
				}
			}
		}
	}
	if(!flag){
		fmt.Println("No record found.");
	}
	}else if (argCount >=1){
		appid = os.Args[1]
		appid = strings.Replace(appid, "/", "", -1)
		req, err := http.NewRequest("GET", baseUrl+applicationApi+"/"+appid, nil)
					req.SetBasicAuth("synthetic-mon","VpYdy5abudqkk3Ts")
					resp, err := client.Do(req)
					if(resp.StatusCode==200){
					if err != nil {
					    fmt.Printf("Error : %s", err)
					}
					decoder1 := json.NewDecoder(resp.Body)
					var applicant App
					decoder1.Decode(&applicant)
					fmt.Println("App Id : "+applicant.App.ID)
					for k,v :=range *applicant.App.Labels{
						fmt.Println(k+" : ",v)
					}
						for i := 0; i < len(applicant.App.Tasks); i++ {
							fmt.Println("Host_"+strconv.Itoa(i)+" : "+applicant.App.Tasks[i].Host)
							fmt.Println("Port_"+strconv.Itoa(i)+" : "+strconv.Itoa(applicant.App.Tasks[i].Ports[0]))
					} 
					}else{
						fmt.Println("No record found.");
					}
	}else{
		os.Exit(0)
	}
	fmt.Println("***********************************************")
}