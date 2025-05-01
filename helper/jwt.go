// Copyright (c) 2025 KAnggara75
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v.2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// @author KAnggara75 on Thu 01/05/25 15.44
// @project api helper
//

package helper

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pakaiwa/api/config"
	"github.com/pakaiwa/api/model/entity"
	"time"
)

func GenerateJWT(user entity.User) (string, error) {

	jwtKey := config.GetJWTKey()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"iat":   time.Now().Unix(),
	})

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return tokenString, nil
}
