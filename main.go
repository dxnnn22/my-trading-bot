package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	// هذا الجزء مهم جداً للخطة المجانية (Web Service)
	// يخبر السيرفر أن البوت "حي" ويعمل
	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Bot is Active and Monitoring Market!")
		})
		port := os.Getenv("PORT")
		if port == "" {
			port = "8080"
		}
		http.ListenAndServe(":"+port, nil)
	}()

	// بيانات الاتصال بموبايلك (سيتم جلبها من إعدادات Render)
	token := os.Getenv("8241919021:AAGfbRDmPUQpMnGTj1R0RmRWjI4K6rPE944")
	chatID := os.Getenv("830076775")

	fmt.Println("البوت بدأ العمل في السحاب...")

	// حلقة المراقبة اللحظية
	for {
		// ملاحظة للخبير: هنا سيتم جلب السعر الحقيقي يوم الاثنين
		fmt.Println("جاري فحص السوق وإرسال التحديث للخبير...")
		
		// إرسال رسالة تجريبية لتأكد من الربط
		sendToTelegram(token, chatID, "تم تحديث الكود! البوت الآن يعمل مجاناً على السحاب.")
		
		// الانتظار لمدة ساعة لكي لا تستهلك الخطة المجانية بسرعة
		time.Sleep(1 * time.Hour)
	}
}

// وظيفة الإرسال لتسهيل الكود
func sendToTelegram(token, chatID, text string) {
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%s&text=%s", token, chatID, text)
	http.Get(apiURL)
}
