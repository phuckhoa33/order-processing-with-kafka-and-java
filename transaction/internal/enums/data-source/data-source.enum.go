package data_source_enum

type Industry string

const (
	IndustryNews       Industry = "News"
	IndustryFinance    Industry = "Finance"
	IndustryTechnology Industry = "Technology"
	IndustryHealth     Industry = "Health"
	IndustryEducation  Industry = "Education"
)
