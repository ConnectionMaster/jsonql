// Code generated by gocc; DO NOT EDIT.

package parser

type (
	actionTable [numStates]actionRow
	actionRow   struct {
		canRecover bool
		actions    [numSymbols]action
	}
)

var actionTab = actionTable{
	actionRow{ // S0
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* * */
			nil,       /* / */
			nil,       /* % */
			nil,       /* ^ */
			shift(5),  /* - */
			shift(6),  /* ! */
			shift(10), /* null */
			shift(12), /* true */
			shift(13), /* false */
			shift(15), /* intLit */
			shift(16), /* floatLit */
			shift(18), /* doubleStringLit */
			shift(19), /* singleStringLit */
			shift(21), /* symbol */
			nil,       /* . */
			nil,       /* [ */
			nil,       /* ] */
		},
	},
	actionRow{ // S1
		canRecover: false,
		actions: [numSymbols]action{
			nil,          /* INVALID */
			accept(true), /* $ */
			shift(22),    /* * */
			shift(23),    /* / */
			shift(24),    /* % */
			nil,          /* ^ */
			nil,          /* - */
			nil,          /* ! */
			nil,          /* null */
			nil,          /* true */
			nil,          /* false */
			nil,          /* intLit */
			nil,          /* floatLit */
			nil,          /* doubleStringLit */
			nil,          /* singleStringLit */
			nil,          /* symbol */
			nil,          /* . */
			nil,          /* [ */
			nil,          /* ] */
		},
	},
	actionRow{ // S2
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(1), /* $, reduce: MulExpr */
			reduce(1), /* *, reduce: MulExpr */
			reduce(1), /* /, reduce: MulExpr */
			reduce(1), /* %, reduce: MulExpr */
			shift(25), /* ^ */
			nil,       /* - */
			nil,       /* ! */
			nil,       /* null */
			nil,       /* true */
			nil,       /* false */
			nil,       /* intLit */
			nil,       /* floatLit */
			nil,       /* doubleStringLit */
			nil,       /* singleStringLit */
			nil,       /* symbol */
			nil,       /* . */
			nil,       /* [ */
			nil,       /* ] */
		},
	},
	actionRow{ // S3
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(5), /* $, reduce: ExpExpr */
			reduce(5), /* *, reduce: ExpExpr */
			reduce(5), /* /, reduce: ExpExpr */
			reduce(5), /* %, reduce: ExpExpr */
			reduce(5), /* ^, reduce: ExpExpr */
			nil,       /* - */
			nil,       /* ! */
			nil,       /* null */
			nil,       /* true */
			nil,       /* false */
			nil,       /* intLit */
			nil,       /* floatLit */
			nil,       /* doubleStringLit */
			nil,       /* singleStringLit */
			nil,       /* symbol */
			nil,       /* . */
			nil,       /* [ */
			nil,       /* ] */
		},
	},
	actionRow{ // S4
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(7), /* $, reduce: UnaryExpr */
			reduce(7), /* *, reduce: UnaryExpr */
			reduce(7), /* /, reduce: UnaryExpr */
			reduce(7), /* %, reduce: UnaryExpr */
			reduce(7), /* ^, reduce: UnaryExpr */
			nil,       /* - */
			nil,       /* ! */
			nil,       /* null */
			nil,       /* true */
			nil,       /* false */
			nil,       /* intLit */
			nil,       /* floatLit */
			nil,       /* doubleStringLit */
			nil,       /* singleStringLit */
			nil,       /* symbol */
			nil,       /* . */
			nil,       /* [ */
			nil,       /* ] */
		},
	},
	actionRow{ // S5
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* * */
			nil,       /* / */
			nil,       /* % */
			nil,       /* ^ */
			nil,       /* - */
			nil,       /* ! */
			shift(10), /* null */
			shift(12), /* true */
			shift(13), /* false */
			shift(15), /* intLit */
			shift(16), /* floatLit */
			shift(18), /* doubleStringLit */
			shift(19), /* singleStringLit */
			shift(21), /* symbol */
			nil,       /* . */
			nil,       /* [ */
			nil,       /* ] */
		},
	},
	actionRow{ // S6
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* * */
			nil,       /* / */
			nil,       /* % */
			nil,       /* ^ */
			shift(5),  /* - */
			shift(6),  /* ! */
			shift(10), /* null */
			shift(12), /* true */
			shift(13), /* false */
			shift(15), /* intLit */
			shift(16), /* floatLit */
			shift(18), /* doubleStringLit */
			shift(19), /* singleStringLit */
			shift(21), /* symbol */
			nil,       /* . */
			nil,       /* [ */
			nil,       /* ] */
		},
	},
	actionRow{ // S7
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(10), /* $, reduce: Expr */
			reduce(10), /* *, reduce: Expr */
			reduce(10), /* /, reduce: Expr */
			reduce(10), /* %, reduce: Expr */
			reduce(10), /* ^, reduce: Expr */
			nil,        /* - */
			nil,        /* ! */
			nil,        /* null */
			nil,        /* true */
			nil,        /* false */
			nil,        /* intLit */
			nil,        /* floatLit */
			nil,        /* doubleStringLit */
			nil,        /* singleStringLit */
			nil,        /* symbol */
			nil,        /* . */
			nil,        /* [ */
			nil,        /* ] */
		},
	},
	actionRow{ // S8
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(11), /* $, reduce: Expr */
			reduce(11), /* *, reduce: Expr */
			reduce(11), /* /, reduce: Expr */
			reduce(11), /* %, reduce: Expr */
			reduce(11), /* ^, reduce: Expr */
			nil,        /* - */
			nil,        /* ! */
			nil,        /* null */
			nil,        /* true */
			nil,        /* false */
			nil,        /* intLit */
			nil,        /* floatLit */
			nil,        /* doubleStringLit */
			nil,        /* singleStringLit */
			nil,        /* symbol */
			shift(28),  /* . */
			shift(29),  /* [ */
			nil,        /* ] */
		},
	},
	actionRow{ // S9
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(19), /* $, reduce: Literal */
			reduce(19), /* *, reduce: Literal */
			reduce(19), /* /, reduce: Literal */
			reduce(19), /* %, reduce: Literal */
			reduce(19), /* ^, reduce: Literal */
			nil,        /* - */
			nil,        /* ! */
			nil,        /* null */
			nil,        /* true */
			nil,        /* false */
			nil,        /* intLit */
			nil,        /* floatLit */
			nil,        /* doubleStringLit */
			nil,        /* singleStringLit */
			nil,        /* symbol */
			nil,        /* . */
			nil,        /* [ */
			nil,        /* ] */
		},
	},
	actionRow{ // S10
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(12), /* $, reduce: NullLiteral */
			reduce(12), /* *, reduce: NullLiteral */
			reduce(12), /* /, reduce: NullLiteral */
			reduce(12), /* %, reduce: NullLiteral */
			reduce(12), /* ^, reduce: NullLiteral */
			nil,        /* - */
			nil,        /* ! */
			nil,        /* null */
			nil,        /* true */
			nil,        /* false */
			nil,        /* intLit */
			nil,        /* floatLit */
			nil,        /* doubleStringLit */
			nil,        /* singleStringLit */
			nil,        /* symbol */
			nil,        /* . */
			nil,        /* [ */
			nil,        /* ] */
		},
	},
	actionRow{ // S11
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(20), /* $, reduce: Literal */
			reduce(20), /* *, reduce: Literal */
			reduce(20), /* /, reduce: Literal */
			reduce(20), /* %, reduce: Literal */
			reduce(20), /* ^, reduce: Literal */
			nil,        /* - */
			nil,        /* ! */
			nil,        /* null */
			nil,        /* true */
			nil,        /* false */
			nil,        /* intLit */
			nil,        /* floatLit */
			nil,        /* doubleStringLit */
			nil,        /* singleStringLit */
			nil,        /* symbol */
			nil,        /* . */
			nil,        /* [ */
			nil,        /* ] */
		},
	},
	actionRow{ // S12
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(13), /* $, reduce: BooleanLiteral */
			reduce(13), /* *, reduce: BooleanLiteral */
			reduce(13), /* /, reduce: BooleanLiteral */
			reduce(13), /* %, reduce: BooleanLiteral */
			reduce(13), /* ^, reduce: BooleanLiteral */
			nil,        /* - */
			nil,        /* ! */
			nil,        /* null */
			nil,        /* true */
			nil,        /* false */
			nil,        /* intLit */
			nil,        /* floatLit */
			nil,        /* doubleStringLit */
			nil,        /* singleStringLit */
			nil,        /* symbol */
			nil,        /* . */
			nil,        /* [ */
			nil,        /* ] */
		},
	},
	actionRow{ // S13
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(14), /* $, reduce: BooleanLiteral */
			reduce(14), /* *, reduce: BooleanLiteral */
			reduce(14), /* /, reduce: BooleanLiteral */
			reduce(14), /* %, reduce: BooleanLiteral */
			reduce(14), /* ^, reduce: BooleanLiteral */
			nil,        /* - */
			nil,        /* ! */
			nil,        /* null */
			nil,        /* true */
			nil,        /* false */
			nil,        /* intLit */
			nil,        /* floatLit */
			nil,        /* doubleStringLit */
			nil,        /* singleStringLit */
			nil,        /* symbol */
			nil,        /* . */
			nil,        /* [ */
			nil,        /* ] */
		},
	},
	actionRow{ // S14
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(21), /* $, reduce: Literal */
			reduce(21), /* *, reduce: Literal */
			reduce(21), /* /, reduce: Literal */
			reduce(21), /* %, reduce: Literal */
			reduce(21), /* ^, reduce: Literal */
			nil,        /* - */
			nil,        /* ! */
			nil,        /* null */
			nil,        /* true */
			nil,        /* false */
			nil,        /* intLit */
			nil,        /* floatLit */
			nil,        /* doubleStringLit */
			nil,        /* singleStringLit */
			nil,        /* symbol */
			nil,        /* . */
			nil,        /* [ */
			nil,        /* ] */
		},
	},
	actionRow{ // S15
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(15), /* $, reduce: NumericLiteral */
			reduce(15), /* *, reduce: NumericLiteral */
			reduce(15), /* /, reduce: NumericLiteral */
			reduce(15), /* %, reduce: NumericLiteral */
			reduce(15), /* ^, reduce: NumericLiteral */
			nil,        /* - */
			nil,        /* ! */
			nil,        /* null */
			nil,        /* true */
			nil,        /* false */
			nil,        /* intLit */
			nil,        /* floatLit */
			nil,        /* doubleStringLit */
			nil,        /* singleStringLit */
			nil,        /* symbol */
			nil,        /* . */
			nil,        /* [ */
			nil,        /* ] */
		},
	},
	actionRow{ // S16
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(16), /* $, reduce: NumericLiteral */
			reduce(16), /* *, reduce: NumericLiteral */
			reduce(16), /* /, reduce: NumericLiteral */
			reduce(16), /* %, reduce: NumericLiteral */
			reduce(16), /* ^, reduce: NumericLiteral */
			nil,        /* - */
			nil,        /* ! */
			nil,        /* null */
			nil,        /* true */
			nil,        /* false */
			nil,        /* intLit */
			nil,        /* floatLit */
			nil,        /* doubleStringLit */
			nil,        /* singleStringLit */
			nil,        /* symbol */
			nil,        /* . */
			nil,        /* [ */
			nil,        /* ] */
		},
	},
	actionRow{ // S17
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(22), /* $, reduce: Literal */
			reduce(22), /* *, reduce: Literal */
			reduce(22), /* /, reduce: Literal */
			reduce(22), /* %, reduce: Literal */
			reduce(22), /* ^, reduce: Literal */
			nil,        /* - */
			nil,        /* ! */
			nil,        /* null */
			nil,        /* true */
			nil,        /* false */
			nil,        /* intLit */
			nil,        /* floatLit */
			nil,        /* doubleStringLit */
			nil,        /* singleStringLit */
			nil,        /* symbol */
			nil,        /* . */
			nil,        /* [ */
			nil,        /* ] */
		},
	},
	actionRow{ // S18
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(17), /* $, reduce: StringLiteral */
			reduce(17), /* *, reduce: StringLiteral */
			reduce(17), /* /, reduce: StringLiteral */
			reduce(17), /* %, reduce: StringLiteral */
			reduce(17), /* ^, reduce: StringLiteral */
			nil,        /* - */
			nil,        /* ! */
			nil,        /* null */
			nil,        /* true */
			nil,        /* false */
			nil,        /* intLit */
			nil,        /* floatLit */
			nil,        /* doubleStringLit */
			nil,        /* singleStringLit */
			nil,        /* symbol */
			nil,        /* . */
			nil,        /* [ */
			nil,        /* ] */
		},
	},
	actionRow{ // S19
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(18), /* $, reduce: StringLiteral */
			reduce(18), /* *, reduce: StringLiteral */
			reduce(18), /* /, reduce: StringLiteral */
			reduce(18), /* %, reduce: StringLiteral */
			reduce(18), /* ^, reduce: StringLiteral */
			nil,        /* - */
			nil,        /* ! */
			nil,        /* null */
			nil,        /* true */
			nil,        /* false */
			nil,        /* intLit */
			nil,        /* floatLit */
			nil,        /* doubleStringLit */
			nil,        /* singleStringLit */
			nil,        /* symbol */
			nil,        /* . */
			nil,        /* [ */
			nil,        /* ] */
		},
	},
	actionRow{ // S20
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(24), /* $, reduce: Identifier */
			reduce(24), /* *, reduce: Identifier */
			reduce(24), /* /, reduce: Identifier */
			reduce(24), /* %, reduce: Identifier */
			reduce(24), /* ^, reduce: Identifier */
			nil,        /* - */
			nil,        /* ! */
			nil,        /* null */
			nil,        /* true */
			nil,        /* false */
			nil,        /* intLit */
			nil,        /* floatLit */
			nil,        /* doubleStringLit */
			nil,        /* singleStringLit */
			nil,        /* symbol */
			reduce(24), /* ., reduce: Identifier */
			reduce(24), /* [, reduce: Identifier */
			nil,        /* ] */
		},
	},
	actionRow{ // S21
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(23), /* $, reduce: ObjectKey */
			reduce(23), /* *, reduce: ObjectKey */
			reduce(23), /* /, reduce: ObjectKey */
			reduce(23), /* %, reduce: ObjectKey */
			reduce(23), /* ^, reduce: ObjectKey */
			nil,        /* - */
			nil,        /* ! */
			nil,        /* null */
			nil,        /* true */
			nil,        /* false */
			nil,        /* intLit */
			nil,        /* floatLit */
			nil,        /* doubleStringLit */
			nil,        /* singleStringLit */
			nil,        /* symbol */
			reduce(23), /* ., reduce: ObjectKey */
			reduce(23), /* [, reduce: ObjectKey */
			nil,        /* ] */
		},
	},
	actionRow{ // S22
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* * */
			nil,       /* / */
			nil,       /* % */
			nil,       /* ^ */
			shift(5),  /* - */
			shift(6),  /* ! */
			shift(10), /* null */
			shift(12), /* true */
			shift(13), /* false */
			shift(15), /* intLit */
			shift(16), /* floatLit */
			shift(18), /* doubleStringLit */
			shift(19), /* singleStringLit */
			shift(21), /* symbol */
			nil,       /* . */
			nil,       /* [ */
			nil,       /* ] */
		},
	},
	actionRow{ // S23
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* * */
			nil,       /* / */
			nil,       /* % */
			nil,       /* ^ */
			shift(5),  /* - */
			shift(6),  /* ! */
			shift(10), /* null */
			shift(12), /* true */
			shift(13), /* false */
			shift(15), /* intLit */
			shift(16), /* floatLit */
			shift(18), /* doubleStringLit */
			shift(19), /* singleStringLit */
			shift(21), /* symbol */
			nil,       /* . */
			nil,       /* [ */
			nil,       /* ] */
		},
	},
	actionRow{ // S24
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* * */
			nil,       /* / */
			nil,       /* % */
			nil,       /* ^ */
			shift(5),  /* - */
			shift(6),  /* ! */
			shift(10), /* null */
			shift(12), /* true */
			shift(13), /* false */
			shift(15), /* intLit */
			shift(16), /* floatLit */
			shift(18), /* doubleStringLit */
			shift(19), /* singleStringLit */
			shift(21), /* symbol */
			nil,       /* . */
			nil,       /* [ */
			nil,       /* ] */
		},
	},
	actionRow{ // S25
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* * */
			nil,       /* / */
			nil,       /* % */
			nil,       /* ^ */
			shift(5),  /* - */
			shift(6),  /* ! */
			shift(10), /* null */
			shift(12), /* true */
			shift(13), /* false */
			shift(15), /* intLit */
			shift(16), /* floatLit */
			shift(18), /* doubleStringLit */
			shift(19), /* singleStringLit */
			shift(21), /* symbol */
			nil,       /* . */
			nil,       /* [ */
			nil,       /* ] */
		},
	},
	actionRow{ // S26
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(8), /* $, reduce: UnaryExpr */
			reduce(8), /* *, reduce: UnaryExpr */
			reduce(8), /* /, reduce: UnaryExpr */
			reduce(8), /* %, reduce: UnaryExpr */
			reduce(8), /* ^, reduce: UnaryExpr */
			nil,       /* - */
			nil,       /* ! */
			nil,       /* null */
			nil,       /* true */
			nil,       /* false */
			nil,       /* intLit */
			nil,       /* floatLit */
			nil,       /* doubleStringLit */
			nil,       /* singleStringLit */
			nil,       /* symbol */
			nil,       /* . */
			nil,       /* [ */
			nil,       /* ] */
		},
	},
	actionRow{ // S27
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(9), /* $, reduce: UnaryExpr */
			reduce(9), /* *, reduce: UnaryExpr */
			reduce(9), /* /, reduce: UnaryExpr */
			reduce(9), /* %, reduce: UnaryExpr */
			reduce(9), /* ^, reduce: UnaryExpr */
			nil,       /* - */
			nil,       /* ! */
			nil,       /* null */
			nil,       /* true */
			nil,       /* false */
			nil,       /* intLit */
			nil,       /* floatLit */
			nil,       /* doubleStringLit */
			nil,       /* singleStringLit */
			nil,       /* symbol */
			nil,       /* . */
			nil,       /* [ */
			nil,       /* ] */
		},
	},
	actionRow{ // S28
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* * */
			nil,       /* / */
			nil,       /* % */
			nil,       /* ^ */
			nil,       /* - */
			nil,       /* ! */
			nil,       /* null */
			nil,       /* true */
			nil,       /* false */
			nil,       /* intLit */
			nil,       /* floatLit */
			nil,       /* doubleStringLit */
			nil,       /* singleStringLit */
			shift(21), /* symbol */
			nil,       /* . */
			nil,       /* [ */
			nil,       /* ] */
		},
	},
	actionRow{ // S29
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* * */
			nil,       /* / */
			nil,       /* % */
			nil,       /* ^ */
			nil,       /* - */
			nil,       /* ! */
			shift(39), /* null */
			shift(41), /* true */
			shift(42), /* false */
			shift(44), /* intLit */
			shift(45), /* floatLit */
			shift(47), /* doubleStringLit */
			shift(48), /* singleStringLit */
			shift(50), /* symbol */
			nil,       /* . */
			nil,       /* [ */
			nil,       /* ] */
		},
	},
	actionRow{ // S30
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(2), /* $, reduce: MulExpr */
			reduce(2), /* *, reduce: MulExpr */
			reduce(2), /* /, reduce: MulExpr */
			reduce(2), /* %, reduce: MulExpr */
			shift(25), /* ^ */
			nil,       /* - */
			nil,       /* ! */
			nil,       /* null */
			nil,       /* true */
			nil,       /* false */
			nil,       /* intLit */
			nil,       /* floatLit */
			nil,       /* doubleStringLit */
			nil,       /* singleStringLit */
			nil,       /* symbol */
			nil,       /* . */
			nil,       /* [ */
			nil,       /* ] */
		},
	},
	actionRow{ // S31
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(3), /* $, reduce: MulExpr */
			reduce(3), /* *, reduce: MulExpr */
			reduce(3), /* /, reduce: MulExpr */
			reduce(3), /* %, reduce: MulExpr */
			shift(25), /* ^ */
			nil,       /* - */
			nil,       /* ! */
			nil,       /* null */
			nil,       /* true */
			nil,       /* false */
			nil,       /* intLit */
			nil,       /* floatLit */
			nil,       /* doubleStringLit */
			nil,       /* singleStringLit */
			nil,       /* symbol */
			nil,       /* . */
			nil,       /* [ */
			nil,       /* ] */
		},
	},
	actionRow{ // S32
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(4), /* $, reduce: MulExpr */
			reduce(4), /* *, reduce: MulExpr */
			reduce(4), /* /, reduce: MulExpr */
			reduce(4), /* %, reduce: MulExpr */
			shift(25), /* ^ */
			nil,       /* - */
			nil,       /* ! */
			nil,       /* null */
			nil,       /* true */
			nil,       /* false */
			nil,       /* intLit */
			nil,       /* floatLit */
			nil,       /* doubleStringLit */
			nil,       /* singleStringLit */
			nil,       /* symbol */
			nil,       /* . */
			nil,       /* [ */
			nil,       /* ] */
		},
	},
	actionRow{ // S33
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(6), /* $, reduce: ExpExpr */
			reduce(6), /* *, reduce: ExpExpr */
			reduce(6), /* /, reduce: ExpExpr */
			reduce(6), /* %, reduce: ExpExpr */
			reduce(6), /* ^, reduce: ExpExpr */
			nil,       /* - */
			nil,       /* ! */
			nil,       /* null */
			nil,       /* true */
			nil,       /* false */
			nil,       /* intLit */
			nil,       /* floatLit */
			nil,       /* doubleStringLit */
			nil,       /* singleStringLit */
			nil,       /* symbol */
			nil,       /* . */
			nil,       /* [ */
			nil,       /* ] */
		},
	},
	actionRow{ // S34
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(25), /* $, reduce: Identifier */
			reduce(25), /* *, reduce: Identifier */
			reduce(25), /* /, reduce: Identifier */
			reduce(25), /* %, reduce: Identifier */
			reduce(25), /* ^, reduce: Identifier */
			nil,        /* - */
			nil,        /* ! */
			nil,        /* null */
			nil,        /* true */
			nil,        /* false */
			nil,        /* intLit */
			nil,        /* floatLit */
			nil,        /* doubleStringLit */
			nil,        /* singleStringLit */
			nil,        /* symbol */
			reduce(25), /* ., reduce: Identifier */
			reduce(25), /* [, reduce: Identifier */
			nil,        /* ] */
		},
	},
	actionRow{ // S35
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* * */
			nil,       /* / */
			nil,       /* % */
			nil,       /* ^ */
			nil,       /* - */
			nil,       /* ! */
			nil,       /* null */
			nil,       /* true */
			nil,       /* false */
			nil,       /* intLit */
			nil,       /* floatLit */
			nil,       /* doubleStringLit */
			nil,       /* singleStringLit */
			nil,       /* symbol */
			nil,       /* . */
			nil,       /* [ */
			shift(51), /* ] */
		},
	},
	actionRow{ // S36
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* * */
			nil,        /* / */
			nil,        /* % */
			nil,        /* ^ */
			nil,        /* - */
			nil,        /* ! */
			nil,        /* null */
			nil,        /* true */
			nil,        /* false */
			nil,        /* intLit */
			nil,        /* floatLit */
			nil,        /* doubleStringLit */
			nil,        /* singleStringLit */
			nil,        /* symbol */
			nil,        /* . */
			nil,        /* [ */
			reduce(10), /* ], reduce: Expr */
		},
	},
	actionRow{ // S37
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* * */
			nil,        /* / */
			nil,        /* % */
			nil,        /* ^ */
			nil,        /* - */
			nil,        /* ! */
			nil,        /* null */
			nil,        /* true */
			nil,        /* false */
			nil,        /* intLit */
			nil,        /* floatLit */
			nil,        /* doubleStringLit */
			nil,        /* singleStringLit */
			nil,        /* symbol */
			shift(52),  /* . */
			shift(53),  /* [ */
			reduce(11), /* ], reduce: Expr */
		},
	},
	actionRow{ // S38
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* * */
			nil,        /* / */
			nil,        /* % */
			nil,        /* ^ */
			nil,        /* - */
			nil,        /* ! */
			nil,        /* null */
			nil,        /* true */
			nil,        /* false */
			nil,        /* intLit */
			nil,        /* floatLit */
			nil,        /* doubleStringLit */
			nil,        /* singleStringLit */
			nil,        /* symbol */
			nil,        /* . */
			nil,        /* [ */
			reduce(19), /* ], reduce: Literal */
		},
	},
	actionRow{ // S39
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* * */
			nil,        /* / */
			nil,        /* % */
			nil,        /* ^ */
			nil,        /* - */
			nil,        /* ! */
			nil,        /* null */
			nil,        /* true */
			nil,        /* false */
			nil,        /* intLit */
			nil,        /* floatLit */
			nil,        /* doubleStringLit */
			nil,        /* singleStringLit */
			nil,        /* symbol */
			nil,        /* . */
			nil,        /* [ */
			reduce(12), /* ], reduce: NullLiteral */
		},
	},
	actionRow{ // S40
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* * */
			nil,        /* / */
			nil,        /* % */
			nil,        /* ^ */
			nil,        /* - */
			nil,        /* ! */
			nil,        /* null */
			nil,        /* true */
			nil,        /* false */
			nil,        /* intLit */
			nil,        /* floatLit */
			nil,        /* doubleStringLit */
			nil,        /* singleStringLit */
			nil,        /* symbol */
			nil,        /* . */
			nil,        /* [ */
			reduce(20), /* ], reduce: Literal */
		},
	},
	actionRow{ // S41
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* * */
			nil,        /* / */
			nil,        /* % */
			nil,        /* ^ */
			nil,        /* - */
			nil,        /* ! */
			nil,        /* null */
			nil,        /* true */
			nil,        /* false */
			nil,        /* intLit */
			nil,        /* floatLit */
			nil,        /* doubleStringLit */
			nil,        /* singleStringLit */
			nil,        /* symbol */
			nil,        /* . */
			nil,        /* [ */
			reduce(13), /* ], reduce: BooleanLiteral */
		},
	},
	actionRow{ // S42
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* * */
			nil,        /* / */
			nil,        /* % */
			nil,        /* ^ */
			nil,        /* - */
			nil,        /* ! */
			nil,        /* null */
			nil,        /* true */
			nil,        /* false */
			nil,        /* intLit */
			nil,        /* floatLit */
			nil,        /* doubleStringLit */
			nil,        /* singleStringLit */
			nil,        /* symbol */
			nil,        /* . */
			nil,        /* [ */
			reduce(14), /* ], reduce: BooleanLiteral */
		},
	},
	actionRow{ // S43
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* * */
			nil,        /* / */
			nil,        /* % */
			nil,        /* ^ */
			nil,        /* - */
			nil,        /* ! */
			nil,        /* null */
			nil,        /* true */
			nil,        /* false */
			nil,        /* intLit */
			nil,        /* floatLit */
			nil,        /* doubleStringLit */
			nil,        /* singleStringLit */
			nil,        /* symbol */
			nil,        /* . */
			nil,        /* [ */
			reduce(21), /* ], reduce: Literal */
		},
	},
	actionRow{ // S44
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* * */
			nil,        /* / */
			nil,        /* % */
			nil,        /* ^ */
			nil,        /* - */
			nil,        /* ! */
			nil,        /* null */
			nil,        /* true */
			nil,        /* false */
			nil,        /* intLit */
			nil,        /* floatLit */
			nil,        /* doubleStringLit */
			nil,        /* singleStringLit */
			nil,        /* symbol */
			nil,        /* . */
			nil,        /* [ */
			reduce(15), /* ], reduce: NumericLiteral */
		},
	},
	actionRow{ // S45
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* * */
			nil,        /* / */
			nil,        /* % */
			nil,        /* ^ */
			nil,        /* - */
			nil,        /* ! */
			nil,        /* null */
			nil,        /* true */
			nil,        /* false */
			nil,        /* intLit */
			nil,        /* floatLit */
			nil,        /* doubleStringLit */
			nil,        /* singleStringLit */
			nil,        /* symbol */
			nil,        /* . */
			nil,        /* [ */
			reduce(16), /* ], reduce: NumericLiteral */
		},
	},
	actionRow{ // S46
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* * */
			nil,        /* / */
			nil,        /* % */
			nil,        /* ^ */
			nil,        /* - */
			nil,        /* ! */
			nil,        /* null */
			nil,        /* true */
			nil,        /* false */
			nil,        /* intLit */
			nil,        /* floatLit */
			nil,        /* doubleStringLit */
			nil,        /* singleStringLit */
			nil,        /* symbol */
			nil,        /* . */
			nil,        /* [ */
			reduce(22), /* ], reduce: Literal */
		},
	},
	actionRow{ // S47
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* * */
			nil,        /* / */
			nil,        /* % */
			nil,        /* ^ */
			nil,        /* - */
			nil,        /* ! */
			nil,        /* null */
			nil,        /* true */
			nil,        /* false */
			nil,        /* intLit */
			nil,        /* floatLit */
			nil,        /* doubleStringLit */
			nil,        /* singleStringLit */
			nil,        /* symbol */
			nil,        /* . */
			nil,        /* [ */
			reduce(17), /* ], reduce: StringLiteral */
		},
	},
	actionRow{ // S48
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* * */
			nil,        /* / */
			nil,        /* % */
			nil,        /* ^ */
			nil,        /* - */
			nil,        /* ! */
			nil,        /* null */
			nil,        /* true */
			nil,        /* false */
			nil,        /* intLit */
			nil,        /* floatLit */
			nil,        /* doubleStringLit */
			nil,        /* singleStringLit */
			nil,        /* symbol */
			nil,        /* . */
			nil,        /* [ */
			reduce(18), /* ], reduce: StringLiteral */
		},
	},
	actionRow{ // S49
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* * */
			nil,        /* / */
			nil,        /* % */
			nil,        /* ^ */
			nil,        /* - */
			nil,        /* ! */
			nil,        /* null */
			nil,        /* true */
			nil,        /* false */
			nil,        /* intLit */
			nil,        /* floatLit */
			nil,        /* doubleStringLit */
			nil,        /* singleStringLit */
			nil,        /* symbol */
			reduce(24), /* ., reduce: Identifier */
			reduce(24), /* [, reduce: Identifier */
			reduce(24), /* ], reduce: Identifier */
		},
	},
	actionRow{ // S50
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* * */
			nil,        /* / */
			nil,        /* % */
			nil,        /* ^ */
			nil,        /* - */
			nil,        /* ! */
			nil,        /* null */
			nil,        /* true */
			nil,        /* false */
			nil,        /* intLit */
			nil,        /* floatLit */
			nil,        /* doubleStringLit */
			nil,        /* singleStringLit */
			nil,        /* symbol */
			reduce(23), /* ., reduce: ObjectKey */
			reduce(23), /* [, reduce: ObjectKey */
			reduce(23), /* ], reduce: ObjectKey */
		},
	},
	actionRow{ // S51
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			reduce(26), /* $, reduce: Identifier */
			reduce(26), /* *, reduce: Identifier */
			reduce(26), /* /, reduce: Identifier */
			reduce(26), /* %, reduce: Identifier */
			reduce(26), /* ^, reduce: Identifier */
			nil,        /* - */
			nil,        /* ! */
			nil,        /* null */
			nil,        /* true */
			nil,        /* false */
			nil,        /* intLit */
			nil,        /* floatLit */
			nil,        /* doubleStringLit */
			nil,        /* singleStringLit */
			nil,        /* symbol */
			reduce(26), /* ., reduce: Identifier */
			reduce(26), /* [, reduce: Identifier */
			nil,        /* ] */
		},
	},
	actionRow{ // S52
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* * */
			nil,       /* / */
			nil,       /* % */
			nil,       /* ^ */
			nil,       /* - */
			nil,       /* ! */
			nil,       /* null */
			nil,       /* true */
			nil,       /* false */
			nil,       /* intLit */
			nil,       /* floatLit */
			nil,       /* doubleStringLit */
			nil,       /* singleStringLit */
			shift(50), /* symbol */
			nil,       /* . */
			nil,       /* [ */
			nil,       /* ] */
		},
	},
	actionRow{ // S53
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* * */
			nil,       /* / */
			nil,       /* % */
			nil,       /* ^ */
			nil,       /* - */
			nil,       /* ! */
			shift(39), /* null */
			shift(41), /* true */
			shift(42), /* false */
			shift(44), /* intLit */
			shift(45), /* floatLit */
			shift(47), /* doubleStringLit */
			shift(48), /* singleStringLit */
			shift(50), /* symbol */
			nil,       /* . */
			nil,       /* [ */
			nil,       /* ] */
		},
	},
	actionRow{ // S54
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* * */
			nil,        /* / */
			nil,        /* % */
			nil,        /* ^ */
			nil,        /* - */
			nil,        /* ! */
			nil,        /* null */
			nil,        /* true */
			nil,        /* false */
			nil,        /* intLit */
			nil,        /* floatLit */
			nil,        /* doubleStringLit */
			nil,        /* singleStringLit */
			nil,        /* symbol */
			reduce(25), /* ., reduce: Identifier */
			reduce(25), /* [, reduce: Identifier */
			reduce(25), /* ], reduce: Identifier */
		},
	},
	actionRow{ // S55
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* * */
			nil,       /* / */
			nil,       /* % */
			nil,       /* ^ */
			nil,       /* - */
			nil,       /* ! */
			nil,       /* null */
			nil,       /* true */
			nil,       /* false */
			nil,       /* intLit */
			nil,       /* floatLit */
			nil,       /* doubleStringLit */
			nil,       /* singleStringLit */
			nil,       /* symbol */
			nil,       /* . */
			nil,       /* [ */
			shift(56), /* ] */
		},
	},
	actionRow{ // S56
		canRecover: false,
		actions: [numSymbols]action{
			nil,        /* INVALID */
			nil,        /* $ */
			nil,        /* * */
			nil,        /* / */
			nil,        /* % */
			nil,        /* ^ */
			nil,        /* - */
			nil,        /* ! */
			nil,        /* null */
			nil,        /* true */
			nil,        /* false */
			nil,        /* intLit */
			nil,        /* floatLit */
			nil,        /* doubleStringLit */
			nil,        /* singleStringLit */
			nil,        /* symbol */
			reduce(26), /* ., reduce: Identifier */
			reduce(26), /* [, reduce: Identifier */
			reduce(26), /* ], reduce: Identifier */
		},
	},
}
