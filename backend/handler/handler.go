package handler

import (
	"education-aws/aws"
	"encoding/json"
	"io/ioutil"
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
	requestBody := new(RequestBody)
	json.NewDecoder(r.Body).Decode(&requestBody)

	file, _, err := r.FormFile(requestBody.FileName)
	if err != nil {
		panic(err)
	}

	handler.S3.Upload(file, requestBody.FileName)

	w.WriteHeader(http.StatusOK)
}

func (handler *Handler) Download(w http.ResponseWriter, r *http.Request) {
	requestBody := new(RequestBody)
	json.NewDecoder(r.Body).Decode(&requestBody)

	file := handler.S3.Download(requestBody.FileName)

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(fileBytes)
}
