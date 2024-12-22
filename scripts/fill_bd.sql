-- Таблица Customer
INSERT INTO Customer (first_name, last_name, email, phone, address, city, country)
VALUES
    ('John', 'Doe', 'john.doe@example.com', '+123456789', '123 Main St', 'Springfield', 'USA'),
    ('Jane', 'Smith', 'jane.smith@example.com', '+987654321', '456 Oak Ave', 'Shelbyville', 'USA'),
    ('Emily', 'Johnson', 'emily.j@example.com', '+112233445', '789 Pine Rd', 'Capital City', 'Canada');

-- Таблица Product
INSERT INTO Product (name, brand, model, price, stock, description)
VALUES
    ('Smartphone', 'TechBrand', 'X100', 799.99, 50, 'High-end smartphone with excellent features'),
    ('Laptop', 'ComputePro', 'L200', 1199.99, 30, 'Lightweight laptop with powerful performance'),
    ('Headphones', 'SoundWave', 'H300', 199.99, 100, 'Wireless headphones with noise cancellation'),
    ('Tablet', 'TechBrand', 'T500', 499.99, 20, 'Portable tablet with high-resolution display'),
    ('Smartwatch', 'WearTime', 'W100', 249.99, 40, 'Smartwatch with fitness tracking'),
    ('Gaming Console', 'GameZone', 'G1', 299.99, 15, 'Popular gaming console with immersive experience'),
    ('Bluetooth Speaker', 'SoundWave', 'S100', 99.99, 60, 'Portable speaker with excellent sound quality'),
    ('External Hard Drive', 'StoragePro', 'HDD1TB', 79.99, 25, '1TB external hard drive for secure data storage');

-- Таблица Category
INSERT INTO Category (name, description)
VALUES
    ('Electronics', 'Devices and gadgets for everyday use'),
    ('Computers', 'Laptops, desktops, and peripherals'),
    ('Accessories', 'Complementary products for electronics'),
    ('Wearables', 'Smartwatches and other wearable technology'),
    ('Gaming', 'Gaming consoles and accessories'),
    ('Storage', 'External drives and storage devices');

-- Таблица Product_Category
INSERT INTO Product_Category (product_id, category_id)
VALUES
    (1, 1), (1, 3), (2, 2), (3, 3),
    (4, 1), (4, 3), (5, 1), (5, 2), (6, 3), (7, 1), (7, 2), (8, 1);

-- Таблица Order
INSERT INTO "order" (customer_id, total_price, status)
VALUES
    (1, 999.98, 'Pending'),
    (2, 199.99, 'Shipped'),
    (3, 749.97, 'Processing'),
    (1, 1399.95, 'Pending');

-- Таблица Order_Item
INSERT INTO Order_Item (order_id, product_id, quantity, price)
VALUES
    (1, 1, 1, 799.99), (1, 3, 1, 199.99), (2, 3, 1, 199.99),
    (3, 1, 1, 799.99), (3, 3, 2, 199.99), (4, 2, 1, 1199.99), (4, 6, 5, 99.99);

-- Таблица Payment
INSERT INTO Payment (order_id, amount, payment_method, payment_status)
VALUES
    (1, 999.98, 'Credit Card', 'Completed'),
    (2, 199.99, 'PayPal', 'Completed');

-- Таблица Review
INSERT INTO Review (product_id, customer_id, rating, comment)
VALUES
    (1, 1, 5, 'Amazing smartphone, highly recommend!'),
    (2, 2, 4, 'Good laptop, but battery life could be better.');

-- Таблица Shipping_Info
INSERT INTO Shipping_Info (order_id, shipping_address, shipping_method, shipping_cost, estimated_delivery_date)
VALUES
    (1, '123 Main St, Springfield, USA', 'Standard', 9.99, '2024-12-05'),
    (2, '456 Oak Ave, Shelbyville, USA', 'Express', 19.99, '2024-12-03');

-- Таблица Role
INSERT INTO Role (name, description)
VALUES
    ('Admin', 'Administrator with full access'),
    ('Customer', 'Regular user with shopping privileges');

-- Таблица User_Role
INSERT INTO User_Role (customer_id, role_id)
VALUES
    (1, 2), (2, 2);

-- Таблица Inventory
INSERT INTO Inventory (product_id, quantity, warehouse_location)
VALUES
    (1, 50, 'Warehouse A'),
    (2, 30, 'Warehouse B'),
    (3, 100, 'Warehouse C');

-- Таблица Supplier
INSERT INTO Supplier (name, contact_name, email, phone, address)
VALUES
    ('Global Supplies Co.', 'Alice Brown', 'alice.b@example.com', '+555123456', '12 Industry Park, Big City'),
    ('TechParts Ltd.', 'Bob Green', 'bob.g@example.com', '+555987654', '45 Tech Drive, Tech City');

-- Таблица Supplier_Product
INSERT INTO Supplier_Product (supplier_id, product_id, supply_price, quantity)
VALUES
    (1, 1, 700.00, 20), (2, 2, 1000.00, 10), (1, 3, 150.00, 50);

-- Добавление новых корзин
INSERT INTO Shopping_Cart (customer_id)
VALUES
    (1), (2), (3);

-- Добавление позиций в корзины
INSERT INTO Cart_Item (cart_id, product_id, quantity)
VALUES
    (1, 4, 1), (1, 5, 2), (2, 6, 1), (2, 7, 1), (3, 8, 3);