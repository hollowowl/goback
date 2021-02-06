package internal

import (
	"arduino-playground.xyz/goback/config"
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(config config.DBConfig) (*Storage, error) {
	db, err := sql.Open(config.Type, config.DSN)
	if err != nil {
		return nil, err
	}
	storage := &Storage{}
	storage.db = db
	return storage, nil
}
func (storage *Storage) getBoardId(boardName string) (int64, error) {
	var id int64
	row := storage.db.QueryRow("SELECT id FROM board WHERE name=? LIMIT 1", boardName)
	switch err := row.Scan(&id); err {
	case sql.ErrNoRows:
		res, err := storage.db.Exec("INSERT board SET name=?", boardName)
		if err != nil {
			return -1, err
		}
		return res.LastInsertId()
	case nil:
		return id, nil
	default:
		return -1, err
	}
}

func (storage *Storage) RegisterIncomingData(data *Data) (*DataID, error) {
	bid, err := storage.getBoardId(data.Board)
	if err != nil {
		return nil, err
	}
	encoded, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	res, err := storage.db.Exec("INSERT received_data SET board_id=?, data=?", bid, encoded)
	if err != nil {
		return nil, err
	}
	rdid, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &DataID{bid, rdid}, nil
}

func (storage *Storage) RegisterDecision(dataID *DataID, d *Decision) error {
	encoded, err := json.Marshal(d)
	if err != nil {
		return err
	}
	if _, err := storage.db.Exec("INSERT decision SET board_id=?, received_data_id=?, decision=?", dataID.BoardID, dataID.ReceivedDataID, encoded); err != nil {
		return err
	}
	return nil
}
