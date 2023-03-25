package usecase

import "test/internal/domain"

func (uc UseCase) GetSettings() (domain.ConfigSetting, error) {
	maxLimitCommon, err := uc.repo.GetLimitSettingInt(domain.MaxLimitCommon)
	if err != nil {
		return domain.ConfigSetting{}, err
	}

	maxLimitIp, err := uc.repo.GetLimitSettingInt(domain.MaxLimitIp)
	if err != nil {
		return domain.ConfigSetting{}, err
	}

	maxLimitLogin, err := uc.repo.GetLimitSettingInt(domain.MaxLimitLogin)
	if err != nil {
		return domain.ConfigSetting{}, err
	}

	maxLimitPassword, err := uc.repo.GetLimitSettingInt(domain.MaxLimitPassword)
	if err != nil {
		return domain.ConfigSetting{}, err
	}

	expairBacket, err := uc.repo.GetLimitSettingInt(domain.ExpairBacket)
	if err != nil {
		return domain.ConfigSetting{}, err
	}

	setting := domain.ConfigSetting{
		MaxLimitCommon: maxLimitCommon,
		MaxLimitIp: maxLimitIp,       
		MaxLimitLogin: maxLimitLogin,   
		MaxLimitPassword: maxLimitPassword,
		ExpairBacket: expairBacket,
	}

	return setting, nil
}

func (uc *UseCase) UpdateSettings(settingUpdate domain.ConfigSetting) error {
	if settingUpdate.ExpairBacket != 0 {
		uc.setting.ExpairBacket = settingUpdate.ExpairBacket
		if err := uc.repo.UpdateSetting(domain.ExpairBacket, settingUpdate.ExpairBacket); err != nil {
			return err
		}
	}

	if settingUpdate.MaxLimitCommon != 0 {
		uc.setting.MaxLimitCommon = settingUpdate.MaxLimitCommon
		if err := uc.repo.UpdateSetting(domain.MaxLimitCommon, settingUpdate.MaxLimitCommon); err != nil {
			return err
		}
	}

	if settingUpdate.MaxLimitIp != 0 {
		uc.setting.MaxLimitIp = settingUpdate.MaxLimitIp
		if err := uc.repo.UpdateSetting(domain.MaxLimitIp, settingUpdate.MaxLimitIp); err != nil {
			return err
		}
	}

	if settingUpdate.MaxLimitLogin != 0 {
		uc.setting.MaxLimitLogin = settingUpdate.MaxLimitLogin
		if err := uc.repo.UpdateSetting(domain.MaxLimitLogin, settingUpdate.MaxLimitLogin); err != nil {
			return err
		}
	}

	if settingUpdate.MaxLimitPassword != 0 {
		uc.setting.MaxLimitPassword = settingUpdate.MaxLimitPassword
		if err := uc.repo.UpdateSetting(domain.MaxLimitPassword, settingUpdate.MaxLimitPassword); err != nil {
			return err
		}
	}

	return nil
}

func (uc UseCase) AddIpToTheList(d domain.ListsActions) error {
	if d.TitleList == domain.ListBlack {
		if err := uc.repo.AddBlackList(d.IP); err != nil {
			return err
		}

		return nil
	}

	if d.TitleList == domain.ListWhite {
		if err := uc.repo.AddWhiteList(d.IP); err != nil {
			return err
		}

		return nil
	}

	return nil
}

func (uc UseCase) RemoveIpToTheList(d domain.ListsActions) error {
	if d.TitleList == domain.ListBlack {
		if err := uc.repo.RemoveFromBlackList(d.IP); err != nil {
			return err
		}

		return nil
	}

	if d.TitleList == domain.ListWhite {
		if err := uc.repo.RemoveFromWhiteList(d.IP); err != nil {
			return err
		}

		return nil
	}

	return nil
}

func (uc UseCase) ResetBucket(resetBucket domain.ResetBucket) error {
	if resetBucket.IP != "" {
		ipKey := joinToFormatIp(resetBucket.IP)
		if err := uc.repo.DeleteByKey(ipKey); err != nil {
			return err
		}
	}

	if resetBucket.Login != "" {
		loginKey := joinToFormatLogin(resetBucket.Login)
		if err := uc.repo.DeleteByKey(loginKey); err != nil {
			return err
		}
	} 

	return nil
}