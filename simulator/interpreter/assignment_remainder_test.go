package interpreter

import (
	"net"
	"testing"
	"time"

	"github.com/ysugimoto/falco/ast"
	"github.com/ysugimoto/falco/simulator/variable"
)

func TestProcessRemainder(t *testing.T) {
	t.Run("left is INTEGER", func(t *testing.T) {
		now := time.Now()
		tests := []struct {
			left    int64
			right   variable.Value
			expect  int64
			isError bool
		}{
			{left: 1002,  right: &variable.Integer{Value: 100}, expect: 2},
			{left: 1002,  right: &variable.Integer{Value: 100, Literal: true}, expect: 2},
			{left: 1002,  right: &variable.Float{Value: 50.0}, expect: 2},
			{left: 1002,  right: &variable.Float{Value: 50.0, Literal: true}, isError: true},
			{left: 1002,  right: &variable.String{Value: "example"}, isError: true},
			{left: 1002,  right: &variable.String{Value: "example", Literal: true}, isError: true},
			{left: 1002,  right: &variable.RTime{Value: 100 * time.Second}, isError: true},
			{left: 1002,  right: &variable.RTime{Value: 100 * time.Second, Literal: true}, isError: true},
			{left: 1002,  right: &variable.Time{Value: now}, isError: true},
			{left: 1002,  right: &variable.Backend{Value: &ast.BackendDeclaration{Name: &ast.Ident{Value: "foo"}}}, isError: true},
			{left: 1002,  right: &variable.Boolean{Value: true}, isError: true},
			{left: 1002,  right: &variable.Boolean{Value: false, Literal: true}, isError: true},
			{left: 1002,  right: &variable.IP{Value: net.ParseIP("127.0.0.1")}, isError: true},
		}

		for i, tt := range tests {
			ip := New(nil)
			left := &variable.Integer{Value: tt.left}
			err := ip.ProcessRemainderAssignment(left, tt.right)
			if tt.isError {
				if err == nil {
					t.Errorf("Index %d: expects error but non-nil", i)
				}
				continue
			}
			if left.Value != tt.expect {
				t.Errorf("Index %d: expect value %d, got %d", i, tt.expect, left.Value)
			}
		}
	})

	t.Run("left is FLOAT", func(t *testing.T) {
		now := time.Now()
		tests := []struct {
			left    float64
			right   variable.Value
			expect  float64
			isError bool
		}{
			{left: 1002.0, right: &variable.Integer{Value: 100}, expect: 2.0},
			{left: 1002.0, right: &variable.Integer{Value: 100, Literal: true}, expect: 2.0},
			{left: 1002.0, right: &variable.Float{Value: 50.0}, expect: 2.0},
			{left: 1002.0, right: &variable.Float{Value: 50.0, Literal: true}, expect: 2.0},
			{left: 1002.0, right: &variable.String{Value: "example"}, isError: true},
			{left: 1002.0, right: &variable.String{Value: "example", Literal: true}, isError: true},
			{left: 1002.0, right: &variable.RTime{Value: 100 * time.Second}, isError: true},
			{left: 1002.0, right: &variable.RTime{Value: 100 * time.Second, Literal: true}, isError: true},
			{left: 1002.0, right: &variable.Time{Value: now}, isError: true},
			{left: 1002.0, right: &variable.Backend{Value: &ast.BackendDeclaration{Name: &ast.Ident{Value: "foo"}}}, isError: true},
			{left: 1002.0, right: &variable.Boolean{Value: true}, isError: true},
			{left: 1002.0, right: &variable.Boolean{Value: false, Literal: true}, isError: true},
			{left: 1002.0, right: &variable.IP{Value: net.ParseIP("127.0.0.1")}, isError: true},
		}

		for i, tt := range tests {
			ip := New(nil)
			left := &variable.Float{Value: tt.left}
			err := ip.ProcessRemainderAssignment(left, tt.right)
			if tt.isError {
				if err == nil {
					t.Errorf("Index %d: expects error but non-nil", i)
				}
				continue
			}
			if left.Value != tt.expect {
				t.Errorf("Index %d: expect value %.2f, got %.2f", i, tt.expect, left.Value)
			}
		}
	})

	t.Run("left is STRING", func(t *testing.T) {
		now := time.Now()
		tests := []struct {
			left    string
			right   variable.Value
			expect  string
			isError bool
		}{
			{left: "left", right: &variable.Integer{Value: 100}, isError: true},
			{left: "left", right: &variable.Integer{Value: 100, Literal: true}, isError: true},
			{left: "left", right: &variable.Float{Value: 50.0}, isError: true},
			{left: "left", right: &variable.Float{Value: 50.0, Literal: true}, isError: true},
			{left: "left", right: &variable.RTime{Value: 100 * time.Second}, isError: true},
			{left: "left", right: &variable.RTime{Value: 100 * time.Second, Literal: true}, isError: true},
			{left: "left", right: &variable.Time{Value: now}, isError: true},
			{left: "left", right: &variable.String{Value: "example"}, isError: true},
			{left: "left", right: &variable.String{Value: "example", Literal: true}, isError: true},
			{left: "left", right: &variable.Backend{Value: &ast.BackendDeclaration{Name: &ast.Ident{Value: "foo"}}}, isError: true},
			{left: "left", right: &variable.Boolean{Value: true}, isError: true},
			{left: "left", right: &variable.Boolean{Value: false, Literal: true}, isError: true},
			{left: "left", right: &variable.IP{Value: net.ParseIP("127.0.0.1")}, isError: true},
		}

		for i, tt := range tests {
			ip := New(nil)
			left := &variable.String{Value: tt.left}
			err := ip.ProcessRemainderAssignment(left, tt.right)
			if tt.isError {
				if err == nil {
					t.Errorf("Index %d: expects error but non-nil", i)
				}
				continue
			}
			if left.Value != tt.expect {
				t.Errorf("Index %d: expect value %s, got %s", i, tt.expect, left.Value)
			}
		}
	})

	t.Run("left is RTIME", func(t *testing.T) {
		now := time.Now()
		tests := []struct {
			left    time.Duration
			right   variable.Value
			expect  time.Duration
			isError bool
		}{
			{left: 1002 * time.Second, right: &variable.Integer{Value: 100}, expect: 2 * time.Second},
			{left: 1002 * time.Second, right: &variable.Integer{Value: 100, Literal: true}, expect: 2 * time.Second},
			{left: 1002 * time.Second, right: &variable.Float{Value: 50.0}, expect: 2 * time.Second},
			{left: 1002 * time.Second, right: &variable.Float{Value: 50.0, Literal: true}, expect: 2 * time.Second},
			{left: 1002 * time.Second, right: &variable.String{Value: "example"}, isError: true},
			{left: 1002 * time.Second, right: &variable.String{Value: "example", Literal: true}, isError: true},
			{left: 1002 * time.Second, right: &variable.RTime{Value: 100 * time.Second}, isError: true},
			{left: 1002 * time.Second, right: &variable.RTime{Value: 100 * time.Second, Literal: true}, isError: true},
			{left: 1002 * time.Second, right: &variable.Time{Value: now}, isError: true},
			{left: 1002 * time.Second, right: &variable.Backend{Value: &ast.BackendDeclaration{Name: &ast.Ident{Value: "foo"}}}, isError: true},
			{left: 1002 * time.Second, right: &variable.Boolean{Value: true}, isError: true},
			{left: 1002 * time.Second, right: &variable.Boolean{Value: false, Literal: true}, isError: true},
			{left: 1002 * time.Second, right: &variable.IP{Value: net.ParseIP("127.0.0.1")}, isError: true},
		}

		for i, tt := range tests {
			ip := New(nil)
			left := &variable.RTime{Value: tt.left}
			err := ip.ProcessRemainderAssignment(left, tt.right)
			if tt.isError {
				if err == nil {
					t.Errorf("Index %d: expects error but non-nil", i)
				}
				continue
			}
			if left.Value != tt.expect {
				t.Errorf("Index %d: expect value %s, got %s", i, tt.expect, left.Value)
			}
		}
	})

	t.Run("left is TIME", func(t *testing.T) {
		now := time.Now()
		now2 := now.Add(10 * time.Second)
		tests := []struct {
			left    time.Time
			right   variable.Value
			expect  time.Time
			isError bool
		}{
			{left: now, right: &variable.Integer{Value: 100}, isError: true},
			{left: now, right: &variable.Integer{Value: 100, Literal: true}, isError: true},
			{left: now, right: &variable.Float{Value: 50.0}, isError: true},
			{left: now, right: &variable.Float{Value: 50.0, Literal: true}, isError: true},
			{left: now, right: &variable.String{Value: "example"}, isError: true},
			{left: now, right: &variable.String{Value: "example", Literal: true}, isError: true},
			{left: now, right: &variable.RTime{Value: 100 * time.Second}, isError: true},
			{left: now, right: &variable.RTime{Value: 100 * time.Second, Literal: true}, isError: true},
			{left: now, right: &variable.Time{Value: now2}, isError: true},
			{left: now, right: &variable.Backend{Value: &ast.BackendDeclaration{Name: &ast.Ident{Value: "foo"}}}, isError: true},
			{left: now, right: &variable.Boolean{Value: true}, isError: true},
			{left: now, right: &variable.Boolean{Value: false, Literal: true}, isError: true},
			{left: now, right: &variable.IP{Value: net.ParseIP("127.0.0.1")}, isError: true},
		}

		for i, tt := range tests {
			ip := New(nil)
			left := &variable.Time{Value: tt.left}
			err := ip.ProcessRemainderAssignment(left, tt.right)
			if tt.isError {
				if err == nil {
					t.Errorf("Index %d: expects error but non-nil", i)
				}
				continue
			}
			if left.Value != tt.expect {
				t.Errorf("Index %d: expect value %s, got %s", i, tt.expect, left.Value)
			}
		}
	})

	t.Run("left is BACKEND", func(t *testing.T) {
		now := time.Now()
		tests := []struct {
			left    string
			right   variable.Value
			expect  string
			isError bool
		}{
			{left: "backend", right: &variable.Integer{Value: 100}, isError: true},
			{left: "backend", right: &variable.Integer{Value: 100, Literal: true}, isError: true},
			{left: "backend", right: &variable.Float{Value: 50.0}, isError: true},
			{left: "backend", right: &variable.Float{Value: 50.0, Literal: true}, isError: true},
			{left: "backend", right: &variable.String{Value: "example"}, isError: true},
			{left: "backend", right: &variable.String{Value: "example", Literal: true}, isError: true},
			{left: "backend", right: &variable.RTime{Value: 100 * time.Second}, isError: true},
			{left: "backend", right: &variable.RTime{Value: 100 * time.Second, Literal: true}, isError: true},
			{left: "backend", right: &variable.Time{Value: now}, isError: true},
			{left: "backend", right: &variable.Backend{Value: &ast.BackendDeclaration{Name: &ast.Ident{Value: "foo"}}}, isError: true},
			{left: "backend", right: &variable.Boolean{Value: true}, isError: true},
			{left: "backend", right: &variable.Boolean{Value: false, Literal: true}, isError: true},
			{left: "backend", right: &variable.IP{Value: net.ParseIP("127.0.0.1")}, isError: true},
		}

		for i, tt := range tests {
			ip := New(nil)
			left := &variable.Backend{Value: &ast.BackendDeclaration{Name: &ast.Ident{Value: tt.left}},
			}
			err := ip.ProcessRemainderAssignment(left, tt.right)
			if tt.isError {
				if err == nil {
					t.Errorf("Index %d: expects error but non-nil", i)
				}
				continue
			}
			if left.Value.Name.Value != tt.expect {
				t.Errorf("Index %d: expect value %s, got %s", i, tt.expect, left.Value.Name.Value)
			}
		}
	})

	t.Run("left is BOOL", func(t *testing.T) {
		now := time.Now()
		tests := []struct {
			left    bool
			right   variable.Value
			expect  bool
			isError bool
		}{
			{left: false, right: &variable.Integer{Value: 100}, isError: true},
			{left: false, right: &variable.Integer{Value: 100, Literal: true}, isError: true},
			{left: false, right: &variable.Float{Value: 50.0}, isError: true},
			{left: false, right: &variable.Float{Value: 50.0, Literal: true}, isError: true},
			{left: false, right: &variable.String{Value: "example"}, isError: true},
			{left: false, right: &variable.String{Value: "example", Literal: true}, isError: true},
			{left: false, right: &variable.RTime{Value: 100 * time.Second}, isError: true},
			{left: false, right: &variable.RTime{Value: 100 * time.Second, Literal: true}, isError: true},
			{left: false, right: &variable.Time{Value: now}, isError: true},
			{left: false, right: &variable.Backend{Value: &ast.BackendDeclaration{Name: &ast.Ident{Value: "foo"}}}, isError: true},
			{left: false, right: &variable.Boolean{Value: true}, isError: true},
			{left: false, right: &variable.Boolean{Value: true, Literal: true}, isError: true},
			{left: false, right: &variable.IP{Value: net.ParseIP("127.0.0.1")}, isError: true},
		}

		for i, tt := range tests {
			ip := New(nil)
			left := &variable.Boolean{Value: tt.left}
			err := ip.ProcessRemainderAssignment(left, tt.right)
			if tt.isError {
				if err == nil {
					t.Errorf("Index %d: expects error but non-nil", i)
				}
				continue
			}
			if left.Value != tt.expect {
				t.Errorf("Index %d: expect value %t, got %t", i, tt.expect, left.Value)
			}
		}
	})

	t.Run("left is IP", func(t *testing.T) {
		now := time.Now()
		v := net.ParseIP("127.0.0.1")
		vv := net.ParseIP("127.0.0.2")
		tests := []struct {
			left    net.IP
			right   variable.Value
			expect  net.IP
			isError bool
		}{
			{left: v, right: &variable.Integer{Value: 100}, isError: true},
			{left: v, right: &variable.Integer{Value: 100, Literal: true}, isError: true},
			{left: v, right: &variable.Float{Value: 50.0}, isError: true},
			{left: v, right: &variable.Float{Value: 50.0, Literal: true}, isError: true},
			{left: v, right: &variable.String{Value: "example"}, isError: true},
			{left: v, right: &variable.String{Value: "example", Literal: true}, isError: true},
			{left: v, right: &variable.RTime{Value: 100 * time.Second}, isError: true},
			{left: v, right: &variable.RTime{Value: 100 * time.Second, Literal: true}, isError: true},
			{left: v, right: &variable.Time{Value: now}, isError: true},
			{left: v, right: &variable.String{Value: "127.0.0.2", Literal: true}, isError: true},
			{left: v, right: &variable.Backend{Value: &ast.BackendDeclaration{Name: &ast.Ident{Value: "foo"}}}, isError: true},
			{left: v, right: &variable.Boolean{Value: true}, isError: true},
			{left: v, right: &variable.Boolean{Value: true, Literal: true}, isError: true},
			{left: v, right: &variable.IP{Value: vv}, isError: true},
		}

		for i, tt := range tests {
			ip := New(nil)
			left := &variable.IP{Value: tt.left}
			err := ip.ProcessRemainderAssignment(left, tt.right)
			if tt.isError {
				if err == nil {
					t.Errorf("Index %d: expects error but non-nil", i)
				}
				continue
			}
			if left.Value.String() != tt.expect.String() {
				t.Errorf("Index %d: expect value %s, got %s", i, tt.expect, left.Value)
			}
		}
	})
}
