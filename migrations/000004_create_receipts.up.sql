CREATE TABLE receipts (
                          id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

                          expense_id UUID NOT NULL,

                          s3_url TEXT NOT NULL,

                          uploaded_at TIMESTAMP NOT NULL DEFAULT NOW(),

                          CONSTRAINT fk_expense
                              FOREIGN KEY(expense_id)
                                  REFERENCES expenses(id)
                                  ON DELETE CASCADE
);
