ALTER TABLE files ADD hmac VARBINARY(32) NOT NULL;
ALTER TABLE files ADD _salt VARBINARY(32) NOT NULL;
ALTER TABLE file_chunks ADD _size INT NOT NULL;