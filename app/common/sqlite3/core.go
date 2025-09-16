package sqlite3

import (
	"database/sql"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	Db *gorm.DB
)

func Init(path string) error {
	// 连接到 SQLite 数据库
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return err
	}
	Db = db
	sqlDb, err := db.DB()
	if err != nil {
		return err
	}
	createCfr(sqlDb)
	createCss(sqlDb)
	return nil
}

func createCfr(db *sql.DB) bool {
	sql := `
CREATE TABLE IF NOT EXISTS "config_forward_rule" (
  "id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
  "name" TEXT,
  "remote_addr" TEXT,
  "local_addr" TEXT,
  "tag" TEXT,
  "sort" integer DEFAULT 0,
  "css_id" integer NOT NULL
);

CREATE INDEX IF NOT EXISTS "config_forward_rule_idx_id" ON "config_forward_rule" ( "id" DESC );
CREATE INDEX IF NOT EXISTS "config_forward_rule_idx_sort" ON "config_forward_rule" ( "sort" DESC );
`
	_, err := db.Exec(sql)
	return err != nil
}

func createCss(db *sql.DB) bool {
	sql := `
CREATE TABLE IF NOT EXISTS "config_ssh_server" (
  "id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
  "name" TEXT,
  "host" TEXT,
  "user" TEXT,
  "pass" TEXT,
  "key_path" TEXT,
  "pass_phrase" TEXT,
  "sort" integer DEFAULT 0
);

CREATE INDEX IF NOT EXISTS "config_ssh_server_idx_id" ON "config_ssh_server" ( "id" DESC );
CREATE INDEX IF NOT EXISTS "config_ssh_server_idx_sort" ON "config_ssh_server" ( "sort" DESC );
`
	_, err := db.Exec(sql)
	return err != nil
}
