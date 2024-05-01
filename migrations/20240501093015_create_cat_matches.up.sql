CREATE TABLE IF NOT EXISTS cat_matches(
  id SERIAL PRIMARY KEY,
  message VARCHAR(255),
  issuer_cat_id INT references cats(id) NOT NULL,
  receiver_cat_id INT references cats(id) NOT NULL,
  status VARCHAR(100) NOT NULL,
  created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMPTZ
);
