package controller

import (
	"context"
	"strconv"

	"github.com/shereifsrf/savespent-api/dao"
	"github.com/shereifsrf/savespent-api/util"
)

const (
	CountKey         = "count"
	SessionKeyFormat = "session_count:%d"
)

func setCount(ctx context.Context) (count int64, err error) {
	currentValue := dao.Get(ctx, CountKey)

	if currentValue != "" {
		count = util.GetInt64(currentValue)
	}
	count += 1

	err = dao.Set(ctx, CountKey, strconv.Itoa(int(count)))
	if err != nil {
		return 0, err
	}

	return count, nil
}
