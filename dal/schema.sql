CREATE DATABASE IF NOT EXISTS executor;

SET
  DATABASE TO executor;

CREATE TABLE IF NOT EXISTS execution_context (
  -- computed message id, hashed differently from transfer or routing info
  id BYTES PRIMARY KEY,
  -- proto serialized bytes of the proto sgn.message.v1.ExecutionContext, 
  exec_ctx BYTES,
  -- see dal.ExecutionStatus
  status INT DEFAULT 1,
  retry_count INT DEFAULT 0,
  chain_id INT,
  src_chain_id INT,
  tx TEXT UNIQUE DEFAULT NULL,
  src_tx TEXT DEFAULT NULL,
  create_time TIMESTAMPTZ,
  update_time TIMESTAMPTZ
);

ALTER TABLE
  execution_context
ADD
  COLUMN IF NOT EXISTS chain_id INT;

ALTER TABLE
  execution_context
ADD
  COLUMN IF NOT EXISTS tx TEXT UNIQUE DEFAULT NULL;

ALTER TABLE
  execution_context
ADD
  COLUMN IF NOT EXISTS src_tx TEXT UNIQUE DEFAULT NULL;

ALTER TABLE
  execution_context
ADD
  COLUMN IF NOT EXISTS src_chain_id INT;

ALTER TABLE
  execution_context
ADD
  COLUMN IF NOT EXISTS create_time TIMESTAMPTZ;

ALTER TABLE
  execution_context
ADD
  COLUMN IF NOT EXISTS update_time TIMESTAMPTZ;

UPDATE
  execution_context
SET
  create_time = now(),
  update_time = now()
WHERE
  create_time IS NULL;

CREATE INDEX IF NOT EXISTS "execution_context_src_tx_key" ON execution_context (src_tx DESC);

CREATE TABLE IF NOT EXISTS monitor_block (
  chid_addr TEXT,
  block_num INT,
  block_idx INT,
  PRIMARY KEY (chid_addr)
);