package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/stark-sim/cephalon-ent/common"
	"github.com/stark-sim/cephalon-ent/pkg/enums"
	"time"
)

type Device struct {
	ent.Schema
}

func (Device) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").StructTag(`json:"user_id,string"`).Default(0).Comment("外键用户 id"),
		field.Int64("gift_mission_config_id").StructTag(`json:"gift_mission_config_id,string"`).Default(0).Comment("外键补贴任务配置 id"),
		field.String("serial_number").StructTag(`json:"serial_number"`).Default("").Comment("设备唯一序列号"),
		field.Enum("state").Values("On", "Down", "Init").Default("Init").Comment("设备状态"),
		field.Int64("sum_cep").StructTag(`json:"sum_cep"`).Default(0).Comment("该设备总获得利润"),
		field.Bool("linking").Default(false).StructTag(`json:"linking"`).Comment("设备是否正在对接中"),
		field.Enum("binding_status").GoType(enums.DeviceBindingStatusInit).Default(string(enums.DeviceBindingStatusInit)).StructTag(`json:"binding_status"`).Comment("设备的绑定状态"),
		field.Enum("status").GoType(enums.DeviceStatusFree).Default(string(enums.DeviceStatusOffline)).StructTag(`json:"status"`).Comment("设备状态"),
		field.String("name").StructTag(`json:"name"`).Default("默认设备名称").Comment("设备名称"),
		field.String("manage_name").StructTag(`json:"manage_name"`).Default("默认管理设备名称").Comment("运维管理设备名称"),
		field.Enum("type").GoType(enums.DeviceTypeOrdinary).Default(string(enums.DeviceTypeOrdinary)).StructTag(`json:"type"`).Comment("设备类型"),
		field.Int64("cores_number").StructTag(`json:"cores_number"`).Default(0).Comment("核心数"),
		field.String("cpu").StructTag(`json:cpu`).Default("").Comment("CPU型号"),
		field.String("cpus").GoType([]string{}).Optional().ValueScanner(common.Bytes2StringSliceValueScanner{}).SchemaType(map[string]string{dialect.Postgres: "bytea"}).StructTag(`json:"cpus"`).Comment("CPU型号"),
		field.Int64("memory").StructTag(`json:"memory"`).Default(0).Comment("内存(单位:G)"),
		field.Float32("disk").SchemaType(map[string]string{dialect.Postgres: "NUMERIC(10,4)"}).StructTag(`json:disk`).Default(0).Comment("硬盘(单位:T)"),
		field.Float("delay").StructTag(`json:"delay"`).Default(0).Comment("延迟(单位:ms)"),
		field.Float("gpu_temperature").StructTag(`json:"gpu_temperature"`).Default(0).Comment("GPU 温度(单位:℃)"),
		field.Float("cpu_temperature").StructTag(`json:"cpu_temperature"`).Default(0).Comment("CPU 温度(单位:℃)"),
		field.Enum("stability").GoType(enums.DeviceStabilityTypeGood).StructTag(`json:"stability"`).Default(string(enums.DeviceStabilityTypeGood)).Comment("设备稳定性"),
		field.String("version").StructTag(`json:"version"`).Default("无版本号").Comment("设备版本"),
		field.Enum("fault").GoType(enums.DeviceFaultTypeOK).StructTag(`json:"fault"`).Default(string(enums.DeviceFaultTypeOK)).Comment("故障信息"),
		field.Enum("rank").GoType(enums.DeviceRankTypeNormal).StructTag(`json:"rank"`).Default(string(enums.DeviceRankTypeNormal)).Comment("设备等级(目前阶段就是黑名单)"),
		field.Int("free_gpu_num").StructTag(`json:"free_gpu_num"`).Default(0).Comment("空闲GPU数量"),
		field.Time("rank_at").Default(time.Now).Optional().Nillable().StructTag(`json:"rank_at"`).Comment("评定设备等级的时刻，带时区"),
		field.Time("stability_at").Default(time.Now).Optional().Nillable().StructTag(`json:"stability_at"`).Comment("判定稳定性的时刻，带时区"),
		field.Time("high_temperature_at").Default(time.Now).Optional().Nillable().StructTag(`json:"high_temperature_at"`).Comment("温度超标的时刻，带时区"),
		field.Enum("hosting_type").GoType(enums.DeviceHostingTypeNo).Default(string(enums.DeviceHostingTypeNo)).StructTag(`json:"hosting_type"`).Comment("托管类型，非托管/半托管/全托管等"),
		field.Enum("mission_tag").GoType(enums.DeviceMissionTagNo).Default(string(enums.DeviceMissionTagNo)).StructTag(`json:"mission_tag"`).Comment("可接特殊任务类型标签"),
		field.Time("last_abnormal_at").Default(time.Now).Optional().Nillable().StructTag(`json:"last_abnormal_at"`).Comment("最后一次异常时刻，带时区（用于设备稳定性升级）"),
	}
}

// Edges of the Device.
func (Device) Edges() []ent.Edge {
	return []ent.Edge{
		// 逻辑外键
		edge.From("user", User.Type).Ref("devices").Field("user_id").Unique().Required(),
		edge.From("gift_mission_config", GiftMissionConfig.Type).Ref("devices").Field("gift_mission_config_id").Unique().Required(),
		edge.To("mission_produce_orders", MissionProduceOrder.Type),
		edge.To("user_devices", UserDevice.Type),
		edge.To("device_gpu_missions", DeviceGpuMission.Type),
		edge.To("frpc_infos", FrpcInfo.Type),
		edge.To("mission_orders", MissionOrder.Type),
		edge.To("mission_productions", MissionProduction.Type),
		edge.To("device_reboot_times", DeviceRebootTime.Type),
		edge.To("trouble_deducts", TroubleDeduct.Type),
		edge.To("device_states", DeviceState.Type),
		edge.To("device_offline_records", DeviceOfflineRecord.Type),
		edge.To("mission_failed_feedbacks", MissionFailedFeedback.Type),
	}
}

// Indexes of Device
func (Device) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
	}
}

// Mixin of Device
func (Device) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (Device) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("设备，对应客户端 core，记录心跳等信息"),
	}
}
