-- Store table
CREATE TABLE store (
    id SERIAL,
    cnpj VARCHAR(20) UNIQUE NOT NULL,
    PRIMARY KEY(id));

-- Person table
CREATE TABLE person (
    id SERIAL,
    most_frequent_store_id INT,
    last_purchase_store_id INT,
    cpf VARCHAR(20) UNIQUE NOT NULL,
    private BOOLEAN DEFAULT false,
    incomplete BOOLEAN DEFAULT false,
    avg_ticket_value INT,
    last_purchase_date DATE,
    last_purchase_ticket INT,
    PRIMARY KEY(id),
    CONSTRAINT fk_most_frequent_store
        FOREIGN KEY(most_frequent_store_id)
            REFERENCES store(id),
    CONSTRAINT fk_last_purchase_store
        FOREIGN KEY(last_purchase_store_id)
            REFERENCES store(id));