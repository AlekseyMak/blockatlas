package binance

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas"
	"net/http"
	"net/url"
	"strconv"
)

// TODO Headers + rate limiting

type Client struct {
	HTTPClient *http.Client
	BaseURL    string
	RPCBaseURL string
}

func (c *Client) GetBlockList(count int) (*BlockList, error) {
	uri := fmt.Sprintf("%s/blocks?page=1&rows=%d",
		c.BaseURL, count)

	res, err := c.HTTPClient.Get(uri)
	if err != nil {
		return nil, err
	}

	if err := getHTTPError(res, "GetBlockList"); err != nil {
		return nil, err
	}

	var blockList BlockList
	err = json.NewDecoder(res.Body).Decode(&blockList)
	if err != nil {
		return nil, err
	} else {
		return &blockList, nil
	}
}

func (c *Client) GetBlockByNumber(num int64) (*TxPage, error) {
	uri := fmt.Sprintf("%s/txs?%s",
		c.BaseURL,
		url.Values{
			"blockHeight": {strconv.FormatInt(num, 10)},
			// Only first 100 transactions of block returned
			// Shouldn't be a problem at the current transaction rate
			"rows":        {"100"},
			"page":        {"1"},
		}.Encode())

	res, err := c.HTTPClient.Get(uri)
	if err != nil {
		return nil, err
	}

	if err := getHTTPError(res, "GetBlockByNumber"); err != nil {
		return nil, err
	}

	stx := new(TxPage)
	err = json.NewDecoder(res.Body).Decode(stx)
	return stx, nil
}

func (c *Client) GetTxsOfAddress(address string, token string) (*TxPage, error) {
	uri := fmt.Sprintf("%s/txs?%s",
		c.BaseURL,
		url.Values{
			"address": {address},
			"rows":    {"100"},
			"page":    {"1"},
		}.Encode())

	res, err := c.HTTPClient.Get(uri)
	if err != nil {
		logrus.WithError(err).Error("Binance: Failed to get transactions")
		return nil, blockatlas.ErrSourceConn
	}

	if err := getHTTPError(res, "GetTxsOfAddress"); err != nil {
		return nil, err
	}

	stx := new(TxPage)
	err = json.NewDecoder(res.Body).Decode(stx)
	return stx, nil
}

func (c *Client) GetTransactionReceipt(hash string) (*Receipt, error) {
	url := fmt.Sprintf("%s/tx/%s?format=json", c.RPCBaseURL, hash)
	println(url)
	res, err := c.HTTPClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	recp := new(Receipt)
	err = json.NewDecoder(res.Body).Decode(recp)
	if err != nil {
		logrus.WithError(err).Error("Binance: Failed to decode transaction receipt API response")
		return nil, blockatlas.ErrSourceConn
	}

	return recp, nil

}

func getHTTPError(res *http.Response, desc string) error {
	switch res.StatusCode {
	case http.StatusBadRequest, http.StatusNotFound:
		return getAPIError(res, desc)
	case http.StatusOK:
		return nil
	default:
		return fmt.Errorf("%s", res.Status)
	}
}

func getAPIError(res *http.Response, desc string) error {
	var sErr Error
	err := json.NewDecoder(res.Body).Decode(&sErr)
	if err != nil {
		logrus.WithError(err).Error("Binance: Failed to get error")
		return blockatlas.ErrSourceConn
	}

	switch sErr.Message {
	case "address is not valid":
		return blockatlas.ErrInvalidAddr
	}

	logrus.WithFields(logrus.Fields {
		"status":  res.StatusCode,
		"code":    sErr.Code,
		"message": sErr.Message,
	}).Error("Binance: Failed to " + desc)
	return blockatlas.ErrSourceConn
}
