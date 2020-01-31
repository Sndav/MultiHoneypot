package mysql

import (
	"Multi-Honeypot/internal/pkg/plugin"
	sqle "github.com/src-d/go-mysql-server"
	"github.com/src-d/go-mysql-server/auth"
	"github.com/src-d/go-mysql-server/memory"
	"github.com/src-d/go-mysql-server/server"
	"github.com/src-d/go-mysql-server/sql"
	"time"
)

func Backend(ctx *plugin.Context) {
	driver := sqle.NewDefault()
	driver.AddDatabase(createTestDatabase())

	config := server.Config{
		Protocol: "unix",
		Address:  ctx.SocksFile,
		Auth:     auth.NewNativeSingle("root", "root", auth.AllPermissions),
	}

	s, err := server.NewDefaultServer(config, driver)
	if err != nil {
		panic(err)
	}

	go s.Start()

	ctx.Mutex.Signal(ctx.SocksFile)
}

func createTestDatabase() *memory.Database {
	const (
		dbName    = "test"
		tableName = "mytable"
	)

	db := memory.NewDatabase(dbName)
	table := memory.NewTable(tableName, sql.Schema{
		{Name: "name", Type: sql.Text, Nullable: false, Source: tableName},
		{Name: "email", Type: sql.Text, Nullable: false, Source: tableName},
		{Name: "phone_numbers", Type: sql.JSON, Nullable: false, Source: tableName},
		{Name: "created_at", Type: sql.Timestamp, Nullable: false, Source: tableName},
	})

	db.AddTable(tableName, table)
	ctx := sql.NewEmptyContext()

	rows := []sql.Row{
		sql.NewRow("John Doe", "john@doe.com", []string{"555-555-555"}, time.Now()),
		sql.NewRow("John Doe", "johnalt@doe.com", []string{}, time.Now()),
		sql.NewRow("Jane Doe", "jane@doe.com", []string{}, time.Now()),
		sql.NewRow("Evil Bob", "evilbob@gmail.com", []string{"555-666-555", "666-666-666"}, time.Now()),
	}

	for _, row := range rows {
		table.Insert(ctx, row)
	}

	return db
}
