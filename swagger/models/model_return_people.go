/*
 * SEP6-movies backend
 *
 * Backend part of the project delivered for SEP6 course - Movies platform  Authors of project:  Konrad Piotrowski (280053) Aleksander Stefan Bialik (280027)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ReturnPeople struct {
	Page int64 `json:"page,omitempty"`

	TotalPages int64 `json:"total_pages,omitempty"`

	TotalResults int64 `json:"total_results,omitempty"`

	People []Person `json:"people,omitempty"`
}
