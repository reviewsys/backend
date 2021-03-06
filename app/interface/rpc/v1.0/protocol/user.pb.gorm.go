// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: app/interface/rpc/v1.0/protocol/user.proto

package protocol

import context "context"
import time "time"

import auth1 "github.com/infobloxopen/atlas-app-toolkit/auth"
import errors1 "github.com/infobloxopen/protoc-gen-gorm/errors"
import field_mask1 "google.golang.org/genproto/protobuf/field_mask"
import gorm1 "github.com/jinzhu/gorm"
import gorm2 "github.com/infobloxopen/atlas-app-toolkit/gorm"
import ptypes1 "github.com/golang/protobuf/ptypes"
import query1 "github.com/infobloxopen/atlas-app-toolkit/query"
import resource1 "github.com/infobloxopen/atlas-app-toolkit/gorm/resource"

import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/timestamp"
import _ "google.golang.org/genproto/protobuf/field_mask"
import _ "github.com/infobloxopen/atlas-app-toolkit/query"
import _ "github.com/infobloxopen/atlas-app-toolkit/rpc/resource"

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = math.Inf

type UserORM struct {
	AccountID string
	CreatedAt time.Time
	DeletedAt *time.Time
	Id        string `gorm:"type:uuid;primary_key"`
	IsAdmin   bool
	Name      string
	TeamId    int64
	UpdatedAt time.Time
}

// TableName overrides the default tablename generated by GORM
func (UserORM) TableName() string {
	return "users"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *User) ToORM(ctx context.Context) (UserORM, error) {
	to := UserORM{}
	var err error
	if prehook, ok := interface{}(m).(UserWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	if v, err := resource1.Decode(&User{}, m.Id); err != nil {
		return to, err
	} else if v != nil {
		to.Id = v.(string)
	}
	if m.CreatedAt != nil {
		var t time.Time
		if t, err = ptypes1.Timestamp(m.CreatedAt); err != nil {
			return to, err
		}
		to.CreatedAt = t
	}
	if m.UpdatedAt != nil {
		var t time.Time
		if t, err = ptypes1.Timestamp(m.UpdatedAt); err != nil {
			return to, err
		}
		to.UpdatedAt = t
	}
	if m.DeletedAt != nil {
		var t time.Time
		if t, err = ptypes1.Timestamp(m.DeletedAt); err != nil {
			return to, err
		}
		to.DeletedAt = &t
	}
	to.TeamId = m.TeamId
	to.Name = m.Name
	to.IsAdmin = m.IsAdmin
	accountID, err := auth1.GetAccountID(ctx, nil)
	if err != nil {
		return to, err
	}
	to.AccountID = accountID
	if posthook, ok := interface{}(m).(UserWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *UserORM) ToPB(ctx context.Context) (User, error) {
	to := User{}
	var err error
	if prehook, ok := interface{}(m).(UserWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	if v, err := resource1.Encode(&User{}, m.Id); err != nil {
		return to, err
	} else {
		to.Id = v
	}
	if to.CreatedAt, err = ptypes1.TimestampProto(m.CreatedAt); err != nil {
		return to, err
	}
	if to.UpdatedAt, err = ptypes1.TimestampProto(m.UpdatedAt); err != nil {
		return to, err
	}
	if m.DeletedAt != nil {
		if to.DeletedAt, err = ptypes1.TimestampProto(*m.DeletedAt); err != nil {
			return to, err
		}
	}
	to.TeamId = m.TeamId
	to.Name = m.Name
	to.IsAdmin = m.IsAdmin
	if posthook, ok := interface{}(m).(UserWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type User the arg will be the target, the caller the one being converted from

// UserBeforeToORM called before default ToORM code
type UserWithBeforeToORM interface {
	BeforeToORM(context.Context, *UserORM) error
}

// UserAfterToORM called after default ToORM code
type UserWithAfterToORM interface {
	AfterToORM(context.Context, *UserORM) error
}

// UserBeforeToPB called before default ToPB code
type UserWithBeforeToPB interface {
	BeforeToPB(context.Context, *User) error
}

// UserAfterToPB called after default ToPB code
type UserWithAfterToPB interface {
	AfterToPB(context.Context, *User) error
}

// DefaultCreateUser executes a basic gorm create call
func DefaultCreateUser(ctx context.Context, in *User, db *gorm1.DB) (*User, error) {
	if in == nil {
		return nil, errors1.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type UserORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type UserORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm1.DB) error
}

// DefaultReadUser executes a basic gorm read call
func DefaultReadUser(ctx context.Context, in *User, db *gorm1.DB) (*User, error) {
	if in == nil {
		return nil, errors1.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if ormObj.Id == "" {
		return nil, errors1.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithBeforeReadApplyQuery); ok {
		if db, err = hook.BeforeReadApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if db, err = gorm2.ApplyFieldSelection(ctx, db, nil, &UserORM{}); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithBeforeReadFind); ok {
		if db, err = hook.BeforeReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	ormResponse := UserORM{}
	if err = db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(UserORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

type UserORMWithBeforeReadApplyQuery interface {
	BeforeReadApplyQuery(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type UserORMWithBeforeReadFind interface {
	BeforeReadFind(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type UserORMWithAfterReadFind interface {
	AfterReadFind(context.Context, *gorm1.DB) error
}

func DefaultDeleteUser(ctx context.Context, in *User, db *gorm1.DB) error {
	if in == nil {
		return errors1.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return err
	}
	if ormObj.Id == "" {
		return errors1.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithBeforeDelete_); ok {
		if db, err = hook.BeforeDelete_(ctx, db); err != nil {
			return err
		}
	}
	err = db.Where(&ormObj).Delete(&UserORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, db)
	}
	return err
}

type UserORMWithBeforeDelete_ interface {
	BeforeDelete_(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type UserORMWithAfterDelete_ interface {
	AfterDelete_(context.Context, *gorm1.DB) error
}

func DefaultDeleteUserSet(ctx context.Context, in []*User, db *gorm1.DB) error {
	if in == nil {
		return errors1.NilArgumentError
	}
	var err error
	keys := []string{}
	for _, obj := range in {
		ormObj, err := obj.ToORM(ctx)
		if err != nil {
			return err
		}
		if ormObj.Id == "" {
			return errors1.EmptyIdError
		}
		keys = append(keys, ormObj.Id)
	}
	if hook, ok := (interface{}(&UserORM{})).(UserORMWithBeforeDeleteSet); ok {
		if db, err = hook.BeforeDeleteSet(ctx, in, db); err != nil {
			return err
		}
	}
	acctId, err := auth1.GetAccountID(ctx, nil)
	if err != nil {
		return err
	}
	err = db.Where("account_id = ? AND id in (?)", acctId, keys).Delete(&UserORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := (interface{}(&UserORM{})).(UserORMWithAfterDeleteSet); ok {
		err = hook.AfterDeleteSet(ctx, in, db)
	}
	return err
}

type UserORMWithBeforeDeleteSet interface {
	BeforeDeleteSet(context.Context, []*User, *gorm1.DB) (*gorm1.DB, error)
}
type UserORMWithAfterDeleteSet interface {
	AfterDeleteSet(context.Context, []*User, *gorm1.DB) error
}

// DefaultStrictUpdateUser clears first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateUser(ctx context.Context, in *User, db *gorm1.DB) (*User, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateUser")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	accountID, err := auth1.GetAccountID(ctx, nil)
	if err != nil {
		return nil, err
	}
	db = db.Where(map[string]interface{}{"account_id": accountID})
	lockedRow := &UserORM{}
	db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("id=?", ormObj.Id).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(UserORMWithBeforeStrictUpdateCleanup); ok {
		if db, err = hook.BeforeStrictUpdateCleanup(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithBeforeStrictUpdateSave); ok {
		if db, err = hook.BeforeStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithAfterStrictUpdateSave); ok {
		if err = hook.AfterStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	if err != nil {
		return nil, err
	}
	return &pbResponse, err
}

type UserORMWithBeforeStrictUpdateCleanup interface {
	BeforeStrictUpdateCleanup(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type UserORMWithBeforeStrictUpdateSave interface {
	BeforeStrictUpdateSave(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type UserORMWithAfterStrictUpdateSave interface {
	AfterStrictUpdateSave(context.Context, *gorm1.DB) error
}

// DefaultPatchUser executes a basic gorm update call with patch behavior
func DefaultPatchUser(ctx context.Context, in *User, updateMask *field_mask1.FieldMask, db *gorm1.DB) (*User, error) {
	if in == nil {
		return nil, errors1.NilArgumentError
	}
	var pbObj User
	var err error
	if hook, ok := interface{}(&pbObj).(UserWithBeforePatchRead); ok {
		if db, err = hook.BeforePatchRead(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbReadRes, err := DefaultReadUser(ctx, &User{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes
	if hook, ok := interface{}(&pbObj).(UserWithBeforePatchApplyFieldMask); ok {
		if db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if _, err := DefaultApplyFieldMaskUser(ctx, &pbObj, in, updateMask, "", db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(UserWithBeforePatchSave); ok {
		if db, err = hook.BeforePatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := DefaultStrictUpdateUser(ctx, &pbObj, db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(UserWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

type UserWithBeforePatchRead interface {
	BeforePatchRead(context.Context, *User, *field_mask1.FieldMask, *gorm1.DB) (*gorm1.DB, error)
}
type UserWithBeforePatchApplyFieldMask interface {
	BeforePatchApplyFieldMask(context.Context, *User, *field_mask1.FieldMask, *gorm1.DB) (*gorm1.DB, error)
}
type UserWithBeforePatchSave interface {
	BeforePatchSave(context.Context, *User, *field_mask1.FieldMask, *gorm1.DB) (*gorm1.DB, error)
}
type UserWithAfterPatchSave interface {
	AfterPatchSave(context.Context, *User, *field_mask1.FieldMask, *gorm1.DB) error
}

// DefaultApplyFieldMaskUser patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskUser(ctx context.Context, patchee *User, patcher *User, updateMask *field_mask1.FieldMask, prefix string, db *gorm1.DB) (*User, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors1.NilArgumentError
	}
	var err error
	for _, f := range updateMask.Paths {
		if f == prefix+"Id" {
			patchee.Id = patcher.Id
			continue
		}
		if f == prefix+"CreatedAt" {
			patchee.CreatedAt = patcher.CreatedAt
			continue
		}
		if f == prefix+"UpdatedAt" {
			patchee.UpdatedAt = patcher.UpdatedAt
			continue
		}
		if f == prefix+"DeletedAt" {
			patchee.DeletedAt = patcher.DeletedAt
			continue
		}
		if f == prefix+"TeamId" {
			patchee.TeamId = patcher.TeamId
			continue
		}
		if f == prefix+"Name" {
			patchee.Name = patcher.Name
			continue
		}
		if f == prefix+"IsAdmin" {
			patchee.IsAdmin = patcher.IsAdmin
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListUser executes a gorm list call
func DefaultListUser(ctx context.Context, db *gorm1.DB, f *query1.Filtering, s *query1.Sorting, p *query1.Pagination, fs *query1.FieldSelection) ([]*User, error) {
	in := User{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db, f, s, p, fs); err != nil {
			return nil, err
		}
	}
	db, err = gorm2.ApplyCollectionOperators(ctx, db, &UserORM{}, &User{}, f, s, p, fs)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db, f, s, p, fs); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("id")
	ormResponse := []UserORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse, f, s, p, fs); err != nil {
			return nil, err
		}
	}
	pbResponse := []*User{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type UserORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm1.DB, *query1.Filtering, *query1.Sorting, *query1.Pagination, *query1.FieldSelection) (*gorm1.DB, error)
}
type UserORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm1.DB, *query1.Filtering, *query1.Sorting, *query1.Pagination, *query1.FieldSelection) (*gorm1.DB, error)
}
type UserORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm1.DB, *[]UserORM, *query1.Filtering, *query1.Sorting, *query1.Pagination, *query1.FieldSelection) error
}
type UserServiceDefaultServer struct {
	DB *gorm1.DB
}

// Create ...
func (m *UserServiceDefaultServer) Create(ctx context.Context, in *CreateUserRequest) (*CreateUserResponse, error) {
	db := m.DB
	if custom, ok := interface{}(in).(UserServiceUserWithBeforeCreate); ok {
		var err error
		if db, err = custom.BeforeCreate(ctx, db); err != nil {
			return nil, err
		}
	}
	res, err := DefaultCreateUser(ctx, in.GetPayload(), db)
	if err != nil {
		return nil, err
	}
	out := &CreateUserResponse{Result: res}
	if custom, ok := interface{}(in).(UserServiceUserWithAfterCreate); ok {
		var err error
		if err = custom.AfterCreate(ctx, out, db); err != nil {
			return nil, err
		}
	}
	return out, nil
}

// UserServiceUserWithBeforeCreate called before DefaultCreateUser in the default Create handler
type UserServiceUserWithBeforeCreate interface {
	BeforeCreate(context.Context, *gorm1.DB) (*gorm1.DB, error)
}

// UserServiceUserWithAfterCreate called before DefaultCreateUser in the default Create handler
type UserServiceUserWithAfterCreate interface {
	AfterCreate(context.Context, *CreateUserResponse, *gorm1.DB) error
}

// Read ...
func (m *UserServiceDefaultServer) Read(ctx context.Context, in *ReadUserRequest) (*ReadUserResponse, error) {
	db := m.DB
	if custom, ok := interface{}(in).(UserServiceUserWithBeforeRead); ok {
		var err error
		if db, err = custom.BeforeRead(ctx, db); err != nil {
			return nil, err
		}
	}
	res, err := DefaultReadUser(ctx, &User{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	out := &ReadUserResponse{Result: res}
	if custom, ok := interface{}(in).(UserServiceUserWithAfterRead); ok {
		var err error
		if err = custom.AfterRead(ctx, out, db); err != nil {
			return nil, err
		}
	}
	return out, nil
}

// UserServiceUserWithBeforeRead called before DefaultReadUser in the default Read handler
type UserServiceUserWithBeforeRead interface {
	BeforeRead(context.Context, *gorm1.DB) (*gorm1.DB, error)
}

// UserServiceUserWithAfterRead called before DefaultReadUser in the default Read handler
type UserServiceUserWithAfterRead interface {
	AfterRead(context.Context, *ReadUserResponse, *gorm1.DB) error
}

// Update ...
func (m *UserServiceDefaultServer) Update(ctx context.Context, in *UpdateUserRequest) (*UpdateUserResponse, error) {
	var err error
	var res *User
	db := m.DB
	if custom, ok := interface{}(in).(UserServiceUserWithBeforeUpdate); ok {
		var err error
		if db, err = custom.BeforeUpdate(ctx, db); err != nil {
			return nil, err
		}
	}
	if in.GetGerogeriGegege() == nil {
		res, err = DefaultStrictUpdateUser(ctx, in.GetPayload(), db)
	} else {
		res, err = DefaultPatchUser(ctx, in.GetPayload(), in.GetGerogeriGegege(), db)
	}
	if err != nil {
		return nil, err
	}
	out := &UpdateUserResponse{Result: res}
	if custom, ok := interface{}(in).(UserServiceUserWithAfterUpdate); ok {
		var err error
		if err = custom.AfterUpdate(ctx, out, db); err != nil {
			return nil, err
		}
	}
	return out, nil
}

// UserServiceUserWithBeforeUpdate called before DefaultUpdateUser in the default Update handler
type UserServiceUserWithBeforeUpdate interface {
	BeforeUpdate(context.Context, *gorm1.DB) (*gorm1.DB, error)
}

// UserServiceUserWithAfterUpdate called before DefaultUpdateUser in the default Update handler
type UserServiceUserWithAfterUpdate interface {
	AfterUpdate(context.Context, *UpdateUserResponse, *gorm1.DB) error
}

// List ...
func (m *UserServiceDefaultServer) List(ctx context.Context, in *ListUserRequest) (*ListUserResponse, error) {
	db := m.DB
	if custom, ok := interface{}(in).(UserServiceUserWithBeforeList); ok {
		var err error
		if db, err = custom.BeforeList(ctx, db); err != nil {
			return nil, err
		}
	}
	pagedRequest := false
	if in.GetPaging().GetLimit() >= 1 {
		in.Paging.Limit++
		pagedRequest = true
	}
	res, err := DefaultListUser(ctx, db, in.Filter, in.OrderBy, in.Paging, in.Fields)
	if err != nil {
		return nil, err
	}
	var resPaging *query1.PageInfo
	if pagedRequest {
		var offset int32
		var size int32 = int32(len(res))
		if size == in.GetPaging().GetLimit() {
			size--
			res = res[:size]
			offset = in.GetPaging().GetOffset() + size
		}
		resPaging = &query1.PageInfo{Offset: offset}
	}
	out := &ListUserResponse{Results: res, PageInfo: resPaging}
	if custom, ok := interface{}(in).(UserServiceUserWithAfterList); ok {
		var err error
		if err = custom.AfterList(ctx, out, db); err != nil {
			return nil, err
		}
	}
	return out, nil
}

// UserServiceUserWithBeforeList called before DefaultListUser in the default List handler
type UserServiceUserWithBeforeList interface {
	BeforeList(context.Context, *gorm1.DB) (*gorm1.DB, error)
}

// UserServiceUserWithAfterList called before DefaultListUser in the default List handler
type UserServiceUserWithAfterList interface {
	AfterList(context.Context, *ListUserResponse, *gorm1.DB) error
}

// Delete ...
func (m *UserServiceDefaultServer) Delete(ctx context.Context, in *DeleteUserRequest) (*DeleteUserResponse, error) {
	db := m.DB
	if custom, ok := interface{}(in).(UserServiceUserWithBeforeDelete); ok {
		var err error
		if db, err = custom.BeforeDelete(ctx, db); err != nil {
			return nil, err
		}
	}
	err := DefaultDeleteUser(ctx, &User{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	out := &DeleteUserResponse{}
	if custom, ok := interface{}(in).(UserServiceUserWithAfterDelete); ok {
		var err error
		if err = custom.AfterDelete(ctx, out, db); err != nil {
			return nil, err
		}
	}
	return out, nil
}

// UserServiceUserWithBeforeDelete called before DefaultDeleteUser in the default Delete handler
type UserServiceUserWithBeforeDelete interface {
	BeforeDelete(context.Context, *gorm1.DB) (*gorm1.DB, error)
}

// UserServiceUserWithAfterDelete called before DefaultDeleteUser in the default Delete handler
type UserServiceUserWithAfterDelete interface {
	AfterDelete(context.Context, *DeleteUserResponse, *gorm1.DB) error
}
