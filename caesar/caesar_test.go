package caesar

import (
	"reflect"
	"testing"
)

func TestDecode(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "success",
			args: args{input: "THIS"},
			want: []string{
				"UIJT", "VJKU", "WKLV", "XLMW", "YMNX", "ZNOY", "AOPZ",
				"BPQA", "CQRB", "DRSC", "ESTD", "FTUE", "GUVF", "HVWG",
				"IWXH", "JXYI", "KYZJ", "LZAK", "MABL", "NBCM", "OCDN",
				"PDEO", "QEFP", "RFGQ", "SGHR", "THIS",
			},
		},
		{
			name: "success with punctuations",
			args: args{input: "THIS _-/\\.!?"},
			want: []string{
				"UIJT _-/\\.!?", "VJKU _-/\\.!?", "WKLV _-/\\.!?", "XLMW _-/\\.!?",
				"YMNX _-/\\.!?", "ZNOY _-/\\.!?", "AOPZ _-/\\.!?", "BPQA _-/\\.!?",
				"CQRB _-/\\.!?", "DRSC _-/\\.!?", "ESTD _-/\\.!?", "FTUE _-/\\.!?",
				"GUVF _-/\\.!?", "HVWG _-/\\.!?", "IWXH _-/\\.!?", "JXYI _-/\\.!?",
				"KYZJ _-/\\.!?", "LZAK _-/\\.!?", "MABL _-/\\.!?", "NBCM _-/\\.!?",
				"OCDN _-/\\.!?", "PDEO _-/\\.!?", "QEFP _-/\\.!?", "RFGQ _-/\\.!?",
				"SGHR _-/\\.!?", "THIS _-/\\.!?",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Decode(tt.args.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseMap(t *testing.T) {
	type args struct {
		m map[string]int
	}
	tests := []struct {
		name string
		args args
		want map[int]string
	}{
		{
			name: "success",
			args: args{
				m: map[string]int{
					"a": 1,
					"b": 3,
					"c": 6,
				},
			},
			want: map[int]string{
				1: "a",
				3: "b",
				6: "c",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseMap(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
