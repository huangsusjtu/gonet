package persistence

import (
	"testing"
)

/*func TestConnPool1(t *testing.T) {
	pool := GetSqlConPool()
	pool.Init()
	pool.Destroy()
}
*/
func TestConnPool2(t *testing.T) {
	pool := GetSqlConPool()
	pool.Init()

	db := pool.GetDb()
	pool.dump()

	pool.RecycleDB(db)
	pool.dump()
	pool.Destroy()
}
