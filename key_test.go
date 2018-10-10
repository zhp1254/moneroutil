package moneroutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeyDerivation(t *testing.T) {
	tests := []struct {
		a, R, kd string
	}{
		{"c702e21a00a3e0d644c91efb350752d27824ea5d8b1649e6fb1b09a4ef833600", "0aa53fcdc6b6b081245d7878ead91cfca92e5d11ac494160a28847bda1518d2f", "ff2159e2c7257853a50a3c18480cbab251e3049b479853bbbf1622dbdf78b2db"},
		{"9f45b178b9d420456e2c985d6f5bcd5633c7908091ba5106d399d5f0520a9d0c", "7b323079b2a90d60cd8dce9671c625ef9db41dea774a52fee1bf60378b2bc4da", "6b3128852ef3285a0b95cf024b271e49a7b0b68382bc3457f37f31564b11a034"},
		{"fac2722e078d47d9bb7131b70c44e39fed54fd86d4d7d0da3a156333ecd96607", "629acdddea2cb95e5a80fc4a7168642b35782c799bd2123ba443fa2f1ce9fd26", "054934acbc6c5a866e22aeb1e83cc81da32f7d455c9b8d030bd2d96e2b0dcfb7"},
		{"ed9d81382a7552dcf0239bc9bc6badf854a58b85a7b14aa38a15e50073888e05", "2f24ad406b72ea0826433df264295b41c48e3f2b13ec8de121f9ef8f92b147e5", "d547031db8ed95a68734291124a824e5c6b76be80c51ef856fe3dc67a4ca622a"},
		{"2647501173cf342b46fff6e24372b2e252fd3a985733ab786904c8b2fd426d0d", "d48eead884b1725524f11ee7e217214dccd98c533b70df22127b51ca53f78cc3", "607e4e23da51c9fa5a2bf7c97d0d2fbc915d36430f2d493bef6cc456c2134910"},
		{"608961996556efc12e021092644b03a313bda413805158a0171578218b89a409", "8e482b843e3a70eb069bf00711c514877d00047ae302a6e21dc5f2172d47fe26", "4cde683ba23303beefdb55573c9839c0574e2450fd9055af792210d28609c535"},
		{"fcc2fc91060bf28bc20ffd27273c52be8c1046eda286ca43254ab4a60b58cd03", "fbd4c90adf8650fe28c713f2f5a8f8ef8a7a4554e9a6be0d108a022f2b3d595a", "20fb9590cb1932fe92658a146645482656c4745068e3c2b0251e2cc25cf07fd2"},
		{"fbe340c23e93c888e58b03fbf5a25b8a82856b847442458e09a9e81397cd3907", "7a303b9e840b6e6dca4d52cadfa36ad19a1258d5ed8cd04025ff5e7f627fb06b", "df490397f974d250582ccab7888ae53a0051735721c7545a7bc1f15b22596c3b"},
		{"517461b60bf3493e7a4fc5d6eabf73cfa28d4874ad06fbc3a28de7d1f269f202", "853f96b792ad7cd8acdb01453a0bf377d829f7849b645f2c328c654934bde9ce", "e164c242a37b7e9d561bea91d1a1bae3de7b55b54b898829127c8edfe79963a9"},
		{"2ce953b24c4df96b9f427020c14dd8c5fb26dc3675db8b1dfbe37bbb93c94c00", "7fa291ba3ba44b514deec366d54039211e8aacc87e9b8bbfb530d9a46b659a7a", "ab019d33792a151854d2cd5e4ae90a6d4cecea7125318c21110a7c73fc0c1c63"},
	}

	for _, test := range tests {
		a, _ := HexToKey(test.a)
		R, _ := HexToKey(test.R)
		expected, _ := HexToKey(test.kd)

		assert.Equal(t, expected, KeyDerivation(&a, &R))
	}
}