package main

import (
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	marathon "github.com/gambol99/go-marathon"
)

var (
	baseUrl        = "https://shipped-tx3-control-01.tx3.shipped-cisco.com/marathon/v2/"
	taskApi        = "tasks"
	applicationApi = "apps"
	user           = "synthetic-mon"
	pwd            = "VpYdy5abudqkk3Ts"
	_seperator     = "===================================================="
)

type App struct {
	App marathon.Application `json:"app"`
}

var (
	Args struct {
		BaseUrl string
		Host    string
		Port    int
		Appid   string
	}
)

func GetArgs() {
	flag.StringVar(&Args.BaseUrl, "ur", baseUrl, "Provide marathon api url")
	flag.StringVar(&Args.Host, "s", "", "Provide Host name need to search")
	flag.IntVar(&Args.Port, "p", 0, "Provide port to search")
	flag.StringVar(&Args.Appid, "a", "", "Provide App id ")

	flag.Parse()

}

func main() {

	GetArgs()

	fmt.Println(_seperator)
	run()
	fmt.Println(_seperator)

}
func run() {
	if len(Args.Host) > 0 {
		if Args.Port == 0 {
			fmt.Printf("\n Error: Please Provide Port number for '%s' Usage s=HOST_ID p=PORT_NUM ", Args)
			return
		}

		//fmt.Println("DEBUG port-- ", port)
		flag := false
		resp, err := GetHttpResponse(Args.BaseUrl+taskApi, "GET")
		if err != nil {
			return
		}
		decoder := json.NewDecoder(resp)
		var test marathon.Tasks
		decoder.Decode(&test)
		for _, t := range test.Tasks {
			//fmt.Println("DEBUG HOST ", t.Host)
			if t.Host == Args.Host {
				for _, p := range t.Ports {
					//fmt.Println("DEBUG port ", p)
					if Args.Port == p {
						resp, err := GetHttpResponse(Args.BaseUrl+applicationApi+t.AppID, "GET")
						if err != nil {
							return
						}
						decoder1 := json.NewDecoder(resp)
						var applicant App
						decoder1.Decode(&applicant)
						fmt.Println("App Id : " + applicant.App.ID)
						for k, v := range *applicant.App.Labels {
							fmt.Println(k+" : ", v)
						}
						flag = true
					}
				}
			}
		}
		if !flag {
			fmt.Println("No record found.")
		}
	} else if len(Args.Appid) > 0 {

		appid := strings.Replace(Args.Appid, "/", "", -1)
		resp, err := GetHttpResponse(fmt.Sprintf("%s%s/%s", Args.BaseUrl, applicationApi, appid), "GET")
		if err != nil {
			return
		}
		decoder1 := json.NewDecoder(resp)
		var applicant App
		decoder1.Decode(&applicant)
		fmt.Println("App Id : " + applicant.App.ID)
		for k, v := range *applicant.App.Labels {
			fmt.Println(k+" : ", v)
		}
		for i, tsk := range applicant.App.Tasks {
			fmt.Println(fmt.Sprintf("Host_%d : %s", i, tsk.Host))
			fmt.Println(fmt.Sprintf("Port_%d : %d", i, tsk.Ports[0]))
		}

	} else {
		os.Exit(0)
	}
}

func GetHttpResponse(ur string, typ string) (io.Reader, error) {
	//fmt.Println("DEBUG URL Call ", url)

	_, err := url.Parse(ur)
	if err != nil {
		fmt.Println("Error: Invalid Url %s", ur)
		return nil, err
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	req, err := http.NewRequest(typ, ur, nil)
	if err != nil {
		fmt.Println("Error: 1", err.Error())
		return nil, err
	}

	req.SetBasicAuth(user, pwd)
	r, e := client.Do(req)
	if e != nil {
		fmt.Println("Error: ", e.Error())
		return nil, e
	}
	if r.StatusCode != http.StatusOK {
		err = fmt.Errorf("Error: Got Invalid response,  StatusCode: %d, Status: %s", r.StatusCode, r.Status)
		fmt.Println(err)
		return nil, err

	}
	return r.Body, nil
}
