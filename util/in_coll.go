package util

type CacheTable map[*[]byte]bool
type Cache map[byte]CacheTable

var cache Cache = make(Cache)

func InColl(val byte, coll *[]byte) bool {
	result := false

	valCacheTable, valInCache := cache[val]
	if !valInCache {
		cache[val] = make(CacheTable)
		valCacheTable = cache[val]
	}

	cachedResult, resultInCache := valCacheTable[coll]
	if resultInCache {
		return cachedResult
	}

	for _, cval := range *coll {
		result = val == cval
		if result {
			break
		}
	}

	valCacheTable[coll] = result
	return result
}
