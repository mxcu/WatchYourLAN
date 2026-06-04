# example create tabe DDL

CREATE TABLE default.watch_your_lan(
    ts     DateTime64(3, 'UTC'),
    ip     IPv4,
    iface  LowCardinality(String),
    name   LowCardinality(String),
    mac    FixedString(17),
    known  UInt8
)
ENGINE = MergeTree
PARTITION BY toYYYYMM(ts)
ORDER BY (ts, ip)
SETTINGS index_granularity = 8192;