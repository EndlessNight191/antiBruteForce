package repository

const (
    whiteList = "whiteList"
    blackList = "blackList"
)

func (r *ClientRepository) CheckWhiteList(ip string) (bool, error) {
    isInWhiteList, err := r.redisClient.SIsMember(whiteList, string(ip)).Result()
    if err != nil {
        return false, err
    }

    return isInWhiteList, nil
}

func (r *ClientRepository) CheckBlackList(ip string) (bool, error) {
    isInBlackList, err := r.redisClient.SIsMember(blackList, string(ip)).Result()
    if err != nil {
        return false, err
    }

    return isInBlackList, nil
}

func (r *ClientRepository) AddWhiteList(ip string) error {
    if err := r.redisClient.SAdd(whiteList, string(ip)).Err(); err != nil {
        return err
    }
    return nil
}

func (r *ClientRepository) AddBlackList(ip string) error {
    if err := r.redisClient.SAdd(blackList, string(ip)).Err(); err != nil {
        return err
    }
    return nil
}

func (r *ClientRepository) RemoveFromWhiteList(ip string) error {
    if err := r.redisClient.SRem(whiteList, string(ip)).Err(); err != nil {
        return err
    }
    return nil
}

func (r *ClientRepository) RemoveFromBlackList(ip string) error {
    if err := r.redisClient.SRem(blackList, string(ip)).Err(); err != nil {
        return err
    }
    return nil
}