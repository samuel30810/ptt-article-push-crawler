# 介紹
這個專案是ptt爬蟲，用來抓特定文章的推文，並透過TG機器人送給指定的用戶

# 目前功能
0. 使用setting.json來控制文章url、telegram bot token 和 telegram chat id
1. 開始時，抓下文章目前所有的推文，並透過TG機器人送給指定的用戶
2. 每十五分鐘確認是否有新推文，有的話透過TG機器人送給指定的用戶