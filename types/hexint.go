package types

import (
	"fmt"

	"github.com/go-zhouxun/xutil/hexutil"
)

type HexInt int64

func (hexInt HexInt) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, hexutil.EncodeUint64(uint64(hexInt)))), nil
}

func (hexInt *HexInt) UnmarshalJSON(data []byte) error {
	data = data[:len(data)-1][1:]
	i, err := hexutil.DecodeUint64(string(data))
	if err != nil {
		return err
	}
	*hexInt = HexInt(int64(i))
	return nil
}
