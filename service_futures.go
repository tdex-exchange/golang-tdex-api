package tdex

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/pkg/errors"
)

func (as *apiService) FuturesOpen(wq FuturesOpenRequest) (*FuturesOpen, error) {
	params := make(map[string]interface{})
	params["cid"] = wq.Cid
	params["side"] = wq.Side
	params["scale"] = wq.Scale
	params["volume"] = wq.Volume
	params["distance"] = wq.Distance
	params["price"] = wq.Price
	params["timely"] = wq.Timely
	params["timelyParam"] = wq.TimelyParam
	params["passive"] = wq.Passive
	params["visible"] = wq.Visible
	params["strategy"] = wq.Strategy
	params["better"] = wq.Better
	params["variable"] = wq.Variable
	params["constant"] = wq.Constant

	res, err := as.request("POST", "/futures/open", params, true, false)
	if err != nil {
		return nil, err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response from FuturesOpen.post")
	}
	defer res.Body.Close()

	log.Println(string(textRes))
	if res.StatusCode != 200 {
		return nil, as.handleError(textRes)
	}

	rawFuturesOpen := FuturesOpen{}

	if err := json.Unmarshal(textRes, &rawFuturesOpen); err != nil {
		return nil, errors.Wrap(err, "rawFuturesOpen unmarshal failed")
	}

	return &rawFuturesOpen, nil
}
func (as *apiService) FuturesClose(wq []FuturesCloseRequest) (*FuturesClose, error) {
	params := make(map[string]interface{})
	params["list"] = wq

	res, err := as.request("POST", "/futures/close", params, true, false)
	if err != nil {
		return nil, err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response from FuturesClose.post")
	}
	defer res.Body.Close()

	log.Println(string(textRes))
	if res.StatusCode != 200 {
		return nil, as.handleError(textRes)
	}

	rawFuturesClose := FuturesClose{}

	if err := json.Unmarshal(textRes, &rawFuturesClose); err != nil {
		return nil, errors.Wrap(err, "rawFuturesClose unmarshal failed")
	}

	return &rawFuturesClose, nil
}
func (as *apiService) FuturesCloseAll(wq []FuturesCloseAllRequest) (*FuturesCloseAll, error) {
	params := make(map[string]interface{})
	params["list"] = wq

	res, err := as.request("POST", "/futures/closeAll", params, true, false)
	if err != nil {
		return nil, err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response from FuturesCloseAll.post")
	}
	defer res.Body.Close()

	log.Println(string(textRes))
	if res.StatusCode != 200 {
		return nil, as.handleError(textRes)
	}

	rawFuturesCloseAll := FuturesCloseAll{}

	if err := json.Unmarshal(textRes, &rawFuturesCloseAll); err != nil {
		return nil, errors.Wrap(err, "rawFuturesCloseAll unmarshal failed")
	}

	return &rawFuturesCloseAll, nil
}
func (as *apiService) FuturesCancel(wq []FuturesCancelRequest) (*FuturesCancel, error) {
	params := make(map[string]interface{})
	params["list"] = wq

	res, err := as.request("POST", "/futures/cancel", params, true, false)
	if err != nil {
		return nil, err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response from FuturesCancel.post")
	}
	defer res.Body.Close()

	log.Println(string(textRes))
	if res.StatusCode != 200 {
		return nil, as.handleError(textRes)
	}

	rawFuturesCancel := FuturesCancel{}

	if err := json.Unmarshal(textRes, &rawFuturesCancel); err != nil {
		return nil, errors.Wrap(err, "rawFuturesCancel unmarshal failed")
	}

	return &rawFuturesCancel, nil
}

func (as *apiService) FuturesSetsl(wq FuturesSetslRequest) (*FuturesSetsl, error) {
	params := make(map[string]interface{})
	params["cid"] = wq.Cid
	params["id"] = wq.ID
	params["distance"] = wq.Distance
	params["price"] = wq.Price
	params["timely"] = wq.Timely
	params["timelyParam"] = wq.TimelyParam
	params["strategy"] = wq.Strategy
	params["variable"] = wq.Variable
	params["constant"] = wq.Constant
	params["passive"] = wq.Passive
	params["visible"] = wq.Visible
	params["better"] = wq.Better

	res, err := as.request("POST", "/futures/setsl", params, true, false)
	if err != nil {
		return nil, err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response from FuturesSetsl.post")
	}
	defer res.Body.Close()

	log.Println(string(textRes))
	if res.StatusCode != 200 {
		return nil, as.handleError(textRes)
	}

	rawFuturesSetsl := FuturesSetsl{}

	if err := json.Unmarshal(textRes, &rawFuturesSetsl); err != nil {
		return nil, errors.Wrap(err, "rawFuturesSetsl unmarshal failed")
	}

	return &rawFuturesSetsl, nil
}
func (as *apiService) FuturesSettp(wq FuturesSettpRequest) (*FuturesSettp, error) {
	params := make(map[string]interface{})
	params["cid"] = wq.Cid
	params["id"] = wq.ID
	params["distance"] = wq.Distance
	params["price"] = wq.Price
	params["timely"] = wq.Timely
	params["timelyParam"] = wq.TimelyParam
	params["strategy"] = wq.Strategy
	params["variable"] = wq.Variable
	params["constant"] = wq.Constant
	params["passive"] = wq.Passive
	params["visible"] = wq.Visible
	params["better"] = wq.Better

	res, err := as.request("POST", "/futures/settp", params, true, false)
	if err != nil {
		return nil, err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response from FuturesSettp.post")
	}
	defer res.Body.Close()

	log.Println(string(textRes))
	if res.StatusCode != 200 {
		return nil, as.handleError(textRes)
	}

	rawFuturesSettp := FuturesSettp{}

	if err := json.Unmarshal(textRes, &rawFuturesSettp); err != nil {
		return nil, errors.Wrap(err, "rawFuturesSettp unmarshal failed")
	}

	return &rawFuturesSettp, nil
}

func (as *apiService) FuturesMerge(wq FuturesMergeRequest) (*FuturesMerge, error) {
	params := make(map[string]interface{})
	params["cid"] = wq.Cid
	params["list"] = wq.List

	res, err := as.request("POST", "/futures/merge", params, true, false)
	if err != nil {
		return nil, err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response from FuturesMerge.post")
	}
	defer res.Body.Close()

	log.Println(string(textRes))
	if res.StatusCode != 200 {
		return nil, as.handleError(textRes)
	}

	rawFuturesMerge := FuturesMerge{}

	if err := json.Unmarshal(textRes, &rawFuturesMerge); err != nil {
		return nil, errors.Wrap(err, "rawFuturesMerge unmarshal failed")
	}

	return &rawFuturesMerge, nil
}
func (as *apiService) FuturesSplit(wq FuturesSplitRequest) (*FuturesSplit, error) {
	params := make(map[string]interface{})
	params["cid"] = wq.Cid
	params["id"] = wq.ID
	params["volume"] = wq.Volume

	res, err := as.request("POST", "/futures/split", params, true, false)
	if err != nil { 
		return nil, err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response from FuturesSplit.post")
	}
	defer res.Body.Close()

	log.Println(string(textRes))
	if res.StatusCode != 200 {
		return nil, as.handleError(textRes)
	}

	rawFuturesSplit := FuturesSplit{}

	if err := json.Unmarshal(textRes, &rawFuturesSplit); err != nil {
		return nil, errors.Wrap(err, "rawFuturesSplit unmarshal failed")
	}

	return &rawFuturesSplit, nil
}
func (as *apiService) FuturesSetup(wq FuturesSetupRequest) (*FuturesSetup, error) {
	params := make(map[string]interface{})
	params["cid"] = wq.Cid
	params["id"] = wq.ID
	params["target"] = wq.Target

	res, err := as.request("POST", "/futures/setup", params, true, false)
	if err != nil {
		return nil, err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response from FuturesSetup.post")
	}
	defer res.Body.Close()

	log.Println(string(textRes))
	if res.StatusCode != 200 {
		return nil, as.handleError(textRes)
	}

	rawFuturesSetup := FuturesSetup{}

	if err := json.Unmarshal(textRes, &rawFuturesSetup); err != nil {
		return nil, errors.Wrap(err, "rawFuturesSetup unmarshal failed")
	}

	return &rawFuturesSetup, nil
}
func (as *apiService) FuturesSchemeSet(wq FuturesSchemeSetRequest) (*FuturesSchemeSet, error) {
	params := make(map[string]interface{})
	params["cid"] = wq.Cid
	params["options"] = wq.Options

	res, err := as.request("PUT", "/futures/scheme", params, true, false)
	if err != nil {
		return nil, err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response from FuturesSchemeSet.post")
	}
	defer res.Body.Close()

	log.Println(string(textRes))
	if res.StatusCode != 200 {
		return nil, as.handleError(textRes)
	}

	rawFuturesSchemeSet := FuturesSchemeSet{}

	if err := json.Unmarshal(textRes, &rawFuturesSchemeSet); err != nil {
		return nil, errors.Wrap(err, "rawFuturesSchemeSet unmarshal failed")
	}

	return &rawFuturesSchemeSet, nil
}
func (as *apiService) FuturesSchemeGet(wq FuturesSchemeGetRequest) (*FuturesSchemeGet, error) {
	params := make(map[string]interface{})
	params["cid"] = wq.Cid

	res, err := as.request("POST", "/futures/scheme", params, true, false)
	if err != nil {
		return nil, err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response from FuturesSchemeGet.post")
	}
	defer res.Body.Close()

	log.Println(string(textRes))
	if res.StatusCode != 200 {
		return nil, as.handleError(textRes)
	}

	rawFuturesSchemeGet := FuturesSchemeGet{}

	if err := json.Unmarshal(textRes, &rawFuturesSchemeGet); err != nil {
		return nil, errors.Wrap(err, "rawFuturesSchemeGet unmarshal failed")
	}

	return &rawFuturesSchemeGet, nil
}
func (as *apiService) FuturesOrders() (*FuturesOrders, error) {
	params := make(map[string]interface{})

	res, err := as.request("POST", "/futures/orders", params, true, false)
	if err != nil {
		return nil, err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response from FuturesOrders.post")
	}
	defer res.Body.Close()

	log.Println(string(textRes))
	if res.StatusCode != 200 {
		return nil, as.handleError(textRes)
	}

	rawFuturesOrders := FuturesOrders{}

	if err := json.Unmarshal(textRes, &rawFuturesOrders); err != nil {
		return nil, errors.Wrap(err, "rawFuturesOrders unmarshal failed")
	}

	return &rawFuturesOrders, nil
}
func (as *apiService) FuturesPosition() (*FuturesPosition, error) {
	params := make(map[string]interface{})

	res, err := as.request("POST", "/futures/position", params, true, false)
	if err != nil {
		return nil, err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response from FuturesPosition.post")
	}
	defer res.Body.Close()

	log.Println(string(textRes))
	if res.StatusCode != 200 {
		return nil, as.handleError(textRes)
	}

	rawFuturesPosition := FuturesPosition{}

	if err := json.Unmarshal(textRes, &rawFuturesPosition); err != nil {
		return nil, errors.Wrap(err, "rawFuturesPosition unmarshal failed")
	}

	return &rawFuturesPosition, nil
}
func (as *apiService) FuturesHistory(wq FuturesHistoryRequest) (*FuturesHistory, error) {
	params := make(map[string]interface{})
	params["page"] = wq.Page
	params["pageSize"] = wq.PageSize

	res, err := as.request("POST", "/futures/history", params, true, false)
	if err != nil {
		return nil, err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response from FuturesHistory.post")
	}
	defer res.Body.Close()

	log.Println(string(textRes))
	if res.StatusCode != 200 {
		return nil, as.handleError(textRes)
	}

	rawFuturesHistory := FuturesHistory{}

	if err := json.Unmarshal(textRes, &rawFuturesHistory); err != nil {
		return nil, errors.Wrap(err, "rawFuturesHistory unmarshal failed")
	}

	return &rawFuturesHistory, nil
}
func (as *apiService) FuturesContract(wq FuturesContractRequest) (*FuturesContract, error) {
	params := make(map[string]interface{})
	params["symbol"] = wq.Symbol

	res, err := as.request("POST", "/futures/contract", params, true, false)
	if err != nil {
		return nil, err
	}
	textRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response from FuturesContract.post")
	}
	defer res.Body.Close()

	log.Println(string(textRes))
	if res.StatusCode != 200 {
		return nil, as.handleError(textRes)
	}

	rawFuturesContract := FuturesContract{}

	if err := json.Unmarshal(textRes, &rawFuturesContract); err != nil {
		return nil, errors.Wrap(err, "rawFuturesContract unmarshal failed")
	}

	return &rawFuturesContract, nil
}
