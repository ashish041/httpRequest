package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strings"
	"sync"
	"time"
)

var httpClient = &http.Client{Timeout: 10 * time.Second}

func getMD5Hash(str string) string {
	hasher := md5.New()
	hasher.Write([]byte(str))
	return hex.EncodeToString(hasher.Sum(nil))
}

func prefixUrl(url string) string {
	if !strings.HasPrefix(url, "http://") &&
		!strings.HasPrefix(url, "https://") {
		url = fmt.Sprintf("http://%s", url)
	}
	return url
}

func httpRequest(url string) (string, error) {
	response, err := httpClient.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	//fmt.Printf("Respose Body: %s \n", string(body))
	return string(body), nil
}

func handleRequest(chanUrl chan []string, wg *sync.WaitGroup) {
	for _, url := range <-chanUrl {
		if url == "" {
			continue
		}
		url = prefixUrl(url)
		response, err := httpRequest(url)
		if err != nil {
			fmt.Printf("Error: %v", err)
			continue
		}
		fmt.Printf("%s %s \n", url, getMD5Hash(response))
	}
	defer wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var procNum int

	flag.IntVar(&procNum, "parallel", 10, "Number of parallel process")
	flag.Parse()
	urls := flag.Args()

	if procNum <= 0 {
		procNum = 10
	}
	chunkLen, chunk := distribute(urls, procNum)
	chanUrl := make(chan []string, chunkLen)
	wg.Add(procNum)
	go func() {
		for _, url := range chunk {
			chanUrl <- url
		}
		close(chanUrl)
	}()
	for i := 0; i < procNum; i++ {
		go handleRequest(chanUrl, &wg)
	}
	wg.Wait()
	fmt.Println("All processing Finished.")
}

func distribute(urls []string, procNum int) (int, [][]string) {
	var chunkLen int
	var chunk [][]string

	if procNum >= len(urls) {
		chunkLen = procNum
	} else {
		chunkLen = int(math.Ceil(float64(len(urls) / procNum)))
	}
	size := (len(urls) + chunkLen - 1) / chunkLen
	for i := 0; i < len(urls); i += size {
		end := i + size
		if end > len(urls) {
			end = len(urls)
		}
		chunk = append(chunk, urls[i:end])
	}
	//for _, slice := range chunk {
	//	fmt.Printf("%#v\n", slice)
	//}
	return chunkLen, chunk
}
