package api_controllers

import (
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/osondoar/vohtk/errors"
	"github.com/osondoar/vohtk/services"

	"code.google.com/p/go-uuid/uuid"
)

type RequestsController struct {
	ApiController
}

func (controller RequestsController) Post(w http.ResponseWriter, r *http.Request) *errors.AppError {
	filename, err := ProcessWavFile(r)
	if err != nil {
		return &errors.AppError{err, fmt.Sprint("Error creating wav file: ", err.Error()), 422}
	}

	transcription, recognizeErr := recognizer.Recognize(filename)
	if recognizeErr != nil {
		return recognizeErr
	}

	controller.Render(w, r, transcription.ToJson())
	return nil
}

// http://golang.org/pkg/mime/multipart/#example_NewReader
func ProcessWavFile(r *http.Request) (string, error) {
	_, params, _ := mime.ParseMediaType(r.Header.Get("Content-Type"))
	var filename string

	mr := multipart.NewReader(r.Body, params["boundary"])
	for {
		p, err := mr.NextPart()

		if err == io.EOF {
			return filename, err
		}
		if err != nil {
			log.Print(err)
			return filename, err
		}
		fileData, err := ioutil.ReadAll(p)
		if err != nil {
			log.Print(err)
			return filename, err
		}
		// fmt.Printf("Media Type: %q. Part: %q. Params: %q\n", mediaType, p.Header.Get("name"), params)
		if filename, err = writeFile(fileData); err != nil {
			return "", err
		}

		return filename, nil

	}
}

func writeFile(data []byte) (string, error) {
	filename := uuid.New()
	file, err := os.Create("tmp/wav/" + filename + ".wav")
	if err != nil {
		log.Printf("Error creating %q: %q", filename, err)
		return "", err
	}

	err = binary.Write(file, binary.LittleEndian, data)
	if err != nil {
		log.Print("binary.Write failed:", err)
		return "", err
	}

	return filename, nil
}
