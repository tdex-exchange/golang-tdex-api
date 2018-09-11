package tdex

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/pkg/errors"
)

func (as *apiService) WalletBalances(bq WalletBalancesRequest) (*WalletBalances, error) {
	params := make(map[string]interface{})
	params["type"] = bq.Type
	res, err := as.request("POST", "/wallet/balances", params, true, false)
	if err != nil {
		return nil, err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response from WalletBalances.post")
	}
	defer res.Body.Close()

	log.Println(string(textRes))
	if res.StatusCode != 200 {
		return nil, as.handleError(textRes)
	}

	rawWalletBalances := WalletBalances{}

	if err := json.Unmarshal(textRes, &rawWalletBalances); err != nil {
		return nil, errors.Wrap(err, "rawWalletBalances unmarshal failed")
	}

	return &rawWalletBalances, nil
}
func (as *apiService) WalletBalance(bq WalletBalanceRequest) (*WalletBalance, error) {
	params := make(map[string]interface{})
	params["currency"] = bq.Currency
	params["type"] = bq.Type

	res, err := as.request("POST", "/wallet/balance", params, true, false)
	if err != nil {
		return nil, err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response from WalletBalance.post")
	}
	defer res.Body.Close()

	log.Println(string(textRes))
	if res.StatusCode != 200 {
		return nil, as.handleError(textRes)
	}

	rawWalletBalance := WalletBalance{}

	if err := json.Unmarshal(textRes, &rawWalletBalance); err != nil {
		return nil, errors.Wrap(err, "rawWalletBalance unmarshal failed")
	}

	return &rawWalletBalance, nil
}

func (as *apiService) WalletWithdraw(wq WalletWithdrawRequest) (*WalletWithdraw, error) {
	params := make(map[string]interface{})
	params["currency"] = wq.Currency
	params["address"] = wq.Address
	params["amount"] = wq.Amount

	res, err := as.request("POST", "/wallet/withdraw", params, true, false)
	if err != nil {
		return nil, err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response from Walletwithdraw.post")
	}
	defer res.Body.Close()

	log.Println(string(textRes))
	if res.StatusCode != 200 {
		return nil, as.handleError(textRes)
	}

	rawWalletWithdraw := WalletWithdraw{}

	if err := json.Unmarshal(textRes, &rawWalletWithdraw); err != nil {
		return nil, errors.Wrap(err, "rawWalletWithdraw unmarshal failed")
	}

	return &rawWalletWithdraw, nil
}
func (as *apiService) WalletSwitch(wq WalletSwitchRequest) (*WalletSwitch, error) {
	params := make(map[string]interface{})
	params["currency"] = wq.Currency
	params["direction"] = wq.Direction
	params["amount"] = wq.Amount

	res, err := as.request("POST", "/wallet/switch", params, true, false)
	if err != nil {
		return nil, err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response from WalletSwitch.post")
	}
	defer res.Body.Close()

	log.Println(string(textRes))
	if res.StatusCode != 200 {
		return nil, as.handleError(textRes)
	}

	rawWalletSwitch := WalletSwitch{}

	if err := json.Unmarshal(textRes, &rawWalletSwitch); err != nil {
		return nil, errors.Wrap(err, "rawWalletSwitch unmarshal failed")
	}

	return &rawWalletSwitch, nil
}
