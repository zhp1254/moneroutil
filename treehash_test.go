package moneroutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTreeHashCnt(t *testing.T) {
	cases := []struct {
		input, expected int
	}{
		{3, 2},
		{4, 2},
		{5, 4},
		{7, 4},
		{9, 8},
		{15, 8},
		{16, 8},
		{17, 16},
		{256, 128},
		{257, 256},
	}

	for _, testCase := range cases {
		assert.Equal(t, testCase.expected, TreeHashCnt(testCase.input))
	}
}

func HexHashesToHashes(hexes []string) []Hash {
	res := make([]Hash, len(hexes))
	for i, h := range hexes {
		res[i], _ = HexToHash(h)
	}

	return res
}

func TestTreeHash(t *testing.T) {
	originalHashes := []string{
		"3e8781dfcd3089ae47ef2f9c6dd9f3056aed2ebc4771481a37c5571535927590",
		"b60cf705578fdbeec876d1bb79efd6a69ca4fc6dcf49fc24d410cfaa2b5ddd1c",
		"07b3c39cf77dd1274a0ec15b7ae76e621d25d8af29cb98aa9630014172420c39",
		"15489fd734b53269e220ad2f61fb47d923b6113389c3c15f3fb058a7121108f0",
		"fbf8f2ff9cc7bdc7a86a62e3c85ff5ef2330dfae737876d2d9301239727741a3",
		"ab30b18235f8a3a73de4d8a2cbe25b95f2df0adb12d9ed0dcbe81d9e3d003ce3",
		"dcbd665b4d7e96b0afa6b390c3dd787081817cec1ee057e76287271a932bc22a",
		"90492c3f3c2daf812afe176237e1d3a2331a94781f120f3fa1121cd6cd7d9bd9",
		"72d5057b49c55b20fefdabf4acdcb21cc5ed7d9f27e018250fd3b089efaa3524",
		"f010183556bcd1fbb6b3f94d8f632e24c04dfeb64160b86f9b7315422d4ba583",
		"36ec9cc0fe6b134bdd72c24214725002f8fd46babc00693d7f0197e351d4c93b",
		"f344bfbc98187f294baae2ff65a4b67c67d07b0401a67dd9b6795426082f403b",
		"246a572f370618980de46a79290f8e8ef70cdbed71f05b23a94ad5e504f080d9",
		"7a55ec43f8aef163e36427462bd23ea73e27e6522cb14c902615750bb8c6c278",
		"c06a8620d4a8fbbc2c84ab24b637b5bb5b264092f8d054296a24a1f04ef0329c",
		"2262013d1be9d1638df54781b43ec82486cfb1cadc591e064040247be36cf548",
		"da389f198568ecc86485cfefbd56bee575e9c8718da6f9bfa7b12fd3d801ff85",
		"f5123c3b731366dcee83fa45c6afe8586eedea0bc4680e87505b295df04c4ff3",
		"c599dcb1daa7c8e0349dce9f90e4cfb40c5d771d33f7b0e7af8a8252bfe69877",
		"eb00b1148e5ff52d41c762ac6ddcf7157a433ebb7e27c48554d663313e0a458a",
		"d5866f489bdd3b013be2eae917c9a2f50e77a8279ef2da3903e420ed6a047cc1",
		"f0ccd2ba067c340db35023f4bf4c5de854b3a5d5f9d117917d728f9147c97e0c",
		"de02fc3a2f1a7e2116e357efd31500ac888e97e3e40f381cdd26c6236fcc729a",
		"ece13410dacfa3b15fb37a70bac32b22823b89dd0abd260ae422e52d13c449c7",
		"78d3ac0536b3f08233af90757d359a2c5f31a637448d71e165ff98eb4c033b45",
		"d2199060bc1b2b151f5aff58cd02b1d5c2388fca44bee4d6acb2bbec386a318f",
	}

	cases := []struct {
		count    int
		expected string
	}{
		{1, "3e8781dfcd3089ae47ef2f9c6dd9f3056aed2ebc4771481a37c5571535927590"},
		{2, "bb44efee25ab3c649e1aaaf42180a8e41b9abb02e1dc9b7d125525fd946e34b9"},
		{3, "7d7e90ad697882b92b89cd53fbd85ee8681b341482089480d35a9d6773eca971"},
		{4, "4a5fb7194da721fb1a1c3f98c44e52bbd88f6176ac4e00bbbb926a3ac21cf317"},
		{5, "56d3608a478272fc3c71e93f17de74d2ae916e31192bc168391885fe5d65f8fe"},
		{6, "4148e07fee3df9a0e0b098292a1c67dcbcdfffabb54af97cb0ebe1a0eab7928b"},
		{7, "5a4f4db5444addb06bf9efad01e038d9815061bc98fffdc95843eb69a4ea9234"},
		{8, "a2686b36bbce51a5b38a2ce2fabf6547b6a1062287edda7f4ce7ae058f2e12c9"},
		{9, "c027bb9efd0c5b522bb1c0f622836fbd4e4425f312d9296e9de002c67c6bafe1"},
		{10, "f3d960ac24714c828f053110cf023698d209f649de1fe72a40458753d17e52af"},
		{11, "194b38a751d7aee16890166b210e83b4be8b406d326381d4010a1fcca34c7d2c"},
		{12, "5f12d959bb02df30b65b532ab6c5c252293447dab912f5a004678a221169c09c"},
		{13, "950f29690a02f432dae5e577b7e0889106b3a4215dc8eb1f97d7885853158406"},
		{14, "2cb00df6c4adf6facfefdc635433b85e45f7a44188fdc3382a74e12044d86481"},
		{15, "db538a846b5226db2dd6999a19429cf73914b5f5b09a0dd1f4b8f6d63e30694b"},
		{16, "218e83666ee47a7c315966a5864ed172c69d54b32d615372d6bcacbfadca2a45"},
		{17, "af2c1ca6dd0b0f6f472320bc26f71b45d8fe606ffdc80d9fbe987d3e237ddc6a"},
		{18, "99d841df0f0f71b82a62743a23bb44e9fd38cd009b0b881e58499b83fa808305"},
		{19, "46b5f25226cb594b3915bf136558e5d0b5640348d82f0567feab2d81bd02e557"},
		{20, "aebf9523b409354b994246af0ee48db6e9162b3822205eb5802b42f3c14d46b4"},
		{21, "a0c4d654b079082c2a0b31f22b8d3bc746ff593da1b98f70f0f4a77bc1a7e138"},
		{22, "86dbf739d44a7c505440b5ae170e3a30fb5c5bf05934c6dd6f4ae9bfb705bc27"},
		{23, "f0f2e5a77a648d9acffa7d40f47a1d1aa12f57066fa2d6719890f2b71e8462aa"},
		{24, "176a792e4374ec355575b67ee9ae17385eeca436beacf7ad5aac611dd58093cb"},
		{25, "af73671becdde78e608ea6d31bbb0cf1db6c2c81527852d694e45d7be945b5d1"},
		{26, "2e52d212c54e10ee7a979dd4a002ebd4ab9a7441a721c2dca3809cca4d59a7dd"},
	}

	for i, test := range cases {
		h := HexHashesToHashes(originalHashes[0:test.count])
		actual := TreeHash(h)
		expected, _ := HexToHash(test.expected)
		assert.Equal(t, expected, actual, "for count: %d (index: %d)", test.count, i)
	}
}
