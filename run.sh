#!/bin/bash

cd "$(dirname "$0")"
echo "🚀 Запуск рассылки: $(date)" >> logs/email_sender.log
./email_sender >> logs/email_sender.log 2>&1
