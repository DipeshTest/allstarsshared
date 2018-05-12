/*
package twilio is a bussiness logic package for FLOGO Twilio connector developed by AllStars team,
this package has the functions RetrieveRecipientList and SendSMS,
RetrieveRecipientList can be used to retrieve all contacts stored/verified by user in his Twilio account,
SendSMS is used to send SMSes to multiple recipients simultaneously.
*/
package twilio

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Twilio struct {
	AccountSid string
	AuthToken  string
	UrlString  string
	MsgData    string
	To         string
	From       string
}

type ResponseData struct {
	StatusCode              int
	Message, Num, ErrorData string
}

type CallerIds struct {
	First_page_uri      string                `json:"first_page_uri"`
	End                 int                   `json:"end"`
	Previous_page_uri   string                `json:"previous_page_uri"`
	Outgoing_caller_ids []outgoing_caller_ids `json:"outgoing_caller_ids"`
	Uri                 string                `json:"uri"`
	Page_size           int                   `json:"page_size"`
	Start               int                   `json:"start"`
	Next_page_uri       string                `json:"next_page_uri"`
	Page                int                   `json:"page"`
}

type outgoing_caller_ids struct {
	Sid, Account_sid, Friendly_name string
	Phone_number                    string `json:"phone_number"`
	date_created, date_updated, uri string
}

func generateResp(resp *http.Response) ResponseData {
	var data map[string]interface{}
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(&data)
	responseMsg := ""
	statusCode := 0
	switch resp.StatusCode {
	case 404:
		{
			statusCode = resp.StatusCode
			responseMsg = resp.Status + "Error from service, please check URL/AccountID" + "Msg from error:" + data["message"].(string) + ",moreinfo from error:" + data["more_info"].(string)
		}
	case 401:
		{
			statusCode = resp.StatusCode
			responseMsg = resp.Status + "Error from service,please check AuthToken" + "Msg from error:" + data["message"].(string) + ",moreinfo from error:" + data["more_info"].(string)
		}
	case 400:
		{
			statusCode = resp.StatusCode
			responseMsg = resp.Status + "Error from service, please check To/From details" + "Msg from error:" + data["message"].(string) + ",moreinfo from error:" + data["more_info"].(string)
		}
	default:
		{
			statusCode = 1001
			responseMsg = "An error has occured, response statusCode from Twilio" + strconv.Itoa(resp.StatusCode) + "Msg from error:" + data["message"].(string) + ",moreinfo from error:" + data["more_info"].(string)

		}
	}
	//fmt.Println(resp.Status, "Response", decoder, "Error message:", err)
	return ResponseData{statusCode, responseMsg, "", ""}
}

func RetrieveRecipientList(smsData Twilio) ResponseData {
	req, _ := http.NewRequest("GET", smsData.UrlString, nil)
	req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(smsData.AccountSid+":"+smsData.AuthToken)))
	client := &http.Client{}
	resp, errResp := client.Do(req)
	if errResp != nil {
		return ResponseData{901, "Failed", "", "Check the connectivity....." + errResp.Error()}
	} else {
		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			var validReceipientList CallerIds
			bodyText, _ := ioutil.ReadAll(resp.Body)
			json.Unmarshal(bodyText, &validReceipientList)
			// /fmt.Println("printing ph num", validReceipientList.Outgoing_caller_ids[0].Phone_number)
			recipients := ""
			for _, num := range validReceipientList.Outgoing_caller_ids {
				recipients = recipients + num.Phone_number + ","
			}
			recipients = recipients[0 : len(recipients)-1]
			return ResponseData{200, "Success", recipients, ""}
		} else {
			return generateResp(resp)
		}

	}
}

func SendSMS(smsData Twilio, jobs <-chan string, results chan<- ResponseData) {

	for j := range jobs {
		msgData := url.Values{}
		msgData.Set("To", j)
		msgData.Set("From", smsData.From)
		msgData.Set("Body", smsData.MsgData)
		msgDataReader := *strings.NewReader(msgData.Encode())

		// Create HTTP request client
		client := &http.Client{}
		req, _ := http.NewRequest("POST", smsData.UrlString, &msgDataReader)
		req.SetBasicAuth(smsData.AccountSid, smsData.AuthToken)
		req.Header.Add("Accept", "application/json")
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		resp, errResp := client.Do(req)
		if errResp != nil {
			errorRespData := ResponseData{901, "Failed", j, "Check the connectivity....." + errResp.Error()}
			results <- errorRespData
		} else {
			if resp.StatusCode >= 200 && resp.StatusCode < 300 {
				response := []string{"Message sent successfully to", j, ",status ", resp.Status, ",statusCode", strconv.Itoa(resp.StatusCode)}
				successResponseData := ResponseData{200, "Success" + strings.Join(response, ":"), j, ""}
				results <- successResponseData
				//}
			} else {
				failureResponseData := generateResp(resp)
				results <- failureResponseData
			}

		}

	}
}
