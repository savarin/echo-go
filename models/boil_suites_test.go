// Code generated by SQLBoiler 4.15.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("EchoLogs", testEchoLogs)
	t.Run("SchemaMigrations", testSchemaMigrations)
}

func TestDelete(t *testing.T) {
	t.Run("EchoLogs", testEchoLogsDelete)
	t.Run("SchemaMigrations", testSchemaMigrationsDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("EchoLogs", testEchoLogsQueryDeleteAll)
	t.Run("SchemaMigrations", testSchemaMigrationsQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("EchoLogs", testEchoLogsSliceDeleteAll)
	t.Run("SchemaMigrations", testSchemaMigrationsSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("EchoLogs", testEchoLogsExists)
	t.Run("SchemaMigrations", testSchemaMigrationsExists)
}

func TestFind(t *testing.T) {
	t.Run("EchoLogs", testEchoLogsFind)
	t.Run("SchemaMigrations", testSchemaMigrationsFind)
}

func TestBind(t *testing.T) {
	t.Run("EchoLogs", testEchoLogsBind)
	t.Run("SchemaMigrations", testSchemaMigrationsBind)
}

func TestOne(t *testing.T) {
	t.Run("EchoLogs", testEchoLogsOne)
	t.Run("SchemaMigrations", testSchemaMigrationsOne)
}

func TestAll(t *testing.T) {
	t.Run("EchoLogs", testEchoLogsAll)
	t.Run("SchemaMigrations", testSchemaMigrationsAll)
}

func TestCount(t *testing.T) {
	t.Run("EchoLogs", testEchoLogsCount)
	t.Run("SchemaMigrations", testSchemaMigrationsCount)
}

func TestHooks(t *testing.T) {
	t.Run("EchoLogs", testEchoLogsHooks)
	t.Run("SchemaMigrations", testSchemaMigrationsHooks)
}

func TestInsert(t *testing.T) {
	t.Run("EchoLogs", testEchoLogsInsert)
	t.Run("EchoLogs", testEchoLogsInsertWhitelist)
	t.Run("SchemaMigrations", testSchemaMigrationsInsert)
	t.Run("SchemaMigrations", testSchemaMigrationsInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("EchoLogs", testEchoLogsReload)
	t.Run("SchemaMigrations", testSchemaMigrationsReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("EchoLogs", testEchoLogsReloadAll)
	t.Run("SchemaMigrations", testSchemaMigrationsReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("EchoLogs", testEchoLogsSelect)
	t.Run("SchemaMigrations", testSchemaMigrationsSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("EchoLogs", testEchoLogsUpdate)
	t.Run("SchemaMigrations", testSchemaMigrationsUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("EchoLogs", testEchoLogsSliceUpdateAll)
	t.Run("SchemaMigrations", testSchemaMigrationsSliceUpdateAll)
}
