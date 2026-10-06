CREATE TABLE expenses (
                          id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

                          user_id UUID NOT NULL,
                          category_id UUID NOT NULL,

                          title VARCHAR(255) NOT NULL,
                          description TEXT,

                          amount NUMERIC(12,2) NOT NULL,

                          expense_date DATE NOT NULL,

                          created_at TIMESTAMP NOT NULL DEFAULT NOW(),
                          updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

                          CONSTRAINT fk_user
                              FOREIGN KEY(user_id)
                                  REFERENCES users(id)
                                  ON DELETE CASCADE,

                          CONSTRAINT fk_category
                              FOREIGN KEY(category_id)
                                  REFERENCES categories(id)
);
