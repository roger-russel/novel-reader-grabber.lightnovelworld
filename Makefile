.PHONY: info
info:
	@go run cmd/novel-grabber/main.go lightnovelworld info -n fantasy-simulator

.PHONY: epub
epub:
	@go run cmd/novel-grabber/main.go lightnovelworld -n fantasy-simulator -d ~/Documents/Books/Novels

.PHONY: link
link:
	@ln -s ~/.novel-grabber ./.novel-grabber

.PHONY: clean/chapter/unfinished
clean/chapter/unfinished:
	@./scripts/clean_chapter_unfinished.sh
