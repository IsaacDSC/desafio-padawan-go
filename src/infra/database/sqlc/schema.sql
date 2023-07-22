CREATE TABLE exchange_rate (
  id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  input varchar(255) NOT NULL,
  output varchar(255) NOT NULL,
  amount_in REAL,
  amount_out REAL,
  rate REAL,
  created_at DATETIME DEFAULT NOW()
);