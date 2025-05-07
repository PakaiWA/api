// Copyright (c) 2025 KAnggara75
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v.2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// @author KAnggara75 on Thu 01/05/25 13.55
// @project api https://github.com/PakaiWA/api/tree/main/api
//

package api

import (
	"fmt"

	"github.com/pakaiwa/api/helper"
	"github.com/pakaiwa/api/model/entity"
)

type UserRs struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func ToUserResponse(user entity.User) UserRs {
	fmt.Println("Invoke ToUserResponse")

	token, err := helper.GenerateJWT(user)
	if err != nil {
		panic(err.Error())
	}

	return UserRs{
		Id:    user.Uuid,
		Email: user.Email,
		Token: token,
	}

}
