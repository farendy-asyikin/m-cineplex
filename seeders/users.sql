INSERT INTO users (name, email, password, role, is_active, created_at, updated_at)
VALUES
    ('Anonycat', 'anocat@gmail.com', '$2a$10$Ibi3s2t1XrKfFZ81Nz9Aw.eGHGj.lwS59.Xap62i.ZPvqfswKMNae', 'superuser', true, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('Garfield', 'garfield@gmail.com', '$2a$10$Ibi3s2t1XrKfFZ81Nz9Aw.eGHGj.lwS59.Xap62i.ZPvqfswKMNae', 'admin', true, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('Casper', 'casper@gmail.com', '$2a$10$Ibi3s2t1XrKfFZ81Nz9Aw.eGHGj.lwS59.Xap62i.ZPvqfswKMNae', 'member', true, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- plain text for login password = password