/*
 * api
 *
 * <br/>https://pca.apigw.ntruss.com/api/v1
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package privateca

type SignCsr struct {

	// PEM 포맷의 CSR(인증서 서명 요청)
	CsrPem *string `json:"csrPem,omitempty"`

	// DER 포맷의 CSR(인증서 서명 요청)
	CsrPemDer *[]byte `json:"csrPemDer,omitempty"`

	// 인증서 유효 기간. 일 단위 (1 <= period <= 3650 혹은 period=Max)
	Period *string `json:"period"`

	// 인증서 키 타입
	KeyType *string `json:"keyType,omitempty"`

	// 인증서 키 Bits
	KeyBits *string `json:"keyBits,omitempty"`
}
