Готово. Ниже — **готовый `README.md` с упором именно на запуск и работу в фоне через systemd и `.env`**. Можно копировать целиком.

---

```md
# T0T0D Cyber Telegram Bot

Telegram-бот на Go с модульной архитектурой, health-check и graceful shutdown.

---

## 📦 Требования

- Linux (Ubuntu/Debian)
- Go ≥ версии, указанной в `go.mod`
- systemd
- Telegram Bot Token

---

## 📂 Структура проекта

```

.
├── cmd/bot            # entrypoint бота
├── internal           # бизнес-логика
├── go.mod
├── go.sum
└── .env               # конфигурация (НЕ коммитится)

````

Точка входа:  
`cmd/bot/main.go`

---

## ⚙️ Конфигурация

Все настройки задаются **через `.env` файл**.

### `.env`
Создай файл:

```bash
nano .env
````

Пример:

```env
BOT_TOKEN=123456:ABCDEF
BOT_ENV=production
LOG_LEVEL=info
```

⚠️ Формат строго `KEY=value`, без пробелов.

---

## ▶️ Запуск локально (проверка)

```bash
cd /tgbots/0t0d/t0t0dcyberbot
go mod tidy
go run ./cmd/bot
```

Или:

```bash
export $(cat .env | xargs)
go run ./cmd/bot
```

---

## 🏗 Сборка бинарника

```bash
go build -o t0t0dcyberbot ./cmd/bot
```

Проверка:

```bash
./t0t0dcyberbot
```

---

## 🔁 Запуск в фоне (systemd)

### 1️⃣ systemd unit

```bash
nano /etc/systemd/system/t0t0dcyberbot.service
```

```ini
[Unit]
Description=T0T0D Cyber Telegram Bot
After=network.target

[Service]
Type=simple
User=root
WorkingDirectory=/tgbots/0t0d/t0t0dcyberbot
EnvironmentFile=/tgbots/0t0d/t0t0dcyberbot/.env
ExecStart=/tgbots/0t0d/t0t0dcyberbot/t0t0dcyberbot
Restart=always
RestartSec=5

StandardOutput=journal
StandardError=journal

[Install]
WantedBy=multi-user.target
```

---

### 2️⃣ Запуск сервиса

```bash
systemctl daemon-reload
systemctl enable t0t0dcyberbot
systemctl start t0t0dcyberbot
```

---

### 3️⃣ Проверка

Статус:

```bash
systemctl status t0t0dcyberbot
```

Логи:

```bash
journalctl -u t0t0dcyberbot -f
```

---

## 🩺 Health Check

Если включён HTTP health endpoint (см. `cmd/bot/health.go`):

```bash
curl http://localhost:8080/health
```

---

## 🔄 Перезапуск

```bash
systemctl restart t0t0dcyberbot
```

---

## 🛑 Остановка

```bash
systemctl stop t0t0dcyberbot
```

Graceful shutdown обрабатывается (см. `shutdown.go`).

---

## ❗ Частые проблемы

### `.env` не читается

```bash
chmod 600 .env
```

### Проверка env вручную

```bash
export $(cat .env | xargs)
./t0t0dcyberbot
```

---

## 📌 Примечания

* Бот рассчитан на запуск через **polling**
* Все настройки централизованы в `.env`
* Логи идут в `journalctl`
* systemd автоматически перезапускает процесс при падении

---

## 📜 License

Private / Internal

```
