// Code generated by ent, DO NOT EDIT.

package price

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
)

const (
	// Label holds the string label denoting the price type in the database.
	Label = "price"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldUpdatedBy holds the string denoting the updated_by field in the database.
	FieldUpdatedBy = "updated_by"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldGpuID holds the string denoting the gpu_id field in the database.
	FieldGpuID = "gpu_id"
	// FieldGpuVersion holds the string denoting the gpu_version field in the database.
	FieldGpuVersion = "gpu_version"
	// FieldMissionType holds the string denoting the mission_type field in the database.
	FieldMissionType = "mission_type"
	// FieldMissionCategory holds the string denoting the mission_category field in the database.
	FieldMissionCategory = "mission_category"
	// FieldMissionBillingType holds the string denoting the mission_billing_type field in the database.
	FieldMissionBillingType = "mission_billing_type"
	// FieldCep holds the string denoting the cep field in the database.
	FieldCep = "cep"
	// FieldOriginalCep holds the string denoting the original_cep field in the database.
	FieldOriginalCep = "original_cep"
	// FieldStartedAt holds the string denoting the started_at field in the database.
	FieldStartedAt = "started_at"
	// FieldFinishedAt holds the string denoting the finished_at field in the database.
	FieldFinishedAt = "finished_at"
	// FieldIsDeprecated holds the string denoting the is_deprecated field in the database.
	FieldIsDeprecated = "is_deprecated"
	// FieldIsSensitive holds the string denoting the is_sensitive field in the database.
	FieldIsSensitive = "is_sensitive"
	// FieldIsHotGpu holds the string denoting the is_hot_gpu field in the database.
	FieldIsHotGpu = "is_hot_gpu"
	// EdgeGpu holds the string denoting the gpu edge name in mutations.
	EdgeGpu = "gpu"
	// Table holds the table name of the price in the database.
	Table = "prices"
	// GpuTable is the table that holds the gpu relation/edge.
	GpuTable = "prices"
	// GpuInverseTable is the table name for the Gpu entity.
	// It exists in this package in order to avoid circular dependency with the "gpu" package.
	GpuInverseTable = "gpus"
	// GpuColumn is the table column denoting the gpu relation/edge.
	GpuColumn = "gpu_id"
)

// Columns holds all SQL columns for price fields.
var Columns = []string{
	FieldID,
	FieldCreatedBy,
	FieldUpdatedBy,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldGpuID,
	FieldGpuVersion,
	FieldMissionType,
	FieldMissionCategory,
	FieldMissionBillingType,
	FieldCep,
	FieldOriginalCep,
	FieldStartedAt,
	FieldFinishedAt,
	FieldIsDeprecated,
	FieldIsSensitive,
	FieldIsHotGpu,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedBy holds the default value on creation for the "created_by" field.
	DefaultCreatedBy int64
	// DefaultUpdatedBy holds the default value on creation for the "updated_by" field.
	DefaultUpdatedBy int64
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultDeletedAt holds the default value on creation for the "deleted_at" field.
	DefaultDeletedAt time.Time
	// DefaultGpuID holds the default value on creation for the "gpu_id" field.
	DefaultGpuID int64
	// DefaultCep holds the default value on creation for the "cep" field.
	DefaultCep int64
	// DefaultOriginalCep holds the default value on creation for the "original_cep" field.
	DefaultOriginalCep int64
	// DefaultIsDeprecated holds the default value on creation for the "is_deprecated" field.
	DefaultIsDeprecated bool
	// DefaultIsSensitive holds the default value on creation for the "is_sensitive" field.
	DefaultIsSensitive bool
	// DefaultIsHotGpu holds the default value on creation for the "is_hot_gpu" field.
	DefaultIsHotGpu bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() int64
)

const DefaultGpuVersion enums.GpuVersion = "RTX2060"

// GpuVersionValidator is a validator for the "gpu_version" field enum values. It is called by the builders before save.
func GpuVersionValidator(gv enums.GpuVersion) error {
	switch gv {
	case "unknown", "RTX2060", "RTX2060Ti", "RTX2070", "RTX2070Ti", "RTX2080", "RTX2080Ti", "RTX3060", "RTX3060Ti", "RTX3070", "RTX3070Ti", "RTX3080", "RTX3080Ti", "RTX3090", "RTX3090Ti", "RTX4060", "RTX4060Ti", "RTX4070", "RTX4070Ti", "RTX4080", "RTX4090", "RTX4090D", "A800", "A100", "V100", "ComputilityKing-I", "Ascend910ProB", "P40":
		return nil
	default:
		return fmt.Errorf("price: invalid enum value for gpu_version field: %q", gv)
	}
}

const DefaultMissionType enums.MissionType = "txt2img"

// MissionTypeValidator is a validator for the "mission_type" field enum values. It is called by the builders before save.
func MissionTypeValidator(mt enums.MissionType) error {
	switch mt {
	case "unknown", "sd_time", "sd_pro_time", "txt2img", "img2img", "jp_time", "wt_time", "extra-single-image", "sd_api", "key_pair", "jp_dk_time", "ssh_time", "sglang_time", "sd_time_plan", "sd_pro_time_plan", "wt_time_plan", "jp_time_plan", "jp_dk_time_plan", "ssh_time_plan", "sglang_time_plan", "sd_tomato_time", "sd_tomato_time_plan", "sd_cmd_time", "sd_cmd_time_plan", "sd_tian_time", "sd_tian_time_plan", "sd_bingo_time", "sd_bingo_time_plan", "fooocus_time", "fooocus_time_plan", "fooocus_lan_que_time", "fooocus_lan_que_time_plan", "fooocus_sha_api_time", "fooocus_sha_api_time_plan", "tabby_time", "tabby_time_plan", "ollama_time", "ollama_time_plan", "jp_conda_time", "jp_conda_time_plan", "jp_ml_time", "jp_ml_time_plan", "sd_cat_time", "sd_cat_time_plan", "sd_fire_time", "sd_fire_time_plan", "comfyui_time", "comfyui_time_plan", "comfyui_pro_time", "comfyui_pro_time_plan", "comfyui_advance_time", "comfyui_advance_time_plan", "jp_dk_3_time", "jp_dk_3_time_plan", "sd_xl_time", "sd_xl_time_plan", "sd_chick_time", "sd_chick_time_plan", "ascend_time", "ascend_time_plan", "sd_wu_shan_time", "sd_wu_shan_time_plan", "sd_lang_time", "sd_lang_time_plan", "sd_zhi_dao_time", "sd_zhi_dao_time_plan", "comfyui_ke_time", "comfyui_ke_time_plan", "comfyui_a_shuo_time", "comfyui_a_shuo_time_plan", "comfyui_jia_ming_time", "comfyui_jia_ming_time_plan", "chatchat_time", "chatchat_time_plan", "chat_tts_time", "chat_tts_time_plan", "lora_time", "lora_time_plan", "lora_flux_time", "lora_flux_time_plan", "lora_flux_gym_time", "lora_flux_gym_time_plan", "fooocus_wu_time", "fooocus_wu_time_plan", "svd_back_time", "svd_back_time_plan", "sd_ji_time", "sd_ji_time_plan", "sd_shang_jin_time", "sd_shang_jin_time_plan", "sd_xiao_chun_time", "sd_xiao_chun_time_plan", "comfyui_wu_time", "comfyui_wu_time_plan", "comfyui_liu_time", "comfyui_liu_time_plan", "sd_qkk_time", "sd_qkk_time_plan", "comfyui_li_time", "comfyui_li_time_plan", "comfyui_nenly_time", "comfyui_nenly_time_plan", "comfyui_ou_time", "comfyui_ou_time_plan", "comfyui_lu_time", "comfyui_lu_time_plan", "comfyui_jiang_time", "comfyui_jiang_time_plan", "comfyui_star_time", "comfyui_star_time_plan", "waiting_time", "waiting_time_plan", "waiting_al_time", "waiting_al_time_plan", "opencl_core_time", "opencl_core_time_plan", "io_paint_time", "io_paint_time_plan":
		return nil
	default:
		return fmt.Errorf("price: invalid enum value for mission_type field: %q", mt)
	}
}

const DefaultMissionCategory enums.MissionCategory = "SD"

// MissionCategoryValidator is a validator for the "mission_category" field enum values. It is called by the builders before save.
func MissionCategoryValidator(mc enums.MissionCategory) error {
	switch mc {
	case "unknown", "SD", "JP", "WT", "JP_DK", "SSH", "SGLANG", "SD_TOMATO", "SD_CMD", "SD_TIAN", "SD_BINGO", "FOOOCUS", "FOOOCUS_LAN_QUE", "FOOOCUS_SHA_API", "TABBY", "OLLAMA", "JP_CONDA", "SD_CAT", "SD_FIRE", "COMFYUI", "SD_XL", "SD_CHICK", "ASCEND", "SD_WU_SHAN", "SD_LANG", "SD_ZHI_DAO", "COMFYUI_KE", "COMFYUI_A_SHUO", "COMFYUI_JIA_MING", "CHATCHAT", "CHATTTS", "LORA", "LORA_FLUX_GYM", "FOOOCUS_WU", "SVD_BACK", "SD_JI", "SD_SHANG_JIN", "SD_XIAO_CHUN", "COMFYUI_WU", "COMFYUI_LIU", "SD_QKK", "COMFYUI_LI", "COMFYUI_NENLY", "COMFYUI_OU", "COMFYUI_LU", "COMFYUI_JIANG", "COMFYUI_STAR", "WAITING", "WAITING_AL", "OPEN_CL", "IO_PAINT":
		return nil
	default:
		return fmt.Errorf("price: invalid enum value for mission_category field: %q", mc)
	}
}

const DefaultMissionBillingType enums.MissionBillingType = "count"

// MissionBillingTypeValidator is a validator for the "mission_billing_type" field enum values. It is called by the builders before save.
func MissionBillingTypeValidator(mbt enums.MissionBillingType) error {
	switch mbt {
	case "unknown", "time", "count", "hold", "volume", "time_plan_hour", "time_plan_day", "time_plan_week", "time_plan_month":
		return nil
	default:
		return fmt.Errorf("price: invalid enum value for mission_billing_type field: %q", mbt)
	}
}

// OrderOption defines the ordering options for the Price queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedBy orders the results by the created_by field.
func ByCreatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedBy, opts...).ToFunc()
}

// ByUpdatedBy orders the results by the updated_by field.
func ByUpdatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedBy, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByGpuID orders the results by the gpu_id field.
func ByGpuID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGpuID, opts...).ToFunc()
}

// ByGpuVersion orders the results by the gpu_version field.
func ByGpuVersion(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGpuVersion, opts...).ToFunc()
}

// ByMissionType orders the results by the mission_type field.
func ByMissionType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMissionType, opts...).ToFunc()
}

// ByMissionCategory orders the results by the mission_category field.
func ByMissionCategory(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMissionCategory, opts...).ToFunc()
}

// ByMissionBillingType orders the results by the mission_billing_type field.
func ByMissionBillingType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMissionBillingType, opts...).ToFunc()
}

// ByCep orders the results by the cep field.
func ByCep(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCep, opts...).ToFunc()
}

// ByOriginalCep orders the results by the original_cep field.
func ByOriginalCep(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOriginalCep, opts...).ToFunc()
}

// ByStartedAt orders the results by the started_at field.
func ByStartedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStartedAt, opts...).ToFunc()
}

// ByFinishedAt orders the results by the finished_at field.
func ByFinishedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFinishedAt, opts...).ToFunc()
}

// ByIsDeprecated orders the results by the is_deprecated field.
func ByIsDeprecated(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsDeprecated, opts...).ToFunc()
}

// ByIsSensitive orders the results by the is_sensitive field.
func ByIsSensitive(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsSensitive, opts...).ToFunc()
}

// ByIsHotGpu orders the results by the is_hot_gpu field.
func ByIsHotGpu(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsHotGpu, opts...).ToFunc()
}

// ByGpuField orders the results by gpu field.
func ByGpuField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newGpuStep(), sql.OrderByField(field, opts...))
	}
}
func newGpuStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(GpuInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, GpuTable, GpuColumn),
	)
}
