package egg_drop

import (
	"math/big"
)

func Height(n, m *big.Int) *big.Int {
	if n.Int64() == 1 || m.Int64() == 1 {
		return m
	}

	var eggThrows = make([][]*big.Int, n.Int64())
	for i := range eggThrows {
		eggThrows[i] = make([]*big.Int, m.Int64())
	}

	for i := int64(0); i < m.Int64(); i++ {
		eggThrows[0][i] = big.NewInt(i + 1)
	}

	for i := int64(1); i < n.Int64(); i++ {
		eggThrows[i][0] = big.NewInt(1)
	}

	for i := int64(1); i < m.Int64(); i++ {
		for j := int64(1); j < n.Int64(); j++ {
			x := big.NewInt(0).Add(eggThrows[j-1][i-1], eggThrows[j][i-1])
			eggThrows[j][i] = x.Add(x, big.NewInt(1))
		}
	}

	return eggThrows[n.Int64()-1][m.Int64()-1]
}
