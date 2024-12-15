#!/bin/bash

# Параметры подключения к базе данных
DB_USER="postgres"
DB_PASSWORD="1234"
DB_NAME="online_store"
DB_HOST="0.0.0.0"
DB_PORT="5432"

# Проверка наличия флагов
CREATE_DB=0
FILL_DB=0

for arg in "$@"; do
  case $arg in
    --create-db)
      CREATE_DB=1
      ;;
    --fill-db)
      FILL_DB=1
      ;;
  esac
done

# Проверяем, доступен ли PostgreSQL
export PGPASSWORD="$DB_PASSWORD"
until psql -h "$DB_HOST" -U "$DB_USER" -p "$DB_PORT" -c '\q'; do
  echo "PostgreSQL не готов. Ждем..."
  sleep 2
done

echo "PostgreSQL готов! Выполняем скрипты..."

# Если флаг --create-db установлен, создаем базу данных
if [ "$CREATE_DB" -eq 1 ]; then
  echo "Создаем базу данных..."

  # Создание базы данных, если она не существует
  psql -h "$DB_HOST" -U "$DB_USER" -p "$DB_PORT" -c "CREATE DATABASE $DB_NAME;"
  echo "База данных $DB_NAME создана!"

  # Выполняем SQL-скрипт для создания таблиц
  echo "Создаем таблицы..."
  psql -h "$DB_HOST" -U "$DB_USER" -p "$DB_PORT" -d "$DB_NAME" -f ./scripts/createdb.sql
fi

# Если флаг заполнения базы данных установлен на 1, выполняем заполнение базы данных
if [ "$FILL_DB" -eq 1 ]; then
  echo "Заполняем базу данных..."

  # Вставка данных для всех таблиц (это пример, замените на ваши данные)
  psql -h "$DB_HOST" -U "$DB_USER" -p "$DB_PORT" -d "$DB_NAME" -f ./scripts/fill_db.sql

  echo "База данных заполнена!"
else
  echo "Заполнение базы данных не выполнено."
fi

echo "Скрипты выполнены успешно!"