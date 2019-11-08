package util

import (
	"fmt"
	"time"

	"infinite-window-micro/constant"
	"infinite-window-micro/pkg/setting"
	"math"
	"math/rand"
)

// Setup Initialize the util
func Setup() {
	jwtSecret = []byte(setting.AppSetting.JwtSecret)
}

//下个自然月开始时间戳
func GetStartTimestampNextMonth() int64 {

	curTime := time.Now()
	dstYear := curTime.Year()
	dstMonth := curTime.Month() + 1
	if dstMonth > time.December {
		dstYear += 1
		dstMonth = time.January
	}

	dstTimeStr := fmt.Sprintf("%04d-%02d-01", dstYear, dstMonth)

	timeObj := ParseDateStringToTimeObj(dstTimeStr)
	if timeObj != nil {
		return timeObj.Unix()
	} else {
		return 0
	}

}

//转换日期字符串为时间对象
func ParseDateStringToTimeObj(timeString string) *time.Time {

	if timeString == "" {
		return nil
	}
	local, _ := time.LoadLocation("Local")
	timeObj, err := time.ParseInLocation(constant.TIME_TO_STRING_DAY, timeString, local)
	if err != nil {
		return nil
	} else {
		return &timeObj
	}

}

//GetHidePhoneNumber 隐藏电话号码
func GetHidePhoneNumber(phone string) string {
	if len(phone) <= 10 {
		return phone
	}
	return phone[:3] + "****" + phone[len(phone)-4:]
}

//RandomStr 随机生成字符串
func RandomStr(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

//某天之后的起始时间
func GetDayStartTime(days int) time.Time {

	year, month, day := time.Now().Date()
	toDayStr := fmt.Sprintf("%04d-%02d-%02d", year, month, day)
	toDayTimeObj := ParseDateStringToTimeObj(toDayStr)

	newTimeObj := toDayTimeObj.Add(time.Hour * 24 * time.Duration(days))

	return newTimeObj
}

//当前自然月开始时间戳
func GetCurrentMonthStartTime() *time.Time {

	curTime := time.Now()
	dstYear := curTime.Year()
	dstMonth := curTime.Month()

	dstTimeStr := fmt.Sprintf("%04d-%02d-01", dstYear, dstMonth)

	return ParseDateStringToTimeObj(dstTimeStr)
}

// Round float取精度
func Round(f float64, n int) float64 {
	n10 := math.Pow10(n)
	return math.Trunc((f+0.5/n10)*n10) / n10
}

// Round2 float取精度
func Round2(f float64) float64 {
	return Round(f, 2)
}

// RemoveRepByMap list 去重
func RemoveRepByMap(slc []uint) []uint {
	result := []uint{}
	tempMap := map[uint]byte{} // 存放不重复主键
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, e)
		}
	}
	return result
}
