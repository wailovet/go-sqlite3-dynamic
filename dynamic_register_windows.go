// Copyright (C) 2019 Samuel Melrose <sam@infitialis.com>.
// Copyright (C) 2017 Hank Shen <swh@admpub.com>.
//
// Based on work by Yasuhiro Matsumoto <mattn.jp@gmail.com>
// https://github.com/mattn/go-sqlite3
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package sqlite3

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/nkbai/go-memorydll"
)

//go:embed memorydll/sqlite3.dll
var dllByte []byte

var (
	modSQLite3 *memorydll.DLL

	dll_sqlite3_libversion           *memorydll.Proc
	dll_sqlite3_libversion_number    *memorydll.Proc
	dll_sqlite3_sourceid             *memorydll.Proc
	dll_sqlite3_errstr               *memorydll.Proc
	dll_sqlite3_errcode              *memorydll.Proc
	dll_sqlite3_extended_errcode     *memorydll.Proc
	dll_sqlite3_errmsg               *memorydll.Proc
	dll_sqlite3_threadsafe           *memorydll.Proc
	dll_sqlite3_open_v2              *memorydll.Proc
	dll_sqlite3_busy_timeout         *memorydll.Proc
	dll_sqlite3_close_v2             *memorydll.Proc
	dll_sqlite3_prepare_v2           *memorydll.Proc
	dll_sqlite3_get_autocommit       *memorydll.Proc
	dll_sqlite3_finalize             *memorydll.Proc
	dll_sqlite3_bind_parameter_count *memorydll.Proc
	dll_sqlite3_bind_parameter_index *memorydll.Proc
	dll_sqlite3_reset                *memorydll.Proc
	dll_sqlite3_bind_null            *memorydll.Proc
	dll_sqlite3_bind_int64           *memorydll.Proc
	dll_sqlite3_bind_int             *memorydll.Proc
	dll_sqlite3_bind_text            *memorydll.Proc
	dll_sqlite3_bind_double          *memorydll.Proc
	dll_sqlite3_bind_blob            *memorydll.Proc
	dll_sqlite3_column_count         *memorydll.Proc
	dll_sqlite3_column_name          *memorydll.Proc
	dll_sqlite3_interrupt            *memorydll.Proc
	dll_sqlite3_clear_bindings       *memorydll.Proc
	dll_sqlite3_step                 *memorydll.Proc
	dll_sqlite3_column_decltype      *memorydll.Proc
	dll_sqlite3_column_type          *memorydll.Proc
	dll_sqlite3_column_int64         *memorydll.Proc
	dll_sqlite3_column_double        *memorydll.Proc
	dll_sqlite3_column_bytes         *memorydll.Proc
	dll_sqlite3_column_blob          *memorydll.Proc
	dll_sqlite3_column_text          *memorydll.Proc
	dll_sqlite3_db_handle            *memorydll.Proc
	dll_sqlite3_last_insert_rowid    *memorydll.Proc
	dll_sqlite3_changes              *memorydll.Proc
)

func registerLibrary() {
	if err := registerDLL(); err != nil {
		libraryRegistered = false
		libraryRegisterError = err
		return
	}

	dll_sqlite3_libversion, _ = modSQLite3.FindProc("sqlite3_libversion")
	dll_sqlite3_libversion_number, _ = modSQLite3.FindProc("sqlite3_libversion_number")
	dll_sqlite3_sourceid, _ = modSQLite3.FindProc("sqlite3_sourceid")

	dll_sqlite3_errstr, _ = modSQLite3.FindProc("sqlite3_errstr")
	dll_sqlite3_errcode, _ = modSQLite3.FindProc("sqlite3_errcode")
	dll_sqlite3_extended_errcode, _ = modSQLite3.FindProc("sqlite3_extended_errcode")
	dll_sqlite3_errmsg, _ = modSQLite3.FindProc("sqlite3_errmsg")
	dll_sqlite3_threadsafe, _ = modSQLite3.FindProc("sqlite3_threadsafe")

	dll_sqlite3_open_v2, _ = modSQLite3.FindProc("sqlite3_open_v2")
	dll_sqlite3_busy_timeout, _ = modSQLite3.FindProc("sqlite3_busy_timeout")
	dll_sqlite3_close_v2, _ = modSQLite3.FindProc("sqlite3_close_v2")

	dll_sqlite3_prepare_v2, _ = modSQLite3.FindProc("sqlite3_prepare_v2")
	dll_sqlite3_get_autocommit, _ = modSQLite3.FindProc("sqlite3_get_autocommit")

	dll_sqlite3_finalize, _ = modSQLite3.FindProc("sqlite3_finalize")
	dll_sqlite3_bind_parameter_count, _ = modSQLite3.FindProc("sqlite3_bind_parameter_count")
	dll_sqlite3_bind_parameter_index, _ = modSQLite3.FindProc("sqlite3_bind_parameter_index")
	dll_sqlite3_reset, _ = modSQLite3.FindProc("sqlite3_reset")
	dll_sqlite3_bind_null, _ = modSQLite3.FindProc("sqlite3_bind_null")
	dll_sqlite3_bind_int64, _ = modSQLite3.FindProc("sqlite3_bind_int64")
	dll_sqlite3_bind_int, _ = modSQLite3.FindProc("sqlite3_bind_int")
	dll_sqlite3_bind_text, _ = modSQLite3.FindProc("sqlite3_bind_text")
	dll_sqlite3_bind_double, _ = modSQLite3.FindProc("sqlite3_bind_double")
	dll_sqlite3_bind_blob, _ = modSQLite3.FindProc("sqlite3_bind_blob")
	dll_sqlite3_column_count, _ = modSQLite3.FindProc("sqlite3_column_count")
	dll_sqlite3_column_name, _ = modSQLite3.FindProc("sqlite3_column_name")
	dll_sqlite3_interrupt, _ = modSQLite3.FindProc("sqlite3_interrupt")
	dll_sqlite3_clear_bindings, _ = modSQLite3.FindProc("sqlite3_clear_bindings")

	dll_sqlite3_step, _ = modSQLite3.FindProc("sqlite3_step")
	dll_sqlite3_column_decltype, _ = modSQLite3.FindProc("sqlite3_column_decltype")
	dll_sqlite3_column_type, _ = modSQLite3.FindProc("sqlite3_column_type")
	dll_sqlite3_column_int64, _ = modSQLite3.FindProc("sqlite3_column_int64")
	dll_sqlite3_column_double, _ = modSQLite3.FindProc("sqlite3_column_double")
	dll_sqlite3_column_bytes, _ = modSQLite3.FindProc("sqlite3_column_bytes")
	dll_sqlite3_column_blob, _ = modSQLite3.FindProc("sqlite3_column_blob")
	dll_sqlite3_column_text, _ = modSQLite3.FindProc("sqlite3_column_text")

	dll_sqlite3_db_handle, _ = modSQLite3.FindProc("sqlite3_db_handle")
	dll_sqlite3_last_insert_rowid, _ = modSQLite3.FindProc("sqlite3_last_insert_rowid")
	dll_sqlite3_changes, _ = modSQLite3.FindProc("sqlite3_changes")

	libraryRegistered = true
}

func registerDLL() error {
	var err error
	modSQLite3, err = memorydll.NewDLL(dllByte, "sqlite3.dll")
	return err
}

func getLibraryPath() (string, error) {
	bPath, dllName := basePath(), "sqlite3.dll"

	if exist, _ := exists(dllName); exist {
		return dllName, nil
	}

	filePath := bPath + dllName
	if exist, _ := exists(filePath); exist {
		return filePath, nil
	}

	filePath = bPath + "support" + string(os.PathSeparator) + dllName
	if exist, _ := exists(filePath); exist {
		return filePath, nil
	}

	return "", fmt.Errorf("%s not found.", dllName)
}
