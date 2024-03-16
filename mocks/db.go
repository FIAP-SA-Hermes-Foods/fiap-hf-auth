package mocks

import (
	"context"
	"database/sql"
	"strings"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Mock DB

type MockDbSQL struct {
	WantResult   sql.Result
	WantRows     *sql.Rows
	WantErr      error
	WantNextRows bool
}

func (m MockDbSQL) Connect() error {
	if m.WantErr != nil && strings.EqualFold("errConnect", m.WantErr.Error()) {
		return m.WantErr
	}
	return nil
}

func (m MockDbSQL) Close() error {
	if m.WantErr != nil && strings.EqualFold("errClose", m.WantErr.Error()) {
		return m.WantErr
	}
	return nil
}

func (m MockDbSQL) PrepareStmt(query string) error {
	if m.WantErr != nil && strings.EqualFold("errPrepareStmt", m.WantErr.Error()) {
		return m.WantErr
	}
	return nil
}

func (m MockDbSQL) GetNextRows() bool {
	return m.WantNextRows
}

func (m MockDbSQL) ExecContext(ctx context.Context, query string, fields ...interface{}) (sql.Result, error) {
	if m.WantErr != nil && strings.EqualFold("errExecContext", m.WantErr.Error()) {
		return nil, m.WantErr
	}
	return m.WantResult, nil
}

func (m MockDbSQL) ExecContextStmt(ctx context.Context, fields ...interface{}) (sql.Result, error) {
	if m.WantErr != nil && strings.EqualFold("errExecContextStmt", m.WantErr.Error()) {
		return nil, m.WantErr
	}

	return m.WantResult, nil
}

func (m MockDbSQL) Query(query string, args ...interface{}) error {
	if m.WantErr != nil && strings.EqualFold("errQuery", m.WantErr.Error()) {
		return m.WantErr
	}
	return nil
}

func (m MockDbSQL) QueryStmt(args ...interface{}) (*sql.Rows, error) {
	if m.WantErr != nil && strings.EqualFold("errQueryStmt", m.WantErr.Error()) {
		return nil, m.WantErr
	}
	return m.WantRows, nil
}

func (m MockDbSQL) QueryRow(args ...interface{}) {
}

func (m MockDbSQL) CloseStmt() error {
	if m.WantErr != nil && strings.EqualFold("errCloseStmt", m.WantErr.Error()) {
		return m.WantErr
	}
	return nil

}

func (m MockDbSQL) Scan(args ...interface{}) error {
	if m.WantErr != nil && strings.EqualFold("errScan", m.WantErr.Error()) {
		return m.WantErr
	}

	return nil
}

func (m MockDbSQL) ScanStmt(args ...interface{}) error {
	if m.WantErr != nil && strings.EqualFold("errScanStmt", m.WantErr.Error()) {
		return m.WantErr
	}

	return nil
}

// Mock DB NoSQL

type MockDbNoSQL struct {
	WantResultScan    *dynamodb.ScanOutput
	WantResultPutItem *dynamodb.PutItemOutput
	WantErr           error
}

func (m MockDbNoSQL) Scan(input *dynamodb.ScanInput) (*dynamodb.ScanOutput, error) {
	if m.WantErr != nil && strings.EqualFold("errScan", m.WantErr.Error()) {
		return nil, m.WantErr
	}

	return m.WantResultScan, nil

}
func (m MockDbNoSQL) PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	if m.WantErr != nil && strings.EqualFold("errPutItem", m.WantErr.Error()) {
		return nil, m.WantErr
	}

	return m.WantResultPutItem, nil
}
