CREATE TABLE IF NOT EXISTS transactions (
    kind String NOT NULL,
	date Date NOT NULL,
    recipient String NOT NULL,
	amount Decimal(18, 2) NOT NULL,
    primary_class String,
    secondary_class String,
    hash String NOT NULL,
	internal UInt8 DEFAULT 0
)
ENGINE = MergeTree
PRIMARY KEY (date, recipient, kind, amount);