package main

import (
	"testing"

	aoc2022 "github.com/pineman/aoc2022"
)

func Test_partOne(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{aoc2022.GetInput(5)}, "ZWHVFWQWW"},
		{"", args{aoc2022.GetBigBoyInput(5)}, "QXWMZOHQIIZEXDCDVRDNYITZTKISAVCDLWNKVBQNGFDXXZKZRUOQAMKJOFOPFFTWQIIVMFOOSGTCLHPXNVRRUIBBPSHGFGULNFAUDSAVQDOMYTPVITJAPYJHZLXGEXCQUGXAPFUCPOZJUDLJSWEJPHWCDWKARGOPMKZRHVDUTHVAVXNUWOODGHEXBIXORGRWPTOPQHEW"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partOne(tt.args.input); got != tt.want {
				t.Errorf("partOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_partTwo(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{aoc2022.GetInput(5)}, "HZFZCCWWV"},
		{"", args{aoc2022.GetBigBoyInput(5)}, "MPRUXFCQFSPJULHGIRCZXCLTVKNUSSCDVWTWOUHSIEBAXFCRMUVZAMBDGLMPCAUXQAIVOXFCSPBTRIPBNKAUKIKBAVNKWKBBDDSIAQNXQJQTKLSNQXMJYIJXAHBEGSJWIAFADPGBECLDRJVRZCVKGHWVZMBAOGGGHAARNZOWPISKTVNUMKACYHXXACEMTGBTTWYPUWSD"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partTwo(tt.args.input); got != tt.want {
				t.Errorf("partTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
