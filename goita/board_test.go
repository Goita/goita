package goita

import (
	"reflect"
	"strings"
	"testing"
)

func TestParseBoard(t *testing.T) {
	type args struct {
		historyString string
	}
	tests := []struct {
		name string
		args args
		// want is input historyString
	}{
		// TODO: Add test cases.
		{"initial", args{"12345678,12345679,11112345,11112345,s1"}},
		{"end of deal", args{"22221678,11111345,11345679,11345345,s1,112,2p,3p,4p,162,2p,3p,4p,172,2p,3p,4p,128"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseBoard(tt.args.historyString).String(); !reflect.DeepEqual(got, tt.args.historyString) {
				t.Errorf("ParseBoard() = %v, want %v", got, tt.args.historyString)
			}
		})
	}
}

func TestBoard_GetPossibleMoves(t *testing.T) {
	tests := []struct {
		name  string
		board string
		want  string
	}{
		// TODO: Add test cases.
		{"max moves", "12345678,12345679,11112345,11112345,s1", "12,13,14,15,16,17,21,23,24,25,26,27,31,32,34,35,36,37,41,42,43,45,46,47,51,52,53,54,56,57,61,62,63,64,65,67,71,72,73,74,75,76,81,82,83,84,85,86,87"},
		{"gon-ou finish", "12345678,12345679,11112345,11112345,s1,113,2p,3p,431,1p,2p,315,4p,156,267,3p,4p,174,242,3p,4p", "p,28"},
		{"end of deal", "12345678,12345679,11112345,11112345,s1,113,2p,3p,431,1p,2p,315,4p,156,267,3p,4p,174,242,3p,4p,128", ""},
		{"king's double-up finish", "12667789,12345543,11112345,11112345,s1,116,2p,3p,4p,126,2p,3p,4p,177,2p,3p,4p", "88"},
		{"finish with yaku", "22235567,12345679,11133448,11111145,s1", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := ParseBoard(tt.board)
			got := b.GetPossibleMoves()
			ret := make([]string, 0)
			for _, v := range got {
				ret = append(ret, v.String())
			}
			moves := strings.Join(ret, ",")
			if len(moves) != len(tt.want) {
				t.Errorf("Board.GetPossibleMoves() = %v, want %v", moves, tt.want)
			}
		})
	}
}

// go test ./goita -bench PossibleMoves -benchmem -benchtime 1s
func Benchmark_PossibleMoves(b *testing.B) {
	buf := make(KomaArray, FieldLength)
	moves := make([]*Move, 0, 64)
	board1 := ParseBoard("12345678,12345679,11112345,11112345,s1")
	board2 := ParseBoard("11244556,12234569,11123378,11113457,s3,371,411,115,2p,3p,4p,145,252,3p,4p,124,2p")

	for i := 0; i < b.N; i++ {
		board1.PossibleMoves(buf, moves)
		board2.PossibleMoves(buf, moves)
	}
}

func TestBoard_Score(t *testing.T) {
	tests := []struct {
		name  string
		board string
		want  int
	}{
		{"not finished", "12345678,12345679,11112345,11112345,s1,113,2p,3p,431,1p,2p,315,4p,156,267,3p,4p,174,242,3p,4p", 0},
		{"end of deal", "12345678,12345679,11112345,11112345,s1,113,2p,3p,431,1p,2p,315,4p,156,267,3p,4p,174,242,3p,4p,128", 50},
		{"king's double-up finish", "12667789,12345543,11112345,11112345,s1,116,2p,3p,4p,126,2p,3p,4p,177,2p,3p,4p,188", 100},
		//{"finish with yaku", "22235567,12345679,11133448,11111145,s1", 30},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := ParseBoard(tt.board)
			if got := b.Score(); got != tt.want {
				t.Errorf("Board.Score() = %v, want %v", got, tt.want)
			}
		})
	}
}
