// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/AllPaste/paste/internal/data/ent/migrate"

	"github.com/AllPaste/paste/internal/data/ent/paste"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Paste is the client for interacting with the Paste builders.
	Paste *PasteClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Paste = NewPasteClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:    ctx,
		config: cfg,
		Paste:  NewPasteClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:    ctx,
		config: cfg,
		Paste:  NewPasteClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Paste.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Paste.Use(hooks...)
}

// PasteClient is a client for the Paste schema.
type PasteClient struct {
	config
}

// NewPasteClient returns a client for the Paste from the given config.
func NewPasteClient(c config) *PasteClient {
	return &PasteClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `paste.Hooks(f(g(h())))`.
func (c *PasteClient) Use(hooks ...Hook) {
	c.hooks.Paste = append(c.hooks.Paste, hooks...)
}

// Create returns a builder for creating a Paste entity.
func (c *PasteClient) Create() *PasteCreate {
	mutation := newPasteMutation(c.config, OpCreate)
	return &PasteCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Paste entities.
func (c *PasteClient) CreateBulk(builders ...*PasteCreate) *PasteCreateBulk {
	return &PasteCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Paste.
func (c *PasteClient) Update() *PasteUpdate {
	mutation := newPasteMutation(c.config, OpUpdate)
	return &PasteUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PasteClient) UpdateOne(pa *Paste) *PasteUpdateOne {
	mutation := newPasteMutation(c.config, OpUpdateOne, withPaste(pa))
	return &PasteUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PasteClient) UpdateOneID(id int64) *PasteUpdateOne {
	mutation := newPasteMutation(c.config, OpUpdateOne, withPasteID(id))
	return &PasteUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Paste.
func (c *PasteClient) Delete() *PasteDelete {
	mutation := newPasteMutation(c.config, OpDelete)
	return &PasteDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *PasteClient) DeleteOne(pa *Paste) *PasteDeleteOne {
	return c.DeleteOneID(pa.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *PasteClient) DeleteOneID(id int64) *PasteDeleteOne {
	builder := c.Delete().Where(paste.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PasteDeleteOne{builder}
}

// Query returns a query builder for Paste.
func (c *PasteClient) Query() *PasteQuery {
	return &PasteQuery{
		config: c.config,
	}
}

// Get returns a Paste entity by its id.
func (c *PasteClient) Get(ctx context.Context, id int64) (*Paste, error) {
	return c.Query().Where(paste.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PasteClient) GetX(ctx context.Context, id int64) *Paste {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *PasteClient) Hooks() []Hook {
	return c.hooks.Paste
}
