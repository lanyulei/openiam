package comparemaps

import "reflect"

/*
  @Author : lanyulei
  @Desc :
*/

// CompareMaps 比较两个多层嵌套的 map
func CompareMaps(m1, m2 map[string]interface{}) bool {
	if len(m1) != len(m2) {
		return false
	}

	for k, v1 := range m1 {
		v2, ok := m2[k]
		if !ok {
			return false
		}

		if !compareValues(v1, v2) {
			return false
		}
	}
	return true
}

// compareValues 比较两个 interface{} 类型的值
func compareValues(v1, v2 interface{}) bool {
	switch v1Typed := v1.(type) {
	case map[string]interface{}:
		v2Typed, ok := v2.(map[string]interface{})
		if !ok || !CompareMaps(v1Typed, v2Typed) {
			return false
		}
	case []interface{}:
		v2Typed, ok := v2.([]interface{})
		if !ok || !compareSlices(v1Typed, v2Typed) {
			return false
		}
	default:
		if !reflect.DeepEqual(v1, v2) {
			return false
		}
	}
	return true
}

// compareSlices 比较两个切片
func compareSlices(s1, s2 []interface{}) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i, v1 := range s1 {
		v2 := s2[i]
		if !compareValues(v1, v2) {
			return false
		}
	}
	return true
}
