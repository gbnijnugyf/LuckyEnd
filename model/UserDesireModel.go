package model

func GetUserCreateDesire(UserID *int) ([]*Desire, error) {
	var desires []*Desire
	err := db.Model(&Desire{}).Where("user_id = ?", *UserID).Find(&desires).Error
	return desires, err
}

func GetUserLightDesire(UserID *int) ([]*Desire, error) {
	var lights []*Desire
	err := db.Model(&Desire{}).Where("light_id = ?", *UserID).Find(&lights).Error
	return lights, err
}
