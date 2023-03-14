package client

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) Execute() {
	log.Print("starting client")
	res, err := c.getUSDBRL()
	if err != nil {
		log.Printf("Error at GetUSDBRL: %s", err.Error())
		return
	}
	log.Printf("USDBRL: %f", res.Bid)
	err = c.save(res)
	if err != nil {
		log.Printf("Error at Save: %s", err.Error())
		return
	}
	log.Print("finishing client")
}

func (c *Client) getUSDBRL() (*USDBRL, error) {
	ctx, cancel := context.WithTimeout(context.Background(), SERVER_TIMEOUT*time.Millisecond)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, SERVER_URL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed: %s", string(body))
	}
	var rate USDBRL
	err = json.Unmarshal(body, &rate)
	if err != nil {
		return nil, err
	}
	return &rate, nil
}

func (c *Client) save(rate *USDBRL) error {
	file, err := os.OpenFile(OUTPUT_FILENAME, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	tmp := template.New("CotacaoTemplate")
	tmp, _ = tmp.Parse("Dólar: {{.Bid}}\n")
	err = tmp.Execute(file, rate)
	if err != nil {
		return err
	}
	return nil
}
