package handler

import (
	"encoding/json"
	"net/http"

	"github.com/fasikawkn/sample-restaurant-rest-api-master/user"
	"github.com/julienschmidt/httprouter"
)

//AdminRoleHandler construct
type AdminRoleHandler struct {
	roleService user.RoleService
}

//NewAdminRoleHadler returns new
func NewAdminRoleHadler(roleSrv user.RoleService) *AdminRoleHandler {
	return &AdminRoleHandler{roleService: roleSrv}
}

//GetRoles handles GET /v1/admin/roles/request
func (arh *AdminRoleHandler) GetRoles(w http.ResponseWriter,
	r *http.Request, _ httprouter.Params) {

	roles, errs := arh.roleService.Roles()

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(roles, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}

//GetSingleRoles handles GET /v1/admin/roles/request
func (arh *AdminRoleHandler) GetSingleRole(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

//PutRole handles PUT /v1/admin/roles/request
func (arh *AdminRoleHandler) PutRole(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

//DeleteRole handles DELETE /v1/admin/roles/:id request
func (arh *AdminRoleHandler) DeleteRole(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
