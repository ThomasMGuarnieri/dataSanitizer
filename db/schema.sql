CREATE TABLE data (
    id SERIAL,
    cnpj VARCHAR NOT NULL,
    most_frequent_store_id INT,
    last_purchase_store_id INT,
    cpf VARCHAR(20) UNIQUE NOT NULL,
    private BOOLEAN DEFAULT false,
    incomplete BOOLEAN DEFAULT false,
    avg_ticket_value INT,
    last_purchase_date DATE,
    last_purchase_ticket INT