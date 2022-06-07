CREATE TABLE Accounts (
  "id" bigserial PRIMARY KEY,
  "owner" varchar NOT NULL,
  "balance" bigint NOT NULL,
  "currency" varchar NOT NULL,
  "created_at" timestamp DEFAULT 'now()'
);

CREATE TABLE Entries (
  "id" bigserial PRIMARY KEY,
  "account_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamp DEFAULT 'now()'
);

CREATE TABLE transfers (
  "id" bigserial PRIMARY KEY,
  "from_account_id" bigint NOT NULL,
  "to_account_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT 'now()'
);

CREATE INDEX ON Accounts ("owner");

CREATE INDEX ON Entries ("account_id");

CREATE INDEX ON transfers ("from_account_id");

CREATE INDEX ON transfers ("to_account_id");

CREATE INDEX ON transfers ("from_account_id", "to_account_id");

COMMENT ON COLUMN Entries."amount" IS 'can be +ve or -ve';

ALTER TABLE Entries ADD FOREIGN KEY ("account_id") REFERENCES Accounts ("id");

ALTER TABLE transfers ADD FOREIGN KEY ("from_account_id") REFERENCES Accounts ("id");

ALTER TABLE transfers ADD FOREIGN KEY ("to_account_id") REFERENCES Accounts ("id");
