lexer grammar langLexer;

tokens { INDENT, DEDENT }


IF: 'if';
FOR: 'for';
COLON: ':';
SEMICOLON: ';';
ASSIGN: ':=';
LET: 'let';
LPAR: '(';
LBRACK: '[';
RBRACK: ']';
RPAR: ')';
DOTDOT: '..';
INTEGER: [0-9]+;
REAL: [0-9]+[.][0-9]+;
IDENTIFIER: [_a-zA-Z]([_a-zA-Z0-9])*;
WhiteSpaces: [\t\u000B\u000C\u0020\u00A0]+ -> channel(HIDDEN);
LineTerminator: [\r\n\u2028\u2029] -> channel(HIDDEN);
