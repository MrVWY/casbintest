package service

import (
	"casbintest/module/model"
)

type CasbinInfo struct {
	Path   string `json:"path" form:"path"`
	Method string `json:"method" form:"method"`
}
type CasbinCreateRequest struct {
	RoleId      string       `json:"role_id" form:"role_id" description:"角色ID"`
	ProId       string       `json:"pro_id" form:"pro_id" description:"项目ID"`
	CasbinInfos []CasbinInfo `json:"casbin_infos" description:"权限模型列表"`
}

type CasbinUpdateRequest struct {
	OldPath    string     `json:"old_path"`
	OldMethod  string     `json:"old_method"`
	CasbinInfo CasbinInfo `json:"casbin_info" description:"权限模型列表"`
}

type CasbinListResponse struct {
	List []CasbinInfo `json:"list" form:"list"`
}

type CasbinListRequest struct {
	RoleID string `json:"role_id" form:"role_id"`
}

func CasbinCreate(param *CasbinCreateRequest) error {
	for _, v := range param.CasbinInfos {
		cm := model.CasbinModel{
			PType:  "p",
			RoleId: param.RoleId,
			ProId:  param.ProId,
			Path:   v.Path,
			Method: v.Method,
		}
		return cm.Create()
	}
	return nil
}

func CasbinList(param *CasbinListRequest) [][]string {
	cm := model.CasbinModel{RoleId: param.RoleID}
	return cm.List()
}
