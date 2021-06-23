.PHONY: info
info:
	@go run cmd/novel-grabber/main.go lightnovelworld info -n the-beginning-after-the-end-novel-09011258

.PHONY: epub
epub:
	@go run cmd/novel-grabber/main.go lightnovelworld -n the-beginning-after-the-end-novel-09011258 -d ~/Documents/Books/Novels

.PHONY: link
link:
	@ln -s ~/.novel-grabber ./.novel-grabber
