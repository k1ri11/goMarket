## Практическая работа 9

### выполнил студент группы ЭФМО-02-24 Трофимович Кирилл

### Таблица Покупатель (Customer)

CREATE TABLE Customer (
customer_id SERIAL PRIMARY KEY,
first_name VARCHAR(50) NOT NULL,
last_name VARCHAR(50) NOT NULL,
email VARCHAR(100) UNIQUE NOT NULL,
phone VARCHAR(20),
address TEXT,
city VARCHAR(50),
country VARCHAR(50),
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

### Таблица Товар (Product)

CREATE TABLE Product (
product_id SERIAL PRIMARY KEY,
name VARCHAR(100) NOT NULL,
brand VARCHAR(50),
model VARCHAR(50),
price DECIMAL(10, 2) NOT NULL,
stock INTEGER DEFAULT 0,
description TEXT,
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

### Таблица Заказ (Order)

CREATE TABLE "Order" (
order_id SERIAL PRIMARY KEY,
customer_id INT REFERENCES Customer(customer_id) ON DELETE CASCADE,
total_price DECIMAL(10, 2) NOT NULL,
status VARCHAR(50) NOT NULL,
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
shipped_at TIMESTAMP
);

### Таблица Позиции Заказа (Order_Item)

CREATE TABLE Order_Item (
order_item_id SERIAL PRIMARY KEY,
order_id INT REFERENCES "Order"(order_id) ON DELETE CASCADE,
product_id INT REFERENCES Product(product_id) ON DELETE CASCADE,
quantity INT NOT NULL,
price DECIMAL(10, 2) NOT NULL
);

### Таблица Оплаты (Payment)

CREATE TABLE Payment (
payment_id SERIAL PRIMARY KEY,
order_id INT REFERENCES "Order"(order_id) ON DELETE CASCADE,
amount DECIMAL(10, 2) NOT NULL,
payment_method VARCHAR(50) NOT NULL,
payment_status VARCHAR(50) DEFAULT 'pending',
payment_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

### Таблица Категория (Category)

CREATE TABLE Category (
category_id SERIAL PRIMARY KEY,
name VARCHAR(100) NOT NULL,
description TEXT
);

### Связующая таблица Товар-Категория (Product_Category)

CREATE TABLE Product_Category (
product_category_id SERIAL PRIMARY KEY,
product_id INT REFERENCES Product(product_id) ON DELETE CASCADE,
category_id INT REFERENCES Category(category_id) ON DELETE CASCADE
);

### Таблица Корзина (Shopping_Cart)

CREATE TABLE Shopping_Cart (
cart_id SERIAL PRIMARY KEY,
customer_id INT REFERENCES Customer(customer_id) ON DELETE CASCADE,
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

### Таблица Позиции в Корзине (Cart_Item)

CREATE TABLE Cart_Item (
cart_item_id SERIAL PRIMARY KEY,
cart_id INT REFERENCES Shopping_Cart(cart_id) ON DELETE CASCADE,
product_id INT REFERENCES Product(product_id) ON DELETE CASCADE,
quantity INT NOT NULL
);

### Таблица Отзывов (Review)

CREATE TABLE Review (
review_id SERIAL PRIMARY KEY,
product_id INT REFERENCES Product(product_id) ON DELETE CASCADE,
customer_id INT REFERENCES Customer(customer_id) ON DELETE CASCADE,
rating INT CHECK (rating >= 1 AND rating <= 5),
comment TEXT,
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

### Таблица Информации о Доставке (Shipping_Info)

CREATE TABLE Shipping_Info (
shipping_id SERIAL PRIMARY KEY,
order_id INT REFERENCES "Order"(order_id) ON DELETE CASCADE,
shipping_address TEXT NOT NULL,
shipping_method VARCHAR(50) NOT NULL,
shipping_cost DECIMAL(10, 2) NOT NULL,
estimated_delivery_date DATE,
actual_delivery_date DATE
);

### Таблица Ролей Пользователей (Role)

CREATE TABLE Role (
role_id SERIAL PRIMARY KEY,
name VARCHAR(50) NOT NULL,
description TEXT
);

### Таблица Связи Пользователь-Роль (User_Role)

CREATE TABLE User_Role (
user_role_id SERIAL PRIMARY KEY,
customer_id INT REFERENCES Customer(customer_id) ON DELETE CASCADE,
role_id INT REFERENCES Role(role_id) ON DELETE CASCADE
);

### Таблица Сессий Пользователей (User_Session)

CREATE TABLE User_Session (
session_id SERIAL PRIMARY KEY,
customer_id INT REFERENCES Customer(customer_id) ON DELETE CASCADE,
session_token TEXT NOT NULL,
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
expires_at TIMESTAMP
);

### Таблица Пользовательских Настроек (User_Settings)

CREATE TABLE User_Settings (
setting_id SERIAL PRIMARY KEY,
customer_id INT REFERENCES Customer(customer_id) ON DELETE CASCADE,
setting_key VARCHAR(50) NOT NULL,
setting_value TEXT
);

### Таблица Складских Запасов (Inventory)

CREATE TABLE Inventory (
inventory_id SERIAL PRIMARY KEY,
product_id INT REFERENCES Product(product_id) ON DELETE CASCADE,
quantity INT NOT NULL,
warehouse_location VARCHAR(100)
);

### Таблица Поставщиков (Supplier)

CREATE TABLE Supplier (
supplier_id SERIAL PRIMARY KEY,
name VARCHAR(100) NOT NULL,
contact_name VARCHAR(50),
email VARCHAR(100),
phone VARCHAR(20),
address TEXT
);

### Таблица Поставляемых Товаров (Supplier_Product)

CREATE TABLE Supplier_Product (
supplier_product_id SERIAL PRIMARY KEY,
supplier_id INT REFERENCES Supplier(supplier_id) ON DELETE CASCADE,
product_id INT REFERENCES Product(product_id) ON DELETE CASCADE,
supply_price DECIMAL(10, 2) NOT NULL,
quantity INT NOT NULL
);
