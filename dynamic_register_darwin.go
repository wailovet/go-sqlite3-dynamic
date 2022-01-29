// Copyright (C) 2019 Samuel Melrose <sam@infitialis.com>.
// Copyright (C) 2017 Hank Shen <swh@admpub.com>.
//
// Based on work by Yasuhiro Matsumoto <mattn.jp@gmail.com>
// https://github.com/mattn/go-sqlite3
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package sqlite3

var (
	raw_sqlite3_libversion           func() []byte                                                                                      // const char *sqlite3_libversion(void);
	raw_sqlite3_libversion_number    func() int32                                                                                       // int sqlite3_libversion_number(void);
	raw_sqlite3_sourceid             func() []byte                                                                                      // const char *sqlite3_sourceid(void);
	raw_sqlite3_errstr               func(i int32) []byte                                                                               // const char *sqlite3_errstr(int);
	raw_sqlite3_errcode              func(handle uintptr) int32                                                                         // int sqlite3_errcode(sqlite3 *db);
	raw_sqlite3_extended_errcode     func(handle uintptr) int32                                                                         // int sqlite3_extended_errcode(sqlite3 *db);
	raw_sqlite3_errmsg               func(handle uintptr) []byte                                                                        // const char *sqlite3_errmsg(sqlite3*);
	raw_sqlite3_threadsafe           func() int32                                                                                       // int sqlite3_threadsafe(void);
	raw_sqlite3_open_v2              func(filename uintptr, outHandle uintptr, flags int32, zVfs uintptr) int32                         // int sqlite3_open_v2(const char *filename, sqlite3 **ppDb, int flags, const char *zVfs);
	raw_sqlite3_busy_timeout         func(handle uintptr, ms int32) int32                                                               // int sqlite3_busy_timeout(sqlite3*, int ms);
	raw_sqlite3_close_v2             func(handle uintptr) int32                                                                         // int sqlite3_close_v2(sqlite3*);
	raw_sqlite3_prepare_v2           func(handle uintptr, zSql []byte, nByte int32, outStmtHandle uintptr, outTailHandle uintptr) int32 // int sqlite3_prepare_v2(sqlite3 *db, const char *zSql, int nByte, sqlite3_stmt **ppStmt, const char **pzTail);
	raw_sqlite3_get_autocommit       func(handle uintptr) int32                                                                         // int sqlite3_get_autocommit(sqlite3*);
	raw_sqlite3_finalize             func(stmtHandle uintptr) int32                                                                     // int sqlite3_finalize(sqlite3_stmt *pStmt);
	raw_sqlite3_bind_parameter_count func(stmtHandle uintptr) int32                                                                     // int sqlite3_bind_parameter_count(sqlite3_stmt*);
	raw_sqlite3_bind_parameter_index func(stmtHandle uintptr, zName []byte) int32                                                       // int sqlite3_bind_parameter_index(sqlite3_stmt*, const char *zName);
	raw_sqlite3_reset                func(stmtHandle uintptr) int32                                                                     // int sqlite3_reset(sqlite3_stmt *pStmt);
	raw_sqlite3_bind_null            func(stmtHandle uintptr, index int32) int32                                                        // int sqlite3_bind_null(sqlite3_stmt*, int);
	raw_sqlite3_bind_int64           func(stmtHandle uintptr, index int32, data int64) int32                                            // int sqlite3_bind_int64(sqlite3_stmt*, int, sqlite3_int64);
	raw_sqlite3_bind_int             func(stmtHandle uintptr, index int32, data int32) int32                                            // int sqlite3_bind_int(sqlite3_stmt*, int, int);
	raw_sqlite3_bind_text            func(stmtHandle uintptr, index int32, data []byte, length int32, overflow uintptr) int32           // int sqlite3_bind_text(sqlite3_stmt*,int,const char*,int,void(*)(void*));
	raw_sqlite3_bind_double          func(stmtHandle uintptr, index int32, data float64) int32                                          // int sqlite3_bind_double(sqlite3_stmt*, int, double);
	raw_sqlite3_bind_blob            func(stmtHandle uintptr, index int32, data uintptr, length int32, overflow uintptr) int32          // int sqlite3_bind_blob(sqlite3_stmt*, int, const void*, int n, void(*)(void*));
	raw_sqlite3_column_count         func(stmtHandle uintptr) int32                                                                     // int sqlite3_column_count(sqlite3_stmt *pStmt);
	raw_sqlite3_column_name          func(stmtHandle uintptr, index int32) []byte                                                       // const char *sqlite3_column_name(sqlite3_stmt*, int N);
	raw_sqlite3_interrupt            func(handle uintptr)                                                                               // void sqlite3_interrupt(sqlite3*);
	raw_sqlite3_clear_bindings       func(stmtHandle uintptr) int32                                                                     // int sqlite3_clear_bindings(sqlite3_stmt*);
	raw_sqlite3_step                 func(stmtHandle uintptr) int32                                                                     // int sqlite3_step(sqlite3_stmt*);
	raw_sqlite3_column_decltype      func(stmtHandle uintptr, index int32) []byte                                                       // const char *sqlite3_column_decltype(sqlite3_stmt*,int);
	raw_sqlite3_column_type          func(stmtHandle uintptr, index int32) int32                                                        // int sqlite3_column_type(sqlite3_stmt*, int iCol);
	raw_sqlite3_column_int64         func(stmtHandle uintptr, index int32) int64                                                        // sqlite3_int64 sqlite3_column_int64(sqlite3_stmt*, int iCol);
	raw_sqlite3_column_double        func(stmtHandle uintptr, index int32) float64                                                      // double sqlite3_column_double(sqlite3_stmt*, int iCol);
	raw_sqlite3_column_bytes         func(stmtHandle uintptr, index int32) int32                                                        // int sqlite3_column_bytes(sqlite3_stmt*, int iCol);
	raw_sqlite3_column_blob          func(stmtHandle uintptr, index int32) uintptr                                                      // const void *sqlite3_column_blob(sqlite3_stmt*, int iCol);
	raw_sqlite3_column_text          func(stmtHandle uintptr, index int32) uintptr                                                      // const unsigned char *sqlite3_column_text(sqlite3_stmt*, int iCol);
	raw_sqlite3_db_handle            func(stmtHandle uintptr) uintptr                                                                   // sqlite3 *sqlite3_db_handle(sqlite3_stmt*);
	raw_sqlite3_last_insert_rowid    func(handle uintptr) int64                                                                         // sqlite3_int64 sqlite3_last_insert_rowid(sqlite3*);
	raw_sqlite3_changes              func(handle uintptr) int32                                                                         // int sqlite3_changes(sqlite3*);
)

func registerLibrary() {

}
