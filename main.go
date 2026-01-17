package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os" // مكتبة os ستستخدم هنا لجلب التوكين من السيرفر
	"time"
)

type PriceResponse struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

func main() {
	// السيرفر سيقرأ هذه البيانات تلقائياً
	token := os.Getenv("8241919021:AAGfbRDmPUQpMnGTj1R0RmRWjI4K6rPE944")
	chatID := os.Getenv("830076775")

	for {
		resp, err := http.Get("https://api.binance.com/api/v3/ticker/price?symbol=BTCUSDT")
		if err != nil {
			time.Sleep(10 * time.Second)
			continue
		}
		
		body, _ := ioutil.ReadAll(resp.Body)
		var priceData PriceResponse
		json.Unmarshal(body, &priceData)
		resp.Body.Close()

		msg := "🚀 تحديث السوق للخبير:\nالسعر الحالي: " + priceData.Price
		apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%s&text=%s", 
			token, chatID, url.QueryEscape(msg))
		
		http.Get(apiURL)

		// التحديث كل ساعة لكي لا تستهلك الرصيد المجاني للسيرفر بسرعة
		time.Sleep(1 * time.Hour) 
	}
}