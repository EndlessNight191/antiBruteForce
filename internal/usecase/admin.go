package usecase

import "test/internal/domain"

func (uc *UseCase) GetSetting(title string) (int) {
	return uc.setting.MaxLimitCommon
}

func (uc *UseCase) UpdateSetting(title string, value int) (int, error) {
	if err := uc.repo.UpdateSetting(title, value); err != nil {
		return value, err
	}

	// uc.setting

	return value, nil
}

func (uc *UseCase) AddIpToTheList(d domain.ListsActions) error {
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

func (uc *UseCase) RemoveIpToTheList(d domain.ListsActions) error {
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

func (uc *UseCase) ResetBucket(bucket domain.ResetBucket) error {

	return nil
}