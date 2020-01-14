package metrics

import (
	"database/sql"
	pb "github.com/33cn/chain33/types"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func NewStore(config string) *Store {
	store := &Store{}
	store.connected = false
	store.connect(config)

	return store
}


type Store struct {
	db *sql.DB
	connected bool
}

func (s *Store) connect(config string) bool  {
	database,err := sql.Open("mysql",config)
	if err != nil{
		return false
	}
	database.SetConnMaxLifetime(100*time.Second)
	database.SetMaxOpenConns(100)
	database.SetMaxIdleConns(16)

	s.db = database
	s.connected = true
	return true
}

func (s *Store) insertMetrics(dstID string,dst string,info *pb.MetricsInfo)  {
	if !s.connected {
		return
	}

	execStr := "insert into Metrics(" +
		"DstID,Dst,MKey,Action,SrcID,Src,Size,Time,Other) " +
		"values(?,?,?,?,?,?,?,?,?)"

	_,err := s.db.Exec(execStr,
		dstID,dst,info.Key,info.Action,
		info.SrcID,info.Src,info.Size,info.Time,info.Other)
	if err != nil{
		//
	}
}