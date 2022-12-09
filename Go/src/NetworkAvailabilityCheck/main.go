package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os/exec"
	"path"
	"strconv"
	"strings"
	"time"
)

const (
	botToken          = ""
	linuxCmd          = "iwgetid"
	linuxArgs         = "--raw"
	checkPeriod       = 15
	sendMessageMethod = "sendMessage"
	tgBotHost         = "api.telegram.org"
	otimofie          = 1462377126
	ntimofie          = 988002557
	msg               = "Інтернет працює"
)

type Client struct {
	host     string
	basePath string
	client   http.Client
}

func New(host string, token string) *Client {
	return &Client{
		host:     host,
		basePath: newBasePath(token),
		client:   http.Client{},
	}
}

func Wrap(msg string, err error) error {
	return fmt.Errorf("%s: %w", msg, err)
}

func WrapIfErr(msg string, err error) error {
	if err == nil {
		return nil
	}
	return Wrap(msg, err)
}

func newBasePath(token string) string {
	return "bot" + token
}

func (c *Client) SendMessage(chatId int, text string) error {
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chatId))
	q.Add("text", text)

	_, err := c.doRequest(sendMessageMethod, q)

	if err != nil {
		return Wrap("can`t send message", err)
	}
	return nil
}

func (c *Client) doRequest(method string, query url.Values) (data []byte, err error) {
	defer func() { err = WrapIfErr("can`t process request", err) }()

	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.basePath, method),
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = query.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}
	return body, nil
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func getNetworkName() string {
	cmd := exec.Command(linuxCmd, linuxArgs)
	stdout, err := cmd.StdoutPipe()
	panicIf(err)

	// start the command after having set up the pipe
	if err := cmd.Start(); err != nil {
		panic(err)
	}
	defer cmd.Wait()

	var str string

	if b, err := ioutil.ReadAll(stdout); err == nil {
		str += (string(b) + "\n")
	}

	name := strings.Replace(str, "\n", "", -1)
	return name
}

func IsOnline() bool {
	timeout := time.Duration(5000 * time.Millisecond)
	client := http.Client{
		Timeout: timeout,
	}
	_, err := client.Get("https://google.com")

	if err != nil {
		return false
	}

	return true
}

func main() {
	status := false
	prevStatus := false

	tg := New(tgBotHost, botToken)

	for {
		status = IsOnline()

		if prevStatus != status {
			prevStatus = status

			if status {
				err := tg.SendMessage(otimofie, msg)
				if err != nil {
					fmt.Println(err)
				}
				err = tg.SendMessage(ntimofie, msg)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
		time.Sleep(checkPeriod * time.Second)
	}
}
