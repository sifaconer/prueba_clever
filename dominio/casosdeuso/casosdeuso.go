package casosdeuso

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func HomologarMoneda(from, to string, amount float64) (*float64, error) {
	access_key := os.Getenv("API_CURRENTLAYER_KEY")
	if access_key == "" {
		return nil, errors.New("no existe api key para https://currencylayer.com, configure la variable de entorno [API_CURRENTLAYER_KEY] con el valor de la key")
	}
	url := fmt.Sprintf("http://api.currencylayer.com/convert?access_key=%s&from=%s&to=%s&amount=%f&format=1", access_key, strings.ToUpper(from), strings.ToUpper(to), amount)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	json.Unmarshal(body, &result)

	if !result["success"].(bool) {
		return nil, errors.New(result["error"].(map[string]interface{})["info"].(string))
	}

	total := result["result"].(float64)
	return &total, nil
}
