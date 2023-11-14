package handlers

//
//func (h *Handler) GetClaims(c *fiber.Ctx) (*token.CustomClaims, error) {
//
//	var claims token.CustomClaims
//
//	strToken := c.Get("Authorization")
//
//	token, err := jwt.Parse(strToken, func(t *jwt.Token) (interface{}, error) {
//		return []byte(h.jwt.SigninKey), nil
//	})
//
//	if err != nil {
//		h.log.Error("invalid access token")
//		return nil, err
//	}
//	rawClaims := token.Claims.(jwt.MapClaims)
//
//	claims.Sub = rawClaims["sub"].(string)
//	claims.Role = rawClaims["role"].(string)
//	claims.Exp = rawClaims["exp"].(float64)
//	fmt.Printf("%T type of value in map %v\n", rawClaims["exp"], rawClaims["exp"])
//	fmt.Printf("%T type of value in map %v\n", rawClaims["iat"], rawClaims["iat"])
//
//	claims.Iat = rawClaims["iat"].(float64)
//
//	var aud = make([]string, len(rawClaims["aud"].([]interface{})))
//
//	for i, v := range rawClaims["aud"].([]interface{}) {
//		aud[i] = v.(string)
//	}
//
//	claims.Aud = aud
//	claims.Iss = rawClaims["iss"].(string)
//
//	return &claims, nil
//
//}
