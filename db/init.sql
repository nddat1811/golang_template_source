-- Bảng SYS_FUNCTION
CREATE TABLE "SYS_FUNCTION" (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    path VARCHAR(255) NOT NULL,
    regex VARCHAR(255),
    description TEXT,
    parent_id INT,
    type VARCHAR(50),
    status VARCHAR(50) NOT NULL,
    icon_url VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    created_by INT,
    updated_by INT
);

-- Bảng SYS_ROLE
CREATE TABLE "SYS_ROLE" (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    description TEXT,
    status VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by INT,
    updated_by INT
);

-- Bảng SYS_ROLE_FUNCTION
CREATE TABLE "SYS_ROLE_FUNCTION" (
    id SERIAL PRIMARY KEY,
    role_id INT NOT NULL,
    function_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    created_by INT,
    updated_by INT,
    FOREIGN KEY (role_id) REFERENCES "SYS_ROLE"(id) ON DELETE CASCADE,
    FOREIGN KEY (function_id) REFERENCES "SYS_FUNCTION"(id) ON DELETE CASCADE
);

-- Bảng SYS_USER
CREATE TABLE "SYS_USER" (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    "identity" VARCHAR(255),
    hash_password VARCHAR(255) NOT NULL,
    phone VARCHAR(15),
    full_name VARCHAR(255) NOT NULL,
    status VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    created_by INT,
    updated_by INT
);

-- Bảng SYS_USER_ROLE
CREATE TABLE "SYS_USER_ROLE" (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    role_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    created_by INT,
    updated_by INT,
    FOREIGN KEY (user_id) REFERENCES "SYS_USER"(id) ON DELETE CASCADE,
    FOREIGN KEY (role_id) REFERENCES "SYS_ROLE"(id) ON DELETE CASCADE
);

-- Thêm các function vào bảng SYS_FUNCTION
INSERT INTO "SYS_FUNCTION" (name, path, description, type, status, created_by, regex)
VALUES 
('Get All Users', '/users', 'Lấy danh sách tất cả người dùng', 'API', 'ACTIVE', 1, NULL),
('Get User Detail', '/users/:id', 'Lấy thông tin chi tiết của người dùng', 'API', 'ACTIVE', 1, '^/users/[a-zA-Z0-9]+$');

-- Thêm các role vào bảng SYS_ROLE
INSERT INTO "SYS_ROLE" (name, description, status, created_by)
VALUES 
('Admin', 'Quản trị viên hệ thống', 'ACTIVE', 1),
('User', 'Người dùng thông thường', 'ACTIVE', 1);

-- Thêm người dùng vào bảng SYS_USER
-- INSERT INTO "SYS_USER" (full_name, email, phone, hash_password)
-- VALUES 
-- ('Admin', 'admin@gmail.com', '0123456789', '$2a$10$j3u/NoMXCHGaAQX92unfHeymF4N0foLDeBqLP7N8wjY9S/gixulG6'),
-- ('User', 'user@gmail.com', '0987654321', '$2a$10$j3u/NoMXCHGaAQX92unfHeymF4N0foLDeBqLP7N8wjY9S/gixulG6');
-- 123456
INSERT INTO "SYS_USER" (email, identity, hash_password, phone, full_name, status, created_by, updated_by) 
VALUES
    ('admin@gmail.com', '045201007111', '$2a$10$2OsbZhkEJn5keRCQr.jfoeCOoqOus.oqnXFdiGL/wbX5God0wSwSa', '0987654321', 'Admin', 'ACTIVE', 1, 1),
    ('affiliate@gmail.com', '045201007112', '$2a$10$2OsbZhkEJn5keRCQr.jfoeCOoqOus.oqnXFdiGL/wbX5God0wSwSa', '0977654321', 'Affiliate', 'ACTIVE', 1, 1),
    ('user@gmail.com', '045201007113', '$2a$10$2OsbZhkEJn5keRCQr.jfoeCOoqOus.oqnXFdiGL/wbX5God0wSwSa', '0967654321', 'User', 'ACTIVE', 1, 1);


-- Phân quyền cho người dùng
-- Admin được gán quyền Admin
INSERT INTO "SYS_USER_ROLE" (user_id, role_id, created_by)
VALUES 
((SELECT id FROM "SYS_USER" WHERE email = 'admin@gmail.com'), 
(SELECT id FROM "SYS_ROLE" WHERE name = 'Admin'), 1);

-- User được gán quyền User
INSERT INTO "SYS_USER_ROLE" (user_id, role_id, created_by)
VALUES 
((SELECT id FROM "SYS_USER" WHERE email = 'user@gmail.com'), 
 (SELECT id FROM "SYS_ROLE" WHERE name = 'User'), 1);


INSERT INTO "SYS_ROLE_FUNCTION" (function_id, role_id, created_by)
VALUES 
((SELECT id FROM "SYS_FUNCTION" WHERE name = 'Get All Users'), 
(SELECT id FROM "SYS_ROLE" WHERE name = 'Admin'), 1);

-- User được gán quyền User
INSERT INTO "SYS_ROLE_FUNCTION" (function_id, role_id, created_by)
VALUES 
((SELECT id FROM "SYS_FUNCTION" WHERE name = 'Get All Users'), 
 (SELECT id FROM "SYS_ROLE" WHERE name = 'User'), 1);


INSERT INTO "SYS_ROLE_FUNCTION" (function_id, role_id, created_by)
VALUES 
((SELECT id FROM "SYS_FUNCTION" WHERE name = 'Get User Detail'), 
(SELECT id FROM "SYS_ROLE" WHERE name = 'Admin'), 1);



CREATE TABLE "USER_DOC" (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    id_card_front INT NOT NULL,
    id_card_back INT NOT NULL,
    portrait_photo INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    created_by INT,
    updated_by INT,
    FOREIGN KEY ("user_id") REFERENCES "SYS_USER"("id") ON DELETE CASCADE
);

CREATE TABLE "USER_CHANGE_HISTORY" (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    "change_type" VARCHAR(255) NOT NULL,
    status VARCHAR(50) NOT NULL,
    approval_time TIMESTAMP,
    approval_id INT NULL,
    note TEXT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    created_by INT,
    updated_by INT,
    FOREIGN KEY ("user_id") REFERENCES "SYS_USER"("id") ON DELETE CASCADE
);

CREATE TABLE "USER_DETAIL_CHANGE" (
    id SERIAL PRIMARY KEY,
    user_change_history_id INT NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    old_value TEXT,
    new_value TEXT,
    FOREIGN KEY ("user_change_history_id") REFERENCES "USER_CHANGE_HISTORY"("id") ON DELETE CASCADE
);

CREATE TABLE "USER_PAYMENT" (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    account_number VARCHAR(255) NOT NULL,
    bank_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    created_by INT,
    updated_by INT,
    FOREIGN KEY ("user_id") REFERENCES "SYS_USER"("id") ON DELETE CASCADE
);


CREATE TABLE "SYS_FILE" (
    id SERIAL PRIMARY KEY,
    share_link TEXT NULL,
    "type" VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    created_by INT NULL
);



CREATE TABLE "PACKAGE" (
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL,
    "price" FLOAT NOT NULL,
    "status" VARCHAR(50) NOT NULL,
    "affiliate_commission" FLOAT NULL,
    "mobifone_commission" FLOAT NULL,
    "agency_commission" FLOAT NULL,
    "cycle" VARCHAR(50) NULL,
    "priority_product" BOOLEAN NOT NULL DEFAULT FALSE,
    "benefit" TEXT NULL,
    "condition" TEXT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    created_by INT,
    updated_by INT
);

CREATE TABLE "WITHDRAWAL" (
    "id" SERIAL PRIMARY KEY,
    "user_id" INT,
    "request_code" VARCHAR(50) UNIQUE NOT NULL,

    "status" VARCHAR(30),
    "amount" FLOAT NOT NULL,
    "processed_by" INT,
    "note" TEXT,

    "created_by" INT,
    "updated_by" INT,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY ("user_id") REFERENCES "SYS_USER"("id") ON DELETE CASCADE
);


CREATE TABLE "ORDER" (
    id SERIAL PRIMARY KEY,
    user_id INT,
    withdrawal_id INT,
    package_id INT,                        

    order_code VARCHAR(30),
    package_code VARCHAR(50),
    price FLOAT,
    commission_affiliate FLOAT,
    commission_mobifone FLOAT,
    commission_agent FLOAT,
    customer_phone VARCHAR(20),
    order_status VARCHAR(50),       
    withdraw_status VARCHAR(50),
    withdraw_time TIMESTAMP,
    withdraw_code VARCHAR(20),
    order_type VARCHAR(50),
    created_by INT,
    updated_by INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY ("user_id") REFERENCES "SYS_USER"("id") ON DELETE CASCADE,
    FOREIGN KEY ("withdrawal_id") REFERENCES "WITHDRAWAL"("id") ON DELETE CASCADE
);