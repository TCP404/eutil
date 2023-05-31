package packer

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type simpleSonStruct struct {
	SF1 string  `json:"sf1"`
	SF2 float64 `json:"sf2"`
}

type SimpleStruct struct {
	F1 string `json:"f1"`
	F2 int    `json:"f2"`
}

type NamingNestedStruct struct {
	F1  string `json:"f1"`
	F2  int    `json:"f2"`
	Sub simpleSonStruct
}

type NamingTagNestedStruct struct {
	F1  string          `json:"f1"`
	F2  int             `json:"f2"`
	Sub simpleSonStruct `json:"son"`
}

type AnonymousNestedStruct struct {
	F1 string `json:"f1"`
	F2 int    `json:"f2"`
	simpleSonStruct
}

type AnonymousTagNestedStruct struct {
	F1              string `json:"f1"`
	F2              int    `json:"f2"`
	simpleSonStruct `json:"son"`
}

type sonStruct struct {
	SF1      string         `json:"sf1"`
	SF2      float64        `json:"sf2"`
	Grandson grandsonStruct `json:"grandson"`
}

type grandsonStruct struct {
	SF1 string `json:"gsf1"`
	SF2 bool   `json:"gsf2"`
	SF3 *int   `json:"gsf3"`
}

type GrandparentStruct struct {
	F1  string    `json:"f1"`
	F2  int       `json:"f2"`
	Son sonStruct `json:"son"`
}
type Link struct {
	Type          string `json:"type"`           // 图片image、视频video、文件file等
	Url           string `json:"url"`            // 内容关联或提取的视频或文件链接
	Name          string `json:"name,omitempty"` // 文件名
	Value         string `json:"value,omitempty"`
	IsManualField        // 是否是人工运营添加的数据
}

// IsManualField 人工录入标识字段
type IsManualField struct {
	IsManual bool `json:"is_manual"`
}

// AggregateCommon 汇聚表共用字段
type AggregateCommon struct {
	Channel       string    `json:"channel"`     // 数据渠道
	Source        string    `json:"source"`      // 来源相关
	EventTime     time.Time `json:"event_time"`  // 事件时间，对于不停更新的场景为第一次插入时间
	InsertTime    time.Time `json:"insert_time"` // 汇聚表更新时间
	CreateTime    time.Time `json:"create_time"`
	IsManualField           // 是否是人工运营添加的数据
}

// ProductCommon 产品相关字段
type ProductCommon struct {
	DataId        string       `json:"data_id"`                  // 数据ID，汇聚表ID
	Company       string       `json:"company"`                  // 公司，支持模糊搜索
	Product       string       `json:"product"`                  // 产品
	Hits          []string     `json:"hits"`                     // 产品命中关键词
	Filters       []string     `json:"filters,omitempty"`        // 产品匹配过滤词
	HitRules      []string     `json:"hit_rules,omitempty"`      // 匹配命中规则名
	FltRules      []string     `json:"flt_rules,omitempty"`      // 匹配过滤规则名
	RuleScore     int32        `json:"rule_score,omitempty"`     // 规则分数
	CategoryRules []string     `json:"category_rules,omitempty"` // 规则分类
	Industry      []string     `json:"industry"`                 // 行业
	IndustryInfo  IndustryInfo `json:"industry_info,omitempty"`  // 行业信息
}

// IndustryInfo 行业字段
type IndustryInfo struct {
	Primary   []string `json:"primary"`
	Secondary []string `json:"secondary"`
}

// OperateCommon 运营相关字段
type OperateCommon struct {
	Comments    string     `json:"comments,omitempty"`     // 备注
	Status      int32      `json:"status,omitempty"`       // 状态：0 待审核, 1 已审核, 2 已删除, 3 审核未通过, 4 失效
	Operator    string     `json:"operator,omitempty"`     // 操作人
	Priority    int32      `json:"priority,omitempty"`     // 优先级：0-100
	Potential   int32      `json:"potential,omitempty"`    // 潜在客户：0 未标记, 1 潜在客户, 2 非潜在客户
	AlertTime   *time.Time `json:"alert_time,omitempty"`   // 预警时间
	AlertStatus int32      `json:"alert_status,omitempty"` // 预警状态：0 未预警, 1已预警
	Origin      int32      `json:"origin,omitempty"`       // 运营数据来源：1 对未命中产品数据运营得到, 2对线上表修改, 3后台重刷数据
	UpdateTime  *time.Time `json:"update_time,omitempty"`  // 运营更新时间，一般时重跑规则会更新数据
	TaskId      int64      `json:"task_id,omitempty"`      // 重刷任务ID
}

type BreachAggBase struct {
	Name           string   `json:"name"`     // 文件名或文章标题
	Type           string   `json:"type"`     // 文件类型
	Author         string   `json:"author"`   // 发布人
	Content        string   `json:"content"`  // 数据泄露内容，允许模糊搜索
	Url            string   `json:"url"`      // 数据泄露来源链接
	Keywords       []string `json:"keywords"` // 命中的敏感关键词
	Links          []Link   `json:"links"`
	Score          int32    `json:"score"` // 数据资产风险分：0 无风险, 10 疑似风险, 50 高风险
	ExtractionCode string   `json:"extraction_code"`
	SimilarId      string   `json:"similar_id"`
	Origin         string   `json:"origin"`
}

// 数据泄露产品命中表
type breachHit struct {
	ProductCommon
	AggregateCommon
	BreachAggBase
}
type BreachOperateBase struct {
	DataType      string `json:"data_type"`      // 数据类型：敏感代码、客户信息等
	BreachChannel string `json:"breach_channel"` // 泄露渠道
	Level         string `json:"level"`          // 危害级别
	DataLevel     string `json:"data_level"`     // 泄露数据量级
	DateRange     string `json:"date_range"`     // 影响时间范围，例如20210909-20210910
	VerifyStatus  int32  `json:"verify_status"`  // 运营验证状态：0 未验证, 1 验证为真, 2 验证为假, 3 无法验证
	ConfirmStatus int32  `json:"confirm_status"` // 客户确认状态：0 未确认, 1 已确认
	EventType     int32  `json:"event_type"`     // 客户确认事件类型：-1 全部, 0 未处置, 1 真实数据, 2 虚假数据, 3 历史数据, 4 误报, 5 忽略
	Warning       int32  `json:"warning"`        // 是否有效数据泄露
	HandleStatus  int32  `json:"handle_status"`  // 处置状态: 是否下架
	UrlStatus     int32  `json:"url_status"`     // 网盘链接状态：0未检测 1有效 2已失效 3检测失败
}
type BreachOperate struct {
	OperateCommon
	BreachOperateBase
}
type BreachAutoAudit struct {
	Status       int32      `json:"status"`                // 状态：0 待审核, 1 已审核, 2 已删除, 3 审核未通过, 4 失效
	VerifyStatus int32      `json:"verify_status"`         // 运营验证状态：0 未验证, 1 验证为真, 2 验证为假, 3 无法验证
	HitRuleID    int32      `json:"hit_rule_id"`           // 命中的规则id
	UpdateTime   *time.Time `json:"update_time,omitempty"` // 自动审核时间
}

// 数据泄露运营表
type breachOpt struct {
	breachHit
	Opt  BreachOperate   `json:"opt"`
	Auto BreachAutoAudit `json:"auto"`
}

// BreachOpt 数据泄露运营表
type BreachOpt struct {
	breachOpt
	Id string `json:"id"`
}

func TestDescribeType(t *testing.T) {
	var (
		simpleStruct             = &SimpleStruct{}
		namingNestedStruct       = &NamingNestedStruct{}
		namingTagNestedStruct    = &NamingTagNestedStruct{}
		anonymousNestedStruct    = &AnonymousNestedStruct{}
		anonymousTagNestedStruct = &AnonymousTagNestedStruct{}
		grandparentStruct        = &GrandparentStruct{}
		breachOpt                = &BreachOpt{}
	)
	type testCase struct {
		got  TypeDescriptor
		want TypeDescriptor
	}
	var (
		tests = make(map[string]testCase, 0)
		got   TypeDescriptor
		err   error
	)

	got, err = DescribeType(breachOpt)
	require.Nil(t, err)
	tests[reflect.TypeOf(breachOpt).Name()] = testCase{
		got: got,
		want: TypeDescriptor{
			// breachOpt.breachHit.ProductCommon
			"data_id":                 fieldInfo{idx: []int{0, 0, 0, 0}, kind: reflect.String},
			"company":                 fieldInfo{idx: []int{0, 0, 0, 1}, kind: reflect.String},
			"product":                 fieldInfo{idx: []int{0, 0, 0, 2}, kind: reflect.String},
			"hits":                    fieldInfo{idx: []int{0, 0, 0, 3}, kind: reflect.Slice},
			"filters":                 fieldInfo{idx: []int{0, 0, 0, 4}, kind: reflect.Slice},
			"hit_rules":               fieldInfo{idx: []int{0, 0, 0, 5}, kind: reflect.Slice},
			"flt_rules":               fieldInfo{idx: []int{0, 0, 0, 6}, kind: reflect.Slice},
			"rule_score":              fieldInfo{idx: []int{0, 0, 0, 7}, kind: reflect.Int32},
			"category_rules":          fieldInfo{idx: []int{0, 0, 0, 8}, kind: reflect.Slice},
			"industry":                fieldInfo{idx: []int{0, 0, 0, 9}, kind: reflect.Slice},
			"industry_info":           fieldInfo{idx: []int{0, 0, 0, 10}, kind: reflect.Struct},
			"industry_info.primary":   fieldInfo{idx: []int{0, 0, 0, 10, 0}, kind: reflect.Slice},
			"industry_info.secondary": fieldInfo{idx: []int{0, 0, 0, 10, 1}, kind: reflect.Slice},

			// breachOpt.breachHit.AggregateCommon
			"channel":     fieldInfo{idx: []int{0, 0, 1, 0}, kind: reflect.String},
			"source":      fieldInfo{idx: []int{0, 0, 1, 1}, kind: reflect.String},
			"event_time":  fieldInfo{idx: []int{0, 0, 1, 2}, kind: reflect.Struct},
			"insert_time": fieldInfo{idx: []int{0, 0, 1, 3}, kind: reflect.Struct},
			"create_time": fieldInfo{idx: []int{0, 0, 1, 4}, kind: reflect.Struct},
			"is_manual":   fieldInfo{idx: []int{0, 0, 1, 5, 0}, kind: reflect.Bool},

			// breachOpt.breachHit.BreachAggBase
			"name":            fieldInfo{idx: []int{0, 0, 2, 0}, kind: reflect.String},
			"type":            fieldInfo{idx: []int{0, 0, 2, 1}, kind: reflect.String},
			"author":          fieldInfo{idx: []int{0, 0, 2, 2}, kind: reflect.String},
			"content":         fieldInfo{idx: []int{0, 0, 2, 3}, kind: reflect.String},
			"url":             fieldInfo{idx: []int{0, 0, 2, 4}, kind: reflect.String},
			"keywords":        fieldInfo{idx: []int{0, 0, 2, 5}, kind: reflect.Slice},
			"links":           fieldInfo{idx: []int{0, 0, 2, 6}, kind: reflect.Slice},
			"score":           fieldInfo{idx: []int{0, 0, 2, 7}, kind: reflect.Int32},
			"extraction_code": fieldInfo{idx: []int{0, 0, 2, 8}, kind: reflect.String},
			"similar_id":      fieldInfo{idx: []int{0, 0, 2, 9}, kind: reflect.String},
			"origin":          fieldInfo{idx: []int{0, 0, 2, 10}, kind: reflect.String},

			// breachOpt.(opt)BreachOperate.OperateCommon
			"opt.comments":     fieldInfo{idx: []int{0, 1, 0, 0}, kind: reflect.String},
			"opt.status":       fieldInfo{idx: []int{0, 1, 0, 1}, kind: reflect.Int32},
			"opt.operator":     fieldInfo{idx: []int{0, 1, 0, 2}, kind: reflect.String},
			"opt.priority":     fieldInfo{idx: []int{0, 1, 0, 3}, kind: reflect.Int32},
			"opt.potential":    fieldInfo{idx: []int{0, 1, 0, 4}, kind: reflect.Int32},
			"opt.alert_time":   fieldInfo{idx: []int{0, 1, 0, 5}, kind: reflect.Ptr},
			"opt.alert_status": fieldInfo{idx: []int{0, 1, 0, 6}, kind: reflect.Int32},
			"opt.origin":       fieldInfo{idx: []int{0, 1, 0, 7}, kind: reflect.Int32},
			"opt.update_time":  fieldInfo{idx: []int{0, 1, 0, 8}, kind: reflect.Ptr},
			"opt.task_id":      fieldInfo{idx: []int{0, 1, 0, 9}, kind: reflect.Int64},

			// breachOpt.(opt)BreachOperate.BreachOperateBase
			"opt.data_type":      fieldInfo{idx: []int{0, 1, 1, 0}, kind: reflect.String},
			"opt.breach_channel": fieldInfo{idx: []int{0, 1, 1, 1}, kind: reflect.String},
			"opt.level":          fieldInfo{idx: []int{0, 1, 1, 2}, kind: reflect.String},
			"opt.data_level":     fieldInfo{idx: []int{0, 1, 1, 3}, kind: reflect.String},
			"opt.date_range":     fieldInfo{idx: []int{0, 1, 1, 4}, kind: reflect.String},
			"opt.verify_status":  fieldInfo{idx: []int{0, 1, 1, 5}, kind: reflect.Int32},
			"opt.confirm_status": fieldInfo{idx: []int{0, 1, 1, 6}, kind: reflect.Int32},
			"opt.event_type":     fieldInfo{idx: []int{0, 1, 1, 7}, kind: reflect.Int32},
			"opt.warning":        fieldInfo{idx: []int{0, 1, 1, 8}, kind: reflect.Int32},
			"opt.handle_status":  fieldInfo{idx: []int{0, 1, 1, 9}, kind: reflect.Int32},
			"opt.url_status":     fieldInfo{idx: []int{0, 1, 1, 10}, kind: reflect.Int32},

			// breachOpt.(auto)BreachAutoAudit
			"auto.status":        fieldInfo{idx: []int{0, 2, 0}, kind: reflect.Int32},
			"auto.verify_status": fieldInfo{idx: []int{0, 2, 1}, kind: reflect.Int32},
			"auto.hit_rule_id":   fieldInfo{idx: []int{0, 2, 2}, kind: reflect.Int32},

			"id": fieldInfo{idx: []int{1}, kind: reflect.String},
		},
	}

	got, err = DescribeType(simpleStruct)
	require.Nil(t, err)
	tests[reflect.TypeOf(simpleStruct).Name()] = testCase{
		got: got,
		want: TypeDescriptor{
			"f1": fieldInfo{idx: []int{0}, kind: reflect.String},
			"f2": fieldInfo{idx: []int{1}, kind: reflect.Int},
		},
	}

	got, err = DescribeType(namingNestedStruct)
	require.Nil(t, err)
	tests[reflect.TypeOf(namingNestedStruct).Name()] = testCase{
		got: got,
		want: TypeDescriptor{
			"f1":      fieldInfo{idx: []int{0}, kind: reflect.String},
			"f2":      fieldInfo{idx: []int{1}, kind: reflect.Int},
			"Sub.sf1": fieldInfo{idx: []int{2, 0}, kind: reflect.String},
			"Sub.sf2": fieldInfo{idx: []int{2, 1}, kind: reflect.Float64},
		},
	}

	got, err = DescribeType(namingTagNestedStruct)
	require.Nil(t, err)
	tests[reflect.TypeOf(namingTagNestedStruct).Name()] = testCase{
		got: got,
		want: TypeDescriptor{
			"f1":      fieldInfo{idx: []int{0}, kind: reflect.String},
			"f2":      fieldInfo{idx: []int{1}, kind: reflect.Int},
			"son.sf1": fieldInfo{idx: []int{2, 0}, kind: reflect.String},
			"son.sf2": fieldInfo{idx: []int{2, 1}, kind: reflect.Float64},
		},
	}

	got, err = DescribeType(anonymousNestedStruct)
	require.Nil(t, err)
	tests[reflect.TypeOf(anonymousNestedStruct).Name()] = testCase{
		got: got,
		want: TypeDescriptor{
			"f1":  fieldInfo{idx: []int{0}, kind: reflect.String},
			"f2":  fieldInfo{idx: []int{1}, kind: reflect.Int},
			"sf1": fieldInfo{idx: []int{2, 0}, kind: reflect.String},
			"sf2": fieldInfo{idx: []int{2, 1}, kind: reflect.Float64},
		},
	}

	got, err = DescribeType(anonymousTagNestedStruct)
	require.Nil(t, err)
	tests[reflect.TypeOf(anonymousTagNestedStruct).Name()] = testCase{
		got: got,
		want: TypeDescriptor{
			"f1":      fieldInfo{idx: []int{0}, kind: reflect.String},
			"f2":      fieldInfo{idx: []int{1}, kind: reflect.Int},
			"son.sf1": fieldInfo{idx: []int{2, 0}, kind: reflect.String},
			"son.sf2": fieldInfo{idx: []int{2, 1}, kind: reflect.Float64},
		},
	}

	got, err = DescribeType(grandparentStruct)
	require.Nil(t, err)
	tests[reflect.TypeOf(grandparentStruct).Name()] = testCase{
		got: got,
		want: TypeDescriptor{
			"f1":                fieldInfo{idx: []int{0}, kind: reflect.String},
			"f2":                fieldInfo{idx: []int{1}, kind: reflect.Int},
			"son.sf1":           fieldInfo{idx: []int{2, 0}, kind: reflect.String},
			"son.sf2":           fieldInfo{idx: []int{2, 1}, kind: reflect.Float64},
			"son.grandson.gsf1": fieldInfo{idx: []int{2, 2, 0}, kind: reflect.String},
			"son.grandson.gsf2": fieldInfo{idx: []int{2, 2, 1}, kind: reflect.Bool},
			"son.grandson.gsf3": fieldInfo{idx: []int{2, 2, 2}, kind: reflect.Ptr},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			if !reflect.DeepEqual(tc.got, tc.want) {
				//t.Errorf("DescribeType() got = %v, want %v", tc.got, tc.want)
				for k, v := range tc.got {
					//t.Log(k)
					//t.Logf("%v, want %v", v, tc.want[k])
					if !reflect.DeepEqual(v, tc.want[k]) {
						t.Errorf("got: %v, want %v", v, tc.want[k])
					}
				}
			}
		})
	}
}
