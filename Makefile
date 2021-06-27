.PHONY: info
info:
	@go run cmd/novel-grabber/main.go lightnovelworld info -n fantasy-simulator

.PHONY: epub/lightnovelworld
epub/lightnovelworld:
	@go run cmd/novel-grabber/main.go lightnovelworld -n tensei-shitara-slime-datta-ken-ln-17031322 -d ~/Documents/Books/Novels

.PHONY: epub/wuxiaworld
epub/wuxiaworld:
	@go run cmd/novel-grabber/main.go wuxiaworld info -n wuxiaworld
	@go run cmd/novel-grabber/main.go wuxiaworld -n martial-world -d ~/Documents/Books/Novels

.PHONY: link
link:
	@ln -s ~/.novel-grabber ./.novel-grabber

.PHONY: clean/chapter/unfinished
clean/chapter/unfinished:
	@./scripts/clean_chapter_unfinished.sh
