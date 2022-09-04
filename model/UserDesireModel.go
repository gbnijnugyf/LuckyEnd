package model

// 查询用户的投递愿望
func GetUserCreateDesire(UserID *int) ([]*Desire, error) {
	var desires []*Desire
	err := db.Model(&Desire{}).Where("user_id = ?", *UserID).Find(&desires).Error
	return desires, err
}

// 查询用户的点亮愿望
func GetUserLightDesire(UserID *int) ([]*Desire, error) {
	var lights []*Desire
	err := db.Model(&Desire{}).Where("light_id = ?", *UserID).Find(&lights).Error
	return lights, err
}
