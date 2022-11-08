package common

/**
 * @Author: QuRuijie
 * @Date: 2022/11/8 22:09
 * @Desc: Only the stars live up to the dark night
 */

func Test[T int | int32](i T) T {
	return i + 1
}
