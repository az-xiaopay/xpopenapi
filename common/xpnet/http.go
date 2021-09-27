package xpnet

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"time"
	"xpopenapi/common/mylog"
	"xpopenapi/common/util"
)

func HttpGet(url string) ([]byte, error) {
	guid := util.GetGUID()

	mylog.Debug(guid, "http请求地址", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	//mylog.Debug("http请求返回", resp)
	//mylog.Debug("http请求状态码", resp.Status)
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	//var r bytes.Buffer
	//json.Indent(&r, result, "", "  ")
	mylog.Debug(guid, "返回结果：", string(result))
	//mylog.Debug("\n\n")

	return result, nil
}

func HttpPost(url string, data interface{}) (ret []byte, err error) {
	// 超时时间：5秒
	client := &http.Client{Timeout: 50 * time.Second}
	jsonStr, err := json.Marshal(data)
	if err != nil {
		mylog.Error(err)
		return nil, err
	}

	guid := util.GetGUID()

	//var b bytes.Buffer
	//json.Indent(&b, jsonStr, "", "  ")
	mylog.Debug(guid, "http请求地址", url, "http请求包体", string(jsonStr))

	resp, err := client.Post(url, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		mylog.Error(err)
		return nil, err
	}
	defer resp.Body.Close()
	//mylog.Debug("http请求返回", resp)
	//mylog.Debug("http请求状态码", resp.Status)
	result, _ := ioutil.ReadAll(resp.Body)

	//var r bytes.Buffer
	//json.Indent(&r, result, "", "  ")
	mylog.Debug(guid, "返回结果：", string(result))
	//mylog.Debug("\n\n")

	return result, nil
}

func HttpPostFile(url string, filename string) (ret []byte, err error) {
	//var b bytes.Buffer
	//json.Indent(&b, jsonStr, "", "  ")
	mylog.Debug("http请求地址", url)
	mylog.Debug("http请求包体", "")

	body_buf := bytes.NewBufferString("")
	body_writer := multipart.NewWriter(body_buf)
	boundary := body_writer.Boundary()

	fs := []io.Reader{}
	var totalFileSize int64

	_, err = body_writer.CreateFormFile("media", filename)
	if err != nil {
		return ret, err
	}
	totalFileSize += int64(body_buf.Len())
	defer body_writer.Close()
	fs = append(fs, body_buf)

	//文件内容
	fh, err := os.Open(filename)
	if err != nil {
		return ret, err
	}
	defer fh.Close()
	fi, err := fh.Stat()
	if err != nil {
		return ret, err
	}

	totalFileSize += fi.Size()
	fs = append(fs, fh)

	//结束
	close_buf := bytes.NewBufferString(fmt.Sprintf("\r\n--%s--\r\n", boundary))
	totalFileSize += int64(close_buf.Len())
	fs = append(fs, close_buf)

	request_reader := io.MultiReader(fs...)

	req, err := http.NewRequest("POST", url, request_reader)
	if err != nil {
		return ret, err
	}

	// Set headers for multipart, and Content Length
	// req.Header.Add("Content-Type", "multipart/form-data; boundary="+boundary)
	req.Header.Add("Content-Type", "multipart/form-data; boundary="+boundary)
	req.ContentLength = totalFileSize

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return ret, err
	}
	defer resp.Body.Close()
	mylog.Debug("http请求返回", resp)
	mylog.Debug("http请求状态码", resp.Status)
	result, _ := ioutil.ReadAll(resp.Body)

	//var r bytes.Buffer
	//json.Indent(&r, result, "", "  ")
	mylog.Debug("返回结果：", string(result))

	return result, nil
}
