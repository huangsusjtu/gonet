package persistence

import (
	. "config"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"sync"
	. "util"
)

//
var logger = NewLogger("database")

type SqlConPool struct {
	pool      []*sql.DB
	available int
	cond      *sync.Cond
	lock      sync.Mutex
	inited    bool
}

var sqlConPool = &SqlConPool{pool: make([]*sql.DB, 8), available: 0, inited: false}

func GetSqlConPool() *SqlConPool {
	return sqlConPool
}

func (scp *SqlConPool) Init() {
	scp.lock.Lock()
	defer scp.lock.Unlock()

	if scp.inited == true {
		logger.Errorln("sqlConPool have been inited")
		return
	}

	scp.cond = sync.NewCond(&scp.lock)
	var err error
	for i, db := range scp.pool {
		if db != nil {
			logger.Warnf("pool slot %i is not nil", i)
			continue
		}

		db, err = sql.Open(DB_TYPE, DB_CONFIG)
		if err != nil {
			logger.Errorln(err)
			continue
		}

		err = db.Ping()
		if err != nil {
			logger.Errorln(err)
			continue
		}

		scp.pool[i] = db
		scp.available++
	}

	scp.inited = true
	logger.Debugln("SqlConPool init ok")
}

func (scp *SqlConPool) Destroy() {
	scp.lock.Lock()
	defer scp.lock.Unlock()

	if scp.inited == false {
		logger.Errorln("sqlConPool have not been inited")
		return
	}

	for i, db := range scp.pool {
		if db != nil {
			db.Close()
		}
		scp.pool[i] = nil
	}
	scp.available = 0
	scp.inited = false
	logger.Debugln("SqlConPool destroy ok")
}

func (scp *SqlConPool) GetDb() *sql.DB {
	var db *sql.DB = nil
	scp.lock.Lock()
	for scp.available == 0 {
		scp.cond.Wait()
	}
	for i, tmpdb := range scp.pool {
		if tmpdb != nil {
			db = tmpdb
			scp.pool[i] = nil
			break
		}
	}
	scp.available--
	scp.lock.Unlock()
	return db
}

func (scp *SqlConPool) RecycleDB(db *sql.DB) {
	if db == nil {
		return
	}

	scp.lock.Lock()
	defer scp.lock.Unlock()
	for i, tmpdb := range scp.pool {
		if tmpdb == nil {
			scp.pool[i] = db
			scp.available++
			if scp.available == 1 {
				scp.cond.Signal()
			}
			return
		}
	}

	db.Close()
}

func (scp *SqlConPool) dump() {
	scp.lock.Lock()
	defer scp.lock.Unlock()

	logger.Debugf("inited %v, available %v\n", scp.inited, scp.available)
	for i, tmpdb := range scp.pool {
		logger.Debugf("i: %v, db: %v", i, tmpdb)
	}
}
