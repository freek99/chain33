package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"

	pb "github.com/33cn/chain33/types"
)


type MetricsDB struct {
	db *sql.DB
}

func (mdb *MetricsDB) connect(dbCfg string) bool  {
	database,err := sql.Open("mysql",dbCfg)
	if err != nil{
		fmt.Printf("Open mysql failed,err:%v\n",err)
		return false
	}
	database.SetConnMaxLifetime(100*time.Second)
	database.SetMaxOpenConns(100)
	database.SetMaxIdleConns(16)

	mdb.db = database
	return true
}

func (mdb *MetricsDB) searrchBroadcastAction(key string) []*pb.MetricsInfo {
	var infos []*pb.MetricsInfo

	queryStr := "select * from Metrics where action=\"recv\" and mkey=?"
	rows, err := mdb.db.Query(queryStr, key)
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	if err != nil {
		fmt.Printf("Query failed,err:%v", err)
		return infos
	}

	for rows.Next() {
		info := &pb.MetricsInfo{}
		id := 0
		err = rows.Scan(&id,
			&info.DstID,&info.Dst,&info.Key, &info.Action,
			&info.SrcID,&info.Src,&info.Size,&info.Time,&info.Other)
		if err != nil {
			fmt.Printf("Scan failed,err:%v", err)
			return infos
		}

		infos = append(infos,info)
	}

	return  infos
}


func (mdb *MetricsDB) searchRewardAction(startAddr string) map[string][]*pb.MetricsInfo {
	replys := make(map[string][]*pb.MetricsInfo)

	queryStr := "select * from Metrics where (action=\"rollback\" or action=\"attach\") and dst=?"
	rows, err := mdb.db.Query(queryStr, startAddr)
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	if err != nil {
		fmt.Printf("Query failed,err:%v", err)
		return replys
	}

	var infos []*pb.MetricsInfo
	for rows.Next() {
		info := &pb.MetricsInfo{}
		id := 0
		err = rows.Scan(&id,
			&info.DstID,&info.Dst,&info.Key, &info.Action,
			&info.SrcID,&info.Src,&info.Size,&info.Time,&info.Other)
		if err != nil {
			fmt.Printf("Scan failed,err:%v", err)
			return replys
		}

		infos = append(infos,info)
	}

	replys[startAddr] = infos

	return  replys
}



