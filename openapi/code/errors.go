package code

import "github.com/databricks/databricks-sdk-go/openapi"

type ExceptionType struct {
	Named
	StatusCode int
	Inherit    *Named
}

type ErrorMappingRule struct {
	Named
	StatusCode int
	ErrorCode  string
}

func (b *Batch) ErrorStatusCodeMapping() (rules []ErrorMappingRule) {
	for _, em := range openapi.ErrorStatusCodeMapping {
		rules = append(rules, ErrorMappingRule{
			StatusCode: em.StatusCode,
			Named: Named{
				Name: em.ErrorCode,
			},
		})
	}
	return rules
}

func (b *Batch) ErrorCodeMapping() (rules []ErrorMappingRule) {
	for _, em := range openapi.ErrorCodeMapping {
		rules = append(rules, ErrorMappingRule{
			ErrorCode: em.ErrorCode,
			Named: Named{
				Name: em.ErrorCode,
			},
		})
	}
	return rules
}

func (b *Batch) ExceptionTypes() []ExceptionType {
	statusCodeMapping := map[int]*Named{}
	exceptionTypes := []ExceptionType{}
	for _, em := range openapi.ErrorStatusCodeMapping {
		statusCodeMapping[em.StatusCode] = &Named{
			Name: em.ErrorCode,
		}
		exceptionTypes = append(exceptionTypes, ExceptionType{
			Named: Named{
				Name:        em.ErrorCode,
				Description: em.Description,
			},
			StatusCode: em.StatusCode,
		})
	}
	for _, em := range openapi.ErrorCodeMapping {
		exceptionTypes = append(exceptionTypes, ExceptionType{
			Named: Named{
				Name:        em.ErrorCode,
				Description: em.Description,
			},
			StatusCode: em.StatusCode,
			Inherit:    statusCodeMapping[em.StatusCode],
		})
	}
	return exceptionTypes
}
