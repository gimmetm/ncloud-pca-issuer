/*
 * api
 *
 * <br/>https://pca.apigw.ntruss.com/api/v1
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package privateca

type CreateEndCert struct {

	// 키 타입
	KeyType *string `json:"keyType"`

	// 일 단위 인증서 유효 기간. (1 <= period <= 3650 또는 period=Max)
	Period *string `json:"period"`

	X509EndCertParameters *X509EndCertParameters `json:"X509EndCertParameters"`
}
