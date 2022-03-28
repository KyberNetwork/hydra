CREATE TABLE IF NOT EXISTS kyber_client_secret(
    pk                                   SERIAL PRIMARY KEY,
	client_id      	                     varchar(255) NOT NULL,
	client_secret_plaintext 	         text NOT NULL,
    created_at                           TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at                           TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX kyber_client_secret_client_id ON kyber_client_secret USING hash (client_id);
