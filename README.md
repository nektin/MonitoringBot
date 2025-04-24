Этот бот проверяет состояние веб-сайтов и уведомляет вас о любых обнаруженных проблемах.
Он периодически отправляет HTTP-запросы на указанные URL-адреса и сообщает, доступны они или нет. 
Если сайт становится недоступным или выдаёт ошибку, бот предупредит вас и предоставит подробную информацию.

# **Инструкция по установке на Linux**
### **1.Запуск программы на Linux**

После того как файл находится на сервере, выполните следующие шаги:

#### **Шаг 1: Сделайте файл исполняемым**

На Linux добавьте права на выполнение:
` chmod +x MonitoringBot`

#### **Шаг 2: Запустите программу**

Запустите программу вручную:
` ./MonitoringBot`

### **2.Настройка автозапуска через systemd**

Чтобы бот автоматически запускался при загрузке системы и мог быть легко управляем, создайте сервис для systemd.

#### **Шаг 1: Создайте файл сервиса**

Создайте файл сервиса в директории /etc/systemd/system/. Например, monitoringbot.service:
` sudo nano /etc/systemd/system/monitoringbot.service`
Добавьте следующее содержимое:
`   [Unit]
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
   WantedBy=multi-user.target`

Замените:
your_user и your_group на имя пользователя и группу, от имени которых будет запущен бот.
/path/to/your/app на путь к директории, где находится ваш исполняемый файл.
ENV_VAR_NAME=value на переменные окружения, если они нужны (например, токен бота).

#### **Шаг 2: Перезагрузите systemd**

После создания файла сервиса обновите конфигурацию systemd:
` sudo systemctl daemon-reload`

#### **Шаг 3: Запустите сервис**

Запустите сервис:
` sudo systemctl start monitoringbot`

#### **Шаг 4: Проверьте статус**

Проверьте, работает ли сервис:
`  sudo systemctl status monitoringbot`
Вы должны увидеть что-то вроде:
`   ● monitoringbot.service - Telegram Monitoring Bot
   Loaded: loaded (/etc/systemd/system/monitoringbot.service; enabled; vendor preset: enabled)
   Active: active (running) since Mon 2023-10-02 12:00:00 UTC; 1min ago
   Main PID: 1234 (MonitoringBot)
   Tasks: 5 (limit: 4915)
   Memory: 10.0M
   CGroup: /system.slice/monitoringbot.service
   └─1234 /path/to/your/app/MonitoringBot`

#### **Шаг 5: Включите автозапуск**

Чтобы сервис запускался автоматически при загрузке системы:
`sudo systemctl enable monitoringbot`