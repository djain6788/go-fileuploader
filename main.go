package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type fileMetaData struct {
	id              int
	fileName        string
	fileDescription string
	mimeType        string
	accountId       string
	recordId        string
	creationTime    time.Time
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	maxFileSize, _ := strconv.Atoi(configuration1.fileSize)
	contentLength := r.ContentLength / 1000000

	if contentLength > int64(maxFileSize) {
		http.Error(w, "request too large", http.StatusExpectationFailed)
		return
	}

	accountId := r.FormValue(configuration1.accountIdFromRequest)
	recordId := r.FormValue(configuration1.recordIdFromRequest)

	if _, err := os.Stat(accountId); os.IsNotExist(err) {
		os.Mkdir(accountId+"/"+recordId, 0700)
	}
	_ = os.Mkdir(accountId+"/"+recordId, 0700)

	fhs := r.MultipartForm.File[configuration1.filesFromRequest]
	for _, fh := range fhs {
		file, err := fh.Open()

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}
		var fmd fileMetaData
		fmd.fileName = fh.Filename
		mediaType, _, err := mime.ParseMediaType(fh.Header.Get("Content-Type"))
		if err != nil {
			fmt.Println(err)
		}
		fmd.mimeType = mediaType
		fmd.accountId = accountId
		fmd.recordId = recordId

		insertFileMetaData(fmd)
		ioutil.WriteFile(accountId+"/"+recordId+"/"+fh.Filename, fileBytes, 0700)
	}

	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

func downloadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Download File Here")

	accountId := r.FormValue("accountId")

	fileIdentifier := r.FormValue("fileIdentifier")

	openfile, err := os.Open(accountId + "/" + fileIdentifier)
	defer openfile.Close() //Close after function return
	if err != nil {
		//File not found, send 404
		http.Error(w, "File not found.", 404)
		return
	}

	FileHeader := make([]byte, 512)
	//Copy the headers into the FileHeader buffer
	openfile.Read(FileHeader)
	//Get content type of file
	FileContentType := http.DetectContentType(FileHeader)

	//Get the file size
	FileStat, _ := openfile.Stat()                     //Get info from file
	FileSize := strconv.FormatInt(FileStat.Size(), 10) //Get file size as a string

	//Send the headers
	w.Header().Set("Content-Disposition", "attachment; filename="+fileIdentifier)
	w.Header().Set("Content-Type", FileContentType)
	w.Header().Set("Content-Length", FileSize)

	openfile.Seek(0, 0)
	io.Copy(w, openfile)
	return
}

func deleteFile(w http.ResponseWriter, r *http.Request) {

	accountId := r.FormValue("accountId")
	fileIdentifier := r.FormValue("fileIdentifier")

	fmt.Println(accountId)
	fmt.Println(fileIdentifier)
}

func setupRoutes() {

	http.HandleFunc("/upload", uploadFile)
	http.HandleFunc("/download", downloadFile)
	http.HandleFunc("/delete", deleteFile)
	http.ListenAndServe(":8080", nil)
}

var configuration1 Configuration

func main() {
	fmt.Println("Hello World")
	configuration1 = initConfiguration()
	fmt.Println(configuration1)
	setupRoutes()
}
