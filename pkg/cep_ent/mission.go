// Code generated by ent, DO NOT EDIT.

package cep_ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/hmackeypair"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/mission"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/missionbatch"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/missionconsumeorder"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/missionfailedfeedback"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/missionkind"
	"github.com/stark-sim/cephalon-ent/pkg/cep_ent/user"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

// 任务，具备任务要求，记录完成情况和结果，金额相关信息在 mission_orders 等订单侧
type Mission struct {
	config `json:"-"`
	// ID of the ent.
	// 19 位雪花 ID
	ID int64 `json:"id,string"`
	// 创建者 ID
	CreatedBy int64 `json:"created_by,string"`
	// 更新者 ID
	UpdatedBy int64 `json:"updated_by,string"`
	// 创建时刻，带时区
	CreatedAt time.Time `json:"created_at"`
	// 更新时刻，带时区
	UpdatedAt time.Time `json:"updated_at"`
	// 软删除时刻，带时区
	DeletedAt time.Time `json:"deleted_at"`
	// 任务类型
	Type enums.MissionType `json:"type"`
	// 外键，任务种类 id
	MissionKindID int64 `json:"mission_kind_id,string"`
	// 任务的请求参数体
	Body string `json:"body"`
	// 回调地址，空字符串表示不回调
	CallBackURL string `json:"call_back_url"`
	// 回调时带上的参数，接收任何类型数据后 json 压缩
	CallBackInfo *string `json:"call_back_info"`
	// 回调时带上的参数，需 json 反序列化后才可使用，所以没有直接 json 序列化字段 (sensitive)
	CallBackData []byte `json:"-"`
	// 任务状态
	Status enums.MissionStatus `json:"status"`
	// 任务结果，pending 表示还没有结果
	Result enums.MissionResult `json:"result"`
	// 新任务状态，整合原状态和结果
	State enums.MissionState `json:"state"`
	// 任务结果资源位置列表序列化
	ResultUrls []string `json:"-"`
	// 任务结果链接列表，json 序列化后存储
	Urls string `json:"urls"`
	// 任务创建者的密钥对 ID
	KeyPairID int64 `json:"key_pair_id"`
	// 外键任务的创建者 id
	UserID int64 `json:"user_id,string"`
	// 外键关联任务批次
	MissionBatchID int64 `json:"mission_batch_id,string"`
	// 任务批次号
	MissionBatchNumber string `json:"mission_batch_number"`
	// 最低可接显卡
	GpuVersion enums.GpuVersion `json:"gpu_version"`
	// 任务单价，按次(count)就是 unit_cep/次，按时(time)就是 unit_cep/分钟
	UnitCep int64 `json:"unit_cep"`
	// 内部功能返回码
	RespStatusCode int32 `json:"resp_status_code"`
	// 返回内容体 json 转 string
	RespBody string `json:"resp_body"`
	// 当 type 为 sd_api 时使用，为转发的 sd 内部接口相对路径
	InnerURI string `json:"inner_uri"`
	// 内部转发接口的请求方式，POST 或者 GET 等
	InnerMethod enums.InnerMethod `json:"inner_method"`
	// 当 type 为 key_pair 时，使用的临时密钥对的键
	TempHmacKey string `json:"temp_hmac_key"`
	// 当 type 为 key_pair 时，使用的临时密钥对的值
	TempHmacSecret string `json:"temp_hmac_secret"`
	// 创建任务时使用了的 二级 hmac_key
	SecondHmacKey string `json:"second_hmac_key"`
	// 某些任务会使用到的验证用户名
	Username string `json:"username"`
	// 某些任务会使用到的验证密码
	Password string `json:"password"`
	// 任务的设备白名单
	WhiteDeviceIds []string `json:"white_device_ids"`
	// 任务的设备黑名单
	BlackDeviceIds []string `json:"black_device_ids"`
	// 任务开始时间
	StartedAt *time.Time `json:"started_at"`
	// 任务结束时间
	FinishedAt *time.Time `json:"finished_at"`
	// 任务到期时间（包时任务才有）
	ExpiredAt *time.Time `json:"expired_at"`
	// 任务释放时刻
	FreeAt time.Time `json:"free_at"`
	// 任务创建方式，user：用户自己通过官网页面创建，load_balance：用户通过创建 Loadbalance 自动创建
	CreateWay enums.CreateWay `json:"create_way"`
	// 任务关闭方式，user：用户自己关闭，balance_not_enough：余额不足自动关闭
	CloseWay enums.CloseWay `json:"close_way"`
	// 用戶关闭任务时间
	ClosedAt *time.Time `json:"closed_at"`
	// 预警次数，任务运行时间超过一定时间会发送预警消息
	WarningTimes int64 `json:"warning_times"`
	// 备注信息
	Remark string `json:"remark"`
	// 是否需要使用鉴权（应用开启后需要账号密码登陆验证）
	UseAuth bool `json:"use_auth"`
	// 外键，重新开机的旧应用 ID
	OldMissionID int64 `json:"old_mission_id,string"`
	// 任务定时关机时间
	TimedShutdown *time.Time `json:"timed_shutdown"`
	// 多GPU支持,使用GPU数量
	GpuNum int `json:"gpu_num"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MissionQuery when eager-loading is set.
	Edges                  MissionEdges `json:"edges"`
	extra_service_missions *int64
	selectValues           sql.SelectValues
}

// MissionEdges holds the relations/edges for other nodes in the graph.
type MissionEdges struct {
	// MissionKind holds the value of the mission_kind edge.
	MissionKind *MissionKind `json:"mission_kind,omitempty"`
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// KeyPair holds the value of the key_pair edge.
	KeyPair *HmacKeyPair `json:"key_pair,omitempty"`
	// MissionBatch holds the value of the mission_batch edge.
	MissionBatch *MissionBatch `json:"mission_batch,omitempty"`
	// OldMission holds the value of the old_mission edge.
	OldMission *Mission `json:"old_mission,omitempty"`
	// MissionKeyPairs holds the value of the mission_key_pairs edge.
	MissionKeyPairs []*MissionKeyPair `json:"mission_key_pairs,omitempty"`
	// MissionConsumeOrder holds the value of the mission_consume_order edge.
	MissionConsumeOrder *MissionConsumeOrder `json:"mission_consume_order,omitempty"`
	// MissionProduceOrders holds the value of the mission_produce_orders edge.
	MissionProduceOrders []*MissionProduceOrder `json:"mission_produce_orders,omitempty"`
	// MissionProductions holds the value of the mission_productions edge.
	MissionProductions []*MissionProduction `json:"mission_productions,omitempty"`
	// MissionOrders holds the value of the mission_orders edge.
	MissionOrders []*MissionOrder `json:"mission_orders,omitempty"`
	// RenewalAgreements holds the value of the renewal_agreements edge.
	RenewalAgreements []*RenewalAgreement `json:"renewal_agreements,omitempty"`
	// MissionExtraServices holds the value of the mission_extra_services edge.
	MissionExtraServices []*MissionExtraService `json:"mission_extra_services,omitempty"`
	// ExtraServices holds the value of the extra_services edge.
	ExtraServices []*ExtraService `json:"extra_services,omitempty"`
	// ExtraServiceOrders holds the value of the extra_service_orders edge.
	ExtraServiceOrders []*ExtraServiceOrder `json:"extra_service_orders,omitempty"`
	// RebootMissions holds the value of the reboot_missions edge.
	RebootMissions []*Mission `json:"reboot_missions,omitempty"`
	// MissionFailedFeedback holds the value of the mission_failed_feedback edge.
	MissionFailedFeedback *MissionFailedFeedback `json:"mission_failed_feedback,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [16]bool
}

// MissionKindOrErr returns the MissionKind value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MissionEdges) MissionKindOrErr() (*MissionKind, error) {
	if e.loadedTypes[0] {
		if e.MissionKind == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: missionkind.Label}
		}
		return e.MissionKind, nil
	}
	return nil, &NotLoadedError{edge: "mission_kind"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MissionEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// KeyPairOrErr returns the KeyPair value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MissionEdges) KeyPairOrErr() (*HmacKeyPair, error) {
	if e.loadedTypes[2] {
		if e.KeyPair == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: hmackeypair.Label}
		}
		return e.KeyPair, nil
	}
	return nil, &NotLoadedError{edge: "key_pair"}
}

// MissionBatchOrErr returns the MissionBatch value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MissionEdges) MissionBatchOrErr() (*MissionBatch, error) {
	if e.loadedTypes[3] {
		if e.MissionBatch == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: missionbatch.Label}
		}
		return e.MissionBatch, nil
	}
	return nil, &NotLoadedError{edge: "mission_batch"}
}

// OldMissionOrErr returns the OldMission value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MissionEdges) OldMissionOrErr() (*Mission, error) {
	if e.loadedTypes[4] {
		if e.OldMission == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: mission.Label}
		}
		return e.OldMission, nil
	}
	return nil, &NotLoadedError{edge: "old_mission"}
}

// MissionKeyPairsOrErr returns the MissionKeyPairs value or an error if the edge
// was not loaded in eager-loading.
func (e MissionEdges) MissionKeyPairsOrErr() ([]*MissionKeyPair, error) {
	if e.loadedTypes[5] {
		return e.MissionKeyPairs, nil
	}
	return nil, &NotLoadedError{edge: "mission_key_pairs"}
}

// MissionConsumeOrderOrErr returns the MissionConsumeOrder value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MissionEdges) MissionConsumeOrderOrErr() (*MissionConsumeOrder, error) {
	if e.loadedTypes[6] {
		if e.MissionConsumeOrder == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: missionconsumeorder.Label}
		}
		return e.MissionConsumeOrder, nil
	}
	return nil, &NotLoadedError{edge: "mission_consume_order"}
}

// MissionProduceOrdersOrErr returns the MissionProduceOrders value or an error if the edge
// was not loaded in eager-loading.
func (e MissionEdges) MissionProduceOrdersOrErr() ([]*MissionProduceOrder, error) {
	if e.loadedTypes[7] {
		return e.MissionProduceOrders, nil
	}
	return nil, &NotLoadedError{edge: "mission_produce_orders"}
}

// MissionProductionsOrErr returns the MissionProductions value or an error if the edge
// was not loaded in eager-loading.
func (e MissionEdges) MissionProductionsOrErr() ([]*MissionProduction, error) {
	if e.loadedTypes[8] {
		return e.MissionProductions, nil
	}
	return nil, &NotLoadedError{edge: "mission_productions"}
}

// MissionOrdersOrErr returns the MissionOrders value or an error if the edge
// was not loaded in eager-loading.
func (e MissionEdges) MissionOrdersOrErr() ([]*MissionOrder, error) {
	if e.loadedTypes[9] {
		return e.MissionOrders, nil
	}
	return nil, &NotLoadedError{edge: "mission_orders"}
}

// RenewalAgreementsOrErr returns the RenewalAgreements value or an error if the edge
// was not loaded in eager-loading.
func (e MissionEdges) RenewalAgreementsOrErr() ([]*RenewalAgreement, error) {
	if e.loadedTypes[10] {
		return e.RenewalAgreements, nil
	}
	return nil, &NotLoadedError{edge: "renewal_agreements"}
}

// MissionExtraServicesOrErr returns the MissionExtraServices value or an error if the edge
// was not loaded in eager-loading.
func (e MissionEdges) MissionExtraServicesOrErr() ([]*MissionExtraService, error) {
	if e.loadedTypes[11] {
		return e.MissionExtraServices, nil
	}
	return nil, &NotLoadedError{edge: "mission_extra_services"}
}

// ExtraServicesOrErr returns the ExtraServices value or an error if the edge
// was not loaded in eager-loading.
func (e MissionEdges) ExtraServicesOrErr() ([]*ExtraService, error) {
	if e.loadedTypes[12] {
		return e.ExtraServices, nil
	}
	return nil, &NotLoadedError{edge: "extra_services"}
}

// ExtraServiceOrdersOrErr returns the ExtraServiceOrders value or an error if the edge
// was not loaded in eager-loading.
func (e MissionEdges) ExtraServiceOrdersOrErr() ([]*ExtraServiceOrder, error) {
	if e.loadedTypes[13] {
		return e.ExtraServiceOrders, nil
	}
	return nil, &NotLoadedError{edge: "extra_service_orders"}
}

// RebootMissionsOrErr returns the RebootMissions value or an error if the edge
// was not loaded in eager-loading.
func (e MissionEdges) RebootMissionsOrErr() ([]*Mission, error) {
	if e.loadedTypes[14] {
		return e.RebootMissions, nil
	}
	return nil, &NotLoadedError{edge: "reboot_missions"}
}

// MissionFailedFeedbackOrErr returns the MissionFailedFeedback value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MissionEdges) MissionFailedFeedbackOrErr() (*MissionFailedFeedback, error) {
	if e.loadedTypes[15] {
		if e.MissionFailedFeedback == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: missionfailedfeedback.Label}
		}
		return e.MissionFailedFeedback, nil
	}
	return nil, &NotLoadedError{edge: "mission_failed_feedback"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Mission) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case mission.FieldCallBackData, mission.FieldResultUrls:
			values[i] = new([]byte)
		case mission.FieldUseAuth:
			values[i] = new(sql.NullBool)
		case mission.FieldID, mission.FieldCreatedBy, mission.FieldUpdatedBy, mission.FieldMissionKindID, mission.FieldKeyPairID, mission.FieldUserID, mission.FieldMissionBatchID, mission.FieldUnitCep, mission.FieldRespStatusCode, mission.FieldWarningTimes, mission.FieldOldMissionID, mission.FieldGpuNum:
			values[i] = new(sql.NullInt64)
		case mission.FieldType, mission.FieldBody, mission.FieldCallBackURL, mission.FieldCallBackInfo, mission.FieldStatus, mission.FieldResult, mission.FieldState, mission.FieldUrls, mission.FieldMissionBatchNumber, mission.FieldGpuVersion, mission.FieldRespBody, mission.FieldInnerURI, mission.FieldInnerMethod, mission.FieldTempHmacKey, mission.FieldTempHmacSecret, mission.FieldSecondHmacKey, mission.FieldUsername, mission.FieldPassword, mission.FieldCreateWay, mission.FieldCloseWay, mission.FieldRemark:
			values[i] = new(sql.NullString)
		case mission.FieldCreatedAt, mission.FieldUpdatedAt, mission.FieldDeletedAt, mission.FieldStartedAt, mission.FieldFinishedAt, mission.FieldExpiredAt, mission.FieldFreeAt, mission.FieldClosedAt, mission.FieldTimedShutdown:
			values[i] = new(sql.NullTime)
		case mission.FieldWhiteDeviceIds:
			values[i] = mission.ValueScanner.WhiteDeviceIds.ScanValue()
		case mission.FieldBlackDeviceIds:
			values[i] = mission.ValueScanner.BlackDeviceIds.ScanValue()
		case mission.ForeignKeys[0]: // extra_service_missions
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Mission fields.
func (m *Mission) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case mission.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = int64(value.Int64)
		case mission.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				m.CreatedBy = value.Int64
			}
		case mission.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				m.UpdatedBy = value.Int64
			}
		case mission.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				m.CreatedAt = value.Time
			}
		case mission.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				m.UpdatedAt = value.Time
			}
		case mission.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				m.DeletedAt = value.Time
			}
		case mission.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				m.Type = enums.MissionType(value.String)
			}
		case mission.FieldMissionKindID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field mission_kind_id", values[i])
			} else if value.Valid {
				m.MissionKindID = value.Int64
			}
		case mission.FieldBody:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field body", values[i])
			} else if value.Valid {
				m.Body = value.String
			}
		case mission.FieldCallBackURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field call_back_url", values[i])
			} else if value.Valid {
				m.CallBackURL = value.String
			}
		case mission.FieldCallBackInfo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field call_back_info", values[i])
			} else if value.Valid {
				m.CallBackInfo = new(string)
				*m.CallBackInfo = value.String
			}
		case mission.FieldCallBackData:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field call_back_data", values[i])
			} else if value != nil {
				m.CallBackData = *value
			}
		case mission.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				m.Status = enums.MissionStatus(value.String)
			}
		case mission.FieldResult:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field result", values[i])
			} else if value.Valid {
				m.Result = enums.MissionResult(value.String)
			}
		case mission.FieldState:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field state", values[i])
			} else if value.Valid {
				m.State = enums.MissionState(value.String)
			}
		case mission.FieldResultUrls:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field result_urls", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &m.ResultUrls); err != nil {
					return fmt.Errorf("unmarshal field result_urls: %w", err)
				}
			}
		case mission.FieldUrls:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field urls", values[i])
			} else if value.Valid {
				m.Urls = value.String
			}
		case mission.FieldKeyPairID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field key_pair_id", values[i])
			} else if value.Valid {
				m.KeyPairID = value.Int64
			}
		case mission.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				m.UserID = value.Int64
			}
		case mission.FieldMissionBatchID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field mission_batch_id", values[i])
			} else if value.Valid {
				m.MissionBatchID = value.Int64
			}
		case mission.FieldMissionBatchNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mission_batch_number", values[i])
			} else if value.Valid {
				m.MissionBatchNumber = value.String
			}
		case mission.FieldGpuVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field gpu_version", values[i])
			} else if value.Valid {
				m.GpuVersion = enums.GpuVersion(value.String)
			}
		case mission.FieldUnitCep:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field unit_cep", values[i])
			} else if value.Valid {
				m.UnitCep = value.Int64
			}
		case mission.FieldRespStatusCode:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field resp_status_code", values[i])
			} else if value.Valid {
				m.RespStatusCode = int32(value.Int64)
			}
		case mission.FieldRespBody:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field resp_body", values[i])
			} else if value.Valid {
				m.RespBody = value.String
			}
		case mission.FieldInnerURI:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field inner_uri", values[i])
			} else if value.Valid {
				m.InnerURI = value.String
			}
		case mission.FieldInnerMethod:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field inner_method", values[i])
			} else if value.Valid {
				m.InnerMethod = enums.InnerMethod(value.String)
			}
		case mission.FieldTempHmacKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field temp_hmac_key", values[i])
			} else if value.Valid {
				m.TempHmacKey = value.String
			}
		case mission.FieldTempHmacSecret:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field temp_hmac_secret", values[i])
			} else if value.Valid {
				m.TempHmacSecret = value.String
			}
		case mission.FieldSecondHmacKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field second_hmac_key", values[i])
			} else if value.Valid {
				m.SecondHmacKey = value.String
			}
		case mission.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				m.Username = value.String
			}
		case mission.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				m.Password = value.String
			}
		case mission.FieldWhiteDeviceIds:
			if value, err := mission.ValueScanner.WhiteDeviceIds.FromValue(values[i]); err != nil {
				return err
			} else {
				m.WhiteDeviceIds = value
			}
		case mission.FieldBlackDeviceIds:
			if value, err := mission.ValueScanner.BlackDeviceIds.FromValue(values[i]); err != nil {
				return err
			} else {
				m.BlackDeviceIds = value
			}
		case mission.FieldStartedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field started_at", values[i])
			} else if value.Valid {
				m.StartedAt = new(time.Time)
				*m.StartedAt = value.Time
			}
		case mission.FieldFinishedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field finished_at", values[i])
			} else if value.Valid {
				m.FinishedAt = new(time.Time)
				*m.FinishedAt = value.Time
			}
		case mission.FieldExpiredAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field expired_at", values[i])
			} else if value.Valid {
				m.ExpiredAt = new(time.Time)
				*m.ExpiredAt = value.Time
			}
		case mission.FieldFreeAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field free_at", values[i])
			} else if value.Valid {
				m.FreeAt = value.Time
			}
		case mission.FieldCreateWay:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field create_way", values[i])
			} else if value.Valid {
				m.CreateWay = enums.CreateWay(value.String)
			}
		case mission.FieldCloseWay:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field close_way", values[i])
			} else if value.Valid {
				m.CloseWay = enums.CloseWay(value.String)
			}
		case mission.FieldClosedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field closed_at", values[i])
			} else if value.Valid {
				m.ClosedAt = new(time.Time)
				*m.ClosedAt = value.Time
			}
		case mission.FieldWarningTimes:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field warning_times", values[i])
			} else if value.Valid {
				m.WarningTimes = value.Int64
			}
		case mission.FieldRemark:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remark", values[i])
			} else if value.Valid {
				m.Remark = value.String
			}
		case mission.FieldUseAuth:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field use_auth", values[i])
			} else if value.Valid {
				m.UseAuth = value.Bool
			}
		case mission.FieldOldMissionID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field old_mission_id", values[i])
			} else if value.Valid {
				m.OldMissionID = value.Int64
			}
		case mission.FieldTimedShutdown:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field timed_shutdown", values[i])
			} else if value.Valid {
				m.TimedShutdown = new(time.Time)
				*m.TimedShutdown = value.Time
			}
		case mission.FieldGpuNum:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field gpu_num", values[i])
			} else if value.Valid {
				m.GpuNum = int(value.Int64)
			}
		case mission.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field extra_service_missions", value)
			} else if value.Valid {
				m.extra_service_missions = new(int64)
				*m.extra_service_missions = int64(value.Int64)
			}
		default:
			m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Mission.
// This includes values selected through modifiers, order, etc.
func (m *Mission) Value(name string) (ent.Value, error) {
	return m.selectValues.Get(name)
}

// QueryMissionKind queries the "mission_kind" edge of the Mission entity.
func (m *Mission) QueryMissionKind() *MissionKindQuery {
	return NewMissionClient(m.config).QueryMissionKind(m)
}

// QueryUser queries the "user" edge of the Mission entity.
func (m *Mission) QueryUser() *UserQuery {
	return NewMissionClient(m.config).QueryUser(m)
}

// QueryKeyPair queries the "key_pair" edge of the Mission entity.
func (m *Mission) QueryKeyPair() *HmacKeyPairQuery {
	return NewMissionClient(m.config).QueryKeyPair(m)
}

// QueryMissionBatch queries the "mission_batch" edge of the Mission entity.
func (m *Mission) QueryMissionBatch() *MissionBatchQuery {
	return NewMissionClient(m.config).QueryMissionBatch(m)
}

// QueryOldMission queries the "old_mission" edge of the Mission entity.
func (m *Mission) QueryOldMission() *MissionQuery {
	return NewMissionClient(m.config).QueryOldMission(m)
}

// QueryMissionKeyPairs queries the "mission_key_pairs" edge of the Mission entity.
func (m *Mission) QueryMissionKeyPairs() *MissionKeyPairQuery {
	return NewMissionClient(m.config).QueryMissionKeyPairs(m)
}

// QueryMissionConsumeOrder queries the "mission_consume_order" edge of the Mission entity.
func (m *Mission) QueryMissionConsumeOrder() *MissionConsumeOrderQuery {
	return NewMissionClient(m.config).QueryMissionConsumeOrder(m)
}

// QueryMissionProduceOrders queries the "mission_produce_orders" edge of the Mission entity.
func (m *Mission) QueryMissionProduceOrders() *MissionProduceOrderQuery {
	return NewMissionClient(m.config).QueryMissionProduceOrders(m)
}

// QueryMissionProductions queries the "mission_productions" edge of the Mission entity.
func (m *Mission) QueryMissionProductions() *MissionProductionQuery {
	return NewMissionClient(m.config).QueryMissionProductions(m)
}

// QueryMissionOrders queries the "mission_orders" edge of the Mission entity.
func (m *Mission) QueryMissionOrders() *MissionOrderQuery {
	return NewMissionClient(m.config).QueryMissionOrders(m)
}

// QueryRenewalAgreements queries the "renewal_agreements" edge of the Mission entity.
func (m *Mission) QueryRenewalAgreements() *RenewalAgreementQuery {
	return NewMissionClient(m.config).QueryRenewalAgreements(m)
}

// QueryMissionExtraServices queries the "mission_extra_services" edge of the Mission entity.
func (m *Mission) QueryMissionExtraServices() *MissionExtraServiceQuery {
	return NewMissionClient(m.config).QueryMissionExtraServices(m)
}

// QueryExtraServices queries the "extra_services" edge of the Mission entity.
func (m *Mission) QueryExtraServices() *ExtraServiceQuery {
	return NewMissionClient(m.config).QueryExtraServices(m)
}

// QueryExtraServiceOrders queries the "extra_service_orders" edge of the Mission entity.
func (m *Mission) QueryExtraServiceOrders() *ExtraServiceOrderQuery {
	return NewMissionClient(m.config).QueryExtraServiceOrders(m)
}

// QueryRebootMissions queries the "reboot_missions" edge of the Mission entity.
func (m *Mission) QueryRebootMissions() *MissionQuery {
	return NewMissionClient(m.config).QueryRebootMissions(m)
}

// QueryMissionFailedFeedback queries the "mission_failed_feedback" edge of the Mission entity.
func (m *Mission) QueryMissionFailedFeedback() *MissionFailedFeedbackQuery {
	return NewMissionClient(m.config).QueryMissionFailedFeedback(m)
}

// Update returns a builder for updating this Mission.
// Note that you need to call Mission.Unwrap() before calling this method if this Mission
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Mission) Update() *MissionUpdateOne {
	return NewMissionClient(m.config).UpdateOne(m)
}

// Unwrap unwraps the Mission entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Mission) Unwrap() *Mission {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("cep_ent: Mission is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Mission) String() string {
	var builder strings.Builder
	builder.WriteString("Mission(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", m.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", m.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(m.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(m.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(m.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", m.Type))
	builder.WriteString(", ")
	builder.WriteString("mission_kind_id=")
	builder.WriteString(fmt.Sprintf("%v", m.MissionKindID))
	builder.WriteString(", ")
	builder.WriteString("body=")
	builder.WriteString(m.Body)
	builder.WriteString(", ")
	builder.WriteString("call_back_url=")
	builder.WriteString(m.CallBackURL)
	builder.WriteString(", ")
	if v := m.CallBackInfo; v != nil {
		builder.WriteString("call_back_info=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("call_back_data=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", m.Status))
	builder.WriteString(", ")
	builder.WriteString("result=")
	builder.WriteString(fmt.Sprintf("%v", m.Result))
	builder.WriteString(", ")
	builder.WriteString("state=")
	builder.WriteString(fmt.Sprintf("%v", m.State))
	builder.WriteString(", ")
	builder.WriteString("result_urls=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("urls=")
	builder.WriteString(m.Urls)
	builder.WriteString(", ")
	builder.WriteString("key_pair_id=")
	builder.WriteString(fmt.Sprintf("%v", m.KeyPairID))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", m.UserID))
	builder.WriteString(", ")
	builder.WriteString("mission_batch_id=")
	builder.WriteString(fmt.Sprintf("%v", m.MissionBatchID))
	builder.WriteString(", ")
	builder.WriteString("mission_batch_number=")
	builder.WriteString(m.MissionBatchNumber)
	builder.WriteString(", ")
	builder.WriteString("gpu_version=")
	builder.WriteString(fmt.Sprintf("%v", m.GpuVersion))
	builder.WriteString(", ")
	builder.WriteString("unit_cep=")
	builder.WriteString(fmt.Sprintf("%v", m.UnitCep))
	builder.WriteString(", ")
	builder.WriteString("resp_status_code=")
	builder.WriteString(fmt.Sprintf("%v", m.RespStatusCode))
	builder.WriteString(", ")
	builder.WriteString("resp_body=")
	builder.WriteString(m.RespBody)
	builder.WriteString(", ")
	builder.WriteString("inner_uri=")
	builder.WriteString(m.InnerURI)
	builder.WriteString(", ")
	builder.WriteString("inner_method=")
	builder.WriteString(fmt.Sprintf("%v", m.InnerMethod))
	builder.WriteString(", ")
	builder.WriteString("temp_hmac_key=")
	builder.WriteString(m.TempHmacKey)
	builder.WriteString(", ")
	builder.WriteString("temp_hmac_secret=")
	builder.WriteString(m.TempHmacSecret)
	builder.WriteString(", ")
	builder.WriteString("second_hmac_key=")
	builder.WriteString(m.SecondHmacKey)
	builder.WriteString(", ")
	builder.WriteString("username=")
	builder.WriteString(m.Username)
	builder.WriteString(", ")
	builder.WriteString("password=")
	builder.WriteString(m.Password)
	builder.WriteString(", ")
	builder.WriteString("white_device_ids=")
	builder.WriteString(fmt.Sprintf("%v", m.WhiteDeviceIds))
	builder.WriteString(", ")
	builder.WriteString("black_device_ids=")
	builder.WriteString(fmt.Sprintf("%v", m.BlackDeviceIds))
	builder.WriteString(", ")
	if v := m.StartedAt; v != nil {
		builder.WriteString("started_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := m.FinishedAt; v != nil {
		builder.WriteString("finished_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := m.ExpiredAt; v != nil {
		builder.WriteString("expired_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("free_at=")
	builder.WriteString(m.FreeAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("create_way=")
	builder.WriteString(fmt.Sprintf("%v", m.CreateWay))
	builder.WriteString(", ")
	builder.WriteString("close_way=")
	builder.WriteString(fmt.Sprintf("%v", m.CloseWay))
	builder.WriteString(", ")
	if v := m.ClosedAt; v != nil {
		builder.WriteString("closed_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("warning_times=")
	builder.WriteString(fmt.Sprintf("%v", m.WarningTimes))
	builder.WriteString(", ")
	builder.WriteString("remark=")
	builder.WriteString(m.Remark)
	builder.WriteString(", ")
	builder.WriteString("use_auth=")
	builder.WriteString(fmt.Sprintf("%v", m.UseAuth))
	builder.WriteString(", ")
	builder.WriteString("old_mission_id=")
	builder.WriteString(fmt.Sprintf("%v", m.OldMissionID))
	builder.WriteString(", ")
	if v := m.TimedShutdown; v != nil {
		builder.WriteString("timed_shutdown=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("gpu_num=")
	builder.WriteString(fmt.Sprintf("%v", m.GpuNum))
	builder.WriteByte(')')
	return builder.String()
}

// Missions is a parsable slice of Mission.
type Missions []*Mission
