package handler

import (
	"education-aws/aws"
	"encoding/json"
	"net/http"
)

type RequestBody struct {
	FileName string `json:"file_name"`
	File     []byte `json:"file,omitempty"`
}

type Handler struct {
	S3 *aws.S3Handler
}

func (handler *Handler) Upload(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	// fileName := r.Form.Get("file_name")
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}

	handler.S3.Upload(file, fileHeader.Filename)

	w.WriteHeader(http.StatusOK)
}

func (handler *Handler) Download(w http.ResponseWriter, r *http.Request) {
	requestBody := new(RequestBody)
	json.NewDecoder(r.Body).Decode(&requestBody)

	file := handler.S3.Download(requestBody.FileName)

	w.WriteHeader(http.StatusOK)

	rq := &RequestBody{
		File:     file,
		FileName: "",
	}

	response, _ := json.Marshal(rq)
	_, err := w.Write([]byte(response))
	if err != nil {
		panic(err)
	}
}
