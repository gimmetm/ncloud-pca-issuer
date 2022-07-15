/*
 * api
 *
 * <br/>https://pca.apigw.ntruss.com/api/v1
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package privateca

type BaseResponse struct {
	Code *ResponseCode `json:"code,omitempty"`

	Msg *string `json:"msg,omitempty"`

	Data *interface{} `json:"data,omitempty"`
}
