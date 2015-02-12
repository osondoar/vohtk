package recognizer

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"

	"github.com/osondoar/vohtk/errors"
	"github.com/osondoar/vohtk/models"
)

type appError struct {
	Error   error
	Message string
	Code    int
}

type Recognizer struct {
}

func AdjustAudio(filename string) error {
	err := exec.Command("sox",
		"-v", "0.3",
		"tmp/wav/"+filename+".wav",
		"tmp/wav/"+filename+"-adjusted.wav",
		"rate", "16k").Run()

	if err != nil {
		return err
	}

	err = exec.Command("mv",
		"tmp/wav/"+filename+"-adjusted.wav",
		"tmp/wav/"+filename+".wav").Run()

	if err != nil {
		return err
	}

	return nil
}

func Recognize(filename string) (models.Transcription, *errors.AppError) {
	transcription := models.Transcription{}

	err := AdjustAudio(filename)
	if err != nil {
		return transcription, &errors.AppError{err, fmt.Sprint("Error running Sox: ", err.Error()), 422}
	}

	var outputBuffer bytes.Buffer
	wavFile := "tmp/wav/" + filename + ".wav"
	mfcFile := "tmp/mfc/" + filename + ".mfc"
	mlfFile := "tmp/mlf/" + filename + ".mlf"

	macrosFile := "htk/model/macros"
	hmmDefsFile := "htk/model/hmmdefs"
	modelConfig := "htk/model/config"
	wdnetFile := "htk/def/wdnet"
	dicGrammarFile := "htk/def/dict_grammar.dic"
	monophonesFile := "htk/training/monophones1"
	wavConfig := "htk/training/wav_config"

	_, err = exec.Command("HCopy", "-A", "-D", "-T", "1", "-C", wavConfig, wavFile, mfcFile).Output()
	if err != nil {
		log.Print("Error running HCopy: ", err.Error())
	}

	c1 := exec.Command("HVite", "-A", "-D", "-T", "1",
		"-H", macrosFile,
		"-H", hmmDefsFile,
		"-C", modelConfig,
		"-l", "'*'",
		"-i", mlfFile,
		"-w", wdnetFile,
		"-p", "0.0",
		"-s", "5.0",
		dicGrammarFile,
		monophonesFile,
		mfcFile)

	c2 := exec.Command("sed", "-n", `s/.*SENT-START\(.*\)SENT-END.*/\1/p`)

	c2.Stdin, _ = c1.StdoutPipe()
	c2.Stdout = &outputBuffer
	// c2.Stderr = os.Stdout // Use it to log errors with HVite
	_ = c2.Start()
	err = c1.Run()
	_ = c2.Wait()

	if err != nil {
		log.Print("Error running HVite: ", err.Error())
		return transcription, &errors.AppError{err, fmt.Sprint("Error running HVite: ", err.Error()), 422}
	}

	log.Print(outputBuffer.String())
	transcription.Text = outputBuffer.String()
	return transcription, nil
}
