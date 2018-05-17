/*
 * Copyright (c) 2016 TFG Co <backend@tfgco.com>
 * Author: TFG Co <backend@tfgco.com>
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of
 * this software and associated documentation files (the "Software"), to deal in
 * the Software without restriction, including without limitation the rights to
 * use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
 * the Software, and to permit persons to whom the Software is furnished to do so,
 * subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
 * FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
 * COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
 * IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
 * CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package interfaces

import (
	"context"
	"io"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

// Queryable is a contract implemented by types that query
// pg and expect go-pg results
type Queryable interface {
	Exec(interface{}, ...interface{}) (orm.Result, error)
	ExecOne(interface{}, ...interface{}) (orm.Result, error)
	Query(interface{}, interface{}, ...interface{}) (orm.Result, error)
	QueryOne(interface{}, interface{}, ...interface{}) (orm.Result, error)
	Model(model ...interface{}) *orm.Query // mesma coisa aqui, temos que ter uma iface de query e não retornar o do pg
}

type ORM interface {
	Select(model interface{}) error
	Insert(model ...interface{}) error
	Update(model ...interface{}) error
	Delete(model interface{}) error

	CopyFrom(r io.Reader, query interface{}, params ...interface{}) (orm.Result, error)
	CopyTo(w io.Writer, query interface{}, params ...interface{}) (orm.Result, error)
	FormatQuery(b []byte, query string, params ...interface{}) []byte
}

// DB represents the contract for a Postgres DB
type DB interface {
	ORM
	Queryable
	Close() error
	Begin() (*pg.Tx, error)
	WithContext(ctx context.Context) *pg.DB // probably should return a DB (interface) as well, not pg.DB
	Context() context.Context
}

// Tx represents the contract for a Postgres Tx
type Tx interface {
	Queryable
	Rollback() error
	Commit() error
}

// TxWrapper is the interface for mocking pg transactions
type TxWrapper interface {
	DbBegin(db DB) (Tx, error)
}
