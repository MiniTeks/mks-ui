package db

import "github.com/go-redis/redis/v8"

func GetValue(rClient *redis.Client, key string) string {
	val, err := rClient.Get(ctx, key).Result()
	if err != nil {
		return "0"
	}
	return val
}
func GetValues(rClient *redis.Client) (*Application, error) {
	dataPoint := []string{"created", "deleted", "active", "failed", "completed"}
	resources := []string{"mksTask", "mksTaskRun", "mksPipelineRun"}
	ans := [3]Mksresource{}
	for j, r := range resources {
		resArr := [5]string{}
		for i, d := range dataPoint {
			val := GetValue(rClient, r+d)
			resArr[i] = val
		}
		ans[j] = Mksresource{
			Created:   resArr[0],
			Deleted:   resArr[1],
			Active:    resArr[2],
			Failed:    resArr[3],
			Completed: resArr[4],
		}
	}
	res := &Application{
		Mkstask:        ans[0],
		Mkstaskrun:     ans[1],
		Mkspipelinerun: ans[2],
	}
	return res, nil
}
