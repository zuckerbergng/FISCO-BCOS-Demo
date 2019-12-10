/*
 * 基于区块链的供应链金融平台
 *
 * Blockchain
 *
 * API version: 1.0.0
 * Contact: chris-ju@qq.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Account struct {

	Id int32 `json:"id,omitempty"`

	Money int32 `json:"money,omitempty"`

	ACompany int32 `json:"ACompany,omitempty"`

	BCompany int32 `json:"BCompany,omitempty"`
}
