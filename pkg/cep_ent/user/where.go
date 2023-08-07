// Code generated by ent, DO NOT EDIT.

package user

import (
	"cephalon-ent/pkg/cep_ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v int64) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v int64) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUpdatedBy, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldDeletedAt, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldName, v))
}

// NickName applies equality check predicate on the "nick_name" field. It's identical to NickNameEQ.
func NickName(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldNickName, v))
}

// Phone applies equality check predicate on the "phone" field. It's identical to PhoneEQ.
func Phone(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPhone, v))
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPassword, v))
}

// AvatarURL applies equality check predicate on the "avatar_url" field. It's identical to AvatarURLEQ.
func AvatarURL(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldAvatarURL, v))
}

// Platform applies equality check predicate on the "platform" field. It's identical to PlatformEQ.
func Platform(v int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPlatform, v))
}

// HmacKey applies equality check predicate on the "hmac_key" field. It's identical to HmacKeyEQ.
func HmacKey(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldHmacKey, v))
}

// HmacSecret applies equality check predicate on the "hmac_secret" field. It's identical to HmacSecretEQ.
func HmacSecret(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldHmacSecret, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v int64) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v int64) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...int64) predicate.User {
	return predicate.User(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...int64) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v int64) predicate.User {
	return predicate.User(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v int64) predicate.User {
	return predicate.User(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v int64) predicate.User {
	return predicate.User(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v int64) predicate.User {
	return predicate.User(sql.FieldLTE(FieldCreatedBy, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v int64) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v int64) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...int64) predicate.User {
	return predicate.User(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...int64) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v int64) predicate.User {
	return predicate.User(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v int64) predicate.User {
	return predicate.User(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v int64) predicate.User {
	return predicate.User(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v int64) predicate.User {
	return predicate.User(sql.FieldLTE(FieldUpdatedBy, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldDeletedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldName, v))
}

// NickNameEQ applies the EQ predicate on the "nick_name" field.
func NickNameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldNickName, v))
}

// NickNameNEQ applies the NEQ predicate on the "nick_name" field.
func NickNameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldNickName, v))
}

// NickNameIn applies the In predicate on the "nick_name" field.
func NickNameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldNickName, vs...))
}

// NickNameNotIn applies the NotIn predicate on the "nick_name" field.
func NickNameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldNickName, vs...))
}

// NickNameGT applies the GT predicate on the "nick_name" field.
func NickNameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldNickName, v))
}

// NickNameGTE applies the GTE predicate on the "nick_name" field.
func NickNameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldNickName, v))
}

// NickNameLT applies the LT predicate on the "nick_name" field.
func NickNameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldNickName, v))
}

// NickNameLTE applies the LTE predicate on the "nick_name" field.
func NickNameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldNickName, v))
}

// NickNameContains applies the Contains predicate on the "nick_name" field.
func NickNameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldNickName, v))
}

// NickNameHasPrefix applies the HasPrefix predicate on the "nick_name" field.
func NickNameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldNickName, v))
}

// NickNameHasSuffix applies the HasSuffix predicate on the "nick_name" field.
func NickNameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldNickName, v))
}

// NickNameEqualFold applies the EqualFold predicate on the "nick_name" field.
func NickNameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldNickName, v))
}

// NickNameContainsFold applies the ContainsFold predicate on the "nick_name" field.
func NickNameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldNickName, v))
}

// PhoneEQ applies the EQ predicate on the "phone" field.
func PhoneEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPhone, v))
}

// PhoneNEQ applies the NEQ predicate on the "phone" field.
func PhoneNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldPhone, v))
}

// PhoneIn applies the In predicate on the "phone" field.
func PhoneIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldPhone, vs...))
}

// PhoneNotIn applies the NotIn predicate on the "phone" field.
func PhoneNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldPhone, vs...))
}

// PhoneGT applies the GT predicate on the "phone" field.
func PhoneGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldPhone, v))
}

// PhoneGTE applies the GTE predicate on the "phone" field.
func PhoneGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldPhone, v))
}

// PhoneLT applies the LT predicate on the "phone" field.
func PhoneLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldPhone, v))
}

// PhoneLTE applies the LTE predicate on the "phone" field.
func PhoneLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldPhone, v))
}

// PhoneContains applies the Contains predicate on the "phone" field.
func PhoneContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldPhone, v))
}

// PhoneHasPrefix applies the HasPrefix predicate on the "phone" field.
func PhoneHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldPhone, v))
}

// PhoneHasSuffix applies the HasSuffix predicate on the "phone" field.
func PhoneHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldPhone, v))
}

// PhoneEqualFold applies the EqualFold predicate on the "phone" field.
func PhoneEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldPhone, v))
}

// PhoneContainsFold applies the ContainsFold predicate on the "phone" field.
func PhoneContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldPhone, v))
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPassword, v))
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldPassword, v))
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldPassword, vs...))
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldPassword, vs...))
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldPassword, v))
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldPassword, v))
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldPassword, v))
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldPassword, v))
}

// PasswordContains applies the Contains predicate on the "password" field.
func PasswordContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldPassword, v))
}

// PasswordHasPrefix applies the HasPrefix predicate on the "password" field.
func PasswordHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldPassword, v))
}

// PasswordHasSuffix applies the HasSuffix predicate on the "password" field.
func PasswordHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldPassword, v))
}

// PasswordEqualFold applies the EqualFold predicate on the "password" field.
func PasswordEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldPassword, v))
}

// PasswordContainsFold applies the ContainsFold predicate on the "password" field.
func PasswordContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldPassword, v))
}

// AvatarURLEQ applies the EQ predicate on the "avatar_url" field.
func AvatarURLEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldAvatarURL, v))
}

// AvatarURLNEQ applies the NEQ predicate on the "avatar_url" field.
func AvatarURLNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldAvatarURL, v))
}

// AvatarURLIn applies the In predicate on the "avatar_url" field.
func AvatarURLIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldAvatarURL, vs...))
}

// AvatarURLNotIn applies the NotIn predicate on the "avatar_url" field.
func AvatarURLNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldAvatarURL, vs...))
}

// AvatarURLGT applies the GT predicate on the "avatar_url" field.
func AvatarURLGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldAvatarURL, v))
}

// AvatarURLGTE applies the GTE predicate on the "avatar_url" field.
func AvatarURLGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldAvatarURL, v))
}

// AvatarURLLT applies the LT predicate on the "avatar_url" field.
func AvatarURLLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldAvatarURL, v))
}

// AvatarURLLTE applies the LTE predicate on the "avatar_url" field.
func AvatarURLLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldAvatarURL, v))
}

// AvatarURLContains applies the Contains predicate on the "avatar_url" field.
func AvatarURLContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldAvatarURL, v))
}

// AvatarURLHasPrefix applies the HasPrefix predicate on the "avatar_url" field.
func AvatarURLHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldAvatarURL, v))
}

// AvatarURLHasSuffix applies the HasSuffix predicate on the "avatar_url" field.
func AvatarURLHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldAvatarURL, v))
}

// AvatarURLEqualFold applies the EqualFold predicate on the "avatar_url" field.
func AvatarURLEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldAvatarURL, v))
}

// AvatarURLContainsFold applies the ContainsFold predicate on the "avatar_url" field.
func AvatarURLContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldAvatarURL, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.User {
	return predicate.User(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.User {
	return predicate.User(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldStatus, vs...))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v Type) predicate.User {
	return predicate.User(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v Type) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...Type) predicate.User {
	return predicate.User(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...Type) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldType, vs...))
}

// PlatformEQ applies the EQ predicate on the "platform" field.
func PlatformEQ(v int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPlatform, v))
}

// PlatformNEQ applies the NEQ predicate on the "platform" field.
func PlatformNEQ(v int) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldPlatform, v))
}

// PlatformIn applies the In predicate on the "platform" field.
func PlatformIn(vs ...int) predicate.User {
	return predicate.User(sql.FieldIn(FieldPlatform, vs...))
}

// PlatformNotIn applies the NotIn predicate on the "platform" field.
func PlatformNotIn(vs ...int) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldPlatform, vs...))
}

// PlatformGT applies the GT predicate on the "platform" field.
func PlatformGT(v int) predicate.User {
	return predicate.User(sql.FieldGT(FieldPlatform, v))
}

// PlatformGTE applies the GTE predicate on the "platform" field.
func PlatformGTE(v int) predicate.User {
	return predicate.User(sql.FieldGTE(FieldPlatform, v))
}

// PlatformLT applies the LT predicate on the "platform" field.
func PlatformLT(v int) predicate.User {
	return predicate.User(sql.FieldLT(FieldPlatform, v))
}

// PlatformLTE applies the LTE predicate on the "platform" field.
func PlatformLTE(v int) predicate.User {
	return predicate.User(sql.FieldLTE(FieldPlatform, v))
}

// HmacKeyEQ applies the EQ predicate on the "hmac_key" field.
func HmacKeyEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldHmacKey, v))
}

// HmacKeyNEQ applies the NEQ predicate on the "hmac_key" field.
func HmacKeyNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldHmacKey, v))
}

// HmacKeyIn applies the In predicate on the "hmac_key" field.
func HmacKeyIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldHmacKey, vs...))
}

// HmacKeyNotIn applies the NotIn predicate on the "hmac_key" field.
func HmacKeyNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldHmacKey, vs...))
}

// HmacKeyGT applies the GT predicate on the "hmac_key" field.
func HmacKeyGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldHmacKey, v))
}

// HmacKeyGTE applies the GTE predicate on the "hmac_key" field.
func HmacKeyGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldHmacKey, v))
}

// HmacKeyLT applies the LT predicate on the "hmac_key" field.
func HmacKeyLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldHmacKey, v))
}

// HmacKeyLTE applies the LTE predicate on the "hmac_key" field.
func HmacKeyLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldHmacKey, v))
}

// HmacKeyContains applies the Contains predicate on the "hmac_key" field.
func HmacKeyContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldHmacKey, v))
}

// HmacKeyHasPrefix applies the HasPrefix predicate on the "hmac_key" field.
func HmacKeyHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldHmacKey, v))
}

// HmacKeyHasSuffix applies the HasSuffix predicate on the "hmac_key" field.
func HmacKeyHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldHmacKey, v))
}

// HmacKeyEqualFold applies the EqualFold predicate on the "hmac_key" field.
func HmacKeyEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldHmacKey, v))
}

// HmacKeyContainsFold applies the ContainsFold predicate on the "hmac_key" field.
func HmacKeyContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldHmacKey, v))
}

// HmacSecretEQ applies the EQ predicate on the "hmac_secret" field.
func HmacSecretEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldHmacSecret, v))
}

// HmacSecretNEQ applies the NEQ predicate on the "hmac_secret" field.
func HmacSecretNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldHmacSecret, v))
}

// HmacSecretIn applies the In predicate on the "hmac_secret" field.
func HmacSecretIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldHmacSecret, vs...))
}

// HmacSecretNotIn applies the NotIn predicate on the "hmac_secret" field.
func HmacSecretNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldHmacSecret, vs...))
}

// HmacSecretGT applies the GT predicate on the "hmac_secret" field.
func HmacSecretGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldHmacSecret, v))
}

// HmacSecretGTE applies the GTE predicate on the "hmac_secret" field.
func HmacSecretGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldHmacSecret, v))
}

// HmacSecretLT applies the LT predicate on the "hmac_secret" field.
func HmacSecretLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldHmacSecret, v))
}

// HmacSecretLTE applies the LTE predicate on the "hmac_secret" field.
func HmacSecretLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldHmacSecret, v))
}

// HmacSecretContains applies the Contains predicate on the "hmac_secret" field.
func HmacSecretContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldHmacSecret, v))
}

// HmacSecretHasPrefix applies the HasPrefix predicate on the "hmac_secret" field.
func HmacSecretHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldHmacSecret, v))
}

// HmacSecretHasSuffix applies the HasSuffix predicate on the "hmac_secret" field.
func HmacSecretHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldHmacSecret, v))
}

// HmacSecretEqualFold applies the EqualFold predicate on the "hmac_secret" field.
func HmacSecretEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldHmacSecret, v))
}

// HmacSecretContainsFold applies the ContainsFold predicate on the "hmac_secret" field.
func HmacSecretContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldHmacSecret, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		p(s.Not())
	})
}
