package server

import (
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type USDBRLService struct{}

func NewUSDBRLService() *USDBRLService {
	initDB()
	return &USDBRLService{}
}

func (s *USDBRLService) USDBRLHandler(w http.ResponseWriter, r *http.Request) {
	usdbrlResponse, err := s.getUSDBRLRate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = usdbrlResponse.Validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = s.save(usdbrlResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	res := ServiceResponse{
		Bid: usdbrlResponse.Bid,
	}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *USDBRLService) getUSDBRLRate() (*USDBRL, error) {
	ctx, cancel := context.WithTimeout(context.Background(), REQUEST_TIMEOUT*time.Millisecond)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, BASE_URL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var r APIResponse
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}
	return &r.USDBRL, nil
}

func (s *USDBRLService) save(r *USDBRL) error {
	ctx, cancel := context.WithTimeout(context.Background(), DB_TIMEOUT*time.Millisecond)
	defer cancel()
	db, err := sql.Open("sqlite3", DB_FILE)
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.ExecContext(
		ctx,
		SQL_INSERT,
		r.Codein,
		r.Name,
		r.High,
		r.Low,
		r.VarBid,
		r.PctChange,
		r.Bid,
		r.Ask,
		r.Timestamp,
		r.CreateDate,
	)
	if err != nil {
		return err
	}
	return nil
}
