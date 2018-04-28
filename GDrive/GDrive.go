package GDrive

import (
	"mime"
	"net/http"
	"os"
	"strings"

	"path/filepath"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	drive "google.golang.org/api/drive/v3"
	"google.golang.org/api/googleapi"
)

func getHTTPClient(token string) *http.Client {
	tok := oauth2.Token{}
	tok.AccessToken = token
	config := oauth2.Config{}
	return config.Client(context.Background(), &tok)
}

func DeleteFile(fileId, token string) (int, string) {
	srv, _ := drive.New(getHTTPClient(token))
	resp := srv.Files.Delete(fileId).Do()
	if resp != nil {
		a := resp.(*googleapi.Error)
		return a.Code, a.Message
	}
	return 200, "File Deleted Successfully"

}
func CreateFile(token, filefullpath, emailAddr, role string, sendNotification bool) (int, string) {
	f, err := os.Open(filefullpath)
	_, name := filepath.Split(filefullpath)
	defer f.Close()
	if err != nil {
		return 101, err.Error()

	} else {
		// if len(strings.TrimSpace(upoloadFileName)) == 0 {
		// 	upoloadFileName = f.Name()
		// }
		srv, err := drive.New(getHTTPClient(token))
		ext := filepath.Ext(f.Name())
		baseMimeType := mime.TypeByExtension(ext)
		convertedMimeType := mime.TypeByExtension(ext)
		file := &drive.File{
			Name:     name,
			MimeType: convertedMimeType,
		}
		res, err := srv.Files.Create(file).Media(f, googleapi.ContentType(baseMimeType)).Do()

		if err != nil {
			a := err.(*googleapi.Error)
			return a.Code, a.Message
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

func ListFile(token, fileName, orderBy string) (int, string) {

	srv, err := drive.New(getHTTPClient(token))

	// if len(strings.TrimSpace(fileName)) == 0 {
	var searchString = ""
	if len(strings.TrimSpace(fileName)) != 0 {
		searchString = "name=\"" + fileName + "\""
	}

	resp, err := srv.Files.List().Fields("kind , files(id,mimeType,name,webContentLink)").Q(searchString).OrderBy(orderBy).Do()

	if err != nil {
		a := err.(*googleapi.Error)
		return a.Code, a.Message
	} else {
		responseBody, _ := resp.MarshalJSON()

		return resp.HTTPStatusCode, string(responseBody)
	}

	return 1000, ""
	// }

}
