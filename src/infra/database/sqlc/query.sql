-- name: CreateInfoExchangeRate :exec
INSERT INTO exchange_rate (input, output, amount_in, amount_out, rate) VALUES (?, ?, ?, ?, ?);
