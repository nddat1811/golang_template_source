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
    created_by INT NOT NULL,
    updated_by INT
);

-- Bảng SYS_LOG
CREATE TABLE "SYS_LOG" (
    id SERIAL PRIMARY KEY,
    action_datetime TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    path_name VARCHAR(255) NOT NULL,
    method VARCHAR(10) NOT NULL,
    ip VARCHAR(45) NOT NULL,
    status_response INT NOT NULL,
    response TEXT,
    description TEXT,
    request_body TEXT,
    request_query TEXT,
    duration FLOAT
);

-- Bảng SYS_ROLE
CREATE TABLE "SYS_ROLE" (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    description TEXT,
    status INT,
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
    created_by INT NOT NULL,
    FOREIGN KEY (role_id) REFERENCES "SYS_ROLE"(id) ON DELETE CASCADE,
    FOREIGN KEY (function_id) REFERENCES "SYS_FUNCTION"(id) ON DELETE CASCADE
);

-- Bảng SYS_USER
CREATE TABLE "SYS_USER" (
    id SERIAL PRIMARY KEY,
    full_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    phone VARCHAR(15),
    hash_password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

-- Bảng SYS_USER_ROLE
CREATE TABLE "SYS_USER_ROLE" (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    role_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    created_by INT NOT NULL,
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
('Admin', 'Quản trị viên hệ thống', 1, 1),
('User', 'Người dùng thông thường', 1, 1);

-- Thêm người dùng vào bảng SYS_USER
INSERT INTO "SYS_USER" (full_name, email, phone, hash_password)
VALUES 
('Admin', 'admin@gmail.com', '0123456789', '$2a$10$j3u/NoMXCHGaAQX92unfHeymF4N0foLDeBqLP7N8wjY9S/gixulG6'),
('User', 'user@gmail.com', '0987654321', '$2a$10$j3u/NoMXCHGaAQX92unfHeymF4N0foLDeBqLP7N8wjY9S/gixulG6');

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

