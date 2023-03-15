package repository

const (
    whiteList = "whiteList"
    blackList = "blackList"
)

func CheckWhiteList(ip string) (bool, error) {
    isInWhiteList, err := RedisClient.SIsMember(whiteList, string(ip)).Result()
    if err != nil {
        return false, err
    }

    return isInWhiteList, nil
}

func CheckBlackList(ip string) (bool, error) {
    isInBlackList, err := RedisClient.SIsMember(blackList, string(ip)).Result()
    if err != nil {
        return false, err
    }

    return isInBlackList, nil
}

func AddWhiteList(ip string) error {
    if err := RedisClient.SAdd(whiteList, string(ip)).Err(); err != nil {
        return err
    }
    return nil
}

func AddBlackList(ip string) error {
    if err := RedisClient.SAdd(blackList, string(ip)).Err(); err != nil {
        return err
    }
    return nil
}