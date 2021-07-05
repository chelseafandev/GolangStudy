package https_server

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang/gddo/httputil/header"
)

type malformedRequest struct {
	status int
	msg    string
}

func (mr *malformedRequest) Error() string {
	return mr.msg
}

/*
parse a JSON request body
reference: https://www.alexedwards.net/blog/how-to-properly-parse-a-json-request-body

예외 처리까지 꼼꼼히 다 해둬서 바로 가져다 사용하기 좋다 :)
*/
func decodeJSONBody(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	if r.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if value != "application/json" {
			msg := "Content-Type header is not application/json"
			return &malformedRequest{status: http.StatusUnsupportedMediaType, msg: msg}
		}
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		case errors.As(err, &syntaxError):
			msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
			return &malformedRequest{status: http.StatusBadRequest, msg: msg}

		case errors.Is(err, io.ErrUnexpectedEOF):
			msg := fmt.Sprintf("Request body contains badly-formed JSON")
			return &malformedRequest{status: http.StatusBadRequest, msg: msg}

		case errors.As(err, &unmarshalTypeError):
			msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
			return &malformedRequest{status: http.StatusBadRequest, msg: msg}

		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			msg := fmt.Sprintf("Request body contains unknown field %s", fieldName)
			return &malformedRequest{status: http.StatusBadRequest, msg: msg}

		case errors.Is(err, io.EOF):
			msg := "Request body must not be empty"
			return &malformedRequest{status: http.StatusBadRequest, msg: msg}

		case err.Error() == "http: request body too large":
			msg := "Request body must not be larger than 1MB"
			return &malformedRequest{status: http.StatusRequestEntityTooLarge, msg: msg}

		default:
			return err
		}
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		msg := "Request body must only contain a single JSON object"
		return &malformedRequest{status: http.StatusBadRequest, msg: msg}
	}

	return nil
}

/*
Request Body에 정의된 JSON 포맷에 따른 구조체 정의
*/
type RequestData struct {
	Url       string
	User_id   string
	User_guid string
}

func send_200_response(w http.ResponseWriter) {
	var response string
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response = "{\"result\":\"ok\"}"
	_, w_err := fmt.Fprintf(w, response)
	if w_err != nil {
		fmt.Println(w_err)
	}
}

func send_400_response(w http.ResponseWriter, resp_msg string, resp_code int) {
	var response string
	// 400 error 처리를 위한 함수를 따로 만들자
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	response = "{\"error\":{\"message\":\""
	response += resp_msg
	response += "\",\"code\":"
	response += fmt.Sprintf("%d", resp_code)
	response += "}}"
	_, w_err := fmt.Fprintf(w, response)
	if w_err != nil {
		fmt.Println(w_err)
	}
}

/*
handler 함수

함수 내부 구현은 처리 내용에 따라 달라질 수 있음 :)
*/
func handleRequest(w http.ResponseWriter, r *http.Request) {

	// decodeJSONBody를 통해 파싱한 데이터를 저장할 구조체 타입 선언
	var rd RequestData

	err := decodeJSONBody(w, r, &rd)
	if err != nil {
		var mr *malformedRequest
		if errors.As(err, &mr) {
			http.Error(w, mr.msg, mr.status) // decodeJSONBody 함수 내부에서 발생한 400 Error(BadRequest)에 대한 응답 처리
		} else {
			log.Println(err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError) // malformedRequest 형태 이외의 Error가 발생한 경우에는 500 Error(Internal Server Error)로 응답 처리
		}
		return
	}

	/* 아래 라인 부터는 무엇을 처리하느냐에 따라 다름 :) */
	var resp_msg string
	if len(rd.Url) <= 0 {
		resp_msg = "url is empty"
		send_400_response(w, resp_msg, 99002)
		return
	}

	if len(rd.User_id) <= 0 {
		resp_msg = "user_id is empty"
		send_400_response(w, resp_msg, 99003)
		return
	}

	if len(rd.User_guid) <= 0 {
		resp_msg = "user_guid is empty"
		send_400_response(w, resp_msg, 99004)
		return
	}

	if !(strings.HasSuffix(rd.Url, "/")) {
		rd.Url += "/"
	}

	host := rd.Url
	onedepth_url := rd.Url
	full_url := rd.Url
	location := "/"

	pos := 0
	found := 0
	tmp_url := rd.Url
	lastPos := strings.LastIndex(tmp_url, "/")
	if lastPos != -1 {
		break_condition := 0
		found = strings.Index(tmp_url, "/")
		for lastPos != found {
			pos = found + 1

			if break_condition == 0 {
				// host
				host = tmp_url[:pos]

				// location
				location += tmp_url[pos:]
			} else if break_condition == 1 {
				// onedepth url
				onedepth_url = tmp_url[:pos]
				fmt.Println(onedepth_url)
				break
			}

			tmp := tmp_url[pos:]
			fmt.Println(tmp)
			found = strings.Index(tmp, "/") + pos
			break_condition += 1
		}

		if lastPos == found {
			// onedepth url
			onedepth_url = tmp_url[:found+1]
		}
	}

	fmt.Printf("\n============ [EDLP I/F] request data ============\n")
	fmt.Printf("session: %s\n", r.RemoteAddr)
	fmt.Printf("url: %s\n", rd.Url)
	fmt.Printf("user_id: %s\n", rd.User_id)
	fmt.Printf("user_guid: %s\n", rd.User_guid)
	fmt.Printf("-------------------------------------------------\n")
	fmt.Printf("host: %s\n", host)
	fmt.Printf("onedept_url: %s\n", onedepth_url)
	fmt.Printf("full_url: %s\n", full_url)
	fmt.Printf("location: %s", location)
	fmt.Printf("\n============ [EDLP I/F] request data ============\n\n")

	send_200_response(w)
}

func StartServer(ip string, port int, location string) {
	// server info (ip, port)
	conn_info := ip + ":" + strconv.Itoa(port)

	// set location & handler
	callback := handleRequest
	http.HandleFunc(location, callback)

	// start tls server
	server_cert := "/home/bluetomorrow/go/src/GolangStudy/SideProject/https/certs/server.crt"
	server_key := "/home/bluetomorrow/go/src/GolangStudy/SideProject/https/certs/server.key"
	err := http.ListenAndServeTLS(conn_info, server_cert, server_key, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
