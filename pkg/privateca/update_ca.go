/*
 * api
 *
 * <br/>https://pca.apigw.ntruss.com/api/v1
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package privateca

type UpdateCa struct {

	// CA 상태 변경
	Status *string `json:"status,omitempty"`

	// 메모 변경
	Memo *string `json:"memo,omitempty"`
}
