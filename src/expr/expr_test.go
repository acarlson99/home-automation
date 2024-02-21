package expr

import (
	"reflect"
	"strconv"
	"strings"
	"testing"

	hpb "github.com/acarlson99/home-automation/proto/go"
)

func TestEvalVar(t *testing.T) {
	type args struct {
		expr *hpb.Var
	}
	tests := []struct {
		name    string
		args    args
		want    *hpb.Primitive
		wantErr bool
	}{
		{
			name: "integer evals to integer",
			args: args{
				expr: &hpb.Var{
					V: &hpb.Var_Prim{
						Prim: &hpb.Primitive{
							V: &hpb.Primitive_Int32{
								Int32: 32,
							},
						},
					},
				},
			},
			want: &hpb.Primitive{
				V: &hpb.Primitive_Int32{
					Int32: 32,
				},
			},
		},
		{
			name: "3 + 10 = 13",
			args: args{
				expr: &hpb.Var{
					V: &hpb.Var_BinOp{
						BinOp: &hpb.BinaryOperator{
							Lhs: &hpb.Var{
								V: &hpb.Var_Prim{
									Prim: &hpb.Primitive{
										V: &hpb.Primitive_Int32{
											Int32: 3,
										},
									},
								},
							},
							Rhs: &hpb.Var{
								V: &hpb.Var_Prim{
									Prim: &hpb.Primitive{
										V: &hpb.Primitive_Int32{
											Int32: 3,
										},
									},
								},
							},
						},
					},
				},
			},
			want: &hpb.Primitive{
				V: &hpb.Primitive_Int32{
					Int32: 6,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EvalVar(tt.args.expr)
			if (err != nil) != tt.wantErr {
				t.Errorf("EvalVar() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EvalVar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cmp(t *testing.T) {
	type args struct {
		op  *hpb.Comparison_Operator
		lhs int
		rhs int
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "9==9=true",
			args: args{
				op:  hpb.Comparison_EQ.Enum(),
				lhs: 9,
				rhs: 9,
			},
			want: true,
		},
		{
			name: "9>9=false",
			args: args{
				op:  hpb.Comparison_GT.Enum(),
				lhs: 9,
				rhs: 9,
			},
			want: false,
		},
		{
			name: "9<9=false",
			args: args{
				op:  hpb.Comparison_LT.Enum(),
				lhs: 9,
				rhs: 9,
			},
			want: false,
		},
		{
			name: "9>=9=true",
			args: args{
				op:  hpb.Comparison_GTEQ.Enum(),
				lhs: 9,
				rhs: 9,
			},
			want: true,
		},
		{
			name: "9<=9=true",
			args: args{
				op:  hpb.Comparison_LTEQ.Enum(),
				lhs: 9,
				rhs: 9,
			},
			want: true,
		},
		{
			name: "9==10=false",
			args: args{
				op:  hpb.Comparison_EQ.Enum(),
				lhs: 9,
				rhs: 10,
			},
			want: false,
		},
		{
			name: "9>10=false",
			args: args{
				op:  hpb.Comparison_GT.Enum(),
				lhs: 9,
				rhs: 10,
			},
			want: false,
		},
		{
			name: "9<10=true",
			args: args{
				op:  hpb.Comparison_LT.Enum(),
				lhs: 9,
				rhs: 10,
			},
			want: true,
		},
		{
			name: "9>=10=false",
			args: args{
				op:  hpb.Comparison_GTEQ.Enum(),
				lhs: 9,
				rhs: 10,
			},
			want: false,
		},
		{
			name: "9<=10=true",
			args: args{
				op:  hpb.Comparison_LTEQ.Enum(),
				lhs: 9,
				rhs: 10,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := cmp(tt.args.op, tt.args.lhs, tt.args.rhs)
			if (err != nil) != tt.wantErr {
				t.Errorf("cmp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("cmp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func varPrim(p *hpb.Primitive) *hpb.Var {
	return &hpb.Var{V: &hpb.Var_Prim{Prim: p}}
}

func TestEvalComparison(t *testing.T) {
	type args struct {
		expr *hpb.Comparison
	}
	type TestCase struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}
	parseStrToComp := func(s string, toStr bool) TestCase {
		// s := "9 == 9 true"
		ss := strings.Split(s, " ")
		op := hpb.Comparison_EQ
		switch ss[1] {
		case "==":
			op = hpb.Comparison_EQ
		case ">":
			op = hpb.Comparison_GT
		case ">=":
			op = hpb.Comparison_GTEQ
		case "<":
			op = hpb.Comparison_LT
		case "<=":
			op = hpb.Comparison_LTEQ
		}
		if !toStr {
			lhs, err := strconv.ParseInt(ss[0], 10, 32)
			if err != nil {
				t.Fatal(err)
			}
			rhs, err := strconv.ParseInt(ss[2], 10, 32)
			if err != nil {
				t.Fatal(err)
			}
			return TestCase{
				name: s,
				args: args{
					expr: &hpb.Comparison{
						Op:  op,
						Lhs: varPrim(&hpb.Primitive{V: &hpb.Primitive_Int32{Int32: int32(lhs)}}),
						Rhs: varPrim(&hpb.Primitive{V: &hpb.Primitive_Int32{Int32: int32(rhs)}}),
					},
				},
				want: ss[3][0] == 't',
			}
		} else {
			return TestCase{
				name: s,
				args: args{
					expr: &hpb.Comparison{
						Op:  op,
						Lhs: varPrim(&hpb.Primitive{V: &hpb.Primitive_String_{String_: ss[0]}}),
						Rhs: varPrim(&hpb.Primitive{V: &hpb.Primitive_String_{String_: ss[2]}}),
					},
				},
				want: ss[3][0] == 't',
			}
		}
	}

	tests := []TestCase{
		parseStrToComp("9 == 9 true", false),
		parseStrToComp("9 > 9 false", false),
		parseStrToComp("9 < 9 false", false),
		parseStrToComp("9 >= 9 true", false),
		parseStrToComp("9 <= 9 true", false),
		parseStrToComp("9 == 10 false", false),
		parseStrToComp("9 > 10 false", false),
		parseStrToComp("9 < 10 true", false),
		parseStrToComp("9 >= 10 false", false),
		parseStrToComp("9 <= 10 true", false),
		parseStrToComp("abc < def true", true),
		parseStrToComp("abc == def false", true),
		parseStrToComp("abc > def false", true),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EvalComparison(tt.args.expr)
			if (err != nil) != tt.wantErr {
				t.Errorf("EvalComparison() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("EvalComparison() = %v, want %v", got, tt.want)
			}
		})
	}
}

func varInt(n int) *hpb.Var {
	return varPrim(&hpb.Primitive{V: &hpb.Primitive_Int32{Int32: int32(n)}})
}

func varFloat(n float32) *hpb.Var {
	return varPrim(&hpb.Primitive{V: &hpb.Primitive_Float{Float: n}})
}

func varBool(b bool) *hpb.Var {
	return varPrim(&hpb.Primitive{V: &hpb.Primitive_Bool{Bool: b}})
}

func varString(s string) *hpb.Var {
	return varPrim(&hpb.Primitive{V: &hpb.Primitive_String_{String_: s}})
}

func TestEvalBinaryOperator(t *testing.T) {
	type args struct {
		expr *hpb.BinaryOperator
	}
	tests := []struct {
		name    string
		args    args
		want    *hpb.Primitive
		wantErr bool
	}{
		{
			name: "20+40",
			args: args{
				expr: &hpb.BinaryOperator{
					Op:  hpb.BinaryOperator_ADD,
					Lhs: varInt(20),
					Rhs: varInt(40),
				},
			},
			want: &hpb.Primitive{V: &hpb.Primitive_Int32{Int32: 60}},
		},
		{
			name: "10^11",
			args: args{
				expr: &hpb.BinaryOperator{
					Op:  hpb.BinaryOperator_XOR,
					Lhs: varInt(10),
					Rhs: varInt(11),
				},
			},
			want: &hpb.Primitive{V: &hpb.Primitive_Int32{Int32: 1}},
		},
		{
			name: "string append",
			args: args{
				expr: &hpb.BinaryOperator{
					Op:  hpb.BinaryOperator_ADD,
					Lhs: varString("abc"),
					Rhs: varString("def"),
				},
			},
			want: &hpb.Primitive{V: &hpb.Primitive_String_{String_: "abcdef"}},
		},
		{
			name: "xor strings should fail",
			args: args{
				expr: &hpb.BinaryOperator{
					Op:  hpb.BinaryOperator_XOR,
					Lhs: varString("abc"),
					Rhs: varString("def"),
				},
			},
			wantErr: true,
		},
		{
			name: "add string float should fail",
			args: args{
				expr: &hpb.BinaryOperator{
					Op:  hpb.BinaryOperator_ADD,
					Lhs: varString("abc"),
					Rhs: varInt(4),
				},
			},
			wantErr: true,
		},
		{
			name: "(10^12)&(6|9)",
			args: args{
				expr: &hpb.BinaryOperator{
					Op: hpb.BinaryOperator_AND,
					Lhs: &hpb.Var{
						V: &hpb.Var_BinOp{
							BinOp: &hpb.BinaryOperator{
								Op:  hpb.BinaryOperator_XOR,
								Lhs: varInt(10),
								Rhs: varInt(12),
							},
						},
					},
					Rhs: &hpb.Var{
						V: &hpb.Var_BinOp{
							BinOp: &hpb.BinaryOperator{
								Op:  hpb.BinaryOperator_OR,
								Lhs: varInt(6),
								Rhs: varInt(9),
							},
						},
					},
				},
			},
			want: &hpb.Primitive{V: &hpb.Primitive_Int32{Int32: 6}},
		},
		{
			name: "cannot xor floats",
			args: args{
				expr: &hpb.BinaryOperator{
					Op: hpb.BinaryOperator_XOR,
					Lhs: &hpb.Var{
						V: &hpb.Var_BinOp{
							BinOp: &hpb.BinaryOperator{
								Op:  hpb.BinaryOperator_ADD,
								Lhs: varFloat(10),
								Rhs: varFloat(12),
							},
						},
					},
					Rhs: &hpb.Var{
						V: &hpb.Var_BinOp{
							BinOp: &hpb.BinaryOperator{
								Op:  hpb.BinaryOperator_ADD,
								Lhs: varFloat(6),
								Rhs: varFloat(9),
							},
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "12.0 `div` 10.0 = 1.2",
			args: args{
				expr: &hpb.BinaryOperator{
					Op:  hpb.BinaryOperator_DIV,
					Lhs: varFloat(12),
					Rhs: varFloat(10),
				},
			},
			want: &hpb.Primitive{V: &hpb.Primitive_Float{Float: 12.0 / 10.0}},
		},
		{
			name: "12 `div` 10 = 1.2",
			args: args{
				expr: &hpb.BinaryOperator{
					Op:  hpb.BinaryOperator_DIV,
					Lhs: varInt(12),
					Rhs: varInt(10),
				},
			},
			want: &hpb.Primitive{V: &hpb.Primitive_Int32{Int32: 12 / 10}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EvalBinaryOperator(tt.args.expr)
			if (err != nil) != tt.wantErr {
				t.Errorf("EvalBinaryOperator() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EvalBinaryOperator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEvalUnaryOperator(t *testing.T) {
	type args struct {
		expr *hpb.UnaryOperator
	}
	tests := []struct {
		name    string
		args    args
		want    *hpb.Primitive
		wantErr bool
	}{
		{
			name: "cast negative to bool should be true",
			args: args{
				expr: &hpb.UnaryOperator{
					Op: hpb.UnaryOperator_CAST_BOOL,
					X:  varInt(-20),
				},
			},
			want: &hpb.Primitive{V: &hpb.Primitive_Bool{Bool: true}},
		},
		{
			name: "--20==20",
			args: args{
				expr: &hpb.UnaryOperator{
					Op: hpb.UnaryOperator_NEG,
					X:  varInt(-20),
				},
			},
			want: &hpb.Primitive{V: &hpb.Primitive_Int32{Int32: 20}},
		},
		{
			name: "cast bool to int should be 1",
			args: args{
				expr: &hpb.UnaryOperator{
					Op: hpb.UnaryOperator_CAST_INT,
					X:  varBool(true),
				},
			},
			want: &hpb.Primitive{V: &hpb.Primitive_Int32{Int32: 1}},
		},
		{
			name: "cast float to int",
			args: args{
				expr: &hpb.UnaryOperator{
					Op: hpb.UnaryOperator_CAST_INT,
					X:  varFloat(42.42),
				},
			},
			want: &hpb.Primitive{V: &hpb.Primitive_Int32{Int32: 42}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EvalUnaryOperator(tt.args.expr)
			if (err != nil) != tt.wantErr {
				t.Errorf("EvalUnaryOperator() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EvalUnaryOperator() = %v, want %v", got, tt.want)
			}
		})
	}
}
