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

type Company struct {
	Id int32 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Trusty bool `json:"trusty,omitempty"`
	Money int32 `json:"money,omitempty"`
}