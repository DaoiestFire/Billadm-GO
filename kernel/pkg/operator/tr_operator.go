package operator

import (
	"sort"

	"github.com/billadm/models/dto"
)

type TrOperator struct {
	trDtos []*dto.TransactionRecordDto
	offset int
	limit  int
}

func NewTrOperator() *TrOperator {
	return &TrOperator{
		trDtos: make([]*dto.TransactionRecordDto, 0),
	}
}

func (t *TrOperator) Add(trDtos []*dto.TransactionRecordDto) *TrOperator {
	t.trDtos = append(t.trDtos, trDtos...)
	return t
}

func (t *TrOperator) FilterByCategoryTags(categoryTags map[string][]string) *TrOperator {
	if len(categoryTags) == 0 {
		return t
	}

	var filtered []*dto.TransactionRecordDto

	for _, tr := range t.trDtos {
		tagsToMatch, exists := categoryTags[tr.Category]
		if !exists {
			continue // 分类不在筛选条件中，跳过
		}

		if len(tagsToMatch) == 0 {
			// 如果没有标签则判定为选中
			filtered = append(filtered, tr)
			continue
		}

		// 将当前记录的 tags 转为 map 便于查找（避免 O(n*m) 嵌套循环）
		recordTagSet := make(map[string]bool)
		for _, tag := range tr.Tags {
			recordTagSet[tag] = true
		}

		// 检查是否包含至少一个目标标签
		matched := false
		for _, requiredTag := range tagsToMatch {
			if recordTagSet[requiredTag] {
				matched = true
				break
			}
		}

		if matched {
			filtered = append(filtered, tr)
		}
	}
	t.trDtos = filtered
	return t
}

func (t *TrOperator) Sort(sortFields []SortField) *TrOperator {
	if len(sortFields) == 0 || len(t.trDtos) <= 1 {
		return t
	}

	toSort := sortableTrDtos{
		data:       t.trDtos,
		sortFields: sortFields,
	}

	sort.Sort(toSort)

	t.trDtos = toSort.data
	return t
}

func (t *TrOperator) Page(offset, limit int) *TrOperator {
	if offset < 0 {
		offset = 0
	}
	if limit < 0 {
		limit = 0
	}
	t.offset = offset
	t.limit = limit
	return t
}

func (t *TrOperator) Summary() *dto.TrQueryResult {
	total := int64(len(t.trDtos))

	// 初始化 trStatistics
	trStatistics := map[string]float64{
		"income":   0.0,
		"expense":  0.0,
		"transfer": 0.0,
	}

	// 遍历所有记录（分页前）进行金额汇总
	for _, tr := range t.trDtos {
		switch tr.TransactionType {
		case "income", "expense", "transfer":
			trStatistics[tr.TransactionType] += tr.Price
		default:
			// 可选：忽略未知类型，或归入其他
		}
	}

	// 执行分页
	var items []*dto.TransactionRecordDto
	if t.limit == 0 {
		// limit=0 表示不分页（或取全部）
		items = t.trDtos
	} else {
		start := t.offset
		end := t.offset + t.limit

		if start > len(t.trDtos) {
			start = len(t.trDtos)
		}
		if end > len(t.trDtos) {
			end = len(t.trDtos)
		}
		if start < end {
			items = t.trDtos[start:end]
		} else {
			items = []*dto.TransactionRecordDto{}
		}
	}

	return &dto.TrQueryResult{
		Items:        items,
		Total:        total,
		TrStatistics: trStatistics,
	}
}
