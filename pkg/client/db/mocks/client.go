// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

package dbmocks

//go:generate minimock -i github.com/a1exCross/common/pkg/client/db.Client -o client.go -n ClientMock -p mocks

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	mm_db "github.com/a1exCross/common/pkg/client/db"
	"github.com/gojuno/minimock/v3"
)

// ClientMock implements db.Client
type ClientMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcClose          func() (err error)
	inspectFuncClose   func()
	afterCloseCounter  uint64
	beforeCloseCounter uint64
	CloseMock          mClientMockClose

	funcDB          func() (d1 mm_db.DB)
	inspectFuncDB   func()
	afterDBCounter  uint64
	beforeDBCounter uint64
	DBMock          mClientMockDB
}

// NewClientMock returns a mock for db.Client
func NewClientMock(t minimock.Tester) *ClientMock {
	m := &ClientMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CloseMock = mClientMockClose{mock: m}

	m.DBMock = mClientMockDB{mock: m}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mClientMockClose struct {
	mock               *ClientMock
	defaultExpectation *ClientMockCloseExpectation
	expectations       []*ClientMockCloseExpectation
}

// ClientMockCloseExpectation specifies expectation struct of the Client.Close
type ClientMockCloseExpectation struct {
	mock *ClientMock

	results *ClientMockCloseResults
	Counter uint64
}

// ClientMockCloseResults contains results of the Client.Close
type ClientMockCloseResults struct {
	err error
}

// Expect sets up expected params for Client.Close
func (mmClose *mClientMockClose) Expect() *mClientMockClose {
	if mmClose.mock.funcClose != nil {
		mmClose.mock.t.Fatalf("ClientMock.Close mock is already set by Set")
	}

	if mmClose.defaultExpectation == nil {
		mmClose.defaultExpectation = &ClientMockCloseExpectation{}
	}

	return mmClose
}

// Inspect accepts an inspector function that has same arguments as the Client.Close
func (mmClose *mClientMockClose) Inspect(f func()) *mClientMockClose {
	if mmClose.mock.inspectFuncClose != nil {
		mmClose.mock.t.Fatalf("Inspect function is already set for ClientMock.Close")
	}

	mmClose.mock.inspectFuncClose = f

	return mmClose
}

// Return sets up results that will be returned by Client.Close
func (mmClose *mClientMockClose) Return(err error) *ClientMock {
	if mmClose.mock.funcClose != nil {
		mmClose.mock.t.Fatalf("ClientMock.Close mock is already set by Set")
	}

	if mmClose.defaultExpectation == nil {
		mmClose.defaultExpectation = &ClientMockCloseExpectation{mock: mmClose.mock}
	}
	mmClose.defaultExpectation.results = &ClientMockCloseResults{err}
	return mmClose.mock
}

// Set uses given function f to mock the Client.Close method
func (mmClose *mClientMockClose) Set(f func() (err error)) *ClientMock {
	if mmClose.defaultExpectation != nil {
		mmClose.mock.t.Fatalf("Default expectation is already set for the Client.Close method")
	}

	if len(mmClose.expectations) > 0 {
		mmClose.mock.t.Fatalf("Some expectations are already set for the Client.Close method")
	}

	mmClose.mock.funcClose = f
	return mmClose.mock
}

// Close implements db.Client
func (mmClose *ClientMock) Close() (err error) {
	mm_atomic.AddUint64(&mmClose.beforeCloseCounter, 1)
	defer mm_atomic.AddUint64(&mmClose.afterCloseCounter, 1)

	if mmClose.inspectFuncClose != nil {
		mmClose.inspectFuncClose()
	}

	if mmClose.CloseMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmClose.CloseMock.defaultExpectation.Counter, 1)

		mm_results := mmClose.CloseMock.defaultExpectation.results
		if mm_results == nil {
			mmClose.t.Fatal("No results are set for the ClientMock.Close")
		}
		return (*mm_results).err
	}
	if mmClose.funcClose != nil {
		return mmClose.funcClose()
	}
	mmClose.t.Fatalf("Unexpected call to ClientMock.Close.")
	return
}

// CloseAfterCounter returns a count of finished ClientMock.Close invocations
func (mmClose *ClientMock) CloseAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmClose.afterCloseCounter)
}

// CloseBeforeCounter returns a count of ClientMock.Close invocations
func (mmClose *ClientMock) CloseBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmClose.beforeCloseCounter)
}

// MinimockCloseDone returns true if the count of the Close invocations corresponds
// the number of defined expectations
func (m *ClientMock) MinimockCloseDone() bool {
	for _, e := range m.CloseMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CloseMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCloseCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcClose != nil && mm_atomic.LoadUint64(&m.afterCloseCounter) < 1 {
		return false
	}
	return true
}

// MinimockCloseInspect logs each unmet expectation
func (m *ClientMock) MinimockCloseInspect() {
	for _, e := range m.CloseMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to ClientMock.Close")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CloseMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCloseCounter) < 1 {
		m.t.Error("Expected call to ClientMock.Close")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcClose != nil && mm_atomic.LoadUint64(&m.afterCloseCounter) < 1 {
		m.t.Error("Expected call to ClientMock.Close")
	}
}

type mClientMockDB struct {
	mock               *ClientMock
	defaultExpectation *ClientMockDBExpectation
	expectations       []*ClientMockDBExpectation
}

// ClientMockDBExpectation specifies expectation struct of the Client.DB
type ClientMockDBExpectation struct {
	mock *ClientMock

	results *ClientMockDBResults
	Counter uint64
}

// ClientMockDBResults contains results of the Client.DB
type ClientMockDBResults struct {
	d1 mm_db.DB
}

// Expect sets up expected params for Client.DB
func (mmDB *mClientMockDB) Expect() *mClientMockDB {
	if mmDB.mock.funcDB != nil {
		mmDB.mock.t.Fatalf("ClientMock.DB mock is already set by Set")
	}

	if mmDB.defaultExpectation == nil {
		mmDB.defaultExpectation = &ClientMockDBExpectation{}
	}

	return mmDB
}

// Inspect accepts an inspector function that has same arguments as the Client.DB
func (mmDB *mClientMockDB) Inspect(f func()) *mClientMockDB {
	if mmDB.mock.inspectFuncDB != nil {
		mmDB.mock.t.Fatalf("Inspect function is already set for ClientMock.DB")
	}

	mmDB.mock.inspectFuncDB = f

	return mmDB
}

// Return sets up results that will be returned by Client.DB
func (mmDB *mClientMockDB) Return(d1 mm_db.DB) *ClientMock {
	if mmDB.mock.funcDB != nil {
		mmDB.mock.t.Fatalf("ClientMock.DB mock is already set by Set")
	}

	if mmDB.defaultExpectation == nil {
		mmDB.defaultExpectation = &ClientMockDBExpectation{mock: mmDB.mock}
	}
	mmDB.defaultExpectation.results = &ClientMockDBResults{d1}
	return mmDB.mock
}

// Set uses given function f to mock the Client.DB method
func (mmDB *mClientMockDB) Set(f func() (d1 mm_db.DB)) *ClientMock {
	if mmDB.defaultExpectation != nil {
		mmDB.mock.t.Fatalf("Default expectation is already set for the Client.DB method")
	}

	if len(mmDB.expectations) > 0 {
		mmDB.mock.t.Fatalf("Some expectations are already set for the Client.DB method")
	}

	mmDB.mock.funcDB = f
	return mmDB.mock
}

// DB implements db.Client
func (mmDB *ClientMock) DB() (d1 mm_db.DB) {
	mm_atomic.AddUint64(&mmDB.beforeDBCounter, 1)
	defer mm_atomic.AddUint64(&mmDB.afterDBCounter, 1)

	if mmDB.inspectFuncDB != nil {
		mmDB.inspectFuncDB()
	}

	if mmDB.DBMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmDB.DBMock.defaultExpectation.Counter, 1)

		mm_results := mmDB.DBMock.defaultExpectation.results
		if mm_results == nil {
			mmDB.t.Fatal("No results are set for the ClientMock.DB")
		}
		return (*mm_results).d1
	}
	if mmDB.funcDB != nil {
		return mmDB.funcDB()
	}
	mmDB.t.Fatalf("Unexpected call to ClientMock.DB.")
	return
}

// DBAfterCounter returns a count of finished ClientMock.DB invocations
func (mmDB *ClientMock) DBAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDB.afterDBCounter)
}

// DBBeforeCounter returns a count of ClientMock.DB invocations
func (mmDB *ClientMock) DBBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDB.beforeDBCounter)
}

// MinimockDBDone returns true if the count of the DB invocations corresponds
// the number of defined expectations
func (m *ClientMock) MinimockDBDone() bool {
	for _, e := range m.DBMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DBMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDBCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDB != nil && mm_atomic.LoadUint64(&m.afterDBCounter) < 1 {
		return false
	}
	return true
}

// MinimockDBInspect logs each unmet expectation
func (m *ClientMock) MinimockDBInspect() {
	for _, e := range m.DBMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to ClientMock.DB")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DBMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDBCounter) < 1 {
		m.t.Error("Expected call to ClientMock.DB")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDB != nil && mm_atomic.LoadUint64(&m.afterDBCounter) < 1 {
		m.t.Error("Expected call to ClientMock.DB")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *ClientMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockCloseInspect()

			m.MinimockDBInspect()
			m.t.FailNow()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *ClientMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *ClientMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCloseDone() &&
		m.MinimockDBDone()
}
