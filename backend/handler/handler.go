package handler

import (
	"education-aws/aws"
	"encoding/json"
	"fmt"
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
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
	}

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}

	err = handler.S3.Upload(file, fileHeader.Filename)

	if err != nil {
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
}

func (handler *Handler) Download(w http.ResponseWriter, r *http.Request) {
	requestBody := new(RequestBody)
	json.NewDecoder(r.Body).Decode(&requestBody)

	file, err := handler.S3.Download(requestBody.FileName)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("can not get file '%s'", requestBody.FileName)))
	}

	w.WriteHeader(http.StatusOK)

	rq := &RequestBody{
		File:     file,
		FileName: "",
	}

	response, _ := json.Marshal(rq)
	_, err = w.Write([]byte(response))
	if err != nil {
		panic(err)
	}
}
