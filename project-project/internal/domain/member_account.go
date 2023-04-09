package domain

import (
	"context"
	"fmt"
	"github.com/axzed/project-common/encrypts"
	"github.com/axzed/project-common/errs"
	"github.com/axzed/project-project/internal/dao"
	"github.com/axzed/project-project/internal/data"
	"github.com/axzed/project-project/internal/repo"
	"github.com/axzed/project-project/pkg/model"
	"time"
)

type AccountDomain struct {
	accountRepo      repo.AccountRepo
	userRpcDomain    *UserRpcDomain
	departmentDomain *DepartmentDomain
}

// AccountList 获取成员账号列表
func (d *AccountDomain) AccountList(
	organizationCode string,
	memberId int64,
	page int64,
	pageSize int64,
	departmentCode string,
	searchType int32) ([]*data.MemberAccountDisplay, int64, *errs.BError) {
	condition := ""
	organizationCodeId := encrypts.DecryptNoErr(organizationCode)
	departmentCodeId := encrypts.DecryptNoErr(departmentCode)
	switch searchType {
	case 1:
		condition = "status = 1"
	case 2:
		condition = "department_code = NULL"
	case 3:
		condition = "status = 0"
	case 4:
		condition = fmt.Sprintf("status = 1 and department_code = %d", departmentCodeId)
	default:
		condition = "status = 1"
	}
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	list, total, err := d.accountRepo.FindList(c, condition, organizationCodeId, departmentCodeId, page, pageSize)
	if err != nil {
		return nil, 0, model.ErrDBFail
	}
	var dList []*data.MemberAccountDisplay
	for _, v := range list {
		display := v.ToDisplay()
		memberInfo, _ := d.userRpcDomain.MemberInfo(c, v.MemberCode)
		display.Avatar = memberInfo.Avatar
		if v.DepartmentCode > 0 {
			department, err := d.departmentDomain.FindDepartmentById(v.DepartmentCode)
			if err != nil {
				return nil, 0, err
			}
			display.Departments = department.Name
		}
		dList = append(dList, display)
	}
	return dList, total, nil
}

// FindAccount 根据成员id获取账号信息
func (d *AccountDomain) FindAccount(memberId int64) (*data.MemberAccount, *errs.BError) {
	account, err := d.accountRepo.FindByMemberId(context.Background(), memberId)
	if err != nil {
		return nil, model.ErrAccountOrPwd
	}
	return account, nil
}

func NewAccountDomain() *AccountDomain {
	return &AccountDomain{
		accountRepo:      dao.NewMemberAccountDao(),
		userRpcDomain:    NewUserRpcDomain(),
		departmentDomain: NewDepartmentDomain(),
	}
}
