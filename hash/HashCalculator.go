package hash

type CityHash struct {

}

func (ch CityHash) CalcHash(a uint64, b uint64) uint64 {
   MAGIC_MULT := uint64(0x4906ba494954cb65)
   return MAGIC_MULT * (a + MAGIC_MULT * b);
}

func (ch CityHash) CalcCatFeatureHash(s string, hashes map[string]uint32, notPresent uint32) uint32{
	if hashes != nil {
		if v,ok := hashes[s]; ok {
			return v
		}
		return notPresent;
	}
	return notPresent;//new CityHash().cityHash64(s) & 0xffffffffl;
}