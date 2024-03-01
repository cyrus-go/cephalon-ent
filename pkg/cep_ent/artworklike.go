// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/artwork"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/artworklike"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/user"
)

// 用户投赞成票给艺术作品；ArtworkLike
type ArtworkLike struct {
	config `json:"-"`
	// ID of the ent.
	// 19 位雪花 ID
	ID int64 `json:"id,string"`
	// 创建者 ID
	CreatedBy int64 `json:"created_by"`
	// 更新者 ID
	UpdatedBy int64 `json:"updated_by"`
	// 创建时刻，带时区
	CreatedAt time.Time `json:"created_at"`
	// 更新时刻，带时区
	UpdatedAt time.Time `json:"updated_at"`
	// 软删除时刻，带时区
	DeletedAt time.Time `json:"deleted_at"`
	// 外键，用户 id
	UserID int64 `json:"user_id"`
	// 外键，作品 id
	ArtworkID int64 `json:"artwork_id"`
	// 投票日期
	Date int32 `json:"date"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ArtworkLikeQuery when eager-loading is set.
	Edges        ArtworkLikeEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ArtworkLikeEdges holds the relations/edges for other nodes in the graph.
type ArtworkLikeEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Artwork holds the value of the artwork edge.
	Artwork *Artwork `json:"artwork,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ArtworkLikeEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// ArtworkOrErr returns the Artwork value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ArtworkLikeEdges) ArtworkOrErr() (*Artwork, error) {
	if e.loadedTypes[1] {
		if e.Artwork == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: artwork.Label}
		}
		return e.Artwork, nil
	}
	return nil, &NotLoadedError{edge: "artwork"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ArtworkLike) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case artworklike.FieldID, artworklike.FieldCreatedBy, artworklike.FieldUpdatedBy, artworklike.FieldUserID, artworklike.FieldArtworkID, artworklike.FieldDate:
			values[i] = new(sql.NullInt64)
		case artworklike.FieldCreatedAt, artworklike.FieldUpdatedAt, artworklike.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ArtworkLike fields.
func (al *ArtworkLike) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case artworklike.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			al.ID = int64(value.Int64)
		case artworklike.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				al.CreatedBy = value.Int64
			}
		case artworklike.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				al.UpdatedBy = value.Int64
			}
		case artworklike.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				al.CreatedAt = value.Time
			}
		case artworklike.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				al.UpdatedAt = value.Time
			}
		case artworklike.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				al.DeletedAt = value.Time
			}
		case artworklike.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				al.UserID = value.Int64
			}
		case artworklike.FieldArtworkID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field artwork_id", values[i])
			} else if value.Valid {
				al.ArtworkID = value.Int64
			}
		case artworklike.FieldDate:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field date", values[i])
			} else if value.Valid {
				al.Date = int32(value.Int64)
			}
		default:
			al.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ArtworkLike.
// This includes values selected through modifiers, order, etc.
func (al *ArtworkLike) Value(name string) (ent.Value, error) {
	return al.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the ArtworkLike entity.
func (al *ArtworkLike) QueryUser() *UserQuery {
	return NewArtworkLikeClient(al.config).QueryUser(al)
}

// QueryArtwork queries the "artwork" edge of the ArtworkLike entity.
func (al *ArtworkLike) QueryArtwork() *ArtworkQuery {
	return NewArtworkLikeClient(al.config).QueryArtwork(al)
}

// Update returns a builder for updating this ArtworkLike.
// Note that you need to call ArtworkLike.Unwrap() before calling this method if this ArtworkLike
// was returned from a transaction, and the transaction was committed or rolled back.
func (al *ArtworkLike) Update() *ArtworkLikeUpdateOne {
	return NewArtworkLikeClient(al.config).UpdateOne(al)
}

// Unwrap unwraps the ArtworkLike entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (al *ArtworkLike) Unwrap() *ArtworkLike {
	_tx, ok := al.config.driver.(*txDriver)
	if !ok {
		panic("cep_ent: ArtworkLike is not a transactional entity")
	}
	al.config.driver = _tx.drv
	return al
}

// String implements the fmt.Stringer.
func (al *ArtworkLike) String() string {
	var builder strings.Builder
	builder.WriteString("ArtworkLike(")
	builder.WriteString(fmt.Sprintf("id=%v, ", al.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", al.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", al.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(al.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(al.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(al.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", al.UserID))
	builder.WriteString(", ")
	builder.WriteString("artwork_id=")
	builder.WriteString(fmt.Sprintf("%v", al.ArtworkID))
	builder.WriteString(", ")
	builder.WriteString("date=")
	builder.WriteString(fmt.Sprintf("%v", al.Date))
	builder.WriteByte(')')
	return builder.String()
}

// ArtworkLikes is a parsable slice of ArtworkLike.
type ArtworkLikes []*ArtworkLike