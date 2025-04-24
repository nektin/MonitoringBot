# 🚀 Telegram Monitoring Bot

Этот бот предназначен для мониторинга доступности веб-сайтов. Он периодически отправляет HTTP-запросы к указанным URL-адресам и уведомляет вас в Telegram о любых проблемах: если сайт становится недоступным или возвращает ошибку — вы получаете уведомление с подробностями.

---

## 📦 Возможности

- 🔄 Регулярная проверка состояния сайтов
- 📬 Мгновенные уведомления в Telegram
- ⚙️ Поддержка автозапуска через `systemd` на Linux

---

## 🐧 Установка и запуск на Linux

### 🔹 1. Запуск программы вручную

#### ✅ Шаг 1: Сделайте файл исполняемым

```bash
chmod +x MonitoringBot
```

#### 🚀 Шаг 2: Запустите программу

```bash
./MonitoringBot
```

---

### 🔹 2. Настройка автозапуска через `systemd`

Чтобы бот автоматически запускался при старте системы и легко управлялся, настройте сервис.

#### 📝 Шаг 1: Создайте файл сервиса

```bash
sudo nano /etc/systemd/system/monitoringbot.service
```

Добавьте следующее содержимое:

```ini
[Unit]
Description=Telegram Monitoring Bot
After=network.target

[Service]
Type=simple
User=your_user
Group=your_group
WorkingDirectory=/path/to/your/app
ExecStart=/path/to/your/app/MonitoringBot
Restart=always
RestartSec=5
Environment=ENV_VAR_NAME=value

[Install]
WantedBy=multi-user.target
```

> 🛠️ **Важно**:  
> - Замените `your_user` и `your_group` на имя пользователя и группу.  
> - Укажите реальный путь вместо `/path/to/your/app`.  
> - Добавьте переменные окружения, если нужно (например, токен Telegram-бота).

#### 🔄 Шаг 2: Перезагрузите systemd

```bash
sudo systemctl daemon-reload
```

#### ▶️ Шаг 3: Запустите сервис

```bash
sudo systemctl start monitoringbot
```

#### 📋 Шаг 4: Проверьте статус

```bash
sudo systemctl status monitoringbot
```

Пример вывода:

```
● monitoringbot.service - Telegram Monitoring Bot
   Loaded: loaded (/etc/systemd/system/monitoringbot.service; enabled; vendor preset: enabled)
   Active: active (running) since Mon 2023-10-02 12:00:00 UTC; 1min ago
   Main PID: 1234 (MonitoringBot)
   Tasks: 5 (limit: 4915)
   Memory: 10.0M
   CGroup: /system.slice/monitoringbot.service
   └─1234 /path/to/your/app/MonitoringBot
```

#### 🔁 Шаг 5: Включите автозапуск при старте системы

```bash
sudo systemctl enable monitoringbot
```

---

## 📫 Обратная связь

Если у вас возникли вопросы или предложения — не стесняйтесь обращаться!

---

## 🛡️ Лицензия

Этот проект распространяется под MIT License.