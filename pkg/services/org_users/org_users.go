package orgusers

import (
	"context"

	"github.com/grafana/grafana/pkg/models"
	orgusersstore "github.com/grafana/grafana/pkg/services/org_users/store"
	"github.com/grafana/grafana/pkg/services/sqlstore"
	"github.com/grafana/grafana/pkg/setting"
)

type Manager interface {
	AddOrgUser(ctx context.Context, cmd *models.AddOrgUserCommand) error
	// UpdateOrgUser(ctx context.Context, cmd *models.UpdateOrgUserCommand) (*models.OrgUser, error)

	GetOrgUsers(ctx context.Context, query *models.GetOrgUsersQuery) error
	// SearchOrgUsers(ctx context.Context, query *models.SearchOrgUsersQuery) error
	RemoveOrgUser(ctx context.Context, cmd *models.RemoveOrgUserCommand) error
}

type ManagerImpl struct {
	store orgusersstore.Store
}

func ProvideService(sqlStore *sqlstore.SQLStore, cfg *setting.Cfg) Manager {
	s := ManagerImpl{store: orgusersstore.NewStore(sqlStore, cfg)}
	return &s
}

func (m *ManagerImpl) AddOrgUser(ctx context.Context, cmd *models.AddOrgUserCommand) error {
	return m.store.AddOrgUser(ctx, cmd)
}

func (m *ManagerImpl) RemoveOrgUser(ctx context.Context, cmd *models.RemoveOrgUserCommand) error {
	return m.store.RemoveOrgUser(ctx, cmd)
}

func (m *ManagerImpl) GetOrgUsers(ctx context.Context, query *models.GetOrgUsersQuery) error {
	return m.store.GetOrgUsers(ctx, query)
}

// 	user, err := m.store.GetUser(ctx, cmd)
// 	if err != nil {
// 		return nil, err
// 	}
// 	cmd.UserId = user.Id

// 	orgUser, err := m.store.Get(ctx, cmd.OrgId, cmd.UserId)
// 	if err != nil {
// 		return nil, err
// 	}
// 	err = m.store.OrgExists(ctx, cmd)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return m.store.Add(ctx, orgUser)
// }

// func (m *ManagerImpl) UpdateOrgUser(ctx context.Context, cmd *models.UpdateOrgUserCommand) (*models.OrgUser, error) {
// 	_, err := m.store.Get(ctx, cmd.OrgId, cmd.UserId)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return m.store.Update(ctx, cmd)
// }

// func (m *ManagerImpl) SearchOrgUsers(ctx context.Context, query *models.SearchOrgUsersQuery) error {
// 	return m.store.SearchOrgUsers(ctx, query)
// }
