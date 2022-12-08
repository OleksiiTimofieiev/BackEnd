package main

import (
	"fmt"
	"io"
	"io/ioutil"

	// "log"
	"net/http"
	"net/url"
	"os/exec"
	"path"
	"strconv"
	"strings"
	"time"
	// tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	botToken          = "5963358002:AAHwHrtf0i464xW0O0nuk2MVhoEagNCiaeU"
	linuxCmd          = "iwgetid"
	linuxArgs         = "--raw"
	checkPeriod       = 15
	sendMessageMethod = "sendMessage"
	tgBotHost         = "api.telegram.org"
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
	_, err := client.Get("https://googlcom")

	if err != nil {
		return false
	}

	return true
}

func main() {
	status := false
	prevStatus := false

	// bot, err := tgbotapi.NewBotAPI(botToken)
	// if err != nil {
	// 	log.Panic(err)
	// }

	tg := New(tgBotHost, botToken)

	for {
		status = IsOnline()
		networkName := getNetworkName()
		if prevStatus != status {
			prevStatus = status

			if status {
				fmt.Println("Connected !!")
				fmt.Println(networkName)

				err := tg.SendMessage(1462377126, "Інтернет працює")
				if err != nil {
					fmt.Println(err)
				}
				err = tg.SendMessage(988002557, "Інтернет працює")
				if err != nil {
					fmt.Println(err)
				}
				// msg := tgbotapi.NewMessage(1462377126, "Інтернет працює")
				// msg2 := tgbotapi.NewMessage(988002557, "Інтернет працює")

				// if _, err := bot.Send(msg); err != nil {
				// 	log.Panic(err)
				// }
				// if _, err := bot.Send(msg2); err != nil {
				// 	log.Panic(err)
				// }
			}
		}
		time.Sleep(checkPeriod * time.Second)
	}
}
