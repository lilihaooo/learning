// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q              = new(Query)
	Bill           *bill
	Fund           *fund
	Lable          *lable
	Myteacher      *myteacher
	MyteacherLable *myteacherLable
	Operation      *operation
	Order          *order
	OrderSubject   *orderSubject
	Recharge       *recharge
	Role           *role
	Subject        *subject
	Teacher        *teacher
	User           *user
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	Bill = &Q.Bill
	Fund = &Q.Fund
	Lable = &Q.Lable
	Myteacher = &Q.Myteacher
	MyteacherLable = &Q.MyteacherLable
	Operation = &Q.Operation
	Order = &Q.Order
	OrderSubject = &Q.OrderSubject
	Recharge = &Q.Recharge
	Role = &Q.Role
	Subject = &Q.Subject
	Teacher = &Q.Teacher
	User = &Q.User
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:             db,
		Bill:           newBill(db, opts...),
		Fund:           newFund(db, opts...),
		Lable:          newLable(db, opts...),
		Myteacher:      newMyteacher(db, opts...),
		MyteacherLable: newMyteacherLable(db, opts...),
		Operation:      newOperation(db, opts...),
		Order:          newOrder(db, opts...),
		OrderSubject:   newOrderSubject(db, opts...),
		Recharge:       newRecharge(db, opts...),
		Role:           newRole(db, opts...),
		Subject:        newSubject(db, opts...),
		Teacher:        newTeacher(db, opts...),
		User:           newUser(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	Bill           bill
	Fund           fund
	Lable          lable
	Myteacher      myteacher
	MyteacherLable myteacherLable
	Operation      operation
	Order          order
	OrderSubject   orderSubject
	Recharge       recharge
	Role           role
	Subject        subject
	Teacher        teacher
	User           user
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:             db,
		Bill:           q.Bill.clone(db),
		Fund:           q.Fund.clone(db),
		Lable:          q.Lable.clone(db),
		Myteacher:      q.Myteacher.clone(db),
		MyteacherLable: q.MyteacherLable.clone(db),
		Operation:      q.Operation.clone(db),
		Order:          q.Order.clone(db),
		OrderSubject:   q.OrderSubject.clone(db),
		Recharge:       q.Recharge.clone(db),
		Role:           q.Role.clone(db),
		Subject:        q.Subject.clone(db),
		Teacher:        q.Teacher.clone(db),
		User:           q.User.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.clone(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.clone(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:             db,
		Bill:           q.Bill.replaceDB(db),
		Fund:           q.Fund.replaceDB(db),
		Lable:          q.Lable.replaceDB(db),
		Myteacher:      q.Myteacher.replaceDB(db),
		MyteacherLable: q.MyteacherLable.replaceDB(db),
		Operation:      q.Operation.replaceDB(db),
		Order:          q.Order.replaceDB(db),
		OrderSubject:   q.OrderSubject.replaceDB(db),
		Recharge:       q.Recharge.replaceDB(db),
		Role:           q.Role.replaceDB(db),
		Subject:        q.Subject.replaceDB(db),
		Teacher:        q.Teacher.replaceDB(db),
		User:           q.User.replaceDB(db),
	}
}

type queryCtx struct {
	Bill           IBillDo
	Fund           IFundDo
	Lable          ILableDo
	Myteacher      IMyteacherDo
	MyteacherLable IMyteacherLableDo
	Operation      IOperationDo
	Order          IOrderDo
	OrderSubject   IOrderSubjectDo
	Recharge       IRechargeDo
	Role           IRoleDo
	Subject        ISubjectDo
	Teacher        ITeacherDo
	User           IUserDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		Bill:           q.Bill.WithContext(ctx),
		Fund:           q.Fund.WithContext(ctx),
		Lable:          q.Lable.WithContext(ctx),
		Myteacher:      q.Myteacher.WithContext(ctx),
		MyteacherLable: q.MyteacherLable.WithContext(ctx),
		Operation:      q.Operation.WithContext(ctx),
		Order:          q.Order.WithContext(ctx),
		OrderSubject:   q.OrderSubject.WithContext(ctx),
		Recharge:       q.Recharge.WithContext(ctx),
		Role:           q.Role.WithContext(ctx),
		Subject:        q.Subject.WithContext(ctx),
		Teacher:        q.Teacher.WithContext(ctx),
		User:           q.User.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	return &QueryTx{q.clone(q.db.Begin(opts...))}
}

type QueryTx struct{ *Query }

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
