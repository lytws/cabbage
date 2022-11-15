package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"cbg/adapters/ahttp"
	"cbg/adapters/gwimpl"
	"cbg/driver"
	"cbg/usecase/interactor"
	"cbg/usecase/repoimpl"
)

func TestAuthByPassword(t *testing.T) {
	// drivers
	dUserdb := driver.NewMockUserDB()
	dAuthInfo := driver.NewMockAuthInfo()
	dHttp := driver.NewMockHttpFramework()
	// presenters
	pAuthHttp := ahttp.NewAuthJsonPresenter(dHttp)
	// gateways
	gwUser := gwimpl.NewMockUser(dUserdb)
	gwAuthInfo := gwimpl.NewMockAuthInfo(dAuthInfo)
	// repos
	repoUser := repoimpl.NewUser(gwUser)
	// interactors
	iAuth := interactor.NewAuth(repoUser, gwAuthInfo, pAuthHttp)
	// controllers
	cAuthHttp := ahttp.NewAuthController(iAuth, dHttp)

	// set mock body data
	mockUserData := dUserdb.GetRandomUser()
	data, _ := json.Marshal(mockUserData)
	ctx := dHttp.ContextWithBody(context.Background(), data)

	// invoke api
	cAuthHttp.AuthByPassword(ctx)
}

func TestAuthByAuthCode(t *testing.T) {
	// drivers
	dUserdb := driver.NewMockUserDB()
	dAuthInfo := driver.NewMockAuthInfo()
	dHttp := driver.NewMockHttpFramework()
	// presenters
	pAuthHttp := ahttp.NewAuthJsonPresenter(dHttp)
	// gateways
	gwUser := gwimpl.NewMockUser(dUserdb)
	gwAuthInfo := gwimpl.NewMockAuthInfo(dAuthInfo)
	// repos
	repoUser := repoimpl.NewUser(gwUser)
	// interactors
	iAuth := interactor.NewAuth(repoUser, gwAuthInfo, pAuthHttp)
	// controllers
	cAuthHttp := ahttp.NewAuthController(iAuth, dHttp)

	// set mock body data
	mockUserData := dUserdb.GetRandomUser()
	authCode := fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
	ctx := dAuthInfo.ContextWithUsername(context.Background(), authCode, mockUserData.Username)
	ctx = dHttp.ContextWithUriParam(ctx, "authCode", authCode)

	// invoke api
	cAuthHttp.AuthByAuthCode(ctx)
}
