package GDrive

import (
	"fmt"
	"mime"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"

	"path/filepath"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	drive "google.golang.org/api/drive/v3"
	"google.golang.org/api/googleapi"
)

func getHTTPClient(token string, timeout string) *http.Client {

	tok := oauth2.Token{}
	tok.AccessToken = token
	config := oauth2.Config{}

	//	s, _ := time.ParseDuration(timeout + "s")
	//ctx, close := context.WithTimeout(context.Background(), 5*time.Second)
	//fmt.Println(ctx.Value(key))
	//	defer close()
	return config.Client(context.Background(), &tok)
}

func DeleteFile(fileId, token string, timeout string) (int, string) {
	srv, _ := drive.New(getHTTPClient(token, timeout))
	s, _ := time.ParseDuration(timeout + "s")
	ctx, close := context.WithTimeout(context.Background(), s)
	defer close()
	resp := srv.Files.Delete(fileId).Context(ctx).Do()
	if resp != nil {
		a := resp.(*googleapi.Error)
		return a.Code, a.Message
	}
	return 200, "File Deleted Successfully"

}
func CreateFile(token, filefullpath, emailAddr, role string, sendNotification bool, timeout string) (int, string) {
	f, err := os.Open(filefullpath)

	_, name := filepath.Split(filefullpath)

	defer f.Close()
	if err != nil {
		return 101, err.Error()

	} else {
		// if len(strings.TrimSpace(upoloadFileName)) == 0 {
		// 	upoloadFileName = f.Name()
		// }
		srv, err := drive.New(getHTTPClient(token, timeout))
		ext := filepath.Ext(f.Name())
		baseMimeType := mime.TypeByExtension(ext)
		convertedMimeType := mime.TypeByExtension(ext)

		file := &drive.File{
			Name:     name,
			MimeType: convertedMimeType,
		}
		s, _ := time.ParseDuration(timeout + "s")
		ctx, close := context.WithTimeout(context.Background(), s)
		defer close()
		res, err := srv.Files.Create(file).Context(ctx).Media(f, googleapi.ContentType(baseMimeType)).Do()
		fmt.Println("error ", reflect.TypeOf(err))
		if err != nil {

			switch reflect.TypeOf(err).String() {
			case "*googleapi.Error":
				{
					a := err.(*googleapi.Error)
					return a.Code, a.Message
				}

			default:
				{
					return 102, err.Error()
				}
			}

		} else {
			uploadSuccess := true
			if len(emailAddr) > 0 {
				permissiondata := &drive.Permission{
					Type:         "user",
					Role:         role,
					EmailAddress: emailAddr,
				}
				_, err := srv.Permissions.Create(res.Id, permissiondata).SendNotificationEmail(sendNotification).Do()
				if err != nil {
					if uploadSuccess {
						a := err.(*googleapi.Error)
						return res.HTTPStatusCode, "File Upload Successful, error while providing permissions," + a.Message

					}
				}
			}
			return res.HTTPStatusCode, "File successfully uploaded, permissions granted successfully" + res.WebContentLink
		}

	}
}

func ListFile(token, fileName, orderBy string, pageSize int64, pageToken string, timeout string) (int, string, int, string) {

	srv, err := drive.New(getHTTPClient(token, timeout))

	// if len(strings.TrimSpace(fileName)) == 0 {
	var searchString = ""
	if len(strings.TrimSpace(fileName)) != 0 {
		searchString = "name=\"" + fileName + "\""
	}
	s, _ := time.ParseDuration(timeout + "s")
	ctx, close := context.WithTimeout(context.Background(), s)
	defer close()
	resp, err := srv.Files.List().Context(ctx).PageSize(pageSize).PageToken(pageToken).Fields("kind, nextPageToken , files(id,mimeType,name,webContentLink)").Q(searchString).OrderBy(orderBy).Do()
	fmt.Println()

	if err != nil {
		a := err.(*googleapi.Error)
		return a.Code, a.Message, 0, ""
	} else {
		responseBody, _ := resp.MarshalJSON()

		return resp.HTTPStatusCode, string(responseBody), len(resp.Files), resp.NextPageToken
	}

	return 1000, "", 0, ""
	// }

}
