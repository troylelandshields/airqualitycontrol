package webhooks

import (
	"context"
	"database/sql"
	"fmt"
)

const (
	allWebhooksQuery = "SELECT webhook FROM webhooks"

	insertQuery = "INSERT INTO webhooks (webhook) VALUES ($1)"
)

type Client struct {
	db *sql.DB
}

func New(db *sql.DB) *Client {
	return &Client{
		db: db,
	}
}

func (c *Client) Create(ctx context.Context, webhook string) error {
	_, err := c.db.ExecContext(ctx, insertQuery, webhook)
	if err != nil {
		return err
	}

	return err
}

func (c *Client) Webhooks(ctx context.Context) ([]string, error) {
	rows, err := c.db.QueryContext(ctx, allWebhooksQuery)
	if err != nil {
		return nil, err
	}

	var webhooks []string
	for rows.Next() {
		var webhook string

		err := rows.Scan(&webhook)
		if err != nil {
			fmt.Println("Error scanning row sorry get over it", err)
			continue
		}

		webhooks = append(webhooks, webhook)
	}

	return webhooks, nil
}
