package casbindemo

import (
	"fmt"
	"github.com/casbin/casbin"
	"github.com/casbin/gorm-adapter"
	"log"
)

// 连接数据库
func Demo(dbtype, dburl string) {
	adapter := gormadapter.NewAdapter(dbtype, dburl, true)
	e := casbin.NewEnforcer("casbindemo/examples/rbac_model.conf", adapter)

	e.LoadPolicy()

	e.AddPolicy("alice", "data1", "read")

	isok := e.Enforce("alice", "data1", "read")
	fmt.Println("isok:", isok)

	//e.SavePolicy()

	subs := e.GetAllSubjects()
	println("len(subs):", len(subs))
	println("subs:", subs)
}

// 需要注意的地方是：策略文件需要在,（逗号）后面加一个空格
//p, superAdmin, gy, project, read
//p, superAdmin, gy, project, write
//p, superAdmin, jn, project, read
//p, superAdmin, jn, project, write
//p, admin, gy, project, read
//p, admin, gy, project, write
//p, admin, jn, asse, read
//p, admin, jn, asse, write
//p, zhuangjia, jn, project, write
//p, zhuangjia, gy, asse, write
//
//g, quyuan, admin, gy
//g, quyuan, admin, jn
//g, wenyin, zhuangjia, gy
//g, shangshang, zhuangjia, jn

// 本地配置文件 不带domain
func Demo2() {
	e := casbin.NewEnforcer("casbindemo/examples/rbac_model.conf", "casbindemo/examples/rbac_model.csv")
	fmt.Println("RBAC test start\n")

	if e.Enforce("superAdmin", "project", "read") {
		log.Println("superAdmin can read project")
	} else {
		log.Fatal("ERROR:superAdmin can not read project")
	}

	permissions := e.GetPermissionsForUser("superAdmin")
	fmt.Println("permission:", permissions)

	users := e.GetUsersForRole("admin")
	fmt.Println("users:", users)

	roles := e.GetRolesForUser("wenyin")
	fmt.Println("roles:", roles)

	if e.AddRoleForUser("ck", "admin") {
		aaa := e.GetRolesForUser("ck")
		fmt.Println("aaa:", aaa)
	}

}

// 本地配置文件 带domain
func Demo3() {
	e := casbin.NewEnforcer("casbindemo/examples/rbac_model_domain.conf", "casbindemo/examples/rbac_model_domain.csv")

	// superAdmin
	if e.Enforce("superAdmin", "gy", "project", "read") {
		log.Println("superAdmin can read project in gy")
	} else {
		log.Fatal("ERROR: superAdmin can not read project in gy")
	}

	//data := e.GetFilteredGroupingPolicy(0, "quyuan", "admin", "gy")
	data := e.GetFilteredGroupingPolicy(1, "admin", "gy")
	res := e.GetFilteredPolicy(0, "superAdmin", "gy", "project", "read") //superAdmin, gy, project, read
	//res = e.GetFilteredPolicy(1, "gy", "project", "read")
	//res = e.GetFilteredPolicy(2, "project", "read")
	//res = e.GetFilteredPolicy(3, "read")

	fmt.Println("data:", data)
	fmt.Println("res:", res)

	res = e.GetFilteredPolicy(0, "admin", "gy")

	var result []string
	for _, d := range res {
		resource, action := d[2], d[3]
		result = append(result, resource+":"+action)
	}

	fmt.Println("result:", result)

	e.AddPolicy("admin", "gy", "users", "create")

	res = e.GetFilteredPolicy(0, "admin", "gy")

	fmt.Println("res:", res)
}
