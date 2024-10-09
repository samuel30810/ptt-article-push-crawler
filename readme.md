# 介紹
這個專案是ptt爬蟲，用來抓特定文章的推文，並透過TG機器人送給指定的用戶
預設情境是抓看板的至頂貼文或live文，如耳機版交易文、NBA比賽live文等

# 目前功能
0. 使用setting.json來控制文章url、telegram bot token、telegram chat id和檢查頻率
1. 開始時，抓下文章目前所有的推文，並透過TG機器人送給指定的用戶
2. 依據檢查頻率確認是否有新推文，有的話透過TG機器人送給指定的用戶