package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

//RequestPaymentToken request data from server
func (p *payRequest) RequestPaymentToken() payResponse {
	//prepare html encoded form
	formData := url.Values{}
	formData.Set("api", p.APIKey)
	formData.Add("amount", strconv.FormatInt(p.Amount, 10))
	formData.Add("redirect", p.Redirect)
	formData.Add("factorNumber", p.FactorNumber)

	//create request object
	req, _ := http.NewRequest("POST", BASEURL+"/test/send", strings.NewReader(formData.Encode()))
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	//lets build our client
	client := &http.Client{}
	res, _ := client.Do(req)
	//and close it asap
	defer res.Body.Close()

	//this is the body of our response
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		//handle the error the way you want , we just panic the hell out of an error :-D
		panic(err)
	}

	//define our response struct object
	var payresponse payResponse

	//decode json response and cast it to a struct
	json.Unmarshal([]byte(body), &payresponse)

	return payresponse
}

func (v *verifyRequest) RequestPaymentVerification() verifyResponse {
	//prepare html encoded form
	formVals := url.Values{}
	formVals.Set("api", v.APIKEY)
	formVals.Add("transId", strconv.FormatInt(v.TransID, 10))

	//create request object
	req, _ := http.NewRequest("POST", BASEURL+"/test/verify", strings.NewReader(formVals.Encode()))
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	//lets build our client
	client := &http.Client{}
	res, _ := client.Do(req)

	//and close it asap
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		//handle the error the way you want , we just panic the hell out of an error :-D
		panic(err)
	}
	//define our response struct object
	var verifyresponse verifyResponse

	//decode json response and cast it to a struct
	json.Unmarshal([]byte(body), &verifyresponse)

	return verifyresponse
}
