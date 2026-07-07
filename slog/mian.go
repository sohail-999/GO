package main

import "github.com/bytedance/gopkg/util/logger"

func main() {

	//	log.Print("log message....")
	//slog.Info("log message...")
	logger.Info(
		"incoming request \n",
		"method", "GET \n",
		"time_taken_ms ", 158, "\n",
		"path", "/hello/world?q=search \n",
		"status ", 200, "\n",
		"user_agent \n", "Googlebot/2.1 (+http://www.google.com/bot.html)",
		//true,
	)

}
