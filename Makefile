# FOR DEVELOPMENT

# just build dev project version
dev-build:
	docker compose -f ./docker-compose.dev.yml build
# just up dev project version
dev-up:
	docker compose -f ./docker-compose.dev.yml up
# just up dev project version in background
dev-upd:
	docker compose -f ./docker-compose.dev.yml up -d
# just down dev project version
dev-down:
	docker compose -f ./docker-compose.dev.yml down
# just ps dev project version
dev-ps:
	docker compose -f ./docker-compose.dev.yml ps


# full restart dev project version in background
dev-full: dev-down dev-build dev-upd
	@sleep 3
	@docker compose -f ./docker-compose.dev.yml ps


# FOR TG BOT

# just build tg_bot project version
tg_bot-build:
	docker compose -f ./docker-compose.tg_bot.yml build
# just up tg_bot project version
tg_bot-up:
	docker compose -f ./docker-compose.tg_bot.yml up
# just up tg_bot project version in background
tg_bot-upd:
	docker compose -f ./docker-compose.tg_bot.yml up -d
# just down tg_bot project version
tg_bot-down:
	docker compose -f ./docker-compose.tg_bot.yml down
# just ps tg_bot project version
tg_bot-ps:
	docker compose -f ./docker-compose.tg_bot.yml ps

# full restart tg_bot project version in background
tg_bot-full: tg_bot-down tg_bot-build tg_bot-upd
	@sleep 3
	@docker compose -f ./docker-compose.tg_bot.yml ps
