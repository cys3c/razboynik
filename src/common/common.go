package common

import (
	"bytes"
	"fmt"
	"fuzzer"
	"io"
	"net/http"
	"os"
	"os/exec"
)

func Upload(path, dir string) {
	bytes, bondary, err := fuzzer.Upload(path, dir)

	if err {
		return
	}

	req, err := fuzzer.NET.PrepareUpload(bytes, bondary)

	if err {
		return
	}

	resp, err := fuzzer.NET.Send(req)

	if err {
		return
	}

	result := fuzzer.NET.GetBodyStr(resp)
	ReceiveOne(result, "File upload successfully")
}

func Download(path, location string) {
	php := fuzzer.Download(path)
	req, err := fuzzer.NET.Prepare(php)

	if err {
		return
	}

	resp, err := fuzzer.NET.Send(req)

	if err {
		return
	}

	ReceiveDownload(resp, location)
}

func ReceiveDownload(resp *http.Response, location string) {
	out, err := os.Create(location)
	defer out.Close()

	if err != nil {
		panic(err)
	}

	_, err = io.Copy(out, resp.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println("Downloaded successfully those byte: ")
	fmt.Println(resp.Body)
}

func ReceiveOne(r, msg string) {
	if r == "1" {
		fmt.Println(msg)
		return
	}

	fmt.Println("An error occured")
}

func Syscall(str string) {
	cmd := exec.Command("bash", "-c", str)

	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput

	err := cmd.Run()

	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("==> Error: %s\n", err))
	}

	fmt.Printf("%s\n", string(cmdOutput.Bytes()))
}